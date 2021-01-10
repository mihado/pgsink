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

var PgAuthMembers = newPgAuthMembersTable()

type pgAuthMembersTable struct {
	postgres.Table

	//Columns
	Roleid      postgres.ColumnString
	Member      postgres.ColumnString
	Grantor     postgres.ColumnString
	AdminOption postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgAuthMembersTable struct {
	pgAuthMembersTable

	EXCLUDED pgAuthMembersTable
}

// AS creates new PgAuthMembersTable with assigned alias
func (a *PgAuthMembersTable) AS(alias string) *PgAuthMembersTable {
	aliasTable := newPgAuthMembersTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgAuthMembersTable() *PgAuthMembersTable {
	return &PgAuthMembersTable{
		pgAuthMembersTable: newPgAuthMembersTableImpl("pg_catalog", "pg_auth_members"),
		EXCLUDED:           newPgAuthMembersTableImpl("", "excluded"),
	}
}

func newPgAuthMembersTableImpl(schemaName, tableName string) pgAuthMembersTable {
	var (
		RoleidColumn      = postgres.StringColumn("roleid")
		MemberColumn      = postgres.StringColumn("member")
		GrantorColumn     = postgres.StringColumn("grantor")
		AdminOptionColumn = postgres.BoolColumn("admin_option")
		allColumns        = postgres.ColumnList{RoleidColumn, MemberColumn, GrantorColumn, AdminOptionColumn}
		mutableColumns    = postgres.ColumnList{RoleidColumn, MemberColumn, GrantorColumn, AdminOptionColumn}
	)

	return pgAuthMembersTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Roleid:      RoleidColumn,
		Member:      MemberColumn,
		Grantor:     GrantorColumn,
		AdminOption: AdminOptionColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
