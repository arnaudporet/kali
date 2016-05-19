// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "strconv"
//#### Types #################################################################//
type Vector []float64
//#### CircShift #############################################################//
func (v *Vector) CircShift(n int) {
    var i int
    var y Vector
    y=make(Vector,len(*v))
    if n>=0 {
        for i=range (*v) {
            y[i]=(*v)[(i+n)%len(*v)]
        }
    } else {
        for i=range (*v) {
            y[(i-n)%len(*v)]=(*v)[i]
        }
    }
    copy((*v),y)
}
//#### Copy ##################################################################//
func (v Vector) Copy() Vector {
    var y Vector
    y=make(Vector,len(v))
    copy(y,v)
    return y
}
//#### Equal #################################################################//
func (v1 Vector) Equal(v2 Vector) bool {
    var i int
    if len(v1)!=len(v2) {
        return false
    } else {
        for i=range v1 {
            if v1[i]!=v2[i] {
                return false
            }
        }
        return true
    }
}
//#### Find ##################################################################//
func (v Vector) Find(x float64) int {
    var i int
    for i=range v {
        if v[i]==x {
            return i
        }
    }
    return -1
}
//#### ItoV ###################################################################//
func ItoV(x []int) Vector {
    var i int
    var y Vector
    for i=range x {
        y=append(y,float64(x[i]))
    }
    return y
}
//#### MinPos ################################################################//
func (v Vector) MinPos() []int {
    var i,imin int
    var y []int
    if len(v)>0 {
        imin=0
        for i=1;i<len(v);i++ {
            if v[i]<v[imin] {
                imin=i
            }
        }
        y=append(y,imin)
        for i=imin+1;i<len(v);i++ {
            if v[i]==v[imin] {
                y=append(y,i)
            }
        }
    }
    return y
}
//#### StoV ##################################################################//
func StoV(s []string) Vector {
    var i int
    var x float64
    var y Vector
    for i=range s {
        x,_=strconv.ParseFloat(s[i],64)
        y=append(y,x)
    }
    return y
}
//#### Sub ###################################################################//
func (v Vector) Sub(x []int) Vector {
    var i int
    var y Vector
    for i=range x {
        y=append(y,v[x[i]])
    }
    return y
}
//#### Sum ###################################################################//
func (v Vector) Sum() float64 {
    var i int
    var y float64
    y=0.0
    for i=range v {
        y+=v[i]
    }
    return y
}
//#### ToI ###################################################################//
func (v Vector) ToI() []int {
    var i int
    var y []int
    for i=range v {
        y=append(y,int(v[i]))
    }
    return y
}
//#### ToM ###################################################################//
func (v Vector) ToM(d int) Matrix {
    var i int
    var y Matrix
    if len(v)>0 {
        switch d {
            case 1:
                y=Matrix{v.Copy()}
            case 2:
                for i=range v {
                    y=append(y,Vector{v[i]})
                }
        }
    }
    return y
}
//#### ToS ###################################################################//
func (v Vector) ToS() []string {
    var i int
    var y []string
    for i=range v {
        y=append(y,strconv.FormatFloat(v[i],'f',-1,64))
    }
    return y
}
