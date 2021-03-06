package db

import (
	"database/sql"
	"errors"
	"time"
)

type DbPool struct {
	connectString string
	poolSize      int
	poolConns     chan *DbPoolConn
}

type DbPoolConn struct {
	conn    *sql.DB
	getTime time.Time
	putTime time.Time
}

var PingIntervalDuration = 900 * time.Second //15分钟ping一次

func NewDbPool(connectString string, poolSize int) (*DbPool, error) {
	dbPool := &DbPool{connectString: connectString, poolSize: poolSize}
	dbPool.poolConns = make(chan *DbPoolConn, dbPool.poolSize)
	flag := make(chan bool, dbPool.poolSize)
	go func() {
		for i := 0; i < dbPool.poolSize; i++ {
			conn := dbPool.newConn()
			dbPool.Put(conn)
			flag <- true
		}
	}()
	for i := 0; i < dbPool.poolSize; i++ {
		<-flag
	}
	return dbPool, nil
}

func (this *DbPool) Get() (*sql.DB, error) {
	select { //判断是否能在3秒内获取连接，如果不能就报错
	case poolConn, ok := <-this.poolConns: //读取通道里的数据库连接，如果读不到就返回报错
		{
			if ok {
				err := poolConn.ping()
				if err != nil {
					poolConn.conn = this.newConn() //重新创建连接
				}
				poolConn.getTime = time.Now()
				return poolConn.conn, nil
			} else {
				return nil, errors.New("数据库连接获取异常, 可能已经被关闭")
			}
		}
	case <-time.After(time.Second * 5): //如果被阻塞5秒仍没有获取到连接，则就返回错误
		return nil, errors.New("获取数据库连接超时")
	}
}

func (this *DbPool) Put(conn *sql.DB) {
	if len(this.poolConns) == this.poolSize {
		conn.Close()
		return
	}
	this.poolConns <- &DbPoolConn{conn: conn, putTime: time.Now()}
}

func (this *DbPool) Free() int {
	return len(this.poolConns)
}

func (this *DbPool) Size() int {
	return this.poolSize
}

func (this *DbPool) newConn() *sql.DB {
	conn, err := sql.Open("mysql", this.connectString)
	if err != nil {
		panic(err)
	}
	return conn
}

func (this *DbPoolConn) ping() error {
	if time.Now().After(this.putTime.Add(PingIntervalDuration)) {
		return this.conn.Ping()
	}
	return nil
}
