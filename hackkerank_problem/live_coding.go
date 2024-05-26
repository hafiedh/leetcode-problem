package hackkerankproblem

import "strconv"

func Calculate(input string) int {
	var result int
	if string(input[0]) == "-" {
		result, _ = strconv.Atoi(string(input[1]))
		result = result * -1
	} else {
		result, _ = strconv.Atoi(string(input[0]))
	}

	for i := range input {
		if i == 0 {
			continue
		}
		if string(input[i]) == "+" {
			i1, _ := strconv.Atoi(string(input[i+1]))
			result = Add(result, i1)
		}
		if string(input[i]) == "-" {
			i2, _ := strconv.Atoi(string(input[i+1]))
			result = Substract(result, i2)
		}
	}
	return result

}

func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}
