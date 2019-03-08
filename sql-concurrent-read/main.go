package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	db, err := sql.Open("mysql", "test:test@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)
	//createTable(db)
	insert(db)

	mesOut := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			var id int
			//mesIn := time.Now()
			err = db.QueryRow("SELECT PersonID FROM Persons;").Scan(&id)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			//fmt.Printf("Time and number: %d -> %s\n", id, time.Since(mesIn))
			wg.Done()
		}(wg)
	}
	wg.Wait()
	fmt.Printf("Time outer: %s\n", time.Since(mesOut))
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func createTable(db *sql.DB) {
	// Prepare statement for inserting data
	stmtCreate, err := db.Prepare(`CREATE TABLE Persons (PersonID int);`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtCreate.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtCreate.Exec()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func insert(db *sql.DB) {
	// Prepare statement for inserting data

	stmtIns, err := db.Prepare("INSERT INTO Persons VALUES(?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(randomInt(1, 1222234234))
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
