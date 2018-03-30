package main

import (
	"fmt"
	"os"
)

func main() {
	// confirm user provided input
	if len(os.Args) < 2 {
		showUsage()
	}
	// get parameters from user
	if len(os.Args) == 2 {
		startWizard()
	} else {
		validateFlags()
	}
	// create project based on given parameters
	createProject()
}

func showUsage() {
	fmt.Println(`Usage:	./fly-creator PROJECT_NAME <OPTIONS> <SUB-OPTIONS>

OPTIONS and SUB-OPTIONS are NOT mandatory, but providing them 
will skip the automation wizard. 

Options:
   -lang string		Select the Programming Language of the project. (allowed: c, go)
   -help		Show usage information

C Sub-Options:
   -libft bool		Get libft from personal GitHub and include in lib directory. (default: false)
   -author string	Define a custom author for the author file (default: nmolina)

Go Sub-Options:
   -web	bool		Create a Go Web project (default: false)`)
	os.Exit(0)
}
