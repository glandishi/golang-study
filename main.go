package main

import "fmt"

var x string = "hello"
const y int = 69

func main(){
	z := 1
	fmt.Printf("value of variable x is %v and type is %T\n",x,x)
	fmt.Printf("value of variable y is %v and type is %T\n",y,y)
	fmt.Printf("value of variable z is %v and type is %T\n",z,z)
}