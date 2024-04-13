import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Param,
  Body,
  HttpCode,
  ParseUUIDPipe,
  ParseEnumPipe,
} from '@nestjs/common';
import { ReportService } from './report.service';
import {
  CreateReportDto,
  ReportResponseDto,
  UpdateReportDto,
  ReportType,
} from '../dtos/report.dto';

@Controller('report/:type')
export class ReportController {
  constructor(private readonly reportService: ReportService) {}
  @Get()
  async getAllReports(
    @Param('type', new ParseEnumPipe(ReportType)) type: string,
  ): Promise<ReportResponseDto[]> {
    const reportType =
      type === 'income' ? ReportType.income : ReportType.expense;

    return this.reportService.getAllReports(reportType);
  }

  @Get(':id')
  async getReportById(
    @Param('type', new ParseEnumPipe(ReportType)) type: string,
    @Param('id', ParseUUIDPipe) id: string,
  ): Promise<ReportResponseDto> {
    const reportType =
      type === 'income' ? ReportType.income : ReportType.expense;

    return await this.reportService.getReportById(reportType, id);
  }

  @Post(':userId')
  async createReport(
    @Body() { source, amount }: CreateReportDto,
    @Param('type', new ParseEnumPipe(ReportType)) type: string,
    @Param('userId', ParseUUIDPipe) userId: string,
  ): Promise<ReportResponseDto> {
    //Get report type enum
    const reportType =
      type === 'income' ? ReportType.income : ReportType.expense;

    return await this.reportService.createReport(userId, reportType, {
      source,
      amount,
    });
  }

  @Put(':id')
  async updateReportById(
    @Param('type', new ParseEnumPipe(ReportType)) type: string,
    @Param('id', ParseUUIDPipe) id: string,
    @Body()
    body: UpdateReportDto,
  ): Promise<ReportResponseDto> {
    const reportType =
      type === 'income' ? ReportType.income : ReportType.expense;

    return this.reportService.updateReport(reportType, id, body);
  }

  @HttpCode(204)
  @Delete(':id')
  async deleteReportById(@Param('id', ParseUUIDPipe) id: string) {
    return this.reportService.deleteReport(id);
  }
}
