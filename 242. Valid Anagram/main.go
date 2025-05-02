func insertionSort(s []byte) {
	for i := 1; i < len(s); i++ {
		temp := s[i]
		j := i - 1
		for j >= 0 && s[j] > temp {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = temp
	}
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sBytes := []byte(s)
	tBytes := []byte(t)

	insertionSort(sBytes)
	insertionSort(tBytes)

	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}

	return true
}
