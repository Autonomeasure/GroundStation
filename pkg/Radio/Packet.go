package Radio

import (
	"github.com/Autonomeasure/GroundStation/pkg"
	"github.com/Autonomeasure/GroundStation/pkg/GPS"
)

type Temperature struct {
	BMP float32 `json:"bmpTemp"`
	MPU float32 `json:"mpuTemp"`
}

type Packet struct {
	Temperature 	Temperature `json:"temperature"`	// Temperature in Celsius
	Pressure 		float32 	`json:"pressure"` 		// Pressure in hPa
	Location 		GPS.Packet	`json:"location"`		// A Location object contains all the information from the GPS module
	Acceleration 	pkg.Vector3 `json:"acceleration"` 	// Acceleration is an instance of Vector3 containing three axis (xyz)
	Gyroscope 		pkg.Vector3 `json:"gyroscope"` 		// Gyroscope is an instance of Vector3 containing three axis (xyz)
}