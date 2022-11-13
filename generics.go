package main

import "fmt"

// based on: https://stackoverflow.com/a/33642734/5572674
type BaseStruct struct {
	Id int `orm:"pk"`
}

func (m *BaseStruct) setId(id int) {
	m.Id = id
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
	setId(id int)
	//constraining a type to a pointer of a given type (from: https://stackoverflow.com/a/70394905/5572674)
	*T
}

// mock similar to orm.Delete()
func Delete(md interface{}, cols ...string) string {
	return fmt.Sprintf("Delete: %+v %T", md, md)
}

func DeleteGenericById[T ISetId[U], U any](id int) string {
	var myU U
	T(&myU).setId(id)
	return Delete(T(&myU))
}
