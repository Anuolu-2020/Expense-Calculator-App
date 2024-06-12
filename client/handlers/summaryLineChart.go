package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetSummaryLineItems(summary summaryReport) []opts.LineData {
	summaryValues := []json.Number{summary.TotalIncome, summary.TotalExpense, summary.NetIncome}

	items := make([]opts.LineData, 0)

	for _, values := range summaryValues {
		items = append(items, opts.LineData{Value: values})
	}

	return items
}

func GenerateSummaryLineChart(summary summaryReport, w http.ResponseWriter) {
	// Create a new line instance
	line := charts.NewLine()

	line.PageTitle = "Expense Tracker - Report Graph"

	summaryFields := []string{"Total Income", "Total Expense", "Net Income"}

	// Put data into instance
	line.SetXAxis(summaryFields).
		AddSeries("", GetSummaryLineItems(summary)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

		// Use the custom renderer
	//line.Renderer = graph.NewSnippetRenderer(line, line.Validate)

	line.Render(w)
}
