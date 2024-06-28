package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetBarSummaryChartItems(summary summaryReport) []opts.BarData {
	summaryValues := []json.Number{summary.TotalIncome, summary.TotalExpense, summary.NetIncome}
	summariesField := []string{"Total Income", "Total Expense", "Net Income"}

	items := make([]opts.BarData, 0)

	for i := range summariesField {
		items = append(items, opts.BarData{Name: summariesField[i], Value: summaryValues[i]})
	}

	return items
}

// Generate Pie chart
func GenerateSummaryBarChart(summary summaryReport, w http.ResponseWriter) {
	summariesField := []string{"Total Income", "Total Expense", "Net Income"}

	bar := charts.NewBar()

	bar.PageTitle = "Expense Tracker - Report Graph"

	bar.SetXAxis(summariesField).AddSeries("", GetBarSummaryChartItems(summary))

	bar.Render(w)
}
