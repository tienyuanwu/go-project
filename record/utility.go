package record

func mapHexagram(table [6]float64, data []float64) (int, int) {
	above := 0
	below := 0
	for i := 0; i < 6; i++ {
		line := 1
		if data[i] > table[i] {
			line = 0
		}

		if i < 3 {
			below |= line << uint(2-i)
		} else {
			above |= line << uint(5-i)
		}
	}
	return above, below
}
