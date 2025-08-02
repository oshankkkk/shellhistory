package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

func ReadHistory(){
	fmt.Println("reading bash history rn.....")
	usr,err:=user.Current()
	if err!=nil{
		fmt.Println("holyshit")
		return
	}
	home:=usr.HomeDir
	fmt.Println(home)
	hisfile,errr:=os.Open(home+"/.bash_history")
	if errr!=nil{
		fmt.Println("holyshit",errr)
		return
	}
	defer hisfile.Close()
	scanner:=bufio.NewScanner(hisfile)
	for scanner.Scan(){
		fmt.Println(scanner.Text())

	}

}








