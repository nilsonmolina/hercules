package main

import (
	"io/ioutil"
	"log"
	"net/mail"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func receive() {
	log.Println("Connecting to IMAP server...")
	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	handleError(err)
	log.Println("Connected")
	// Don't forget to logout
	defer c.Logout()
	// Login
	handleError(c.Login("42.nmolina@gmail.com", "42Hercules"))
	log.Println("Logged in")
	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	handleError(err)
	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(mbox.Messages, mbox.Messages)
	// Get the whole message body
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}
	messages := make(chan *imap.Message, 1)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()
	// Print Message Body
	log.Println("Last message:")
	msg := <-messages
	r := msg.GetBody(section)

	if r == nil {
		log.Fatal("Server did not return message body!")
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	log.Println("Date:", header.Get("Date"))
	log.Println("From:", header.Get("From"))
	log.Println("To:", header.Get("To"))
	log.Println("Subject:", header.Get("Subject"))

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Body:")
	log.Print("\t", string(body[:]))

	log.Println("Done!")
}
