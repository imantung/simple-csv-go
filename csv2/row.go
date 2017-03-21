package csv2

import "fmt"

// Row is wrapper for csv row
type Row struct {
	RowNumber      int
	Columns        []string
	ColumnIndexMap map[string]int
}

// Get get value of column by index
func (r *Row) Get(i int) (value string, err error) {
	if i < len(r.Columns) {
		value = r.Columns[i]
	} else {
		err = fmt.Errorf("out of index: %d", i)
	}

	return
}

// GetByName get value of column by name
func (r *Row) GetByName(name string) (value string, err error) {

	index, ok := r.ColumnIndexMap[name]
	if !ok {
		err = fmt.Errorf("%s is not defined for this csv", name)
		return
	}

	value, err = r.Get(index)
	return
}
