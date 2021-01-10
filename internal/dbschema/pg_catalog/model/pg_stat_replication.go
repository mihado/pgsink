//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type PgStatReplication struct {
	Pid             *int32
	Usesysid        *string
	Usename         *string
	ApplicationName *string
	ClientAddr      *string
	ClientHostname  *string
	ClientPort      *int32
	BackendStart    *time.Time
	BackendXmin     *string
	State           *string
	SentLsn         *string
	WriteLsn        *string
	FlushLsn        *string
	ReplayLsn       *string
	WriteLag        *string
	FlushLag        *string
	ReplayLag       *string
	SyncPriority    *int32
	SyncState       *string
	ReplyTime       *time.Time
}
