package main 

import (
        "fmt"
        "os"
        "net"
        "bufio"
        "strings"
)

func main() {
        // open connections
        conn, err := net.Dial("tcp", "localhost:50000") // client use net.Dial() to create a connection with server
        checkError(err)

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
                checkError(err)
        }
}

func checkError(err error) {
        if err != nil {
                panic("Error: " + err.Error())    
        }
}
