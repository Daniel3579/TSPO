package dtos

import (
	"tspo/db"
	pb "tspo/proto/gen"
)

// Original structure
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

// Wrapper
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

type OverviewDB struct {
	*Overview
}

// tableNameable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––

func (s *OverviewDB) TableName() string {
	return "overview"
}

// Databaseable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––

func (s *OverviewDB) Columns() []string {
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

func (s *OverviewDB) InsertableValues() []any {
	return []any{
		s.Symbol,
		s.AssetType,
		s.Name,
		s.Description,
		s.CIK,
		s.Exchange,
		s.Currency,
		s.Country,
		s.Sector,
		s.Industry,
		s.Address,
		s.OfficialSite,
		s.FiscalYearEnd,
		s.LatestQuarter,
		s.MarketCapitalization,
		s.EBITDA,
		s.PERatio,
		s.PEGRatio,
		s.BookValue,
		s.DividendPerShare,
		s.DividendYield,
		s.EPS,
		s.RevenuePerShareTTM,
		s.ProfitMargin,
		s.OperatingMarginTTM,
		s.ReturnOnAssetsTTM,
		s.ReturnOnEquityTTM,
		s.RevenueTTM,
		s.GrossProfitTTM,
		s.DilutedEPSTTM,
		s.QuarterlyEarningsGrowthYOY,
		s.QuarterlyRevenueGrowthYOY,
		s.AnalystTargetPrice,
		s.AnalystRatingStrongBuy,
		s.AnalystRatingBuy,
		s.AnalystRatingHold,
		s.AnalystRatingSell,
		s.AnalystRatingStrongSell,
		s.TrailingPE,
		s.ForwardPE,
		s.PriceToSalesRatioTTM,
		s.PriceToBookRatio,
		s.EVToRevenue,
		s.EVToEBITDA,
		s.Beta,
		s.WeekHigh52,
		s.WeekLow52,
		s.DayMovingAverage50,
		s.DayMovingAverage200,
		s.SharesOutstanding,
		s.SharesFloat,
		s.PercentInsiders,
		s.PercentInstitutions,
		s.DividendDate,
		s.ExDividendDate,
	}
}

func (s *OverviewDB) SelectableValues() []any {
	return []any{
		&s.Symbol,
		&s.AssetType,
		&s.Name,
		&s.Description,
		&s.CIK,
		&s.Exchange,
		&s.Currency,
		&s.Country,
		&s.Sector,
		&s.Industry,
		&s.Address,
		&s.OfficialSite,
		&s.FiscalYearEnd,
		&s.LatestQuarter,
		&s.MarketCapitalization,
		&s.EBITDA,
		&s.PERatio,
		&s.PEGRatio,
		&s.BookValue,
		&s.DividendPerShare,
		&s.DividendYield,
		&s.EPS,
		&s.RevenuePerShareTTM,
		&s.ProfitMargin,
		&s.OperatingMarginTTM,
		&s.ReturnOnAssetsTTM,
		&s.ReturnOnEquityTTM,
		&s.RevenueTTM,
		&s.GrossProfitTTM,
		&s.DilutedEPSTTM,
		&s.QuarterlyEarningsGrowthYOY,
		&s.QuarterlyRevenueGrowthYOY,
		&s.AnalystTargetPrice,
		&s.AnalystRatingStrongBuy,
		&s.AnalystRatingBuy,
		&s.AnalystRatingHold,
		&s.AnalystRatingSell,
		&s.AnalystRatingStrongSell,
		&s.TrailingPE,
		&s.ForwardPE,
		&s.PriceToSalesRatioTTM,
		&s.PriceToBookRatio,
		&s.EVToRevenue,
		&s.EVToEBITDA,
		&s.Beta,
		&s.WeekHigh52,
		&s.WeekLow52,
		&s.DayMovingAverage50,
		&s.DayMovingAverage200,
		&s.SharesOutstanding,
		&s.SharesFloat,
		&s.PercentInsiders,
		&s.PercentInstitutions,
		&s.DividendDate,
		&s.ExDividendDate,
	}
}

// Sliceable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *Overview) ToDatabaseableSlice() []db.Databaseable {
	return []db.Databaseable{&OverviewDB{s}}
}

// Grpcable implemented
// –––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (s *OverviewDB) ToResponse() any {
	v := s.InsertableValues()
	return &pb.OverviewResponse{
		Symbol:                     v[0].(string),
		Assettype:                  v[1].(string),
		Name:                       v[2].(string),
		Description:                v[3].(string),
		Cik:                        v[4].(string),
		Exchange:                   v[5].(string),
		Currency:                   v[6].(string),
		Country:                    v[7].(string),
		Sector:                     v[8].(string),
		Industry:                   v[9].(string),
		Address:                    v[10].(string),
		Officialsite:               v[11].(string),
		Fiscalyearend:              v[12].(string),
		Latestquarter:              v[13].(string),
		Marketcapitalization:       v[14].(string),
		Ebitda:                     v[15].(string),
		Peratio:                    v[16].(string),
		Pegratio:                   v[17].(string),
		Bookvalue:                  v[18].(string),
		Dividendpershare:           v[19].(string),
		Dividendyield:              v[20].(string),
		Eps:                        v[21].(string),
		Revenuepersharettm:         v[22].(string),
		Profitmargin:               v[23].(string),
		Operatingmarginttm:         v[24].(string),
		Returnonassetsttm:          v[25].(string),
		Returnonequityttm:          v[26].(string),
		Revenuettm:                 v[27].(string),
		Grossprofitttm:             v[28].(string),
		Dilutedepsttm:              v[29].(string),
		Quarterlyearningsgrowthyoy: v[30].(string),
		Quarterlyrevenuegrowthyoy:  v[31].(string),
		Analysttargetprice:         v[32].(string),
		Analystratingstrongbuy:     v[33].(string),
		Analystratingbuy:           v[34].(string),
		Analystratinghold:          v[35].(string),
		Analystratingsell:          v[36].(string),
		Analystratingstrongsell:    v[37].(string),
		Trailingpe:                 v[38].(string),
		Forwardpe:                  v[39].(string),
		Pricetosalesratiottm:       v[40].(string),
		Pricetobookratio:           v[41].(string),
		Evtorevenue:                v[42].(string),
		Evtoebitda:                 v[43].(string),
		Beta:                       v[44].(string),
		Weekhigh52:                 v[45].(string),
		Weeklow52:                  v[46].(string),
		Daymovingaverage50:         v[47].(string),
		Daymovingaverage200:        v[48].(string),
		Sharesoutstanding:          v[49].(string),
		Sharesfloat:                v[50].(string),
		Percentinsiders:            v[51].(string),
		Percentinstitutions:        v[52].(string),
		Dividenddate:               v[53].(string),
		Exdividenddate:             v[54].(string),
	}
}
