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

@Controller('report')
export class ReportController {
  constructor(private readonly reportService: ReportService) { }
  @Get(':userId')
  async getAllReportsByUserId(
    @Param('userId', ParseUUIDPipe) userId: string,
  ): Promise<any> {
    const report = await this.reportService.getAllReportsByUserId(userId);

    return { results: report };
  }

  @Get(':userId/:id')
  async getAReport(
    @Param('userId', ParseUUIDPipe) userId: string,
    @Param('id', ParseUUIDPipe) id: string,
  ) {
    const report = await this.reportService.getAReport(userId, id);

    return report;
  }

  @Put(':userId/:id')
  async updateAReport(
    @Body() data: UpdateReportDto,
    @Param('userId', ParseUUIDPipe) userId: string,
    @Param('id', ParseUUIDPipe) id: string,
  ) {
    const report = await this.reportService.updateAReport(userId, id, data);

    return report;
  }

  @Get('type/:reportType')
  async getAllReportsType(
    @Param('reportType', new ParseEnumPipe(ReportType)) type: string,
  ): Promise<any> {
    const reportType =
      type === 'income' ? ReportType.income : ReportType.expense;

    const report = await this.reportService.getAllReportsType(reportType);

    return { results: report };
  }

  @Get('type/:reportType/:userId')
  async getReportTypeByUserId(
    @Param('reportType', new ParseEnumPipe(ReportType)) type: string,
    @Param('userId', ParseUUIDPipe) userId: string,
  ): Promise<any> {
    const reportType =
      type === 'income' ? ReportType.income : ReportType.expense;

    const report = await this.reportService.getReportTypeByUserId(
      reportType,
      userId,
    );

    return { results: report };
  }

  @Post('type/:reportType/:userId')
  async createReport(
    @Body() { source, amount }: CreateReportDto,
    @Param('reportType', new ParseEnumPipe(ReportType)) type: string,
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

  @Put('type/:reportType/:id')
  async updateReportById(
    @Param('reportType', new ParseEnumPipe(ReportType)) type: string,
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
