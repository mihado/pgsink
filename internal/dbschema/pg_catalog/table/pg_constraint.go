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

var PgConstraint = newPgConstraintTable()

type pgConstraintTable struct {
	postgres.Table

	//Columns
	Oid           postgres.ColumnString
	Conname       postgres.ColumnString
	Connamespace  postgres.ColumnString
	Contype       postgres.ColumnString
	Condeferrable postgres.ColumnBool
	Condeferred   postgres.ColumnBool
	Convalidated  postgres.ColumnBool
	Conrelid      postgres.ColumnString
	Contypid      postgres.ColumnString
	Conindid      postgres.ColumnString
	Conparentid   postgres.ColumnString
	Confrelid     postgres.ColumnString
	Confupdtype   postgres.ColumnString
	Confdeltype   postgres.ColumnString
	Confmatchtype postgres.ColumnString
	Conislocal    postgres.ColumnBool
	Coninhcount   postgres.ColumnInteger
	Connoinherit  postgres.ColumnBool
	Conkey        postgres.ColumnString
	Confkey       postgres.ColumnString
	Conpfeqop     postgres.ColumnString
	Conppeqop     postgres.ColumnString
	Conffeqop     postgres.ColumnString
	Conexclop     postgres.ColumnString
	Conbin        postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgConstraintTable struct {
	pgConstraintTable

	EXCLUDED pgConstraintTable
}

// AS creates new PgConstraintTable with assigned alias
func (a *PgConstraintTable) AS(alias string) *PgConstraintTable {
	aliasTable := newPgConstraintTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgConstraintTable() *PgConstraintTable {
	return &PgConstraintTable{
		pgConstraintTable: newPgConstraintTableImpl("pg_catalog", "pg_constraint"),
		EXCLUDED:          newPgConstraintTableImpl("", "excluded"),
	}
}

func newPgConstraintTableImpl(schemaName, tableName string) pgConstraintTable {
	var (
		OidColumn           = postgres.StringColumn("oid")
		ConnameColumn       = postgres.StringColumn("conname")
		ConnamespaceColumn  = postgres.StringColumn("connamespace")
		ContypeColumn       = postgres.StringColumn("contype")
		CondeferrableColumn = postgres.BoolColumn("condeferrable")
		CondeferredColumn   = postgres.BoolColumn("condeferred")
		ConvalidatedColumn  = postgres.BoolColumn("convalidated")
		ConrelidColumn      = postgres.StringColumn("conrelid")
		ContypidColumn      = postgres.StringColumn("contypid")
		ConindidColumn      = postgres.StringColumn("conindid")
		ConparentidColumn   = postgres.StringColumn("conparentid")
		ConfrelidColumn     = postgres.StringColumn("confrelid")
		ConfupdtypeColumn   = postgres.StringColumn("confupdtype")
		ConfdeltypeColumn   = postgres.StringColumn("confdeltype")
		ConfmatchtypeColumn = postgres.StringColumn("confmatchtype")
		ConislocalColumn    = postgres.BoolColumn("conislocal")
		ConinhcountColumn   = postgres.IntegerColumn("coninhcount")
		ConnoinheritColumn  = postgres.BoolColumn("connoinherit")
		ConkeyColumn        = postgres.StringColumn("conkey")
		ConfkeyColumn       = postgres.StringColumn("confkey")
		ConpfeqopColumn     = postgres.StringColumn("conpfeqop")
		ConppeqopColumn     = postgres.StringColumn("conppeqop")
		ConffeqopColumn     = postgres.StringColumn("conffeqop")
		ConexclopColumn     = postgres.StringColumn("conexclop")
		ConbinColumn        = postgres.StringColumn("conbin")
		allColumns          = postgres.ColumnList{OidColumn, ConnameColumn, ConnamespaceColumn, ContypeColumn, CondeferrableColumn, CondeferredColumn, ConvalidatedColumn, ConrelidColumn, ContypidColumn, ConindidColumn, ConparentidColumn, ConfrelidColumn, ConfupdtypeColumn, ConfdeltypeColumn, ConfmatchtypeColumn, ConislocalColumn, ConinhcountColumn, ConnoinheritColumn, ConkeyColumn, ConfkeyColumn, ConpfeqopColumn, ConppeqopColumn, ConffeqopColumn, ConexclopColumn, ConbinColumn}
		mutableColumns      = postgres.ColumnList{OidColumn, ConnameColumn, ConnamespaceColumn, ContypeColumn, CondeferrableColumn, CondeferredColumn, ConvalidatedColumn, ConrelidColumn, ContypidColumn, ConindidColumn, ConparentidColumn, ConfrelidColumn, ConfupdtypeColumn, ConfdeltypeColumn, ConfmatchtypeColumn, ConislocalColumn, ConinhcountColumn, ConnoinheritColumn, ConkeyColumn, ConfkeyColumn, ConpfeqopColumn, ConppeqopColumn, ConffeqopColumn, ConexclopColumn, ConbinColumn}
	)

	return pgConstraintTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Oid:           OidColumn,
		Conname:       ConnameColumn,
		Connamespace:  ConnamespaceColumn,
		Contype:       ContypeColumn,
		Condeferrable: CondeferrableColumn,
		Condeferred:   CondeferredColumn,
		Convalidated:  ConvalidatedColumn,
		Conrelid:      ConrelidColumn,
		Contypid:      ContypidColumn,
		Conindid:      ConindidColumn,
		Conparentid:   ConparentidColumn,
		Confrelid:     ConfrelidColumn,
		Confupdtype:   ConfupdtypeColumn,
		Confdeltype:   ConfdeltypeColumn,
		Confmatchtype: ConfmatchtypeColumn,
		Conislocal:    ConislocalColumn,
		Coninhcount:   ConinhcountColumn,
		Connoinherit:  ConnoinheritColumn,
		Conkey:        ConkeyColumn,
		Confkey:       ConfkeyColumn,
		Conpfeqop:     ConpfeqopColumn,
		Conppeqop:     ConppeqopColumn,
		Conffeqop:     ConffeqopColumn,
		Conexclop:     ConexclopColumn,
		Conbin:        ConbinColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
