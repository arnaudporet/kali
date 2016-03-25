// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "strconv"
//#### Types #################################################################//
type Vector []float64
//#### Cat ###################################################################//
func (v1 *Vector) Cat(v2 Vector) {
    var i int
    for i=range v2 {
        (*v1)=append((*v1),v2[i])
    }
}
//#### CircShift #############################################################//
func (v *Vector) CircShift(n int) {
    var i int
    var y Vector
    y=make(Vector,len(*v))
    for i=range (*v) {
        y[i]=(*v)[(i+n)%len(*v)]
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
//#### Pos ##################################################################//
func (v Vector) Pos(x []int) Vector {
    var i int
    var y Vector
    if len(v)>0 {
        for i=range x {
            y=append(y,v[x[i]])
        }
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
//#### ToV ###################################################################//
func ToV(x []int) Vector {
    var i int
    var y Vector
    for i=range x {
        y=append(y,float64(x[i]))
    }
    return y
}
