// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "encoding/csv"
    "os"
)
type Matrix []Vector
func (m Matrix) Append(v Vector,d int) Matrix {
    var (
        i int
        y Matrix
    )
    y=m.Copy()
    if d==1 {
        y=append(y,v.Copy())
    } else if d==2 {
        for i=range v {
            y[i]=append(y[i],v[i])
        }
    }
    return y
}
func (m1 Matrix) Cat(m2 Matrix,d int) Matrix {
    var (
        i int
        y Matrix
    )
    y=m1.Copy()
    if d==1 {
        for i=range m2 {
            y=append(y,m2[i].Copy())
        }
    } else if d==2 {
        for i=range m2 {
            y[i]=y[i].Cat(m2[i])
        }
    }
    return y
}
func (m Matrix) CircShift(n,d int) Matrix {
    var (
        i int
        y Matrix
    )
    if d==1 {
        y=m.T()
        for i=range y {
            y[i]=y[i].CircShift(n)
        }
        y=y.T()
    } else if d==2 {
        y=m.Copy()
        for i=range y {
            y[i]=y[i].CircShift(n)
        }
    }
    return y
}
func (m Matrix) Cols(pos []int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(m))
    for i=range m {
        y[i]=m[i].Sub(pos)
    }
    return y
}
func (m Matrix) Copy() Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(m))
    for i=range m {
        y[i]=m[i].Copy()
    }
    return y
}
func (m1 Matrix) Equal(m2 Matrix) bool {
    var i int
    if len(m1)!=len(m2) {
        return false
    } else {
        for i=range m1 {
            if !m1[i].Equal(m2[i]) {
                return false
            }
        }
        return true
    }
}
func (m Matrix) Find(v Vector,d int) int {
    var (
        i int
        z Matrix
    )
    if d==1 {
        for i=range m {
            if m[i].Equal(v) {
                return i
            }
        }
    } else if d==2 {
        z=m.T()
        for i=range z {
            if z[i].Equal(v) {
                return i
            }
        }
    }
    return -1
}
func IntToMatrix(x [][]int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(x))
    for i=range x {
        y[i]=IntToVector(x[i])
    }
    return y
}
func LoadMatrix(filename string) Matrix {
    var (
        z [][]string
        reader *csv.Reader
        file *os.File
    )
    file,_=os.Open(filename)
    reader=csv.NewReader(file)
    reader.Comma=','
    reader.Comment=0
    reader.FieldsPerRecord=-1
    reader.LazyQuotes=false
    reader.TrimLeadingSpace=true
    z,_=reader.ReadAll()
    file.Close()
    return StringToMatrix(z)
}
func MakeMatrix(n,m int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,n)
    for i=range y {
        y[i]=make(Vector,m)
    }
    return y
}
func (m Matrix) Save(filename string) {
    var (
        writer *csv.Writer
        file *os.File
    )
    file,_=os.Create(filename)
    writer=csv.NewWriter(file)
    writer.Comma=','
    writer.UseCRLF=false
    writer.WriteAll(m.ToString())
    file.Close()
}
func (m Matrix) Size(d int) int {
    if d==1 {
        return len(m)
    } else if d==2 {
        return len(m[0])
    }
    return -1// here for "missing return at end of function"
}
func StringToMatrix(s [][]string) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(s))
    for i=range s {
        y[i]=StringToVector(s[i])
    }
    return y
}
func (m Matrix) T() Matrix {
    var (
        i,j int
        y Matrix
    )
    y=MakeMatrix(m.Size(2),m.Size(1))
    for i=range m {
        for j=range m[i] {
            y[j][i]=m[i][j]
        }
    }
    return y
}
func (m Matrix) ToInt() [][]int {
    var (
        i int
        y [][]int
    )
    y=make([][]int,len(m))
    for i=range m {
        y[i]=m[i].ToInt()
    }
    return y
}
func (m Matrix) ToString() [][]string {
    var (
        i int
        y [][]string
    )
    y=make([][]string,len(m))
    for i=range m {
        y[i]=m[i].ToString()
    }
    return y
}
