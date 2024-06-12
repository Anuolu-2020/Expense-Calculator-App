package handlers

import (
	"log"
	"net/http"
	"slices"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func GetLineItems(reports Response) []opts.LineData {
	summedReports := make([]Results, 0)

	for _, report := range reports.Results {
		found := false

		for idx, summedReport := range summedReports {
			if summedReport.Source == report.Source {
				summedReports[idx].Amount += report.Amount
				found = true
				break
			}
		}

		if !found {
			summedReports = append(summedReports, report)
		}
	}

	items := make([]opts.LineData, 0)

	for _, report := range summedReports {
		items = append(items, opts.LineData{Value: report.Amount})
	}

	return items
}

func GenerateLineChart(reports Response, w http.ResponseWriter) {
	// Create a new line instance
	line := charts.NewLine()

	line.PageTitle = "Expense Tracker - Report Graph"

	var reportsSources []string

	for _, report := range reports.Results {
		if slices.Contains(reportsSources, report.Source) {
			continue
		} else {
			reportsSources = append(reportsSources, report.Source)
		}
	}

	// Put data into instance
	line.SetXAxis(reportsSources).
		AddSeries("", GetLineItems(reports)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

		// Use the custom renderer
	//line.Renderer = graph.NewSnippetRenderer(line, line.Validate)

	line.Render(w)
}

func GenerateReportLineChart(userId string, w http.ResponseWriter) {
	reports, err := GetReports(userId)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.SendErrorResponse(w, "An Error Occurred", http.StatusInternalServerError)
	}

	GenerateLineChart(*reports, w)
}

func GenerateReportTypeLineChart(
	reportsType string,
	userId string,
	w http.ResponseWriter,
) {
	reports, err := GetReportsType(reportsType, userId)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.SendErrorResponse(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	GenerateLineChart(*reports, w)
}
