package problem

func AddBinary(a string, b string) string {
	aArr := []rune(a)
	bArr := []rune(b)

	aLen := len(aArr)
	bLen := len(bArr)

	var result string
	var carry int
	for i := 0; i < aLen || i < bLen; i++ {
		var aVal, bVal int
		if i < aLen {
			aVal = int(aArr[aLen-1-i] - '0')
		}

		if i < bLen {
			bVal = int(bArr[bLen-1-i] - '0')
		}

		sum := aVal + bVal + carry
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}

		if sum%2 == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}
