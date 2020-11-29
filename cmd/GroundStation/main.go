package main

import (
	"bufio"
	"fmt"
	"github.com/Autonomeasure/GroundStation/pkg/Radio"
	"log"
)

func main() {
	// Create a connection to the serial port
	serialPort, err := Radio.OpenSerial("/dev/serial0", 9600)

	defer serialPort.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Create a scanner for the serial port
	scanner := bufio.NewScanner(bufio.NewReader(serialPort))

	// Keep reading the data
	for scanner.Scan() {
		// Print the incoming data
		fmt.Println(scanner.Text())
	}
}
