package game

func compareSlices(in []int, out []int) (bool, int) {
	for index := 0; index < len(in); index++ {
		if in[index] != out[index] {
			return false, index
		}
	}
	return true, 0
}

// Find  takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
