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

var PgCast = newPgCastTable()

type pgCastTable struct {
	postgres.Table

	//Columns
	Oid         postgres.ColumnString
	Castsource  postgres.ColumnString
	Casttarget  postgres.ColumnString
	Castfunc    postgres.ColumnString
	Castcontext postgres.ColumnString
	Castmethod  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgCastTable struct {
	pgCastTable

	EXCLUDED pgCastTable
}

// AS creates new PgCastTable with assigned alias
func (a *PgCastTable) AS(alias string) *PgCastTable {
	aliasTable := newPgCastTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgCastTable() *PgCastTable {
	return &PgCastTable{
		pgCastTable: newPgCastTableImpl("pg_catalog", "pg_cast"),
		EXCLUDED:    newPgCastTableImpl("", "excluded"),
	}
}

func newPgCastTableImpl(schemaName, tableName string) pgCastTable {
	var (
		OidColumn         = postgres.StringColumn("oid")
		CastsourceColumn  = postgres.StringColumn("castsource")
		CasttargetColumn  = postgres.StringColumn("casttarget")
		CastfuncColumn    = postgres.StringColumn("castfunc")
		CastcontextColumn = postgres.StringColumn("castcontext")
		CastmethodColumn  = postgres.StringColumn("castmethod")
		allColumns        = postgres.ColumnList{OidColumn, CastsourceColumn, CasttargetColumn, CastfuncColumn, CastcontextColumn, CastmethodColumn}
		mutableColumns    = postgres.ColumnList{OidColumn, CastsourceColumn, CasttargetColumn, CastfuncColumn, CastcontextColumn, CastmethodColumn}
	)

	return pgCastTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Oid:         OidColumn,
		Castsource:  CastsourceColumn,
		Casttarget:  CasttargetColumn,
		Castfunc:    CastfuncColumn,
		Castcontext: CastcontextColumn,
		Castmethod:  CastmethodColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
