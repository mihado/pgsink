//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var ImportJobs = newImportJobsTable()

type importJobsTable struct {
	postgres.Table

	//Columns
	ID             postgres.ColumnInteger
	SubscriptionID postgres.ColumnString
	TableName      postgres.ColumnString
	Cursor         postgres.ColumnString
	CompletedAt    postgres.ColumnTimestampz
	ExpiredAt      postgres.ColumnTimestampz
	UpdatedAt      postgres.ColumnTimestampz
	CreatedAt      postgres.ColumnTimestampz
	Error          postgres.ColumnString
	Schema         postgres.ColumnString
	ErrorCount     postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ImportJobsTable struct {
	importJobsTable

	EXCLUDED importJobsTable
}

// AS creates new ImportJobsTable with assigned alias
func (a *ImportJobsTable) AS(alias string) *ImportJobsTable {
	aliasTable := newImportJobsTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newImportJobsTable() *ImportJobsTable {
	return &ImportJobsTable{
		importJobsTable: newImportJobsTableImpl("pgsink", "import_jobs"),
		EXCLUDED:        newImportJobsTableImpl("", "excluded"),
	}
}

func newImportJobsTableImpl(schemaName, tableName string) importJobsTable {
	var (
		IDColumn             = postgres.IntegerColumn("id")
		SubscriptionIDColumn = postgres.StringColumn("subscription_id")
		TableNameColumn      = postgres.StringColumn("table_name")
		CursorColumn         = postgres.StringColumn("cursor")
		CompletedAtColumn    = postgres.TimestampzColumn("completed_at")
		ExpiredAtColumn      = postgres.TimestampzColumn("expired_at")
		UpdatedAtColumn      = postgres.TimestampzColumn("updated_at")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		ErrorColumn          = postgres.StringColumn("error")
		SchemaColumn         = postgres.StringColumn("schema")
		ErrorCountColumn     = postgres.IntegerColumn("error_count")
		allColumns           = postgres.ColumnList{IDColumn, SubscriptionIDColumn, TableNameColumn, CursorColumn, CompletedAtColumn, ExpiredAtColumn, UpdatedAtColumn, CreatedAtColumn, ErrorColumn, SchemaColumn, ErrorCountColumn}
		mutableColumns       = postgres.ColumnList{SubscriptionIDColumn, TableNameColumn, CursorColumn, CompletedAtColumn, ExpiredAtColumn, UpdatedAtColumn, CreatedAtColumn, ErrorColumn, SchemaColumn, ErrorCountColumn}
	)

	return importJobsTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		ID:             IDColumn,
		SubscriptionID: SubscriptionIDColumn,
		TableName:      TableNameColumn,
		Cursor:         CursorColumn,
		CompletedAt:    CompletedAtColumn,
		ExpiredAt:      ExpiredAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		CreatedAt:      CreatedAtColumn,
		Error:          ErrorColumn,
		Schema:         SchemaColumn,
		ErrorCount:     ErrorCountColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
