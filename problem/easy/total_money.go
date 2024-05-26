package problem

func TotalMoney(n int) int {

	total := (28*(n/7) + 7*(n/7)*(n/7-1)/2) + (n%7)*(n/7) + (n%7)*(n%7+1)/2

	return total
}
