//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgAggregate struct {
	Aggfnoid         string
	Aggkind          string
	Aggnumdirectargs int16
	Aggtransfn       string
	Aggfinalfn       string
	Aggcombinefn     string
	Aggserialfn      string
	Aggdeserialfn    string
	Aggmtransfn      string
	Aggminvtransfn   string
	Aggmfinalfn      string
	Aggfinalextra    bool
	Aggmfinalextra   bool
	Aggfinalmodify   string
	Aggmfinalmodify  string
	Aggsortop        string
	Aggtranstype     string
	Aggtransspace    int32
	Aggmtranstype    string
	Aggmtransspace   int32
	Agginitval       *string
	Aggminitval      *string
}
