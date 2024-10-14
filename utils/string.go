package utils

func Reverse(input ...string) string {
	var result string
	for _, s := range input {
		for i := len(s) - 1; i >= 0; i-- {
			result += string(s[i])
		}
	}
	return result
}

func ReverseInt(input int) int {
	var result int
	for input > 0 {
		remainder := input % 10
		result = result*10 + remainder
		input /= 10
	}
	return result
}
