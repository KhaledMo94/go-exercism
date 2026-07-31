// Package proverb generates the proverbial rhyme from a list of inputs.
package proverb

import "fmt"

// Proverb returns the proverb lines for the given rhyme words.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}

	proverb := make([]string, 0, len(rhyme))

	for i := 0; i < len(rhyme)-1; i++ {
		proverb = append(proverb, fmt.Sprintf(
			"For want of a %s the %s was lost.",
			rhyme[i],
			rhyme[i+1],
		))
	}

	proverb = append(proverb, fmt.Sprintf(
		"And all for the want of a %s.",
		rhyme[0],
	))

	return proverb
}
