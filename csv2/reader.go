package csv2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reader is wrapper to read csv
type Reader struct {
	Path           string
	Separator      string
	File           *os.File
	Scanner        *bufio.Scanner
	HeaderColumns  []string
	ColumnIndexMap map[string]int
	Headers        []string
	RowNumber      int
}

// Open csv file
func (c *Reader) Open(path, separator string, withHeader bool) (err error) {
	c.Path = path
	c.Separator = separator
	c.File, err = os.Open(path)
	if err != nil {
		return
	}

	c.Scanner = bufio.NewScanner(c.File)

	if withHeader {
		if c.Next() {
			c.HeaderColumns = c.Read()
			c.ColumnIndexMap = make(map[string]int)

			for i, column := range c.HeaderColumns {
				c.ColumnIndexMap[column] = i
			}
		} else {
			err = fmt.Errorf("No header on empty file: %s", path)
			return
		}
	}

	return
}

// Validate if expetectedColumns is exist
func (c *Reader) Validate(expectedColumns []string) (ok bool, err error) {

loop_expected_columns:
	for _, expectedColumn := range expectedColumns {
		for _, column := range c.HeaderColumns {
			if column == expectedColumn {
				continue loop_expected_columns
			}
		}

		ok = false
		err = fmt.Errorf("Expected Column is missing: %s", expectedColumn)
		return
	}

	ok = true
	return
}

// Next check if next line available
func (c *Reader) Next() bool {
	return c.Scanner.Scan()
}

// Read to get all columns
func (c *Reader) Read() (columns []string) {
	columns = strings.Split(c.Scanner.Text(), c.Separator)

	for i, column := range columns {
		columns[i] = RemoveDoubleQuote(column)
	}

	return
}

// ReadAsRow to get as row
func (c *Reader) ReadAsRow() (row *Row, err error) {
	if c.ColumnIndexMap == nil {
		err = fmt.Errorf("Columns is not defined")
		return
	}

	columns := c.Read()

	row = &Row{
		RowNumber:      c.RowNumber,
		Columns:        columns,
		ColumnIndexMap: c.ColumnIndexMap,
	}

	c.RowNumber++

	return
}

// Close the file
func (c *Reader) Close() {
	if c.File != nil {
		c.File.Close()
	}
}
