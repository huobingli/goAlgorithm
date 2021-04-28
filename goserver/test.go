package main

import (
	"database/sql"
	"strconv"
	"time"
)

package main

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
"log"
"strconv"
"time"
)

func main() {
	mysql()
}
//结构体
type Job struct {
	db *sql.DB
	ch chan int
	total int
	n int
}


func mysql(){
	db, err := sql.Open("mysql", "root:123456@/test_big?charset=utf8")
	if err != nil {
		fmt.Println("访问数据库出错", err)
		return
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Second * 500) //设置连接超时500秒
	db.SetMaxOpenConns(100)                 //设置最大连接数

	total := 25000
	gonum:=400
	fmt.Println("====start=====")
	start := time.Now()
	// 测试插入数据库的功能,每次最多同步20个工作协程
	jobChan := make(chan Job, 20)
	go worker(jobChan)
	//统计使用次数
	ch := make(chan int,gonum)
	for n := 0; n < gonum; n++ {
		job:=Job{
			db:db,
			ch:ch,
			total:total,
			n:n,
		}
		jobChan <- job
	}
	ii:=0
	for{
		<-ch
		ii++
		if(ii>=gonum){
			break;
		}
	}

	end := time.Now()
	curr := end.Sub(start)
	fmt.Println("run time:", curr)
}

func worker(jobChan <-chan Job) {
	for job := range jobChan {
		go sqlExec(job)
	}
}

func sqlExec(job Job)  {
	buf:=make([]byte,0,job.total)
	buf=append(buf," insert into student(user_name) values "...)
	for i := 0; i < job.total; i++ {
		myid, _ := uuid.NewV4()
		userName := myid.String()
		if i == job.total-1 {
			buf=append(buf,"('" + userName + "');"...)
		} else {
			buf=append(buf,"('" + userName + "'),"...)
		}
	}
	ss:=string(buf)
	fmt.Println("第" + strconv.Itoa(job.n) + "次插入2.5万条数据！")
	_, err := job.db.Exec(ss)
	checkErr(err)
	fmt.Println("完成---" + strconv.Itoa(job.n) + "次插入2.5万条数据！")
	job.ch<- 1
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}