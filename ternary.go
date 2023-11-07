package goternary

// An alternative function for ternary operation
func Ternary(condition bool, leftValue, rightValue any) any {
	if condition {
		return leftValue
	}
	return rightValue
}
