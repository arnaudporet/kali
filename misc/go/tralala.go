
// clear && golang-go run tralala.go

package main

import (
    "fmt"
)

func main() {
    x:=[][]int{
        {1,2},
        {3,4},
    }
    fmt.Println(x)
    fmt.Println(transpose(x))
}

func transpose(x [][]int) [][]int {
    y:=[][]int{}
    for i,_:=range x {
        y=append(y,[]int{})
        for j,_:=range x[i] {
            y[len(y)-1]=append(y[len(y)-1],x[j][i])
        }
    }
    return y
}





