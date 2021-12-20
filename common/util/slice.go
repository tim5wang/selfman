package util

func IsInSlice(ss []string, des string) bool {
	for _, s := range ss {
		if s == des {
			return true
		}
	}
	return false
}
