package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	c := ConfigDB{
		Host:     "localhost",
		Port:     5432,
		Username: "robinson",
		Password: "4V32Am?dLak@Z*z",
		DBName:   "testdb",
	}

	source := makeSourceName(c)
	db, err := sql.Open("postgres", source)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	err = createTableTasks(db)
	checkError(err)

	// err = dropTableTasks(db)
	// checkError(err)

	//checkError(removeTasks(db))

	// err = appendRandTasks(db, 1)
	// checkError(err)

	// err = testAppendMulti(db)
	// checkError(err)

	ts, err := getTasks(db)
	checkError(err)
	for _, t := range ts {
		fmt.Printf("%+v\n", t)
	}

	// ts, err := getTasksLoc(db, [2]int{0, 10})
	// checkError(err)
	// for _, t := range ts {
	// 	fmt.Printf("%+v\n", t)
	// }

	count, err := numberOfTasks(db)
	checkError(err)
	fmt.Println("numberOfTasks:", count)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ConfigDB struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func makeSourceName(c ConfigDB) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Password, c.DBName)
}
