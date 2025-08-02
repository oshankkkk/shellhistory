package main

import (
	"fmt"
	"log"
	"os"
)

func Keylistner(){
logfile,err:= os.OpenFile("example.txt",os.O_WRONLY| os.O_APPEND |os.O_CREATE,0644 )
if err!=nil{
	fmt.Println("something is wrong",err)
	return
}
defer logfile.Close()
daemonlogger:=log.New(logfile,"LOG",log.Ldate|log.Ltime)
daemonlogger.Println("Keylistner started")
} 
