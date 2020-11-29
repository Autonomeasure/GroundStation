package main

import "github.com/Autonomeasure/GroundStation/pkg/Radio"

func main() {
	Radio.OpenSerial("/dev/serial0", 9600)
}
