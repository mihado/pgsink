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

var PgLargeobjectMetadata = newPgLargeobjectMetadataTable()

type pgLargeobjectMetadataTable struct {
	postgres.Table

	//Columns
	Oid      postgres.ColumnString
	Lomowner postgres.ColumnString
	Lomacl   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgLargeobjectMetadataTable struct {
	pgLargeobjectMetadataTable

	EXCLUDED pgLargeobjectMetadataTable
}

// AS creates new PgLargeobjectMetadataTable with assigned alias
func (a *PgLargeobjectMetadataTable) AS(alias string) *PgLargeobjectMetadataTable {
	aliasTable := newPgLargeobjectMetadataTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgLargeobjectMetadataTable() *PgLargeobjectMetadataTable {
	return &PgLargeobjectMetadataTable{
		pgLargeobjectMetadataTable: newPgLargeobjectMetadataTableImpl("pg_catalog", "pg_largeobject_metadata"),
		EXCLUDED:                   newPgLargeobjectMetadataTableImpl("", "excluded"),
	}
}

func newPgLargeobjectMetadataTableImpl(schemaName, tableName string) pgLargeobjectMetadataTable {
	var (
		OidColumn      = postgres.StringColumn("oid")
		LomownerColumn = postgres.StringColumn("lomowner")
		LomaclColumn   = postgres.StringColumn("lomacl")
		allColumns     = postgres.ColumnList{OidColumn, LomownerColumn, LomaclColumn}
		mutableColumns = postgres.ColumnList{OidColumn, LomownerColumn, LomaclColumn}
	)

	return pgLargeobjectMetadataTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Oid:      OidColumn,
		Lomowner: LomownerColumn,
		Lomacl:   LomaclColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
