package hackkerankproblem

func SlowestKey(keytimes [][]int32) string {
	// Write your code here
	var maxDuration int32
	var maxKey string
	for i := 0; i < len(keytimes); i++ {
		var duration int32
		if i == 0 {
			duration = keytimes[i][1]
		} else {
			duration = keytimes[i][1] - keytimes[i-1][1]
		}
		if duration > maxDuration {
			maxDuration = duration
			maxKey = string('a' + keytimes[i][0])
		}
	}
	return maxKey
}
