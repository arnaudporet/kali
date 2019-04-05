// Copyright (C) 2013-2019 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "math"
    "math/rand"
    "sort"
    "strconv"
)
type Vector []float64
func (v Vector) Arrang(k int) Vector {
    // with repetitions
    var (
        i int
        y Vector
    )
    y=make(Vector,k)
    for i=range y {
        y[i]=v[rand.Intn(len(v))]
    }
    return y
}
func (v Vector) Arrangs(k,narrang int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,NearInt(math.Min(float64(narrang),math.Pow(float64(len(v)),float64(k)))))
    for i=range y {
        for {
            y[i]=v.Arrang(k)
            if y[:i].FindRow(y[i])==-1 {
                break
            }
        }
    }
    return y
}
func (v Vector) Combi(k int) Vector {
    // without repetitions
    var (
        i int
        z []int
        y Vector
    )
    y=make(Vector,k)
    z=rand.Perm(len(v))
    for i=range y {
        y[i]=v[z[i]]
    }
    sort.Float64s(y)
    return y
}
func (v Vector) Combis(k,ncombi int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,NearInt(math.Min(float64(ncombi),math.Gamma(float64(len(v)+1))/(math.Gamma(float64(k+1))*math.Gamma(float64(len(v)-k+1))))))
    for i=range y {
        for {
            y[i]=v.Combi(k)
            if y[:i].FindRow(y[i])==-1 {
                break
            }
        }
    }
    return y
}
func (v Vector) Copy() Vector {
    var y Vector
    y=make(Vector,len(v))
    copy(y,v)
    return y
}
func (v1 Vector) Eq(v2 Vector) bool {
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
func FillVect(n int,x float64) Vector {
    var (
        i int
        y Vector
    )
    y=make(Vector,n)
    for i=range y {
        y[i]=x
    }
    return y
}
func (v Vector) Find(x float64) int {
    var i int
    for i=range v {
        if v[i]==x {
            return i
        }
    }
    return -1
}
func IntToVect(x []int) Vector {
    var (
        i int
        y Vector
    )
    y=make(Vector,len(x))
    for i=range x {
        y[i]=float64(x[i])
    }
    return y
}
func (v Vector) Shoot(pos []int,val Vector) Vector {
    var (
        i int
        y Vector
    )
    y=v.Copy()
    for i=range pos {
        y[pos[i]]=val[i]
    }
    return y
}
func (v Vector) Space(n int) Matrix {
    var (
        i1,i2,m int
        z Vector
        y Matrix
    )
    y=v.ToCol()
    for i1=1;i1<n;i1++ {
        m=NearInt(math.Pow(float64(len(v)),float64(i1)))
        for i2=0;i2<len(v)-1;i2++ {
            y=append(y,y[:m].Copy()...)
        }
        z=FillVect(m,v[0])
        for i2=1;i2<len(v);i2++ {
            z=append(z,FillVect(m,v[i2])...)
        }
        y=y.AddCol(z)
    }
    return y
}
func StrToVect(s []string) Vector {
    var (
        i int
        y Vector
    )
    y=make(Vector,len(s))
    for i=range s {
        y[i],_=strconv.ParseFloat(s[i],64)
    }
    return y
}
func (v Vector) Sum() float64 {
    var (
        i int
        y float64
    )
    y=0.0
    for i=range v {
        y+=v[i]
    }
    return y
}
func (v1 Vector) Sup(v2 Vector) bool {
    // according to the lexicographical order
    var i int
    for i=0;i<NearInt(math.Min(float64(len(v1)),float64(len(v2))));i++ {
        if v1[i]<v2[i] {
            return false
        } else if v1[i]>v2[i] {
            return true
        }
    }
    return len(v1)>len(v2)
}
func (v Vector) ToCol() Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(v))
    for i=range v {
        y[i]=Vector{v[i]}
    }
    return y
}
func (v Vector) ToInt() []int {
    var (
        i int
        y []int
    )
    y=make([]int,len(v))
    for i=range v {
        y[i]=NearInt(v[i])
    }
    return y
}
func (v Vector) ToStr() []string {
    var (
        i int
        y []string
    )
    y=make([]string,len(v))
    for i=range v {
        y[i]=strconv.FormatFloat(v[i],'f',-1,64)
    }
    return y
}
