package db

import (
	"testing"
)

func TestDbPool(t *testing.T) {
	dbPool, err := NewDbPool("tester:3335688kingsoft@tcp(10.20.216.117:3306)/test?charset=utf8&", 10)
	conn, err := dbPool.Get()
	if err != nil {
		println(err)
	}

	println(dbPool.Free(), dbPool.Size())

	dbPool.Put(conn)

	println(dbPool.Free(), dbPool.Size())
}
