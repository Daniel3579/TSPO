package db

type Sliceable interface {
	ToDatabaseableSlice() []Databaseable
}

type Databaseable interface {
	TableName() string
	Columns() []string
	insertable
	selectable
}

type insertable interface {
	InsertableValues() []any
}

type selectable interface {
	SelectableValues() []any
}
