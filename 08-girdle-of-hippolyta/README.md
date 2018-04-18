# Labour 08: Girdle of Hippolyta
### **WORK IN PROGESS** - *(code does not work yet)*
**Goal**  
For this labour you write a program to send mail via SMTP. Learn about how email works and and can be used in a program. You must use a langauge other than BASH.

**Mandatory**  
Send an email out to warriors to join your army! As mentioned before, write a program that uses SMTP to send email.

## **Commands to Run**
**_*Note:_** *Still not done writing the code!*  

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