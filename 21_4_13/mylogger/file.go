package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel //日志等级
	filePath    string   //日志文件路径
	fileName    string   //日志文件名
	fileobj     *os.File
	errFileobj  *os.File
	maxFileSize int64 //最大文件个数
}

//构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {

	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	f := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f.initFile() //按文件路径和名字将文件打开
	if err != nil {
		panic(err)
	}
	return f

}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}

	//日志文件都已经打开了
	f.fileobj = fileObj
	f.errFileobj = errFileObj
	return nil
}

func (f *FileLogger) enable(logLevel LogLevel) bool {

	return logLevel >= f.Level
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}

	fmt.Printf("get max size :%d file real size:%d\n", f.maxFileSize, fileInfo.Size())
	//当前文件大小大于等于文件的最大容量，返回true
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//分割日志文件

	//备份一下文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	//fmt.Printf("log file name:%s\n", f.filePath)

	logName := path.Join(f.filePath, fileInfo.Name())
	//logName := f.filePath + fileInfo.Name()
	//fmt.Printf("log file name:%s\n", logName)
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	//fmt.Printf(" back log file name:%s\n", newLogName)

	//关闭当前文件
	file.Close()
	//fmt.Println("close file sucess")
	//拼接一个文件备份的名字

	er := os.Rename(logName, newLogName)
	if er != nil {
		fmt.Println("rename failed")
		panic(err)
	}
	//打开一个新的文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}

	return fileObj, nil

}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {

	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funName, fileName, lineNo := getInfo(3)

		if f.checkSize(f.fileobj) {

			newFile, err := f.splitFile(f.fileobj)
			if err != nil {
				fmt.Printf("open new log file failed,err:%v\n", err)
				return
			}

			f.fileobj = newFile
		}
		fmt.Println("open new log file sucess")
		fmt.Fprintf(f.fileobj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), levelTostr[int(lv)], funName, fileName, lineNo, msg)

		if lv >= ERROR {

			if f.checkSize(f.errFileobj) {
				newFile, err := f.splitFile(f.errFileobj)
				if err != nil {
					return
				}
				f.errFileobj = newFile
			}
			//若要记录的日志大于等于ERROR级别，还是在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileobj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), levelTostr[int(lv)], funName, fileName, lineNo, msg)

		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {

	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {

	f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {

	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {

	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {

	f.log(FATAL, format, a...)
}

/*
func (f *FileLogger) close() {

	f.fileobj.Close()
	f.errFileobj.Close()
}
*/
