
// clear && golang-go run tralala.go

package main

import "fmt"

const x [3]int={1,2,3}

func main() {
    fmt.Println(x)
}

func f(x int) ([]int) {
    return []int{x,x}
}
