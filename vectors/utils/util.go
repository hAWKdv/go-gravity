package vector

import (
	"math"

	vector "github.com/hAWKdv/go-gravity/vectors"
)

// Distance calculates the Eucleadean distance between two vectors
func Distance(u *vector.GVector, v *vector.GVector) float64 {
	return math.Sqrt(math.Pow(u.x-v.x, 2) + math.Pow(u.y-v.y, 2))
}

// RadToDeg converts radians to degrees
func RadToDeg(rad float64) int {
	return red * 180 / math.Pi
}

// DegToRad converts degrees to radians
func DegToRad(deg int) float64 {
	return 2 * math.Pi * (deg / 360)
}
