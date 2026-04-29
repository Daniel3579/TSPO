package dtos

import (
	"tspo/db"
	pb "tspo/proto/gen"
)

// Original structure
type BalanceSheet struct {
	Symbol           string           `json:"symbol"`
	AnnualReports    []BalanceReports `json:"annualReports"`
	QuarterlyReports []BalanceReports `json:"quarterlyReports"`
}

type BalanceReports struct {
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

// Wrapper
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type BalanceSheetDB struct {
	Symbol string
	*BalanceReports
}

// tableNameable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *BalanceSheetDB) TableName() string {
	return "balance_sheet"
}

// Databaseable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *BalanceSheetDB) Columns() []string {
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

func (s *BalanceSheetDB) InsertableValues() []any {
	return []any{
		s.Symbol,
		s.FiscalDateEnding,
		s.ReportedCurrency,
		s.TotalAssets,
		s.TotalCurrentAssets,
		s.CashAndCashEquivalentsAtCarryingValue,
		s.CashAndShortTermInvestments,
		s.Inventory,
		s.CurrentNetReceivables,
		s.TotalNonCurrentAssets,
		s.PropertyPlantEquipment,
		s.AccumulatedDepreciationAmortizationPPE,
		s.IntangibleAssets,
		s.IntangibleAssetsExcludingGoodwill,
		s.Goodwill,
		s.Investments,
		s.LongTermInvestments,
		s.ShortTermInvestments,
		s.OtherCurrentAssets,
		s.OtherNonCurrentAssets,
		s.TotalLiabilities,
		s.TotalCurrentLiabilities,
		s.CurrentAccountsPayable,
		s.DeferredRevenue,
		s.CurrentDebt,
		s.ShortTermDebt,
		s.TotalNonCurrentLiabilities,
		s.CapitalLeaseObligations,
		s.LongTermDebt,
		s.CurrentLongTermDebt,
		s.LongTermDebtNoncurrent,
		s.ShortLongTermDebtTotal,
		s.OtherCurrentLiabilities,
		s.OtherNonCurrentLiabilities,
		s.TotalShareholderEquity,
		s.TreasuryStock,
		s.RetainedEarnings,
		s.CommonStock,
		s.CommonStockSharesOutstanding,
	}
}

func (s *BalanceSheetDB) SelectableValues() []any {
	return []any{
		&s.Symbol,
		&s.FiscalDateEnding,
		&s.ReportedCurrency,
		&s.TotalAssets,
		&s.TotalCurrentAssets,
		&s.CashAndCashEquivalentsAtCarryingValue,
		&s.CashAndShortTermInvestments,
		&s.Inventory,
		&s.CurrentNetReceivables,
		&s.TotalNonCurrentAssets,
		&s.PropertyPlantEquipment,
		&s.AccumulatedDepreciationAmortizationPPE,
		&s.IntangibleAssets,
		&s.IntangibleAssetsExcludingGoodwill,
		&s.Goodwill,
		&s.Investments,
		&s.LongTermInvestments,
		&s.ShortTermInvestments,
		&s.OtherCurrentAssets,
		&s.OtherNonCurrentAssets,
		&s.TotalLiabilities,
		&s.TotalCurrentLiabilities,
		&s.CurrentAccountsPayable,
		&s.DeferredRevenue,
		&s.CurrentDebt,
		&s.ShortTermDebt,
		&s.TotalNonCurrentLiabilities,
		&s.CapitalLeaseObligations,
		&s.LongTermDebt,
		&s.CurrentLongTermDebt,
		&s.LongTermDebtNoncurrent,
		&s.ShortLongTermDebtTotal,
		&s.OtherCurrentLiabilities,
		&s.OtherNonCurrentLiabilities,
		&s.TotalShareholderEquity,
		&s.TreasuryStock,
		&s.RetainedEarnings,
		&s.CommonStock,
		&s.CommonStockSharesOutstanding,
	}
}

// Sliceable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *BalanceSheet) ToDatabaseableSlice() []db.Databaseable {
	symbol := s.Symbol
	qurterlyReports := s.QuarterlyReports
	slice := make([]db.Databaseable, len(qurterlyReports))

	for i, v := range qurterlyReports {
		instance := &BalanceSheetDB{symbol, &v}
		slice[i] = instance
	}

	return slice
}

// Grpcable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *BalanceSheetDB) ToResponse() any {
	v := s.InsertableValues()
	return &pb.SingleBalanceSheetResponse{
		Symbol:                                 v[0].(string),
		Fiscaldateending:                       v[1].(string),
		Reportedcurrency:                       v[2].(string),
		Totalassets:                            v[3].(string),
		Totalcurrentassets:                     v[4].(string),
		Cashandcashequivalentsatcarryingvalue:  v[5].(string),
		Cashandshortterminvestments:            v[6].(string),
		Inventory:                              v[7].(string),
		Currentnetreceivables:                  v[8].(string),
		Totalnoncurrentassets:                  v[9].(string),
		Propertyplantequipment:                 v[10].(string),
		Accumulateddepreciationamortizationppe: v[11].(string),
		Intangibleassets:                       v[12].(string),
		Intangibleassetsexcludinggoodwill:      v[13].(string),
		Goodwill:                               v[14].(string),
		Investments:                            v[15].(string),
		Longterminvestments:                    v[16].(string),
		Shortterminvestments:                   v[17].(string),
		Othercurrentassets:                     v[18].(string),
		Othernoncurrentassets:                  v[19].(string),
		Totalliabilities:                       v[20].(string),
		Totalcurrentliabilities:                v[21].(string),
		Currentaccountspayable:                 v[22].(string),
		Deferredrevenue:                        v[23].(string),
		Currentdebt:                            v[24].(string),
		Shorttermdebt:                          v[25].(string),
		Totalnoncurrentliabilities:             v[26].(string),
		Capitalleaseobligations:                v[27].(string),
		Longtermdebt:                           v[28].(string),
		Currentlongtermdebt:                    v[29].(string),
		Longtermdebtnoncurrent:                 v[30].(string),
		Shortlongtermdebttotal:                 v[31].(string),
		Othercurrentliabilities:                v[32].(string),
		Othernoncurrentliabilities:             v[33].(string),
		Totalshareholderequity:                 v[34].(string),
		Treasurystock:                          v[35].(string),
		Retainedearnings:                       v[36].(string),
		Commonstock:                            v[37].(string),
		Commonstocksharesoutstanding:           v[38].(string),
	}
}
