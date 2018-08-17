package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "root@/test?charset=utf8")
	checkErr(err)
	//stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?")
	//checkErr(err)
	//res, err := stmt.Exec("dalong", "研发", "2012-09-12")
	//checkErr(err)
	//id, err := res.LastInsertId()
	//checkErr(err)
	//fmt.Println(id)

	//stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//res, err := stmt.Exec("wenlong", "1")
	//rownum, err := res.RowsAffected()
	//fmt.Println(rownum)

	//rows, err := db.Query("select * from userinfo")
	//checkErr(err)
	//
	//for rows.Next(){
	//	var uid int
	//	var username string
	//	var departname string
	//	var created string
	//	err = rows.Scan(&uid, &username, &departname, &created)
	//	checkErr(err)
	//	fmt.Println(uid)
	//	fmt.Println(username)
	//	fmt.Println(departname)
	//	fmt.Println(created)
	//}

	stmt, err := db.Prepare("delete from userinfo WHERE uid=?")
	checkErr(err)
	res, err := stmt.Exec("1")
	rownum, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(rownum)


}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}