
// clear && golang-go run tralala.go

package main

import (
    "math"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(int64(time.Now().Nanosecond()))
    generate_arrangement(10,10000)
}

func generate_arrangement(k,n_arrang int) (arrang_mat [][]int) {
    ////////////////////    /!\ only with repetition /!\    ////////////////////
    var i1,j1,i2 int
    var arrang []int
    var in_arrang_mat bool
    var z []bool
    var max_arrang int
    max_arrang=int(math.Min(float64(n_arrang),math.Pow(float64(2),float64(k))))
    for i1=1;i1<=max_arrang;i1++ {
        for {
            arrang=[]int{}
            for j1=1;j1<=k;j1++ {arrang=append(arrang,rand.Intn(2))}
            in_arrang_mat=false
            for i2=range arrang_mat {
                z=[]bool{}
                for j1=range arrang {z=append(z,arrang[j1]==arrang_mat[i2][j1])}
                if all(z) {
                    in_arrang_mat=true
                    break
                }
            }
            if !in_arrang_mat {
                arrang_mat=append(arrang_mat,arrang)
                break
            }
        }
    }
    return
}

func all(x []bool) bool {
    var i1 int
    for i1=range x {if x[i1]==false {return false}}
    return true
}
