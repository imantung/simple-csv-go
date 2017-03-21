package csv2

// RemoveDoubleQuote remove double quote in string
func RemoveDoubleQuote(s string) string {
	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}
