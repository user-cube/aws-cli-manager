package ui

import (
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

// TableData represents a row of data in a table
type TableData struct {
	Columns []interface{}
}

// TableOptions contains options for table styling
type TableOptions struct {
	Headers []string
	Rows    []TableData
}

// RenderTable renders a table with the provided headers and rows
func RenderTable(options TableOptions) {
	headerColor := color.New(color.FgHiCyan, color.Bold).SprintFunc()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Convert string headers to colored header row
	headerRow := make(table.Row, len(options.Headers))
	for i, header := range options.Headers {
		headerRow[i] = headerColor(header)
	}
	t.AppendHeader(headerRow)

	// Add all rows
	for _, rowData := range options.Rows {
		t.AppendRow(table.Row(rowData.Columns))
	}

	t.SetStyle(table.StyleLight)
	t.Render()
}
