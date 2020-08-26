# csvtostruct
Go Package to convert CSV fields to Struct

# Usage

```
type testParser1 struct {
	Field1 string `csv:"field1"`
	Field2 int    `csv:"field2"`
}
headerFields := []string{"field1" , "field2"}
row := 0
for {
		record, err := testFile.Read()
		if err == io.EOF {
			break
		}
		
    newCSVParser := csv.NewCSVStructer(&testParser1{}, headerFields)
    if row == 0 {
      if !csv.ValidateHeaders(record) {
        break
        }
      row+=1
    }
    var parser testParser1
    err := csv.ScanStruct(record , &parser)
    // Now parser struct will contain the csv record
	}
