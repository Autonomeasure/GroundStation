package GPS

import "math"

/*
private double angleFromCoordinate(double lat1, double long1, double lat2,
        double long2) {

    double dLon = (long2 - long1);

    double y = Math.sin(dLon) * Math.cos(lat2);
    double x = Math.cos(lat1) * Math.sin(lat2) - Math.sin(lat1)
            * Math.cos(lat2) * Math.cos(dLon);

    double brng = Math.atan2(y, x);

    brng = Math.toDegrees(brng);
    brng = (brng + 360) % 360;
    brng = 360 - brng; // count degrees counter-clockwise - remove to make clockwise

    return brng;
}
 */

func CourseTo(lat1 float32, lon1 float32, lat2 float32, lon2 float32) float32 {
	// Returns the course in degrees (north = 0, west = 270) from position 1 to position 2

	dLon := lon2 - lon1
	var y float32 = float32(math.Sin(float64(dLon)) * math.Cos(float64(lat2)))
	var x float32 = float32(math.Cos(float64(lat1))*math.Sin(float64(lat2)) - math.Sin(float64(lat1)) * math.Cos(float64(lat2)) * math.Cos(float64(dLon)))
	a := math.Atan2(float64(y), float64(x))

	a = a * 180 / math.Pi
	a = math.Mod(a + 360, 360)
	a = 360 - a
	return float32(a)
}
