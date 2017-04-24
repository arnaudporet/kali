// Copyright (C) 2013-2017 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "encoding/csv"
    "os"
)
type Matrix []Vector
func (m Matrix) AddCol(v Vector) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(m))
    for i=range v {
        y[i]=append(m[i].Copy(),v[i])
    }
    return y
}
func (m Matrix) CircRows(n int) Matrix {
    // n>0
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(m))
    for i=range m {
        y[i]=m[(i+n)%len(m)].Copy()
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
func (m1 Matrix) Eq(m2 Matrix) bool {
    var i int
    if len(m1)!=len(m2) {
        return false
    } else {
        for i=range m1 {
            if !m1[i].Eq(m2[i]) {
                return false
            }
        }
        return true
    }
}
func (m Matrix) FindRow(v Vector) int {
    var i int
    for i=range m {
        if m[i].Eq(v) {
            return i
        }
    }
    return -1
}
func LoadMat(filename string) Matrix {
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
    return StrToMat(z)
}
func MakeMat(n,m int) Matrix {
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
func (m Matrix) MinRow() int {
    // according to the lexicographical order
    var i,imin int
    imin=0
    for i=1;i<len(m);i++ {
        if m[imin].Sup(m[i]) {
            imin=i
        }
    }
    return imin
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
    writer.WriteAll(m.ToStr())
    file.Close()
}
func (m Matrix) SortRows() Matrix {
    // according to the lexicographical order
    var (
        i,imin int
        y,z Matrix
    )
    y=make(Matrix,len(m))
    z=m.Copy()
    for i=range y {
        imin=z.MinRow()
        y[i]=z[imin].Copy()
        z=append(z[:imin],z[imin+1:]...)
    }
    return y
}
func StrToMat(s [][]string) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(s))
    for i=range s {
        y[i]=StrToVect(s[i])
    }
    return y
}
func (m Matrix) T() Matrix {
    var (
        i,j int
        y Matrix
    )
    y=MakeMat(len(m[0]),len(m))
    for i=range m {
        for j=range m[i] {
            y[j][i]=m[i][j]
        }
    }
    return y
}
func (m Matrix) ToStr() [][]string {
    var (
        i int
        y [][]string
    )
    y=make([][]string,len(m))
    for i=range m {
        y[i]=m[i].ToStr()
    }
    return y
}
