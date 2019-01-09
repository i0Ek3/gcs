package main

import (
        "fmt"
        "net"
)

func main() {
        fmt.Println("Starting the server...")
        // create listener which is a var to listen and accept requests from client
        listener, err := net.Listen("tcp", "localhost:50000") // 127.0.0.1:50000 based on tcp
        if err != nil {
                fmt.Println("Error listening", err.Error())
                return // end program
        }

        // listen connections from client
        for {
                conn, err := listener.Accept() // waitting for requests
                if err != nil {
                        fmt.Println("Error acceppting", err.Error())
                        return 
                }
                go doServerStuff(conn) // create different goroutines for every connection
        }
}

func doServerStuff(conn net.Conn) { 
        for {
                buf := make([]byte, 512) 
                len, err := conn.Read(buf) // len(data)
                if err != nil {
                        fmt.Println("Error reading", err.Error())
                        return
                }
                //fmt.Printf("Received data: [")
                //fmt.Printf("%v", string(buf[:len]))
                //fmt.Println("]")
                fmt.Printf("Received data: [ %v ] \n", string(buf[:len]))
        }
}


