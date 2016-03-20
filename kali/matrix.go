// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
//#### Types #################################################################//
type Matrix []Vector
//#### Cat ###################################################################//
func (m1 *Matrix) Cat(m2 Matrix,d int) {
    var i int
    switch d {
        case 1:
            for i=range m2 {
                (*m1)=append((*m1),m2[i].Copy())
            }
        case 2:
            for i=range m2 {
                (*m1)[i].Cat(m2[i])
            }
    }
}
//#### CircShift #############################################################//
func (m *Matrix) CircShift(n int) {
    var i int
    for i=range (*m) {
        (*m)[i].CircShift(n)
    }
}
//#### Col ###################################################################//
func (m Matrix) Col(j int) Vector {
    var i int
    var y Vector
    for i=range m {
        y=append(y,m[i][j])
    }
    return y
}
//#### Copy ##################################################################//
func (m Matrix) Copy() Matrix {
    var i int
    var y Matrix
    for i=range m {
        y=append(y,m[i].Copy())
    }
    return y
}
//#### Equal #################################################################//
func (m1 Matrix) Equal(m2 Matrix) bool {
    var i int
    if m1.Size(1)!=m2.Size(1) || m1.Size(2)!=m2.Size(2) {
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
//#### Find ##################################################################//
func (m Matrix) Find(v Vector,d int) int {
    var i int
    switch d {
        case 1:
            for i=range m {
                if m[i].Equal(v) {
                    return i
                }
            }
        case 2:
            for i=range m[0] {
                if m.Col(i).Equal(v) {
                    return i
                }
            }
    }
    return -1
}
//#### Size ##################################################################//
func (m Matrix) Size(d int) int {
    switch d {
        case 1: return len(m)
        case 2: return len(m[0])
    }
    return -1
}
//#### Sub ###################################################################//
func (m Matrix) Sub(rows,cols []int) Matrix {
    var i,j int
    var z Vector
    var y Matrix
    for i=range rows {
        z=Vector{}
        for j=range cols {
            z=append(z,m[rows[i]][cols[j]])
        }
        y=append(y,z.Copy())
    }
    return y
}
//#### T #####################################################################//
func (m Matrix) T() Matrix {
    var j int
    var y Matrix
    for j=range m[0] {
        y=append(y,m.Col(j))
    }
    return y
}