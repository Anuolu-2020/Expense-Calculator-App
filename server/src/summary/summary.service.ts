import { Injectable } from '@nestjs/common';
import { ReportType } from 'src/dtos/report.dto';
import { ReportService, Report } from 'src/report/report.service';

@Injectable()
export class SummaryService {
  constructor(private readonly reportService: ReportService) { }
  async calculateSummary(id: string) {
    const getTotalExpense = await this.reportService.getReportTypeByUserId(
      ReportType.expense,
      id,
    );

    let totalExpense = 0;
    if (Array.isArray(getTotalExpense)) {
      totalExpense = getTotalExpense.reduce(
        (sum: number, report: Report) => sum + report.amount,
        0,
      );
    }

    const getTotalIncome = await this.reportService.getReportTypeByUserId(
      ReportType.income,
      id,
    );

    let totalIncome = 0;
    if (Array.isArray(getTotalIncome)) {
      totalIncome = getTotalIncome.reduce(
        (sum: number, report: Report) => sum + report.amount,
        0,
      );
    }

    return {
      totalIncome,
      totalExpense,
      netIncome: totalIncome - totalExpense,
    };
  }
}
