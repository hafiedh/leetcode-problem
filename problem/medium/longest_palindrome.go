package medium

func LongestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if isPalindrome(s[i:j]) {
				if len(s[i:j]) > len(result) {
					result = s[i:j]
				}
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
