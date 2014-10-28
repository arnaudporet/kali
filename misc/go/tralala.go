
// clear && golang-go run tralala.go

package main

import "fmt"

// func concatenate(x,y [][]int) [][]int {
//     for i,_:=range x {
//         x[i]=append(x[i],y[i]...)
//     }
//     return x
// }

func main() {
    // x:=[][]int{
    //     {1,2},
    //     {5,6},
    // }
    // y:=[][]int{
    //     {3,4},
    //     {7,8},
    // }
    // x=concatenate(x,y)
    // fmt.Println(x)
    a1:=[][]bool{
        {true,true,false},
        {true,false,true},
    }
    a2:=[][]bool{
        {true,false,true},
        {true,true,false},
    }
    fmt.Println(compare_attractor(a1,a2))
}

func compare_attractor_set(A1,A2 [][][]bool) bool {
    if len(A1)!=len(A2) {return true} else {
        in_2:=[]bool{}
        for _,a1:=range A1 {
            z:=false
            for _,a2:=range A2 {
                if !compare_attractor(a1,a2) {
                    z=true
                    break
                }
            }
            in_2=append(in_2,z)
        }
        return !all(in_2)
    }
}

func compare_attractor(a1,a2 [][]bool) bool {
    var start1,start2 int
    if len(a1[0])!=len(a2[0]) {return true} else {
        start_found:=false
        for j1,_:=range a1[0] {
            for j2,_:=range a2[0] {
                z:=[]bool{}
                for i,_:=range a1 {z=append(z,a1[i][j1]==a2[i][j2])}
                if all(z) {
                    start_found=true
                    start1=j1
                    start2=j2
                    break
                }
            }
            if start_found {break}
        }
        if !start_found {return true} else {
            for j:=1;j<=len(a1[0])-1;j++ {
                z:=[]bool{}
                for i,_:=range a1 {z=append(z,a1[i][(start1+j)%len(a1[0])]==a2[i][(start2+j)%len(a2[0])])}
                if !all(z) {return true}
            }
            return false
        }
    }
}

func all(x []bool) bool {
    for _,value:=range x {if value==false {return false}}
    return true
}

func any(x []bool) bool {
    for _,value:=range x {if value==true {return true}}
    return false
}
