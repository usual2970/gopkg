package tel

import "regexp"

var wildcardReg = regexp.MustCompile(`(\d{3})\d{4}(\d{4})`)

func TrimRegion(tel string) string {
	return tel[4:]
}

func Wildcard(tel string) string {
	placeholder := "****"
	l := len(tel)
	suffix := tel[l-4:]

	prefix := ""
	if l > 8 {
		prefix = tel[:l-8]
	}

	return prefix + placeholder + suffix
}
