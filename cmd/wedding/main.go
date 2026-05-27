package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
	"leinadium.dev/wedding/internal/client"
	"leinadium.dev/wedding/internal/models"
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

var attendeesCmd = &cobra.Command{
	Use:   "attendees",
	Short: "Manage attendees",
}

var attendeesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all attendees",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(url)
		cli.SetAuth(getAuth())
		attendees, err := cli.Attendees()
		if err != nil {
			return err
		}
		fmt.Printf("Found %d attendees\n", len(attendees))
		for _, attendee := range attendees {
			fmt.Println("---------")
			fmt.Println("ID:", attendee.ID)
			fmt.Println("Name:", attendee.Name)
			fmt.Println("Is Child:", attendee.IsChild)
			confirmedStatus := "no response"
			if attendee.Confirmed.Valid {
				confirmedStatus = fmt.Sprint(attendee.Confirmed.Bool)
			}
			fmt.Println("Confirmed:", confirmedStatus)
			fmt.Println("Updated At:", attendee.UpdatedAt)
		}
		return nil
	},
}

var attendeesDeleteFlag struct {
	ID string
}
var attendeesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an attendee",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(url)
		cli.SetAuth(getAuth())
		if err := cli.DeleteAttendee(models.AttendeeID(attendeesDeleteFlag.ID)); err != nil {
			return err
		}
		fmt.Println("Attendee deleted successfully")
		return nil
	},
}

var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite management commands",
}

var inviteCreateFlag struct {
	Phone    string
	Attendee []string
}
var inviteCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(url)
		cli.SetAuth(getAuth())
		invite := models.Invite{
			Phone:     inviteCreateFlag.Phone,
			Attendees: make([]models.Attendee, 0, len(inviteCreateFlag.Attendee)),
		}
		for _, name := range inviteCreateFlag.Attendee {
			invite.Attendees = append(invite.Attendees, models.Attendee{Name: name})
		}
		id, err := cli.CreateInvite(invite)
		if err != nil {
			return err
		}
		fmt.Println("Invite created successfully with ID:", id)
		return nil
	},
}

var inviteDeleteFlag struct {
	ID string
}
var inviteDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(url)
		cli.SetAuth(getAuth())
		if inviteDeleteFlag.ID == "" {
			return errors.New("invite ID is required")
		}
		err := cli.DeleteInvite(models.InviteID(inviteDeleteFlag.ID))
		if err != nil {
			return err
		}
		fmt.Println("Invite deleted successfully")
		return nil
	},
}

var rootCmd = &cobra.Command{
	Use:   "wedding",
	Short: "Wedding CLI",
	Long:  `A CLI for managing wedding attendees and purchases.`,
}

func init() {
	// root command flags
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "http://localhost:8080/v1", "API URL")

	// invite
	inviteCreateCmd.Flags().StringVar(&inviteCreateFlag.Phone, "phone", "", "phone number of the invitee")
	inviteCreateCmd.Flags().StringSliceVar(&inviteCreateFlag.Attendee, "name", nil, "add attendee name (repeatable)")
	inviteDeleteCmd.Flags().StringVar(&inviteDeleteFlag.ID, "id", "", "invite ID to delete")
	inviteCmd.AddCommand(inviteCreateCmd)
	inviteCmd.AddCommand(inviteDeleteCmd)
	rootCmd.AddCommand(inviteCmd)

	attendeesDeleteCmd.Flags().StringVar(&attendeesDeleteFlag.ID, "id", "", "attendee ID to delete")
	attendeesCmd.AddCommand(attendeesDeleteCmd)
	attendeesCmd.AddCommand(attendeesGetCmd)
	rootCmd.AddCommand(attendeesCmd)

	// commands without flags
	rootCmd.AddCommand(productsCmd)
	rootCmd.AddCommand(purchasesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
