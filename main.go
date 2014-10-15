package main

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/citysir/golib/db"
)

func main() {
	fmt.Println("Hello, golib")
	testDbPool()
}

func testDbPool() {
	dbPool, err := db.NewDbPool("tester:3335688kingsoft@tcp(10.20.216.117:3306)/test?charset=utf8&", 10)
	conn, err := dbPool.Get()
	if err != nil {
		println(err)
	}

	println(dbPool.Free(), dbPool.Size())

	dbPool.Put(conn)

	println(dbPool.Free(), dbPool.Size())
}
