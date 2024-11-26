package hackkerankeasy

func AppendAndDelete(s string, t string, k int32) string {

	sArr := []rune(s)
	tArr := []rune(t)

	sLen := len(sArr)
	tLen := len(tArr)

	var result string
	if sLen+tLen <= int(k) {
		result = "Yes"
	} else {
		var i int
		for i = 0; i < sLen && i < tLen; i++ {
			if sArr[i] != tArr[i] {
				break
			}
		}
		if sLen+tLen-2*i > int(k) {
			result = "No"
		} else if (sLen+tLen-2*i)%2 == int(k)%2 {
			result = "Yes"
		} else if sLen+tLen-int(k) < 0 {
			result = "Yes"
		} else {
			result = "No"
		}
	}

	return result
}
