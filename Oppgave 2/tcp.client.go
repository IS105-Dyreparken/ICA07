package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	// kobler til denne socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// leser inn input fra stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// sender til socket
		fmt.Fprintf(conn, text+"\n")
		// lytter til svar
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
