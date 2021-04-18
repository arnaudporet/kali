// Copyright (C) 2013-2021 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

// WARNING The functions in the present file do not handle exceptions and
// errors. Instead, they assume that such handling is performed upstream by the
// <do*> top-level functions of kali. Consequently, they should not be used as
// is outside of kali.

package kali
import (
    "fmt"
    "math"
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
func ComputeTherapeuticBullets(S,Targ,Moda Matrix,f func(Vector) Vector,nSteps,maxTry int,th float64,Aphysio,Apatho,Aversus AttractorSet,upd string) (BulletSet,bool) {
    var (
        success bool
        i,j,maxForward int
        sizes []float64
        b Bullet
        Aphyrsus,Atest AttractorSet
        Btherap BulletSet
    )
    Aphyrsus=append(Aphysio.Copy(),Aversus.Copy()...)
    for i=range Aphyrsus {
        sizes=append(sizes,float64(len(Aphyrsus[i].States)))
    }
    maxForward=int(math.Round(Max(sizes...)))
    b.Gain=make(Vector,2)
    b.Gain[0]=Aphysio.Cover(Apatho).Sum()
    for i=range Targ {
        b.Targ=Targ[i].Copy()
        for j=range Moda {
            b.Moda=Moda[j].Copy()
            Atest,success=ComputeAttractorSet(S,f,b,nSteps,maxForward,maxTry,upd,"patho",Aphysio)
            if success {
                b.Gain[1]=Aphysio.Cover(Atest).Sum()
                if b.IsTherapeutic(Atest,Aversus,th) {
                    b.Cover=Aphyrsus.Cover(Atest)
                    Btherap=append(Btherap,b.Copy())
                }
            }
        }
    }
    if len(Btherap)!=0 {
        success=true
        Btherap=Btherap.Sort()
    } else {
        success=false
    }
    return Btherap,success
}
func (b Bullet) Copy() Bullet {
    var (
        y Bullet
    )
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
func (b Bullet) IsTherapeutic(Atest,Aversus AttractorSet,th float64) bool {
    var (
        y bool
        i int
    )
    if (b.Gain[1]-b.Gain[0])<th {
        y=false
    } else {
        y=true
        for i=range Atest {
            if strings.Contains(Atest[i].Name,"patho") && Aversus.Find(Atest[i])==-1 {
                y=false
                break
            }
        }
    }
    return y
}
func (B BulletSet) Report(nodes,physioNames,pathoNames []string,quiet bool) {
    var (
        i,j int
        sep,report string
        lines,bullet []string
        file *os.File
    )
    sep=strings.Repeat("-",80)
    lines=append(lines,"B_therap")
    lines=append(lines,sep)
    for i=range B {
        bullet=make([]string,len(B[i].Targ))
        for j=range bullet {
            bullet[j]=nodes[int(math.Round(B[i].Targ[j]))]+"["+strconv.FormatFloat(B[i].Moda[j],'f',-1,64)+"]"
        }
        lines=append(lines,"Bullet: "+strings.Join(bullet," "))
        lines=append(lines,"Gain: "+strconv.FormatFloat(B[i].Gain[0],'f',-1,64)+"% --> "+strconv.FormatFloat(B[i].Gain[1],'f',-1,64)+"%")
        lines=append(lines,"Physiological basins:")
        for j=range physioNames {
            lines=append(lines,"    "+physioNames[j]+": "+strconv.FormatFloat(B[i].Cover[j],'f',-1,64)+"%")
        }
        lines=append(lines,"Pathological basins:")
        for j=range pathoNames {
            lines=append(lines,"    "+pathoNames[j]+": "+strconv.FormatFloat(B[i].Cover[j+len(physioNames)],'f',-1,64)+"%")
        }
        lines=append(lines,sep)
    }
    lines=append(lines,"Found therapeutic bullets: "+strconv.FormatInt(int64(len(B)),10))
    report=strings.Join(lines,"\n")
    file,_=os.Create("B_therap.txt")
    _,_=file.WriteString(report+"\n")
    file.Close()
    if !quiet {
        fmt.Println(report)
    }
}
func (B BulletSet) Sort() BulletSet {
    var (
        rePass bool
        i int
        y BulletSet
    )
    y=B.Copy()
    for {
        rePass=false
        for i=0;i<len(y)-1;i++ {
            if y[i].Targ.Sup(y[i+1].Targ) || (y[i].Targ.Eq(y[i+1].Targ) && y[i].Moda.Sup(y[i+1].Moda)) {
                y=y.Swap(i,i+1)
                rePass=true
            }
        }
        if !rePass {
            break
        }
    }
    return y
}
func (B BulletSet) Swap(i,j int) BulletSet {
    var (
        y BulletSet
    )
    y=B.Copy()
    y[i]=B[j].Copy()
    y[j]=B[i].Copy()
    return y
}
