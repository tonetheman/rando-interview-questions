package main

import (
	"fmt"
)

type intholder struct {
	key int
	next * intholder
}

func makeNewInt(val int) *intholder {
	return &intholder{val,nil}
}

func insert(newint * intholder, root * intholder) *intholder{
	if root == nil {
		return root
	}
	for i:=root; i!=nil; i = i.next {
		if i.next == nil {
			i.next = newint
			return root
		}
	}
	return root
}

func pr(root * intholder) {
	count :=0
	for i:=root;i!=nil;i=i.next {
		fmt.Println(i,count)
		count++
	}
}

func main() {
	a := makeNewInt(10)
	root := a
	fmt.Println("root and a",root,a,"before any nodes added")
	b := makeNewInt(20)

	root = insert(root,b)
	fmt.Println("root after append",root)
	c := makeNewInt(30)
	root = insert(root,c)
	pr(root)
}