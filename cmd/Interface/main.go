/*
 * CanSat 2020-2021 Team Autonomeasure - GroundStation | Interface
 * Author: Joep van Dijk
 * First mission:
 *    Our first mission is to be able to land properly, collect data from the BMP280 sensor and send that data to the ground station at least one time per second.
 *
 * Second mission:
 *    Our second mission is to make a sustainable data collection station.
 *
 * Links:
 *    GitHub Autonomeasure:           https://github.com/Autonomeasure
 *    GitHub GroundStation repo:      https://github.com/Autonomeasure/GroundStation
 *    GitHub Can repo:                https://github.com/Autonomeasure/Can
 *    Instagram:                      https://instagram.com/Autonomeasure/
 *
 * This is the "interface" application
 * The interface will do the following things:
 *	- Receive the incoming data from the APC220 radio module (Serial/UART communication)
 *		- Save the received data to the database so the webserver / interface application can access it
 */
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

	defer func() {
		e := serialPort.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	// Create a scanner for the serial port
	scanner := bufio.NewScanner(bufio.NewReader(serialPort))

	// Keep reading the data
	for scanner.Scan() {
		// Print the incoming data
		var input = scanner.Text()
		fmt.Println(input)
		// Create a Radio.Packet object and print it
		var p = Radio.Decode(input)
		fmt.Printf("%+v\n", p)
	}
}
