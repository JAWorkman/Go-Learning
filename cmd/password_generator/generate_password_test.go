package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const testLength = 50

// userInput replaces os.Stdin with a pipe containing the given input.
// Returns a cleanup function that restores the original stdin.
func simulateUserInput(input string) func() {
	originalStdin := os.Stdin
	r, w, _ := os.Pipe()
	_, err := w.WriteString(input)
	if err != nil {
		fmt.Printf("Error writing to pipe: %v\n", err)
	}

	err = w.Close()

	if err != nil {
		fmt.Printf("Error closing write end of pipe: %v\n", err)
	}

	os.Stdin = r
	return func() { os.Stdin = originalStdin }
}

// containsOnly checks that every character in s exists in allowed.
func containsOnly(s, allowed string) bool {
	for _, ch := range s {
		if !strings.ContainsRune(allowed, ch) {
			return false
		}
	}
	return true
}

// containsNone checks that no character in s exists in forbidden.
func containsNone(s, forbidden string) bool {
	for _, ch := range s {
		if strings.ContainsRune(forbidden, ch) {
			return false
		}
	}
	return true
}

func assertContainsOnly(t *testing.T, password, allowed string) {
	t.Helper()
	if !containsOnly(password, allowed) {
		t.Errorf("password contains unexpected characters: %s", password)
	}
}

func assertContainsNone(t *testing.T, password, forbidden, label string) {
	t.Helper()
	if !containsNone(password, forbidden) {
		t.Errorf("password should not contain %s: %s", label, password)
	}
}

func assertLength(t *testing.T, password string, expected int) {
	t.Helper()
	if len(password) != expected {
		t.Errorf("expected password of length %d, got password of length %d.", expected, len(password))
	}
}

func TestGeneratePasswordFromCharSet(t *testing.T) {
	t.Run("returns password of requested length", func(t *testing.T) {
		password := generatePasswordFromCharSet(20, AllChars)
		assertLength(t, password, 20)
	})

	t.Run("returns empty string for zero length", func(t *testing.T) {
		password := generatePasswordFromCharSet(0, AllChars)
		if password != "" {
			t.Errorf("expected empty string, got %q", password)
		}
	})

	t.Run("only uses characters from the given set", func(t *testing.T) {
		charSet := "abc"
		password := generatePasswordFromCharSet(100, charSet)
		assertContainsOnly(t, password, charSet)
	})
}

func TestGeneratePassword(t *testing.T) {
	t.Run("option 1: random password with all characters", func(t *testing.T) {
		userInput := simulateUserInput("1\n")
		defer userInput()

		length := 12
		password := generatePassword(length)

		assertLength(t, password, length)
		assertContainsOnly(t, password, AllChars)
	})

	t.Run("option 2: password with custom length", func(t *testing.T) {
		customLength := 25
		userInput := simulateUserInput(fmt.Sprintf("2\n %d\n", customLength))
		defer userInput()
		password := generatePassword(customLength) // initial length ignored, replaced by user input 25

		assertLength(t, password, customLength)
		assertContainsOnly(t, password, AllChars)
	})

	t.Run("option 3: password from custom character set", func(t *testing.T) {
		userInput := simulateUserInput("3\nXYZ123\n")
		defer userInput()

		password := generatePassword(testLength)

		assertLength(t, password, testLength)
		assertContainsOnly(t, password, "XYZ123")
	})

	t.Run("option 4: no lowercase letters", func(t *testing.T) {
		userInput := simulateUserInput("4\n")
		defer userInput()

		password := generatePassword(testLength)

		assertLength(t, password, testLength)
		assertContainsNone(t, password, LowercaseChars, "lowercase letters")
	})

	t.Run("option 5: no uppercase letters", func(t *testing.T) {
		userInput := simulateUserInput("5\n")
		defer userInput()

		password := generatePassword(testLength)

		assertLength(t, password, testLength)
		assertContainsNone(t, password, UppercaseChars, "uppercase letters")
	})

	t.Run("option 6: no digits", func(t *testing.T) {
		userInput := simulateUserInput("6\n")
		defer userInput()

		password := generatePassword(testLength)

		assertLength(t, password, testLength)
		assertContainsNone(t, password, Digits, "digits")
	})

	t.Run("option 7: no special characters", func(t *testing.T) {
		userInput := simulateUserInput("7\n")
		defer userInput()

		password := generatePassword(testLength)

		assertLength(t, password, testLength)
		assertContainsNone(t, password, SpecialChars, "special characters")
	})

	t.Run("invalid choice retries until valid", func(t *testing.T) {
		userInput := simulateUserInput("9\n0\n1\n") // two invalid choices, then option 1
		defer userInput()

		password := generatePassword(testLength)

		assertLength(t, password, testLength)
		assertContainsOnly(t, password, AllChars)
	})
}
