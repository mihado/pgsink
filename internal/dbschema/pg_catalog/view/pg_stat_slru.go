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

var PgStatSlru = newPgStatSlruTable()

type pgStatSlruTable struct {
	postgres.Table

	//Columns
	Name        postgres.ColumnString
	BlksZeroed  postgres.ColumnInteger
	BlksHit     postgres.ColumnInteger
	BlksRead    postgres.ColumnInteger
	BlksWritten postgres.ColumnInteger
	BlksExists  postgres.ColumnInteger
	Flushes     postgres.ColumnInteger
	Truncates   postgres.ColumnInteger
	StatsReset  postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgStatSlruTable struct {
	pgStatSlruTable

	EXCLUDED pgStatSlruTable
}

// AS creates new PgStatSlruTable with assigned alias
func (a *PgStatSlruTable) AS(alias string) *PgStatSlruTable {
	aliasTable := newPgStatSlruTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgStatSlruTable() *PgStatSlruTable {
	return &PgStatSlruTable{
		pgStatSlruTable: newPgStatSlruTableImpl("pg_catalog", "pg_stat_slru"),
		EXCLUDED:        newPgStatSlruTableImpl("", "excluded"),
	}
}

func newPgStatSlruTableImpl(schemaName, tableName string) pgStatSlruTable {
	var (
		NameColumn        = postgres.StringColumn("name")
		BlksZeroedColumn  = postgres.IntegerColumn("blks_zeroed")
		BlksHitColumn     = postgres.IntegerColumn("blks_hit")
		BlksReadColumn    = postgres.IntegerColumn("blks_read")
		BlksWrittenColumn = postgres.IntegerColumn("blks_written")
		BlksExistsColumn  = postgres.IntegerColumn("blks_exists")
		FlushesColumn     = postgres.IntegerColumn("flushes")
		TruncatesColumn   = postgres.IntegerColumn("truncates")
		StatsResetColumn  = postgres.TimestampzColumn("stats_reset")
		allColumns        = postgres.ColumnList{NameColumn, BlksZeroedColumn, BlksHitColumn, BlksReadColumn, BlksWrittenColumn, BlksExistsColumn, FlushesColumn, TruncatesColumn, StatsResetColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, BlksZeroedColumn, BlksHitColumn, BlksReadColumn, BlksWrittenColumn, BlksExistsColumn, FlushesColumn, TruncatesColumn, StatsResetColumn}
	)

	return pgStatSlruTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Name:        NameColumn,
		BlksZeroed:  BlksZeroedColumn,
		BlksHit:     BlksHitColumn,
		BlksRead:    BlksReadColumn,
		BlksWritten: BlksWrittenColumn,
		BlksExists:  BlksExistsColumn,
		Flushes:     FlushesColumn,
		Truncates:   TruncatesColumn,
		StatsReset:  StatsResetColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
