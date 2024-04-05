import { ReportService } from 'src/report/report.service';
export declare class SummaryService {
    private readonly reportService;
    constructor(reportService: ReportService);
    calculateSummary(): {
        totalIncome: number;
        totalExpense: number;
        netIncome: number;
    };
}
