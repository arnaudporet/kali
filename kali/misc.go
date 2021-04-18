// Copyright (C) 2013-2021 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

// WARNING The functions in the present file do not handle exceptions and
// errors. Instead, they assume that such handling is performed upstream by the
// <do*> top-level functions of kali. Consequently, they should not be used as
// is outside of kali.

package kali
import (
    "fmt"
    "math/rand"
    "os"
    "strings"
)
func Align(s []string,filler string) []string {
    var (
        wMax,i int
        y []string
    )
    wMax=MaxLen(s)
    y=make([]string,len(s))
    for i=range s {
        y[i]=s[i]+strings.Repeat(filler,wMax-len(s[i]))
    }
    return y
}
func CheckF(f func(Vector) Vector,vals Vector,n int,ok *bool) {
    var (
        i int
    )
    defer Recover(ok)
    for i=range vals {
        _=f(MakeVect(n,vals[i]))
    }
}
func Exist(fileName string) bool {
    var (
        err error
    )
    _,err=os.Stat(fileName)
    return !os.IsNotExist(err)
}
func GetInt(prompt string,deck []int) int {
    var (
        err error
        x int
        vDeck Vector
    )
    vDeck=IntToVect(deck)
    for {
        fmt.Print(prompt)
        _,err=fmt.Scan(&x)
        if err!=nil {
            panic("GetInt(prompt,deck): "+err.Error())
        } else if (len(vDeck)==0) || (vDeck.Find(float64(x))!=-1) {
            break
        } else {
            fmt.Println("must be in ["+strings.Join(vDeck.ToStr(),",")+"]")
        }
    }
    return x
}
func GoForward(x0 Vector,f func(Vector) Vector,b Bullet,maxForward int) (Matrix,bool) {
    var (
        skipped bool
        i,j int
        forward,newCheck,toCheck,succ Matrix
    )
    skipped=false
    forward=append(forward,x0.Copy())
    newCheck=append(newCheck,x0.Copy())
    for {
        toCheck=newCheck.Copy()
        newCheck=Matrix{}
        for i=range toCheck {
            succ=toCheck[i].Succ(f,b)
            for j=range succ {
                if forward.FindRow(succ[j])==-1 {
                    forward=append(forward,succ[j].Copy())
                    if (maxForward>0) && (len(forward)>maxForward) {
                        skipped=true
                        break
                    } else {
                        newCheck=append(newCheck,succ[j].Copy())
                    }
                }
            }
            if skipped {
                break
            }
        }
        if skipped || (len(newCheck)==0) {
            break
        }
    }
    if !skipped {
        forward=forward.SortRows()
    }
    return forward,skipped
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
    var (
        y,i int
    )
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
    if a==b {
        y=[]int{a}
    } else {
        y=make([]int,b-a)
        for i=range y {
            y[i]=a+i
        }
    }
    return y
}
func ReachCycle(x0 Vector,f func(Vector) Vector,b Bullet) Matrix {
    var (
        i int
        x Vector
        y Matrix
    )
    y=append(y,x0.Copy())
    x=x0.Copy()
    for {
        x=f(x)
        x=x.Shoot(b)
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
func Recover(ok *bool) {
    var (
        recov interface{}
    )
    recov=recover()
    if recov!=nil {
        fmt.Println(recov)
        (*ok)=false
    } else {
        (*ok)=true
    }
}
func (x Vector) Succ(f func(Vector) Vector,b Bullet) Matrix {
    var (
        i int
        y,z Vector
        succ Matrix
    )
    y=f(x)
    for i=range y {
        z=x.Copy()
        z[i]=y[i]
        z=z.Shoot(b)
        if succ.FindRow(z)==-1 {
            succ=append(succ,z.Copy())
        }
    }
    return succ
}
func Walk(x0 Vector,f func(Vector) Vector,b Bullet,nSteps int) Vector {
    var (
        k,i int
        x,y Vector
    )
    x=x0.Copy()
    for k=0;k<nSteps;k++ {
        y=f(x)
        i=rand.Intn(len(y))
        x[i]=y[i]
        x=x.Shoot(b)
    }
    return x
}
