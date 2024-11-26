package hackkerankeasy

func GetMaxBookCombination(listPrices []int32, money int32) int32 {
	// example:
	// listPrices = [100000, 60000, 30000,45000, 50000]
	// money = 250000

	// listPrices = [50000, 100000, 100000, 75000,50000]
	// money = 50000
	// result = 1

	// 1. sort listPrices and remove duplicate
	// 2. loop listPrices
	// 3. if money - listPrices[i] >= 0, money = money - listPrices[i], result++
	// 4. if money - listPrices[i] < 0, break

	// 1. sort listPrices and remove duplicate
	listPrices = sortAndRemoveDuplicate(listPrices)

	// 2. loop listPrices
	var result int32 = 0
	for i := 0; i < len(listPrices); i++ {
		// 3. if money - listPrices[i] >= 0, money = money - listPrices[i], result++
		if money-listPrices[i] >= 0 {
			money = money - listPrices[i]
			result++
		} else {
			// 4. if money - listPrices[i] < 0, break
			break
		}
	}

	return result
}

func sortAndRemoveDuplicate(listPrices []int32) []int32 {
	// 1. sort listPrices
	// 2. remove duplicate
	// 3. return listPrices

	// 1. sort listPrices
	for i := 0; i < len(listPrices); i++ {
		for j := 0; j < len(listPrices); j++ {
			if listPrices[i] < listPrices[j] {
				listPrices[i], listPrices[j] = listPrices[j], listPrices[i]
			}
		}
	}

	// 2. remove duplicate
	var newListPrices []int32
	for i := 0; i < len(listPrices); i++ {
		if i == 0 {
			newListPrices = append(newListPrices, listPrices[i])
		} else {
			if listPrices[i] != listPrices[i-1] {
				newListPrices = append(newListPrices, listPrices[i])
			}
		}
	}

	// 3. return listPrices
	return newListPrices
}
