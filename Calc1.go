package main
import "fmt"

func main(){
var  x, y int
var result int
var a string

fmt.Printf("請輸入:")
fmt.Scan(&x)
fmt.Printf("請輸入運算符號:")
fmt.Scan(&a)
fmt.Printf("請輸入:")
fmt.Scan(&y)
//fmt.Println(x, a, y)
//result = fmt.Sprintln(x,a,y)
if a=="+"{
	result = x+y
	fmt.Println("答案:",result)
}
if a == "-"{
	result = x-y
	fmt.Println("答案:",result)
}
if a == "/"{
	result = x/y
	fmt.Println("答案:",result)
}
}