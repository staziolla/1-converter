package bins

import (
	"time"
)


type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinWithDb struct {
	Bin
	db Db
}	



func CreateBin(id string, private bool, createdAt time.Time, name string) *Bin {
	b := Bin{Id: id, Private: private, CreatedAt: createdAt, Name: name}
	return &b
}
