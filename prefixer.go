package main

import "fmt"


type Node struct{
children map[string]*Node
isEndLetter bool

}

func Nodeinit()*Node{
return &Node{children: make(map[string]*Node),isEndLetter: false}
}

type Tree struct{
Root *Node
}

func TreeInit()*Tree{
return &Tree{
Root: Nodeinit(),}
}
func(t *Tree)InsertLetter(word string){
	currentnode:=t.Root
	for _,value:=range word{
		letter:=string(value)
		if _,exits:= currentnode.children[letter]; exits{
			currentnode=currentnode.children[letter]	
		}else{
			newnode:=Nodeinit()
			currentnode.children[letter]=newnode
			currentnode=currentnode.children[letter]

		}
		
	} 
currentnode.isEndLetter=true
		fmt.Println("hari oya inserted ")
}


func helperShowtree(key string, currentnode *Node) {
	fmt.Println(key)
	if len(currentnode.children) > 0 {
		for key, node := range currentnode.children {
			helperShowtree(key, node)
		}
	}
}

func(t *Tree) Showtree(){
	currentnode:=t.Root
	for key,node:=range currentnode.children{

helperShowtree(key,node)
	}
}







