package main

import (
	"fmt"
	"math/rand"
	"myGoPractices/gitrepo/exceldemo/model"
	"sync"
	"time"
)

const (
	mainDbUserName  = "root"
	mainDbUserPwd   = "1058471169"
	mainDbHost      = "127.0.0.1:3306"
	mainDbName      = "ncu"
	mainDbCfgName   = "main"
	mainDbMaxIdle   = 10
	mainDbMaxActive = 20

	numberGoroutines = 10 // 要使用的 goroutine 的数量
)

type UserInfo struct {
	Name      string
	Age       int
	Gender    int
	RandomNum int
}

var (
	wg sync.WaitGroup
	// mutex sync.Mutex
)

func main() {
	begin := time.Now()
	wg.Add(100)

	model.InitMysql(mainDbUserName, mainDbUserPwd, mainDbHost, mainDbName, mainDbCfgName, mainDbMaxIdle, mainDbMaxActive)

	users := make([]*UserInfo, 0, 1000)

	num := 100000
	for i := 0; i < 10; i++ {
		value := i
		go func() {
			for j := value; j < 1000; j = j + 10 {

				for k := j * 1000; k < j*1000+1000; k++ {
					users = append(users, generate())
				}
				insert(users)
				users = make([]*UserInfo, 0, 1000)
			}
			/*
				for i := value; i <= num; i++ {
					users = append(users, generate())
					if i%2000 == 0 {
						insert(users)
						users = make([]*UserInfo, 0)
					}
				}
			*/
		}()
	}

	fmt.Println("\n数据生成耗时:", time.Since(begin))

	insBegin := time.Now()
	fmt.Printf("准备插入 %d 条数据 \n", num*numberGoroutines)

	fmt.Println("数据插入耗时:", time.Since(insBegin))

	fmt.Println("程序运行总时间:", time.Since(begin))
	wg.Wait()
}

func generate() *UserInfo {
	userInfo := new(UserInfo)

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(10) + 1
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	name := string(result)
	age := r.Intn(100)
	gender := r.Intn(2)
	randomNum := r.Intn(1000)

	userInfo.Name = name
	userInfo.Age = age
	userInfo.Gender = gender
	userInfo.RandomNum = randomNum
	return userInfo
}

func insert(users []*UserInfo) error {
	defer wg.Done()
	db, err := model.GetMysqlCon("main")

	if err != nil {
		fmt.Println("\n 获取数据库连接失败:", err)
		return err
	}
	sqlExec := model.StringBuilder{}
	sqlExec.Append("INSERT INTO user_info")
	sqlExec.Append("(`name`,`age`,`gender`,`random_num`) VALUES ")
	params := make([]interface{}, 0, 10*len(users))
	n := 1
	m := len(users)

	for _, v := range users {
		sqlExec.Append("(?,?,?,?)")
		if n < m {
			sqlExec.Append(",")
		}

		params = append(params, v.Name, v.Age, v.Gender, v.RandomNum)
		n++
	}

	result, err := db.Exec(sqlExec.ToString(), params...)
	if err != nil {
		fmt.Println("插入信息失败:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("获取受影响的行数失败:", err)
		return err
	}
	fmt.Println("受影响的行数为:", rowsAffected)

	return nil
}
