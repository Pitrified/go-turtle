package turtle

import "math"

// Convert degrees to radians.
func Deg2rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// Convert radians to degrees.
func Rad2deg(rad float64) float64 {
	return rad / math.Pi * 180
}

// int can be abs too!
func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
