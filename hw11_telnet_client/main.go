package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var timeout time.Duration

const DefaultTimeout = 10

func main() {
	flag.DurationVar(&timeout, "timeout", DefaultTimeout*time.Second, "Connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("need address and port")
	}

	host := args[0]
	port := args[1]
	socket := net.JoinHostPort(host, port)

	client := NewTelnetClient(socket, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatalf("error on connection: %s", err)
	}

	defer func(client TelnetClient) {
		err := client.Close()
		if err != nil {
			log.Fatalf("error on closing connection: %s", err)
		}
	}(client)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go sender(client, cancel)
	go reciver(client, cancel)

	<-ctx.Done()
	cancel()
}

func sender(client TelnetClient, cancel context.CancelFunc) {
	if err := client.Send(); err != nil {
		cancel()
	}
}

func reciver(client TelnetClient, cancel context.CancelFunc) {
	if err := client.Receive(); err != nil {
		cancel()
	}
}
