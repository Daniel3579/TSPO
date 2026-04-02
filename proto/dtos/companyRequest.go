package dtos

import (
	"tspo/db"
	pb "tspo/proto/gen"
)

type CompanyRequest struct {
	*pb.CompanyRequest
}

func (s *CompanyRequest) ToDatabaseableSlice() []db.Databaseable {
	return []db.Databaseable{s}
}

func (req *CompanyRequest) TableName() string {
	return "overview"
}

func (req *CompanyRequest) Columns() []string {
	return []string{"*"}
}

func (req *CompanyRequest) InsertableValues() []any {
	return []any{req.Name}
}

func (req *CompanyRequest) SelectableValues() []any {
	return nil
}

type OverviewResponse struct {
	*pb.OverviewResponse
}

func (s *OverviewResponse) ToDatabaseableSlice() []db.Databaseable {
	return []db.Databaseable{s}
}

func (res *OverviewResponse) TableName() string {
	return ""
}

func (res *OverviewResponse) Columns() []string {
	return nil
}

func (res *OverviewResponse) InsertableValues() []any {
	return nil
}

func (res *OverviewResponse) SelectableValues() []any {
	return []any{
		&res.Symbol,
		&res.Assettype,
		&res.Name,
		&res.Description,
		&res.Cik,
		&res.Exchange,
		&res.Currency,
		&res.Country,
		&res.Sector,
		&res.Industry,
		&res.Address,
		&res.Officialsite,
		&res.Fiscalyearend,
		&res.Latestquarter,
		&res.Marketcapitalization,
		&res.Ebitda,
		&res.Peratio,
		&res.Pegratio,
		&res.Bookvalue,
		&res.Dividendpershare,
		&res.Dividendyield,
		&res.Eps,
		&res.Revenuepersharettm,
		&res.Profitmargin,
		&res.Operatingmarginttm,
		&res.Returnonassetsttm,
		&res.Returnonequityttm,
		&res.Revenuettm,
		&res.Grossprofitttm,
		&res.Dilutedepsttm,
		&res.Quarterlyearningsgrowthyoy,
		&res.Quarterlyrevenuegrowthyoy,
		&res.Analysttargetprice,
		&res.Analystratingstrongbuy,
		&res.Analystratingbuy,
		&res.Analystratinghold,
		&res.Analystratingsell,
		&res.Analystratingstrongsell,
		&res.Trailingpe,
		&res.Forwardpe,
		&res.Pricetosalesratiottm,
		&res.Pricetobookratio,
		&res.Evtorevenue,
		&res.Evtoebitda,
		&res.Beta,
		&res.Weekhigh52,
		&res.Weeklow52,
		&res.Daymovingaverage50,
		&res.Daymovingaverage200,
		&res.Sharesoutstanding,
		&res.Sharesfloat,
		&res.Percentinsiders,
		&res.Percentinstitutions,
		&res.Dividenddate,
		&res.Exdividenddate,
	}
}
