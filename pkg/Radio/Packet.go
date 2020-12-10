package Radio

import (
	"fmt"
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
	ID				uint32 		`json:"id"`
	Temperature 	Temperature `json:"temperature"`	// Temperature in Celsius
	Pressure 		float32 	`json:"pressure"` 		// Pressure in hPa
	GPS 			GPS.Packet	`json:"gps"`			// A GPS object contains all the information from the GPS module
	Acceleration 	pkg.Vector3 `json:"acceleration"` 	// Acceleration is an instance of Vector3 containing three axis (xyz)
	Gyroscope 		pkg.Vector3 `json:"gyroscope"` 		// Gyroscope is an instance of Vector3 containing three axis (xyz)
	Time 			[]byte		`json:"time"`			// Time is a char array containing three items: hh:mm:ss
}

// Decode the received message into the Packet struct and return a Packet
func Decode(input string) Packet {
	var p Packet

	//fmt.Println([]byte(input))

	s := strings.Split(input, ";")
	//fmt.Printf("%+v\n", s)

	id, _ := strconv.ParseUint(s[0], 10, 32)
	p.ID = uint32(id)
	bmpTemp, _ := strconv.ParseInt(s[1], 10, 32)
	bmpTempF := float32(bmpTemp) / 100.0
	fmt.Println(s[0])
	fmt.Println(bmpTemp)
	fmt.Println(bmpTempF)
	p.Temperature.BMP =  bmpTempF
	mpuTemp, _ := strconv.ParseInt(s[2], 10, 32)
	mpuTempF := float32(mpuTemp) / 100.0
	p.Temperature.MPU =  mpuTempF
	pressure, _ := strconv.ParseFloat(s[3],32)
	p.Pressure = float32(pressure)
	lat, _ := strconv.ParseFloat(s[4],32)
	p.GPS.Latitude = float32(lat)
	lon, _ := strconv.ParseFloat(s[5],32)
	p.GPS.Longitude = float32(lon)
	alt, _ := strconv.ParseFloat(s[6], 32)
	p.GPS.Altitude = float32(alt)
	speed, _ := strconv.ParseFloat(s[7], 32)
	p.GPS.Speed = float32(speed) / 100.0
	ax, _ := strconv.ParseInt(s[8], 10, 32)
	ay, _ := strconv.ParseInt(s[9], 10, 32)
	az, _ := strconv.ParseInt(s[10], 10, 32)
	p.Acceleration.X = float32(ax) / 100.0
	p.Acceleration.Y = float32(ay) / 100.0
	p.Acceleration.Z = float32(az) / 100.0
	gx, _ := strconv.ParseInt(s[11], 10, 32)
	gy, _ := strconv.ParseInt(s[12], 10, 32)
	gz, _ := strconv.ParseInt(s[13], 10, 32)
	p.Gyroscope.X = float32(gx) / 100.0
	p.Gyroscope.Y = float32(gy) / 100.0
	p.Gyroscope.Z = float32(gz) / 100.0
	time := s[14]
	p.Time = []byte(time)
	//p.Time[0] = string(rune(time)[0])
	//p.Time[1] = time[1]
	//p.Time[2] = time[2]



	//// Get all the different items and put it in the Packet object
	//for _, str := range s {
	//	splitted := strings.Split(str, "=")
	//	if splitted[0] == "id" {
	//		id, _ := strconv.ParseUint(splitted[1], 10, 32)
	//		p.ID = uint32(id)
	//	} else if splitted[0] == "tm" {
	//		temp, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Temperature.MPU = float32(temp)
	//		continue
	//	} else if splitted[0] == "tb" {
	//		temp, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Temperature.BMP = float32(temp)
	//		continue
	//	} else if splitted[0] == "p" {
	//		pressure, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Pressure = float32(pressure)
	//		continue
	//	} else if splitted[0] == "ax" {
	//		ax, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Acceleration.X = float32(ax)
	//		continue
	//	} else if splitted[0] == "ay" {
	//		ay, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Acceleration.Y = float32(ay)
	//		continue
	//	} else if splitted[0] == "az" {
	//		az, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Acceleration.Z = float32(az)
	//		continue
	//	} else if splitted[0] == "gx" {
	//		gx, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Gyroscope.X = float32(gx)
	//		continue
	//	} else if splitted[0] == "gy" {
	//		gy, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Gyroscope.Y = float32(gy)
	//		continue
	//	} else if splitted[0] == "gz" {
	//		gz, _ := strconv.ParseFloat(splitted[1], 32)
	//		p.Gyroscope.Z = float32(gz)
	//		continue
	//	}
	//}

	return p
}