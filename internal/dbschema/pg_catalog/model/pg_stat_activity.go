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

type PgStatActivity struct {
	Datid           *string
	Datname         *string
	Pid             *int32
	LeaderPid       *int32
	Usesysid        *string
	Usename         *string
	ApplicationName *string
	ClientAddr      *string
	ClientHostname  *string
	ClientPort      *int32
	BackendStart    *time.Time
	XactStart       *time.Time
	QueryStart      *time.Time
	StateChange     *time.Time
	WaitEventType   *string
	WaitEvent       *string
	State           *string
	BackendXid      *string
	BackendXmin     *string
	Query           *string
	BackendType     *string
}
