import { Module } from '@nestjs/common';
import { SummaryModule } from './summary/summary.module';
import { ReportModule } from './report/report.module';
import { ConfigModule } from '@nestjs/config';
import { PrismaService } from './prisma/prisma.service';
import * as Joi from 'joi';
import { AppService } from './app.service';
import { AppController } from './app.controller';

@Module({
  imports: [
    SummaryModule,
    ReportModule,
    ConfigModule.forRoot({
      isGlobal: true,
      validationSchema: Joi.object({
        DATABASE_URL: Joi.string().required(),
        PORT: Joi.number().required(),
      }),
      validationOptions: {
        abortEarly: true,
      },
    }),
  ],
  controllers: [AppController],
  providers: [PrismaService, AppService],
})
export class AppModule { }
