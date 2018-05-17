/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   proxy.c                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nmolina <nmolina@student.42.us.org>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/05/12 00:39:59 by nmolina           #+#    #+#             */
/*   Updated: 2018/05/17 10:27:10 by nmolina          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <errno.h>
#include <fcntl.h>

#define BIND_ADDR       INADDR_ANY
#define BIND_PORT		1337

#define SOCKS5_VERSION  0x05
#define SOCKS5_NO_AUTH  0x00
#define SOCKS5_TCP_CNN  0x01
#define SOCKS5_IPV4     0x01

static int	create_server_sfd(void)
{
	int					sfd;
	struct sockaddr_in	addr;

	/* Create a server socket */
	if ((sfd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
		printf("Error calling socket(): %s\n", strerror(errno));
		return -1;
	}
	/* Bind to an address to service requests */
	addr.sin_family = AF_INET;
	addr.sin_addr.s_addr = BIND_ADDR;
	addr.sin_port = htons(BIND_PORT);
	if (bind(sfd, (struct sockaddr *)&addr, sizeof (addr)) == -1) {
		printf("Error calling bind(): %s\n", strerror(errno));
		close(sfd);
		return -1;
	}
	/* Notify kernel how many pending requests can queue up */
	if (listen(sfd, 8) == -1) {
		printf("Error calling listen(): %s\n", strerror(errno));
		close(sfd);
		return -1;
	}
	return sfd;
}

static int	handle_client_greeting(int c_sfd)
{
	static const char success[] = { SOCKS5_VERSION, SOCKS5_NO_AUTH };
	unsigned char buff[64];
	int len;
	int i;
	int has_no_auth;

	len = recv(c_sfd, buff, sizeof(buff), 0);
	if (len == -1)
		return 0;
	/* Check the version of the client */
	if (buff[0] != SOCKS5_VERSION) {
		printf("Client didn't use SOCKS5\n");
		return 0;
	}
	/* Check the number of authentication methods */
	if (buff[1] == 0) {
		printf("Client has 0 authentication methods o_O\n");
		return 0;
	}
	/* Check that at least one of the supported auth methods is no auth */
	for (i = 1, has_no_auth = 0; i <= buff[1]; ++i) {
		if (buff[1 + i] == SOCKS5_NO_AUTH) {
			has_no_auth = 1;
			break;
		}
	}
	if (has_no_auth == 0) {
		printf("Client does not support no auth\n");
		return 0;
	}
	/* Send our succesful response to the client
	 * field1 - SOCKS VERSION
	 * field2 - chosen auth version
	 */
	send(c_sfd, success, sizeof(success), 0);
	return 1;
}

static int	handle_client_conn_req(int c_sfd, uint32_t *addr, uint16_t *port)
{
	char	buff[10];
	int		len;
	
	len = recv(c_sfd, buff, sizeof(buff), 0);
	if (len != 10) {
		printf("Incorrect client connection request length!\n");
		return 0;
	}
	if (buff[0] != SOCKS5_VERSION) {
		printf("Version not 5\n");
		return 0;
	}
	if (buff[1] != SOCKS5_TCP_CNN) {
		printf("We only support TCP connections\n");
		return 0;
	}
	if (buff[3] != SOCKS5_IPV4) {
		printf("We only support IPv4 - %04x\n", buff[3]);
		return 0;
	}
	*addr = *((uint32_t *)&buff[4]);
	/* You might also want to support domain name resolution but LZ */
	*port = *((uint16_t *)&buff[8]);
	return 1;
}

static int	proxy_traffic(int c_sfd, uint32_t addr, uint16_t port)
{
	struct	sockaddr_in saddr;
	int 	s_sfd;
	static const char connected[] = { 
		SOCKS5_VERSION, 
		0x00, 
		0x00, 
		0x01, 
		0x00,0x00,0x00,0x00, 
		0x00, 0x00 };

	printf("- socket\n");	
	if ((s_sfd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
		printf("Failed to create socket: %s\n", strerror(errno));
		return 0;
	}

	printf("- connecting\n");
	saddr.sin_family = AF_INET;
	saddr.sin_addr.s_addr = addr;
	saddr.sin_port = port;
	if (connect(s_sfd, (struct sockaddr *)&saddr, sizeof (saddr)) == -1) {
		printf("Failed to connect to target: %s\n", strerror(errno));
		close(s_sfd);
		return 0;
	}
	printf("We're connected to the remote side port %d!\n", ntohs(port));
	send(c_sfd, connected, sizeof(connected), 0);
	

	/* Make the socket nonblocking so we can read from both in and out
	 * fds at the same time without blocking */
	fcntl(c_sfd, F_SETFL, O_NONBLOCK | fcntl(c_sfd, F_GETFL, 0));
	fcntl(s_sfd, F_SETFL, O_NONBLOCK | fcntl(s_sfd, F_GETFL, 0));
	while (1) {
		char buff[1024];
		int len;

		/* Read traffic from the client. If there was data to read
		 * then len would be greater than 0. If len is -1 and errno is
		 * EAGAIN or EWOULDBLOCK it just means that there was no data
		 * to read at this time, and we should move on to trying to read
		 * from the server */	
		len = recv(c_sfd, buff, sizeof(buff), 0);
		//printf(" - reading from client - len: %d | c_sfd: %d\n", len, c_sfd);				
		if (len < 1) {
			if (errno != EAGAIN && errno != EWOULDBLOCK) {
				printf("Error reading from client: %s\n", strerror(errno));
				close(s_sfd);
				return 0;
			}
		} else {
			printf(" - sending to server - len: %d | c_sfd: %d\n", len, c_sfd);
			send(s_sfd, buff, len, 0);
		}

		// printf(" - reading from server\n");
		/* Read traffic from the server. IF there was data to read then
		 * len would be greater than 0 etc etc etc */
		len = recv(s_sfd, buff, sizeof(buff), 0);
		//printf(" - reading from server - len: %d | s_sfd: %d\n", len, s_sfd);	
		if (len < 1) {
			if (errno != EAGAIN && errno != EWOULDBLOCK) {
				printf("Error reading from server: %s\n", strerror(errno));
				close(s_sfd);
				return 0;
			}
		} else {
			printf(" - sending to client - len: %d | s_sfd: %d\n", len, s_sfd);
			send(c_sfd, buff, len, 0);
		}
		usleep(100);
	}
	return (0);
}

static void	handle_client(int c_sfd, const struct sockaddr_in *addr_client)
{
	uint32_t addr;
	uint16_t port;

	if (handle_client_greeting(c_sfd) == 0) {
		printf(" - error: handle_client_greeting == 0 \n");		
		close(c_sfd);
		return;
	}
	if (handle_client_conn_req(c_sfd, &addr, &port) == 0) {
		printf(" - error: handle_client_conn_req == 0 \n");
		close(c_sfd);
		return;
	}
	printf("- Handling proxy_traffic\n");
	if (proxy_traffic(c_sfd, addr, port) == 0) {
		printf(" - error: proxy_traffic == 0\n");		
		close(c_sfd);
		return;
	}

	printf("Finished yo\n");
	close(c_sfd);
}

int			main(void)
{
	int					sfd;
	int					c_sfd;
	struct sockaddr_in	addr;
	socklen_t			slen;

	if ((sfd = create_server_sfd()) == -1)
		return EXIT_FAILURE;

	while (1) {
		if ((c_sfd = accept(sfd, (struct sockaddr *)&addr, &slen)) == -1) {
            printf("Failed to accept client: %s\n", strerror(errno));
            close(sfd);
            return EXIT_FAILURE;
        }

        switch (fork()) {
            case -1:
                printf("Failed fork to handle client: %s\n", strerror(errno));
                return EXIT_FAILURE;
            case 0:
                printf("Client accepted! Handling client now\n");
                handle_client(c_sfd, &addr);
                exit(EXIT_SUCCESS);
            default:
                break;
        }
	}
}