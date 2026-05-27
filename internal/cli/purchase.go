package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"leinadium.dev/wedding/internal/client"
)

var purchasesCmd = &cobra.Command{
	Use:   "purchases",
	Short: "Get all purchases",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(globalFlagURL)
		cli.SetAuth(Auth())
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

func init() {
	rootCmd.AddCommand(purchasesCmd)
}
