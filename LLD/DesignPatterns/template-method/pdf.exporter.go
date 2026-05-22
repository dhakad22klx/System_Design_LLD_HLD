package templatemethod

import "fmt"

// PDF — overrides footer hook as well
type PDFExporter struct {
	BaseExporter
}

func (p *PDFExporter) writeHeader(data ReportData) {
	fmt.Println("Writing PDF header with logo")
}

func (p *PDFExporter) writeDataRows(data ReportData) {
	fmt.Println("Writing PDF table rows")
}

func (p *PDFExporter) writeFooter(data ReportData) { // overrides hook
	fmt.Println("Writing PDF footer: Page 1 of N")
}
