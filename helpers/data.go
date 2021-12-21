package helpers

func ContainsInt(slice []int, value int) int {
	for i, v := range slice {
		if value == v {
			return i
		}
	}
	return -1
}

func ContainsString(slice []string, value string) int {
	for i, v := range slice {
		if value == v {
			return i
		}
	}
	return -1
}
