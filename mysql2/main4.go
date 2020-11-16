package main

//https://tutorialedge.net/golang/golang-mysql-tutorial/

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Db *sql.DB
}

func (dbw DbWorker) getTbCount(sqls []string, c chan int) {
	defer close(c)
	for _, sqlStr := range sqls {
		cnt := 0
		fmt.Println(sqlStr)
		dbw.Db.QueryRow(sqlStr).Scan(&cnt)
		c <- cnt
	}
}

func main() {
	dbw := DbWorker{}
	dbw.Db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cms1")
	defer dbw.Db.Close()

	sqls := []string{
		"select count(*) as cnt from article",
		"select count(*) as cnt from vbaidu",
		"select count(*) as cnt from vbaidu_xiaopin",
	}
	result := []int{}
	c := make(chan int)
	go dbw.getTbCount(sqls, c)
	for res := range c {
		result = append(result, res)
	}
	fmt.Println(result)
}
