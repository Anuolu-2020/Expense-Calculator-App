import { ReportType } from 'src/data';
export declare class CreateReportDto {
    amount: number;
    source: string;
}
export declare class UpdateReportDto {
    amount: number;
    source: string;
}
export declare class ReportResponseDto {
    id: string;
    source: string;
    amount: number;
    created_at: Date;
    updated_at: Date;
    type: ReportType;
    transformCreatedAt(): Date;
    constructor(partial: Partial<ReportResponseDto>);
}
