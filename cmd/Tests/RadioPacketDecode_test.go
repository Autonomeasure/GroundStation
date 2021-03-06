package Tests

import (
	"github.com/Autonomeasure/GroundStation/pkg"
	"github.com/Autonomeasure/GroundStation/pkg/GPS"
	"github.com/Autonomeasure/GroundStation/pkg/Radio"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRadioPacketDecode(t *testing.T) {
	// Check if a normal, valid input gets parsed correctly without any errors
	input := "0;2010;2030;1013.25;0.000000000000;1.000000000000;1000.00;1000.00;170;200;450;18000;9000;1000;AAAA;\n"
	t.Log("Input string: " + input)
	packet, err := Radio.Decode(input)
	var testPacket = Radio.Packet{
		ID: 0,
		Temperature: Radio.Temperature{
			BMP: 20.10,
			MPU: 20.30,
		},
		Pressure: 1013.25,
		BMPAltitude: 1000.00,
		GPS: GPS.Packet{
			Latitude: 0.000000000000,
			Longitude: 1.000000000000,
			Altitude: 1000.00,
		},
		Acceleration: pkg.Vector3{
			X: 1.70,
			Y: 2.00,
			Z: 4.50,
		}, Gyroscope: pkg.Vector3{
			X: 180.00,
			Y: 90.00,
			Z: 10.00,
		},
		Time: []byte{'A', 'A', 'A', 'A'},
	}

	t.Logf("Packet: \n%+v\n", packet)
	t.Logf("TestPacket: \n%+v\n", testPacket)


	if !cmp.Equal(packet, testPacket) && err == nil {
		t.Error("Radio.Decode was incorrect")
		t.Error("Input string: 0;2010;2030;1013.25;0.000000000000;1.000000000000;1000.00;1000.00;170;200;450;18000;9000;1000;AAAA;\n")
		t.Error("Packet: ")
		t.Errorf("%+v\n", packet)
		t.Error("TestPacket: ")
		t.Errorf("%+v\n", testPacket)
		t.Error("Error: ", err)
		t.Failed()
	}



	// Check if an invalid input gives the expected error
	input = "0;2010;2030;1013.25;0.000000;1.000000;1000.00;340;170;200;450;18000;9000;1000;\n"
	t.Log("Input string: " + input)
	_, err = Radio.Decode(input)
	t.Log("err: ", err)
	if err == nil {
		t.Error("Invalid input was given but no error returned")
		t.Failed()
	}
	t.Log("Error was handled correctly")
}
