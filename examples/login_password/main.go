package main

import (
	"flag"
	"fmt"
	"github.com/timvosch/gospotlib/pkg/connection"
	"github.com/timvosch/gospotlib/pkg/pb"
	"log"
)

var (
	fUsername = flag.String("u", "", "Spotify username")
	fPassword = flag.String("p", "", "Spotify password")
)

func main() {
	flag.Parse()
	if *fUsername == "" || *fPassword == "" {
		fmt.Println("Usage ./login_password -u <username> -p <password>")
		return
	}
	credentials := pb.CredentialsFromUsernamePassword(*fUsername, *fPassword)

	log.Println("[Main] Logging in")
	con, err := connection.Connect(credentials, "device_name")
	if err != nil {
		log.Printf("[Main] Got error: %v\n", err)
		return
	}
	log.Printf("[Main] Login successful: %v\n", con)
}
