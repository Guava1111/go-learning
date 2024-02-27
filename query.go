package main
import "fmt"

func query (x []int,a int)bool {
var ans bool
	for t:=0;t<=len(x)-1;t++{
	    ans =x[t]==a
		if ans == true{
			return ans

		}
	}
	return (ans)
}


func main(){
var y[]int=[]int{10,21,13,14,51}
var z  int
fmt.Scanln(&z)
var ann bool=query(y,z)
fmt.Println(ann)
}

