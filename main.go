package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "vinam"
	dbPass := "vinam"
	dbName := "goblog"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	csvFile, _ := os.Open("people.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	db := dbConn()
	fmt.Println(time.Now())
	flag := false
	for {
		fmt.Println(flag)
		line, error := reader.Read()
		if len(line) > 0 && flag == true {
			insForm, err := db.Prepare("INSERT INTO names(firstname, lastname) VALUES(?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(line[0], line[1])
		}
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		flag = true
	}
	fmt.Println(time.Now())
}
