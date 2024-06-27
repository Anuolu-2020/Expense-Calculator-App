import { Module } from '@nestjs/common';
import { SummaryModule } from './summary/summary.module';
import { ReportModule } from './report/report.module';
import { ConfigModule } from '@nestjs/config';
import { PrismaService } from './prisma/prisma.service';
import * as Joi from 'joi';
import { AppService } from './app.service';
import { AppController } from './app.controller';
import {
  ThrottlerGuard,
  ThrottlerModule,
  hours,
  seconds,
} from '@nestjs/throttler';
import { ThrottlerStorageRedisService } from 'nestjs-throttler-storage-redis';
import { APP_GUARD } from '@nestjs/core';

@Module({
  imports: [
    ThrottlerModule.forRoot({
      throttlers: [
        {
          limit: Number(process.env.RATE_LIMIT_PER_MINUTE),
          ttl: seconds(Number(process.env.THROTTLE_TTL1)),
        },
        {
          limit: Number(process.env.RATE_LIMIT_PER_HOUR),
          ttl: hours(Number(process.env.THROTTLE_TTL2)),
        },
      ],
      storage: new ThrottlerStorageRedisService(process.env.REDIS_URL),
    }),
    SummaryModule,
    ReportModule,
    ConfigModule.forRoot({
      isGlobal: true,
      validationSchema: Joi.object({
        DATABASE_URL: Joi.string().required(),
        PORT: Joi.number().required(),
        REDIS_URL: Joi.string().required(),
        RATE_LIMIT_PER_MINUTE: Joi.string().required(),
        THROTTLE_TTL1: Joi.string().required(),
        RATE_LIMIT_PER_HOUR: Joi.string().required(),
        THROTTLE_TTL2: Joi.string().required(),
      }),
      validationOptions: {
        abortEarly: true,
      },
    }),
  ],
  controllers: [AppController],
  providers: [
    PrismaService,
    AppService,
    {
      provide: APP_GUARD,
      useClass: ThrottlerGuard,
    },
  ],
})
export class AppModule { }
