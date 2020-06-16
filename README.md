
# CPD in File Logger

## Package Logger 
To import and use this package do:
```
go get github.com/chapdast/Logger
```
and just add import on top on Go File.
```
import "github.com/chapdast/Logger"
```

    type Logger struct {...}
    func New(logRoot string, debug bool, daily bool) *Logger
---
    func (l *Logger) Custom(Tag string, msg ...interface{})
    func (l *Logger) Debug(msg ...interface{})
    func (l *Logger) Error(msg ...interface{})
    func (l *Logger) Info(msg ...interface{})
    func (l *Logger) Warn(msg ...interface{})


### Examples

    l := 
    Logger.New("logs/","chapdast", true, true)
	
    l.Info("This","is", "a info Message")
	l.Debug("This is A Debug Message")
	l.Warn("This is A Warn Message")
	l.Error("This is A Error Message")
	l.Custom("cpd", "This is a Custom Tag")

this will log following to console and to file:

    2020/June/16 9:34:25    [info]  [[This is a info Message]]
    2020/June/16 9:34:25    [debug] [[This is A Debug Message]]
    2020/June/16 9:34:25    [warn]  [[This is A Warn Message]]
    2020/June/16 9:34:25    [error] [[This is A Error Message]]
    2020/June/16 9:34:25    [cpd]   [[This is a Custom Tag]]