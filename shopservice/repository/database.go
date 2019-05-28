package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	location string
}

var itemsBucket = []byte("items")

func New() *BoltDB {
	db := &BoltDB{
		location: "/etc/docker-demo/items.db",
	}

	db.createBuckets()
	return db
}

func (b *BoltDB) createBuckets() error {
	db, err := bolt.Open(b.location, 0644, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(itemsBucket)
		return err
	})
}

func (b *BoltDB) addItem(i Item) error {
	db, err := bolt.Open(b.location, 0644, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(itemsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", itemsBucket)
		}

		key := []byte(i.ID)
		value, err := json.Marshal(i)

		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})
}

func (b *BoltDB) getAllItems() ([]Item, error) {
	db, err := bolt.Open(b.location, 0644, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var items []Item
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(itemsBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", itemsBucket)
		}

		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var i Item
			err = json.Unmarshal(v, &i)
			if err != nil {
				return err
			}

			items = append(items, i)
		}

		return nil

	})

	return items, err
}
