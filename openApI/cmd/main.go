package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// asker := &pkg.Asker{Model: openai.GPT3Dot5Turbo0301}
	// asker.CreateClient()
	// asker.Req = asker.CreateCompletionRequest("Tell me the recipe of a milkashake", nil)
	// asker.GetChatCompletitionStream()
	// asker.PrintResults()
	final_value := 0
	for {
		var input string
		fmt.Scanln(&input)

		final_value += Find_Value(input)
		fmt.Println("final value is", final_value)
	}

}

func Read_line() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	return input
}

func wordToDigit(word string) int {
	switch word {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1 // indicates not a valid spelled-out number
	}
}

func extractDigits(input string) (int, int) {
	numWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var firstDigit, lastDigit int
	var foundFirst, foundLast bool

	// Check for the first digit or spelled-out number
	for i, r := range input {
		if unicode.IsDigit(r) {
			firstDigit = int(r - '0')
			foundFirst = true
			break
		} else {
			for _, word := range numWords {
				if strings.HasPrefix(input[i:], word) {
					firstDigit = wordToDigit(word)
					foundFirst = true
					break
				}
			}
		}
		if foundFirst {
			break
		}
	}

	// Check for the last digit or spelled-out number in reverse order
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(input[i])) {
			lastDigit = int(input[i] - '0')
			foundLast = true
			break
		} else {
			for _, word := range numWords {
				if i >= len(word)-1 && input[i-len(word)+1:i+1] == word {
					lastDigit = wordToDigit(word)
					foundLast = true
					break
				}
			}
		}
		if foundLast {
			break
		}
	}

	return firstDigit, lastDigit
}

// Updated Find_Value function
func Find_Value(input string) int {
	firstDigit, lastDigit := extractDigits(input)
	return firstDigit*10 + lastDigit
}
