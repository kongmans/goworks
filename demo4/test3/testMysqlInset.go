package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//打开连接

	db, err := sql.Open("mysql", "root:root123456@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	//预处理ｓｑｌ
	stmt, err := db.Prepare("insert into people values (default,?,?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	r, err := stmt.Exec("普金", "俄罗斯")
	if err != nil {
		fmt.Println("写入失败！")
		return
	}

	//获取结果
	count, _ := r.RowsAffected()
	if count > 0 {
		fmt.Println("新增成功！")
	} else {
		fmt.Println("新增失败！")
	}

}
