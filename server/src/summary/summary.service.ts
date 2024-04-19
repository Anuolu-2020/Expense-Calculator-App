import { Injectable } from '@nestjs/common';
import { ReportType } from 'src/dtos/report.dto';
import { ReportService, Report } from 'src/report/report.service';

@Injectable()
export class SummaryService {
  constructor(private readonly reportService: ReportService) { }
  async calculateSummary(id: string) {
    const getTotalExpense = await this.reportService.getReportById(
      ReportType.expense,
      id,
    );

    const totalExpense = getTotalExpense.reduce(
      (sum: number, report: Report) => sum + report.amount,
      0,
    );

    const getTotalIncome = await this.reportService.getReportById(
      ReportType.income,
      id,
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
