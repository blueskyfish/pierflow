package helper

// IIF is a utility function that returns one of two values based on a condition.
func IIF[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
