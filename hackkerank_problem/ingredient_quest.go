package hackkerankeasy

// let's say recipe = "MOCHI" and inredients = "IHCOMIHCOM"
// result = 2 (because there are only 2 mochi in inredients)

// let's say recipe = "DORAYAKI" and inredients = "DDDOOORRRAAAAYYKKIIBBPPXX"
// result = 2 (because there are only 2 dorayaki in inredients)

func Cook(recipe string, inredients string) int32 {
	var recipeMap = make(map[string]int)
	var ingredientsMap = make(map[string]int)
	var count int32

	for i := 0; i < len(recipe); i++ {
		recipeMap[string(recipe[i])]++
	}

	for i := 0; i < len(inredients); i++ {
		ingredientsMap[string(inredients[i])]++
	}

	// check if a;

	for key, value := range recipeMap {
		if ingredientsMap[key] < value {
			count = 0
			break
		} else {
			count = int32(ingredientsMap[key] / value)
		}

	}

	return count

}
