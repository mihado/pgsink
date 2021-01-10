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

var PgType = newPgTypeTable()

type pgTypeTable struct {
	postgres.Table

	//Columns
	Oid            postgres.ColumnString
	Typname        postgres.ColumnString
	Typnamespace   postgres.ColumnString
	Typowner       postgres.ColumnString
	Typlen         postgres.ColumnInteger
	Typbyval       postgres.ColumnBool
	Typtype        postgres.ColumnString
	Typcategory    postgres.ColumnString
	Typispreferred postgres.ColumnBool
	Typisdefined   postgres.ColumnBool
	Typdelim       postgres.ColumnString
	Typrelid       postgres.ColumnString
	Typelem        postgres.ColumnString
	Typarray       postgres.ColumnString
	Typinput       postgres.ColumnString
	Typoutput      postgres.ColumnString
	Typreceive     postgres.ColumnString
	Typsend        postgres.ColumnString
	Typmodin       postgres.ColumnString
	Typmodout      postgres.ColumnString
	Typanalyze     postgres.ColumnString
	Typalign       postgres.ColumnString
	Typstorage     postgres.ColumnString
	Typnotnull     postgres.ColumnBool
	Typbasetype    postgres.ColumnString
	Typtypmod      postgres.ColumnInteger
	Typndims       postgres.ColumnInteger
	Typcollation   postgres.ColumnString
	Typdefaultbin  postgres.ColumnString
	Typdefault     postgres.ColumnString
	Typacl         postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgTypeTable struct {
	pgTypeTable

	EXCLUDED pgTypeTable
}

// AS creates new PgTypeTable with assigned alias
func (a *PgTypeTable) AS(alias string) *PgTypeTable {
	aliasTable := newPgTypeTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgTypeTable() *PgTypeTable {
	return &PgTypeTable{
		pgTypeTable: newPgTypeTableImpl("pg_catalog", "pg_type"),
		EXCLUDED:    newPgTypeTableImpl("", "excluded"),
	}
}

func newPgTypeTableImpl(schemaName, tableName string) pgTypeTable {
	var (
		OidColumn            = postgres.StringColumn("oid")
		TypnameColumn        = postgres.StringColumn("typname")
		TypnamespaceColumn   = postgres.StringColumn("typnamespace")
		TypownerColumn       = postgres.StringColumn("typowner")
		TyplenColumn         = postgres.IntegerColumn("typlen")
		TypbyvalColumn       = postgres.BoolColumn("typbyval")
		TyptypeColumn        = postgres.StringColumn("typtype")
		TypcategoryColumn    = postgres.StringColumn("typcategory")
		TypispreferredColumn = postgres.BoolColumn("typispreferred")
		TypisdefinedColumn   = postgres.BoolColumn("typisdefined")
		TypdelimColumn       = postgres.StringColumn("typdelim")
		TyprelidColumn       = postgres.StringColumn("typrelid")
		TypelemColumn        = postgres.StringColumn("typelem")
		TyparrayColumn       = postgres.StringColumn("typarray")
		TypinputColumn       = postgres.StringColumn("typinput")
		TypoutputColumn      = postgres.StringColumn("typoutput")
		TypreceiveColumn     = postgres.StringColumn("typreceive")
		TypsendColumn        = postgres.StringColumn("typsend")
		TypmodinColumn       = postgres.StringColumn("typmodin")
		TypmodoutColumn      = postgres.StringColumn("typmodout")
		TypanalyzeColumn     = postgres.StringColumn("typanalyze")
		TypalignColumn       = postgres.StringColumn("typalign")
		TypstorageColumn     = postgres.StringColumn("typstorage")
		TypnotnullColumn     = postgres.BoolColumn("typnotnull")
		TypbasetypeColumn    = postgres.StringColumn("typbasetype")
		TyptypmodColumn      = postgres.IntegerColumn("typtypmod")
		TypndimsColumn       = postgres.IntegerColumn("typndims")
		TypcollationColumn   = postgres.StringColumn("typcollation")
		TypdefaultbinColumn  = postgres.StringColumn("typdefaultbin")
		TypdefaultColumn     = postgres.StringColumn("typdefault")
		TypaclColumn         = postgres.StringColumn("typacl")
		allColumns           = postgres.ColumnList{OidColumn, TypnameColumn, TypnamespaceColumn, TypownerColumn, TyplenColumn, TypbyvalColumn, TyptypeColumn, TypcategoryColumn, TypispreferredColumn, TypisdefinedColumn, TypdelimColumn, TyprelidColumn, TypelemColumn, TyparrayColumn, TypinputColumn, TypoutputColumn, TypreceiveColumn, TypsendColumn, TypmodinColumn, TypmodoutColumn, TypanalyzeColumn, TypalignColumn, TypstorageColumn, TypnotnullColumn, TypbasetypeColumn, TyptypmodColumn, TypndimsColumn, TypcollationColumn, TypdefaultbinColumn, TypdefaultColumn, TypaclColumn}
		mutableColumns       = postgres.ColumnList{OidColumn, TypnameColumn, TypnamespaceColumn, TypownerColumn, TyplenColumn, TypbyvalColumn, TyptypeColumn, TypcategoryColumn, TypispreferredColumn, TypisdefinedColumn, TypdelimColumn, TyprelidColumn, TypelemColumn, TyparrayColumn, TypinputColumn, TypoutputColumn, TypreceiveColumn, TypsendColumn, TypmodinColumn, TypmodoutColumn, TypanalyzeColumn, TypalignColumn, TypstorageColumn, TypnotnullColumn, TypbasetypeColumn, TyptypmodColumn, TypndimsColumn, TypcollationColumn, TypdefaultbinColumn, TypdefaultColumn, TypaclColumn}
	)

	return pgTypeTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Oid:            OidColumn,
		Typname:        TypnameColumn,
		Typnamespace:   TypnamespaceColumn,
		Typowner:       TypownerColumn,
		Typlen:         TyplenColumn,
		Typbyval:       TypbyvalColumn,
		Typtype:        TyptypeColumn,
		Typcategory:    TypcategoryColumn,
		Typispreferred: TypispreferredColumn,
		Typisdefined:   TypisdefinedColumn,
		Typdelim:       TypdelimColumn,
		Typrelid:       TyprelidColumn,
		Typelem:        TypelemColumn,
		Typarray:       TyparrayColumn,
		Typinput:       TypinputColumn,
		Typoutput:      TypoutputColumn,
		Typreceive:     TypreceiveColumn,
		Typsend:        TypsendColumn,
		Typmodin:       TypmodinColumn,
		Typmodout:      TypmodoutColumn,
		Typanalyze:     TypanalyzeColumn,
		Typalign:       TypalignColumn,
		Typstorage:     TypstorageColumn,
		Typnotnull:     TypnotnullColumn,
		Typbasetype:    TypbasetypeColumn,
		Typtypmod:      TyptypmodColumn,
		Typndims:       TypndimsColumn,
		Typcollation:   TypcollationColumn,
		Typdefaultbin:  TypdefaultbinColumn,
		Typdefault:     TypdefaultColumn,
		Typacl:         TypaclColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
