
// clear && golang-go run tralala.go

package main

import (
    "fmt"
    "strings"
)

func main() {
    x:="y"
    fmt.Scanf("%s",&x)
    x=strings.ToLower(x)
    fmt.Println(x==("n" || "no"))
}
