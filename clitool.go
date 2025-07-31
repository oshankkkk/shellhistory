package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)
func Talk(){
fmt.Println("helo appachchie")
scanner:=bufio.NewReader(os.Stdin)
fmt.Println("enter something yarr")
userinput,err:=scanner.ReadString('\n')

if err!=nil{fmt.Println("kelaa una")
return
}
fmt.Println(userinput)

}
func ReadFile(){
fmt.Println("this is a readfile method") 
data,err:=os.ReadFile("textfile.txt")
if err!=nil{
fmt.Println("hukao yako")
return

}
fmt.Println(string(data))
}
func ReadBigFile(){
fmt.Println("Read big file now")
data,err:=os.Open("README.md")
if err!=nil{
	fmt.Println("holyshit")
return
}
defer data.Close()
scanner:=bufio.NewScanner(data)
for scanner.Scan(){
fmt.Println(scanner.Text())
}
// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}


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








