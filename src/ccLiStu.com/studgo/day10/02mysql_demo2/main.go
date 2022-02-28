package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
// go连接MySQL示例
var db *sql.DB // 是一个连接对象

func initDB() (err error)  {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/user"
	//连接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                              //dsn 格式不正确的时候会报错
		return
	}
	err = db.Ping()  //尝试连接数据库
	if err != nil {
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	return
}

type user struct {
	id int
	name string
	age int
}

//查询
func queryOne(id int)  {
	var u1 user
	//1. 写查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id=?;`
	//2.执行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age) //  从连接池里拿一个连接出来去数据库查询记录

	//3.拿到结果
	//rowObj //必须对rowObj对象
	//4.打印结果
	fmt.Printf("u1:%#v\n", u1)
}

func insert()  {
	
}
func main()  {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功！")

}

//

//