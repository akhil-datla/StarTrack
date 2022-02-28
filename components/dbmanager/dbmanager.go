package dbmanager

import (
	"github.com/asdine/storm"	
)

var db *storm.DB

//Open opens the database
func Open(name string) error {
	var err error
	db, err = storm.Open(name)
	return err
}

//AutoCreateStruct creates a table for the struct
func AutoCreateStruct(data interface{}) error {
	err := db.Init(data)
	return err
}

//Save saves the data (struct)
func Save(data interface{}) error {
	err := db.Save(data)
	return err
}

//Query executes a query given a struct
func Query(fieldName string, value, to interface{}) error {
	err := db.One(fieldName, value, to)
	return err
}

func QueryAll(to interface{}) error {
	err := db.All(to)
	return err
}

//Update updates the data (struct)
func Update(data interface{}) error {
	err := db.Update(data)
	return err
}

//Delete deletes the data (struct)
func Delete(data interface{}) error {
	err := db.DeleteStruct(data)
	return err
}

//Close closes the database
func Close() error {
	return db.Close()
}
