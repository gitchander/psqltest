package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gitchander/psqltest/random"
)

// Aliases

const tableNameTasks = "tasks"

type Task struct {
	ID        int
	Alias     string
	Timestamp time.Time
	Groups    string
	Number    string
	FieldU64  uint64
}

func createTableTasks(db *sql.DB) error {

	cis := []ColumnInfo{
		ColumnInfo{
			ColumnName: "task_id",
			DataType:   "serial",
			Constraints: []string{
				PrimaryKey,
			},
		},
		ColumnInfo{
			ColumnName: "task_alias",
			DataType:   "VARCHAR(50)",
			Constraints: []string{
				NotNull,
			},
		},
		ColumnInfo{
			ColumnName: "timestamp",
			DataType:   "TIMESTAMP WITH TIME ZONE",
			Constraints: []string{
				NotNull,
			},
		},
		ColumnInfo{
			ColumnName: "groups",
			DataType:   "VARCHAR(256)",
			Constraints: []string{
				NotNull,
			},
		},
		ColumnInfo{
			ColumnName: "number",
			DataType:   "VARCHAR(40)",
			Constraints: []string{
				NotNull,
			},
		},
		ColumnInfo{
			ColumnName: "field_u64",
			DataType:   "BIGINT",
			Constraints: []string{
				NotNull,
			},
		},
	}

	query := querier.createTable(tableNameTasks, cis)

	fmt.Println("query:", query)

	_, err := db.Exec(query)
	return err
}

func dropTableTasks(db *sql.DB) error {
	query := querier.dropTable(tableNameTasks)
	_, err := db.Exec(query)
	return err
}

func randTasks(db *sql.DB, n int) ([]Task, error) {
	ts := make([]Task, n)
	r := random.NewRandNow()
	for i := range ts {
		t := randTask(r)
		ts[i] = *t
	}
	return ts, nil
}

func appendRandTasks(db *sql.DB, n int) error {
	r := random.NewRandNow()
	for i := 0; i < n; i++ {
		t := randTask(r)
		err := appendTask(db, t)
		if err != nil {
			return err
		}
	}
	return nil
}

// ------------------------------------------------------------------------------
func appendTask(db *sql.DB, t *Task) error {

	var ti taskInserter
	var (
		columns = ti.Columns()
		args    = ti.Args(t)
	)

	query := querier.appendRecord(tableNameTasks, columns)
	_, err := db.Exec(query, args...)
	return err
}

func testAppendMulti(db *sql.DB) error {
	ts, err := randTasks(db, 13)
	if err != nil {
		return err
	}
	return appendMulti(db, ts)
}

func appendMulti(db *sql.DB, ts []Task) error {
	v := tasksInsert{
		tasks: ts,
	}
	return insertIntoMulti(db, tableNameTasks, v)
}

// ------------------------------------------------------------------------------
// taskArgumenter, taskInserter
type taskInserter struct{}

func (taskInserter) Columns() []string {
	return []string{
		"task_alias",
		"timestamp",
		"groups",
		"number",
		"field_u64",
	}
}

func (taskInserter) appendTask(args []interface{}, t *Task) []interface{} {
	return append(args,
		t.Alias,
		t.Timestamp,
		t.Groups,
		t.Number,
		t.FieldU64,
	)
}

func (v taskInserter) Args(t *Task) []interface{} {
	return v.appendTask(nil, t)
}

// ------------------------------------------------------------------------------
type tasksInsert struct {
	taskInserter
	tasks []Task
}

var _ MultiParams = tasksInsert{}

func (x tasksInsert) Append(args []interface{}, i int) []interface{} {
	t := &(x.tasks[i])
	return x.appendTask(args, t)
}

func (x tasksInsert) Len() int {
	return len(x.tasks)
}

// ------------------------------------------------------------------------------
func appendTaskSample(db *sql.DB) error {

	now := time.Now()
	t := &Task{
		Alias:     "alias3",
		Timestamp: now,
		Groups:    "/group2/subgroup3",
	}

	fmt.Printf("task: %+v\n", t)

	return appendTask(db, t)
}

func removeTasks(db *sql.DB) error {
	query := querier.deleteAllRecords(tableNameTasks)
	_, err := db.Exec(query)
	return err
}

func getTasks(db *sql.DB) ([]Task, error) {
	query := querier.selectAll(tableNameTasks)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tasksFromRows(rows)
}

func getTasksLoc(db *sql.DB, loc [2]int) ([]Task, error) {
	query, args := querier.selectRange(tableNameTasks, "task_id", loc)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tasksFromRows(rows)
}

func tasksFromRows(rows *sql.Rows) ([]Task, error) {
	var (
		t  Task
		ts []Task
	)
	for rows.Next() {
		err := rows.Scan(
			&(t.ID),
			&(t.Alias),
			&(t.Timestamp),
			&(t.Groups),
			&(t.Number),
			&(t.FieldU64),
		)
		if err != nil {
			return ts, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func numberOfTasks(db *sql.DB) (int, error) {
	return selectCountRows(db, tableNameTasks)
}

func selectCountRows(db *sql.DB, tableName string) (int, error) {

	query := querier.numberOfRecords(tableName)
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !(rows.Next()) {
		return 0, errors.New("there are no records")
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
