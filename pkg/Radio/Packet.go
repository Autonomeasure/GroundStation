package Radio

import (
	"fmt"
	"github.com/Autonomeasure/GroundStation/pkg"
	"github.com/Autonomeasure/GroundStation/pkg/GPS"
	"strings"
)

type Temperature struct {
	BMP float32 `json:"bmpTemp"`
	MPU float32 `json:"mpuTemp"`
}

type Packet struct {
	Temperature 	Temperature `json:"temperature"`	// Temperature in Celsius
	Pressure 		float32 	`json:"pressure"` 		// Pressure in hPa
	GPS 			GPS.Packet	`json:"gps"`			// A GPS object contains all the information from the GPS module
	Acceleration 	pkg.Vector3 `json:"acceleration"` 	// Acceleration is an instance of Vector3 containing three axis (xyz)
	Gyroscope 		pkg.Vector3 `json:"gyroscope"` 		// Gyroscope is an instance of Vector3 containing three axis (xyz)
}

// Decode the received message into the Packet struct and return a Packet
func Decode(input string) Packet {
	var p Packet

	s := strings.Split(input, ";")
	//fmt.Println(s)

	for _, i := range s {
		fmt.Println(i)
	}

	return p
}