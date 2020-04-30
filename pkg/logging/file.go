package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt 	= "txt"
	TimeFormat = "20060102"
)

func getLogFilePath()string{
	return fmt.Sprintf("%s",logSavePath)
}

func getFileFullPath()string{
	prefixPath := getLogFilePath()
	suffixPath :=fmt.Sprintf("%s%s.%s",LogSaveName,time.Now().Format(TimeFormat ),LogFileExt)
	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
}

func openLogFile(filePath string)*os.File{
	_,err:=os.Stat( filePath)

	switch  {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("permission denied:%v",err)
	}

	//O_RDONLY int = syscall.O_RDONLY // 以只读模式打开文件
	//O_WRONLY int = syscall.O_WRONLY // 以只写模式打开文件
	//O_RDWR   int = syscall.O_RDWR   // 以读写模式打开文件
	//// The remaining values may be or'ed in to control behavior.
	//O_APPEND int = syscall.O_APPEND // 在写入时将数据追加到文件中
	//O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件
	//O_EXCL   int = syscall.O_EXCL   // 使用O_CREATE时，文件必须不存在
	//O_SYNC   int = syscall.O_SYNC   // 同步IO
	//O_TRUNC  int = syscall.O_TRUNC  // 如果可以，打开时

	handle,err:=os.OpenFile(filePath , os.O_APPEND |os.O_CREATE |os.O_WRONLY,0644)
	if err!=nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

func mkDir(){
	dir,_:=os.Getwd()  //获取当前目录 get working directory
	err:=os.MkdirAll(dir+"/"+getLogFilePath(),os.ModePerm)		//MkdirAll创建一个名为path的目录， 以及任何必要的父母，并且返回零， 否则返回错误。 权限位perm（umask之前）用于所有 /MkdirAll创建的目录。 如果path已经是一个目录，MkdirAll将不执行任何操作 并且返回零。
	if err!=nil{
		panic(err)
	}
}