package medium

import "fmt"

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			res[i+j+1] += x * y
		}
	}

	fmt.Println(res)

	for i := m + n - 1; i > 0; i-- {
		res[i-1] += res[i] / 10
		res[i] %= 10
	}
	idx := 0
	if res[0] == 0 {
		idx = 1
	}
	ans := ""
	for ; idx < m+n; idx++ {
		ans += fmt.Sprint(res[idx])
	}
	return ans
}
