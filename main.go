package main
import "fmt"
func main(){
	fmt.Println("main function")
	tree:=TreeInit()
	tree.InsertLetter("oshan")
	tree.Showtree()
}
