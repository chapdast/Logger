package Logger

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Logger struct {
	logPath string
	debug bool
	dailySepration bool
	lock *sync.Mutex
	extention string
}

// Create And Return a New Logger
func New(logRoot string, debug bool, daily bool) *Logger{
	return &Logger{
		logPath: logRoot,
		debug: debug,
		dailySepration: daily,
		lock:      &sync.Mutex{},
		extention: ".cpd.log",
	}
}

func(l *Logger) log(tag string, msg ...interface{}) error{
	l.lock.Lock()
	defer l.lock.Unlock()
	if _, err := os.Stat(l.logPath); os.IsNotExist(err) {
		err := os.Mkdir(l.logPath, 0755)
		if err != nil {
			return err
		}
	}
	t := time.Now()
	var name = "log_"
	if l.dailySepration {
		name = strconv.Itoa(t.Year()) + "_" + t.Month().String() + "_" + strconv.Itoa(t.Day()) + l.extention
	} else {
		name = strconv.Itoa(t.Year()) + "_" + t.Month().String() + "_" + strconv.Itoa(t.Day()) + "_" + strconv.Itoa(t.Hour()) + l.extention
	}
	var file *os.File

	if _, err := os.Stat(l.logPath + name); os.IsNotExist(err) {
		file, err = os.Create(l.logPath + name)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(l.logPath+name, syscall.O_CREAT|syscall.O_APPEND|syscall.O_WRONLY, 0755)
		if err != nil {
			return err
		}
	}
	defer file.Close()
	logRepo :=fmt.Sprint(t.Year(), "/", t.Month().String(), "/", t.Day(), " ", t.Hour(), ":", t.Minute(), ":", t.Second(), "\t","[" + tag + "]\t", msg, "\n")
	file.WriteString(logRepo)
	if l.debug{
		fmt.Print(logRepo);
	}
	



	return nil
}

// Log Messages to LogFile under Info tag
func (l *Logger) Info(msg ...interface{}){
	if err:= l.log("info", msg); err!=nil{
		fmt.Println("ERROR", err);
	}
	
}

// Log Messages to LogFile under Debug tag
func (l *Logger) Debug(msg ...interface{}){
	if err:= l.log("debug", msg); err!=nil{
		fmt.Println("ERROR", err);
	}
}

// Log Messages to LogFile under Warn tag
func (l *Logger) Warn(msg ...interface{}){
	if err:= l.log("warn", msg); err!=nil{
		fmt.Println("ERROR", err);
	}
	
}


// Log Messages to LogFile under Error tag
func (l *Logger) Error(msg ...interface{}){
	if err:= l.log("error", msg); err!=nil{
		fmt.Println("ERROR", err);
	}
}

// Log Messages to LogFile under custom defined tag
func (l *Logger) Custom(Tag string, msg ...interface{}){
	if err:= l.log(Tag, msg); err!=nil{
		fmt.Println("ERROR", err);
	}
}
