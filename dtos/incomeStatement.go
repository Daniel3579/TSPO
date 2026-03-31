package dtos

type IncomeStatement struct {
	Symbol           string          `json:"symbol"`
	AnnualReports    []IncomeReports `json:"annualReports"`
	QuarterlyReports []IncomeReports `json:"quarterlyReports"`
}

type IncomeReports struct {
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

func (s IncomeReports) GetValues() []interface{} {
	return []interface{}{
		s.FiscalDateEnding,
		s.ReportedCurrency,
		s.GrossProfit,
		s.TotalRevenue,
		s.CostOfRevenue,
		s.CostofGoodsAndServicesSold,
		s.OperatingIncome,
		s.SellingGeneralAndAdministrative,
		s.ResearchAndDevelopment,
		s.OperatingExpenses,
		s.InvestmentIncomeNet,
		s.NetInterestIncome,
		s.InterestIncome,
		s.InterestExpense,
		s.NonInterestIncome,
		s.OtherNonOperatingIncome,
		s.Depreciation,
		s.DepreciationAndAmortization,
		s.IncomeBeforeTax,
		s.IncomeTaxExpense,
		s.InterestAndDebtExpense,
		s.NetIncomeFromContinuingOperations,
		s.ComprehensiveIncomeNetOfTax,
		s.Ebit,
		s.Ebitda,
		s.NetIncome,
	}
}

func (req *IncomeStatement) TableName() string {
	return "income_statement"
}

func (req *IncomeStatement) Columns() []string {
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

// func (req *IncomeStatement) Values() []interface{} {
// 	req.QuarterlyReports

// 	return []interface{}{
// 		req.Symbol,
// 		req.QuarterlyReports
// 	}
// }
