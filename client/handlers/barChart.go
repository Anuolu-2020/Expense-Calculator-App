package handlers

import (
	"log"
	"net/http"
	"slices"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func GetBarItems(reports Response) []opts.BarData {
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

	items := make([]opts.BarData, 0)

	for _, report := range summedReports {
		items = append(items, opts.BarData{Value: report.Amount})
	}

	return items
}

func GenerateBarChart(reports Response, w http.ResponseWriter) {
	// Create a new line instance
	bar := charts.NewBar()

	bar.PageTitle = "Expense Tracker - Report Graph"

	var reportsSources []string

	for _, report := range reports.Results {
		if slices.Contains(reportsSources, report.Source) {
			continue
		} else {
			reportsSources = append(reportsSources, report.Source)
		}
	}

	// Put data into instance
	bar.SetXAxis(reportsSources).
		AddSeries("", GetBarItems(reports))

		// Use the custom renderer
	//line.Renderer = graph.NewSnippetRenderer(line, line.Validate)

	bar.Render(w)
}

func GenerateReportBarChart(userId string, w http.ResponseWriter) {
	reports, err := GetReports(userId)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	GenerateBarChart(*reports, w)
}

func GenerateReportTypeBarChart(
	reportsType string,
	userId string,
	w http.ResponseWriter,
) {
	reports, err := GetReportsType(reportsType, userId)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	GenerateBarChart(*reports, w)
}
