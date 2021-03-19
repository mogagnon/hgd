package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hgd/api"
	"os"
	"time"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an event.",
	Run: func(cmd *cobra.Command, args []string) {
		createEvent(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createEvent(what string, tags string) error {
	client := api.CreateHgClient(Token)

	err := client.CreateEvent(api.Event{
		What: what,
		When: time.Now().Unix(),
		Tags: tags,
	})

	if err != nil {
		fmt.Println("\033[31m", fmt.Errorf("unable to create event%s: %v", what, err))
		os.Exit(0)
	}

	fmt.Println("\033[32m", "event created successfully")

	return nil
}
