package hackkerankeasy

import "strconv"

func SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	res := make([]string, 0)
	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			if start == end {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			start = nums[i]
			end = nums[i]
		}
	}
	if start == end {
		res = append(res, strconv.Itoa(start))
	} else {
		res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
	return res
}
