# Labour 09: Cattle of Geryon
**Goal**  
For this labour you will write a program like siege. Learn about how siege works and can be used to benchmark a server. The point of this program is to simulate placing a server "under siege." You can use any language.

**Mandatory**  
Send an email out to warriors to join your army! As mentioned before, write a program that uses SMTP to send email.

## **Commands to Run**   

**Send email using Wizard**
```
$ ./herc-mail
```
**Send email using defaults**
```
$ ./herc-mail -to 42.nmolina@gmail.com
```
**Send an HTML Email - _(open email using mail app)_**
```
$ ./herc-mail -html true -to 42.nmolina@gmail.com
```
**Receive email**
```
$ ./herc-mail -get
```
**Usage**
```
Usage:	./herc-mail <OPTIONS>

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
            (No Input) - gets latest email from 42.nmolina@gmail.com
```

## SMTP
SMTP(Simple Mail Transfer Protocol) is the standard for passing email messages from one mail server to another. 
There are two ports you’ll need to be aware of for SMTP:

- **Port 25:** This is the default SMTP port. It is not encrypted.
- **Port 465 / 587:** The default port for using SMTP through SSL.

An SMTP server understands very simple text commands like EHLO, MAIL, RCPT and DATA. The most common commands are:

- **HELO** - introduce yourself
- **EHLO** - introduce yourself and request extended mode  
- **MAIL FROM:** - specify the sender  
- **RCPT TO:** - specify the recipient  
- **DATA** - specify the body of the message (To, From and Subject should be the first three lines.)
- **RSET** - reset
- **QUIT** - quit the session
- **HELP** - get help on commands
- **VRFY** - verify an address
- **EXPN** - expand an address
- **VERB** - verbose