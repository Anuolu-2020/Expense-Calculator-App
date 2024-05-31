package handlers

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func GetLineItems(reports Response) []opts.LineData {
	items := make([]opts.LineData, 0)

	for _, report := range reports.Results {
		items = append(items, opts.LineData{Value: report.Amount})
	}

	return items
}

func GenerateLineChart(reports Response) *charts.Line {
	// Create a new line instance
	line := charts.NewLine()

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Reports Chart Analysis",
			Subtitle: "Analysis of your reports so far",
		}))

	reportsSources := make([]string, 0)
	// var reportsType string

	for _, report := range reports.Results {
		reportsSources = append(reportsSources, report.Source)
	}

	// Put data into instance
	line.SetXAxis(reportsSources).
		AddSeries("", GetLineItems(reports))

	return line
}

func (h Handler) LineChart(w http.ResponseWriter, r *http.Request) {
	reports, err := GetReports(r)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.ServeErrorPage(w, r)
	}

	line := GenerateLineChart(*reports)

	line.Render(w)
}
