package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "8081"
)

func main() {
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Println("Connected to PokeBat Game Server")

	// Game loop
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(message)

		if strings.Contains(message, "Choose your starting Pokémon:") || strings.Contains(message, "Your turn!") || strings.Contains(message, "Choose a Pokémon to switch to:") {
			input := prompt("")
			conn.Write([]byte(input + "\n"))
		}

		if strings.Contains(message, "Congratulations! You won the battle.") || strings.Contains(message, "You lost the battle.") {
			break
		}
	}
}

func prompt(promptMsg string) string {
	fmt.Print(promptMsg)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}
