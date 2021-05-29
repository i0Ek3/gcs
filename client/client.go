package main 

import (
        "fmt"
        "os"
        "net"
        "bufio"
        "strings"

        e "err"
)

func main() {
        // open connections
        // client use net.Dial() to create a connection with server
        conn, err := net.Dial("tcp", "localhost:50000") 
        e.checkError(err)

        // read buffer from os.Stdin 
        inputReader := bufio.NewReader(os.Stdin)
        fmt.Println("Enter your name: ")
        clientName, _ := inputReader.ReadString('\n')
        trimmedClient := strings.Trim(clientName, "\n")

        // send data to server util it's program exit
        for {
                fmt.Println("Enter your words to send to server, type q to quit.")
                input, _ := inputReader.ReadString('\n')
                trimmedInput := strings.Trim(input, "\n")
                if trimmedInput == "q" {
                        return 
                }
                
                // after slice write input to server
                _, err = conn.Write([]byte(trimmedClient + " : " + trimmedInput))
                e.checkError(err)
        }
}
