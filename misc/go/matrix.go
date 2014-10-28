
// clear && golang-go run matrix.go

package main

import "fmt"

type matrix struct {
    mat [][]int
}

func (x *matrix) shape() []int {
    return []int{len(x.mat),len(x.mat[0])}
}

func (x *matrix) row() int {
    return len(x.mat)
}

func (x *matrix) col() int {
    return len(x.mat[0])
}

func (x *matrix) getrow(i int) []int {
    return x.mat[i-1]
}

func (x *matrix) getcol(i int) []int {
    z:=[]int{}
    for key,_:=range x.mat {
        z=append(z,x.mat[key][i-1])
    }
    return z
}

func main() {
    M:=matrix{
        [][]int{
            {1,2,3},
            {4,5,6},
        },
    }
    fmt.Println(M.mat)
    fmt.Println(M.shape())
    fmt.Println(M.row())
    fmt.Println(M.col())
    fmt.Println(M.getrow(2))
    fmt.Println(M.getcol(2))
}
