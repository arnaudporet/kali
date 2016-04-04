// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "math"
import "math/rand"
import "sort"
//#### GenArrang #############################################################//
func (v *Vector) GenArrang(deck Vector,k int) {
    var i int
    (*v)=Vector{}
    if len(deck)>0 && k>0 {
        for i=0;i<k;i++ {
            (*v)=append((*v),deck[rand.Intn(len(deck))])
        }
    }
}
//#### GenArrangMat ##########################################################//
func (m *Matrix) GenArrangMat(deck Vector,k,narrang int) {
    var inMat bool
    var i1,i2 int
    var arrang Vector
    (*m)=Matrix{}
    if len(deck)>0 && k>0 && narrang>0 {
        for i1=0;i1<int(math.Min(float64(narrang),math.Pow(float64(len(deck)),float64(k))));i1++ {
            for {
                arrang.GenArrang(deck,k)
                inMat=false
                for i2=range (*m) {
                    if (*m)[i2].Equal(arrang) {
                        inMat=true
                        break
                    }
                }
                if !inMat {
                    (*m)=append((*m),arrang.Copy())
                    break
                }
            }
        }
    }
}
//#### GenCombi ##############################################################//
func (v *Vector) GenCombi(deck Vector,k int) {
    var i int
    var z []int
    (*v)=Vector{}
    if len(deck)>0 && k>0 {
        z=rand.Perm(len(deck))
        for i=range z[:k] {
            (*v)=append((*v),deck[z[i]])
        }
        sort.Float64s(*v)
    }
}
//#### GenCombiMat ###########################################################//
func (m *Matrix) GenCombiMat(deck Vector,k,ncombi int) {
    var inMat bool
    var i1,i2 int
    var combi Vector
    (*m)=Matrix{}
    if len(deck)>0 && k>0 && ncombi>0 {
        for i1=0;i1<int(math.Min(float64(ncombi),math.Gamma(float64(len(deck)+1))/(math.Gamma(float64(k+1))*math.Gamma(float64(len(deck)-k+1)))));i1++ {
            for {
                combi.GenCombi(deck,k)
                inMat=false
                for i2=range (*m) {
                    if (*m)[i2].Equal(combi) {
                        inMat=true
                        break
                    }
                }
                if !inMat  {
                    (*m)=append((*m),combi.Copy())
                    break
                }
            }
        }
    }
}
//#### GenS ##################################################################//
func (m *Matrix) GenS(deck Vector,n int) {
    var i1,i2,i3 int
    var z Vector
    var Z Matrix
    (*m)=Matrix{}
    if len(deck)>0 && n>0 {
        for i1=0;i1<n;i1++ {
            Z=(*m).Copy()
            for i2=0;i2<len(deck)-1;i2++ {
                (*m).Cat(Z,2)
            }
            z=Vector{}
            for i2=range deck {
                for i3=0;i3<int(math.Pow(float64(len(deck)),float64(i1)));i3++ {
                    z=append(z,deck[i2])
                }
            }
            (*m)=append((*m),z.Copy())
        }
    }
}
