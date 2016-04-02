// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "os"
import "strings"
//#### Exist #################################################################//
func Exist(file string) bool {
    var err error
    _,err=os.Stat(file)
    return !os.IsNotExist(err)
}
//#### FillToMaxLen ##########################################################//
func FillToMaxLen(s []string) []string {
    var wmax,i int
    var z []string
    if len(s)>0 {
        z=make([]string,len(s))
        copy(z,s)
        wmax=len(z[0])
        for i=1;i<len(z);i++ {
            if wmax<len(z[i]) {
                wmax=len(z[i])
            }
        }
        for i=range z {
            for len(z[i])<wmax {
                z[i]+=" "
            }
        }
    }
    return z
}
//#### Min ###################################################################//
func Min(x ...float64) float64 {
    var i int
    var y float64
    y=x[0]
    for i=1;i<len(x);i++ {
        if y>x[i] {
            y=x[i]
        }
    }
    return y
}
//#### Max ###################################################################//
func Max(x ...float64) float64 {
    var i int
    var y float64
    y=x[0]
    for i=1;i<len(x);i++ {
        if y<x[i] {
            y=x[i]
        }
    }
    return y
}
//#### Prompt ################################################################//
func Prompt(message string,deck Vector) float64 {
    var x float64
    for {
        fmt.Print(message)
        fmt.Scan(&x)
        if len(deck)==0 || deck.Find(x)>-1 {
            return x
        } else {
            fmt.Println("\nERROR: must be in ["+strings.Join(deck.ToS(),",")+"]")
        }
    }
}
//#### Range #################################################################//
func Range(a,b int) []int {
    var i int
    var y []int
    for i=0;i<b-a;i++ {
        y=append(y,a+i)
    }
    return y
}
