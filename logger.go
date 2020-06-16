package Logger

import "sync"

type Logger struct {
	LogPath string
	Debug bool
	DailySepration bool
	ML *sync.Mutex
	extention string
}

func New(LogRootPath string, debug bool, daily bool) *Logger{
	
	return &Logger{
		LogPath: dump,
		Debug: debug,
		DailySepration: daily,
		ML:      &sync.Mutex{},
		extention: ".cpd.log"
	}
}

func(l *Logger) log(tag string, msg interface{}...){
	if _, err := os.Stat(l.LogPath); os.IsNotExist(err) {
		err := os.Mkdir(l.LogPath, 0755)
		if err != nil {
			return err
		}
	}
	t := time.Now()
	var name = "log_"
	if l.DailySepration {
		name = strconv.Itoa(t.Year()) + "_" + t.Month().String() + "_" + strconv.Itoa(t.Day()) + l.extention
	} else {
		name = strconv.Itoa(t.Year()) + "_" + t.Month().String() + "_" + strconv.Itoa(t.Day()) + "_" + strconv.Itoa(t.Hour()) + l.extention
	}
	var file *os.File

	if _, err := os.Stat(l.LogPath + name); os.IsNotExist(err) {
		file, err = os.Create(l.LogPath + name)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(l.LogPath+name, syscall.O_CREAT|syscall.O_APPEND|syscall.O_WRONLY, 0755)
		if err != nil {
			return err
		}
	}
	defer file.Close()
	file.WriteString(fmt.Sprint(t.Year(), "/", t.Month().String(), "/", t.Day(), " ", t.Hour(), ":", t.Minute(), ":", t.Second(), "\t",
		"[" + tag + "]", msg, "\n"))
	l.ML.Unlock()
	return nil


}
func (l *Logger) Info(msg interface{}...){
	l.log("info", msg)
}

func (l *Logger) Debug(msg interface{}...){
	l.log("debug", msg)
}

func (l *Logger) Warn(msg interface{}...){
	l.log("debug", msg)
}

func (l *Logger) Error(msg interface{}...){
	l.log("error", msg)
}

func (l *Logger) Custom(Tag string, msg interface{}...){
	l.log(tag, msg);
}
