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
