import { Module } from '@nestjs/common';
//import { APP_INTERCEPTOR } from '@nestjs/core';
import { SummaryModule } from './summary/summary.module';
import { ReportModule } from './report/report.module';
import { ConfigModule } from '@nestjs/config';
import { PrismaService } from './prisma/prisma.service';
import * as Joi from 'joi';

@Module({
  imports: [
    SummaryModule,
    ReportModule,
    ConfigModule.forRoot({
      isGlobal: true,
      validationSchema: Joi.object({ DATABASE_URL: Joi.string().required() }),
      validationOptions: {
        abortEarly: true,
      },
    }),
  ],
  controllers: [],
  providers: [PrismaService],
})
export class AppModule { }
