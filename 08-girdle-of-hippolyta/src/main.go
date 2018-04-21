package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	var e hercSMTP
	if len(os.Args) == 1 {
		e = newSMTPEmail()
		startWizard(&e)
	} else {
		validateFlags()
		e = newSMTPEmail()
	}
	// send email
	send(e)
}

func validateFlags() {
	if os.Args[1] == "-get" {
		receive()
		os.Exit(0)
	} else if len(os.Args)%2 == 0 {
		showUsage()
	}
	arg := 1
	for arg < len(os.Args) {
		flag := os.Args[arg]
		value := os.Args[arg+1]
		if v, ok := flags[flag]; !ok {
			showUsage()
		} else {
			switch v.(type) {
			case string:
				flags[flag] = value
			case bool:
				flags[flag], err = strconv.ParseBool(value)
				handleErrorMessage(err, fmt.Sprintf("The '%v' flag requires a value of type bool", flag))
			}
		}
		arg = arg + 2 // '+ 2' because flags come in {-flag value} pairs (ex. -src true)
	}
}

func showUsage() {
	fmt.Println(`Usage:	./herc-mail <OPTIONS>

Herc-Mail is a basic SMTP email client. All emails will be sent
from '42.nmolina@gmail.com' for the purposes of this project.
        
Options:
-to     Change recipient 
            (Default: "42-hercules@mailinator.com")
-body   Change email body 
            (Default: "This email was sent using herc-mail.")
-sub    Change email subject 
            (Default: "Hello From Herc-Mail")
-help   Show usage information
			(No Input)

*Experimental:
-html   Send email with HTML/CSS styling (allowed: true/false)
            (Default: false)
-get	Receive email using IMAP (MUST BE FIRST OPTION)
            (No Input) - gets latest email from 42.nmolina@gmail.com`)
	os.Exit(0)
}

var err error

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
