package repository

import (
	"log"

	"github.com/satori/go.uuid"
)

var db *BoltDB

func init() {
	db = New()

	items, _ := GetAllItems()
	if len(items) != 0 {
		return
	}

	CreateItem(Item{Name: "Oathkeeper", Description: "Valyrian steel. Wielded by the Lord Commander of the Kingsguard.", Price: 153240})
	CreateItem(Item{Name: "Widowmaker", Description: "Valyrian steel. Whereabouts unknown.", Price: 450321})
}

func CreateItem(i Item) error {
	i.ID = uuid.NewV4().String()

	err := db.addItem(i)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetAllItems() ([]Item, error) {
	return db.getAllItems()
}
