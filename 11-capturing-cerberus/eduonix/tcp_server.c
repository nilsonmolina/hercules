/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   tcp_server.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nmolina <nmolina@student.42.us.org>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/05/11 23:01:24 by nmolina           #+#    #+#             */
/*   Updated: 2018/05/11 23:22:27 by nmolina          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>

int main()
{
	struct sockaddr_in  server_address;
	char                server_message[256] = "You have reached the server!";
	int                 server_socket;
	int					client_socket;

	// create the server socket
	server_socket = socket(AF_INET, SOCK_STREAM, 0);

	// define the server address
	server_address.sin_family = AF_INET;
	server_address.sin_addr.s_addr = INADDR_ANY;
	server_address.sin_port = htons(9002);

	// bind the socket to our specified IP and port
	bind(server_socket, (struct sockaddr *)&server_address, sizeof(server_address));

	// accept connections
	listen(server_socket, 5);
	client_socket = accept(server_socket, NULL, NULL);

	// send the message
	send(client_socket, server_message, sizeof(server_message), 0);

	// close the socket
	close(server_socket);

	return (0);
}