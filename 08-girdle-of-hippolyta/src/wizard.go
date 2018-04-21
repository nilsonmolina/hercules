package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func startWizard(e *hercSMTP) {
	// set clear commands for different OS's
	initClear()
	fmt.Println(`
--------------------------------------------------------
                  Hercules Labour 08
                  girdle of hippolyta
--------------------------------------------------------
Welcome to the herc-mail wizard! This wizard will walk 
you through the various options to build your email. 

Tips:
- Leaving blank choices, will pick the default.
- At any point, you can exit the wizard by typing 'quit'.

Press 'enter' to continue`)
	// await for any input
	getInput()
	askFrom(e)
	askTo(e)
	askSubject(e)
	askBody(e)
	confirm(e)
	callClear()
}

func askFrom(e *hercSMTP) {
	callClear()
	fmt.Printf("* MUST BE GMAIL & MUST ALLOW UNSAFE APPS! *\n\n'FROM' what email would you like to send from?\n\t- default: %v\n\n", e.email.From)
	input := getInput()

	if input != "" {
		e.user = input
		e.email.From = input
		askPass(e)
	}
}

func askPass(e *hercSMTP) {
	fmt.Printf("PASSWORD: ")
	input := getInput()

	if input != "" {
		e.pass = input
	}
}

func askTo(e *hercSMTP) {
	callClear()
	fmt.Printf("'TO' what email would you like to send?\n\t- default: %v\n\n", e.email.To)
	input := getInput()

	if input != "" {
		e.email.To = input
	}
}

func askSubject(e *hercSMTP) {
	callClear()
	fmt.Printf("What is the 'SUBJECT' of your email?\n\t- default: %v\n\n", e.email.Subject)
	input := getInput()

	if input != "" {
		e.email.Subject = input
	}
}

func askBody(e *hercSMTP) {
	callClear()
	fmt.Printf("What is the 'BODY' of your email?\n\t- default: %v\n\n", e.email.Body)
	input := getInput()

	if input != "" {
		e.email.Body = input
	}
}

func confirm(e *hercSMTP) {
	callClear()
	fmt.Println("The following are your settings:")
	fmt.Printf("\nFROM:\n   - %v\n", e.email.From)
	fmt.Printf("\nTO:\n   - %v\n", e.email.To)
	fmt.Printf("\nSUBJECT:\n   - %v\n", e.email.Subject)
	fmt.Printf("\nBODY:\n   - %v\n", e.email.Body)
	fmt.Println("\nDo you still want to proceed?\n1. Yes\n2. No")

	input := getInput()
	if input == "2" {
		os.Exit(0)
	} else if input != "1" {
		confirm(e)
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
