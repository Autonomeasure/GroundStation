package Tests

import (
	"fmt"
	"github.com/Autonomeasure/GroundStation/pkg/GPS"
	"testing"
)

func TestGPSCourseTo(t *testing.T) {
	var lat1 = 0.0
	var lon1 = 0.0
	var lat2 = 1.0
	var lon2 = 1.0
	fmt.Println(GPS.CourseTo(float32(lat1), float32(lon1), float32(lat2), float32(lon2)))
}