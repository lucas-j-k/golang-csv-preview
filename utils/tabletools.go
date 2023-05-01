package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// PrintTableToStdout takes an array of string headers, an array of row arrays an a rowLimit.
// Renders a table to the terminal displaying the data
func PrintTableToStdout(headers []string, rows [][]string, rowLimit int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	columnCount := len(headers)

	headerRow := table.Row{}

	for _, headerText := range headers {
		headerRow = append(headerRow, headerText)
	}

	t.AppendHeader(headerRow)

	for i, row := range rows {
		// break once rowLimit is hit
		if i >= rowLimit {
			break
		}

		tableRow := table.Row{}
		// pull all the string values out of the csv row and insert into the new table row
		for i := 0; i < columnCount; i++ {
			tableRow = append(tableRow, row[i])
		}
		t.AppendRow(tableRow)
	}

	/*
	*	Render the table
	 */
	t.AppendSeparator()
	t.Render()
}
