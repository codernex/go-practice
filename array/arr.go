package array

func Filter[K comparable](arr []K, cb func(item K) bool) []K {
	newArr := []K{}
	for _, item := range arr {
		if cb(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

func Map[K comparable](arr []K, cb func(item K, index int, array []K) any) []K {
	newArr := []K{}

	for i, item := range arr {
		newArr = append(newArr, cb(item, i, arr))
	}

	return newArr
}

func Push[K comparable](arr *[]K, elem ...K) {
	newArr := make([]K, len(*arr), len(*arr)*3)

	for _, item := range elem {
		newArr = append(*arr, item)
	}
	*arr = newArr

}
