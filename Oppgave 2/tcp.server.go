package main

import "net"
import "fmt"
import "bufio"
import "strings" // Bare nødvendig nedenfor for prøvebehandling

func main() {

	fmt.Println("Launching server...")

	// Lytter på alle grensesnitt
	ln, _ := net.Listen("tcp", ":8081")

	// godtar tilkobling på port
	conn, _ := ln.Accept()

	// kjører loop "forver" (eller frem til vi trykker på ctrl-c)
	for {
		// vil lytte etter melding for å behandle slutt i newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// output melding er motatt
		newmessage := strings.ToUpper(message)
		// sender ny streng tilbake til klient
		conn.Write([]byte(newmessage + "\n"))
	}
}
