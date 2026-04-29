package dtos

import (
	"tspo/db"
	pb "tspo/proto/gen"
)

// Original structure
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

// Wrapper
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type IncomeStatementDB struct {
	Symbol string
	*IncomeReports
}

// tableNameable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *IncomeStatementDB) TableName() string {
	return "income_statement"
}

// Databaseable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *IncomeStatementDB) Columns() []string {
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

func (s *IncomeStatementDB) InsertableValues() []any {
	return []any{
		s.Symbol,
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

func (s *IncomeStatementDB) SelectableValues() []any {
	return []any{
		&s.Symbol,
		&s.FiscalDateEnding,
		&s.ReportedCurrency,
		&s.GrossProfit,
		&s.TotalRevenue,
		&s.CostOfRevenue,
		&s.CostofGoodsAndServicesSold,
		&s.OperatingIncome,
		&s.SellingGeneralAndAdministrative,
		&s.ResearchAndDevelopment,
		&s.OperatingExpenses,
		&s.InvestmentIncomeNet,
		&s.NetInterestIncome,
		&s.InterestIncome,
		&s.InterestExpense,
		&s.NonInterestIncome,
		&s.OtherNonOperatingIncome,
		&s.Depreciation,
		&s.DepreciationAndAmortization,
		&s.IncomeBeforeTax,
		&s.IncomeTaxExpense,
		&s.InterestAndDebtExpense,
		&s.NetIncomeFromContinuingOperations,
		&s.ComprehensiveIncomeNetOfTax,
		&s.Ebit,
		&s.Ebitda,
		&s.NetIncome,
	}
}

// Sliceable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *IncomeStatement) ToDatabaseableSlice() []db.Databaseable {
	symbol := s.Symbol
	qurterlyReports := s.QuarterlyReports
	slice := make([]db.Databaseable, len(qurterlyReports))

	for i, v := range qurterlyReports {
		instance := &IncomeStatementDB{symbol, &v}
		slice[i] = instance
	}

	return slice
}

// Grpcable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *IncomeStatementDB) ToResponse() any {
	v := s.InsertableValues()
	return &pb.SingleIncomeStatementResponse{
		Symbol:                            v[0].(string),
		Fiscaldateending:                  v[1].(string),
		Reportedcurrency:                  v[2].(string),
		Grossprofit:                       v[3].(string),
		Totalrevenue:                      v[4].(string),
		Costofrevenue:                     v[5].(string),
		Costofgoodsandservicessold:        v[6].(string),
		Operatingincome:                   v[7].(string),
		Sellinggeneralandadministrative:   v[8].(string),
		Researchanddevelopment:            v[9].(string),
		Operatingexpenses:                 v[10].(string),
		Investmentincomenet:               v[11].(string),
		Netinterestincome:                 v[12].(string),
		Interestincome:                    v[13].(string),
		Interestexpense:                   v[14].(string),
		Noninterestincome:                 v[15].(string),
		Othernonoperatingincome:           v[16].(string),
		Depreciation:                      v[17].(string),
		Depreciationandamortization:       v[18].(string),
		Incomebeforetax:                   v[19].(string),
		Incometaxexpense:                  v[20].(string),
		Interestanddebtexpense:            v[21].(string),
		Netincomefromcontinuingoperations: v[22].(string),
		Comprehensiveincomenetoftax:       v[23].(string),
		Ebit:                              v[24].(string),
		Ebitda:                            v[25].(string),
		Netincome:                         v[26].(string),
	}
}
