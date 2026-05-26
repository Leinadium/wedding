package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"leinadium.dev/wedding/internal/client"
)

var url string

func getAuth() string {
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

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Get all available products",
	RunE: func(cmd *cobra.Command, args []string) error {
		products, err := client.New(url).Products()
		if err != nil {
			return err
		}
		fmt.Printf("Found %d products\n", len(products))
		for _, product := range products {
			fmt.Println("---------")
			fmt.Println("StripeID:", product.StripeID)
			fmt.Printf("Name: %s\n", product.Name)
			fmt.Printf("Price: %d\n", product.PriceBRL)
			fmt.Printf("Purchased: %t\n", product.Purchased)
		}
		return nil
	},
}

var purchasesCmd = &cobra.Command{
	Use:   "purchases",
	Short: "Get all purchases",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(url)
		cli.SetAuth(getAuth())
		purchases, err := cli.Purchases()
		if err != nil {
			return err
		}
		fmt.Printf("Found %d purchases\n", len(purchases))
		for _, purchase := range purchases {
			fmt.Println("---------")
			fmt.Println("ID:", purchase.ID)
			fmt.Println("Email:", purchase.Email)
			fmt.Println("ProductID:", purchase.ProductID)
			fmt.Println("ProductName:", purchase.ProductName)
			fmt.Println("Price:", purchase.Price)
		}
		return nil
	},
}

var rootCmd = &cobra.Command{
	Use:   "wedding",
	Short: "Wedding CLI",
	Long:  `A CLI for managing wedding attendees and purchases.`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "http://localhost:8080/v1", "API URL")

	rootCmd.AddCommand(productsCmd)
	rootCmd.AddCommand(purchasesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
