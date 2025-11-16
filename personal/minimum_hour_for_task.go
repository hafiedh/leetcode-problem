package personal

func MinimumHourForTask(tasks []int, gap int) int {
	lastTime := make(map[int]int)
	hours := 0
	for _, task := range tasks {
		if last, ok := lastTime[task]; ok && hours-last <= gap {
			hours = last + gap
		}
		hours++
		lastTime[task] = hours
	}
	return hours
}

// MinimumHourForTaskRearrange menghitung jam minimum untuk menyelesaikan semua tugas dengan jeda (gap), dengan memperbolehkan penataan ulang tugas.
func MinimumHourForTaskRearrange(tasks []int, gap int) int {
	if len(tasks) == 0 {
		return 0
	}
	// Hitung frekuensi setiap tugas
	taskCount := make(map[int]int)
	maxFreq := 0
	for _, t := range tasks {
		taskCount[t]++
		if taskCount[t] > maxFreq {
			maxFreq = taskCount[t]
		}
	}
	// Hitung berapa banyak tugas yang memiliki frekuensi maksimum
	maxCount := 0
	for _, v := range taskCount {
		if v == maxFreq {
			maxCount++
		}
	}
	// Rumus: (maxFreq-1)*(gap+1) + maxCount
	minTime := (maxFreq-1)*(gap+1) + maxCount
	if minTime < len(tasks) {
		return len(tasks)
	}
	return minTime
}
