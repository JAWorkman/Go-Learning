package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	LowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	UppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits         = "0123456789"
	SpecialChars   = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	AllChars       = LowercaseChars + UppercaseChars + Digits + SpecialChars
)

// generatePassword displays a menu to the user and returns a password based on their selection.
// If the user enters an invalid choice, the menu is displayed again.
func generatePassword(length int) string {
	for {
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
		if _, err := fmt.Scanln(&userChoice); err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 7.")
			continue
		}

		// Generate password based on user's choice
		switch userChoice {
		case 1:
			// Use all characters with the default length
			return generatePasswordFromCharSet(length, AllChars)
		case 2:
			// Prompt user for a custom length
			fmt.Print("Enter the desired password length: ")
			if _, err := fmt.Scanln(&length); err != nil {
				fmt.Println("Invalid input. Using default length.")
			}
			return generatePasswordFromCharSet(length, AllChars)
		case 3:
			// Prompt user for a custom character set
			fmt.Print("Enter the character set to use for password generation: ")
			var charSet string
			if _, err := fmt.Scanln(&charSet); err != nil {
				fmt.Println("Invalid input. Using default character set.")
				charSet = AllChars
			}
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
		}
	}
}

// generatePasswordFromCharSet builds a random password of the given length
// by picking random characters from the provided character set.
func generatePasswordFromCharSet(length int, charSet string) string {
	var password strings.Builder
	for i := 0; i < length; i++ {
		// Pick a random character from the character set
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		// Append that character to the password
		password.WriteByte(charSet[randomIndex.Int64()])
	}
	return password.String()
}

func main() {
	// Generate a password with a random default length between 8 and 16
	randomLength, _ := rand.Int(rand.Reader, big.NewInt(9))
	password := generatePassword(int(randomLength.Int64()) + 8) // random length between 8 and 16
	fmt.Printf("Generated Password: %s\n", password)
}
