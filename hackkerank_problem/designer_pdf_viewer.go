package hackkerankproblem

func DesignerPdfViewer(h []int32, word string) int32 {
	var max int32
	for _, c := range word {
		if h[c-97] > max {
			max = h[c-97]
		}
	}
	return max * int32(len(word))
}

func DesignerPdfViewer2(h []int32, word string) int32 {
	var max int32
	mDictionary := map[string]int32{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
	}
	for _, v := range word {
		if _, ok := mDictionary[string(v)]; ok {
			if max < h[mDictionary[string(v)]] {
				max = h[mDictionary[string(v)]]
			}
		}
	}
	return max * int32(len(word))
}
