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

var PgGroup = newPgGroupTable()

type pgGroupTable struct {
	postgres.Table

	//Columns
	Groname  postgres.ColumnString
	Grosysid postgres.ColumnString
	Grolist  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgGroupTable struct {
	pgGroupTable

	EXCLUDED pgGroupTable
}

// AS creates new PgGroupTable with assigned alias
func (a *PgGroupTable) AS(alias string) *PgGroupTable {
	aliasTable := newPgGroupTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgGroupTable() *PgGroupTable {
	return &PgGroupTable{
		pgGroupTable: newPgGroupTableImpl("pg_catalog", "pg_group"),
		EXCLUDED:     newPgGroupTableImpl("", "excluded"),
	}
}

func newPgGroupTableImpl(schemaName, tableName string) pgGroupTable {
	var (
		GronameColumn  = postgres.StringColumn("groname")
		GrosysidColumn = postgres.StringColumn("grosysid")
		GrolistColumn  = postgres.StringColumn("grolist")
		allColumns     = postgres.ColumnList{GronameColumn, GrosysidColumn, GrolistColumn}
		mutableColumns = postgres.ColumnList{GronameColumn, GrosysidColumn, GrolistColumn}
	)

	return pgGroupTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Groname:  GronameColumn,
		Grosysid: GrosysidColumn,
		Grolist:  GrolistColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
