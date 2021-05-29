package main

import (
        "fmt"
        "net"
        "os"
        "time"
        "flag"
        "syscall"
        "strings"

        e "err"
)

var mapUsers map[string]int
const maxRead = 250

func main() {
        mapUsers = make(map[string]int)
        fmt.Println("Starting the server...")
        
        flag.Parse()
        if flag.NArg() != 2 {
                panic("Usage: ./server host port")  
        }
        host := fmt.Sprintf("%s", flag.Arg(0))
        port := fmt.Sprintf("%s", flag.Arg(1))
        listener := initServer(host, port)

        // listen connections from client
        for {
                // waitting for requests
                conn, err := listener.Accept() 
                e.checkError(err, "Accept: ")
                // create different goroutines for every connection
                go doServerStuff(conn) 
        }
}

func initServer(host, port string) *net.TCPListener {
        serverAddr, err := net.ResolveTCPAddr("tcp", host + ":" + port)
        e.checkError(err, "Cannot resolving address: " + host + ":" + port +" ")
        listener, err := net.ListenTCP("tcp", serverAddr)
        e.checkError(err, "ListenTCP: ")
        fmt.Println("Listening to: ", listener.Addr().String())
        return listener 
}

func doServerStuff(conn net.Conn) {  
        connFrom := conn.RemoteAddr().String()
        fmt.Println("Connection from: ", connFrom)

        for {
                buf := make([]byte, maxRead + 1) 
                _, err := conn.Read(buf[0:maxRead])
                buf[maxRead] = 0
                
                switch err {
                case nil:
                        fmt.Printf("Received data: [ %v ] \n", string(buf))
                case syscall.EAGAIN:
                        continue
                default:
                        goto CLOSE
                }

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
        }
CLOSE:
        err := conn.Close()
        fmt.Println("Closed connection: ", connFrom)
        e.checkError(err, "Close: ")
}

func listUser() {
        fmt.Println("This is the client list: 1=active, 0=inactive")
        for key, value := range mapUsers {
                fmt.Printf("User %s is %d\n", key, value)    
        }
}
