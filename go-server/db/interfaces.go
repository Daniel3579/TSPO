package db

type tableNameable interface {
	TableName() string
}

type Databaseable interface {
	tableNameable
	Columns() []string
	InsertableValues() []any
	SelectableValues() []any
}

type Sliceable interface {
	ToDatabaseableSlice() []Databaseable
}

type Requestable[T Databaseable] interface {
	tableNameable
	Params() (string, error)
	GetDatabaseable() T
}

type Grpcable interface {
	ToResponse() any
}
