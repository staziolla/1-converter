package bins

import (
	"encoding/json"
	"fmt"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte) error
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

type BinsListWithDb struct {
	Bins []Bin
	db   Db
}

func CreateBin(id string, private bool, createdAt time.Time, name string) *Bin {
	b := Bin{Id: id, Private: private, CreatedAt: createdAt, Name: name}
	return &b
}

func CreateBinsListWithDb(db Db) *BinsListWithDb {

	data, err := db.Read()
	if err != nil {
		binsList := BinsListWithDb{Bins: []Bin{}, db: db}
		return &binsList
	}
	var binsload []Bin
	err = json.Unmarshal(data, &binsload)
	if err != nil {
		fmt.Println("Не удалось разобрать файл ")
		return &BinsListWithDb{Bins: []Bin{}, db: db}
	}

	return &BinsListWithDb{Bins: binsload, db: db}

	//	binsList := BinsListWithDb{Bins: []Bin{}, db: db}
	//	return &binsList
}

func (bl *BinsListWithDb) Append(id string, private bool, createdAt time.Time, name string) error {
	bl.Bins = append(bl.Bins, *CreateBin(id, private, createdAt, name))
	fmt.Println(len(bl.Bins))
	return nil
}

func (bl *BinsListWithDb) ShowData() {
	fmt.Println(bl.db)
}

func (bl *BinsListWithDb) Save() error {
	data, err := json.Marshal(bl.Bins)
	if err != nil {
		fmt.Println("Ошибка сериализации " + err.Error())
		return err
	}
	err = bl.db.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи " + err.Error())
		return err
	}

	return nil

}

// func CreateBinWithDb(db Db, id string, private bool, createdAt time.Time, name string) *BinWithDb {
// 	//b := Bin{Id: id, Private: private, CreatedAt: createdAt, Name: name}
// 	b := BinWithDb{
// 		Bin: Bin{Id: id, Private: private, CreatedAt: createdAt, Name: name},
// 		db: db,
// 	}
// 	return &b
// }
