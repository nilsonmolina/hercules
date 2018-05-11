#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <arpa/inet.h>
#include <errno.h>
#include <fcntl.h>

#define BIND_ADDR  INADDR_ANY
#define BIND_PORT  1337

#define SOCKS5_VERSION  0x05
#define SOCKS5_NO_AUTH  0x00
#define SOCKS5_TCP_CNN  0x01
#define SOCKS5_IPV4     0x01

static int
create_server_fd(void)
{
    int fd;
    struct sockaddr_in addr;

    /* Create a server socket */
    if ((fd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
        printf("Error calling socket(): %s\n", strerror(errno));
        return -1;
    }

    /* Bind to an address to service requests */
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = INADDR_ANY;
    addr.sin_port = htons(BIND_PORT);
    if (bind(fd, (struct sockaddr *)&addr, sizeof (addr)) == -1) {
        printf("Error calling bind(): %s\n", strerror(errno));
        close(fd);
        return -1;
    }

    /* Notify kernel how many pending requests can queue up */
    if (listen(fd, 8) == -1) {
        printf("Error calling listen(): %s\n", strerror(errno));
        close(fd);
        return -1;
    }

    return fd;
}

static int
handle_client_greeting(int fd)
{
    static const char success[] = { SOCKS5_VERSION, SOCKS5_NO_AUTH };
    unsigned char buff[64];
    int len;
    int i;
    int has_no_auth;

    len = recv(fd, buff, sizeof (buff), 0);
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
    send(fd, success, sizeof(success), 0);

    return 1;
}

static int
handle_client_conn_req(int fd, uint32_t *addr, uint16_t *port)
{
    char buff[10];
    int len;
    
    len = recv(fd, buff, sizeof (buff), 0);
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
        printf("We only support IPv4\n");
        return 0;
    }
    *addr = *((uint32_t *)buff[4]);
    /* You might also want to support domain name resolution but LZ */

    *port = *((uint16_t *)buff[8]);

    return 0;
}

static int
proxy_traffic(int fd, uint32_t addr, uint16_t port)
{
    struct sockaddr_in daddr;
    int ofd;

    if ((ofd = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
        printf("Failed to create socket: %s\n", strerror(errno));
        return 0;
    }

    daddr.sin_family = AF_INET;
    daddr.sin_addr.s_addr = addr;
    daddr.sin_port = port;
    if (connect(ofd, (struct sockaddr *)&daddr, sizeof (daddr)) == -1) {
        printf("Failed to connect to target: %s\n", strerror(errno));
        close(ofd);
        return 0;
    }

    /* Make the socket nonblocking so we can read from both in and out
     * fds at the same time without blocking */
    fcntl(fd, F_SETFL, O_NONBLOCK | fcntl(fd, F_GETFL, 0));
    fcntl(ofd, F_SETFL, O_NONBLOCK | fcntl(ofd, F_GETFL, 0));
    while (1) {
        char buff[1024];
        int len;

        /* Read traffic from the client. If there was data to read
         * then len would be greater than 0. If len is -1 and errno is
         * EAGAIN or EWOULDBLOCK it just means that there was no data
         * to read at this time, and we should move on to trying to read
         * from the server */
        len = recv(fd, buff, sizeof (buff), 0);
        if (len == -1) {
            if (errno != EAGAIN && errno != EWOULDBLOCK) {
                printf("Error reading from client: %s\n", strerror(errno));
                close(ofd);
                return 0;
            }
        } else {
            send(ofd, buff, len, 0);
        }

        /* Read traffic from the server. IF there was data to read then
         * len would be greater than 0 etc etc etc */
        len = recv(ofd, buff, sizeof (buff), 0);
        if (len == -1) {
            if (errno != EAGAIN && errno != EWOULDBLOCK) {
                printf("Error reading from server: %s\n", strerror(errno));
                close(ofd);
                return 0;
            }
        } else {
            send(fd, buff, len, 0);
        }
    }
}

static void
handle_client(int fd, const struct sockaddr_in *addr)
{
    uint32_t addr;
    uint16_t port;

    if (handle_client_greeting(fd) == 0) {
        close(fd);
        return;
    }
    if (handle_client_conn_req(fd, &addr, &port) == 0) {
        close(fd);
        return;
    }
    if (proxy_traffic(fd, addr, port) == 0) {
        close(fd);
        return;
    }

    printf("Finished yo\n");
    close(fd);
}

int
main(void)
{
    int sfd;

    if ((sfd = create_server_fd()) == -1)
        return EXIT_FAILURE;

    while (1) {
        int fd;
        struct sockaddr_in addr;
        socklen_t slen;

        if ((fd = accept(sfd, (struct sockaddr *)&addr, &slen)) == -1) {
            printf("Failed to accept client: %s\n", strerror(errno));
            close(sfd);
            return EXIT_FAILURE;
        }

        switch (fork()) {
            case -1:
                printf("Failed to fork to handle client: %s\n", strerror(errno));
                return EXIT_FAILURE;
            case 0:
                handle_client(fd, &addr);
                exit(EXIT_SUCCESS);
            default:
                break;
        }
    }
}