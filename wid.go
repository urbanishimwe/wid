package wid

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"unicode"
)

var (
	digits = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

// Generate Id in the format: {first word} {2 digits} {second word} {1 digit}.
//
// Error occurs if we can't read from cryptographically secure random number generator.
func Generate() (string, error) {
	id, err := GenerateUpper()
	return strings.ToLower(id), err
}

// GenerateUpper is like Generate but characters of words are uppercased randomly.
func GenerateUpper() (string, error) {

	// guider is a map to the following random properties in this order:
	// indices of words from list(4 bytes), capitalization flags(2 bytes), and 3 digits indices(3 byte).
	// giving more than 7776 * 10 * 10 * 7776 * 10 unique ids.
	var guider [9]byte
	_, err := rand.Read(guider[:])
	if err != nil {
		return "", err
	}

	firstWord := wordsList[binary.BigEndian.Uint16(guider[:2])%uint16(len(wordsList))]
	secondWord := wordsList[binary.BigEndian.Uint16(guider[2:4])%uint16(len(wordsList))]
	capitalizeFlags := binary.BigEndian.Uint16(guider[4:6])
	firstDigit := digits[int(guider[6])%len(digits)]
	secondDigit := digits[int(guider[7])%len(digits)]
	thirdDigit := digits[int(guider[8])%len(digits)]

	firstWord, secondWord = capitalize([]byte(firstWord), []byte(secondWord), capitalizeFlags)

	return fmt.Sprintf("%s%s%s%s%s", firstWord, firstDigit, secondDigit, secondWord, thirdDigit), nil
}

// capitalization is based on bit-set flag at a certain index.
func capitalize(firstWord, secondWord []byte, capitalizeFlags uint16) (string, string) {

	// capitalize first word
	for i := 0; i < len(firstWord); i++ {
		if flagSet(capitalizeFlags, i) {
			firstWord[i] = byte(unicode.ToUpper(rune(firstWord[i])))
		}
	}

	// capitalize second word
	capitalizeFlags >>= len(firstWord)
	for i := 0; i < len(secondWord); i++ {
		if flagSet(capitalizeFlags, i) {
			secondWord[i] = byte(unicode.ToUpper(rune(secondWord[i])))
		}
	}

	return string(firstWord), string(secondWord)
}

func flagSet(flags uint16, i int) bool {
	return math.MaxUint16&(flags&(1<<i)) != 0
}
