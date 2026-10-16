package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)

	n := len(runes)
	if n == 0 {
		return s
	}

	shift %= n
	if shift < 0 {
		shift += n
	}

	result := make([]rune, n)

	for i, r := range runes {
		newIndex := i - shift
		if newIndex < 0 {
			newIndex += n
		}
		result[newIndex] = r
	}

	return string(result)
}
