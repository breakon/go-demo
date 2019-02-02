package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 调用 db.Query 执行 SQL 语句, 此方法会返回一个 Rows 作为查询的结果
// 通过 rows.Next() 迭代查询数据.
// 通过 rows.Scan() 读取每一行的值
// 调用 db.Close() 关闭查询
const (
	userName = "root"
	password = "djl123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "mytest" //连接的数据库名字
)

/**
主函数
*/
func main() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	var name string
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := sql.Open("mysql", path) //用户：密码@主机/数据库
	checkErr(err)
	// err = db.QueryRow("select name from user where id = 1", 0).Scan(&name)
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)

	rows, err := db.Query("select * from `user` where id>1")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		fmt.Println(id, name, age)
	}
	fmt.Println(name)

	//Err返回可能的、在迭代时出现的错误。Err需在显式或隐式调用Close方法后调用。
	err = rows.Err()
	if err != nil {
		fmt.Println("other error:", err)
		return
	}
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		log.Fatal(err)
		panic(err)
	}
}

// create table userinfo (

// 	uid int(10) not null auto_increment,

// 	username varchar(64) null default null,

// 	departname varchar(64) null default null,

// 	create data null default null,

// 	primary key(uid))
