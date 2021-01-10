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

var RoleUsageGrants = newRoleUsageGrantsTable()

type roleUsageGrantsTable struct {
	postgres.Table

	//Columns
	Grantor       postgres.ColumnString
	Grantee       postgres.ColumnString
	ObjectCatalog postgres.ColumnString
	ObjectSchema  postgres.ColumnString
	ObjectName    postgres.ColumnString
	ObjectType    postgres.ColumnString
	PrivilegeType postgres.ColumnString
	IsGrantable   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type RoleUsageGrantsTable struct {
	roleUsageGrantsTable

	EXCLUDED roleUsageGrantsTable
}

// AS creates new RoleUsageGrantsTable with assigned alias
func (a *RoleUsageGrantsTable) AS(alias string) *RoleUsageGrantsTable {
	aliasTable := newRoleUsageGrantsTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newRoleUsageGrantsTable() *RoleUsageGrantsTable {
	return &RoleUsageGrantsTable{
		roleUsageGrantsTable: newRoleUsageGrantsTableImpl("information_schema", "role_usage_grants"),
		EXCLUDED:             newRoleUsageGrantsTableImpl("", "excluded"),
	}
}

func newRoleUsageGrantsTableImpl(schemaName, tableName string) roleUsageGrantsTable {
	var (
		GrantorColumn       = postgres.StringColumn("grantor")
		GranteeColumn       = postgres.StringColumn("grantee")
		ObjectCatalogColumn = postgres.StringColumn("object_catalog")
		ObjectSchemaColumn  = postgres.StringColumn("object_schema")
		ObjectNameColumn    = postgres.StringColumn("object_name")
		ObjectTypeColumn    = postgres.StringColumn("object_type")
		PrivilegeTypeColumn = postgres.StringColumn("privilege_type")
		IsGrantableColumn   = postgres.StringColumn("is_grantable")
		allColumns          = postgres.ColumnList{GrantorColumn, GranteeColumn, ObjectCatalogColumn, ObjectSchemaColumn, ObjectNameColumn, ObjectTypeColumn, PrivilegeTypeColumn, IsGrantableColumn}
		mutableColumns      = postgres.ColumnList{GrantorColumn, GranteeColumn, ObjectCatalogColumn, ObjectSchemaColumn, ObjectNameColumn, ObjectTypeColumn, PrivilegeTypeColumn, IsGrantableColumn}
	)

	return roleUsageGrantsTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Grantor:       GrantorColumn,
		Grantee:       GranteeColumn,
		ObjectCatalog: ObjectCatalogColumn,
		ObjectSchema:  ObjectSchemaColumn,
		ObjectName:    ObjectNameColumn,
		ObjectType:    ObjectTypeColumn,
		PrivilegeType: PrivilegeTypeColumn,
		IsGrantable:   IsGrantableColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
