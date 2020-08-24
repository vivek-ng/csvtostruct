package main

import (
	"fmt"

	"github.com/vivek-ng/csvtostruct/csv"
)

type Person struct {
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}

func main() {
	p := Person{}
	headers := []string{"name", "age"}
	csvScanner, err := csv.NewCSVStructer(p, headers)
	if err != nil {
		fmt.Println(err)
	}
	b := csvScanner.ValidateHeaders()
	fmt.Println(b)
	csvScanner.ScanStruct([]string{"vivek", "34"}, &p)
	fmt.Println(p)
}
