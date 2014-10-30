
// clear && golang-go run tralala.go

package main

import (
    "fmt"
    "os"
    "encoding/csv"
    "strconv"
)

func main() {
    A:=[][][]bool{}
    set_name:="tralala.csv"
    csv_file,_:=os.Open(set_name)
    csv_reader:=csv.NewReader(csv_file)
    csv_reader.FieldsPerRecord=-1
    s,_:=csv_reader.ReadAll()
    csv_file.Close()
    s_bis:=[][]bool{}
    for i,_:=range s {
        if i>=1 {
            s_bis=append(s_bis,[]bool{})
            for j,_:=range s[i] {
                z,_:=strconv.ParseBool(s[i][j])
                s_bis[len(s_bis)-1]=append(s_bis[len(s_bis)-1],z)
            }
        }
    }
    n_attractor,_:=strconv.ParseInt(s[0][0],10,0)
    n_line:=(len(s)-1)/int(n_attractor)
    for i:=1;i<=int(n_attractor);i++ {A=append(A,s_bis[(i-1)*n_line:i*n_line])}
    for i1,_:=range A {
        for i2,_:=range A[i1] {
            fmt.Println(A[i1][i2])
        }
        fmt.Println("-------------------------")
    }
}
