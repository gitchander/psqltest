package main

import (
	"fmt"
	"strings"
)

type baseQuerier struct{}

func (baseQuerier) createTable(tableName string, cis []ColumnInfo) string {
	vs := make([]string, len(cis))
	for i, ci := range cis {
		vs[i] = ci.String()
	}
	columnsBlock := framing(strings.Join(vs, ",\n"))

	op := "CREATE TABLE"
	if true {
		op = op + " IF NOT EXISTS"
	}

	return fmt.Sprintf("%s %s %s;", op, tableName, columnsBlock)
}

func (baseQuerier) dropTable(tableName string) string {
	return fmt.Sprintf("DROP TABLE %s;", tableName)
}

func (baseQuerier) selectAll(tableName string) string {
	return fmt.Sprintf("SELECT * FROM %s;", tableName)
}

func (baseQuerier) deleteAllRecords(tableName string) string {
	return fmt.Sprintf("DELETE FROM %s;", tableName)
}

func (baseQuerier) appendRecord(tableName string, fieldNames []string) string {
	const delim = ", "
	var (
		fieldsBlock = framing(strings.Join(fieldNames, delim))
		binds       = makeRecordBinds(len(fieldNames))
	)
	return fmt.Sprintf("INSERT INTO %s %s values %s;", tableName, fieldsBlock, binds)
}

func (baseQuerier) numberOfRecords(tableName string) (query string) {
	return fmt.Sprintf("SELECT COUNT(*) FROM %s;", tableName)
}

func (baseQuerier) selectRange(tableName string, fieldName string,
	loc [2]int) (query string, args []interface{}) {

	// orderBlock

	// ORDER BY fieldName ASC
	// ORDER BY fieldName DESC - reverse order

	reverse := false
	orderBlock := fmt.Sprintf("ORDER BY %s", fieldName)
	if reverse {
		orderBlock += " DESC"
	} else {
		orderBlock += " ASC"
	}

	//--------------------------------------------------------------------------
	// rangeBlock
	var (
		offset = loc[0]
		limit  = loc[1] - loc[0]
	)
	args = []interface{}{
		offset, limit,
	}
	rangeBlock := "OFFSET $1 LIMIT $2"

	//--------------------------------------------------------------------------
	query = fmt.Sprintf("SELECT * FROM %s %s %s;", tableName, orderBlock, rangeBlock)
	return query, args
}

var querier = baseQuerier{}
