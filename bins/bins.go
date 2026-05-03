package bins

import (
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func CreateBin(id string, private bool, createdAt time.Time, name string) *Bin {
	b := Bin{Id: id, Private: private, CreatedAt: createdAt, Name: name}
	return &b
}
