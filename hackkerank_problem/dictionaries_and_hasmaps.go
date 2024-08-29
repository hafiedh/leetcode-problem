package hackkerankproblem

import "fmt"

// Harold is a kidnapper who wrote a ransom note, but now he is worried it will be traced back to him through his handwriting. He found a magazine and wants to know if he can cut out whole words from it and use them to create an untraceable replica of his ransom note. The words in his note are case-sensitive and he must use only whole words available in the magazine. He cannot use substrings or concatenation to create the words he needs.

// Given the words in the magazine and the words in the ransom note, print Yes if he can replicate his ransom note exactly using whole words from the magazine; otherwise, print No.

func checkMagazine(magazine []string, note []string) {
	magazineMap := make(map[string]int)
	for _, word := range magazine {
		magazineMap[word]++
	}

	for _, word := range note {
		if magazineMap[word] == 0 {
			fmt.Println("No")
			return
		}
		magazineMap[word]--
	}

	fmt.Println("Yes")
}
