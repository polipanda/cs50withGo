package week1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

const AMEX = "AMEX"
const INVALID = "INVALID\n"
const MASTERCARD = "MASTERCARD\n"
const VISA = "VISA\n"

func ValidateCreditCardNumbers(number int) string {
	if validCheckSum(number) != nil {
		return INVALID
	}

	return determineCreditCardCarrier(number)
}

func validCheckSum(number int) error {
	sNum := strconv.Itoa(number)

	isEven := len(sNum)%2 == 0
	sum := 0
	for i, digit := range sNum {
		if isEven {
			if i%2 == 0 {
				d := int(digit-'0') * 2
				if d > 10 {
					d -= 9
				}
				sum += d
			} else {
				sum += int(digit - '0')
			}
		} else {
			if i%2 == 0 {
				sum += int(digit - '0')
			} else {
				d := int(digit-'0') * 2
				if d > 10 {
					d -= 9
				}
				sum += d
			}
		}
	}

	if sum%10 != 0 {
		return fmt.Errorf("invalid check sum")
	}

	return nil
}

func determineCreditCardCarrier(number int) string {
	numLength := intLength(number)
	validAMEX := []int{34, 37}
	if slices.Contains(validAMEX, firstNDigits(number, 2)) && numLength == 15 {
		return AMEX
	}

	validMasterCard := []int{22}
	if slices.Contains(validMasterCard, firstNDigits(number, 2)) && numLength == 16 {
		return MASTERCARD
	}

	validVISA := []int{4}
	if slices.Contains(validVISA, firstNDigits(number, 1)) && (numLength == 13 || numLength == 16) {
		return VISA
	}

	return INVALID
}

func intLength(number int) int {
	if number == 0 {
		return 1
	}
	count := 0
	for number != 0 {
		number /= 10
		count++
	}

	return count
}

func firstNDigits(number int, digits int) int {
	var i int
	for i = number; i >= int(math.Pow10(digits)); i = i / 10 {
	}

	return i
}
