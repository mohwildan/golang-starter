package helpers

func InArrayString(s string, haystack []string) bool {
	for _, str := range haystack {
		if s == str {
			return true
		}
	}
	return false
}
