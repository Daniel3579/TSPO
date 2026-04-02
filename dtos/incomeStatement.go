package dtos

import "tspo/db"

type IncomeStatement struct {
	Symbol           string          `json:"symbol"`
	AnnualReports    []incomeReports `json:"annualReports"`
	QuarterlyReports []incomeReports `json:"quarterlyReports"`
}

type incomeReports struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       string `json:"grossProfit"`
	TotalRevenue                      string `json:"totalRevenue"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	CostofGoodsAndServicesSold        string `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   string `json:"operatingIncome"`
	SellingGeneralAndAdministrative   string `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment"`
	OperatingExpenses                 string `json:"operatingExpenses"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	InterestIncome                    string `json:"interestIncome"`
	InterestExpense                   string `json:"interestExpense"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	Depreciation                      string `json:"depreciation"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	InterestAndDebtExpense            string `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              string `json:"ebit"`
	Ebitda                            string `json:"ebitda"`
	NetIncome                         string `json:"netIncome"`
}

func (s *IncomeStatement) ToDatabaseableSlice() []db.Databaseable {
	symbol := s.Symbol
	qurterlyReports := s.QuarterlyReports
	slice := make([]db.Databaseable, len(qurterlyReports))

	for i, v := range qurterlyReports {
		instance := &incomeStatementDB{symbol, &v}
		slice[i] = instance
	}

	return slice
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type incomeStatementDB struct {
	Symbol string `json:"symbol"`
	*incomeReports
}

// –––––––––––––––––––––––––––––––––––––––––––––––––––––

func (req *incomeStatementDB) TableName() string {
	return "income_statement"
}

func (req *incomeStatementDB) Columns() []string {
	return []string{
		"symbol",
		"fiscalDateEnding",
		"reportedCurrency",
		"grossProfit",
		"totalRevenue",
		"costOfRevenue",
		"costofGoodsAndServicesSold",
		"operatingIncome",
		"sellingGeneralAndAdministrative",
		"researchAndDevelopment",
		"operatingExpenses",
		"investmentIncomeNet",
		"netInterestIncome",
		"interestIncome",
		"interestExpense",
		"nonInterestIncome",
		"otherNonOperatingIncome",
		"depreciation",
		"depreciationAndAmortization",
		"incomeBeforeTax",
		"incomeTaxExpense",
		"interestAndDebtExpense",
		"netIncomeFromContinuingOperations",
		"comprehensiveIncomeNetOfTax",
		"ebit",
		"ebitda",
		"netIncome",
	}
}

// –––––––––––––––––––––––––––––––––––––––––––––––––––––

func (req *incomeStatementDB) InsertableValues() []any {
	return []any{
		req.Symbol,
		req.FiscalDateEnding,
		req.ReportedCurrency,
		req.GrossProfit,
		req.TotalRevenue,
		req.CostOfRevenue,
		req.CostofGoodsAndServicesSold,
		req.OperatingIncome,
		req.SellingGeneralAndAdministrative,
		req.ResearchAndDevelopment,
		req.OperatingExpenses,
		req.InvestmentIncomeNet,
		req.NetInterestIncome,
		req.InterestIncome,
		req.InterestExpense,
		req.NonInterestIncome,
		req.OtherNonOperatingIncome,
		req.Depreciation,
		req.DepreciationAndAmortization,
		req.IncomeBeforeTax,
		req.IncomeTaxExpense,
		req.InterestAndDebtExpense,
		req.NetIncomeFromContinuingOperations,
		req.ComprehensiveIncomeNetOfTax,
		req.Ebit,
		req.Ebitda,
		req.NetIncome,
	}
}

// –––––––––––––––––––––––––––––––––––––––––––––––––––––

func (res *incomeStatementDB) SelectableValues() []any {
	return []any{
		&res.Symbol,
		&res.FiscalDateEnding,
		&res.ReportedCurrency,
		&res.GrossProfit,
		&res.TotalRevenue,
		&res.CostOfRevenue,
		&res.CostofGoodsAndServicesSold,
		&res.OperatingIncome,
		&res.SellingGeneralAndAdministrative,
		&res.ResearchAndDevelopment,
		&res.OperatingExpenses,
		&res.InvestmentIncomeNet,
		&res.NetInterestIncome,
		&res.InterestIncome,
		&res.InterestExpense,
		&res.NonInterestIncome,
		&res.OtherNonOperatingIncome,
		&res.Depreciation,
		&res.DepreciationAndAmortization,
		&res.IncomeBeforeTax,
		&res.IncomeTaxExpense,
		&res.InterestAndDebtExpense,
		&res.NetIncomeFromContinuingOperations,
		&res.ComprehensiveIncomeNetOfTax,
		&res.Ebit,
		&res.Ebitda,
		&res.NetIncome,
	}
}

// func NewIncomeStatementDB(symbol string, vals ...any) *incomeStatementDB {
// 	i := &incomeStatementDB{Symbol: symbol}
//
// 	i.FiscalDateEnding = vals[0].(string)
// 	i.ReportedCurrency = vals[1].(string)
// 	i.GrossProfit = vals[2].(string)
// 	i.TotalRevenue = vals[3].(string)
// 	i.CostOfRevenue = vals[4].(string)
// 	i.CostofGoodsAndServicesSold = vals[5].(string)
// 	i.OperatingIncome = vals[6].(string)
// 	i.SellingGeneralAndAdministrative = vals[7].(string)
// 	i.ResearchAndDevelopment = vals[8].(string)
// 	i.OperatingExpenses = vals[9].(string)
// 	i.InvestmentIncomeNet = vals[10].(string)
// 	i.NetInterestIncome = vals[11].(string)
// 	i.InterestIncome = vals[12].(string)
// 	i.InterestExpense = vals[13].(string)
// 	i.NonInterestIncome = vals[14].(string)
// 	i.OtherNonOperatingIncome = vals[15].(string)
// 	i.Depreciation = vals[16].(string)
// 	i.DepreciationAndAmortization = vals[17].(string)
// 	i.IncomeBeforeTax = vals[18].(string)
// 	i.IncomeTaxExpense = vals[19].(string)
// 	i.InterestAndDebtExpense = vals[20].(string)
// 	i.NetIncomeFromContinuingOperations = vals[21].(string)
// 	i.ComprehensiveIncomeNetOfTax = vals[22].(string)
// 	i.Ebit = vals[23].(string)
// 	i.Ebitda = vals[24].(string)
// 	i.NetIncome = vals[25].(string)
//
// 	return i
// }

// func (s *incomeReports) getValues() []any {
// 	return []any{
// 		s.FiscalDateEnding,
// 		s.ReportedCurrency,
// 		s.GrossProfit,
// 		s.TotalRevenue,
// 		s.CostOfRevenue,
// 		s.CostofGoodsAndServicesSold,
// 		s.OperatingIncome,
// 		s.SellingGeneralAndAdministrative,
// 		s.ResearchAndDevelopment,
// 		s.OperatingExpenses,
// 		s.InvestmentIncomeNet,
// 		s.NetInterestIncome,
// 		s.InterestIncome,
// 		s.InterestExpense,
// 		s.NonInterestIncome,
// 		s.OtherNonOperatingIncome,
// 		s.Depreciation,
// 		s.DepreciationAndAmortization,
// 		s.IncomeBeforeTax,
// 		s.IncomeTaxExpense,
// 		s.InterestAndDebtExpense,
// 		s.NetIncomeFromContinuingOperations,
// 		s.ComprehensiveIncomeNetOfTax,
// 		s.Ebit,
// 		s.Ebitda,
// 		s.NetIncome,
// 	}
// }

// func (s *IncomeStatement) ToDatabaseableSlice() []Databaseable {
// 	symbol := s.Symbol
// 	qurterlyReports := s.QuarterlyReports
// 	slice := make([]Databaseable, len(qurterlyReports))
//
// 	for i, v := range qurterlyReports {
// 		values := v.getValues()
// 		instance := NewIncomeStatementDB(symbol, values...)
// 		slice[i] = instance
// 	}
//
// 	return slice
// }
