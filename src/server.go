package main

import (
        "fmt"
        "net"
        "os"
        "time"
        "strings"
)

var mapUsers map[string]int

func main() {
        mapUsers = make(map[string]int)

        fmt.Println("Starting the server...")
        
        // create listener which is a var to listen and accept requests from client
        listener, err := net.Listen("tcp", "localhost:50000") // 127.0.0.1:50000 based on tcp
        checkError(err)
        
        // listen connections from client
        for {
                conn, err := listener.Accept() // waitting for requests
                checkError(err)
                go doServerStuff(conn) // create different goroutines for every connection
        }
}

func doServerStuff(conn net.Conn) {  
        for {
                buf := make([]byte, 512) 
                _, err := conn.Read(buf)
                checkError(err)
                input := string(buf)
                
                if strings.Contains(input, ": shutdown") {
                        fmt.Println("Server shutting down...")
                        os.Exit(0)
                }

                if strings.Contains(input, ": list") {
                        listUser()
                }

                if strings.Contains(input, ": time") {
                        t := time.Now()
                        fmt.Println(t.Format(time.ANSIC))
                }
                
                if strings.Contains(input, ": help") {
                        fmt.Println("You can use these commands as follows:")
                        fmt.Println("   help           To check avaiable commands.")
                        fmt.Println("   shutdown       To close the connection with server.")
                        fmt.Println("   list           To check the users under the connection.")
                        fmt.Println("   time           To check the current time.")
                }

                index := strings.Index(input, ":")
                clientName := input[0 : index-1]
                mapUsers[string(clientName)] = 1
                fmt.Printf("Received data: [ %v ] \n", string(buf))
        }
}

func listUser() {
        fmt.Println("This is the client list: 1=active, 0=inactive")
        for key, value := range mapUsers {
                fmt.Printf("User %s is %d\n", key, value)    
        }
}

func checkError(err error) {
        if err != nil {
            panic("Error: " + err.Error())
        }
}
