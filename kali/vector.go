// Copyright (C) 2013-2016 Arnaud Poret
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
func (v1 Vector) Cat(v2 Vector) Vector {
    var (
        i int
        y Vector
    )
    y=v1.Copy()
    for i=range v2 {
        y=append(y,v2[i])
    }
    return y
}
func (v Vector) CircShift(n int) Vector {
    var (
        i int
        y Vector
    )
    y=make(Vector,len(v))
    for i=range v {
        y[i]=v[(i+n)%len(v)]
    }
    return y
}
func (v Vector) Copy() Vector {
    var y Vector
    y=make(Vector,len(v))
    copy(y,v)
    return y
}
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
func (v Vector) Fill(x float64) Vector {
    var (
        i int
        y Vector
    )
    y=make(Vector,len(v))
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
func (v Vector) GenArrangMat(k,narrang int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,int(math.Min(float64(narrang),math.Pow(float64(len(v)),float64(k)))))
    for i=range y {
        for {
            y[i]=v.GenArrangVect(k)
            if y[:i].Find(y[i],1)==-1 {
                break
            }
        }
    }
    return y
}
func (v Vector) GenArrangVect(k int) Vector {
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
func (v Vector) GenCombiMat(k,ncombi int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,int(math.Min(float64(ncombi),math.Gamma(float64(len(v)+1))/(math.Gamma(float64(k+1))*math.Gamma(float64(len(v)-k+1))))))
    for i=range y {
        for {
            y[i]=v.GenCombiVect(k)
            if y[:i].Find(y[i],1)==-1 {
                break
            }
        }
    }
    return y
}
func (v Vector) GenCombiVect(k int) Vector {
    // without repetitions
    var y Vector
    y=v.Sub(rand.Perm(len(v))[:k])
    sort.Float64s(y)
    return y
}
func (v Vector) GenS(n int) Matrix {
    var (
        i1,i2,m int
        z Vector
        y Matrix
    )
    y=v.ToMatrix(2)
    for i1=1;i1<n;i1++ {
        m=int(math.Pow(float64(len(v)),float64(i1)))
        for i2=0;i2<len(v)-1;i2++ {
            y=y.Cat(y[:m],1)
        }
        z=make(Vector,m).Fill(v[0])
        for i2=1;i2<len(v);i2++ {
            z=z.Cat(make(Vector,m).Fill(v[i2]))
        }
        y=y.Append(z,2)
    }
    return y
}
func IntToVector(x []int) Vector {
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
func (v Vector) MinPos() []int {
    var (
        imin,i int
        y []int
    )
    imin=0
    for i=1;i<len(v);i++ {
        if v[imin]>v[i] {
            imin=i
        }
    }
    y=[]int{imin}
    for i=imin+1;i<len(v);i++ {
        if v[i]==v[imin] {
            y=append(y,i)
        }
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
        y[int(b.Targ[i])]=b.Moda[i]
    }
    return y
}
func StringToVector(s []string) Vector {
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
func (v Vector) Sub(pos []int) Vector {
    var (
        i int
        y Vector
    )
    y=make(Vector,len(pos))
    for i=range pos {
        y[i]=v[pos[i]]
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
func (v Vector) ToInt() []int {
    var (
        i int
        y []int
    )
    y=make([]int,len(v))
    for i=range v {
        y[i]=int(v[i])
    }
    return y
}
func (v Vector) ToMatrix(d int) Matrix {
    var (
        i int
        y Matrix
    )
    if d==1 {
        y=Matrix{v.Copy()}
    } else if d==2 {
        y=make(Matrix,len(v))
        for i=range v {
            y[i]=Vector{v[i]}
        }
    }
    return y
}
func (v Vector) ToString() []string {
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
