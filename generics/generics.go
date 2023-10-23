package main

// this is repetitive!!!
func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStrSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// solution
// Type parameters
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// Constraints in types
// allow us to write generics that operate on certain interface types
type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}
