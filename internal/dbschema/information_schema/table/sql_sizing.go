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

var SQLSizing = newSQLSizingTable()

type sQLSizingTable struct {
	postgres.Table

	//Columns
	SizingID       postgres.ColumnInteger
	SizingName     postgres.ColumnString
	SupportedValue postgres.ColumnInteger
	Comments       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SQLSizingTable struct {
	sQLSizingTable

	EXCLUDED sQLSizingTable
}

// AS creates new SQLSizingTable with assigned alias
func (a *SQLSizingTable) AS(alias string) *SQLSizingTable {
	aliasTable := newSQLSizingTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newSQLSizingTable() *SQLSizingTable {
	return &SQLSizingTable{
		sQLSizingTable: newSQLSizingTableImpl("information_schema", "sql_sizing"),
		EXCLUDED:       newSQLSizingTableImpl("", "excluded"),
	}
}

func newSQLSizingTableImpl(schemaName, tableName string) sQLSizingTable {
	var (
		SizingIDColumn       = postgres.IntegerColumn("sizing_id")
		SizingNameColumn     = postgres.StringColumn("sizing_name")
		SupportedValueColumn = postgres.IntegerColumn("supported_value")
		CommentsColumn       = postgres.StringColumn("comments")
		allColumns           = postgres.ColumnList{SizingIDColumn, SizingNameColumn, SupportedValueColumn, CommentsColumn}
		mutableColumns       = postgres.ColumnList{SizingIDColumn, SizingNameColumn, SupportedValueColumn, CommentsColumn}
	)

	return sQLSizingTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		SizingID:       SizingIDColumn,
		SizingName:     SizingNameColumn,
		SupportedValue: SupportedValueColumn,
		Comments:       CommentsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
