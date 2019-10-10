// Common utils (Contains()/Filter()/..) for common types of slices
package sliceutil

// Go really should have generics..

func ContainsString(items []string, item string) bool {
	for _, candidate := range items {
		if candidate == item {
			return true
		}
	}

	return false
}

func FilterString(items []string, fn func(item string) bool) []string {
	altered := []string{}

	for _, item := range items {
		if fn(item) {
			altered = append(altered, item)
		}
	}

	return altered
}

func ContainsInt(items []int, item int) bool {
	for _, candidate := range items {
		if candidate == item {
			return true
		}
	}

	return false
}

func FilterInt(items []int, fn func(item int) bool) []int {
	altered := []int{}

	for _, item := range items {
		if fn(item) {
			altered = append(altered, item)
		}
	}

	return altered
}
