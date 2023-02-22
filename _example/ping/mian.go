package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/zhangguojvn/gobci"
)

func main() {
	db, err := sql.Open(
		"gobci",
		fmt.Sprintf(
			"%s/%s@%s:%s/%s",
			os.Getenv("USERNAME"),
			os.Getenv("PASSWORD"),
			os.Getenv("OBIP"),
			os.Getenv("OBPORT"),
			os.Getenv("OBDATABASE"),
		),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic("Ping error:" + err.Error())
	}
	row, err := db.Query("SELECT * FROM DUAL")
	if err != nil {
		panic("Ping error:" + err.Error())
	}
	defer row.Close()
	for row.Next() {
		var reuslt interface{}
		err = row.Scan(&reuslt)
		if err != nil {
			panic("Ping error:" + err.Error())
		}
		fmt.Print(reuslt)
	}
}
