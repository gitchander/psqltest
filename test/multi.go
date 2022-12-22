package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

// https://go.dev/doc/database/sql-injection

// Binds - placeholders

type MultiParams interface {
	Columns() []string
	Append(args []interface{}, i int) []interface{}
	Len() int
}

type InsertParameters struct {
	Columns string // (column1, column2, …)
	Binds   string // ($1, $2, ...), ($7, $8, ...), ...
	Args    []interface{}
}

func insertIntoMulti(db *sql.DB, tableName string, mp MultiParams) error {

	v := makeInsertParameters(mp)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", tableName, v.Columns, v.Binds)
	fmt.Println(query)

	_, err := db.Exec(query, v.Args...)
	return err
}

func makeInsertParameters(mp MultiParams) InsertParameters {

	const delim = ", "

	columns := mp.Columns()

	totalRecords := mp.Len()

	fieldsPerRecord := len(columns)

	var b strings.Builder

	nextNumber := nextNumberFunc()

	args := make([]interface{}, 0, (totalRecords * fieldsPerRecord))

	for i := 0; i < totalRecords; i++ {
		writeRecordBinds(&b, fieldsPerRecord, nextNumber)
		b.WriteString(delim)

		args = mp.Append(args, i)
	}

	binds := strings.TrimSuffix(b.String(), delim)

	return InsertParameters{
		Columns: strings.Join(columns, delim),
		Binds:   binds,
		Args:    args,
	}
}

func writeRecordBinds(b *strings.Builder, fieldsPerRecord int, nextNumber func() int) {
	const delim = ", "
	b.WriteByte('(')
	for j := 0; j < fieldsPerRecord; j++ {
		if j > 0 {
			b.WriteString(delim)
		}
		if false {
			b.WriteByte('$')
			b.WriteString(strconv.Itoa(nextNumber()))
		} else {
			fmt.Fprintf(b, "$%d", nextNumber())
		}
	}
	b.WriteByte(')')
}

// ($1, $2, ... $n)
func makeRecordBinds(n int) string {
	var b strings.Builder
	writeRecordBinds(&b, n, nextNumberFunc())
	return b.String()
}

func nextNumberFunc() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}
