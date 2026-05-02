package rtos

import (
	"fmt"
	"tspo/db"
	pb "tspo/proto/gen"
)

// Wrapper
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––
type FinancialRequestDB[T db.Databaseable] struct {
	Table              string
	NewRowDatabaseable func() T
	*pb.FinancialRequest
}

// tableNameable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––
func (req *FinancialRequestDB[T]) TableName() string {
	return req.Table
}

// Requestable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––
func (req *FinancialRequestDB[T]) Params() (string, error) {
	symbol := req.GetSymbol()
	if symbol == "" {
		return "", fmt.Errorf("Не передан параметр \"symbol\"")
	}

	date := req.GetDate()
	dateRange := req.GetRange()
	if date == nil && dateRange == nil {
		return "", fmt.Errorf("Не передан параметр \"timeframe\"")
	}
	var params string

	if date != nil {
		quarter := date.GetQuarter()
		if quarter == 0 {
			return "", fmt.Errorf("Не передан параметр \"quarter\"")
		}

		year := date.GetYear()
		if year == 0 {
			return "", fmt.Errorf("Не передан параметр \"year\"")
		}

		start := fmt.Sprintf("%d-%02d", year, quarter*3-2)
		end := fmt.Sprintf("%d-%02d", year, quarter*3+1)

		params = fmt.Sprintf("'%s' and '%s'", start, end)
	}

	if dateRange != nil {
		startQuarter := dateRange.GetStart().GetQuarter()
		if startQuarter == 0 {
			return "", fmt.Errorf("Не передан параметр \"startQuarter\"")
		}

		startYear := dateRange.GetStart().GetYear()
		if startYear == 0 {
			return "", fmt.Errorf("Не передан параметр \"startYear\"")
		}

		endQuarter := dateRange.GetEnd().GetQuarter()
		if endQuarter == 0 {
			return "", fmt.Errorf("Не передан параметр \"endQuarter\"")
		}

		endYear := dateRange.GetEnd().GetYear()
		if endYear == 0 {
			return "", fmt.Errorf("Не передан параметр \"endYear\"")
		}

		start := fmt.Sprintf("%d-%02d", startYear, startQuarter*3-2)
		end := fmt.Sprintf("%d-%02d", endYear, endQuarter*3+1)

		params = fmt.Sprintf("'%s' and '%s'", start, end)
	}

	return fmt.Sprintf("Where symbol='%s' and fiscalDateEnding between %s", symbol, params), nil
}

func (req *FinancialRequestDB[T]) GetDatabaseable() T {
	return req.NewRowDatabaseable()
}

// 2024-01-01-00:00 on — 2024-04-01-00:00 off Q1
// 2024-04-01-00:00 on — 2024-07-01-00:00 off Q2
// 2024-07-01-00:00 on — 2024-10-01-00:00 off Q3
// 2024-10-01-00:00 on — 2024-13-01-00:00 off Q4

// 2024-01 — 2024-04 Q1 Q*3-2 – Q*3+1
// 2024-04 — 2024-07 Q2 Q*3-2 – Q*3+1
// 2024-07 — 2024-10 Q3 Q*3-2 – Q*3+1
// 2024-10 — 2024-13 Q4 Q*3-2 – Q*3+1
