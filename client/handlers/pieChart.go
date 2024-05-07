package handlers

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

func GetPieChartItems(reports Response) []opts.PieData {
	items := make([]opts.PieData, 0)

	for _, report := range reports.Results {
		items = append(items, opts.PieData{Name: report.Source, Value: report.Amount})
	}

	return items
}

func GeneratePieChart(reports Response) *charts.Pie {
	pie := charts.NewPie()

	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Your reports so far"}),
	)

	pie.AddSeries("pie", GetPieChartItems(reports))

	return pie
}

func (h Handler) PieChart(w http.ResponseWriter, r *http.Request) {
	reports, err := GetReports(r)
	if err != nil {
		log.Printf("Error occurred while retrieving reports: %v", err)
		pkg.ServeErrorPage(w, r)
	}

	bar := GeneratePieChart(*reports)

	bar.Render(w)
}
