package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	// confirm user provided input
	if len(os.Args) != 2 {
		showUsage()
	}
	// construct email
	e := defaultSMTPEmail()
	// send email
	send(e)
}

func showUsage() {
	fmt.Println(`Usage:	./go-herc-mail <OPTIONS>
        
Options:
-to     Change recipient 
            (Default: "42-hercules@mailinator.com")
-body   Change email body 
            (Default: "This email was sent from go-herc-mail.")
-sub    Change email subject 
            (Default: "Hello Inbox")
-html   Send email with HTML/CSS styling
            (Default: false)
-help   Show usage information`)
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
