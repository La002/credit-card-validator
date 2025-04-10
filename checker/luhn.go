package checker

import (
	"fmt"
	"strconv"
)

func CheckNumberValid(n string) bool {
	if len(n) < 13 || len(n) > 16 {
		fmt.Println("Number length not valid")
		return false
	}
	var total int
	reverseIndex := len(n) - 1
	for i, digitRune := range n {
		digit, err := strconv.Atoi(string(digitRune))
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return false
		}
		if (reverseIndex-i)%2 == 1 { // Double every second digit from the right
			digit *= 2
			if digit > 9 {
				digit = digit/10 + digit%10
			}
		}
		total += digit
	}
	if total%10 != 0 {
		return false
	}
	return true
}

func GetNetwork(n string) string {
	switch string(n[0]) {
	case "4":
		return "Visa"
	case "5":
		return "Master"
	case "6":
		return "Discover - Discover is tied up with RuPay in India"
	case "3":
		if string(n[1]) == "7" {
			return "American Express"
		}
	}
	return ""
}
