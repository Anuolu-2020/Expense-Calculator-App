import { Controller, Get, Param, ParseUUIDPipe } from '@nestjs/common';
import { SummaryService } from './summary.service';

@Controller('summary')
export class SummaryController {
  constructor(private readonly summaryService: SummaryService) { }
  @Get(':userId')
  getSummary(@Param('userId', ParseUUIDPipe) id: string) {
    return this.summaryService.calculateSummary(id);
  }
}
