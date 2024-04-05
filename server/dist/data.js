"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.data = exports.ReportType = void 0;
var ReportType;
(function (ReportType) {
    ReportType["INCOME"] = "income";
    ReportType["EXPENSE"] = "expense";
})(ReportType || (exports.ReportType = ReportType = {}));
exports.data = {
    report: [{
            id: "uuid1",
            source: "salary",
            amount: 2500,
            created_at: new Date,
            updated_at: new Date,
            type: ReportType.INCOME
        }, {
            id: "uuid2",
            source: "Youtube",
            amount: 7500,
            created_at: new Date,
            updated_at: new Date,
            type: ReportType.INCOME
        }, {
            id: "uuid3",
            source: "salary",
            amount: 4500,
            created_at: new Date,
            updated_at: new Date,
            type: ReportType.INCOME
        }],
};
//# sourceMappingURL=data.js.map