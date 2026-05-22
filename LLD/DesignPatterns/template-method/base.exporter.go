package templatemethod

import "fmt"

// Base hook — default no-op footer, embedded in each exporter
type BaseExporter struct{}

func (b *BaseExporter) writeFooter(data ReportData) {
	// default: no footer
	fmt.Println("Base writeFooter Called")
}
