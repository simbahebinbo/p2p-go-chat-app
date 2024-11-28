package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

func main() {
	// Create a host that listens on all interfaces
	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), // Listen on all interfaces, random port
	)
	if err != nil {
		log.Fatalf("Failed to create libp2p host: %v", err)
	}
	defer host.Close()

	// Print the host details - this is what you need to share with the other peer
	fmt.Println("My peer ID:", host.ID())
	fmt.Println("My multiaddresses:")
	for _, addr := range host.Addrs() {
		fmt.Printf("\t%s/p2p/%s\n", addr, host.ID())
	}

	// Listen for incoming messages on a custom protocol (e.g., "/chat/1.0")
	host.SetStreamHandler("/chat/1.0", handleStream)

	// Start a goroutine for reading input from the user and sending messages
	go sendMessage(host)

	// Keep the program running to maintain the peer (indefinitely)
	select {}
}

// handleStream is the function that processes incoming messages from other peers.
func handleStream(stream network.Stream) {
	defer stream.Close()

	// Read the incoming message
	buf := make([]byte, 1024)
	n, err := stream.Read(buf)
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}

	// Clear the current line and print received message
	fmt.Printf("\r%s\rReceived message: %s\n", strings.Repeat(" ", 50), string(buf[:n]))
	fmt.Print("Enter message (or 'exit' to quit): ")
}

// sendMessage reads user input and sends it to a peer.
func sendMessage(host host.Host) {
	reader := bufio.NewReader(os.Stdin)
	var peerID peer.ID

	// First, establish connection
	for peerID == "" {
		fmt.Print("Enter peer multiaddr (or 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting...")
			return
		}

		// Parse the multiaddr
		addr, err := peer.AddrInfoFromString(input)
		if err != nil {
			fmt.Println("Invalid multiaddr:", err)
			continue
		}

		// Connect to the peer
		err = host.Connect(context.Background(), *addr)
		if err != nil {
			fmt.Printf("Failed to connect to peer: %v\n", err)
			continue
		}

		peerID = addr.ID
		fmt.Println("Connected successfully!")
	}

	// Now just loop for messages
	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			fmt.Println("Exiting...")
			return
		}

		// Open stream and send message
		stream, err := host.NewStream(context.Background(), peerID, "/chat/1.0")
		if err != nil {
			fmt.Printf("Failed to open stream: %v\n", err)
			continue
		}

		_, err = stream.Write([]byte(message))
		if err != nil {
			fmt.Printf("Failed to send message: %v\n", err)
		}
		stream.Close()
	}
}
