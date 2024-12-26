package hackkerankeasy

import (
	"math"
)

func SherlockAndSquare(a int32, b int32) int32 {
	sqrtA := math.Ceil(math.Sqrt(float64(a)))
	sqrtB := math.Floor(math.Sqrt(float64(b)))
	return int32(sqrtB - sqrtA + 1)
}
