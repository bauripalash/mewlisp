package main

import (
	"mewlisp/mewlisp"
)

func main(){
	//fmt.Println("Mewl - If Cats knew lisp")
	sample := "(define x (+ 10 11))"
	a := mewlisp.NewParser(sample)
	a.Debug()	
	
}
