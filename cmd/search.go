package cmd

import (
	"example/go-csv-preview/utils"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var searchByValue = &cobra.Command{
	Use:   "search",
	Short: "Search file for a given string",
	Long:  `Based on a column name and a search term, print out rows which match the search parameters`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(2), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {

		var columnName string
		var searchTerm string

		// [0] : target column name
		if args[0] != "" {
			columnName = args[0]
		}

		// [1] : target column name
		if args[1] != "" {
			searchTerm = args[1]
		}

		// retrieve flags
		rowLimit, _ := rootCmd.Flags().GetInt("row-limit")
		filePath, _ := rootCmd.Flags().GetString("file")

		details, err := utils.ReadFile(filePath)

		if err != nil {
			fmt.Printf("Unable to read csv file: [%s]\n\n", filePath)
			return
		}

		// check that the target column header exists in the parsed CSV file
		_, ok := details.HeaderIndexMap[strings.ToUpper(columnName)]
		if !ok {
			fmt.Println("Invalid column name")
			return
		}

		// retrieve the indexes of any rows where the given column matches the search term
		matchedRows := details.SearchColumn(columnName, searchTerm, rowLimit)

		// no matching rows found
		if len(matchedRows) == 0 {
			fmt.Println("No results found")
			return
		}

		// print the result rows. Provide a default rowLimit to guard against overly generic searches
		utils.PrintTableToStdout(details.Headers, matchedRows, rowLimit)

	},
}

func init() {
	rootCmd.AddCommand(searchByValue)
}
