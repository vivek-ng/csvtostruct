package csv

import (
	"errors"
	"reflect"
	"strconv"
)

type CSVStruct struct {
	data    interface{}
	headers []string
}

func NewCSVStructer(input interface{}, headers []string) (*CSVStruct, error) {
	c := CSVStruct{
		data:    input,
		headers: headers,
	}
	return &c, nil
}

func (c *CSVStruct) ValidateHeaders() bool {
	if c.data == nil {
		return false
	}
	s := reflect.TypeOf(c.data)
	if len(c.headers) != s.NumField() {
		return false
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		currentTag := f.Tag.Get("csv")
		ind := isPresent(currentTag, c.headers)
		if !ind {
			return false
		}
	}
	return true
}

func (c *CSVStruct) ScanStruct(csvRow []string, inputStruct interface{}) error { //nolint: gocyclo
	s := reflect.ValueOf(inputStruct)
	if s.Kind() != reflect.Ptr {
		return errors.New("input should be a pointer to a struct")
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return errors.New("input should be a pointer to a struct")
	}
	//fmt.Println(s.String())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().Kind() {
		case reflect.String:
			f.SetString(csvRow[i])
		case reflect.Int:
			ival, err := strconv.ParseInt(csvRow[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		case reflect.Bool:
			ival, err := strconv.ParseBool(csvRow[i])
			if err != nil {
				return err
			}
			f.SetBool(ival)
		}
	}
	return nil

}

// func main() {
// 	test := []string{"Vivek", "45"}
// 	p := Person{}
// 	ScanStruct(test, &p)
// }

func isPresent(val string, allValues []string) bool {

	for _, value := range allValues {
		if val == value {
			return true
		}
	}

	return false
}
