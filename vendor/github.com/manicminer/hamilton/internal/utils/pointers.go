package utils

// BoolPtr returns a pointer to the provided boolean variable.
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtr returns a pointer to the provided string variable.
func StringPtr(s string) *string {
	return &s
}
