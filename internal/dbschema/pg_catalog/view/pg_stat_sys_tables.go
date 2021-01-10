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

var PgStatSysTables = newPgStatSysTablesTable()

type pgStatSysTablesTable struct {
	postgres.Table

	//Columns
	Relid            postgres.ColumnString
	Schemaname       postgres.ColumnString
	Relname          postgres.ColumnString
	SeqScan          postgres.ColumnInteger
	SeqTupRead       postgres.ColumnInteger
	IdxScan          postgres.ColumnInteger
	IdxTupFetch      postgres.ColumnInteger
	NTupIns          postgres.ColumnInteger
	NTupUpd          postgres.ColumnInteger
	NTupDel          postgres.ColumnInteger
	NTupHotUpd       postgres.ColumnInteger
	NLiveTup         postgres.ColumnInteger
	NDeadTup         postgres.ColumnInteger
	NModSinceAnalyze postgres.ColumnInteger
	NInsSinceVacuum  postgres.ColumnInteger
	LastVacuum       postgres.ColumnTimestampz
	LastAutovacuum   postgres.ColumnTimestampz
	LastAnalyze      postgres.ColumnTimestampz
	LastAutoanalyze  postgres.ColumnTimestampz
	VacuumCount      postgres.ColumnInteger
	AutovacuumCount  postgres.ColumnInteger
	AnalyzeCount     postgres.ColumnInteger
	AutoanalyzeCount postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgStatSysTablesTable struct {
	pgStatSysTablesTable

	EXCLUDED pgStatSysTablesTable
}

// AS creates new PgStatSysTablesTable with assigned alias
func (a *PgStatSysTablesTable) AS(alias string) *PgStatSysTablesTable {
	aliasTable := newPgStatSysTablesTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgStatSysTablesTable() *PgStatSysTablesTable {
	return &PgStatSysTablesTable{
		pgStatSysTablesTable: newPgStatSysTablesTableImpl("pg_catalog", "pg_stat_sys_tables"),
		EXCLUDED:             newPgStatSysTablesTableImpl("", "excluded"),
	}
}

func newPgStatSysTablesTableImpl(schemaName, tableName string) pgStatSysTablesTable {
	var (
		RelidColumn            = postgres.StringColumn("relid")
		SchemanameColumn       = postgres.StringColumn("schemaname")
		RelnameColumn          = postgres.StringColumn("relname")
		SeqScanColumn          = postgres.IntegerColumn("seq_scan")
		SeqTupReadColumn       = postgres.IntegerColumn("seq_tup_read")
		IdxScanColumn          = postgres.IntegerColumn("idx_scan")
		IdxTupFetchColumn      = postgres.IntegerColumn("idx_tup_fetch")
		NTupInsColumn          = postgres.IntegerColumn("n_tup_ins")
		NTupUpdColumn          = postgres.IntegerColumn("n_tup_upd")
		NTupDelColumn          = postgres.IntegerColumn("n_tup_del")
		NTupHotUpdColumn       = postgres.IntegerColumn("n_tup_hot_upd")
		NLiveTupColumn         = postgres.IntegerColumn("n_live_tup")
		NDeadTupColumn         = postgres.IntegerColumn("n_dead_tup")
		NModSinceAnalyzeColumn = postgres.IntegerColumn("n_mod_since_analyze")
		NInsSinceVacuumColumn  = postgres.IntegerColumn("n_ins_since_vacuum")
		LastVacuumColumn       = postgres.TimestampzColumn("last_vacuum")
		LastAutovacuumColumn   = postgres.TimestampzColumn("last_autovacuum")
		LastAnalyzeColumn      = postgres.TimestampzColumn("last_analyze")
		LastAutoanalyzeColumn  = postgres.TimestampzColumn("last_autoanalyze")
		VacuumCountColumn      = postgres.IntegerColumn("vacuum_count")
		AutovacuumCountColumn  = postgres.IntegerColumn("autovacuum_count")
		AnalyzeCountColumn     = postgres.IntegerColumn("analyze_count")
		AutoanalyzeCountColumn = postgres.IntegerColumn("autoanalyze_count")
		allColumns             = postgres.ColumnList{RelidColumn, SchemanameColumn, RelnameColumn, SeqScanColumn, SeqTupReadColumn, IdxScanColumn, IdxTupFetchColumn, NTupInsColumn, NTupUpdColumn, NTupDelColumn, NTupHotUpdColumn, NLiveTupColumn, NDeadTupColumn, NModSinceAnalyzeColumn, NInsSinceVacuumColumn, LastVacuumColumn, LastAutovacuumColumn, LastAnalyzeColumn, LastAutoanalyzeColumn, VacuumCountColumn, AutovacuumCountColumn, AnalyzeCountColumn, AutoanalyzeCountColumn}
		mutableColumns         = postgres.ColumnList{RelidColumn, SchemanameColumn, RelnameColumn, SeqScanColumn, SeqTupReadColumn, IdxScanColumn, IdxTupFetchColumn, NTupInsColumn, NTupUpdColumn, NTupDelColumn, NTupHotUpdColumn, NLiveTupColumn, NDeadTupColumn, NModSinceAnalyzeColumn, NInsSinceVacuumColumn, LastVacuumColumn, LastAutovacuumColumn, LastAnalyzeColumn, LastAutoanalyzeColumn, VacuumCountColumn, AutovacuumCountColumn, AnalyzeCountColumn, AutoanalyzeCountColumn}
	)

	return pgStatSysTablesTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Relid:            RelidColumn,
		Schemaname:       SchemanameColumn,
		Relname:          RelnameColumn,
		SeqScan:          SeqScanColumn,
		SeqTupRead:       SeqTupReadColumn,
		IdxScan:          IdxScanColumn,
		IdxTupFetch:      IdxTupFetchColumn,
		NTupIns:          NTupInsColumn,
		NTupUpd:          NTupUpdColumn,
		NTupDel:          NTupDelColumn,
		NTupHotUpd:       NTupHotUpdColumn,
		NLiveTup:         NLiveTupColumn,
		NDeadTup:         NDeadTupColumn,
		NModSinceAnalyze: NModSinceAnalyzeColumn,
		NInsSinceVacuum:  NInsSinceVacuumColumn,
		LastVacuum:       LastVacuumColumn,
		LastAutovacuum:   LastAutovacuumColumn,
		LastAnalyze:      LastAnalyzeColumn,
		LastAutoanalyze:  LastAutoanalyzeColumn,
		VacuumCount:      VacuumCountColumn,
		AutovacuumCount:  AutovacuumCountColumn,
		AnalyzeCount:     AnalyzeCountColumn,
		AutoanalyzeCount: AutoanalyzeCountColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
