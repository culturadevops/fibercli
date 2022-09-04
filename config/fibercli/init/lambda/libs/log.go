package libs

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	LogInfo  *log.Logger
	LogError *log.Logger
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
func Write(nivel int, typems string, format string, a ...interface{}) string {
	_, file, line, _ := runtime.Caller(nivel)
	spliting := strings.Split(file, "/")
	x := len(spliting)
	infofile := spliting[x-1]

	info := fmt.Sprintf(format, a...)
	text := fmt.Sprintf("%s:%d %v", infofile, line, info)
	fmt.Println(typems, text)
	return text

}
func Error(nivel int, format string, a ...interface{}) {
	texto := Write(nivel, "【Error】:", format, a...)
	color.Red(texto)
	LogError.Println(texto)
	LogInfo.Println(texto)
}
func Info(nivel int, format string, a ...interface{}) {
	texto := Write(nivel, "【Info】:", format, a...)
	color.Yellow(texto)
	LogInfo.Println(texto)
}
func init() {
	os.Mkdir("logs", 0755)
	os.Mkdir("logs/error", 0755)
	logFile, err := os.OpenFile("./logs/run_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln("open log file failed", err)
	}
	logFileError, err1 := os.OpenFile("./logs/error/run_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err1 != nil {
		log.Fatalln("open log file 'error' failed ", err)
	}
	//日志
	LogInfo = log.New(io.MultiWriter(logFile), "【Info】:", log.Ldate|log.Ltime)        //LogInfo.Println(1, 2, 3)
	LogError = log.New(io.MultiWriter(logFileError), "【Error】:", log.Ldate|log.Ltime) //LogError.Println(4, 5, 6)

}
