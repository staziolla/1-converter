package bins

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func CreateBin(id string, private bool, createdAt time.Time, name string) *Bin {
	b := Bin{id: id, private: private, createdAt: createdAt, name: name}
	return &b
}
