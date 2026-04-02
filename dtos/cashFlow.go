package dtos

import "tspo/db"

type CashFlow struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []cashFlowReports `json:"annualReports"`
	QuarterlyReports []cashFlowReports `json:"quarterlyReports"`
}

type cashFlowReports struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         string `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            string `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              string `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   string `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      string `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       string `json:"capitalExpenditures"`
	ChangeInReceivables                                       string `json:"changeInReceivables"`
	ChangeInInventory                                         string `json:"changeInInventory"`
	ProfitLoss                                                string `json:"profitLoss"`
	CashflowFromInvestment                                    string `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     string `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     string `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            string `json:"dividendPayout"`
	DividendPayoutCommonStock                                 string `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            string `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`
	StockBasedCompensation                                    string `json:"stockBasedCompensation"`
	ChangeInCashAndCashEquivalents                            string `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      string `json:"changeInExchangeRate"`
	NetIncome                                                 string `json:"netIncome"`
}

func (s *CashFlow) ToDatabaseableSlice() []db.Databaseable {
	symbol := s.Symbol
	qurterlyReports := s.QuarterlyReports
	slice := make([]db.Databaseable, len(qurterlyReports))

	for i, v := range qurterlyReports {
		instance := &cashFlowDB{symbol, &v}
		slice[i] = instance
	}

	return slice
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type cashFlowDB struct {
	Symbol string `json:"symbol"`
	*cashFlowReports
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (req *cashFlowDB) TableName() string {
	return "cash_flow"
}

func (req *cashFlowDB) Columns() []string {
	return []string{
		"symbol",
		"fiscalDateEnding",
		"reportedCurrency",
		"operatingCashflow",
		"paymentsForOperatingActivities",
		"proceedsFromOperatingActivities",
		"changeInOperatingLiabilities",
		"changeInOperatingAssets",
		"depreciationDepletionAndAmortization",
		"capitalExpenditures",
		"changeInReceivables",
		"changeInInventory",
		"profitLoss",
		"cashflowFromInvestment",
		"cashflowFromFinancing",
		"proceedsFromRepaymentsOfShortTermDebt",
		"paymentsForRepurchaseOfCommonStock",
		"paymentsForRepurchaseOfEquity",
		"paymentsForRepurchaseOfPreferredStock",
		"dividendPayout",
		"dividendPayoutCommonStock",
		"dividendPayoutPreferredStock",
		"proceedsFromIssuanceOfCommonStock",
		"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet",
		"proceedsFromIssuanceOfPreferredStock",
		"proceedsFromRepurchaseOfEquity",
		"proceedsFromSaleOfTreasuryStock",
		"stockBasedCompensation",
		"changeInCashAndCashEquivalents",
		"changeInExchangeRate",
		"netIncome",
	}
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (req *cashFlowDB) InsertableValues() []any {
	return []any{
		req.Symbol,
		req.FiscalDateEnding,
		req.ReportedCurrency,
		req.OperatingCashflow,
		req.PaymentsForOperatingActivities,
		req.ProceedsFromOperatingActivities,
		req.ChangeInOperatingLiabilities,
		req.ChangeInOperatingAssets,
		req.DepreciationDepletionAndAmortization,
		req.CapitalExpenditures,
		req.ChangeInReceivables,
		req.ChangeInInventory,
		req.ProfitLoss,
		req.CashflowFromInvestment,
		req.CashflowFromFinancing,
		req.ProceedsFromRepaymentsOfShortTermDebt,
		req.PaymentsForRepurchaseOfCommonStock,
		req.PaymentsForRepurchaseOfEquity,
		req.PaymentsForRepurchaseOfPreferredStock,
		req.DividendPayout,
		req.DividendPayoutCommonStock,
		req.DividendPayoutPreferredStock,
		req.ProceedsFromIssuanceOfCommonStock,
		req.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,
		req.ProceedsFromIssuanceOfPreferredStock,
		req.ProceedsFromRepurchaseOfEquity,
		req.ProceedsFromSaleOfTreasuryStock,
		req.StockBasedCompensation,
		req.ChangeInCashAndCashEquivalents,
		req.ChangeInExchangeRate,
		req.NetIncome,
	}
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (res *cashFlowDB) SelectableValues() []any {
	return []any{
		&res.Symbol,
		&res.FiscalDateEnding,
		&res.ReportedCurrency,
		&res.OperatingCashflow,
		&res.PaymentsForOperatingActivities,
		&res.ProceedsFromOperatingActivities,
		&res.ChangeInOperatingLiabilities,
		&res.ChangeInOperatingAssets,
		&res.DepreciationDepletionAndAmortization,
		&res.CapitalExpenditures,
		&res.ChangeInReceivables,
		&res.ChangeInInventory,
		&res.ProfitLoss,
		&res.CashflowFromInvestment,
		&res.CashflowFromFinancing,
		&res.ProceedsFromRepaymentsOfShortTermDebt,
		&res.PaymentsForRepurchaseOfCommonStock,
		&res.PaymentsForRepurchaseOfEquity,
		&res.PaymentsForRepurchaseOfPreferredStock,
		&res.DividendPayout,
		&res.DividendPayoutCommonStock,
		&res.DividendPayoutPreferredStock,
		&res.ProceedsFromIssuanceOfCommonStock,
		&res.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,
		&res.ProceedsFromIssuanceOfPreferredStock,
		&res.ProceedsFromRepurchaseOfEquity,
		&res.ProceedsFromSaleOfTreasuryStock,
		&res.StockBasedCompensation,
		&res.ChangeInCashAndCashEquivalents,
		&res.ChangeInExchangeRate,
		&res.NetIncome,
	}
}
