package go_practice

func Filter[K comparable](arr []K, cb func(item K) bool) []K {
	newArr := []K{}
	for _, item := range arr {
		if cb(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}
