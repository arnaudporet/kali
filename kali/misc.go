// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "math/rand"
    "os"
    "strings"
)
func Align(s []string,filler string) []string {
    var (
        wmax,i int
        y []string
    )
    wmax=MaxLen(s)
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
func GetInt(prompt string,deck []int) int {
    var (
        x int
        v Vector
    )
    v=IntToVect(deck)
    for {
        fmt.Print(prompt)
        fmt.Scan(&x)
        if len(deck)==0 || v.Find(float64(x))!=-1 {
            return x
        } else {
            fmt.Println("\nERROR: must be in ["+strings.Join(v.ToStr(),",")+"]")
        }
    }
}
func GoForward(f func(Vector) Vector,x0 Vector,b Bullet) Matrix {
    // asynchronous only
    var (
        i int
        x,y,z Vector
        fwd,stack Matrix
    )
    fwd=Matrix{x0.Copy()}
    stack=Matrix{x0.Copy()}
    for {
        x=stack[len(stack)-1].Copy()
        stack=stack[:len(stack)-1]
        y=f(x)
        for i=range y {
            z=x.Copy()
            z[i]=y[i]
            z=z.Shoot(b.Targ.ToInt(),b.Moda)
            if fwd.FindRow(z)==-1 {
                fwd=append(fwd,z.Copy())
                stack=append(stack,z.Copy())
            }
        }
        if len(stack)==0 {
            break
        }
    }
    return fwd.SortRows()
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
func MaxLen(s []string) int {
    var i,y int
    y=len(s[0])
    for i=1;i<len(s);i++ {
        if y<len(s[i]) {
            y=len(s[i])
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
func ReachCycle(f func(Vector) Vector,x0 Vector,b Bullet) Matrix {
    // synchronous only
    var (
        i int
        x Vector
        y Matrix
    )
    y=Matrix{x0.Copy()}
    x=x0.Copy()
    for {
        x=f(x).Shoot(b.Targ.ToInt(),b.Moda)
        i=y.FindRow(x)
        if i!=-1 {
            y=y[i:]
            break
        } else {
            y=append(y,x.Copy())
        }
    }
    return y.CircRows(y.MinRow())
}
func Walk(f func(Vector) Vector,x0 Vector,b Bullet,kmax int) Vector {
    // asynchronous only
    var (
        k,i int
        x,y Vector
    )
    x=x0.Copy()
    for k=0;k<kmax;k++ {
        y=f(x)
        i=rand.Intn(len(x))
        x[i]=y[i]
        x=x.Shoot(b.Targ.ToInt(),b.Moda)
    }
    return x
}
