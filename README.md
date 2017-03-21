# Simple CSV Go #

A simpler way to read csv on go

### Install ###
```
$ go get github.com/imantung/simple-csv-go
```

### Open Csv ###
```
reader := csv2.Reader{}
err := reader.Open("path.of.csv", ";", true)
if err != nil {
  log.Fatal(err)
}
defer reader.Close()
```

### Validate ###
```
expectedColumn := []string{"column1", "column2"}
if ok, err := reader.Validate(expectedColumn); !ok {
  log.Fatalf("Invalid csv: %s", err.Error())
}
```

### Read All ###
Feel like read file in usual way
```
for reader.Next() {
  row, err := reader.Read()
  if err != nil {
    log.Printf("This row is ignored due to error: %s\n", err.Error())
  }

  column1, _ := row[0]
  column2, _ := row[1]

  log.Printf("'column1':%s 'column2':%s", rowNumber, name, age)
}

log.Println("Done!")
```

The reader will save index of column so you can get the value by column name.
```
for reader.Next() {
  row, err := reader.ReadAsRow()
  if err != nil {
    log.Printf("This row is ignored due to error: %s\n", err.Error())
  }

  rowNumber := row.RowNumber
  column1, _ := row.GetByName("column1")
  column2, _ := row.GetByName("column2")

  log.Printf("Line:%d 'column1':%s 'column2':%s", rowNumber, name, age)
}

log.Println("Done!")
```





