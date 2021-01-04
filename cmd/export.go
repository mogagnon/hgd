package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"hgd/api"
	"hgd/io"
	"os"
	"path/filepath"
	"strings"
)

var directory bool

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a dashboard to Hosted Graphite from a json file. Will update if the dashboard already exist.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("required a json file")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		export(args[0])
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().BoolVarP(&directory, "all", "a", false, "To export all graph from a directory.")
}

func export(file string) error {
	var err error

	client := api.CreateHgClient(Token)

	if directory {
		filepath.Walk(file, func(filePath string, info os.FileInfo, err error) error {
			if strings.Contains(filePath, ".json") {
				tmp, _ := io.LoadDashboard(filePath)
				err = createOfUpdate(&client, tmp)

				if err != nil {
					fmt.Println("\033[31m", fmt.Errorf("unable to export %s to hosted graphite: %v", filePath, err))
					os.Exit(0)
				}
			}
			return nil
		})
	} else {
		dashboard, _ := io.LoadDashboard(file)
		err = createOfUpdate(&client, dashboard)
	}

	if err != nil {
		fmt.Println("\033[31m", fmt.Errorf("unable to export %s to hosted graphite: %v", file, err))
		os.Exit(0)
	}

	fmt.Println("\033[32m", "dashboard exported successfully")

	return nil
}

func createOfUpdate(client *api.HgClient, dashboard *api.DashboardTemplate) error {
	if client.Exist(dashboard.Dashboard.Title) {
		return client.Update(dashboard)
	} else {
		return client.Create(dashboard)
	}
}
