package db

import (
	"testing"
)

func TestDbPool(t *testing.T) {
	dbPool := NewDbPool("tester:3335688kingsoft@tcp(10.20.216.117:3306)/test?charset=utf8&", 10)
	conn, err := dbPool.Get()
	if err != nil {
		println(err)
	}

}
