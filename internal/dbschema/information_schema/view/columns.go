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

var Columns = newColumnsTable()

type columnsTable struct {
	postgres.Table

	//Columns
	TableCatalog           postgres.ColumnString
	TableSchema            postgres.ColumnString
	TableName              postgres.ColumnString
	ColumnName             postgres.ColumnString
	OrdinalPosition        postgres.ColumnInteger
	ColumnDefault          postgres.ColumnString
	IsNullable             postgres.ColumnString
	DataType               postgres.ColumnString
	CharacterMaximumLength postgres.ColumnInteger
	CharacterOctetLength   postgres.ColumnInteger
	NumericPrecision       postgres.ColumnInteger
	NumericPrecisionRadix  postgres.ColumnInteger
	NumericScale           postgres.ColumnInteger
	DatetimePrecision      postgres.ColumnInteger
	IntervalType           postgres.ColumnString
	IntervalPrecision      postgres.ColumnInteger
	CharacterSetCatalog    postgres.ColumnString
	CharacterSetSchema     postgres.ColumnString
	CharacterSetName       postgres.ColumnString
	CollationCatalog       postgres.ColumnString
	CollationSchema        postgres.ColumnString
	CollationName          postgres.ColumnString
	DomainCatalog          postgres.ColumnString
	DomainSchema           postgres.ColumnString
	DomainName             postgres.ColumnString
	UdtCatalog             postgres.ColumnString
	UdtSchema              postgres.ColumnString
	UdtName                postgres.ColumnString
	ScopeCatalog           postgres.ColumnString
	ScopeSchema            postgres.ColumnString
	ScopeName              postgres.ColumnString
	MaximumCardinality     postgres.ColumnInteger
	DtdIdentifier          postgres.ColumnString
	IsSelfReferencing      postgres.ColumnString
	IsIdentity             postgres.ColumnString
	IdentityGeneration     postgres.ColumnString
	IdentityStart          postgres.ColumnString
	IdentityIncrement      postgres.ColumnString
	IdentityMaximum        postgres.ColumnString
	IdentityMinimum        postgres.ColumnString
	IdentityCycle          postgres.ColumnString
	IsGenerated            postgres.ColumnString
	GenerationExpression   postgres.ColumnString
	IsUpdatable            postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ColumnsTable struct {
	columnsTable

	EXCLUDED columnsTable
}

// AS creates new ColumnsTable with assigned alias
func (a *ColumnsTable) AS(alias string) *ColumnsTable {
	aliasTable := newColumnsTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newColumnsTable() *ColumnsTable {
	return &ColumnsTable{
		columnsTable: newColumnsTableImpl("information_schema", "columns"),
		EXCLUDED:     newColumnsTableImpl("", "excluded"),
	}
}

func newColumnsTableImpl(schemaName, tableName string) columnsTable {
	var (
		TableCatalogColumn           = postgres.StringColumn("table_catalog")
		TableSchemaColumn            = postgres.StringColumn("table_schema")
		TableNameColumn              = postgres.StringColumn("table_name")
		ColumnNameColumn             = postgres.StringColumn("column_name")
		OrdinalPositionColumn        = postgres.IntegerColumn("ordinal_position")
		ColumnDefaultColumn          = postgres.StringColumn("column_default")
		IsNullableColumn             = postgres.StringColumn("is_nullable")
		DataTypeColumn               = postgres.StringColumn("data_type")
		CharacterMaximumLengthColumn = postgres.IntegerColumn("character_maximum_length")
		CharacterOctetLengthColumn   = postgres.IntegerColumn("character_octet_length")
		NumericPrecisionColumn       = postgres.IntegerColumn("numeric_precision")
		NumericPrecisionRadixColumn  = postgres.IntegerColumn("numeric_precision_radix")
		NumericScaleColumn           = postgres.IntegerColumn("numeric_scale")
		DatetimePrecisionColumn      = postgres.IntegerColumn("datetime_precision")
		IntervalTypeColumn           = postgres.StringColumn("interval_type")
		IntervalPrecisionColumn      = postgres.IntegerColumn("interval_precision")
		CharacterSetCatalogColumn    = postgres.StringColumn("character_set_catalog")
		CharacterSetSchemaColumn     = postgres.StringColumn("character_set_schema")
		CharacterSetNameColumn       = postgres.StringColumn("character_set_name")
		CollationCatalogColumn       = postgres.StringColumn("collation_catalog")
		CollationSchemaColumn        = postgres.StringColumn("collation_schema")
		CollationNameColumn          = postgres.StringColumn("collation_name")
		DomainCatalogColumn          = postgres.StringColumn("domain_catalog")
		DomainSchemaColumn           = postgres.StringColumn("domain_schema")
		DomainNameColumn             = postgres.StringColumn("domain_name")
		UdtCatalogColumn             = postgres.StringColumn("udt_catalog")
		UdtSchemaColumn              = postgres.StringColumn("udt_schema")
		UdtNameColumn                = postgres.StringColumn("udt_name")
		ScopeCatalogColumn           = postgres.StringColumn("scope_catalog")
		ScopeSchemaColumn            = postgres.StringColumn("scope_schema")
		ScopeNameColumn              = postgres.StringColumn("scope_name")
		MaximumCardinalityColumn     = postgres.IntegerColumn("maximum_cardinality")
		DtdIdentifierColumn          = postgres.StringColumn("dtd_identifier")
		IsSelfReferencingColumn      = postgres.StringColumn("is_self_referencing")
		IsIdentityColumn             = postgres.StringColumn("is_identity")
		IdentityGenerationColumn     = postgres.StringColumn("identity_generation")
		IdentityStartColumn          = postgres.StringColumn("identity_start")
		IdentityIncrementColumn      = postgres.StringColumn("identity_increment")
		IdentityMaximumColumn        = postgres.StringColumn("identity_maximum")
		IdentityMinimumColumn        = postgres.StringColumn("identity_minimum")
		IdentityCycleColumn          = postgres.StringColumn("identity_cycle")
		IsGeneratedColumn            = postgres.StringColumn("is_generated")
		GenerationExpressionColumn   = postgres.StringColumn("generation_expression")
		IsUpdatableColumn            = postgres.StringColumn("is_updatable")
		allColumns                   = postgres.ColumnList{TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, OrdinalPositionColumn, ColumnDefaultColumn, IsNullableColumn, DataTypeColumn, CharacterMaximumLengthColumn, CharacterOctetLengthColumn, NumericPrecisionColumn, NumericPrecisionRadixColumn, NumericScaleColumn, DatetimePrecisionColumn, IntervalTypeColumn, IntervalPrecisionColumn, CharacterSetCatalogColumn, CharacterSetSchemaColumn, CharacterSetNameColumn, CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, DomainCatalogColumn, DomainSchemaColumn, DomainNameColumn, UdtCatalogColumn, UdtSchemaColumn, UdtNameColumn, ScopeCatalogColumn, ScopeSchemaColumn, ScopeNameColumn, MaximumCardinalityColumn, DtdIdentifierColumn, IsSelfReferencingColumn, IsIdentityColumn, IdentityGenerationColumn, IdentityStartColumn, IdentityIncrementColumn, IdentityMaximumColumn, IdentityMinimumColumn, IdentityCycleColumn, IsGeneratedColumn, GenerationExpressionColumn, IsUpdatableColumn}
		mutableColumns               = postgres.ColumnList{TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, OrdinalPositionColumn, ColumnDefaultColumn, IsNullableColumn, DataTypeColumn, CharacterMaximumLengthColumn, CharacterOctetLengthColumn, NumericPrecisionColumn, NumericPrecisionRadixColumn, NumericScaleColumn, DatetimePrecisionColumn, IntervalTypeColumn, IntervalPrecisionColumn, CharacterSetCatalogColumn, CharacterSetSchemaColumn, CharacterSetNameColumn, CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, DomainCatalogColumn, DomainSchemaColumn, DomainNameColumn, UdtCatalogColumn, UdtSchemaColumn, UdtNameColumn, ScopeCatalogColumn, ScopeSchemaColumn, ScopeNameColumn, MaximumCardinalityColumn, DtdIdentifierColumn, IsSelfReferencingColumn, IsIdentityColumn, IdentityGenerationColumn, IdentityStartColumn, IdentityIncrementColumn, IdentityMaximumColumn, IdentityMinimumColumn, IdentityCycleColumn, IsGeneratedColumn, GenerationExpressionColumn, IsUpdatableColumn}
	)

	return columnsTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		TableCatalog:           TableCatalogColumn,
		TableSchema:            TableSchemaColumn,
		TableName:              TableNameColumn,
		ColumnName:             ColumnNameColumn,
		OrdinalPosition:        OrdinalPositionColumn,
		ColumnDefault:          ColumnDefaultColumn,
		IsNullable:             IsNullableColumn,
		DataType:               DataTypeColumn,
		CharacterMaximumLength: CharacterMaximumLengthColumn,
		CharacterOctetLength:   CharacterOctetLengthColumn,
		NumericPrecision:       NumericPrecisionColumn,
		NumericPrecisionRadix:  NumericPrecisionRadixColumn,
		NumericScale:           NumericScaleColumn,
		DatetimePrecision:      DatetimePrecisionColumn,
		IntervalType:           IntervalTypeColumn,
		IntervalPrecision:      IntervalPrecisionColumn,
		CharacterSetCatalog:    CharacterSetCatalogColumn,
		CharacterSetSchema:     CharacterSetSchemaColumn,
		CharacterSetName:       CharacterSetNameColumn,
		CollationCatalog:       CollationCatalogColumn,
		CollationSchema:        CollationSchemaColumn,
		CollationName:          CollationNameColumn,
		DomainCatalog:          DomainCatalogColumn,
		DomainSchema:           DomainSchemaColumn,
		DomainName:             DomainNameColumn,
		UdtCatalog:             UdtCatalogColumn,
		UdtSchema:              UdtSchemaColumn,
		UdtName:                UdtNameColumn,
		ScopeCatalog:           ScopeCatalogColumn,
		ScopeSchema:            ScopeSchemaColumn,
		ScopeName:              ScopeNameColumn,
		MaximumCardinality:     MaximumCardinalityColumn,
		DtdIdentifier:          DtdIdentifierColumn,
		IsSelfReferencing:      IsSelfReferencingColumn,
		IsIdentity:             IsIdentityColumn,
		IdentityGeneration:     IdentityGenerationColumn,
		IdentityStart:          IdentityStartColumn,
		IdentityIncrement:      IdentityIncrementColumn,
		IdentityMaximum:        IdentityMaximumColumn,
		IdentityMinimum:        IdentityMinimumColumn,
		IdentityCycle:          IdentityCycleColumn,
		IsGenerated:            IsGeneratedColumn,
		GenerationExpression:   GenerationExpressionColumn,
		IsUpdatable:            IsUpdatableColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
