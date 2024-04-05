import { SummaryService } from './summary.service';
export declare class SummaryController {
    private readonly summaryService;
    constructor(summaryService: SummaryService);
    getSummary(): {
        totalIncome: number;
        totalExpense: number;
        netIncome: number;
    };
}
