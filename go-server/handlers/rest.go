package handlers

import (
	"encoding/json"
	"net/http"
	pb "tspo/proto/gen"
)

type CompanyRequest struct {
	Symbol string `json:"symbol"`
}

type Quarter struct {
	Year    int32 `json:"year"`
	Quarter int32 `json:"quarter"`
}

type QuarterRange struct {
	Start Quarter `json:"start"`
	End   Quarter `json:"end"`
}

type FinancialRequest struct {
	Symbol string        `json:"symbol"`
	Date   *Quarter      `json:"date,omitempty"`
	Range  *QuarterRange `json:"range,omitempty"`
}

type HttpServer struct {
	GrpcSrv *Server
}

func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func (h *HttpServer) OverviewHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody CompanyRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	grpcReq := &pb.CompanyRequest{Symbol: reqBody.Symbol}
	resp, err := h.GrpcSrv.Overview(r.Context(), grpcReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *HttpServer) IncomeHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody FinancialRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if (reqBody.Date == nil && reqBody.Range == nil) || (reqBody.Date != nil && reqBody.Range != nil) {
		http.Error(w, "must specify either 'date' or 'range'", http.StatusBadRequest)
		return
	}

	var grpcReq *pb.FinancialRequest = &pb.FinancialRequest{Symbol: reqBody.Symbol}

	if reqBody.Date != nil {
		grpcReq.Timeframe = &pb.FinancialRequest_Date{Date: &pb.Quarter{Year: reqBody.Date.Year, Quarter: reqBody.Date.Quarter}}
	}
	if reqBody.Range != nil {
		grpcReq.Timeframe = &pb.FinancialRequest_Range{Range: &pb.QuarterRange{
			Start: &pb.Quarter{Year: reqBody.Range.Start.Year, Quarter: reqBody.Range.Start.Quarter},
			End:   &pb.Quarter{Year: reqBody.Range.End.Year, Quarter: reqBody.Range.End.Quarter},
		}}
	}

	resp, err := h.GrpcSrv.IncomeStatement(r.Context(), grpcReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Преобразование oneof в удобный JSON
	var result interface{}
	switch v := resp.Response.(type) {
	case *pb.IncomeStatementResponse_Single:
		result = v.Single
	case *pb.IncomeStatementResponse_Multiple:
		result = v.Multiple
	}

	json.NewEncoder(w).Encode(result)
}

func (h *HttpServer) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody FinancialRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if (reqBody.Date == nil && reqBody.Range == nil) || (reqBody.Date != nil && reqBody.Range != nil) {
		http.Error(w, "must specify either 'date' or 'range'", http.StatusBadRequest)
		return
	}

	var grpcReq *pb.FinancialRequest = &pb.FinancialRequest{Symbol: reqBody.Symbol}

	if reqBody.Date != nil {
		grpcReq.Timeframe = &pb.FinancialRequest_Date{Date: &pb.Quarter{Year: reqBody.Date.Year, Quarter: reqBody.Date.Quarter}}
	}
	if reqBody.Range != nil {
		grpcReq.Timeframe = &pb.FinancialRequest_Range{Range: &pb.QuarterRange{
			Start: &pb.Quarter{Year: reqBody.Range.Start.Year, Quarter: reqBody.Range.Start.Quarter},
			End:   &pb.Quarter{Year: reqBody.Range.End.Year, Quarter: reqBody.Range.End.Quarter},
		}}
	}

	resp, err := h.GrpcSrv.BalanceSheet(r.Context(), grpcReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Преобразование oneof в удобный JSON
	var result interface{}
	switch v := resp.Response.(type) {
	case *pb.BalanceSheetResponse_Single:
		result = v.Single
	case *pb.BalanceSheetResponse_Multiple:
		result = v.Multiple
	}

	json.NewEncoder(w).Encode(result)
}

func (h *HttpServer) CashFlowHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody FinancialRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if (reqBody.Date == nil && reqBody.Range == nil) || (reqBody.Date != nil && reqBody.Range != nil) {
		http.Error(w, "must specify either 'date' or 'range'", http.StatusBadRequest)
		return
	}

	var grpcReq *pb.FinancialRequest = &pb.FinancialRequest{Symbol: reqBody.Symbol}

	if reqBody.Date != nil {
		grpcReq.Timeframe = &pb.FinancialRequest_Date{Date: &pb.Quarter{Year: reqBody.Date.Year, Quarter: reqBody.Date.Quarter}}
	}
	if reqBody.Range != nil {
		grpcReq.Timeframe = &pb.FinancialRequest_Range{Range: &pb.QuarterRange{
			Start: &pb.Quarter{Year: reqBody.Range.Start.Year, Quarter: reqBody.Range.Start.Quarter},
			End:   &pb.Quarter{Year: reqBody.Range.End.Year, Quarter: reqBody.Range.End.Quarter},
		}}
	}

	resp, err := h.GrpcSrv.CashFlow(r.Context(), grpcReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Преобразование oneof в удобный JSON
	var result interface{}
	switch v := resp.Response.(type) {
	case *pb.CashFlowResponse_Single:
		result = v.Single
	case *pb.CashFlowResponse_Multiple:
		result = v.Multiple
	}

	json.NewEncoder(w).Encode(result)
}
