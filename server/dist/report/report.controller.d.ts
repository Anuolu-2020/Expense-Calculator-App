import { ReportService } from './report.service';
import { CreateReportDto, ReportResponseDto, UpdateReportDto } from '../dtos/report.dto';
export declare class ReportController {
    private readonly reportService;
    constructor(reportService: ReportService);
    getAllReports(type: string): ReportResponseDto[];
    getReportById(type: string, id: string): ReportResponseDto;
    createReport({ source, amount }: CreateReportDto, type: string): ReportResponseDto;
    updateReportById(type: string, id: string, body: UpdateReportDto): ReportResponseDto;
    deleteReportById(id: string): void;
}
