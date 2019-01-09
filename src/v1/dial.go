package main

import (
        "fmt"
        "net"
        "os"
)

// We can use net.Dial() to send and recieve data, then return a Conn type interface once you connect to remote system.
// Dial() abstracted the network structure and it's transport. 
// We can use this public interface whether network protocol type is IPv4 or IPv6 and TCP or UDP. 

func main() {
        conn, err := net.Dial("tcp", "18.235.228.29:80")
        checkConnection(conn, err)
        //conn, err = net.Dial("udp", "ipv4:port")
        //checkConnection(conn, err)
        //conn, err = net.Dial("tcp", "ipv6:port")
        //checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
        if err != nil {
                fmt.Printf("error %v connecting!", err) 
                os.Exit(1)
        }
        fmt.Printf("Connection is made with %v\n", conn)
}

