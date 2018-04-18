package main

import (
	"fmt"
	"net/smtp"
)

func send(e hercSMTP) {

	msg := "From: " + e.email.from + "\n" +
		"To: " + e.email.to + "\n" +
		"Subject: " + e.email.subject + "\n\n" +
		e.email.body

	handleError(smtp.SendMail(e.address,
		smtp.PlainAuth("", e.user, e.pass, e.host),
		e.email.from, []string{e.email.to}, []byte(msg)))

	fmt.Printf("smtp email sent! see it at https://www.mailinator.com/v2/inbox.jsp?zone=public&query=%v\n", e.email.to)
}

type email struct {
	from    string
	to      string
	subject string
	body    string
}

type hercSMTP struct {
	host    string
	port    string
	address string
	user    string
	pass    string
	email   email
}

func defaultSMTPEmail() hercSMTP {
	return hercSMTP{
		host:    "smtp.gmail.com",
		port:    "587",
		address: "smtp.gmail.com:587",
		user:    "42.nmolina@gmail.com",
		pass:    "42Hercules",
		email: email{
			from:    "42.nmolina@gmail.com",
			to:      "42-hercules@mailinator.com",
			subject: "Hello Inbox",
			body:    "This email was sent from go-herc-mail.",
		},
	}
}
