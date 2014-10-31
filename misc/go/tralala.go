
// clear && golang-go run tralala.go

package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    y:=generate_arrangement(3,10000)
    x:=transpose(y)
    for _,line:=range x {fmt.Println(line)}
}

func transpose(x [][]bool) [][]bool {
    y:=[][]bool{}
    for j1:=range x[0] {
        y=append(y,[]bool{})
        for i1:=range x {y[j1]=append(y[j1],x[i1][j1])}
    }
    return y
}

func generate_arrangement(k,n_arrang int) [][]bool {
    ////////////////////    /!\ only with repetition /!\    ////////////////////
    arrang_mat:=[][]bool{}
    for i1:=1;i1<=int(math.Min(float64(n_arrang),math.Pow(float64(2),float64(k))));i1++ {
        for {
            arrang:=[]bool{}
            for j1:=1;j1<=k;j1++ {
                z:=false
                if rand.Intn(2)==1 {z=true}
                arrang=append(arrang,z)
            }
            in_arrang_mat:=false
            for i2:=range arrang_mat {
                z:=[]bool{}
                for j1:=range arrang {z=append(z,arrang[j1]==arrang_mat[i2][j1])}
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
    return arrang_mat
}

func all(x []bool) bool {
    for i1:=range x {if x[i1]==false {return false}}
    return true
}
