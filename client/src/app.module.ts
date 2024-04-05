import { ClassSerializerInterceptor, Module } from '@nestjs/common';
//import { APP_INTERCEPTOR } from '@nestjs/core';
import { SummaryModule } from './summary/summary.module';
import { ReportModule } from './report/report.module';

@Module({
  imports: [SummaryModule, ReportModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
