package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDeleteGenericById(t *testing.T) {

	TEST_DB_PATH := os.TempDir() + "/_generics_test.db"
	os.Remove(TEST_DB_PATH)
	db, err := gorm.Open(sqlite.Open(TEST_DB_PATH), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&BaseStruct{})
	db.AutoMigrate(&ExtendedStruct1{})
	db.AutoMigrate(&ExtendedStruct2{})
	db.AutoMigrate(&ExtendedStruct3{})

	// Create
	db.Create(&BaseStruct{ID: 1})
	db.Create(&BaseStruct{ID: 2})
	db.Create(&BaseStruct{ID: 3})

	db.Create(&ExtendedStruct1{BaseStruct: BaseStruct{ID: 4}})
	db.Create(&ExtendedStruct1{BaseStruct: BaseStruct{ID: 5}})
	db.Create(&ExtendedStruct1{BaseStruct: BaseStruct{ID: 6}})

	db.Create(&ExtendedStruct2{BaseStruct: BaseStruct{ID: 7}})
	db.Create(&ExtendedStruct2{BaseStruct: BaseStruct{ID: 8}})
	db.Create(&ExtendedStruct2{BaseStruct: BaseStruct{ID: 9}})

	db.Create(&ExtendedStruct3{ExtendedStruct1: ExtendedStruct1{BaseStruct: BaseStruct{ID: 10}}})
	db.Create(&ExtendedStruct3{ExtendedStruct1: ExtendedStruct1{BaseStruct: BaseStruct{ID: 11}}})
	db.Create(&ExtendedStruct3{ExtendedStruct1: ExtendedStruct1{BaseStruct: BaseStruct{ID: 12}}})

	assert := assert.New(t)

	var baseStruct BaseStruct
	result := db.First(&baseStruct, 1) // find BaseStruct with integer primary key
	assert.Equal(int64(1), result.RowsAffected)
	assert.Equal(nil, result.Error)

	var extendedStruct1 ExtendedStruct1
	result = db.First(&extendedStruct1, 4) // find ExtendedStruct1 with integer primary key
	assert.Equal(int64(1), result.RowsAffected)
	assert.Equal(nil, result.Error)

	var extendedStruct2 ExtendedStruct2
	result = db.First(&extendedStruct2, 7) // find ExtendedStruct2 with integer primary key
	assert.Equal(int64(1), result.RowsAffected)
	assert.Equal(nil, result.Error)

	var extendedStruct3 ExtendedStruct3
	result = db.First(&extendedStruct3, 10) // find ExtendedStruct3 with integer primary key
	assert.Equal(int64(1), result.RowsAffected)
	assert.Equal(nil, result.Error)

	assert.Equal("Delete: &{ID:1} *main.BaseStruct", DeleteGenericById[*BaseStruct](db, 1))
	result = db.First(&baseStruct, 1) // find BaseStruct with integer primary key
	assert.Zero(result.RowsAffected)
	assert.True(errors.Is(result.Error, gorm.ErrRecordNotFound), "The BaseStruct should not be able to be found after it was deleted")

	assert.Equal("Delete: &{BaseStruct:{ID:4} foo:} *main.ExtendedStruct1", DeleteGenericById[*ExtendedStruct1](db, 4))
	result = db.First(&extendedStruct1, 4) // find ExtendedStruct1 with integer primary key
	assert.Zero(result.RowsAffected)
	assert.True(errors.Is(result.Error, gorm.ErrRecordNotFound), "The ExtendedStruct1 should not be able to be found after it was deleted")

	assert.Equal("Delete: &{BaseStruct:{ID:7} bar:} *main.ExtendedStruct2", DeleteGenericById[*ExtendedStruct2](db, 7))
	result = db.First(&extendedStruct2, 7) // find ExtendedStruct2 with integer primary key
	assert.Zero(result.RowsAffected)
	assert.True(errors.Is(result.Error, gorm.ErrRecordNotFound), "The ExtendedStruct2 should not be able to be found after it was deleted")

	assert.Equal("Delete: &{ExtendedStruct1:{BaseStruct:{ID:10} foo:} baz:} *main.ExtendedStruct3", DeleteGenericById[*ExtendedStruct3](db, 10))
	result = db.First(&extendedStruct3, 10) // find ExtendedStruct3 with integer primary key
	assert.Zero(result.RowsAffected)
	assert.True(errors.Is(result.Error, gorm.ErrRecordNotFound), "The ExtendedStruct3 should not be able to be found after it was deleted")

}
