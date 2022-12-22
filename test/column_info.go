package main

import (
	"strings"
)

type ColumnInfo struct {
	ColumnName        string
	DataType          string
	CompressionMethod string
	Constraints       []string
}

// ColumnInfo Constraints:
const (
	PrimaryKey = "PRIMARY KEY"
	NotNull    = "NOT NULL"
)

func (ci ColumnInfo) String() string {
	cs := []string{
		ci.ColumnName,
		ci.DataType,
	}
	if ci.CompressionMethod != "" {
		cs = append(cs, ci.CompressionMethod)
	}
	cs = append(cs, ci.Constraints...)
	return strings.Join(cs, " ")
}
