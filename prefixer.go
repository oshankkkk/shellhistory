package main

import "fmt"
type Tree struct{


	Content string
	Children []*Tree
}
func Newtree(value string) *Tree{
return &Tree{Content: value} 
}

// why cant we just go &Tree{Content: value}?
func(t *Tree) InsertNode(node string){
	child :=Newtree(node)
	t.Children=append(t.Children, child)
}

func(t *Tree) ShowTree(){
	fmt.Println("Gonna print the tree now")
	for _,value:=range t.Children{
	fmt.Println(*&value.Content ," node")
	}
}
