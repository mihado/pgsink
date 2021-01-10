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

var DomainConstraints = newDomainConstraintsTable()

type domainConstraintsTable struct {
	postgres.Table

	//Columns
	ConstraintCatalog postgres.ColumnString
	ConstraintSchema  postgres.ColumnString
	ConstraintName    postgres.ColumnString
	DomainCatalog     postgres.ColumnString
	DomainSchema      postgres.ColumnString
	DomainName        postgres.ColumnString
	IsDeferrable      postgres.ColumnString
	InitiallyDeferred postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type DomainConstraintsTable struct {
	domainConstraintsTable

	EXCLUDED domainConstraintsTable
}

// AS creates new DomainConstraintsTable with assigned alias
func (a *DomainConstraintsTable) AS(alias string) *DomainConstraintsTable {
	aliasTable := newDomainConstraintsTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newDomainConstraintsTable() *DomainConstraintsTable {
	return &DomainConstraintsTable{
		domainConstraintsTable: newDomainConstraintsTableImpl("information_schema", "domain_constraints"),
		EXCLUDED:               newDomainConstraintsTableImpl("", "excluded"),
	}
}

func newDomainConstraintsTableImpl(schemaName, tableName string) domainConstraintsTable {
	var (
		ConstraintCatalogColumn = postgres.StringColumn("constraint_catalog")
		ConstraintSchemaColumn  = postgres.StringColumn("constraint_schema")
		ConstraintNameColumn    = postgres.StringColumn("constraint_name")
		DomainCatalogColumn     = postgres.StringColumn("domain_catalog")
		DomainSchemaColumn      = postgres.StringColumn("domain_schema")
		DomainNameColumn        = postgres.StringColumn("domain_name")
		IsDeferrableColumn      = postgres.StringColumn("is_deferrable")
		InitiallyDeferredColumn = postgres.StringColumn("initially_deferred")
		allColumns              = postgres.ColumnList{ConstraintCatalogColumn, ConstraintSchemaColumn, ConstraintNameColumn, DomainCatalogColumn, DomainSchemaColumn, DomainNameColumn, IsDeferrableColumn, InitiallyDeferredColumn}
		mutableColumns          = postgres.ColumnList{ConstraintCatalogColumn, ConstraintSchemaColumn, ConstraintNameColumn, DomainCatalogColumn, DomainSchemaColumn, DomainNameColumn, IsDeferrableColumn, InitiallyDeferredColumn}
	)

	return domainConstraintsTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		ConstraintCatalog: ConstraintCatalogColumn,
		ConstraintSchema:  ConstraintSchemaColumn,
		ConstraintName:    ConstraintNameColumn,
		DomainCatalog:     DomainCatalogColumn,
		DomainSchema:      DomainSchemaColumn,
		DomainName:        DomainNameColumn,
		IsDeferrable:      IsDeferrableColumn,
		InitiallyDeferred: InitiallyDeferredColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
