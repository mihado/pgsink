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

var PgStatWalReceiver = newPgStatWalReceiverTable()

type pgStatWalReceiverTable struct {
	postgres.Table

	//Columns
	Pid                postgres.ColumnInteger
	Status             postgres.ColumnString
	ReceiveStartLsn    postgres.ColumnString
	ReceiveStartTli    postgres.ColumnInteger
	WrittenLsn         postgres.ColumnString
	FlushedLsn         postgres.ColumnString
	ReceivedTli        postgres.ColumnInteger
	LastMsgSendTime    postgres.ColumnTimestampz
	LastMsgReceiptTime postgres.ColumnTimestampz
	LatestEndLsn       postgres.ColumnString
	LatestEndTime      postgres.ColumnTimestampz
	SlotName           postgres.ColumnString
	SenderHost         postgres.ColumnString
	SenderPort         postgres.ColumnInteger
	Conninfo           postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgStatWalReceiverTable struct {
	pgStatWalReceiverTable

	EXCLUDED pgStatWalReceiverTable
}

// AS creates new PgStatWalReceiverTable with assigned alias
func (a *PgStatWalReceiverTable) AS(alias string) *PgStatWalReceiverTable {
	aliasTable := newPgStatWalReceiverTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgStatWalReceiverTable() *PgStatWalReceiverTable {
	return &PgStatWalReceiverTable{
		pgStatWalReceiverTable: newPgStatWalReceiverTableImpl("pg_catalog", "pg_stat_wal_receiver"),
		EXCLUDED:               newPgStatWalReceiverTableImpl("", "excluded"),
	}
}

func newPgStatWalReceiverTableImpl(schemaName, tableName string) pgStatWalReceiverTable {
	var (
		PidColumn                = postgres.IntegerColumn("pid")
		StatusColumn             = postgres.StringColumn("status")
		ReceiveStartLsnColumn    = postgres.StringColumn("receive_start_lsn")
		ReceiveStartTliColumn    = postgres.IntegerColumn("receive_start_tli")
		WrittenLsnColumn         = postgres.StringColumn("written_lsn")
		FlushedLsnColumn         = postgres.StringColumn("flushed_lsn")
		ReceivedTliColumn        = postgres.IntegerColumn("received_tli")
		LastMsgSendTimeColumn    = postgres.TimestampzColumn("last_msg_send_time")
		LastMsgReceiptTimeColumn = postgres.TimestampzColumn("last_msg_receipt_time")
		LatestEndLsnColumn       = postgres.StringColumn("latest_end_lsn")
		LatestEndTimeColumn      = postgres.TimestampzColumn("latest_end_time")
		SlotNameColumn           = postgres.StringColumn("slot_name")
		SenderHostColumn         = postgres.StringColumn("sender_host")
		SenderPortColumn         = postgres.IntegerColumn("sender_port")
		ConninfoColumn           = postgres.StringColumn("conninfo")
		allColumns               = postgres.ColumnList{PidColumn, StatusColumn, ReceiveStartLsnColumn, ReceiveStartTliColumn, WrittenLsnColumn, FlushedLsnColumn, ReceivedTliColumn, LastMsgSendTimeColumn, LastMsgReceiptTimeColumn, LatestEndLsnColumn, LatestEndTimeColumn, SlotNameColumn, SenderHostColumn, SenderPortColumn, ConninfoColumn}
		mutableColumns           = postgres.ColumnList{PidColumn, StatusColumn, ReceiveStartLsnColumn, ReceiveStartTliColumn, WrittenLsnColumn, FlushedLsnColumn, ReceivedTliColumn, LastMsgSendTimeColumn, LastMsgReceiptTimeColumn, LatestEndLsnColumn, LatestEndTimeColumn, SlotNameColumn, SenderHostColumn, SenderPortColumn, ConninfoColumn}
	)

	return pgStatWalReceiverTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Pid:                PidColumn,
		Status:             StatusColumn,
		ReceiveStartLsn:    ReceiveStartLsnColumn,
		ReceiveStartTli:    ReceiveStartTliColumn,
		WrittenLsn:         WrittenLsnColumn,
		FlushedLsn:         FlushedLsnColumn,
		ReceivedTli:        ReceivedTliColumn,
		LastMsgSendTime:    LastMsgSendTimeColumn,
		LastMsgReceiptTime: LastMsgReceiptTimeColumn,
		LatestEndLsn:       LatestEndLsnColumn,
		LatestEndTime:      LatestEndTimeColumn,
		SlotName:           SlotNameColumn,
		SenderHost:         SenderHostColumn,
		SenderPort:         SenderPortColumn,
		Conninfo:           ConninfoColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
