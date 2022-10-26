package ftest

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

// Setup test, create table, insert seed and return func to tear down table
func Setup(filename string, db *sql.DB) func() {
	fmt.Println("*** SETUP")
	createTable(filename, db)
	insertSeed(filename, db)
	fn := func() {
		fmt.Println("***TEARDOWN")
		p := fmt.Sprintf("%s/tests/tear/%s.sql", rootPath(), filename)
		c, err := ioutil.ReadFile(p)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(string(c))
		if err != nil {
			panic(err)
		}
	}
	return fn
}

func rootPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "..")
}

func createTable(fn string, db *sql.DB) {
	fmt.Println("    *** Create table")
	p := fmt.Sprintf("%s/tests/setup/%s.sql", rootPath(), fn)
	c, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(c))
	if err != nil {
		panic(err)
	}
}

func insertSeed(fn string, db *sql.DB) {
	fmt.Println("    *** Insert seed")
	p := fmt.Sprintf("%s/tests/seed/%s.sql", rootPath(), fn)
	c, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(c))
	if err != nil {
		panic(err)
	}
}
