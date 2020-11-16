package main

//https://tutorialedge.net/golang/golang-mysql-tutorial/

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Db     *sql.DB
	Result []int
}

func main() {
	dbw := DbWorker{}
	dbw.Db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cms1")
	defer dbw.Db.Close()

	var wg sync.WaitGroup
	sqls := []string{
		"select count(*) as cnt from article",
		"select count(*) as cnt from vbaidu",
		"select count(*) as cnt from vbaidu_xiaopin",
		"select count(*) as cnt from vbaidu_xiaopin",
		"select count(*) as cnt from vbaidu_xiaopin",
		"select count(*) as cnt from vbaidu_xiaopin",
	}
	dbw.Result = make([]int, len(sqls))
	for idx, sqlStr := range sqls {
		wg.Add(1)
		go func(sqlStr string, idx int, dbw *DbWorker) {
			defer wg.Done()
			cnt := 0
			dbw.Db.QueryRow(sqlStr).Scan(&cnt)
			// fmt.Println(sqlStr, cnt)
			dbw.Result[idx] = cnt
		}(sqlStr, idx, &dbw)
	}

	wg.Wait()
	fmt.Println(dbw.Result)
}
