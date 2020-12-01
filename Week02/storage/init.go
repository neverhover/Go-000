package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var dbConn *sql.DB

func GetDB() *sql.DB {
	return dbConn
}

func CloseDB() {
	dbConn.Close()
}

func InitDB() {
	dsn := "./dist/week02.sqlite3"
	os.Remove(dsn)
	dbIns, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
	dbIns.Ping()
	dbConn = dbIns
	mockInit()
	log.Printf("DB stats:%+v\n", dbIns.Stats())
	// Any other operations
}

// mockInit init mock data
func mockInit() {
	sqlStmt := `
	create table user (id text not null primary key, name text);
	`
	db := GetDB()
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into user(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 1; i < 100; i++ {
		_, err = stmt.Exec(fmt.Sprintf("%d", 1000+i),
			fmt.Sprintf("Tony-%04d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

