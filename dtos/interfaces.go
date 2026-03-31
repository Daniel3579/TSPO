package dtos

type Insertable interface {
	TableName() string
	Columns() []string
	Values() []interface{}
}

type Scannable interface {
	ScanValues() []interface{}
}
