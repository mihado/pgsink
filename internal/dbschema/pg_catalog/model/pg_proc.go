//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgProc struct {
	Oid             string
	Proname         string
	Pronamespace    string
	Proowner        string
	Prolang         string
	Procost         float32
	Prorows         float32
	Provariadic     string
	Prosupport      string
	Prokind         string
	Prosecdef       bool
	Proleakproof    bool
	Proisstrict     bool
	Proretset       bool
	Provolatile     string
	Proparallel     string
	Pronargs        int16
	Pronargdefaults int16
	Prorettype      string
	Proargtypes     string
	Proallargtypes  *string
	Proargmodes     *string
	Proargnames     *string
	Proargdefaults  *string
	Protrftypes     *string
	Prosrc          string
	Probin          *string
	Proconfig       *string
	Proacl          *string
}
