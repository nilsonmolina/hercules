package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func startWizard() {
	// set clear commands for different OS's
	initClear()
	fmt.Println(`
--------------------------------------------------------
                  Hercules Labour 05
                   stymphalian-bird
--------------------------------------------------------
Welcome to the hercules automation tool wizard!
This wizard will walk you through the various options
to properly build your project structure. 

Tips:
- Leaving blank choices, will pick the default.
- At any point, you can exit the wizard by typing 'quit'.

Press 'enter' to continue`)
	// await for any input
	getInput()
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
	callClear()
}

func askLanguage() {
	callClear()
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
	callClear()
	fmt.Println(`
Do you want your remote github libft to be included?
  - will be downloaded from http://github.com/nilsonmolina/libft.git
1. Yes
2. No (default)`)
	input := getInput()

	if input == "1" {
		project.flags["-libft"] = true
	} else if input == "2" || input == "" {
		project.flags["-libft"] = false
	} else {
		askLibft()
	}
}

func askAuthor() {
	callClear()
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
	callClear()
	fmt.Println(`
Would you like your go files to be in a src directory?
1. Yes
2. No (default)`)
	input := getInput()

	if input == "1" {
		project.flags["-src"] = true
	} else if input == "2" || input == "" {
		project.flags["-src"] = false
	} else {
		askWeb()
	}
}

func confirm() {
	callClear()
	fmt.Println("The following are your settings:")
	fmt.Printf("\nproject:\n   - %v\n", project.name)
	fmt.Printf("language:\n   - %v\n", project.language)
	for k, v := range project.flags {
		fmt.Printf("%v:\n   - %v\n", k[1:], v)
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

// type inquiry struct {
// 	question string
// 	handler  func()
// }

// func ask(i inquiry) {
// 	callClear()
// 	fmt.Println(i.question)
// 	i.handler()
// }

var clear map[string]func()

func initClear() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {
		value()
	}
}
