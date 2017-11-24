// Copyright (C) 2013-2017 Arnaud Poret
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
func ComputeTherapeuticBullets(fpatho func(Vector) Vector,S,Targ,Moda Matrix,kmax,threshold,sync int,Aphysio,Apatho,Aversus AttractorSet) BulletSet {
    var (
        skipped bool
        i1,i2,maxfwd int
        sizes []float64
        Aphyrsus,Atest AttractorSet
        b Bullet
        Btherap BulletSet
    )
    Aphyrsus=append(Aphysio.Copy(),Aversus.Copy()...)
    for i1=range Aphyrsus {
        sizes=append(sizes,float64(len(Aphyrsus[i1].States)))
    }
    maxfwd=NearInt(Max(sizes...))
    b.Gain=make(Vector,2)
    b.Gain[0]=Aphysio.Cover(Apatho).Sum()
    for i1=range Targ {
        for i2=range Moda {
            b.Targ=Targ[i1].Copy()
            b.Moda=Moda[i2].Copy()
            Atest,skipped=ComputeAttractorSet(fpatho,S,b,kmax,1,sync,maxfwd,Aphysio)
            if !skipped {
                b.Gain[1]=Aphysio.Cover(Atest).Sum()
                if b.IsTherapeutic(Atest,Aversus,threshold) {
                    b.Cover=Aphyrsus.Cover(Atest)
                    Btherap=append(Btherap,b.Copy())
                }
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
func (b Bullet) IsTherapeutic(Atest,Aversus AttractorSet,threshold int) bool {
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
func (Btherap BulletSet) Report(nodes,physionames,pathonames []string) {
    var (
        i1,i2 int
        report string
        bullet []string
        file *os.File
    )
    report="B_therap\n"+strings.Repeat("-",80)+"\n"
    for i1=range Btherap {
        report+="Bullet: "
        bullet=make([]string,len(Btherap[i1].Targ))
        for i2=range bullet {
            bullet[i2]=nodes[NearInt(Btherap[i1].Targ[i2])]+"["+strconv.FormatFloat(Btherap[i1].Moda[i2],'f',-1,64)+"]"
        }
        report+=strings.Join(bullet," ")+"\nGain: "+strconv.FormatFloat(Btherap[i1].Gain[0],'f',-1,64)+"% --> "+strconv.FormatFloat(Btherap[i1].Gain[1],'f',-1,64)+"%\nPhysiological basins:\n"
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
        i int
        y BulletSet
    )
    y=B.Copy()
    for {
        repass=false
        for i=0;i<len(y)-1;i++ {
            if y[i].Targ.Sup(y[i+1].Targ) || (y[i].Targ.Eq(y[i+1].Targ) && y[i].Moda.Sup(y[i+1].Moda)) {
                y=y.Swap(i,i+1)
                repass=true
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
