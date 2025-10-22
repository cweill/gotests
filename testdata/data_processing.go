package testdata

// FilterPositive returns only positive numbers from the slice
func FilterPositive(numbers []int) []int {
	if numbers == nil {
		return nil
	}

	if len(numbers) == 0 {
		return []int{}
	}

	result := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if num > 0 {
			result = append(result, num)
		}
	}

	return result
}

// GroupByLength groups strings by their length
func GroupByLength(words []string) map[int][]string {
	if words == nil {
		return nil
	}

	result := make(map[int][]string)

	for _, word := range words {
		length := len(word)
		result[length] = append(result[length], word)
	}

	return result
}

// Deduplicate removes duplicate strings while preserving order
func Deduplicate(items []string) []string {
	if items == nil {
		return nil
	}

	if len(items) == 0 {
		return []string{}
	}

	seen := make(map[string]bool)
	result := make([]string, 0, len(items))

	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// SumByKey sums integer values grouped by string keys
func SumByKey(data map[string][]int) map[string]int {
	if data == nil {
		return nil
	}

	result := make(map[string]int)

	for key, values := range data {
		sum := 0
		for _, val := range values {
			sum += val
		}
		result[key] = sum
	}

	return result
}

// MergeUnique merges two slices and removes duplicates
func MergeUnique(a, b []string) []string {
	if a == nil && b == nil {
		return nil
	}

	seen := make(map[string]bool)
	result := make([]string, 0, len(a)+len(b))

	// Add from first slice
	for _, item := range a {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	// Add from second slice
	for _, item := range b {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Partition splits a slice into two based on a predicate
func Partition(numbers []int, threshold int) (below, aboveOrEqual []int) {
	if numbers == nil {
		return nil, nil
	}

	below = make([]int, 0)
	aboveOrEqual = make([]int, 0)

	for _, num := range numbers {
		if num < threshold {
			below = append(below, num)
		} else {
			aboveOrEqual = append(aboveOrEqual, num)
		}
	}

	return below, aboveOrEqual
}
