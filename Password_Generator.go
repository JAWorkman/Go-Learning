package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// generatePassword displays a menu to the user and returns a password based on their selection.
func generatePassword(length int) string {
	// Define character sets for password generation
	LowercaseChars := "abcdefghijklmnopqrstuvwxyz"
	UppercaseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits := "0123456789"
	SpecialChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	allChars := LowercaseChars + UppercaseChars + Digits + SpecialChars

	// Display menu options to the user
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

	// Read the user's menu selection
	var userChoice int
	fmt.Scanln(&userChoice)

	// Generate password based on user's choice
	switch userChoice {
	case 1:
		// Use all characters with the default length
		return generatePasswordFromCharSet(length, allChars)
	case 2:
		// Prompt user for a custom length
		fmt.Print("Enter the desired password length: ")
		fmt.Scanln(&length)
		return generatePasswordFromCharSet(length, allChars)
	case 3:
		// Prompt user for a custom character set
		fmt.Print("Enter the character set to use for password generation: ")
		var charSet string
		fmt.Scanln(&charSet)
		return generatePasswordFromCharSet(length, charSet)
	case 4:
		// Exclude lowercase letters
		return generatePasswordFromCharSet(length, UppercaseChars+Digits+SpecialChars)
	case 5:
		// Exclude uppercase letters
		return generatePasswordFromCharSet(length, LowercaseChars+Digits+SpecialChars)
	case 6:
		// Exclude digits
		return generatePasswordFromCharSet(length, LowercaseChars+UppercaseChars+SpecialChars)
	case 7:
		// Exclude special characters
		return generatePasswordFromCharSet(length, LowercaseChars+UppercaseChars+Digits)
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
		return ""
	}
}

// generatePasswordFromCharSet builds a random password of the given length
// by picking random characters from the provided character set.
func generatePasswordFromCharSet(length int, charSet string) string {
	var password strings.Builder
	for i := 0; i < length; i++ {
		// Pick a random character from the character set
		randomIndex := rand.Intn(len(charSet))
		// Append that character to the password
		password.WriteByte(charSet[randomIndex])
	}
	return password.String()
}

func main() {
	// Generate a password with a random default length between 8 and 16
	password := generatePassword(rand.Intn(9) + 8)
	fmt.Printf("\nGenerated Password: %s\n", password)
}
