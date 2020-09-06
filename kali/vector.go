// Copyright (C) 2013-2020 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

// WARNING The functions in the present file do not handle exceptions and
// errors. Instead, they assume that such handling is performed upstream by the
// <do*> top-level functions of kali. Consequently, they should not be used as
// is outside of kali.

package kali
import (
    "math"
    "math/rand"
    "sort"
    "strconv"
)
type Vector []float64
func (v Vector) Arrangs(k,n int) Matrix {
    var (
        i,j int
        z Vector
        y Matrix
    )
    y=make(Matrix,int(math.Round(math.Min(float64(n),math.Pow(float64(len(v)),float64(k))))))
    z=make(Vector,k)
    for i=range y {
        for {
            for j=range z {
                z[j]=v[rand.Intn(len(v))]
            }
            if y[:i].FindRow(z)==-1 {
                break
            }
        }
        y[i]=z.Copy()
    }
    return y
}
func (v Vector) Combis(k,n int) Matrix {
    var (
        i,j int
        pos []int
        z Vector
        y Matrix
    )
    y=make(Matrix,int(math.Round(math.Min(float64(n),math.Gamma(float64(len(v)+1))/(math.Gamma(float64(k+1))*math.Gamma(float64(len(v)-k+1)))))))
    z=make(Vector,k)
    for i=range y {
        for {
            pos=rand.Perm(len(v))
            for j=range z {
                z[j]=v[pos[j]]
            }
            sort.Float64s(z)
            if y[:i].FindRow(z)==-1 {
                break
            }
        }
        y[i]=z.Copy()
    }
    return y
}
func (v Vector) Copy() Vector {
    var (
        y Vector
    )
    y=make(Vector,len(v))
    copy(y,v)
    return y
}
func (v1 Vector) Eq(v2 Vector) bool {
    var (
        y bool
        i int
    )
    if len(v1)!=len(v2) {
        y=false
    } else {
        y=true
        for i=range v1 {
            if v1[i]!=v2[i] {
                y=false
                break
            }
        }
    }
    return y
}
func (v Vector) Find(x float64) int {
    var (
        y,i int
    )
    y=-1
    for i=range v {
        if v[i]==x {
            y=i
            break
        }
    }
    return y
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
func MakeVect(n int,x float64) Vector {
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
func (v Vector) Shoot(b Bullet) Vector {
    var (
        i int
        y Vector
    )
    y=v.Copy()
    for i=range b.Targ {
        y[int(math.Round(b.Targ[i]))]=b.Moda[i]
    }
    return y
}
func (v Vector) Space(n int) Matrix {
    var (
        i,j,m int
        z Vector
        y Matrix
    )
    y=v.ToCol()
    for i=1;i<n;i++ {
        m=int(math.Round(math.Pow(float64(len(v)),float64(i))))
        for j=0;j<len(v)-1;j++ {
            y=append(y,y[:m]...)
        }
        z=MakeVect(m,v[0])
        for j=1;j<len(v);j++ {
            z=append(z,MakeVect(m,v[j])...)
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
    y=0
    for i=range v {
        y+=v[i]
    }
    return y
}
func (v1 Vector) Sup(v2 Vector) bool {
    var (
        y bool
        i int
    )
    y=len(v1)>len(v2)
    for i=0;i<int(math.Round(math.Min(float64(len(v1)),float64(len(v2)))));i++ {
        if v1[i]<v2[i] {
            y=false
            break
        } else if v1[i]>v2[i] {
            y=true
            break
        }
    }
    return y
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
        y[i]=int(math.Round(v[i]))
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
