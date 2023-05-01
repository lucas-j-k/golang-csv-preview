package utils

import (
	"encoding/csv"
	"os"
	"strings"
)

// FileDetails describes a parsed CSV file.
// Includes the rows and headers, along with computed metadata
type FileDetails struct {
	Headers        []string
	Rows           [][]string
	ColumnCount    int
	RowCount       int
	HeaderIndexMap map[string]int
}

// GetColumnIndexes takes an array of header names and returns an array of column indexes
// corresponding to each header string.
func (file FileDetails) GetColumnIndexes(headerNames []string) []int {

	// store the indexes of the specified headers
	headerIndexes := []int{}

	// for each incoming header, store the index against the header in the map
	for _, key := range headerNames {
		capped := strings.ToUpper(key)
		index, ok := file.HeaderIndexMap[capped]
		if ok {
			headerIndexes = append(headerIndexes, index)
		}
	}

	return headerIndexes
}

// GetValuesAtColumns takes an array of column indexes, and returns a slice of transformed CSV rows,
// containing only the fields at the specified indexes.
func (file FileDetails) GetValuesAtColumns(columnIndexes []int) ([]string, [][]string) {

	headers := []string{}
	rows := [][]string{}

	// extract header values at indexes
	for _, indexValue := range columnIndexes {
		headerAtIndex := file.Headers[indexValue]
		headers = append(headers, headerAtIndex)
	}

	// extract row values at the specified indexes
	for _, row := range file.Rows {
		filteredRow := []string{}
		for _, indexValue := range columnIndexes {
			filteredRow = append(filteredRow, row[indexValue])
		}
		rows = append(rows, filteredRow)
	}

	return headers, rows
}

// SearchColumn takes a column header string, and a search term. It searches for all rows in the CSV file where the value in
// the target column matches the search term.
func (file FileDetails) SearchColumn(columnName string, searchTerm string, rowLimit int) [][]string {

	cappedColName := strings.ToUpper(columnName)
	columnIndex, ok := file.HeaderIndexMap[cappedColName]

	rows := [][]string{}

	// invalid column name, break early
	if !ok {
		return rows
	}

	rowIndexes := []int{}

	// loop through all rows, if the value at the column index in the row matches the searchTerm, add the index to the result
	for i, line := range file.Rows {
		// break search early if we have already hit the rowLimit for matching rows
		if len(rowIndexes) == rowLimit {
			break
		}

		cellToMatch := line[columnIndex]
		if strings.EqualFold(cellToMatch, searchTerm) {
			rowIndexes = append(rowIndexes, i)
		}
	}

	// pull out the rows at the found indexes
	for _, indexValue := range rowIndexes {
		rows = append(rows, file.Rows[indexValue])
	}
	return rows
}

// ReadFile parses a CSV file at a given system file path.
// It stores the header strings and rows in a struct, alongside the row and column counts, and
// a map of each column header to the column index.
func ReadFile(fileName string) (*FileDetails, error) {

	// attempt to load in the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	details := FileDetails{}

	// store header strings
	headers, _ := reader.Read()
	details.Headers = headers
	details.ColumnCount = len(headers)

	// store the column index against each header in a map
	headerIndexMap := make(map[string]int)
	for i, headerName := range headers {
		capped := strings.ToUpper(headerName)
		headerIndexMap[capped] = i
	}

	// store rows and total rows
	rows, _ := reader.ReadAll()
	details.Rows = rows
	details.RowCount = len(rows)
	details.HeaderIndexMap = headerIndexMap

	return &details, nil
}
