import { Injectable } from '@nestjs/common';
import { ReportType } from 'src/dtos/report.dto';
import { ReportService, Report } from 'src/report/report.service';

@Injectable()
export class SummaryService {
  constructor(private readonly reportService: ReportService) {}
  async calculateSummary() {
    const getTotalExpense = await this.reportService.getAllReports(
      ReportType.expense,
    );

    const totalExpense = getTotalExpense.reduce(
      (sum: number, report: Report) => sum + report.amount,
      0,
    );

    const getTotalIncome = await this.reportService.getAllReports(
      ReportType.income,
    );

    const totalIncome = getTotalIncome.reduce(
      (sum: number, report: Report) => sum + report.amount,
      0,
    );

    return {
      totalIncome,
      totalExpense,
      netIncome: totalIncome - totalExpense,
    };
  }
}
