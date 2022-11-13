package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Delete: &{BaseStruct:{Id:11} foo:} *main.ExtendedStruct1", DeleteGenericById[*ExtendedStruct1](11))
	assert.Equal("Delete: &{BaseStruct:{Id:22} bar:} *main.ExtendedStruct2", DeleteGenericById[*ExtendedStruct2](22))
	assert.Equal("Delete: &{ExtendedStruct1:{BaseStruct:{Id:33} foo:} baz:} *main.ExtendedStruct3", DeleteGenericById[*ExtendedStruct3](33))
}
