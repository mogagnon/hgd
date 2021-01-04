package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Token string

var rootCmd = &cobra.Command{
	Use:   "hgd",
	Short: "hgd is a simple utility to manage dashboard in hosted graphite.",
	Long:  `hdg main purpose is to facilitate the deployment of dashboard in hosted graphite.`,
}

func init() {
	rootCmd.Flags()
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "The token for Hosted Graphite.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
