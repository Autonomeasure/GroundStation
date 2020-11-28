package GroundStation

import "github.com/Autonomeasure/GroundStation/pkg/Radio"

func main() {
	Radio.OpenSerial("/dev/ttyS0", 9600)
}
