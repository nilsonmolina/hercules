/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   armegeddon.c                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nmolina <nmolina@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/04/14 22:16:46 by nmolina           #+#    #+#             */
/*   Updated: 2018/04/15 15:02:28 by nmolina          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>
#include <stdlib.h>

void	pooper(void)
{
	system("for ((i=0; 1; i++)); do echo dung > poo$i.crap; sleep 0.5; done");
}

void	loud_mouth(void)
{
	while (1)
		system("echo :D; sleep 0.05");
}

void	fart_face(void)
{
	while (1)
	{
		system("open -a Preview ./assets/death.jpg; sleep 0.85");
		system("pkill Preview");
	}
}

int		main(void)
{
	if (fork())
		pooper();
	else
		fork() ? fart_face() : loud_mouth();
	return (0);
}
