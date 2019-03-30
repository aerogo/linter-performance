package performance

import (
	"fmt"

	"github.com/blitzprog/color"
)

// printResult shows the result for a route on the terminal.
func printResult(label string, responseTime int64, responseSize float64) {
	faint := color.New(color.Faint).SprintFunc()

	// Response size color
	var responseSizeColor func(a ...interface{}) string

	switch {
	case responseSize < 15:
		responseSizeColor = color.New(color.FgGreen).SprintFunc()
	case responseSize < 100:
		responseSizeColor = color.New(color.FgYellow).SprintFunc()
	default:
		responseSizeColor = color.New(color.FgRed).SprintFunc()
	}

	// Response time color
	var responseTimeColor func(a ...interface{}) string

	switch {
	case responseTime < 10:
		responseTimeColor = color.New(color.FgGreen).SprintFunc()
	case responseTime < 100:
		responseTimeColor = color.New(color.FgYellow).SprintFunc()
	default:
		responseTimeColor = color.New(color.FgRed).SprintFunc()
	}

	// Output
	fmt.Printf("%-67s %s %s %s %s\n", color.BlueString(label), responseSizeColor(fmt.Sprintf("%6.0f", responseSize)), faint("KB"), responseTimeColor(fmt.Sprintf("%7d", responseTime)), faint("ms"))
}
