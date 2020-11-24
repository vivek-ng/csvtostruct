package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestParser1 struct {
	Field1 string `csv:"field1"`
	Field2 int    `csv:"field2"`
	TestParser2
}

type TestParser2 struct {
	Field3 string `csv:"field3"`
}

type TestParser3 struct {
	Field1 int `csv:"field1"`
	testParser4
}

type testParser4 struct {
	Field2 string `csv:"field2"`
}

func TestParser(t *testing.T) {
	newParser, _ := NewCSVStructer(&TestParser1{}, []string{"field1", "field2", "field3"})
	isValid := newParser.ValidateHeaders([]string{"field1", "field2", "field3"})
	assert.Equal(t, isValid, true)
	var parser TestParser1
	err := newParser.ScanStruct([]string{"apple", "43", "banana"}, &parser)
	assert.Nil(t, err)
	assert.Equal(t, parser.Field1, "apple")
	assert.Equal(t, parser.Field2, 43)
	assert.Equal(t, parser.Field3, "banana")
}

func TestParser_Error(t *testing.T) {
	newParser, err := NewCSVStructer(&TestParser1{}, []string{"field1", "field2"})
	assert.Nil(t, err)
	isValid := newParser.ValidateHeaders([]string{"field3", "field2"})
	assert.Equal(t, isValid, false)
	var parser TestParser1
	err = newParser.ScanStruct([]string{"apple", "banana"}, &parser)
	assert.Error(t, err)
}

func TestParser_UnexportedFields(t *testing.T) {
	newParser, err := NewCSVStructer(&TestParser3{}, []string{"field1", "field2"})
	assert.Nil(t, err)
	isValid := newParser.ValidateHeaders([]string{"field1", "field2"})
	assert.Equal(t, isValid, true)
	var parser TestParser3
	err = newParser.ScanStruct([]string{"10", "banana"}, &parser)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "struct contains unexported fields")
}
