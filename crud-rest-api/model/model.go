package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"strconv"
)

var db *gorm.DB

type Record struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func ConnectDatabase() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Record{}) // ok for development, suggest using Goblin for more complex requirements.
}

func DisconnectDatabase() {
	db.Close()
}

func CreateRecord(record *Record) (*Record, error) {
	record.ID = strconv.Itoa(rand.Intn(1000000))
	db.Create(record)
	return record, nil
}
