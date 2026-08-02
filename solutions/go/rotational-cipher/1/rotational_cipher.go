package rotationalcipher

func shiftLetter(b byte, shiftKey int) byte {
	var base byte
	switch {
	case b >= 'A' && b <= 'Z':
		base = 'A'
	case b >= 'a' && b <= 'z':
		base = 'a'
	default:
		return b
	}
	return base + byte((int(b-base)+shiftKey)%26)
}

func RotationalCipher(plain string, shiftKey int) string {
	out := []byte(plain)
	for i, b := range out {
		out[i] = shiftLetter(b, shiftKey)
	}
	return string(out)
}
