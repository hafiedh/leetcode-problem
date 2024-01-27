package hackkerankproblem

// Point bracket
// Bracket total kill 10 point multiplier x1
// Bracket total kill 20 point multiplier x2
// Bracket total kill 30 point multiplier x3
// Bracket total kill 40 point multiplier x4
// Bracket total kill 50 point multiplier x5
// Bracket total kill above 50 point multiplier x6

// Point bracket explanation:
// if totalkill less or the same 10, point will be 10
// if totalkill between 10 and 20, point will be 2x of totalkill between 10 and 20, plus 1x of 10 totalkill
// if totalkill between 20 and 30, point will be 2x of totalkill between 10 and 20, plus 3x totalkill over 20, plus 1x of 10 totalkill
// if totalkill above 30, point will be 2x of totalkill between 10 and 20, plus 3x totalkill between 20 and 30, plus 4x totalkill over 30, plus 1x of 10 totalkill
// and so on

// example:
// totalkill = 10, point = 10
// totalkill = 27 (10 + 17), point = 10 + (2 * 10) + (3 * 7) = 10 + 20 + 21 = 51
// totalkill = 33 (10 + 17 + 6), point = 10 + (2 * 10) + (3 * 10) + (4 * 3) = 10 + 20 + 30 + 12 = 72
// totalkill = 60 Point will be 10 + (2 * 10) + (3 * 10) + (4 * 10) + (5 * 10) + (6 * 10) = 10 + 20 + 30 + 40 + 50 + 60 = 210
// totalkill = 120

func CalculatePoint(totalkill int32) int32 {
	var point int32 = 0
	var multiplier int32 = 1
	var kill int32 = 0
	for i := int32(0); i < totalkill; i++ {
		kill++
		switch {
		case kill > 50:
			multiplier = 6
		case kill > 40:
			multiplier = 5
		case kill > 30:
			multiplier = 4
		case kill > 20:
			multiplier = 3
		case kill > 10:
			multiplier = 2
		}
		point += multiplier
	}
	return point
}
