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

var TablePrivileges = newTablePrivilegesTable()

type tablePrivilegesTable struct {
	postgres.Table

	//Columns
	Grantor       postgres.ColumnString
	Grantee       postgres.ColumnString
	TableCatalog  postgres.ColumnString
	TableSchema   postgres.ColumnString
	TableName     postgres.ColumnString
	PrivilegeType postgres.ColumnString
	IsGrantable   postgres.ColumnString
	WithHierarchy postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TablePrivilegesTable struct {
	tablePrivilegesTable

	EXCLUDED tablePrivilegesTable
}

// AS creates new TablePrivilegesTable with assigned alias
func (a *TablePrivilegesTable) AS(alias string) *TablePrivilegesTable {
	aliasTable := newTablePrivilegesTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newTablePrivilegesTable() *TablePrivilegesTable {
	return &TablePrivilegesTable{
		tablePrivilegesTable: newTablePrivilegesTableImpl("information_schema", "table_privileges"),
		EXCLUDED:             newTablePrivilegesTableImpl("", "excluded"),
	}
}

func newTablePrivilegesTableImpl(schemaName, tableName string) tablePrivilegesTable {
	var (
		GrantorColumn       = postgres.StringColumn("grantor")
		GranteeColumn       = postgres.StringColumn("grantee")
		TableCatalogColumn  = postgres.StringColumn("table_catalog")
		TableSchemaColumn   = postgres.StringColumn("table_schema")
		TableNameColumn     = postgres.StringColumn("table_name")
		PrivilegeTypeColumn = postgres.StringColumn("privilege_type")
		IsGrantableColumn   = postgres.StringColumn("is_grantable")
		WithHierarchyColumn = postgres.StringColumn("with_hierarchy")
		allColumns          = postgres.ColumnList{GrantorColumn, GranteeColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, PrivilegeTypeColumn, IsGrantableColumn, WithHierarchyColumn}
		mutableColumns      = postgres.ColumnList{GrantorColumn, GranteeColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, PrivilegeTypeColumn, IsGrantableColumn, WithHierarchyColumn}
	)

	return tablePrivilegesTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Grantor:       GrantorColumn,
		Grantee:       GranteeColumn,
		TableCatalog:  TableCatalogColumn,
		TableSchema:   TableSchemaColumn,
		TableName:     TableNameColumn,
		PrivilegeType: PrivilegeTypeColumn,
		IsGrantable:   IsGrantableColumn,
		WithHierarchy: WithHierarchyColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
