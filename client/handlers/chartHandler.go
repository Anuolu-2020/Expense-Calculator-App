package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"

	"github.com/Anuolu-2020/Expense-Calculator-App/pkg"
)

type ReportChartBody struct {
	ReportType string `json:"reportType" validate:"required"`
	ChartType  string `json:"chartType"  validate:"required"`
}

type summaryChartRequestBody struct {
	ChartType string `json:"summaryChartType" validate:"required"`
}

func (h Handler) ReportChart(w http.ResponseWriter, r *http.Request) {
	allowedReportValues := []string{"reports", "income", "expense"}
	allowedChartValues := []string{"piechart", "linechart"}

	userId := r.PathValue("userId")

	var reportBody ReportChartBody

	err := json.NewDecoder(r.Body).Decode(&reportBody)
	if err != nil {
		log.Printf("Json Error: %v", err)
		pkg.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if ok, err := pkg.ValidateInputs(reportBody); !ok {
		pkg.ValidationError(w, http.StatusBadRequest, err)
		return
	}

	if ok := slices.Contains(allowedReportValues, reportBody.ReportType); !ok {
		pkg.SendErrorResponse(w, "Invalid reportType value", http.StatusBadRequest)
	}

	if ok := slices.Contains(allowedChartValues, reportBody.ChartType); !ok {
		pkg.SendErrorResponse(w, "Invalid chartType value", http.StatusBadRequest)
	}

	// If it's All reports
	if reportBody.ReportType == "reports" {
		if reportBody.ChartType == "linechart" {
			GenerateReportLineChart(userId, w)
			return
		} else {
			GenerateReportPieChart(userId, w)
			return
		}
	}

	// If it's expense or income
	if reportBody.ChartType == "linechart" {
		GenerateReportTypeLineChart(reportBody.ReportType, userId, w)
		return
	} else {
		GenerateReportTypePieChart(reportBody.ReportType, userId, w)
		return
	}
}

func (h Handler) ReportSummaryChart(w http.ResponseWriter, r *http.Request) {
	allowedChartValues := []string{"piechart", "linechart"}

	var response summaryReport

	userId := r.PathValue("userId")

	var summaryReportBody summaryChartRequestBody

	err := json.NewDecoder(r.Body).Decode(&summaryReportBody)
	if err != nil {
		log.Printf("Json Error: %v", err)
		pkg.SendErrorResponse(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if ok, err := pkg.ValidateInputs(summaryReportBody); !ok {
		pkg.ValidationError(w, http.StatusBadRequest, err)
		return
	}

	if ok := slices.Contains(allowedChartValues, summaryReportBody.ChartType); !ok {
		pkg.SendErrorResponse(w, "Invalid chartType value", http.StatusBadRequest)
	}

	// Get Summaries
	resp, err := GetSummary(userId)
	if err != nil {
		log.Printf("An Error Occurred: %v", err)
		pkg.SendErrorResponse(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("An Error Ocurred: %v", err)
		pkg.SendErrorResponse(w, "An Error Occurred", http.StatusInternalServerError)
		return
	}

	if summaryReportBody.ChartType == "piechart" {
		GenerateSummaryPieChart(response, w)
		return
	} else {
		GenerateSummaryLineChart(response, w)
		return
	}
}
