package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	codes := map[string]int{
		"black": 0,
		"brown": 1,
		"red": 2,
		"orange":3,
		"yellow":4,
		"green": 5,
		"blue": 6,
		"violet":7,
		"grey": 8,
		"white": 9,
	}

	if len(colors) == 0 {
		return 0
	}

	if len(colors) == 1 {
		return codes[colors[0]]
	}

	return 10*codes[colors[0]] + codes[colors[1]]
}
