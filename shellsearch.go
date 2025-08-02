package main

import (
	"fmt" 
	"os"
"os/user"

)

func Readshellhistory() (*os.File, error){
	currentuser,err:=user.Current()
	bashhistoryfile:=currentuser.HomeDir	

	bashhistory, err:=os.Open(bashhistoryfile+"/.bash_history")
	if err!=nil{
		fmt.Println("kela una")
return nil, err
	}

	return bashhistory,nil
}
func SearchForMatches(){


}
