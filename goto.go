package main

import "fmt"


func main(){
var n int = 67
i := 2
check:
if i == n {
goto next;
} else {
j := 2
check1:
if i != j{
if i % j ==0 {
i += 1
goto check
}
j +=1
goto check1
} else {
fmt.Printf("Prime no is %d \n",i)
i +=1
goto check;
}
}
next:
}