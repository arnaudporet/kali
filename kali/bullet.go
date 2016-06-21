// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "os"
    "strconv"
    "strings"
)
type Bullet struct {
    Targ Vector
    Moda Vector
    Gain Vector
    Cover Vector
}
type BulletSet []Bullet
func (b Bullet) Assess(Atest,Aversus AttractorSet,threshold int) bool {
    var i int
    if b.Gain[1]-b.Gain[0]>=float64(threshold) {
        for i=range Atest {
            if strings.Contains(Atest[i].Name,"patho") && Aversus.Find(Atest[i])==-1 {
                return false
            }
        }
        return true
    } else {
        return false
    }
}
func ComputeTherapeuticBullets(fpatho func(Matrix,int) Vector,S,Targ,Moda Matrix,Aphysio,Apatho,Aversus AttractorSet,threshold int) BulletSet {
    var (
        i1,i2 int
        Atest AttractorSet
        b Bullet
        Btherap BulletSet
    )
    b.Gain=make(Vector,2)
    b.Gain[0]=Aphysio.Cover(Apatho).Sum()
    for i1=range Targ {
        for i2=range Moda {
            b.Targ=Targ[i1].Copy()
            b.Moda=Moda[i2].Copy()
            Atest=ComputeAttractorSet(fpatho,S,b,Aphysio,1)
            b.Gain[1]=Aphysio.Cover(Atest).Sum()
            if b.Assess(Atest,Aversus,threshold) {
                b.Cover=Aphysio.Cat(Aversus).Cover(Atest)
                Btherap=append(Btherap,b.Copy())
            }
        }
    }
    return Btherap.Sort()
}
func (b Bullet) Copy() Bullet {
    var y Bullet
    y.Targ=b.Targ.Copy()
    y.Moda=b.Moda.Copy()
    y.Gain=b.Gain.Copy()
    y.Cover=b.Cover.Copy()
    return y
}
func (B BulletSet) Copy() BulletSet {
    var (
        i int
        y BulletSet
    )
    y=make(BulletSet,len(B))
    for i=range B {
        y[i]=B[i].Copy()
    }
    return y
}
func (Btherap BulletSet) Report(nodes,physionames,pathonames []string) {
    var (
        i1,i2 int
        report string
        s []string
        file *os.File
    )
    report="B_therap\n"+strings.Repeat("-",80)+"\n"
    for i1=range Btherap {
        report+="Bullet: "
        s=make([]string,len(Btherap[i1].Targ))
        for i2=range Btherap[i1].Targ {
            s[i2]=nodes[int(Btherap[i1].Targ[i2])]+"["+strconv.FormatFloat(Btherap[i1].Moda[i2],'f',-1,64)+"]"
        }
        report+=strings.Join(s," ")+"\nGain: "+strconv.FormatFloat(Btherap[i1].Gain[0],'f',-1,64)+"% --> "+strconv.FormatFloat(Btherap[i1].Gain[1],'f',-1,64)+"%\nPhysiological basins:\n"
        for i2=range physionames {
            report+="    "+physionames[i2]+": "+strconv.FormatFloat(Btherap[i1].Cover[i2],'f',-1,64)+"%\n"
        }
        report+="Pathological basins:\n"
        for i2=range pathonames {
            report+="    "+pathonames[i2]+": "+strconv.FormatFloat(Btherap[i1].Cover[len(physionames)+i2],'f',-1,64)+"%\n"
        }
        report+=strings.Repeat("-",80)+"\n"
    }
    report+="Found therapeutic bullets: "+strconv.FormatInt(int64(len(Btherap)),10)
    fmt.Println("\n"+report)
    file,_=os.Create("B_therap.txt")
    file.WriteString(report+"\n")
    file.Close()
}
func (B BulletSet) Sort() BulletSet {
    var (
        repass bool
        i1,i2 int
        y BulletSet
    )
    y=B.Copy()
    for {
        repass=false
        for i1=0;i1<len(y)-1;i1++ {
            if y[i1].Targ.Equal(y[i1+1].Targ) {
                for i2=range y[i1].Moda {
                    if y[i1].Moda[i2]>y[i1+1].Moda[i2] {
                        y=y.Swap(i1,i1+1)
                        repass=true
                        break
                    } else if y[i1].Moda[i2]<y[i1+1].Moda[i2] {
                        break
                    }
                }
            } else {
                for i2=range y[i1].Targ {
                    if y[i1].Targ[i2]>y[i1+1].Targ[i2] {
                        y=y.Swap(i1,i1+1)
                        repass=true
                        break
                    } else if y[i1].Targ[i2]<y[i1+1].Targ[i2] {
                        break
                    }
                }
            }
        }
        if !repass {
            return y
        }
    }
}
func (B BulletSet) Swap(i,j int) BulletSet {
    var y BulletSet
    y=B.Copy()
    y[i]=B[j].Copy()
    y[j]=B[i].Copy()
    return y
}
