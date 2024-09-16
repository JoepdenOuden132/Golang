package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var lowercase = "abcdefghijklmnopqrstuvwxyz"
var uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digits = "0123456789"
var special = "!@#$%^&*()_+-=[]{}|;:',.<>/?"

func main() {
	useLowercase, err := askYesandNo("Do you want lowercase letters? (yes or no)")
	if err != nil {
		fmt.Println(err)
		return
	}

	useUppercase, err := askYesandNo("Do you want uppercase letters? (yes or no)")
	if err != nil {
		fmt.Println(err)
		return
	}

	useDigits, err := askYesandNo("Do you want digits? (yes or no)")
	if err != nil {
		fmt.Println(err)
		return
	}

	useSpecial, err := askYesandNo("Do you want special characters? (yes or no)")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User wants lowercase letters:", useLowercase)
	fmt.Println("User wants uppercase letters:", useUppercase)
	fmt.Println("User wants digits:", useDigits)
	fmt.Println("User wants special characters:", useSpecial)

	length, err := askPasswordLength()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Password length:", length)

	var characterPool string
	if useLowercase {
		characterPool += "abcdefghijklmnopqrstuvwxyz"
	}
	if useUppercase {
		characterPool += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useSpecial {
		characterPool += "0123456789"
	}
	if useDigits {
		characterPool += "!@#$%^&*()_+-=[]{}|;:',.<>/?"
	}

	if len(characterPool) == 0 {
		fmt.Println("No character types selected. Cannot generate a password.")
		return
	}

	password, err := passwordGenerator(length, characterPool)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println(password)
}

func askYesandNo(question string) (bool, error) {
	var response string

	fmt.Println(question)
	_, err := fmt.Scanln(&response)

	if err != nil {
		return false, err
	}
	response = strings.ToLower(response)

	if response == "yes" || response == "y" {
		return true, nil
	} else if response == "no" || response == "n" {
		return false, nil
	} else {
		return false, fmt.Errorf("invalid input: please enter 'yes' or 'no'")
	}

}

func askPasswordLength() (int, error) {
	var lengthStr string
	fmt.Println("how long do you want your password to be? (minimal lenght of 8)")

	_, err := fmt.Scanln(&lengthStr)

	if err != nil {
		return 0, err
	}

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return 0, fmt.Errorf("Error, something went wrong")
	}

	if length < 8 {
		return 0, fmt.Errorf("Error, something went wrong")
	}

	return length, nil

}

func passwordGenerator(lenght int, characterpool string) (string, error) {
	fmt.Println("Generating Password!!")
	var password strings.Builder
	for i := 0; i < lenght; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(characterpool))))
		if err != nil {
			return "", err
		}
		password.WriteByte(characterpool[index.Int64()])
	}
	return password.String(), nil

}
