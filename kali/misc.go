// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "os"
    "strings"
)
func Align(s []string,filler string) []string {
    var (
        wmax,i int
        y []string
    )
    wmax=len(s[0])
    for i=1;i<len(s);i++ {
        if wmax<len(s[i]) {
            wmax=len(s[i])
        }
    }
    y=make([]string,len(s))
    for i=range s {
        y[i]=s[i]+strings.Repeat(filler,wmax-len(s[i]))
    }
    return y
}
func Exist(filename string) bool {
    var err error
    _,err=os.Stat(filename)
    return !os.IsNotExist(err)
}
func GetInt(message string,deck []int) int {
    var (
        x int
        v Vector
    )
    v=IntToVector(deck)
    for {
        fmt.Print(message)
        fmt.Scan(&x)
        if len(deck)==0 || v.Find(float64(x))!=-1 {
            return x
        } else {
            fmt.Println("\nERROR: must be in ["+strings.Join(v.ToString(),",")+"]")
        }
    }
}
func Max(x ...float64) float64 {
    var (
        i int
        y float64
    )
    y=x[0]
    for i=1;i<len(x);i++ {
        if y<x[i] {
            y=x[i]
        }
    }
    return y
}
func Min(x ...float64) float64 {
    var (
        i int
        y float64
    )
    y=x[0]
    for i=1;i<len(x);i++ {
        if y>x[i] {
            y=x[i]
        }
    }
    return y
}
func Range(a,b int) []int {
    var (
        i int
        y []int
    )
    y=make([]int,b-a)
    for i=range y {
        y[i]=a+i
    }
    return y
}
