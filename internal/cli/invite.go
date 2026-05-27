package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"leinadium.dev/wedding/internal/client"
	"leinadium.dev/wedding/internal/models"
)

var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite management commands",
}

var inviteListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all invites",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(globalFlagURL)
		cli.SetAuth(Auth())
		invites, err := cli.Invites()
		if err != nil {
			return err
		}
		for _, invite := range invites {
			var names strings.Builder
			for _, attendee := range invite.Attendees {
				names.WriteString(attendee.Name)
				names.WriteString(", ")
			}

			fmt.Println("---------")
			fmt.Println("ID:", invite.ID)
			fmt.Println("Phone:", invite.Phone)
			fmt.Println("Names:", names.String())
		}
		return nil
	},
}

var inviteCreateFlag struct {
	Phone    string
	Attendee []string
}
var inviteCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(globalFlagURL)
		cli.SetAuth(Auth())
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
		cli := client.New(globalFlagURL)
		cli.SetAuth(Auth())
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

func init() {
	inviteCreateCmd.Flags().StringVar(&inviteCreateFlag.Phone, "phone", "", "phone number of the invitee")
	inviteCreateCmd.Flags().StringSliceVar(&inviteCreateFlag.Attendee, "name", nil, "add attendee name (repeatable)")
	inviteDeleteCmd.Flags().StringVar(&inviteDeleteFlag.ID, "id", "", "invite ID to delete")
	inviteCmd.AddCommand(inviteCreateCmd)
	inviteCmd.AddCommand(inviteDeleteCmd)
	inviteCmd.AddCommand(inviteListCmd)
	rootCmd.AddCommand(inviteCmd)
}
