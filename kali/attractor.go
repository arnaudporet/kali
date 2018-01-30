// Copyright (C) 2013-2018 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
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
func ComputeAttractor(f func(Vector) Vector,x0 Vector,b Bullet,kmax,sync,maxfwd int) (Attractor,bool) {
    var (
        skipped bool
        a Attractor
    )
    if sync==1 {
        a.States=ReachCycle(f,x0,b)
    } else if sync==0 {
        for {
            a.States,skipped=GoForward(f,Walk(f,x0,b,kmax),b,maxfwd)
            if skipped {
                return Attractor{},true
            }
            if a.IsTerminal(f,b) {
                break
            }
        }
    }
    return a,false
}
func ComputeAttractorSet(f func(Vector) Vector,S Matrix,b Bullet,kmax,setting,sync,maxfwd int,RefSet AttractorSet) (AttractorSet,bool) {
    var (
        skipped bool
        i,inA int
        name string
        a Attractor
        A AttractorSet
    )
    if setting==0 {
        name="a_physio"
    } else if setting==1 {
        name="a_patho"
    }
    for i=range S {
        a,skipped=ComputeAttractor(f,S[i],b,kmax,sync,maxfwd)
        if skipped {
            return AttractorSet{},true
        }
        inA=A.Find(a)
        if inA!=-1 {
            A[inA].Basin+=1.0
        } else {
            a.Basin=1.0
            A=append(A,a.Copy())
        }
    }
    for i=range A {
        A[i].Basin=100.0*A[i].Basin/float64(len(S))
    }
    return A.Sort().SetNames(name,RefSet),false
}
func (a Attractor) Copy() Attractor {
    var y Attractor
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
            y[i]=0.0
        }
    }
    return y
}
func (A AttractorSet) Find(a Attractor) int {
    var i int
    for i=range A {
        if A[i].States.Eq(a.States) {
            return i
        }
    }
    return -1
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
func (Apatho AttractorSet) GetVersus() AttractorSet {
    var (
        i int
        Aversus AttractorSet
    )
    for i=range Apatho {
        if strings.Contains(Apatho[i].Name,"patho") {
            Aversus=append(Aversus,Apatho[i].Copy())
        }
    }
    return Aversus
}
func (a Attractor) IsTerminal(f func(Vector) Vector,b Bullet) bool {
    // asynchronous only
    var (
        i int
        fwd Matrix
    )
    for i=range a.States {
        fwd,_=GoForward(f,a.States[i],b,-9)
        if !fwd.Eq(a.States) {
            return false
        }
    }
    return true
}
func LoadAttractorSet(setting int) AttractorSet {
    var (
        i1,i2 int
        n int64
        filename string
        s []string
        A AttractorSet
        reader *csv.Reader
        file *os.File
    )
    if setting==0 {
        filename="A_physio.csv"
    } else if setting==1 {
        filename="A_patho.csv"
    } else if setting==2 {
        filename="A_versus.csv"
    }
    file,_=os.Open(filename)
    reader=csv.NewReader(file)
    reader.Comma=','
    reader.Comment=0
    reader.FieldsPerRecord=-1
    reader.LazyQuotes=false
    reader.TrimLeadingSpace=true
    s,_=reader.Read()
    n,_=strconv.ParseInt(s[0],10,0)
    A=make(AttractorSet,int(n))
    s,_=reader.Read()
    n,_=strconv.ParseInt(s[0],10,0)
    for i1=range A {
        A[i1].States=make(Matrix,int(n))
        s,_=reader.Read()
        A[i1].Name=s[0]
        s,_=reader.Read()
        A[i1].Basin,_=strconv.ParseFloat(s[0],64)
        for i2=range A[i1].States {
            s,_=reader.Read()
            A[i1].States[i2]=StrToVect(s)
        }
        A[i1].States=A[i1].States.T()
    }
    file.Close()
    return A
}
func (A AttractorSet) Report(nodes []string,setting int) {
    var (
        npoint,ncycle,i1,i2 int
        name,report string
        aligned []string
        states [][]string
        file *os.File
    )
    if setting==0 {
        name="A_physio"
    } else if setting==1 {
        name="A_patho"
    } else if setting==2 {
        name="A_versus"
    }
    aligned=Align(nodes," ")
    npoint=0
    ncycle=0
    report=name+"={"+strings.Join(A.GetNames(),",")+"}\n"+strings.Repeat("-",80)+"\n"
    for i1=range A {
        if len(A[i1].States)==1 {
            npoint+=1
        } else {
            ncycle+=1
        }
        report+="Name: "+A[i1].Name+"\nBasin: "+strconv.FormatFloat(A[i1].Basin,'f',-1,64)+"%\nMatrix:\n"
        states=A[i1].States.T().ToStr()
        for i2=range states {
            report+="    "+aligned[i2]+" "+strings.Join(states[i2]," ")+"\n"
        }
        report+=strings.Repeat("-",80)+"\n"
    }
    report+="Found attractors: "+strconv.FormatInt(int64(len(A)),10)+"\n    points: "+strconv.FormatInt(int64(npoint),10)+"\n    cycles: "+strconv.FormatInt(int64(ncycle),10)
    fmt.Println("\n"+report)
    file,_=os.Create(name+".txt")
    file.WriteString(report+"\n")
    file.Close()
    A.Save(setting)
}
func (A AttractorSet) Save(setting int) {
    var (
        i1,i2 int
        filename string
        s [][]string
        states Matrix
        writer *csv.Writer
        file *os.File
    )
    if setting==0 {
        filename="A_physio.csv"
    } else if setting==1 {
        filename="A_patho.csv"
    } else if setting==2 {
        filename="A_versus.csv"
    }
    s=append(s,[]string{strconv.FormatInt(int64(len(A)),10)})
    if len(A)==0 {
        s=append(s,[]string{"0"})
    } else {
        s=append(s,[]string{strconv.FormatInt(int64(len(A[0].States[0])),10)})
    }
    for i1=range A {
        s=append(s,[]string{A[i1].Name})
        s=append(s,[]string{strconv.FormatFloat(A[i1].Basin,'f',-1,64)})
        states=A[i1].States.T()
        for i2=range states {
            s=append(s,states[i2].ToStr())
        }
    }
    file,_=os.Create(filename)
    writer=csv.NewWriter(file)
    writer.Comma=','
    writer.UseCRLF=false
    writer.WriteAll(s)
    file.Close()
}
func (A AttractorSet) SetNames(name string,RefSet AttractorSet) AttractorSet {
    var (
        k,i,inRef int
        y AttractorSet
    )
    y=A.Copy()
    k=1
    for i=range y {
        inRef=RefSet.Find(y[i])
        if inRef!=-1 {
            y[i].Name=RefSet[inRef].Name
        } else {
            y[i].Name=name+strconv.FormatInt(int64(k),10)
            k+=1
        }
    }
    return y
}
func (A AttractorSet) Sort() AttractorSet {
    var (
        repass bool
        i int
        y AttractorSet
    )
    y=A.Copy()
    for {
        repass=false
        for i=0;i<len(y)-1;i++ {
            if len(y[i].States)>len(y[i+1].States) || (len(y[i].States)==len(y[i+1].States) && y[i].States[0].Sup(y[i+1].States[0])) {
                y=y.Swap(i,i+1)
                repass=true
            }
        }
        if !repass {
            return y
        }
    }
}
func (A AttractorSet) Swap(i,j int) AttractorSet {
    var y AttractorSet
    y=A.Copy()
    y[i]=A[j].Copy()
    y[j]=A[i].Copy()
    return y
}
