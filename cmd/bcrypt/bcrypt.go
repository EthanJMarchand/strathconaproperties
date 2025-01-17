package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func hash(password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing password: %s\n", err)
	}
	fmt.Println(string(bytes))
}

func Compare(hash, password string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("Failed to pass with error: %s\n", err)
		return
	}
	fmt.Println("Password matches")
}

func main() {
	args := os.Args
	switch args[1] {
	case "hash":
		hash(args[2])
	case "compare":
		Compare(args[2], args[3])
	default:
		fmt.Printf("Invalid Command: %s\n", args[1])
		return
	}

}
