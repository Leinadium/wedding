package cli

import (
	"github.com/spf13/cobra"
)

var globalFlagURL string

var rootCmd = &cobra.Command{
	Use:   "wedding",
	Short: "Wedding CLI",
	Long:  `A CLI for managing wedding attendees and purchases.`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&globalFlagURL, "url", "u", "http://localhost:8080/v1", "API URL")
}

func Execute() error {
	return rootCmd.Execute()
}
