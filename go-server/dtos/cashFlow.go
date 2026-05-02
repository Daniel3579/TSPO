package dtos

import (
	"tspo/db"
	pb "tspo/proto/gen"
)

// Original structure
type CashFlow struct {
	Symbol           string            `json:"symbol"`
	AnnualReports    []CashFlowReports `json:"annualReports"`
	QuarterlyReports []CashFlowReports `json:"quarterlyReports"`
}

type CashFlowReports struct {
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

// Wrapper
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type CashFlowDB struct {
	Symbol string
	*CashFlowReports
}

// tableNameable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *CashFlowDB) TableName() string {
	return "cash_flow"
}

// Databaseable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *CashFlowDB) Columns() []string {
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

func (s *CashFlowDB) InsertableValues() []any {
	return []any{
		s.Symbol,
		s.FiscalDateEnding,
		s.ReportedCurrency,
		s.OperatingCashflow,
		s.PaymentsForOperatingActivities,
		s.ProceedsFromOperatingActivities,
		s.ChangeInOperatingLiabilities,
		s.ChangeInOperatingAssets,
		s.DepreciationDepletionAndAmortization,
		s.CapitalExpenditures,
		s.ChangeInReceivables,
		s.ChangeInInventory,
		s.ProfitLoss,
		s.CashflowFromInvestment,
		s.CashflowFromFinancing,
		s.ProceedsFromRepaymentsOfShortTermDebt,
		s.PaymentsForRepurchaseOfCommonStock,
		s.PaymentsForRepurchaseOfEquity,
		s.PaymentsForRepurchaseOfPreferredStock,
		s.DividendPayout,
		s.DividendPayoutCommonStock,
		s.DividendPayoutPreferredStock,
		s.ProceedsFromIssuanceOfCommonStock,
		s.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,
		s.ProceedsFromIssuanceOfPreferredStock,
		s.ProceedsFromRepurchaseOfEquity,
		s.ProceedsFromSaleOfTreasuryStock,
		s.StockBasedCompensation,
		s.ChangeInCashAndCashEquivalents,
		s.ChangeInExchangeRate,
		s.NetIncome,
	}
}

func (s *CashFlowDB) SelectableValues() []any {
	return []any{
		&s.Symbol,
		&s.FiscalDateEnding,
		&s.ReportedCurrency,
		&s.OperatingCashflow,
		&s.PaymentsForOperatingActivities,
		&s.ProceedsFromOperatingActivities,
		&s.ChangeInOperatingLiabilities,
		&s.ChangeInOperatingAssets,
		&s.DepreciationDepletionAndAmortization,
		&s.CapitalExpenditures,
		&s.ChangeInReceivables,
		&s.ChangeInInventory,
		&s.ProfitLoss,
		&s.CashflowFromInvestment,
		&s.CashflowFromFinancing,
		&s.ProceedsFromRepaymentsOfShortTermDebt,
		&s.PaymentsForRepurchaseOfCommonStock,
		&s.PaymentsForRepurchaseOfEquity,
		&s.PaymentsForRepurchaseOfPreferredStock,
		&s.DividendPayout,
		&s.DividendPayoutCommonStock,
		&s.DividendPayoutPreferredStock,
		&s.ProceedsFromIssuanceOfCommonStock,
		&s.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet,
		&s.ProceedsFromIssuanceOfPreferredStock,
		&s.ProceedsFromRepurchaseOfEquity,
		&s.ProceedsFromSaleOfTreasuryStock,
		&s.StockBasedCompensation,
		&s.ChangeInCashAndCashEquivalents,
		&s.ChangeInExchangeRate,
		&s.NetIncome,
	}
}

// Sliceable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *CashFlow) ToDatabaseableSlice() []db.Databaseable {
	symbol := s.Symbol
	qurterlyReports := s.QuarterlyReports
	slice := make([]db.Databaseable, len(qurterlyReports))

	for i, v := range qurterlyReports {
		instance := &CashFlowDB{symbol, &v}
		slice[i] = instance
	}

	return slice
}

// Grpcable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *CashFlowDB) ToResponse() any {
	v := s.InsertableValues()
	return &pb.SingleCashFlowResponse{
		Symbol:                                v[0].(string),
		Fiscaldateending:                      v[1].(string),
		Reportedcurrency:                      v[2].(string),
		Operatingcashflow:                     v[3].(string),
		Paymentsforoperatingactivities:        v[4].(string),
		Proceedsfromoperatingactivities:       v[5].(string),
		Changeinoperatingliabilities:          v[6].(string),
		Changeinoperatingassets:               v[7].(string),
		Depreciationdepletionandamortization:  v[8].(string),
		Capitalexpenditures:                   v[9].(string),
		Changeinreceivables:                   v[10].(string),
		Changeininventory:                     v[11].(string),
		Profitloss:                            v[12].(string),
		Cashflowfrominvestment:                v[13].(string),
		Cashflowfromfinancing:                 v[14].(string),
		Proceedsfromrepaymentsofshorttermdebt: v[15].(string),
		Paymentsforrepurchaseofcommonstock:    v[16].(string),
		Paymentsforrepurchaseofequity:         v[17].(string),
		Paymentsforrepurchaseofpreferredstock: v[18].(string),
		Dividendpayout:                        v[19].(string),
		Dividendpayoutcommonstock:             v[20].(string),
		Dividendpayoutpreferredstock:          v[21].(string),
		Proceedsfromissuanceofcommonstock:     v[22].(string),
		Proceedsfromissuanceoflongtermdebtandcapitalsecuritiesnet: v[23].(string),
		Proceedsfromissuanceofpreferredstock:                      v[24].(string),
		Proceedsfromrepurchaseofequity:                            v[25].(string),
		Proceedsfromsaleoftreasurystock:                           v[26].(string),
		Stockbasedcompensation:                                    v[27].(string),
		Changeincashandcashequivalents:                            v[28].(string),
		Changeinexchangerate:                                      v[29].(string),
		Netincome:                                                 v[30].(string),
	}
}
