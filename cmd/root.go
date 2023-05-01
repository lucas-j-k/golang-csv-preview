package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-csv-preview",
	Short: "Command line preview for CSV files",
	Long:  `Command line preview for CSV files`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().Int("row-limit", 5, "Define maximum number of rows to return")
	rootCmd.PersistentFlags().String("file", "", "Path to CSV file")
}
