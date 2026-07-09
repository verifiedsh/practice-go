package main

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	prevUpper := false
	for i, r := range s {
		isLower, isUpper := r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z'
		if !isLower && !isUpper || isUpper && prevUpper {
			return s
		}
		if isUpper && i != 0 {
			result += "_"
		}
		if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
			return s
		}
		result += string(r)
	}
	return result
}
