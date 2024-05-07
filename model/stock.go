package model

import "time"

type Stock struct {
	Date      time.Time
	Kode      string
	LocalIS   uint64
	LocalCP   uint64
	LocalPF   uint64
	LocalIB   uint64
	LocalID   uint64
	LocalMF   uint64
	LocalSC   uint64
	LocalFD   uint64
	LocalOT   uint64
	ForeignIS uint64
	ForeignCP uint64
	ForeignPF uint64
	ForeignIB uint64
	ForeignID uint64
	ForeignMF uint64
	ForeignSC uint64
	ForeignFD uint64
	ForeignOT uint64
}
