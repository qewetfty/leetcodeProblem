package utils

func Reverse(s string) string {
	l := len(s)
	b := []byte(s)
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	return string(b)
}
