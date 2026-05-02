package handlers

import (
	"context"
	"tspo/db"
	"tspo/dtos"
	pb "tspo/proto/gen"
	"tspo/proto/rtos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedDataServiceServer
}

func (s *Server) Overview(ctx context.Context, req *pb.CompanyRequest) (*pb.OverviewResponse, error) {
	company := req.GetSymbol()

	if company == "" {
		return nil, status.Error(codes.InvalidArgument, "company required")
	}

	companyRequestDB := &rtos.CompanyRequestDB{
		Table:          "overview",
		CompanyRequest: req,
	}

	sel, err := db.Select(companyRequestDB)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Select failed: %v", err)
	}

	if len(sel) == 0 {
		return nil, status.Errorf(codes.NotFound, "Not found")

	} else {
		res := sel[0].ToResponse().(*pb.OverviewResponse)
		return res, nil
	}
}

func (s *Server) IncomeStatement(ctx context.Context, req *pb.FinancialRequest) (*pb.IncomeStatementResponse, error) {
	company := req.GetSymbol()

	if company == "" {
		return nil, status.Error(codes.InvalidArgument, "company required")
	}

	companyRequestDB := &rtos.FinancialRequestDB[*dtos.IncomeStatementDB]{
		Table: "income_statement",
		NewRowDatabaseable: func() *dtos.IncomeStatementDB {
			return &dtos.IncomeStatementDB{
				IncomeReports: &dtos.IncomeReports{},
			}
		},
		FinancialRequest: req,
	}

	sel, err := db.Select(companyRequestDB)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Select failed: %v", err)
	}

	if len(sel) == 0 {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}

	mul_res := make([]*pb.SingleIncomeStatementResponse, len(sel))

	for i, v := range sel {
		res := v.ToResponse().(*pb.SingleIncomeStatementResponse)
		mul_res[i] = res
	}

	if len(sel) == 1 {
		return &pb.IncomeStatementResponse{
			Response: &pb.IncomeStatementResponse_Single{Single: mul_res[0]},
		}, nil

	} else {
		multiple := &pb.MultipleIncomeStatementResponse{
			Response: mul_res,
		}
		return &pb.IncomeStatementResponse{
			Response: &pb.IncomeStatementResponse_Multiple{Multiple: multiple},
		}, nil
	}
}

func (s *Server) BalanceSheet(ctx context.Context, req *pb.FinancialRequest) (*pb.BalanceSheetResponse, error) {
	company := req.GetSymbol()

	if company == "" {
		return nil, status.Error(codes.InvalidArgument, "company required")
	}

	companyRequestDB := &rtos.FinancialRequestDB[*dtos.BalanceSheetDB]{
		Table: "balance_sheet",
		NewRowDatabaseable: func() *dtos.BalanceSheetDB {
			return &dtos.BalanceSheetDB{
				BalanceReports: &dtos.BalanceReports{},
			}
		},
		FinancialRequest: req,
	}

	sel, err := db.Select(companyRequestDB)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Select failed: %v", err)
	}

	if len(sel) == 0 {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}

	mul_res := make([]*pb.SingleBalanceSheetResponse, len(sel))

	for i, v := range sel {
		res := v.ToResponse().(*pb.SingleBalanceSheetResponse)
		mul_res[i] = res
	}

	if len(sel) == 1 {
		return &pb.BalanceSheetResponse{
			Response: &pb.BalanceSheetResponse_Single{Single: mul_res[0]},
		}, nil

	} else {
		multiple := &pb.MultipleBalanceSheetResponse{
			Response: mul_res,
		}
		return &pb.BalanceSheetResponse{
			Response: &pb.BalanceSheetResponse_Multiple{Multiple: multiple},
		}, nil
	}
}

func (s *Server) CashFlow(ctx context.Context, req *pb.FinancialRequest) (*pb.CashFlowResponse, error) {
	company := req.GetSymbol()

	if company == "" {
		return nil, status.Error(codes.InvalidArgument, "company required")
	}

	companyRequestDB := &rtos.FinancialRequestDB[*dtos.CashFlowDB]{
		Table: "cash_flow",
		NewRowDatabaseable: func() *dtos.CashFlowDB {
			return &dtos.CashFlowDB{
				CashFlowReports: &dtos.CashFlowReports{},
			}
		},
		FinancialRequest: req,
	}

	sel, err := db.Select(companyRequestDB)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Select failed: %v", err)
	}

	if len(sel) == 0 {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}

	mul_res := make([]*pb.SingleCashFlowResponse, len(sel))

	for i, v := range sel {
		res := v.ToResponse().(*pb.SingleCashFlowResponse)
		mul_res[i] = res
	}

	if len(sel) == 1 {
		return &pb.CashFlowResponse{
			Response: &pb.CashFlowResponse_Single{Single: mul_res[0]},
		}, nil

	} else {
		multiple := &pb.MultipleCashFlowResponse{
			Response: mul_res,
		}
		return &pb.CashFlowResponse{
			Response: &pb.CashFlowResponse_Multiple{Multiple: multiple},
		}, nil
	}
}
