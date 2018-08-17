//package main
//
//import (
//	"database/sql"
//	"fmt"
//	//"time"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func main() {
//	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
//	checkErr(err)
//
//	//插入数据
//	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
//	checkErr(err)
//
//	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
//	checkErr(err)
//
//	id, err := res.LastInsertId()
//	checkErr(err)
//
//	fmt.Println(id)
//	//更新数据
//	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
//	checkErr(err)
//
//	res, err = stmt.Exec("astaxieupdate", id)
//	checkErr(err)
//
//	affect, err := res.RowsAffected()
//	checkErr(err)
//
//	fmt.Println(affect)
//
//	//查询数据
//	rows, err := db.Query("SELECT * FROM userinfo")
//	checkErr(err)
//
//	for rows.Next() {
//		var uid int
//		var username string
//		var department string
//		var created string
//		err = rows.Scan(&uid, &username, &department, &created)
//		checkErr(err)
//		fmt.Println(uid)
//		fmt.Println(username)
//		fmt.Println(department)
//		fmt.Println(created)
//	}
//
//	//删除数据
//	stmt, err = db.Prepare("delete from userinfo where uid=?")
//	checkErr(err)
//
//	res, err = stmt.Exec(id)
//	checkErr(err)
//
//	affect, err = res.RowsAffected()
//	checkErr(err)
//
//	fmt.Println(affect)
//
//	db.Close()
//
//}
//
//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}


package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
	"time"
)

type Userinfo struct {
	Uid     int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username    string
	Departname  string
	Created     time.Time
}

type User struct {
	Uid         int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Userinfo),new(User), new(Profile), new(Tag))
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
