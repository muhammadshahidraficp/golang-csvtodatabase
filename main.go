package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

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
	flag := false
	new := 0
	count := 0
	for {
		line, error := reader.Read()
		if len(line) > 0 && flag == true {
			selDB, err := db.Query("SELECT * FROM name WHERE firstname=? and lastname=?", line[0], line[1])
			for selDB.Next() {
				count++
			}
			if err != nil {
				panic(err.Error())
			}

		}
		if len(line) > 0 && flag == true {
			insForm, err := db.Prepare("INSERT INTO name(firstname, lastname) VALUES(?,?)")
			//fmt.Println(line[0], line[1])
			new++
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
	defer db.Close()
	fmt.Println("Match from database ", count)
	fmt.Println("Inserted ", new)
}
