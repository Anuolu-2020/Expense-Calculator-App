export declare enum ReportType {
    INCOME = "income",
    EXPENSE = "expense"
}
export declare const data: Data;
interface Data {
    report: {
        id: string;
        source: string;
        amount: number;
        created_at: Date;
        updated_at: Date;
        type: ReportType;
    }[];
}
export {};
