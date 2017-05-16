package main

 import (
         "log"
         "net"
         "fmt"
 )

 func main() {
         hostName := "localhost"
         portNum := "6000"

         service := hostName + ":" + portNum

         RemoteAddr, err := net.ResolveUDPAddr("udp", service)



         conn, err := net.DialUDP("udp", nil, RemoteAddr)


         if err != nil {
                 log.Fatal(err)
         }

         log.Printf("Established connection to %s \n", service)
         log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
         log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

         defer conn.Close()


         message := []byte("Møte Fr 5.5 14:45 Flåklypa")

         _, err = conn.Write(message)

         if err != nil {
                 log.Println(err)
         }


         buffer := make([]byte, 1024)
         n, addr, err := conn.ReadFromUDP(buffer)

         fmt.Println("UDP Server : ", addr)
         fmt.Println("Received from UDP server : ", string(buffer[:n]))

 }
