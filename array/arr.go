package array

// Filter function takes a slice of elements and a callback function
// that accepts an element of the slice and returns a boolean value.
// It filters the input slice based on the provided callback, retaining
// only the elements for which the callback returns true.
//
// Parameters:
//   - arr: A slice of elements of type K that needs to be filtered.
//   - cb: A callback function that takes an element of type K as its
//     parameter and returns a boolean value indicating whether
//     the element should be included in the filtered result.
//
// Returns:
//   - []K: A new slice containing only the elements from the input
//     slice for which the callback function returned true.
//
// Example:
//
//	inputSlice := []int{1, 2, 3, 4, 5}
//	isEven := func(n int) bool { return n%2 == 0 }
//	result := Filter(inputSlice, isEven)
//	// result will be []int{2, 4} since only the even numbers
//	// satisfy the filtering condition.
//
// Note:
//   - The type parameter K must be a comparable type (implements
//     equality and ordering operations) for the function to work.
//   - The original input slice 'arr' remains unchanged, and the
//     function returns a new slice with the filtered elements.
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

// Push function adds one or more elements to the end of a slice.
// It modifies the original slice in place, updating its length and
// capacity as needed.
//
// Parameters:
//   - arr: A pointer to a slice of elements of type K that needs to
//     be extended by appending new elements.
//   - elem: One or more elements of type K to be added to the end of
//     the slice.
//
// Returns:
//
//	None
//
// Example:
//
//	originalSlice := []int{1, 2, 3}
//	Push(&originalSlice, 4, 5)
//	// originalSlice will be modified to []int{1, 2, 3, 4, 5}
//
// Note:
//   - The type parameter K must be a comparable type (implements
//     equality and ordering operations) for the function to work.
//   - The function updates the original slice in place and does not
//     return a new slice.
func Push[K comparable](arr *[]K, elem ...K) {
	newArr := make([]K, len(*arr), len(*arr)+len(elem))

	for _, item := range elem {
		newArr = append(*arr, item)
	}
	*arr = newArr
}
