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

var PgPolicies = newPgPoliciesTable()

type pgPoliciesTable struct {
	postgres.Table

	//Columns
	Schemaname postgres.ColumnString
	Tablename  postgres.ColumnString
	Policyname postgres.ColumnString
	Permissive postgres.ColumnString
	Roles      postgres.ColumnString
	Cmd        postgres.ColumnString
	Qual       postgres.ColumnString
	WithCheck  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgPoliciesTable struct {
	pgPoliciesTable

	EXCLUDED pgPoliciesTable
}

// AS creates new PgPoliciesTable with assigned alias
func (a *PgPoliciesTable) AS(alias string) *PgPoliciesTable {
	aliasTable := newPgPoliciesTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgPoliciesTable() *PgPoliciesTable {
	return &PgPoliciesTable{
		pgPoliciesTable: newPgPoliciesTableImpl("pg_catalog", "pg_policies"),
		EXCLUDED:        newPgPoliciesTableImpl("", "excluded"),
	}
}

func newPgPoliciesTableImpl(schemaName, tableName string) pgPoliciesTable {
	var (
		SchemanameColumn = postgres.StringColumn("schemaname")
		TablenameColumn  = postgres.StringColumn("tablename")
		PolicynameColumn = postgres.StringColumn("policyname")
		PermissiveColumn = postgres.StringColumn("permissive")
		RolesColumn      = postgres.StringColumn("roles")
		CmdColumn        = postgres.StringColumn("cmd")
		QualColumn       = postgres.StringColumn("qual")
		WithCheckColumn  = postgres.StringColumn("with_check")
		allColumns       = postgres.ColumnList{SchemanameColumn, TablenameColumn, PolicynameColumn, PermissiveColumn, RolesColumn, CmdColumn, QualColumn, WithCheckColumn}
		mutableColumns   = postgres.ColumnList{SchemanameColumn, TablenameColumn, PolicynameColumn, PermissiveColumn, RolesColumn, CmdColumn, QualColumn, WithCheckColumn}
	)

	return pgPoliciesTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Schemaname: SchemanameColumn,
		Tablename:  TablenameColumn,
		Policyname: PolicynameColumn,
		Permissive: PermissiveColumn,
		Roles:      RolesColumn,
		Cmd:        CmdColumn,
		Qual:       QualColumn,
		WithCheck:  WithCheckColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
