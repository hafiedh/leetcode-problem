package personal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func Check6MonthOrLater(startDate time.Time, targetDate time.Time) bool {
	timeDiff := targetDate.Sub(startDate)
	monthDiff := int(timeDiff.Hours() / 24 / 30)
	return monthDiff >= 6
}

func EncryptUrl(url string) (encryptedUrl string, err error) {
	key := []byte("testingtesting!!") // should be 16, 24 or 32 bytes and .env
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return
	}
	encryptedUrl = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(url), nil))
	return
}

func DecryptUrl(encryptedUrl string) (url string, err error) {
	key := []byte("testingtesting!!") // should be 16, 24 or 32
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}
	decodeBase64, err := base64.URLEncoding.DecodeString(encryptedUrl)
	if err != nil {
		return
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := []byte(decodeBase64)[:nonceSize], []byte(decodeBase64)[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}
	url = string(plaintext)
	return
}

func MaxTransaction(transactions []int32) int32 {
	if len(transactions) == 0 {
		return 0
	}
	var maxTr, balance, indexFirst int32

	for i, tr := range transactions {
		if tr > 0 {
			indexFirst = int32(i)
			balance = tr
			maxTr = 1
			break
		}
	}

	for i := indexFirst + 1; i < int32(len(transactions)); i++ {
		if balance+transactions[i] >= 0 {
			balance += transactions[i]
			maxTr++
		}
	}
	return maxTr
}

func FindFistNonRepetableCharacter(s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range s {
		if m[r] == 1 {
			return string(r)
		}
	}
	return "_"
}

// example input = [1,2,3,2,1] remove pair of same number
// output = [2,1] first how many pair removed, second is the remaining number

//example input = []int{2, 1, 3, 2, 1, 2, 2, 3, 3}
//output = []int{4,1} first how many pair removed, second is the remaining number

func CountPairAndCheckRemaining(arr []int) []int {
	m := make(map[int]int)
	var count, remaning int
	for _, val := range arr {
		m[val]++
		if m[val]%2 == 0 {
			count++
			remaning--
		} else {
			remaning++
		}
	}

	return []int{count, remaning}
}

func ReverseNumber(n int) int {
	var rev int
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}
