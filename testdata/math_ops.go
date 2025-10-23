package testdata

import "errors"

// Clamp restricts a value to be within a specified range
func Clamp(value, min, max int) int {
	if min > max {
		// Swap if min and max are reversed
		min, max = max, min
	}

	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Average calculates the average of a slice of integers
func Average(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot calculate average of empty slice")
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return float64(sum) / float64(len(numbers)), nil
}

// Factorial calculates the factorial of n
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial not defined for negative numbers")
	}

	if n > 20 {
		return 0, errors.New("factorial too large (max 20)")
	}

	if n == 0 || n == 1 {
		return 1, nil
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result, nil
}

// GCD calculates the greatest common divisor using Euclidean algorithm
func GCD(a, b int) int {
	// Handle negative numbers
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Handle zero cases
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check for divisors up to sqrt(n)
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

// AbsDiff returns the absolute difference between two integers
func AbsDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}
