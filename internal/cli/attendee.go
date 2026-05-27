package cli

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"leinadium.dev/wedding/internal/client"
)

var attendeesCmd = &cobra.Command{
	Use:   "attendees",
	Short: "Manage attendees",
}

var attendeesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all attendees",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli := client.New(globalFlagURL)
		cli.SetAuth(Auth())
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
		cli := client.New(globalFlagURL)
		cli.SetAuth(Auth())
		id := uuid.MustParse(attendeesDeleteFlag.ID)

		if err := cli.DeleteAttendee(id); err != nil {
			return err
		}
		fmt.Println("Attendee deleted successfully")
		return nil
	},
}

func init() {
	attendeesDeleteCmd.Flags().StringVar(&attendeesDeleteFlag.ID, "id", "", "attendee ID to delete")
	attendeesCmd.AddCommand(attendeesDeleteCmd)
	attendeesCmd.AddCommand(attendeesGetCmd)
	rootCmd.AddCommand(attendeesCmd)
}
