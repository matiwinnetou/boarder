package utils

type Table struct {
	name string
}

func NewTable(n string) *Table {
	return &Table{name: n}
}
