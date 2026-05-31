package cli

import (
	"fmt"
	"syscall"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"golang.org/x/term"
)

var myEnv struct {
	AuthSecret string `env:"AUTH_SECRET"`
	ServerURL  string `env:"SERVER_URL"`
}

func init() {
	_ = godotenv.Load()
	_ = env.Parse(&myEnv)
}

func Auth() string {
	if myEnv.AuthSecret != "" {
		return authFromEnv()
	}
	fmt.Println("Auth secret not found in environment")
	return authFromTerminal()
}

func authFromEnv() string {
	return myEnv.AuthSecret
}

func authFromTerminal() string {
	fmt.Print("Secret: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nError reading password:", err)
		return ""
	}

	// Move to a new line after the user presses Enter
	fmt.Println()
	return string(bytePassword)
}
