// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "os"
import "strconv"
import "strings"
//#### Types #################################################################//
type Bullet struct {
    Targ Vector
    Moda Vector
    Gain Vector
    Cover Vector
}
type Bset []Bullet
//#### Assess ################################################################//
func (b Bullet) Assess(Atest,Aversus Aset) bool {
    var i int
    if b.Gain[1]>b.Gain[0]+1.0 {// +1% to avoid round-off errors
        for i=range Atest {
            if Atest[i].IsPatho() && Aversus.Find(Atest[i])<0 {
                return false
            }
        }
        return true
    } else {
        return false
    }
}
//#### Compute ###############################################################//
func (B *Bset) Compute(fpatho func(Matrix,int) Vector,S,Targ,Moda Matrix,Aphysio,Apatho,Aversus Aset) {
    var i1,i2 int
    var b Bullet
    var Atest Aset
    (*B)=Bset{}
    b.Gain=make(Vector,2)
    b.Gain[0]=Aphysio.Covers(Apatho).Sum()
    for i1=range Targ {
        for i2=range Moda {
            b.Targ=Targ[i1].Copy()
            b.Moda=Moda[i2].Copy()
            Atest.Compute(fpatho,S,b,Aphysio,1)
            b.Gain[1]=Aphysio.Covers(Atest).Sum()
            if b.Assess(Atest,Aversus) {
                b.Cover=Union(Aphysio,Aversus).Covers(Atest)
                (*B)=append((*B),b.Copy())
            }
        }
    }
    (*B).Sort()
}
//#### Copy ##################################################################//
func (b Bullet) Copy() Bullet {
    var y Bullet
    y.Targ=b.Targ.Copy()
    y.Moda=b.Moda.Copy()
    y.Gain=b.Gain.Copy()
    y.Cover=b.Cover.Copy()
    return y
}
//#### Report ################################################################//
func (B Bset) Report(nodes []string,Aphysio,Aversus Aset) {
    var i1,i2,save int
    var report string
    var z []string
    var file *os.File
    report=""
    for i1=range B {
        report+="Bullet: "
        z=[]string{}
        for i2=range B[i1].Targ {
            z=append(z,nodes[int(B[i1].Targ[i2])]+"["+strconv.FormatFloat(B[i1].Moda[i2],'f',-1,64)+"]")
        }
        report+=strings.Join(z," ")+"\nGain: "+strconv.FormatFloat(B[i1].Gain[0],'f',-1,64)+"% --> "+strconv.FormatFloat(B[i1].Gain[1],'f',-1,64)+"%\nPhysiological basins:\n"
        for i2=range Aphysio {
            report+="    "+Aphysio[i2].Name+": "+strconv.FormatFloat(B[i1].Cover[i2],'f',-1,64)+"%\n"
        }
        report+="Pathological basins:\n"
        for i2=range Aversus {
            report+="    "+Aversus[i2].Name+": "+strconv.FormatFloat(B[i1].Cover[len(Aphysio)+i2],'f',-1,64)+"%\n"
        }
        report+=strings.Repeat("-",80)+"\n"
    }
    report+="Found therapeutic bullets: "+strconv.FormatInt(int64(len(B)),10)+"\n"
    fmt.Println("\n"+report)
    save=int(Prompt("Save? (optional) [0/1] ",Vector{0.0,1.0}))
    if save==1 {
        file,_=os.Create("B_therap.txt")
        file.WriteString(report)
        file.Close()
        fmt.Println("\nINFO: report saved as B_therap.txt")
    }
}
//#### Sort ##################################################################//
func (B *Bset) Sort() {
    var repass bool
    var i1,i2 int
    for {
        repass=false
        for i1=0;i1<len(*B)-1;i1++ {
            if (*B)[i1].Targ.Equal((*B)[i1+1].Targ) {
                for i2=range (*B)[i1].Moda {
                    if (*B)[i1].Moda[i2]>(*B)[i1+1].Moda[i2] {
                        (*B).Swap(i1,i1+1)
                        repass=true
                        break
                    } else if (*B)[i1].Moda[i2]<(*B)[i1+1].Moda[i2] {
                        break
                    }
                }
            } else {
                for i2=range (*B)[i1].Targ {
                    if (*B)[i1].Targ[i2]>(*B)[i1+1].Targ[i2] {
                        (*B).Swap(i1,i1+1)
                        repass=true
                        break
                    } else if (*B)[i1].Targ[i2]<(*B)[i1+1].Targ[i2] {
                        break
                    }
                }
            }
        }
        if !repass {
            break
        }
    }
}
//#### Swap ##################################################################//
func (B *Bset) Swap(i1,i2 int) {
    var b Bullet
    if len(*B)>0 {
        b=(*B)[i1].Copy()
        (*B)[i1]=(*B)[i2].Copy()
        (*B)[i2]=b
    }
}
