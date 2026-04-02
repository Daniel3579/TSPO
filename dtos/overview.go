package dtos

import (
	"tspo/db"
)

type Overview struct {
	Symbol                     string `json:"Symbol"`
	AssetType                  string `json:"AssetType"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	CIK                        string `json:"CIK"`
	Exchange                   string `json:"Exchange"`
	Currency                   string `json:"Currency"`
	Country                    string `json:"Country"`
	Sector                     string `json:"Sector"`
	Industry                   string `json:"Industry"`
	Address                    string `json:"Address"`
	OfficialSite               string `json:"OfficialSite"`
	FiscalYearEnd              string `json:"FiscalYearEnd"`
	LatestQuarter              string `json:"LatestQuarter"`
	MarketCapitalization       string `json:"MarketCapitalization"`
	EBITDA                     string `json:"EBITDA"`
	PERatio                    string `json:"PERatio"`
	PEGRatio                   string `json:"PEGRatio"`
	BookValue                  string `json:"BookValue"`
	DividendPerShare           string `json:"DividendPerShare"`
	DividendYield              string `json:"DividendYield"`
	EPS                        string `json:"EPS"`
	RevenuePerShareTTM         string `json:"RevenuePerShareTTM"`
	ProfitMargin               string `json:"ProfitMargin"`
	OperatingMarginTTM         string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string `json:"ReturnOnEquityTTM"`
	RevenueTTM                 string `json:"RevenueTTM"`
	GrossProfitTTM             string `json:"GrossProfitTTM"`
	DilutedEPSTTM              string `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string `json:"AnalystTargetPrice"`
	AnalystRatingStrongBuy     string `json:"AnalystRatingStrongBuy"`
	AnalystRatingBuy           string `json:"AnalystRatingBuy"`
	AnalystRatingHold          string `json:"AnalystRatingHold"`
	AnalystRatingSell          string `json:"AnalystRatingSell"`
	AnalystRatingStrongSell    string `json:"AnalystRatingStrongSell"`
	TrailingPE                 string `json:"TrailingPE"`
	ForwardPE                  string `json:"ForwardPE"`
	PriceToSalesRatioTTM       string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           string `json:"PriceToBookRatio"`
	EVToRevenue                string `json:"EVToRevenue"`
	EVToEBITDA                 string `json:"EVToEBITDA"`
	Beta                       string `json:"Beta"`
	WeekHigh52                 string `json:"52WeekHigh"`
	WeekLow52                  string `json:"52WeekLow"`
	DayMovingAverage50         string `json:"50DayMovingAverage"`
	DayMovingAverage200        string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	SharesFloat                string `json:"SharesFloat"`
	PercentInsiders            string `json:"PercentInsiders"`
	PercentInstitutions        string `json:"PercentInstitutions"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
}

func (s *Overview) ToDatabaseableSlice() []db.Databaseable {
	return []db.Databaseable{&overviewDB{s}}
}

// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type overviewDB struct {
	*Overview
}

// ––––––––––––––––––––––––––––––––––––––––––––––

func (req *overviewDB) TableName() string {
	return "overview"
}

func (req *overviewDB) Columns() []string {
	return []string{
		"symbol",
		"assetType",
		"name",
		"description",
		"cIK",
		"exchange",
		"currency",
		"country",
		"sector",
		"industry",
		"address",
		"officialSite",
		"fiscalYearEnd",
		"latestQuarter",
		"marketCapitalization",
		"eBITDA",
		"pERatio",
		"pEGRatio",
		"bookValue",
		"dividendPerShare",
		"dividendYield",
		"ePS",
		"revenuePerShareTTM",
		"profitMargin",
		"operatingMarginTTM",
		"returnOnAssetsTTM",
		"returnOnEquityTTM",
		"revenueTTM",
		"grossProfitTTM",
		"dilutedEPSTTM",
		"quarterlyEarningsGrowthYOY",
		"quarterlyRevenueGrowthYOY",
		"analystTargetPrice",
		"analystRatingStrongBuy",
		"analystRatingBuy",
		"analystRatingHold",
		"analystRatingSell",
		"analystRatingStrongSell",
		"trailingPE",
		"forwardPE",
		"priceToSalesRatioTTM",
		"priceToBookRatio",
		"eVToRevenue",
		"eVToEBITDA",
		"beta",
		"weekHigh52",
		"weekLow52",
		"dayMovingAverage50",
		"dayMovingAverage200",
		"sharesOutstanding",
		"sharesFloat",
		"percentInsiders",
		"percentInstitutions",
		"dividendDate",
		"exDividendDate",
	}
}

// ––––––––––––––––––––––––––––––––––––––––––––––

func (req *overviewDB) InsertableValues() []any {
	return []any{
		req.Symbol,
		req.AssetType,
		req.Name,
		req.Description,
		req.CIK,
		req.Exchange,
		req.Currency,
		req.Country,
		req.Sector,
		req.Industry,
		req.Address,
		req.OfficialSite,
		req.FiscalYearEnd,
		req.LatestQuarter,
		req.MarketCapitalization,
		req.EBITDA,
		req.PERatio,
		req.PEGRatio,
		req.BookValue,
		req.DividendPerShare,
		req.DividendYield,
		req.EPS,
		req.RevenuePerShareTTM,
		req.ProfitMargin,
		req.OperatingMarginTTM,
		req.ReturnOnAssetsTTM,
		req.ReturnOnEquityTTM,
		req.RevenueTTM,
		req.GrossProfitTTM,
		req.DilutedEPSTTM,
		req.QuarterlyEarningsGrowthYOY,
		req.QuarterlyRevenueGrowthYOY,
		req.AnalystTargetPrice,
		req.AnalystRatingStrongBuy,
		req.AnalystRatingBuy,
		req.AnalystRatingHold,
		req.AnalystRatingSell,
		req.AnalystRatingStrongSell,
		req.TrailingPE,
		req.ForwardPE,
		req.PriceToSalesRatioTTM,
		req.PriceToBookRatio,
		req.EVToRevenue,
		req.EVToEBITDA,
		req.Beta,
		req.WeekHigh52,
		req.WeekLow52,
		req.DayMovingAverage50,
		req.DayMovingAverage200,
		req.SharesOutstanding,
		req.SharesFloat,
		req.PercentInsiders,
		req.PercentInstitutions,
		req.DividendDate,
		req.ExDividendDate,
	}
}

// ––––––––––––––––––––––––––––––––––––––––––––––

func (res *overviewDB) SelectableValues() []any {
	return []any{
		&res.Symbol,
		&res.AssetType,
		&res.Name,
		&res.Description,
		&res.CIK,
		&res.Exchange,
		&res.Currency,
		&res.Country,
		&res.Sector,
		&res.Industry,
		&res.Address,
		&res.OfficialSite,
		&res.FiscalYearEnd,
		&res.LatestQuarter,
		&res.MarketCapitalization,
		&res.EBITDA,
		&res.PERatio,
		&res.PEGRatio,
		&res.BookValue,
		&res.DividendPerShare,
		&res.DividendYield,
		&res.EPS,
		&res.RevenuePerShareTTM,
		&res.ProfitMargin,
		&res.OperatingMarginTTM,
		&res.ReturnOnAssetsTTM,
		&res.ReturnOnEquityTTM,
		&res.RevenueTTM,
		&res.GrossProfitTTM,
		&res.DilutedEPSTTM,
		&res.QuarterlyEarningsGrowthYOY,
		&res.QuarterlyRevenueGrowthYOY,
		&res.AnalystTargetPrice,
		&res.AnalystRatingStrongBuy,
		&res.AnalystRatingBuy,
		&res.AnalystRatingHold,
		&res.AnalystRatingSell,
		&res.AnalystRatingStrongSell,
		&res.TrailingPE,
		&res.ForwardPE,
		&res.PriceToSalesRatioTTM,
		&res.PriceToBookRatio,
		&res.EVToRevenue,
		&res.EVToEBITDA,
		&res.Beta,
		&res.WeekHigh52,
		&res.WeekLow52,
		&res.DayMovingAverage50,
		&res.DayMovingAverage200,
		&res.SharesOutstanding,
		&res.SharesFloat,
		&res.PercentInsiders,
		&res.PercentInstitutions,
		&res.DividendDate,
		&res.ExDividendDate,
	}
}
