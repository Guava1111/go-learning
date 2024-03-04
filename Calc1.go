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
   
    //parts := strings.Split(x,"+") 找到字串裡的+剃除後印出
    parts := strings.Split(x,"")//一個一個分開印出
     num1,_ := strconv.Atoi(parts[0])
     num2,_ := strconv.Atoi(parts[2])
    fmt.Println(num1)
    fmt.Println(num2)
    if parts[1] == "+"{
        
        reult = num1 + num2
        fmt.Println(reult)
       }
    if parts[1] == "-"{
        reult = num1 - num2
        fmt.Println(reult)
       }
    if parts[1] == "/"{
        reult = num1 / num2
        fmt.Println(reult)
    }
    if parts[1] == "*"{
        reult = num1*num2
        fmt.Println(reult)
    }


}

