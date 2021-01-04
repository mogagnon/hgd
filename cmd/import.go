package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"hgd/api"
	"hgd/io"
	"os"
)

var output string

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a dashboard locally.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("required a dashboard name")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		importDashboard(output, args[0])
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVarP(&output, "output", "o", "", "Json output file for the dashboard.")
}

func importDashboard(outputFile string, id string) error {
	client := api.CreateHgClient(Token)

	dashboard, err := client.Get(id)

	if err != nil {
		fmt.Println("\033[31m", fmt.Errorf("unable to import %s from hosted graphite: %v", id, err))
		os.Exit(0)
	}

	io.SaveDashboard(dashboard, outputFile)

	fmt.Println("\033[32m", "dashboard imported successfully")

	return nil
}
