package templatemethod

import "fmt"

/*
Template Method is a behavioral design pattern that defines the skeleton of an algorithm in the superclass
but lets subclasses override specific steps of the algorithm without changing its structure.
----
Template Method is a behavioral design pattern that allows you to define a skeleton of an algorithm in a base class
and let subclasses override the steps without changing the overall algorithm’s structure.
----
It’s particularly useful in situations where:

You have a well-defined sequence of steps to perform a task.
Some parts of the process are shared across all implementations.
You want to allow subclasses to customize specific steps without rewriting the whole algorithm.

//////|| Strategy vs Template Method ||\\\\\\

The real question to ask yourself:

"Do I need to enforce a sequence of steps, or do I just need the end result?"

If the process matters — open, validate, execute, log —
and you want subclasses to only tweak specific steps → Template Method

If only the outcome matters and the algorithm can do whatever it wants internally → Strategy

---

We use an abstract class rather than an interface because the
base class needs to provide a concrete template method with real logic (the step ordering).
An interface cannot contain a method that calls other methods in sequence while enforcing that sequence.

This is one of the few patterns where inheritance is the right tool, not composition.
*/

type ReportData struct{}

func (r ReportData) GetHeaders() []string {
	return []string{"ID", "Name", "Value"}
}

func (r ReportData) GetRows() []map[string]interface{} {
	return []map[string]interface{}{
		{"ID": 1, "Name": "Item A", "Value": 100.0},
		{"ID": 2, "Name": "Item B", "Value": 150.5},
		{"ID": 3, "Name": "Item C", "Value": 75.25},
	}
}

func TestTemplateMethod() {
	// CSV export
	csv := &AbstractReportExporter{writer: &CSVExporter{}}
	csv.ExportReport(ReportData{}, "report.csv")

	fmt.Println()
	// PDF export
	pdf := &AbstractReportExporter{writer: &PDFExporter{}}
	pdf.ExportReport(ReportData{}, "report.pdf")
}
