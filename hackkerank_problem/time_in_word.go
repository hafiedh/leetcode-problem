package hackkerankproblem

func TimeInWords(h int32, m int32) string {
	var result string
	var hours []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve"}
	var minutes []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen",
		"twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six",
		"twenty seven", "twenty eight", "twenty nine", "half"}
	if m == 0 {
		result = hours[h-1] + " o' clock"
	} else if m == 1 {
		result = minutes[m-1] + " minute past " + hours[h-1]
	} else if m == 15 || m == 30 {
		result = minutes[m-1] + " past " + hours[h-1]
	} else if m < 30 {
		result = minutes[m-1] + " minutes past " + hours[h-1]
	} else if m == 45 {
		result = minutes[60-m-1] + " to " + hours[h]
	} else {
		result = minutes[60-m-1] + " minutes to " + hours[h]
	}
	return result
}
