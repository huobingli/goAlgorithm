package fileoper

// 解析文件内容 存到数据库
import (
	"bufio"
	"fmt"
	"io"
	"os"
)


//arr [4]string = {"17_SXXX.txt", "33_SXXX.txt", "17_XY", "33_XY"}
//file_array = []
//file_path = "/file"

type stockdata struct {
	stock 	string
	SXXX	float64
	XY		float64
}

func loadfile() {
	
}

func Openfile() {
	file, err := os.Open("17_XY.txt")
    if err != nil {
        fmt.Println("文件打开失败 = ", err)
    }
    //及时关闭 file 句柄，否则会有内存泄漏
    defer file.Close()
    //创建一个 *Reader ， 是带缓冲的
    reader := bufio.NewReader(file)
    // for {
        str, err := reader.ReadString('\n') //读到一个换行就结束
        if err == io.EOF {                  //io.EOF 表示文件的末尾
            //break
        }
        fmt.Print(str)
    // }
    fmt.Println("文件读取结束...")
}