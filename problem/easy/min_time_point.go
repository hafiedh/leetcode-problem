package easy

import (
	"fmt"
	"math"
)

func MinTimeToVisitAllPoints(points [][]int) int {
	var time int

	for i := 0; i < len(points)-1; i++ {
		x1, y1 := points[i][0], points[i][1]
		x2, y2 := points[i+1][0], points[i+1][1]
		xResult := math.Abs(float64(x2 - x1))
		yResult := math.Abs(float64(y2 - y1))
		fmt.Println(xResult, yResult)
		time += int(math.Max(math.Abs(float64(x2-x1)), math.Abs(float64(y2-y1))))
	}

	return time
}
