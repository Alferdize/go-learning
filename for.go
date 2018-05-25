package main 

import "fmt"

func main(){
var a int
var b int = 15

numbers := [10]int{1,1,1,1,12,24,5,2,1,6}

for a := 0; a < 10; a++{
fmt.Printf("VALUEVof a: %d\n",a)
}

for a < b{
a++
fmt.Printf("VALUE OF a %d\n",a)
}
for i,x := range numbers{
fmt.Printf("VALUE OF x = %d at %d\n",x,i)
}
}