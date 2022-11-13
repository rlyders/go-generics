package main

import (
	"fmt"

	"gorm.io/gorm"
)

// based on: https://stackoverflow.com/a/33642734/5572674
type BaseStruct struct {
	ID uint
}

func (m *BaseStruct) setId(id uint) {
	m.ID = id
}

type ExtendedStruct1 struct {
	BaseStruct
	foo string
}

type ExtendedStruct2 struct {
	BaseStruct
	bar string
}

type ExtendedStruct3 struct {
	ExtendedStruct1
	baz string
}

type ISetId[T any] interface {
	setId(id uint)
	//constraining a type to a pointer of a given type (from: https://stackoverflow.com/a/70394905/5572674)
	*T
}

// wrapper for gorm.Delete() to check type being deleted
func Delete(g *gorm.DB, md interface{}, cols ...string) string {
	ck := fmt.Sprintf("Delete: %+v %T", md, md)
	g.Delete(md)
	return ck
}

func DeleteGenericById[T ISetId[U], U any](g *gorm.DB, id uint) string {
	var myU U
	T(&myU).setId(id)
	return Delete(g, T(&myU))
}
