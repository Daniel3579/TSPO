package dtos

import "tspo/db"

type BalanceSheet struct {
	Symbol           string           `json:"symbol"`
	AnnualReports    []balanceReports `json:"annualReports"`
	QuarterlyReports []balanceReports `json:"quarterlyReports"`
}

type balanceReports struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            string `json:"totalAssets"`
	TotalCurrentAssets                     string `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  string `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            string `json:"cashAndShortTermInvestments"`
	Inventory                              string `json:"inventory"`
	CurrentNetReceivables                  string `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  string `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 string `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE string `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       string `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      string `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               string `json:"goodwill"`
	Investments                            string `json:"investments"`
	LongTermInvestments                    string `json:"longTermInvestments"`
	ShortTermInvestments                   string `json:"shortTermInvestments"`
	OtherCurrentAssets                     string `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  string `json:"otherNonCurrentAssets"`
	TotalLiabilities                       string `json:"totalLiabilities"`
	TotalCurrentLiabilities                string `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 string `json:"currentAccountsPayable"`
	DeferredRevenue                        string `json:"deferredRevenue"`
	CurrentDebt                            string `json:"currentDebt"`
	ShortTermDebt                          string `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             string `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                string `json:"capitalLeaseObligations"`
	LongTermDebt                           string `json:"longTermDebt"`
	CurrentLongTermDebt                    string `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 string `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 string `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                string `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             string `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 string `json:"totalShareholderEquity"`
	TreasuryStock                          string `json:"treasuryStock"`
	RetainedEarnings                       string `json:"retainedEarnings"`
	CommonStock                            string `json:"commonStock"`
	CommonStockSharesOutstanding           string `json:"commonStockSharesOutstanding"`
}

func (s *BalanceSheet) ToDatabaseableSlice() []db.Databaseable {
	symbol := s.Symbol
	qurterlyReports := s.QuarterlyReports
	slice := make([]db.Databaseable, len(qurterlyReports))

	for i, v := range qurterlyReports {
		instance := &balanceSheetDB{symbol, &v}
		slice[i] = instance
	}

	return slice
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type balanceSheetDB struct {
	Symbol string `json:"symbol"`
	*balanceReports
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––

func (req *balanceSheetDB) TableName() string {
	return "balance_sheet"
}

func (req *balanceSheetDB) Columns() []string {
	return []string{
		"symbol",
		"fiscalDateEnding",
		"reportedCurrency",
		"totalAssets",
		"totalCurrentAssets",
		"cashAndCashEquivalentsAtCarryingValue",
		"cashAndShortTermInvestments",
		"inventory",
		"currentNetReceivables",
		"totalNonCurrentAssets",
		"propertyPlantEquipment",
		"accumulatedDepreciationAmortizationPPE",
		"intangibleAssets",
		"intangibleAssetsExcludingGoodwill",
		"goodwill",
		"investments",
		"longTermInvestments",
		"shortTermInvestments",
		"otherCurrentAssets",
		"otherNonCurrentAssets",
		"totalLiabilities",
		"totalCurrentLiabilities",
		"currentAccountsPayable",
		"deferredRevenue",
		"currentDebt",
		"shortTermDebt",
		"totalNonCurrentLiabilities",
		"capitalLeaseObligations",
		"longTermDebt",
		"currentLongTermDebt",
		"longTermDebtNoncurrent",
		"shortLongTermDebtTotal",
		"otherCurrentLiabilities",
		"otherNonCurrentLiabilities",
		"totalShareholderEquity",
		"treasuryStock",
		"retainedEarnings",
		"commonStock",
		"commonStockSharesOutstanding",
	}
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––

func (req *balanceSheetDB) InsertableValues() []any {
	return []any{
		req.Symbol,
		req.FiscalDateEnding,
		req.ReportedCurrency,
		req.TotalAssets,
		req.TotalCurrentAssets,
		req.CashAndCashEquivalentsAtCarryingValue,
		req.CashAndShortTermInvestments,
		req.Inventory,
		req.CurrentNetReceivables,
		req.TotalNonCurrentAssets,
		req.PropertyPlantEquipment,
		req.AccumulatedDepreciationAmortizationPPE,
		req.IntangibleAssets,
		req.IntangibleAssetsExcludingGoodwill,
		req.Goodwill,
		req.Investments,
		req.LongTermInvestments,
		req.ShortTermInvestments,
		req.OtherCurrentAssets,
		req.OtherNonCurrentAssets,
		req.TotalLiabilities,
		req.TotalCurrentLiabilities,
		req.CurrentAccountsPayable,
		req.DeferredRevenue,
		req.CurrentDebt,
		req.ShortTermDebt,
		req.TotalNonCurrentLiabilities,
		req.CapitalLeaseObligations,
		req.LongTermDebt,
		req.CurrentLongTermDebt,
		req.LongTermDebtNoncurrent,
		req.ShortLongTermDebtTotal,
		req.OtherCurrentLiabilities,
		req.OtherNonCurrentLiabilities,
		req.TotalShareholderEquity,
		req.TreasuryStock,
		req.RetainedEarnings,
		req.CommonStock,
		req.CommonStockSharesOutstanding,
	}
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––

func (res *balanceSheetDB) SelectableValues() []any {
	return []any{
		&res.Symbol,
		&res.FiscalDateEnding,
		&res.ReportedCurrency,
		&res.TotalAssets,
		&res.TotalCurrentAssets,
		&res.CashAndCashEquivalentsAtCarryingValue,
		&res.CashAndShortTermInvestments,
		&res.Inventory,
		&res.CurrentNetReceivables,
		&res.TotalNonCurrentAssets,
		&res.PropertyPlantEquipment,
		&res.AccumulatedDepreciationAmortizationPPE,
		&res.IntangibleAssets,
		&res.IntangibleAssetsExcludingGoodwill,
		&res.Goodwill,
		&res.Investments,
		&res.LongTermInvestments,
		&res.ShortTermInvestments,
		&res.OtherCurrentAssets,
		&res.OtherNonCurrentAssets,
		&res.TotalLiabilities,
		&res.TotalCurrentLiabilities,
		&res.CurrentAccountsPayable,
		&res.DeferredRevenue,
		&res.CurrentDebt,
		&res.ShortTermDebt,
		&res.TotalNonCurrentLiabilities,
		&res.CapitalLeaseObligations,
		&res.LongTermDebt,
		&res.CurrentLongTermDebt,
		&res.LongTermDebtNoncurrent,
		&res.ShortLongTermDebtTotal,
		&res.OtherCurrentLiabilities,
		&res.OtherNonCurrentLiabilities,
		&res.TotalShareholderEquity,
		&res.TreasuryStock,
		&res.RetainedEarnings,
		&res.CommonStock,
		&res.CommonStockSharesOutstanding,
	}
}
