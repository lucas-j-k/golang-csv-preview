package cmd

import (
	"example/go-csv-preview/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview lines of a CSV file",
	Long:  `Preview lines of a csv file. Displays all columns. Accepts a --row-limit flag, to specify how many lines to print.`,
	Run: func(cmd *cobra.Command, args []string) {

		// retrieve flags
		rowLimit, _ := rootCmd.Flags().GetInt("row-limit")
		filePath, _ := rootCmd.Flags().GetString("file")

		details, err := utils.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Unable to read csv file: [%s]\n\n", filePath)
			return
		}

		// print the data to terminal
		utils.PrintTableToStdout(details.Headers, details.Rows, rowLimit)
	},
}

var previewColumns = &cobra.Command{
	Use:   "preview-columns",
	Short: "Preview values for one or more column(s)",
	Long:  `Preview values for one or more columns(s) by name`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {

		var columnNames []string

		if args[0] != "" {
			columnNames = strings.Split(args[0], ",")
		}

		// retrieve flags
		rowLimit, _ := rootCmd.Flags().GetInt("row-limit")
		filePath, _ := rootCmd.Flags().GetString("file")

		details, err := utils.ReadFile(filePath)

		if err != nil {
			fmt.Printf("Unable to read csv file: [%s]\n\n", filePath)
			return
		}

		// retrieve the index of each column specified
		targetIndexes := details.GetColumnIndexes(columnNames)

		// no column headers in the file match the provided flag options
		if len(targetIndexes) == 0 {
			fmt.Println("No matching columns found")
		}

		// get only the columns at the specified indexes
		headers, rows := details.GetValuesAtColumns(targetIndexes)

		utils.PrintTableToStdout(headers, rows, rowLimit)
	},
}

func init() {
	rootCmd.AddCommand(previewCmd)
	rootCmd.AddCommand(previewColumns)
}
