package csv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vivek-ng/csvtostruct/csv"
)

type testParser1 struct {
	Field1 string `csv:"field1"`
	Field2 int    `csv:"field2"`
}

func TestParser(t *testing.T) {
	newParser, _ := csv.NewCSVStructer(&testParser1{}, []string{"field1", "field2"})
	isValid := newParser.ValidateHeaders([]string{"field1", "field2"})
	assert.Equal(t, isValid, true)
	var parser testParser1
	err := newParser.ScanStruct([]string{"apple", "43"}, &parser)
	assert.Nil(t, err)
	assert.Equal(t, parser.Field1, "apple")
	assert.Equal(t, parser.Field2, 43)
}
