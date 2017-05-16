package main

 import (
         "fmt"
         "log"
         "net"

 )

 func handleUDPConnection(conn *net.UDPConn) {


         buffer := make([]byte, 1024)

         n, addr, err := conn.ReadFromUDP(buffer)

         fmt.Println("UDP client : ", addr)
         fmt.Println("Received from UDP client :  ", string(buffer[:n]))

         if err != nil {
                 log.Fatal(err)
         }


         message := []byte("Møte Fr 5.5 14:45 Flåklypa")
         _, err = conn.WriteToUDP(message, addr)

         if err != nil {
                 log.Println(err)
         }

 }

 func main() {
         hostName := "localhost"
         portNum := "6000"
         service := hostName + ":" + portNum

         udpAddr, err := net.ResolveUDPAddr("udp4", service)

         if err != nil {
                 log.Fatal(err)
         }


         ln, err := net.ListenUDP("udp", udpAddr)

         if err != nil {
                 log.Fatal(err)
         }

         fmt.Println("UDP server up and listening on port 6000")

         defer ln.Close()

         for {

                 handleUDPConnection(ln)
         }

 }
