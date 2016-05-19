// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "math"
import "math/rand"
import "sort"
//#### GenArrang #############################################################//
func GenArrang(deck Vector,k int) Vector {
    var i int
    var arrang Vector
    if len(deck)>0 && k>0 {
        for i=0;i<k;i++ {
            arrang=append(arrang,deck[rand.Intn(len(deck))])
        }
    }
    return arrang
}
//#### GenArrangMat ##########################################################//
func GenArrangMat(deck Vector,k,narrang int) Matrix {
    var inArrang bool
    var i1,i2 int
    var arrang Vector
    var Arrang Matrix
    if len(deck)>0 && k>0 && narrang>0 {
        for i1=0;i1<int(math.Min(float64(narrang),math.Pow(float64(len(deck)),float64(k))));i1++ {
            for {
                arrang=GenArrang(deck,k)
                inArrang=false
                for i2=range Arrang {
                    if Arrang[i2].Equal(arrang) {
                        inArrang=true
                        break
                    }
                }
                if !inArrang {
                    Arrang=append(Arrang,arrang.Copy())
                    break
                }
            }
        }
    }
    return Arrang
}
//#### GenCombi ##############################################################//
func GenCombi(deck Vector,k int) Vector {
    var i int
    var z []int
    var combi Vector
    if len(deck)>0 && k>0 {
        z=rand.Perm(len(deck))
        for i=0;i<k;i++ {
            combi=append(combi,deck[z[i]])
        }
        sort.Float64s(combi)
    }
    return combi
}
//#### GenCombiMat ###########################################################//
func GenCombiMat(deck Vector,k,ncombi int) Matrix {
    var inCombi bool
    var i1,i2 int
    var combi Vector
    var Combi Matrix
    if len(deck)>0 && k>0 && ncombi>0 {
        for i1=0;i1<int(math.Min(float64(ncombi),math.Gamma(float64(len(deck)+1))/(math.Gamma(float64(k+1))*math.Gamma(float64(len(deck)-k+1)))));i1++ {
            for {
                combi=GenCombi(deck,k)
                inCombi=false
                for i2=range Combi {
                    if Combi[i2].Equal(combi) {
                        inCombi=true
                        break
                    }
                }
                if !inCombi {
                    Combi=append(Combi,combi.Copy())
                    break
                }
            }
        }
    }
    return Combi
}
//#### GenS ##################################################################//
func GenS(deck Vector,n int) Matrix {
    var i1,i2,i3 int
    var z Vector
    var Z,S Matrix
    if len(deck)>0 && n>0 {
        for i1=0;i1<n;i1++ {
            Z=S.Copy()
            for i2=0;i2<len(deck)-1;i2++ {
                S.Cat(Z.Copy(),2)
            }
            z=Vector{}
            for i2=range deck {
                for i3=0;i3<int(math.Pow(float64(len(deck)),float64(i1)));i3++ {
                    z=append(z,deck[i2])
                }
            }
            S=append(S,z.Copy())
        }
    }
    return S
}
