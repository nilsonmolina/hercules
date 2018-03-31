package main

import (
	"fmt"
	"os"
	"regexp"
	"errors"

	"github.com/fatih/color"
)

func main() {
	// confirm user provided input
	if len(os.Args) < 2 {
		showUsage()
	}
	// check for proper filename
	if !isAlpha(string(os.Args[1][0])) {
		handleError(errors.New("Improper project name - first character must be alphanumeric"))
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
   -lang string		Select the Programming Language of the project. (allowed: c, go, html5)
   -help		Show usage information

C Sub-Options:
   -libft bool		Get libft from personal GitHub and include in lib directory. (default: false)
   -author string	Define a custom author for the author file (default: nmolina)

Go Sub-Options:
   -src	bool		Put all go files in a src directory (default: false)`)
	os.Exit(0)
}

var err error
var isAlpha = regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString

func handleError(err error) {
	if err != nil {
		color.Red("error: %v", err)
		os.Exit(1)
	}
}

func handleErrorMessage(err error, message string) {
	if err != nil {
		color.Red("error: %v", message)
		os.Exit(1)
	}
}

func handleUsageError(err error) {
	if err != nil {
		showUsage()
	}
}
