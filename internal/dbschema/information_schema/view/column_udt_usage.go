//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/postgres"
)

var ColumnUdtUsage = newColumnUdtUsageTable()

type columnUdtUsageTable struct {
	postgres.Table

	//Columns
	UdtCatalog   postgres.ColumnString
	UdtSchema    postgres.ColumnString
	UdtName      postgres.ColumnString
	TableCatalog postgres.ColumnString
	TableSchema  postgres.ColumnString
	TableName    postgres.ColumnString
	ColumnName   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ColumnUdtUsageTable struct {
	columnUdtUsageTable

	EXCLUDED columnUdtUsageTable
}

// AS creates new ColumnUdtUsageTable with assigned alias
func (a *ColumnUdtUsageTable) AS(alias string) *ColumnUdtUsageTable {
	aliasTable := newColumnUdtUsageTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newColumnUdtUsageTable() *ColumnUdtUsageTable {
	return &ColumnUdtUsageTable{
		columnUdtUsageTable: newColumnUdtUsageTableImpl("information_schema", "column_udt_usage"),
		EXCLUDED:            newColumnUdtUsageTableImpl("", "excluded"),
	}
}

func newColumnUdtUsageTableImpl(schemaName, tableName string) columnUdtUsageTable {
	var (
		UdtCatalogColumn   = postgres.StringColumn("udt_catalog")
		UdtSchemaColumn    = postgres.StringColumn("udt_schema")
		UdtNameColumn      = postgres.StringColumn("udt_name")
		TableCatalogColumn = postgres.StringColumn("table_catalog")
		TableSchemaColumn  = postgres.StringColumn("table_schema")
		TableNameColumn    = postgres.StringColumn("table_name")
		ColumnNameColumn   = postgres.StringColumn("column_name")
		allColumns         = postgres.ColumnList{UdtCatalogColumn, UdtSchemaColumn, UdtNameColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn}
		mutableColumns     = postgres.ColumnList{UdtCatalogColumn, UdtSchemaColumn, UdtNameColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn}
	)

	return columnUdtUsageTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		UdtCatalog:   UdtCatalogColumn,
		UdtSchema:    UdtSchemaColumn,
		UdtName:      UdtNameColumn,
		TableCatalog: TableCatalogColumn,
		TableSchema:  TableSchemaColumn,
		TableName:    TableNameColumn,
		ColumnName:   ColumnNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
