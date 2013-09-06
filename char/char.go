package char

func IsAlphanumeric(b byte) bool {
	if '0' <= b && b <= '9' {
		return true
	}
	if 'A' <= b && b <= 'Z' {
		return true
	}
	if 'a' <= b && b <= 'z' {
		return true
	}
	return false
}
