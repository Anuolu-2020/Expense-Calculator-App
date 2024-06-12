package handlers

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func GetPieChartItems(reports Response) []opts.PieData {
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

	items := make([]opts.PieData, 0)

	for _, report := range summedReports {
		items = append(items, opts.PieData{Name: report.Source, Value: report.Amount})
	}

	return items
}

// Generate Pie chart
func GeneratePieChart(reports Response, w http.ResponseWriter) {
	pie := charts.NewPie()

	pie.PageTitle = "Expense Tracker - Report Graph"

	pie.AddSeries("pie", GetPieChartItems(reports)).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{b}: {c}"}),
	)

	// Use the custom Renderer
	//	pie.Renderer = graph.NewSnippetRenderer(pie, pie.Validate)

	pie.Render(w)
}

func GenerateReportPieChart(userId string, w http.ResponseWriter) {
	reports, err := GetReports(userId)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.SendErrorResponse(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	GeneratePieChart(*reports, w)
}

func GenerateReportTypePieChart(reportType string, userId string, w http.ResponseWriter) {
	reports, err := GetReportsType(reportType, userId)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.SendErrorResponse(w, "An Error occurred", http.StatusInternalServerError)
		return
	}

	GeneratePieChart(*reports, w)
}
