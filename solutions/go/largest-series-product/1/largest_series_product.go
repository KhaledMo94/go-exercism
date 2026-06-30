package largestseriesproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	if span == 0 {
		return 1, nil
	}

	if span > len(digits) {
		return 0, errors.New("span must not exceed string length")
	}

	for _, ch := range digits {
		if ch < '0' || ch > '9' {
			return 0, errors.New("digits input must only contain digits")
		}
	}

	var max int64

	for start := 0; start <= len(digits)-span; start++ {
		product := int64(1)

		for i := start; i < start+span; i++ {
			product *= int64(digits[i] - '0')
		}

		if product > max {
			max = product
		}
	}

	return max, nil
}