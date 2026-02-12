package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func generatePassword(length int) string {
	LowercaseChars := "abcdefghijklmnopqrstuvwxyz"
	UppercaseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits := "0123456789"
	SpecialChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	allChars := LowercaseChars + UppercaseChars + Digits + SpecialChars

	fmt.Print("Go Password Generator\n")
	fmt.Print(strings.Repeat("=", 60))
	fmt.Print("\nChoose from the following options:\n")
	fmt.Print("1. Generate a Random Password.\n")
	fmt.Print("2. Generate a Password of Specific Length\n")
	fmt.Print("3. Generate a Password from a Specific Character Set\n")
	fmt.Print("4. Generate a Password Without Any Lowercase Letters\n")
	fmt.Print("5. Generate a Password Without Any Uppercase Letters\n")
	fmt.Print("6. Generate a Password Without Any Digits\n")
	fmt.Print("7. Generate a Password Without Any Special Characters\n")
	fmt.Print("Enter your choice (1-7): ")
	var userChoice int
	fmt.Scanln(&userChoice)

	switch userChoice {
	// Generate Random Passowrd
	case 1:
		return generatePasswordFromCharSet(length, allChars)
	// Generate Random Password of Specified Length
	case 2:
		fmt.Print("Enter the desired password length: ")
		fmt.Scanln(&length)
		return generatePasswordFromCharSet(length, allChars)
	// Generate Random Password Based on User Typed Characters
	case 3:
		fmt.Print("Enter the character set to use for password generation: ")
		var charSet string
		fmt.Scanln(&charSet)
		return generatePasswordFromCharSet(length, charSet)
	// Generate Random Password Without Lowercase Letters
	case 4:
		return generatePasswordFromCharSet(length, UppercaseChars+Digits+SpecialChars)
	// Generate Random Password Without Uppercase Letters
	case 5:
		return generatePasswordFromCharSet(length, LowercaseChars+Digits+SpecialChars)
	// Generate Random Password Without Digits
	case 6:
		return generatePasswordFromCharSet(length, LowercaseChars+UppercaseChars+SpecialChars)
	// Generate Random Password Without Special Characters
	case 7:
		return generatePasswordFromCharSet(length, LowercaseChars+UppercaseChars+Digits)
	// Message if Input is Invalid
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
		return ""
	}
}

func generatePasswordFromCharSet(length int, charSet string) string {
	var password strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charSet))
		password.WriteByte(charSet[randomIndex])
	}
	return password.String()
}

func main() {
	// Minimum Default Password Length of 8, Maximum Default Password Length of 16
	password := generatePassword(rand.Intn(9) + 8)
	fmt.Printf("\nGenerated Password: %s\n", password)
}
