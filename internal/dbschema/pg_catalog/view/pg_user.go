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

var PgUser = newPgUserTable()

type pgUserTable struct {
	postgres.Table

	//Columns
	Usename      postgres.ColumnString
	Usesysid     postgres.ColumnString
	Usecreatedb  postgres.ColumnBool
	Usesuper     postgres.ColumnBool
	Userepl      postgres.ColumnBool
	Usebypassrls postgres.ColumnBool
	Passwd       postgres.ColumnString
	Valuntil     postgres.ColumnTimestampz
	Useconfig    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgUserTable struct {
	pgUserTable

	EXCLUDED pgUserTable
}

// AS creates new PgUserTable with assigned alias
func (a *PgUserTable) AS(alias string) *PgUserTable {
	aliasTable := newPgUserTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgUserTable() *PgUserTable {
	return &PgUserTable{
		pgUserTable: newPgUserTableImpl("pg_catalog", "pg_user"),
		EXCLUDED:    newPgUserTableImpl("", "excluded"),
	}
}

func newPgUserTableImpl(schemaName, tableName string) pgUserTable {
	var (
		UsenameColumn      = postgres.StringColumn("usename")
		UsesysidColumn     = postgres.StringColumn("usesysid")
		UsecreatedbColumn  = postgres.BoolColumn("usecreatedb")
		UsesuperColumn     = postgres.BoolColumn("usesuper")
		UsereplColumn      = postgres.BoolColumn("userepl")
		UsebypassrlsColumn = postgres.BoolColumn("usebypassrls")
		PasswdColumn       = postgres.StringColumn("passwd")
		ValuntilColumn     = postgres.TimestampzColumn("valuntil")
		UseconfigColumn    = postgres.StringColumn("useconfig")
		allColumns         = postgres.ColumnList{UsenameColumn, UsesysidColumn, UsecreatedbColumn, UsesuperColumn, UsereplColumn, UsebypassrlsColumn, PasswdColumn, ValuntilColumn, UseconfigColumn}
		mutableColumns     = postgres.ColumnList{UsenameColumn, UsesysidColumn, UsecreatedbColumn, UsesuperColumn, UsereplColumn, UsebypassrlsColumn, PasswdColumn, ValuntilColumn, UseconfigColumn}
	)

	return pgUserTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Usename:      UsenameColumn,
		Usesysid:     UsesysidColumn,
		Usecreatedb:  UsecreatedbColumn,
		Usesuper:     UsesuperColumn,
		Userepl:      UsereplColumn,
		Usebypassrls: UsebypassrlsColumn,
		Passwd:       PasswdColumn,
		Valuntil:     ValuntilColumn,
		Useconfig:    UseconfigColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
