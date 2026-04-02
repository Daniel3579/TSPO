package handlers

import (
	"context"
	pb "tspo/proto/gen"
)

type Server struct {
	pb.UnimplementedDataServiceServer
}

func (s *Server) Overview(ctx context.Context, req *pb.CompanyRequest) (*pb.OverviewResponse, error) {
	return nil, nil
}

func (s *Server) IncomeStatement(ctx context.Context, req *pb.FinancialRequest) (*pb.IncomeStatementResponse, error) {
	return nil, nil
}

func (s *Server) BalanceSheet(ctx context.Context, req *pb.FinancialRequest) (*pb.BalanceSheetResponse, error) {
	return nil, nil
}

func (s *Server) CashFlow(ctx context.Context, req *pb.FinancialRequest) (*pb.CashFlowResponse, error) {
	return nil, nil
}
