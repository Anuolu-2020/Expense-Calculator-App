package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetPieSummaryChartItems(summary summaryReport) []opts.PieData {
	summaryValues := []json.Number{summary.TotalIncome, summary.TotalExpense, summary.NetIncome}
	summariesField := []string{"Total Income", "Total Expense", "Net Income"}

	items := make([]opts.PieData, 0)

	for i := range summariesField {
		items = append(items, opts.PieData{Name: summariesField[i], Value: summaryValues[i]})
	}

	return items
}

// Generate Pie chart
func GenerateSummaryPieChart(summary summaryReport, w http.ResponseWriter) {
	pie := charts.NewPie()

	pie.PageTitle = "Expense Tracker - Report Graph"

	pie.AddSeries("pie", GetPieSummaryChartItems(summary)).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{Show: true, Formatter: "{b}: {c}"}),
	)

	pie.Render(w)
}
