package utils

// StrLimit reduces string length to a certain length
func StrLimit(str string, limit int) string {
	if len(str) <= limit {
		return str
	}

	return string([]byte(str)[:limit]) + "..."
}
