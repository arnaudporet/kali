// Copyright (C) 2013-2021 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

// WARNING The functions in the present file do not handle exceptions and
// errors. Instead, they assume that such handling is performed upstream by the
// <do*> top-level functions of kali. Consequently, they should not be used as
// is outside of kali.

package kali
import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "strings"
)
type Attractor struct {
    Name string
    Basin float64
    States Matrix
}
type AttractorSet []Attractor
func ComputeAttractor(x0 Vector,f func(Vector) Vector,b Bullet,nSteps,maxForward,maxTry int,upd string) (Attractor,bool) {
    var (
        skipped,success bool
        i int
        a Attractor
    )
    if upd=="sync" {
        a.States=ReachCycle(x0,f,b)
        success=true
    } else if upd=="async" {
        success=false
        for i=0;i<maxTry;i++ {
            a.States,skipped=GoForward(Walk(x0,f,b,nSteps),f,b,maxForward)
            if !skipped {
                if a.IsTerminal(f,b) {
                    success=true
                    break
                }
            }
        }
    }
    return a,success
}
func ComputeAttractorSet(S Matrix,f func(Vector) Vector,b Bullet,nSteps,maxForward,maxTry int,upd,setting string,refSet AttractorSet) (AttractorSet,bool) {
    var (
        success bool
        n,i,inA int
        name string
        a Attractor
        A AttractorSet
    )
    if setting=="physio" {
        name="a_physio"
    } else if setting=="patho" {
        name="a_patho"
    }
    n=0
    for i=range S {
        a,success=ComputeAttractor(S[i],f,b,nSteps,maxForward,maxTry,upd)
        if success {
            n+=1
            inA=A.Find(a)
            if inA!=-1 {
                A[inA].Basin+=1
            } else {
                a.Basin=1
                A=append(A,a.Copy())
            }
        }
    }
    if len(A)!=0 {
        success=true
        for i=range A {
            A[i].Basin=100*A[i].Basin/float64(n)
        }
        A=A.Sort().SetNames(name,refSet)
    } else {
        success=false
    }
    return A,success
}
func (a Attractor) Copy() Attractor {
    var (
        y Attractor
    )
    y.Name=a.Name
    y.Basin=a.Basin
    y.States=a.States.Copy()
    return y
}
func (A AttractorSet) Copy() AttractorSet {
    var (
        i int
        y AttractorSet
    )
    y=make(AttractorSet,len(A))
    for i=range A {
        y[i]=A[i].Copy()
    }
    return y
}
func (A1 AttractorSet) Cover(A2 AttractorSet) Vector {
    var (
        i,in2 int
        y Vector
    )
    y=make(Vector,len(A1))
    for i=range A1 {
        in2=A2.Find(A1[i])
        if in2!=-1 {
            y[i]=A2[in2].Basin
        } else {
            y[i]=0
        }
    }
    return y
}
func (A AttractorSet) Find(a Attractor) int {
    var (
        y,i int
    )
    y=-1
    for i=range A {
        if A[i].States.Eq(a.States) {
            y=i
            break
        }
    }
    return y
}
func (A AttractorSet) GetNames() []string {
    var (
        i int
        y []string
    )
    y=make([]string,len(A))
    for i=range A {
        y[i]=A[i].Name
    }
    return y
}
func (A AttractorSet) GetVersus() AttractorSet {
    var (
        i int
        Aversus AttractorSet
    )
    for i=range A {
        if strings.Contains(A[i].Name,"patho") {
            Aversus=append(Aversus,A[i].Copy())
        }
    }
    return Aversus
}
func (a Attractor) IsTerminal(f func(Vector) Vector,b Bullet) bool {
    var (
        y,skipped bool
        maxForward,i int
        forward Matrix
    )
    y=true
    maxForward=len(a.States)
    for i=range a.States {
        forward,skipped=GoForward(a.States[i],f,b,maxForward)
        if skipped || !forward.Eq(a.States) {
            y=false
            break
        }
    }
    return y
}
func LoadAttractorSet(setting string) AttractorSet {
    var (
        i,j int
        n,m int64
        fileName string
        lines [][]string
        A AttractorSet
        file *os.File
        reader *csv.Reader
    )
    if setting=="physio" {
        fileName="A_physio.csv"
    } else if setting=="patho" {
        fileName="A_patho.csv"
    } else if setting=="versus" {
        fileName="A_versus.csv"
    }
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
    n,_=strconv.ParseInt(lines[0][0],10,0)
    m,_=strconv.ParseInt(lines[0][1],10,0)
    A=make(AttractorSet,int(n))
    for i=range A {
        A[i].Name=lines[1+i*(int(m)+1)][0]
        A[i].Basin,_=strconv.ParseFloat(lines[1+i*(int(m)+1)][1],64)
        A[i].States=make(Matrix,int(m))
        for j=range A[i].States {
            A[i].States[j]=StrToVect(lines[2+j+i*(int(m)+1)])
        }
        A[i].States=A[i].States.T()
    }
    return A
}
func (A AttractorSet) Report(nodes []string,setting string,quiet bool) {
    var (
        nPoint,nCycle,i,j int
        name,sep,report string
        aligned,lines []string
        states [][]string
        file *os.File
    )
    if setting=="physio" {
        name="A_physio"
    } else if setting=="patho" {
        name="A_patho"
    } else if setting=="versus" {
        name="A_versus"
    }
    nPoint=0
    nCycle=0
    sep=strings.Repeat("-",80)
    aligned=Align(nodes," ")
    lines=append(lines,name+"={"+strings.Join(A.GetNames(),",")+"}")
    lines=append(lines,sep)
    for i=range A {
        if len(A[i].States)==1 {
            nPoint+=1
        } else if len(A[i].States)>1 {
            nCycle+=1
        }
        lines=append(lines,"Name: "+A[i].Name)
        lines=append(lines,"Basin: "+strconv.FormatFloat(A[i].Basin,'f',-1,64)+"%")
        lines=append(lines,"States:")
        states=A[i].States.T().ToStr()
        for j=range states {
            lines=append(lines,"    "+aligned[j]+" "+strings.Join(states[j]," "))
        }
        lines=append(lines,sep)
    }
    lines=append(lines,"Found attractors: "+strconv.FormatInt(int64(len(A)),10))
    lines=append(lines,"    points: "+strconv.FormatInt(int64(nPoint),10))
    lines=append(lines,"    cycles: "+strconv.FormatInt(int64(nCycle),10))
    report=strings.Join(lines,"\n")
    file,_=os.Create(name+".txt")
    _,_=file.WriteString(report+"\n")
    file.Close()
    A.Save(len(nodes),setting)
    if !quiet {
        fmt.Println(report)
    }
}
func (A AttractorSet) Save(n int,setting string) {
    var (
        i,j int
        fileName string
        lines [][]string
        states Matrix
        file *os.File
        writer *csv.Writer
    )
    if setting=="physio" {
        fileName="A_physio.csv"
    } else if setting=="patho" {
        fileName="A_patho.csv"
    } else if setting=="versus" {
        fileName="A_versus.csv"
    }
    lines=append(lines,[]string{strconv.FormatInt(int64(len(A)),10),strconv.FormatInt(int64(n),10)})
    for i=range A {
        lines=append(lines,[]string{A[i].Name,strconv.FormatFloat(A[i].Basin,'f',-1,64)})
        states=A[i].States.T()
        for j=range states {
            lines=append(lines,states[j].ToStr())
        }
    }
    file,_=os.Create(fileName)
    writer=csv.NewWriter(file)
    writer.Comma=','
    writer.UseCRLF=false
    _=writer.WriteAll(lines)
    file.Close()
}
func (A AttractorSet) SetNames(name string,refSet AttractorSet) AttractorSet {
    var (
        k,i,inRef int
        y AttractorSet
    )
    y=A.Copy()
    k=1
    for i=range y {
        inRef=refSet.Find(y[i])
        if inRef!=-1 {
            y[i].Name=refSet[inRef].Name
        } else {
            y[i].Name=name+strconv.FormatInt(int64(k),10)
            k+=1
        }
    }
    return y
}
func (A AttractorSet) Sort() AttractorSet {
    var (
        rePass bool
        i int
        y AttractorSet
    )
    y=A.Copy()
    for {
        rePass=false
        for i=0;i<len(y)-1;i++ {
            if y[i].States.Sup(y[i+1].States) {
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
func (A AttractorSet) Swap(i,j int) AttractorSet {
    var (
        y AttractorSet
    )
    y=A.Copy()
    y[i]=A[j].Copy()
    y[j]=A[i].Copy()
    return y
}
