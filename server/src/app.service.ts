import { Injectable } from '@nestjs/common';

@Injectable()
export class AppService {
  getHelloWorld() {
    return { message: 'Welcome to Expense App Server' };
  }
}
