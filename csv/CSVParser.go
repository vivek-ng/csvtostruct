package csv

import (
	"errors"
	"reflect"
	"strconv"
)

type CSVStruct struct {
	headers []string
}

func NewCSVStructer(input interface{}, headers []string) (*CSVStruct, error) {
	c := CSVStruct{
		headers: headers,
	}
	return &c, nil
}

func (c *CSVStruct) ValidateHeaders(csvHeaders []string) bool {
	if len(c.headers) != len(csvHeaders) {
		return false
	}
	for _, val := range csvHeaders {
		if !isPresent(val, c.headers) {
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
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.Type().Kind() == reflect.Struct {
			err := c.ScanStruct(csvRow, f.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}
		csvTag := reflect.TypeOf(inputStruct).Elem().Field(i).Tag.Get("csv")
		idx := index(csvTag, c.headers)
		if idx == -1 {
			continue
		}
		switch f.Type().Kind() {
		case reflect.String:
			f.SetString(csvRow[idx])
		case reflect.Int:
			ival, err := strconv.ParseInt(csvRow[idx], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		case reflect.Bool:
			ival, err := strconv.ParseBool(csvRow[idx])
			if err != nil {
				return err
			}
			f.SetBool(ival)
		}
	}
	return nil

}

func isPresent(val string, allValues []string) bool {

	for _, value := range allValues {
		if val == value {
			return true
		}
	}

	return false
}

func index(val string, allValues []string) int {
	for ind, value := range allValues {
		if value == val {
			return ind
		}
	}
	return -1
}
