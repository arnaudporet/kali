// Copyright (C) 2013-2019 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

// WARNING The functions in the present file do not handle exceptions and
// errors. Instead, they assume that such handling is performed upstream by the
// <do*> top-level functions of kali. Consequently, they should not be used as
// is outside of kali.

package kali
import (
    "encoding/csv"
    "math"
    "os"
)
type Matrix []Vector
func (m Matrix) AddCol(v Vector) Matrix {
    var (
        i int
        y Matrix
    )
    y=m.Copy()
    for i=range v {
        y[i]=append(y[i],v[i])
    }
    return y
}
func (m Matrix) CircRows(n int) Matrix {
    var (
        i int
        y Matrix
    )
    y=make(Matrix,len(m))
    for i=range m {
        y[i]=m[int(math.Round(math.Mod(float64(i+n),float64(len(m)))))].Copy()
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
    var (
        y bool
        i int
    )
    if len(m1)!=len(m2) {
        y=false
    } else {
        y=true
        for i=range m1 {
            if !m1[i].Eq(m2[i]) {
                y=false
                break
            }
        }
    }
    return y
}
func (m Matrix) FindRow(v Vector) int {
    var (
        y,i int
    )
    y=-1
    for i=range m {
        if m[i].Eq(v) {
            y=i
            break
        }
    }
    return y
}
func LoadMat(fileName string) Matrix {
    var (
        lines [][]string
        file *os.File
        reader *csv.Reader
    )
    file,_=os.Open(fileName)
    reader=csv.NewReader(file)
    reader.Comma=','
    reader.Comment=0
    reader.FieldsPerRecord=-1
    reader.LazyQuotes=false
    reader.TrimLeadingSpace=true
    reader.ReuseRecord=true
    lines,_=reader.ReadAll()
    file.Close()
    return StrToMat(lines)
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
    var (
        iMin,i int
    )
    iMin=0
    for i=1;i<len(m);i++ {
        if m[iMin].Sup(m[i]) {
            iMin=i
        }
    }
    return iMin
}
func (m Matrix) Save(fileName string) {
    var (
        file *os.File
        writer *csv.Writer
    )
    file,_=os.Create(fileName)
    writer=csv.NewWriter(file)
    writer.Comma=','
    writer.UseCRLF=false
    _=writer.WriteAll(m.ToStr())
    file.Close()
}
func (m Matrix) SortRows() Matrix {
    var (
        i,iMin int
        y,z Matrix
    )
    y=make(Matrix,len(m))
    z=m.Copy()
    for i=range m {
        iMin=z.MinRow()
        y[i]=z[iMin].Copy()
        z=append(z[:iMin],z[iMin+1:]...)
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
func (m1 Matrix) Sup(m2 Matrix) bool {
    var (
        y bool
        i int
    )
    y=len(m1)>len(m2)
    for i=0;i<int(math.Round(math.Min(float64(len(m1)),float64(len(m2)))));i++ {
        if m2[i].Sup(m1[i]) {
            y=false
            break
        } else if m1[i].Sup(m2[i]) {
            y=true
            break
        }
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
