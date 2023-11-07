package goternary

import "reflect"

// An alternative function for ternary operation
func Ternary(condition bool, leftValue, rightValue any) any {
	if condition {
		return leftValue
	}
	return rightValue
}

// An alternative function for ternary operation, returns value if pointer is received
func TernaryValue(condition bool, leftValue, rightValue any) any {
	result := Ternary(condition, leftValue, rightValue)

	return getValueFromPointer(result)
}

func getValueFromPointer(v any) any {
	// Use reflection to dereference a pointer and return the value
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		return value.Elem().Interface()
	}
	return v
}
