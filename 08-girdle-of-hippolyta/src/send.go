package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"strings"
)

// Email :
type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

type hercSMTP struct {
	host    string
	port    string
	address string
	user    string
	pass    string
	email   Email
}

var flags = map[string]interface{}{
	"-to":   "42-hercules@mailinator.com",
	"-body": "This email was sent from herc-mail.",
	"-sub":  "Hello From Herc-Mail",
	"-html": false,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("assets/email.html"))
}

func send(e hercSMTP) {
	var msg string
	log.Println("Preparing email.")
	if flags["-html"] == true {
		msg = makeHTMLMsg(e)
	} else {
		msg = makeTextMsg(e)
	}
	log.Println("Sending to SMTP server...")
	handleError(smtp.SendMail(e.address,
		smtp.PlainAuth("", e.user, e.pass, e.host),
		e.email.From, []string{e.email.To}, []byte(msg)))
	log.Println("Email Sent!")
	if strings.Contains(e.email.To, "@mailinator.com") {
		fmt.Printf("See the email at https://www.mailinator.com and search up %v\n", e.email.To)
	} else {
		fmt.Printf("smtp email was successfully sent to %v\n", e.email.To)
	}
}

func makeTextMsg(e hercSMTP) string {
	return "From: " + e.email.From + "\n" +
		"To: " + e.email.To + "\n" +
		"Subject: " + e.email.Subject + "\n\n" +
		e.email.Body
}

func makeHTMLMsg(e hercSMTP) string {
	buffer := new(bytes.Buffer)
	handleError(tpl.Execute(buffer, e.email))
	MIME := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	return "To: " + e.email.From + "\r\nSubject: " + e.email.Subject + "\r\n" + MIME + "\r\n" + buffer.String()
}

func newSMTPEmail() hercSMTP {
	return hercSMTP{
		host:    "smtp.gmail.com",
		port:    "587",
		address: "smtp.gmail.com:587",
		user:    "42.nmolina@gmail.com",
		pass:    "42Hercules",
		email: Email{
			From:    "42.nmolina@gmail.com",
			To:      flags["-to"].(string),
			Subject: flags["-sub"].(string),
			Body:    flags["-body"].(string),
		},
	}
}
