package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"leinadium.dev/wedding/internal/client"
)

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Get all available products",
	RunE: func(cmd *cobra.Command, args []string) error {
		products, err := client.New(globalFlagURL).Products()
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

func init() {
	rootCmd.AddCommand(productsCmd)
}
