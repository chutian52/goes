package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main(){

	//fmt.Println("Hello world")

	db, err := sql.Open("mysql", "root:123456@tcp(mysql.dev.svc.cluster.local:3307)/starter_go?charset=utf8mb4")
	if err != nil {
		log.Fatalln("创建mysql连接失败", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("hello world")


}
