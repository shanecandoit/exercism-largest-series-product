package lsproduct

import (
	"errors"
	"fmt"
)

/*
Given a string of digits, calculate the largest product for a
contiguous substring of digits of length n.

For example, for the input '1027839564', the largest product for a series of
3 digits is 270 (9 * 5 * 6), and the largest product for a series of
5 digits is 7560 (7 * 8 * 3 * 9 * 5).

Note that these series are only required to occupy adjacent positions
in the input; the digits need not be numerically consecutive.

For the input '73167176531330624919225119674426574742355349194934',
the largest product for a series of 6 digits is 23520.
*/

// LargestSeriesProduct find the largest product of a contiguous set of digits of span length
func LargestSeriesProduct(digits string, span int) (int, error) {
	fmt.Println("LargestSeriesProduct", digits, span)
	// var a,b,c =0,0,0
	// maxProd := 0
	fmt.Println("len digits", len(digits), "span", span)

	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}

	maxProd := -1
	// fmt.Println("len(digits)-span", len(digits)-span)
	for i := 0; i < len(digits)-span+1; i++ {
		chunk := digits[i : i+span]
		// fmt.Println("check", chunk)

		prod := 1
		for _, ch := range chunk {
			digit := int(ch) - int('0')
			if digit < 0 || digit > 9 {
				return -1, errors.New("digits input must only contain digits")
			}
			// fmt.Println(".", j, ch)

			prod *= digit
		}
		if prod > maxProd {
			fmt.Println("new maxProd", maxProd, "from", chunk)
			maxProd = prod
		}
	}

	fmt.Println("LargestSeriesProduct =", maxProd)
	return maxProd, nil
}
