/*
  Warnings:

  - Changed the type of `type` on the `Report` table. No cast exists, the column would be dropped and recreated, which cannot be done if there is data, since the column is required.

*/
-- CreateEnum
CREATE TYPE "ReportType" AS ENUM ('EXPENSE', 'INCOME');

-- AlterTable
ALTER TABLE "Report" DROP COLUMN "type",
ADD COLUMN     "type" "ReportType" NOT NULL;

-- DropEnum
DROP TYPE "REPORTTYPE";
