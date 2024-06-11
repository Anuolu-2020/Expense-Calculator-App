import { Injectable, NotFoundException } from '@nestjs/common';
import { ReportType } from 'src/dtos/report.dto';
import { ReportResponseDto } from '../dtos/report.dto';
import { PrismaService } from 'src/prisma/prisma.service';

export interface Report {
  source: string;
  amount: number;
}

interface UpdateReport {
  source?: string;
  amount?: number;
}

@Injectable()
export class ReportService {
  constructor(private readonly prisma: PrismaService) { }
  async getAReport(userId: string, id: string) {
    const report = await this.prisma.reports.findFirst({
      where: { user_id: userId, id },
    });

    return report;
  }

  async updateAReport(
    userId: string,
    id: string,
    { source, amount }: UpdateReport,
  ) {
    const report = await this.prisma.reports.update({
      where: { user_id: userId, id },
      data: { source, amount },
    });

    return report;
  }

  async getAllReportsByUserId(userId: string): Promise<any> {
    const report = await this.prisma.reports.findMany({
      where: { user_id: userId },
    });

    return report;
  }

  async getAllReportsType(type: ReportType): Promise<ReportResponseDto[]> {
    const report = await this.prisma.reports.findMany({
      where: {
        type,
      },
    });

    /*if (report.length === 0) {
      return { message: 'No Report Record Yet' };
    }*/

    return report;
  }

  async getReportTypeByUserId(
    type: ReportType,
    id: string,
  ): Promise<ReportResponseDto[]> {
    console.log(`User Id: ${id}`);

    const report = await this.prisma.reports.findMany({
      where: {
        type,
        user_id: id,
      },
    });

    /*if (report.length === 0) {
      return { message: 'No Report Records Yet' };
    }*/

    return report;
  }

  async createReport(
    userId: string,
    type: ReportType,
    { source, amount }: Report,
  ): Promise<ReportResponseDto> {
    //save report
    const report = await this.prisma.reports.create({
      data: {
        source,
        amount,
        type,
        user_id: userId,
      },
    });

    return report;
  }

  async updateReport(
    type: ReportType,
    id: string,
    data: UpdateReport,
  ): Promise<ReportResponseDto> {
    const report = await this.prisma.reports.findUnique({
      where: { id, type },
    });

    if (!report) {
      throw new NotFoundException('Report not found');
    }

    const newReport = await this.prisma.reports.update({
      where: { id, type },
      data,
    });

    return newReport;
  }

  async deleteReport(id: string) {
    const report = await this.prisma.reports.findUnique({ where: { id } });

    if (!report) {
      throw new NotFoundException('Report not found');
    }

    return await this.prisma.reports.delete({ where: { id } });
  }
}
