import { ReportType } from '../data';
import { ReportResponseDto } from '../dtos/report.dto';
interface Report {
    source: string;
    amount: number;
}
interface UpdateReport {
    source?: string;
    amount?: number;
}
export declare class ReportService {
    getAllReports(type: ReportType): ReportResponseDto[];
    getReportById(type: ReportType, id: string): ReportResponseDto;
    createReport(type: ReportType, { source, amount }: Report): ReportResponseDto;
    updateReport(type: ReportType, id: string, body: UpdateReport): ReportResponseDto;
    deleteReport(id: string): void;
}
export {};
