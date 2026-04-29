package rtos

import (
	"fmt"
	"tspo/dtos"
	pb "tspo/proto/gen"
)

// Wrapper
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––
type CompanyRequestDB struct {
	Table string
	*pb.CompanyRequest
}

// tableNameable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––
func (req *CompanyRequestDB) TableName() string {
	return req.Table
}

// Requestable implemented
// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

func (req *CompanyRequestDB) Params() (string, error) {
	symbol := req.GetSymbol()
	if symbol == "" {
		return "", fmt.Errorf("Не передан параметр \"symbol\"")
	}

	return fmt.Sprintf("Where symbol='%s'", symbol), nil
}

func (req *CompanyRequestDB) GetDatabaseable() *dtos.OverviewDB {
	return &dtos.OverviewDB{Overview: &dtos.Overview{}}
}
