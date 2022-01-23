package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"userid"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test01?charset=utf8mb4")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	// defer Db.Close() // 注意这行代码要写在上面err判断的下面
}

func main() {

	defer Db.Close()
	/*插入*/
	// r, err_insert := Db.Exec("insert into person(userid , username, sex, email)values(? ,?, ?, ?)", 12, "stu008", "女", "stu08@qq.com")
	r, err_insert := Db.Exec("insert into student(class_id , name)values(?, ?)", 1, "stu008")
	if err_insert != nil {
		fmt.Println("insert failed, ", err_insert)
		return
	}
	fmt.Println("r:", r)
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert success! ", id)
	/*查询*/
	// var p []Person
	// // err_select := Db.Select(&p, "select username ,sex, email from person where userid=?", 1)
	// err_select := Db.Select(&p, "select userid,username ,sex, email from person ")
	// if err_select != nil {
	// 	fmt.Println("select failed ", err_select)
	// 	return
	// }
	// fmt.Println("select方法一：", p)

	// p := Person{}
	// rows, err_select := Db.Query("select username ,sex, email from person where userid=?", 1)
	// if err_select != nil {
	// 	fmt.Println("select failed ", err_select)
	// 	return
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&p.Username, &p.Sex, &p.Email)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println("select方法二", &p.Username, &p.Sex, &p.Email)
	// }

}
