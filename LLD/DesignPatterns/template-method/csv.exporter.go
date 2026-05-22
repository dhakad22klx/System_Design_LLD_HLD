package templatemethod

import (
	"fmt"
	"strings"
)

// CSV — must implement header + rows, inherits no-op footer
type CSVExporter struct {
	BaseExporter // gets default writeFooter
}

func (c *CSVExporter) writeHeader(data ReportData) {
	fmt.Println("Writing CSV header:", strings.Join(data.GetHeaders(), ","))
}

func (c *CSVExporter) writeDataRows(data ReportData) {
	fmt.Println("Writing CSV Rows:")
	for _, row := range data.GetRows() {
		values := make([]string, 0, len(data.GetHeaders()))
		for _, h := range data.GetHeaders() {
			values = append(values, fmt.Sprint(row[h]))
		}
		fmt.Println("CSV: " + strings.Join(values, ","))
	}
}
