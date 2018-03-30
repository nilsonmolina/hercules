package main

import (
	"bufio"
	"fmt"
	"os"
)

func startWizard() {
	fmt.Println(`
--------------------------------------------------------
                  Hercules Labour 05
                   stymphalian-bird
--------------------------------------------------------
Welcome to the fly-creator's automation tool wizard!
This wizard will walk you through the various options
to properly build your project structure. 

Tips:
- Leaving blank choices, will pick the default.
- At any point, you can exit the wizard by typing 'quit'.`)

	askLanguage()
	if project.language == "c" {
		askLibft()
		askAuthor()
	} else if project.language == "go" {
		askWeb()
	} else {
		os.Exit(1)
	}
	confirm()
}

func askLanguage() {
	fmt.Println(`
What Programming Language will your project be in?
1. C (default)
2. Go`)
	input := getInput()

	if input == "1" || input == "" {
		setCProject(os.Args[1])
	} else if input == "2" {
		setGoProject(os.Args[1])
	} else {
		askLanguage()
	}
}

func askLibft() {
	fmt.Println(`
Do you want your remote github libft to be included?
  - will be downloaded from http://github.com/nilsonmolina/libft.git
1. Yes
2. No (default)`)
	input := getInput()

	if input == "1" {
		project.flags["-libft"] = "true"
	} else if input == "2" || input == "" {
		project.flags["-libft"] = "false"
	} else {
		askLibft()
	}
}

func askAuthor() {
	fmt.Println(`
What author name would you like to use?
 - default: nmolina
	`)
	input := getInput()

	if input == "" {
		project.flags["-author"] = "nmolina"
	} else {
		project.flags["-author"] = input
	}
}

func askWeb() {
	fmt.Println(`
Will this be a golang web project?
1. Yes
2. No (default)`)
	input := getInput()

	if input == "1" {
		project.flags["web"] = "true"
	} else if input == "2" || input == "" {
		project.flags["web"] = "false"
	} else {
		askWeb()
	}
}

func confirm() {
	fmt.Println("The following are your settings:")
	fmt.Printf("\nproject:\t%v\n", project.name)
	fmt.Printf("language:\t%v\n", project.language)
	if project.language == "c" {
		fmt.Printf("libft:\t\t%v\n", project.flags["libft"])
		fmt.Printf("author:\t\t%v\n", project.flags["author"])
	} else if project.language == "go" {
		fmt.Printf("web:\t\t%v\n", project.flags["web"])
	}
	fmt.Println("\nDo you still want to proceed?\n1. Yes\n2. No")

	input := getInput()
	if input == "2" {
		os.Exit(0)
	} else if input != "1" {
		confirm()
	}
}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ln := scanner.Text()

	if ln == "quit" {
		os.Exit(0)
	}
	return ln
}
