package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	db := openDB()
	query(db)
	// queryRow(db)
	// insert(db)

}

func insert(db *sql.DB) {
	res, err := db.Exec("insert into user(name,age) values(?,?)", "xiaoxiao", 16)
	if err != nil {
		log.Fatalln(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("受影响的行数:", count)
}

func queryRow(db *sql.DB) {
	users := make([]*User, 0)

	stmt, err := db.Prepare("select id,name,age from user where id = ?")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	rows := stmt.QueryRow(2)

	var user User
	err = rows.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		log.Fatalln(err)
	}
	users = append(users, &user)

	for _, u := range users {
		log.Printf("%#v \n", u)
	}
}

func query(db *sql.DB) {
	users := make([]*User, 0)

	stmt, err := db.Prepare("select id,name,age from user where id = ?")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(2)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, &user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, u := range users {
		log.Printf("%#v \n", u)
	}
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1058471169@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("成功打开了数据库")
	return db
}
