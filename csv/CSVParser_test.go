package csv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vivek-ng/csvtostruct/csv"
)

type TestParser1 struct {
	Field1 string `csv:"field1"`
	Field2 int    `csv:"field2"`
	TestParser2
}

type TestParser2 struct {
	Field3 string `csv:"field3"`
}

func TestParser(t *testing.T) {
	newParser, _ := csv.NewCSVStructer(&TestParser1{}, []string{"field1", "field2", "field3"})
	isValid := newParser.ValidateHeaders([]string{"field1", "field2", "field3"})
	assert.Equal(t, isValid, true)
	var parser TestParser1
	err := newParser.ScanStruct([]string{"apple", "43", "banana"}, &parser)
	assert.Nil(t, err)
	assert.Equal(t, parser.Field1, "apple")
	assert.Equal(t, parser.Field2, 43)
	assert.Equal(t, parser.Field3, "banana")
}

func TestParser_error(t *testing.T) {
	newParser, err := csv.NewCSVStructer(&TestParser1{}, []string{"field1", "field2"})
	assert.Nil(t, err)
	isValid := newParser.ValidateHeaders([]string{"field3", "field2"})
	assert.Equal(t, isValid, false)
	var parser TestParser1
	err = newParser.ScanStruct([]string{"apple", "banana"}, &parser)
	assert.Error(t, err)
}
