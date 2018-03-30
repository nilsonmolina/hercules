package main

// c files
var cMain = `#include "main.h"

int main()
{
	write(1, "Hello World!\n", 13);
	return (0);
}`

var cMakefile = `NAME = main

# file names
SRC = main.c 
OBJ = $(SRC:.c=.o)

# directories
SRCDIR = src
OBJDIR = objs

# files with their paths
SRCS = $(addprefix $(SRCDIR)/, $(SRC))
OBJS = $(addprefix $(OBJDIR)/, $(OBJ))

# compiler
CC = gcc
CFLAGS = -c -Wall -Werror -Wextra 
LIBS = 
HEADERS = -I includes

# prevent name collisions with files in the directory.
.PHONY: all clean fclean re

all: $(NAME)

$(OBJDIR)/%.o: $(SRCDIR)/%.c
	@mkdir -p $(OBJDIR)
	@$(CC) $(CFLAGS) $(HEADERS) $< -o $@

$(NAME): $(OBJS)
	@$(CC) $(OBJS) $(LIBS) -o $@
	@echo "- $(NAME) built and ready"

clean:
	@/bin/rm -rf $(OBJDIR)
	@echo "- $(NAME) cleaned"

fclean: clean
	@/bin/rm -f $(NAME)
	@echo "- $(NAME) fcleaned"

re: fclean all
`
var cHeaderFile = `#ifndef MAIN_H
# define MAIN_H

// Libraries
# include <unistd.h>

// VARIABLES


// STRUCTS


// FUNCTIONS


#endif
`

// go files
var goMain = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
`
