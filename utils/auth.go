package utils

import (
	"errors"
	"os"
	"strings"
)

const tokenEnvVar = "PWCTL_TOKEN"

func GetToken() (string, error) {
	token := os.Getenv(tokenEnvVar)
	if token == "" {
		return "", errors.New("token not found in environment variable")
	}

	return token, nil
}

// StoreToken stores the token to the environment variable.
func StoreToken(token string) error {

	err := os.Setenv(tokenEnvVar, token)
	if err != nil {
		return err
	}

	return nil
}

// Implementing the system variable feature will need some interactions with the mentor
// as it has pros and cons of its own. So just for the sake to implement other commands
// I have also implemented the token storage in a file.
const tokenFile = "./utils/token.txt"

func GetTokenFile() (string, error) {
	tokenBytes, err := os.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}

	token := strings.TrimSpace(string(tokenBytes))
	if token == "" {
		return "", errors.New("token not found in file")
	}

	return token, nil
}

// StoreToken stores the token to the file.
func StoreTokenFile(token string) error {
	err := os.WriteFile(tokenFile, []byte(token), 0644)
	if err != nil {
		return err
	}

	return nil
}
