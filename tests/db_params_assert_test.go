package tests

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func getConn() (*sql.DB, error) {
	userName := "root"
	userPwd := "1058471169"
	host := "127.0.0.1:3306"
	dbName := ""
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local", userName, userPwd, host, dbName))
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
func BenchmarkDbWithAssert(B *testing.B) {
	conn, _ := getConn()
	tmp := make(map[string]interface{})
	stuNo1 := "6103214026"
	tmp["no1"] = stuNo1

	for i := 0; i < B.N; i++ {
		params := make([]interface{}, 0)
		params = append(params, tmp["no1"].(string))
		conn.Query("select student_no,stu_name,birthday from student_info where student in (?) ", params)
	}
}
func BenchmarkDbWithoutAssert(B *testing.B) {
	conn, _ := getConn()
	tmp := make(map[string]interface{})
	stuNo1 := "6103214026"
	tmp["no1"] = stuNo1

	for i := 0; i < B.N; i++ {
		params := make([]interface{}, 0)
		params = append(params, tmp["no1"])
		conn.Query("select student_no,stu_name,birthday from student_info where student in (?) ", params)
	}
}
