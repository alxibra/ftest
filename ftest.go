package ftest

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
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
		qs := splitQuery(string(c))
		for _, q := range qs {
			_, err = db.Exec(q)
			if err != nil {
				panic(err)
			}
		}
	}
	return fn
}

func rootPath() string {
	_, filename, _, _ := runtime.Caller(6)
	fmt.Println(filename)
	return path.Join(path.Dir(filename), "..")
}

func createTable(fn string, db *sql.DB) {
	fmt.Println("    *** Create table")
	p := fmt.Sprintf("%s/tests/setup/%s.sql", rootPath(), fn)
	c, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	qs := splitQuery(string(c))
	for _, q := range qs {
		_, err = db.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func insertSeed(fn string, db *sql.DB) {
	fmt.Println("    *** Insert seed")
	p := fmt.Sprintf("%s/tests/seed/%s.sql", rootPath(), fn)
	c, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	qs := splitQuery(string(c))
	for _, q := range qs {
		_, err = db.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func splitQuery(str string) []string {
	cleanup := strings.TrimRight(str, "\t\r\n")
	qs := strings.Split(cleanup, ";")
	qs = qs[:len(qs)-1]
	for i, q := range qs {
		qs[i] = fmt.Sprintf("%s;", q)
	}
	return qs
}
