package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/timvosch/gospotlib/pkg/connection"
	"github.com/timvosch/gospotlib/pkg/discovery"
)

var (
	deviceName = "gospot"
)

func main() {
	log.Println("[Main] Starting...")
	defer log.Println("[Main] Stopping...")

	disc := discovery.New(deviceName)
	go disc.Start()
	defer disc.Shutdown()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sigC:
			return
		case creds := <-disc.Credentials():
			{
				log.Printf("[Main] Received credentials for: %s\n", *creds.Username)
				conn, err := connection.Connect(creds, disc.DeviceID())
				if err != nil {
					log.Printf("[Main] Failed to connect: %s\n", err)
					continue
				}
				log.Printf("[Main] Connected and authenticated succesfully!")
				// TODO: use authenticated connection to interact with spotify API's
				conn.Close()
			}
		}
	}
}
