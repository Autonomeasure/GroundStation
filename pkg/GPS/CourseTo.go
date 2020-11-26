package GPS

import "math"

func CourseTo(lat1 float32, lon1 float32, lat2 float32, lon2 float32) {
	// Returns the course in degrees (north = 0, west = 270) from position 1 to posotion two

	dlon float32 = (lon2 - lon1) * (math.Pi / 180)

	
}
