//package main
//
//import (
//	"bufio"
//	"fmt"
//	"math/rand"
//	"net"
//	"os"
//	"strings"
//	"time"
//)
//
//func main() {
//	// Start the TCP server at the given port
//	arguments := os.Args
//	if len(arguments) == 1 {
//		fmt.Println("Please provide a port number!")
//		return
//	}
//
//	PORT := ":" + arguments[1]
//	l, err := net.Listen("tcp4", PORT)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer l.Close()
//	rand.Seed(time.Now().Unix())
//
//	for {
//		c, err := l.Accept()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		go handleConnection(c)
//	}
//}
//
//func handleConnection(c net.Conn) {
//	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
//	for {
//		netData, err := bufio.NewReader(c).ReadString('\n')
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		fmt.Print("Incoming: "); fmt.Print(netData)
//		data := strings.TrimSpace(netData)
//		if strings.Contains(netData, "PING") {
//			_, err = c.Write([]byte("PONG!")); if err != nil {
//				panic(err)
//			}
//			continue
//		}
//		if data == "ALL" {
//			// Return all the data
//		}
//
//		//result := strconv.Itoa(rand.Int()) + "\n"
//		//_, err = c.Write([]byte(string(result))); if err != nil {
//		//	panic(err)
//		//}
//	}
//}
package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp4"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}