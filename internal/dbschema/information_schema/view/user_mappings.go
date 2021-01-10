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

var UserMappings = newUserMappingsTable()

type userMappingsTable struct {
	postgres.Table

	//Columns
	AuthorizationIdentifier postgres.ColumnString
	ForeignServerCatalog    postgres.ColumnString
	ForeignServerName       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserMappingsTable struct {
	userMappingsTable

	EXCLUDED userMappingsTable
}

// AS creates new UserMappingsTable with assigned alias
func (a *UserMappingsTable) AS(alias string) *UserMappingsTable {
	aliasTable := newUserMappingsTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newUserMappingsTable() *UserMappingsTable {
	return &UserMappingsTable{
		userMappingsTable: newUserMappingsTableImpl("information_schema", "user_mappings"),
		EXCLUDED:          newUserMappingsTableImpl("", "excluded"),
	}
}

func newUserMappingsTableImpl(schemaName, tableName string) userMappingsTable {
	var (
		AuthorizationIdentifierColumn = postgres.StringColumn("authorization_identifier")
		ForeignServerCatalogColumn    = postgres.StringColumn("foreign_server_catalog")
		ForeignServerNameColumn       = postgres.StringColumn("foreign_server_name")
		allColumns                    = postgres.ColumnList{AuthorizationIdentifierColumn, ForeignServerCatalogColumn, ForeignServerNameColumn}
		mutableColumns                = postgres.ColumnList{AuthorizationIdentifierColumn, ForeignServerCatalogColumn, ForeignServerNameColumn}
	)

	return userMappingsTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		AuthorizationIdentifier: AuthorizationIdentifierColumn,
		ForeignServerCatalog:    ForeignServerCatalogColumn,
		ForeignServerName:       ForeignServerNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
