package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hgd/api"
	"os"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Dashboard.",
	Run: func(cmd *cobra.Command, args []string) {
		deleteDashboard(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteDashboard(name string) error {
	client := api.CreateHgClient(Token)

	err := client.Delete(name)

	if err != nil {
		fmt.Println("\033[31m", fmt.Errorf("unable to delete %s: %v", name, err))
		os.Exit(0)
	}

	fmt.Println("\033[32m", "dashboard deleted successfully")

	return nil
}
