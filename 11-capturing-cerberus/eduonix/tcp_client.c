/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tcp_client.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nmolina <nmolina@student.42.us.org>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/05/11 22:23:29 by nmolina           #+#    #+#             */
/*   Updated: 2018/05/12 13:12:30 by nmolina          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>

int     main()
{
	struct sockaddr_in  server_address;
	char                server_response[256];
	int                 network_socket;
	int                 conn_status;
	
	// create a socket
	network_socket = socket(AF_INET, SOCK_STREAM, 0);

	// specify an address for the socket
	server_address.sin_family = AF_INET;
	server_address.sin_addr.s_addr = INADDR_ANY;
	server_address.sin_port = htons(9002);

	// attempt connection
	conn_status = connect(network_socket, 
		(struct sockaddr *)&server_address, sizeof(server_address));
	if (conn_status == -1) {
		printf("error making a connection to the remote socket \n\n");
	}

	// recieve data from the server
	recv(network_socket, &server_response, sizeof(server_response), 0);
	printf("The server sent the data: \n - %s\n", server_response);

	return (0);
}