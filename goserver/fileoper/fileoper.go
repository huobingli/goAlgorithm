package fileoper

import "database/sql"

// https://www.jianshu.com/p/9b5cd762e256
// https://blog.csdn.net/xianchanghuang/article/details/84766135
// 解析文件内容 存到数据库
import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


//arr [4]string = {"17_SXXX.txt", "33_SXXX.txt", "17_XY", "33_XY"}
//file_array = []
//file_path = "/file"

type StockData struct {
	db *sql.DB
	ch chan int
	stock 	string
	SXXX	float64
	XY		float64
}

type DataItem struct {
	stock 	string
	SXXX	float64
	XY		float64
}

func loadfile() {
	
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sqlExec(data StockData)  {
	//strsql := "INSERT INTO test(n1, n2, n3) VALUES ?"
	//strsql = append(strsql," insert into stock(code, sxxx, xy) values ", data.stock, data.SXXX, data.XY)
	strsql := fmt.Sprintf("insert into stock(code, sxxx, xy) values ", data.stock, data.SXXX, data.XY)

	ss:=string(buf)
	fmt.Println("第" + strconv.Itoa(data.n) + "次插入2.5万条数据！")
	_, err := data.db.Exec(ss)
	checkErr(err)
	fmt.Println("完成---" + strconv.Itoa(data.n) + "次插入2.5万条数据！")
	data.ch<- 1
}

func worker(dataChan <-chan StockData) {
	for data := range dataChan {
		go sqlExec(data)
	}
}

func adddata() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", username, password, ipaddrees, port, dbname, charset)
	db, err := sqlx.Open("mysql", conn)
	//database, err := sqlx.Open("mysql", "root:1111@tcp(47.114.171.118:3306)/stock")
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}

	defer db.Close()
	db.SetConnMaxLifetime(time.Second * 500) //设置连接超时500秒
	db.SetMaxOpenConns(100)                 //设置最大连接数

	total := 25000
	gonum:=400
	start := time.Now()
	// 测试插入数据库的功能,每次最多同步20个工作协程
	stock := make(chan StockData, 20)
	go worker(stock)
	//统计使用次数
	ch := make(chan int,gonum)
	for n := 0; n < gonum; n++ {
		job:=StockData{
			db:db,
			ch:ch,
			stock:code,
			SXXX:sxxx,
			XY:xy,
		}
		stock <- job
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

	//return database
}

func Openfile() {
	dataArray := map[uint]DataItem {}
 	file, err := os.Open("D:\\GitProject\\goAlgorithm\\goserver\\fileoper\\17_XY.txt")
    if err != nil {
        fmt.Println("文件打开失败 = ", err)
    }
    //及时关闭 file 句柄，否则会有内存泄漏
    defer file.Close()
    //创建一个 *Reader ， 是带缓冲的
    reader := bufio.NewReader(file)
    i := 0
    for {
        str, err := reader.ReadString('\n') //读到一个换行就结束
        if err == io.EOF {                  //io.EOF 表示文件的末尾
            break
        }
		//str = str[1:]
		arr := strings.Split(str,";")
		if len(arr) > 3 {
			fmt.Print(str)
		}

		v2, err := strconv.ParseFloat(arr[2], 64)
		v3, err := strconv.ParseFloat(arr[3], 64)
		dataArray[i] = DataItem{arr[0], v2, v3}
		i++
		// add dataArray
    }
    fmt.Println("文件读取结束...")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", username, password, ipaddrees, port, dbname, charset)
	db, err := sqlx.Open("mysql", conn)
	//database, err := sqlx.Open("mysql", "root:1111@tcp(47.114.171.118:3306)/stock")
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}

	for data := range dataArray {
		db.Exec("insert into stock(code, sxxx, xy) values ", data.stock, data.SXXX, data.XY)
	}

}