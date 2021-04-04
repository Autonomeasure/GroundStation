package Tests

import (
	"github.com/Autonomeasure/GroundStation/pkg/GPS"
	"testing"
)

func TestGPSDistance(t *testing.T) {
	//var lat1 = 52.082050
	//var lon1 = 5.238720
	//var lat2 = 52.3547498
	//var lon2 = 4.8339211

	var lat1 = 0.0
	var lon1 = 0.0
	var lat2 = 2.0
	var lon2 = 1.0

	GPS.Distance(lat1, lon1, lat2, lon2, "K")
}
