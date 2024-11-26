package hackkerankeasy

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1 > y2 {
		return 10000
	}
	if y1 == y2 && m1 > m2 {
		return (m1 - m2) * 500
	}
	if y1 == y2 && m1 == m2 && d1 > d2 {
		return (d1 - d2) * 15
	}
	return 0
}
