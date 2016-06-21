// Copyright (C) 2013-2016 Arnaud Poret
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
    Matrix Matrix
}
type AttractorSet []Attractor
func (A1 AttractorSet) Cat(A2 AttractorSet) AttractorSet {
    var (
        i int
        y AttractorSet
    )
    y=A1.Copy()
    for i=range A2 {
        y=append(y,A2[i].Copy())
    }
    return y
}
func ComputeAttractor(f func(Matrix,int) Vector,x0 Vector,b Bullet) Attractor {
    var (
        k,j int
        z Vector
        x Matrix
        a Attractor
    )
    x=x0.ToMatrix(2)
    k=0
    for {
        z=f(x,k).Shoot(b)
        j=x.Find(z,2)
        if j!=-1 {
            a.Matrix=x.Cols(Range(j,k+1))
            return a.Sort()
        } else {
            x=x.Append(z,2)
            k+=1
        }
    }
}
func ComputeAttractorSet(f func(Matrix,int) Vector,S Matrix,b Bullet,Ref AttractorSet,setting int) AttractorSet {
    var (
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
        a=ComputeAttractor(f,S[i],b)
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
    return A.Sort().SetNames(Ref,name)
}
func (a Attractor) Copy() Attractor {
    var y Attractor
    y.Name=a.Name
    y.Basin=a.Basin
    y.Matrix=a.Matrix.Copy()
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
        if A[i].Matrix.Equal(a.Matrix) {
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
func LoadAttractorSet(setting int) AttractorSet {
    var (
        i1,i2 int
        n int64
        filename string
        s []string
        a Attractor
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
    a.Matrix=make(Matrix,int(n))
    for i1=range A {
        s,_=reader.Read()
        a.Name=s[0]
        s,_=reader.Read()
        a.Basin,_=strconv.ParseFloat(s[0],64)
        for i2=range a.Matrix {
            s,_=reader.Read()
            a.Matrix[i2]=StringToVector(s)
        }
        A[i1]=a.Copy()
    }
    file.Close()
    return A
}
func (A AttractorSet) Report(nodes []string,setting int) {
    var (
        npoint,ncycle,i1,i2 int
        name,report string
        aligned []string
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
        if A[i1].Matrix.Size(2)==1 {
            npoint+=1
        } else {
            ncycle+=1
        }
        report+="Name: "+A[i1].Name+"\nBasin: "+strconv.FormatFloat(A[i1].Basin,'f',-1,64)+"%\nMatrix:\n"
        for i2=range A[i1].Matrix {
            report+="    "+aligned[i2]+" "+strings.Join(A[i1].Matrix[i2].ToString()," ")+"\n"
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
        s=append(s,[]string{strconv.FormatInt(int64(len(A[0].Matrix)),10)})
    }
    for i1=range A {
        s=append(s,[]string{A[i1].Name})
        s=append(s,[]string{strconv.FormatFloat(A[i1].Basin,'f',-1,64)})
        for i2=range A[i1].Matrix {
            s=append(s,A[i1].Matrix[i2].ToString())
        }
    }
    file,_=os.Create(filename)
    writer=csv.NewWriter(file)
    writer.Comma=','
    writer.UseCRLF=false
    writer.WriteAll(s)
    file.Close()
}
func (A AttractorSet) SetNames(Ref AttractorSet,name string) AttractorSet {
    var (
        k,i,inRef int
        y AttractorSet
    )
    y=A.Copy()
    k=1
    for i=range A {
        inRef=Ref.Find(A[i])
        if inRef!=-1 {
            y[i].Name=Ref[inRef].Name
        } else {
            y[i].Name=name+strconv.FormatInt(int64(k),10)
            k+=1
        }
    }
    return y
}
func (a Attractor) Sort() Attractor {
    var (
        i int
        jmin Vector
        y Attractor
    )
    y=a.Copy()
    jmin=IntToVector(Range(0,y.Matrix.Size(2)))
    for i=range y.Matrix {
        jmin=jmin.Sub(y.Matrix[i].Sub(jmin.ToInt()).MinPos())
        if len(jmin)==1 {
            y.Matrix=y.Matrix.CircShift(int(jmin[0]),2)
            break
        }
    }
    return y
}
func (A AttractorSet) Sort() AttractorSet {
    var (
        repass bool
        i1,i2 int
        y AttractorSet
    )
    y=A.Copy()
    for {
        repass=false
        for i1=0;i1<len(y)-1;i1++ {
            if y[i1].Matrix.Size(2)>y[i1+1].Matrix.Size(2) {
                y=y.Swap(i1,i1+1)
                repass=true
            } else if y[i1].Matrix.Size(2)==y[i1+1].Matrix.Size(2) {
                for i2=range y[i1].Matrix {
                    if y[i1].Matrix[i2][0]>y[i1+1].Matrix[i2][0] {
                        y=y.Swap(i1,i1+1)
                        repass=true
                        break
                    } else if y[i1].Matrix[i2][0]<y[i1+1].Matrix[i2][0] {
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
func (A AttractorSet) Swap(i,j int) AttractorSet {
    var y AttractorSet
    y=A.Copy()
    y[i]=A[j].Copy()
    y[j]=A[i].Copy()
    return y
}
