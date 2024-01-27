package hackkerankproblem

import "fmt"

func FindMinimumDifference(packats []int32, participants int32) string {
	// example:
	// packats = [2, 20, 13, 31, 10, 12, 17, 16, 27]
	// participants = 3
	// result = 3 [10, 12, 13]

	// exaample 2:
	// packats = [10, 6, 8, 12,15,6, 14 16]
	// participants = 4
	// result = 4 [6, 6, 8, 10]

	// 1. sort packats
	// 2. find minimum difference
	// 3. return minimum difference
	var result string
	// 1. sort packats
	for i := 0; i < len(packats); i++ {
		for j := 0; j < len(packats); j++ {
			if packats[i] < packats[j] {
				packats[i], packats[j] = packats[j], packats[i]
			}
		}
	}

	// 2. find minimum difference
	var minimumDifference int32 = 0
	var tempMinimumDifference int32 = 0
	var tempMinimumDifferenceIndex int32 = 0
	for i := int32(0); i < participants-1; i++ {
		tempMinimumDifference = packats[i+1] - packats[i]
		if i == 0 {
			minimumDifference = tempMinimumDifference
		} else {
			if minimumDifference > tempMinimumDifference {
				minimumDifference = tempMinimumDifference
				tempMinimumDifferenceIndex = i
			}
		}
	}

	// 3. return minimum difference
	var res []int32
	for i := tempMinimumDifferenceIndex; i < tempMinimumDifferenceIndex+participants; i++ {
		res = append(res, packats[i])
	}

	for i := 0; i < len(res); i++ {
		if i == 0 {
			result = result + string(res[i])
		} else {
			result = result + " " + string(res[i])
		}
	}
	fmt.Println(result)

	return result
}
