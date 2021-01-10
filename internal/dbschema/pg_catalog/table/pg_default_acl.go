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

var PgDefaultACL = newPgDefaultACLTable()

type pgDefaultACLTable struct {
	postgres.Table

	//Columns
	Oid             postgres.ColumnString
	Defaclrole      postgres.ColumnString
	Defaclnamespace postgres.ColumnString
	Defaclobjtype   postgres.ColumnString
	Defaclacl       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgDefaultACLTable struct {
	pgDefaultACLTable

	EXCLUDED pgDefaultACLTable
}

// AS creates new PgDefaultACLTable with assigned alias
func (a *PgDefaultACLTable) AS(alias string) *PgDefaultACLTable {
	aliasTable := newPgDefaultACLTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgDefaultACLTable() *PgDefaultACLTable {
	return &PgDefaultACLTable{
		pgDefaultACLTable: newPgDefaultACLTableImpl("pg_catalog", "pg_default_acl"),
		EXCLUDED:          newPgDefaultACLTableImpl("", "excluded"),
	}
}

func newPgDefaultACLTableImpl(schemaName, tableName string) pgDefaultACLTable {
	var (
		OidColumn             = postgres.StringColumn("oid")
		DefaclroleColumn      = postgres.StringColumn("defaclrole")
		DefaclnamespaceColumn = postgres.StringColumn("defaclnamespace")
		DefaclobjtypeColumn   = postgres.StringColumn("defaclobjtype")
		DefaclaclColumn       = postgres.StringColumn("defaclacl")
		allColumns            = postgres.ColumnList{OidColumn, DefaclroleColumn, DefaclnamespaceColumn, DefaclobjtypeColumn, DefaclaclColumn}
		mutableColumns        = postgres.ColumnList{OidColumn, DefaclroleColumn, DefaclnamespaceColumn, DefaclobjtypeColumn, DefaclaclColumn}
	)

	return pgDefaultACLTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Oid:             OidColumn,
		Defaclrole:      DefaclroleColumn,
		Defaclnamespace: DefaclnamespaceColumn,
		Defaclobjtype:   DefaclobjtypeColumn,
		Defaclacl:       DefaclaclColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
