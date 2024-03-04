package main
import (
    "fmt"
    "strings"
    "strconv"
    
)

func main(){
    var x string
    var reult int
    fmt.Scanln(&x)
   
    if strings.Contains(x,"+"){
		parts := strings.Split(x,"+")
		num1,_ := strconv.Atoi(parts[0])
		num2,_ := strconv.Atoi(parts[1])
		reult = num1 + num2
		fmt.Println(reult)
	}
	if strings.Contains(x,"-"){
		parts := strings.Split(x,"-")
		num1,_ := strconv.Atoi(parts[0])
		num2,_ := strconv.Atoi(parts[1])
		reult = num1 - num2
		fmt.Println(reult)
	}
	if strings.Contains(x,"*"){
		parts := strings.Split(x,"*")
		num1,_ := strconv.Atoi(parts[0])
		num2,_ := strconv.Atoi(parts[1])
		reult = num1 * num2
		fmt.Println(reult)
	}
	if strings.Contains(x,"/"){
		parts := strings.Split(x,"/")
		num1,_ := strconv.Atoi(parts[0])
		num2,_ := strconv.Atoi(parts[1])
		reult = num1 / num2
		fmt.Println(reult)
	}
}
