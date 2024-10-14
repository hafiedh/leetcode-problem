package hackkerankproblem

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	var result int64
	if bc > wc+z {
		result = int64(wc)*int64(b+w) + int64(z)*int64(b)
	} else if wc > bc+z {
		result = int64(bc)*int64(b+w) + int64(z)*int64(w)
	} else {
		result = int64(bc)*int64(b) + int64(wc)*int64(w)
	}
	return result
}
