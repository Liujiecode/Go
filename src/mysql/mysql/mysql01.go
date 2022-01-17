package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Main_mysql01() {
	data, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test01")
	if err != nil {
		fmt.Println("连接出错了！")
	} else {
		fmt.Println(data)
	}
	fmt.Println("ja")
}
