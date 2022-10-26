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
	_, filename1, _, _ := runtime.Caller(1)
	_, filename2, _, _ := runtime.Caller(0)
	_, filename3, _, _ := runtime.Caller(2)
	_, filename4, _, _ := runtime.Caller(3)
	_, filename5, _, _ := runtime.Caller(4)
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename1)
	fmt.Println(filename2)
	fmt.Println(filename3)
	fmt.Println(filename4)
	fmt.Println(filename5)
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
