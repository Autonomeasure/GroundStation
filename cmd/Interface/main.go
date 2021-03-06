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
	"fmt"
	"github.com/Autonomeasure/GroundStation/pkg/Radio"
	"log"
)

//var database Database.Database

func main() {
	// Create a connection to the serial port
	serialPort, err := Radio.OpenSerial("\\COM3", 9600)

	//database.Open()

	defer func() {
		e := serialPort.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}


	buf := make([]byte, 1024)

	for {
		n, err := serialPort.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		_, _, err = Radio.Decode(buf, n)
		if err != nil {
			fmt.Println(err)
		}

	}

	//
	//// Create a scanner for the serial port
	//scanner := bufio.NewScanner(bufio.NewReader(serialPort))
	////err = Radio.ConfigureRadioModule(serialPort)
	////if err != nil {
	////	log.Fatal(err)
	////}
	//
	//// Keep reading the data
	//for scanner.Scan() {
	//	// Save the incoming data
	//	var input = scanner.Text()
	//	//fmt.Println([]byte(input))
	//	//fmt.Println(input)
	//	// Create a Radio.Packet object and print it
	//	p, err := Radio.Decode(input)
	//	if err != nil {
	//		//fmt.Println(err)
	//		continue
	//	}
	//	//fmt.Printf("%+v\n\n", p)
	//	fmt.Printf("ID: %d | bTemp: %.2f | mTemp: %.2f | pressure: %.2f | time: %s (%d:%d:%d:%d) | lat: %.12f | lon: %.12f | gps_alt: %.2f | bmp_alt: %.2f | gx: %.2f | gy: %.2f | gz %.2f | ax: %.2f | ay: %.2f | az %.2f\r", p.ID, p.Temperature.BMP, p.Temperature.MPU, p.Pressure, p.Time, int(p.Time[0]) - 1, int(p.Time[1]) - 1, int(p.Time[2]) - 1, int(p.Time[3]) - 1, p.GPS.Latitude, p.GPS.Longitude, p.GPS.Altitude, p.BMPAltitude, p.Gyroscope.X, p.Gyroscope.Y, p.Gyroscope.Z, p.Acceleration.X, p.Acceleration.Y, p.Acceleration.Z)
	//	//go database.SaveRadioPacket(p)
	//}
}
