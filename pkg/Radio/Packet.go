package Radio

import (
	"errors"
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
	BMPAltitude		float32		`json:"bmpAltitude"` 	// Altitude based on the pressure from the BMP module
	GPS 			GPS.Packet	`json:"gps"`			// A GPS object contains all the information from the GPS module
	Acceleration 	pkg.Vector3 `json:"acceleration"` 	// Acceleration is an instance of Vector3 containing three axis (xyz)
	Gyroscope 		pkg.Vector3 `json:"gyroscope"` 		// Gyroscope is an instance of Vector3 containing three axis (xyz)
	Time 			[]byte		`json:"time"`			// Time is a char array containing three items: hh:mm:ss
}

// Decode the received message into the Packet struct and return a Packet
func Decode(input string) (Packet, error) {
	var p Packet

	s := strings.Split(input, ";")

	if len(s) < 15 {
		return p, errors.New("invalid packet was received")
	}

	id, _ := strconv.ParseUint(s[0], 10, 32)
	p.ID = uint32(id)
	bmpTemp, _ := strconv.ParseInt(s[1], 10, 32)
	bmpTempF := float32(bmpTemp) / 100.0
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
	bmpAlt, _ := strconv.ParseFloat(s[7], 32)
	p.BMPAltitude = float32(bmpAlt)
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
	if len(time) != 4 {
		//p.Time = []byte{'A', 'A', 'A', 'A'}
		p.Time = []byte{0, 0, 0, 0}
	}

	return p, nil
}