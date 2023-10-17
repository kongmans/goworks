package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Age  int
}

var db *sql.DB

// 初始化数据库连接
func initDB() (err error) {
	// sql.Open参数 第一个参数为数据库类型，例如：mysql、oracle等，第二个参数为指定数据源。格式：数据库账号:密码@连接方式(数据库所在的ip地址)/数据库名
	db, err = sql.Open("mysql", "root:root123456@tcp(127.0.0.1)/test")
	if err != nil {
		fmt.Printf("db open err : %s\n", err)
		return err
	}
	// 检测数据源配置是否合法使用db.Ping()方法
	err = db.Ping()
	if err != nil {
		fmt.Printf("db ping err :%s\n", err)
		return err
	}
	return nil
}

// 查询一条数据
func findOne(name string) {
	var u User
	err := db.QueryRow("select id,name,age from `user` where name= ?", name).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("findOne data failed err:%s\n", err)
		return
	}
	fmt.Printf("findOne User info %v\n", u)
}

// 查询多条数据
func findsData() {
	var s []User
	rows, err := db.Query("select id,name,age from `user`")
	if err != nil {
		fmt.Printf("findsData failed err:%s\n", err)
		return
	}
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("findsData scan failed err:%s\n", err)
			return
		}
		fmt.Printf("findsData User info %v\n", u)
		s = append(s, u)
	}
	fmt.Printf("%v\n", s)
}

// 插入一条数据
func insertData(name string, age int) {
	exec, err := db.Exec("insert into user(name,age) values (?,?)", name, age)
	if err != nil {
		fmt.Printf("exec insert failed err:%s\n", err)
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Printf("exec insert failed err:%s\n", err)
		return
	}
	fmt.Printf("insert data id is : %d\n", id)
}

// 更新数据
func updateData(id, age int, name string) {
	ret, err := db.Exec("update user set name = ?,age = ? where id = ?", name, age, id)
	if err != nil {
		fmt.Printf("update failed err:%s\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update RowsAffected failed err:%s\n", err)
		return
	}
	fmt.Printf("update success rows:%d\n", affected)
}

// 删除数据
func delData(id int) {
	ret, err := db.Exec("delete from user where id = ?", id)
	if err != nil {
		fmt.Printf("del failed err:%s\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed err:%s\n", err)
		return
	}
	fmt.Printf("delete success rows:%d\n", affected)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connection mysql db failed:%s", err)
		return
	}
	fmt.Println("connection mysql db success")
	fmt.Println("init db")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	fmt.Println("start insert")
	insertData("张三", 10)
	insertData("李四", 3)
	insertData("王五", 17)
	fmt.Println("insert end")
	fmt.Println("-----------------------------------")
	fmt.Println("start get")
	findOne("张三")
	fmt.Println("get end")
	fmt.Println("-----------------------------------")
	updateData(43, 100, "3333")
	fmt.Println("-----------------------------------")
	delData(45)
	fmt.Println("-----------------------------------")
	findsData()

}
