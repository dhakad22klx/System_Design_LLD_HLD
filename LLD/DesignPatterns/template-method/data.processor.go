package templatemethod

import (
	"fmt"
	"strings"
)

// DataProcessor defines the interface for steps
// and provides the template method Process
// (Go doesn't have abstract classes, so we use interface + embedding)

/*
Go has no virtual dispatch. Embedding promotes methods but doesn't override them polymorphically.
The skeleton calls the base methods, always.

The fix — interface as the abstract contract
Pull the varying steps into an interface. The skeleton accepts the interface, so dispatch works correctly.
*/
// Abstract steps — each exporter MUST implement these

type ReportWriter interface {
	writeHeader(data ReportData)
	writeDataRows(data ReportData)
	writeFooter(data ReportData) // hook — provide default in base
}

// Skeleton lives here — owns the fixed flow
type AbstractReportExporter struct {
	writer ReportWriter // ← dispatch goes through interface
}

func (a *AbstractReportExporter) ExportReport(data ReportData, filePath string) {
	a.prepareData(data)
	a.openFile(filePath)
	a.writer.writeHeader(data)   // calls concrete implementation
	a.writer.writeDataRows(data) // calls concrete implementation
	a.writer.writeFooter(data)   // calls concrete implementation
	a.closeFile(filePath)
	fmt.Println("Export complete: " + filePath)
}

// Shared concrete steps
func (a *AbstractReportExporter) prepareData(data ReportData) {
	
	fmt.Println("Preparing report data..." + strings.Join(data.GetHeaders(), ","))
}
func (a *AbstractReportExporter) openFile(filePath string) {
	fmt.Println("Opening file: " + filePath)
}
func (a *AbstractReportExporter) closeFile(filePath string) {
	fmt.Println("Closing file: " + filePath)
}
