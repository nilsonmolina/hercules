package main

import (
	"fmt"
	"os"
	"strconv"
)

func validateOptions() {
	// check required -lang flag
	if os.Args[2] != "-lang" || len(os.Args) < 4 {
		showUsage()
	}
	if os.Args[3] != "c" && os.Args[3] != "go" {
		showUsage()
	}
	project.language = os.Args[3]
	// check remaining flags
	arg := 4
	var err error
	for arg < (len(os.Args)) { // the '+ 1' is to ensure we have a flag and value
		if os.Args[arg] == "-libft" {
			confirmLanguage(os.Args[arg], "c")
			project.libft, err = strconv.ParseBool(os.Args[arg+1])
			if err != nil {
				invalidValue(os.Args[arg], "boolean")
			}
		} else if os.Args[arg] == "-author" {
			confirmLanguage(os.Args[arg], "c")
			project.author = os.Args[arg+1]
		} else if os.Args[arg] == "-web" {
			confirmLanguage(os.Args[arg], "go")
			project.web, err = strconv.ParseBool(os.Args[arg+1])
			if err != nil {
				invalidValue(os.Args[arg], "boolean")
			}
		} else {
			showUsage()
		}
		arg = arg + 2
	}
}

func confirmLanguage(givenFlag string, requiredLanguage string) {
	if project.language != requiredLanguage {
		fmt.Printf("The '%v' flag only applies to the '%v' language.\n", givenFlag, requiredLanguage)
		os.Exit(1)
	}
}

func invalidValue(givenFlag string, expectedType string) {
	fmt.Printf("Error: expected a %v value for '%v' flag\n", expectedType, givenFlag)
	os.Exit(1)
}
