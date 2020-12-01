package Radio

import (
	"github.com/Autonomeasure/GroundStation/pkg"
	"github.com/Autonomeasure/GroundStation/pkg/GPS"
	"strconv"
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

	for _, str := range s {
		//fmt.Println(i)
		splitted := strings.Split(str, "=")
		if splitted[0] == "tm" {
			temp, _ := strconv.ParseFloat(splitted[1], 32)
			p.Temperature.MPU = float32(temp)
			continue
		} else if splitted[0] == "tb" {
			temp, _ := strconv.ParseFloat(splitted[1], 32)
			p.Temperature.BMP = float32(temp)
			continue
		} else if splitted[0] == "p" {
			pressure, _ := strconv.ParseFloat(splitted[1], 32)
			p.Pressure = float32(pressure)
			continue
		} else if splitted[0] == "ax" {
			ax, _ := strconv.ParseFloat(splitted[1], 32)
			p.Acceleration.X = float32(ax)
			continue
		} else if splitted[0] == "ay" {
			ay, _ := strconv.ParseFloat(splitted[1], 32)
			p.Acceleration.Y = float32(ay)
			continue
		} else if splitted[0] == "az" {
			az, _ := strconv.ParseFloat(splitted[1], 32)
			p.Acceleration.Z = float32(az)
			continue
		} else if splitted[0] == "gx" {
			gx, _ := strconv.ParseFloat(splitted[1], 32)
			p.Gyroscope.X = float32(gx)
			continue
		} else if splitted[0] == "gy" {
			gy, _ := strconv.ParseFloat(splitted[1], 32)
			p.Gyroscope.Y = float32(gy)
			continue
		} else if splitted[0] == "gz" {
			gz, _ := strconv.ParseFloat(splitted[1], 32)
			p.Gyroscope.Z = float32(gz)
			continue
		}
	}

	return p
}