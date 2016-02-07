package main

import (
	"flag"
	"log"

	"github.com/twstrike/coyim/xmpp"
)

var user string
var password string
var recipient string
var message string

func init() {
	flag.StringVar(&user, "user", "", "JID of sender.")
	flag.StringVar(&password, "password", "", "Password of sender.")
	flag.StringVar(&recipient, "to", "", "JID of recipient.")
	flag.StringVar(&message, "msg", "", "Message to send.")
}

func main() {
	flag.Parse()

	if len(user) == 0 {
		log.Println("user is required")
		flag.Usage()
		return
	}

	if len(password) == 0 {
		log.Println("password is required")
		flag.Usage()
		return
	}

	if len(recipient) == 0 {
		log.Println("recipient is required")
		flag.Usage()
		return
	}

	if len(message) == 0 {
		log.Println("message is required")
		flag.Usage()
		return
	}

	log.Printf("Logging in with %s...", user)
	dialer := &xmpp.Dialer{
		JID:      user,
		Password: password,
	}
	conn, err := dialer.Dial()
	if err != nil {
		log.Fatalf("Error connecting: %s", err)
		return
	}
	defer conn.Close()

	log.Printf("Sending message to %s...", recipient)
	conn.Send(recipient, message)

	log.Printf("Closing connection (30s timeout).")
}
