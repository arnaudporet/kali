// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "encoding/csv"
import "fmt"
import "os"
import "strconv"
import "strings"
//#### Types #################################################################//
type Attractor struct {
    Name string
    Basin float64
    Mat Matrix
}
type Aset []Attractor
//#### Compute ###############################################################//
func (a *Attractor) Compute(f func(Matrix,int) Vector,x0 Vector,b Bullet) {
    var i,j,k int
    var y Vector
    var x Matrix
    (*a)=Attractor{}
    x=x0.VtoM(2)
    k=0
    for {
        y=f(x,k)
        for i=range b.Targ {
            y[int(b.Targ[i])]=b.Moda[i]
        }
        j=x.Find(y,2)
        if j>-1 {
            (*a).Mat=x.Sub(Range(0,x.Size(1)),Range(j,k+1))
            (*a).Sort()
            break
        } else {
            x.Cat(y.VtoM(2),2)
            k+=1
        }
    }
}
//#### Compute ###############################################################//
func (A *Aset) Compute(f func(Matrix,int) Vector,S Matrix,b Bullet,Ref Aset,setting int) {
    var i,inA int
    var a Attractor
    (*A)=Aset{}
    for i=range S[0] {
        a.Compute(f,S.Col(i),b)
        inA=(*A).Find(a)
        if inA>-1 {
            (*A)[inA].Basin+=1.0
        } else {
            a.Basin=1.0
            (*A)=append((*A),a.Copy())
        }
    }
    for i=range (*A) {
        (*A)[i].Basin=100.0*(*A)[i].Basin/float64(S.Size(2))
    }
    (*A).Sort()
    (*A).Name(Ref,setting)
}
//#### Copy ##################################################################//
func (a Attractor) Copy() Attractor {
    var y Attractor
    y.Name=a.Name
    y.Basin=a.Basin
    y.Mat=a.Mat.Copy()
    return y
}
//#### Covers ################################################################//
func (A1 Aset) Covers(A2 Aset) Vector {
    var i,in2 int
    var y Vector
    for i=range A1 {
        in2=A2.Find(A1[i])
        if in2>-1 {
            y=append(y,A2[in2].Basin)
        } else {
            y=append(y,0.0)
        }
    }
    return y
}
//#### Find ##################################################################//
func (A Aset) Find(a Attractor) int {
    var i int
    for i=range A {
        if A[i].Mat.Equal(a.Mat) {
            return i
        }
    }
    return -1
}
//#### GetNames ##############################################################//
func (A Aset) GetNames() []string {
    var i int
    var names []string
    for i=range A {
        names=append(names,A[i].Name)
    }
    return names
}
//#### IsPatho ###############################################################//
func (a Attractor) IsPatho() bool {
    return strings.Contains(a.Name,"patho")
}
//#### Load ##################################################################//
func (A *Aset) Load(setting int) {
    var i1,i2 int
    var n,m int64
    var filename string
    var s []string
    var a Attractor
    var file *os.File
    var reader *csv.Reader
    switch setting {
        case 0: filename="A_physio.csv"
        case 1: filename="A_patho.csv"
        case 2: filename="A_versus.csv"
    }
    if !Exist(filename) {
        fmt.Println("\nERROR: unable to load "+filename)
    } else {
        (*A)=Aset{}
        file,_=os.Open(filename)
        reader=csv.NewReader(file)
        reader.Comma=','
        reader.Comment=0
        reader.FieldsPerRecord=-1
        reader.LazyQuotes=false
        reader.TrimLeadingSpace=true
        s,_=reader.Read()
        n,_=strconv.ParseInt(s[0],10,0)
        if int(n)>0 {
            s,_=reader.Read()
            m,_=strconv.ParseInt(s[0],10,0)
        }
        for i1=0;i1<int(n);i1++ {
            s,_=reader.Read()
            a.Name=s[0]
            s,_=reader.Read()
            a.Basin,_=strconv.ParseFloat(s[0],64)
            a.Mat=Matrix{}
            for i2=0;i2<int(m);i2++ {
                s,_=reader.Read()
                a.Mat=append(a.Mat,StoV(s))
            }
            (*A)=append((*A),a.Copy())
        }
        file.Close()
        fmt.Println("\nINFO: "+filename+" loaded")
    }
}
//#### Name ##################################################################//
func (A *Aset) Name(Ref Aset,setting int) {
    var i,k,inRef int
    var name string
    switch setting {
        case 0: name="a_physio"
        case 1: name="a_patho"
    }
    k=1
    for i=range (*A) {
        inRef=Ref.Find((*A)[i])
        if inRef>-1 {
            (*A)[i].Name=Ref[inRef].Name
        } else {
            (*A)[i].Name=name+strconv.FormatInt(int64(k),10)
            k+=1
        }
    }
}
//#### Report ################################################################//
func (A Aset) Report(setting int,nodes []string) {
    var npoint,ncycle,i1,i2 int
    var filename,report string
    var nodesfilled []string
    var file *os.File
    nodesfilled=FillToMaxLen(nodes)
    npoint=0
    ncycle=0
    switch setting {
        case 0:
            filename="A_physio.txt"
            report="A_physio={"
        case 1:
            filename="A_patho.txt"
            report="A_patho={"
        case 2:
            filename="A_versus.txt"
            report="A_versus={"
    }
    report+=strings.Join(A.GetNames(),",")+"}\n"+strings.Repeat("-",80)+"\n"
    for i1=range A {
        if A[i1].Mat.Size(2)==1 {
            npoint+=1
        } else {
            ncycle+=1
        }
        report+="Name: "+A[i1].Name+"\nBasin: "+strconv.FormatFloat(A[i1].Basin,'f',-1,64)+"%\nMatrix:\n"
        for i2=range A[i1].Mat {
            report+="    "+nodesfilled[i2]+" "+strings.Join(A[i1].Mat[i2].VtoS()," ")+"\n"
        }
        report+=strings.Repeat("-",80)+"\n"
    }
    report+="Found attractors: "+strconv.FormatInt(int64(len(A)),10)+"\n    points: "+strconv.FormatInt(int64(npoint),10)+"\n    cycles: "+strconv.FormatInt(int64(ncycle),10)
    fmt.Println("\n"+report)
    if Exist(filename) {
        fmt.Println("\nWARNING: "+filename+" will be overwritten")
    }
    file,_=os.Create(filename)
    file.WriteString(report+"\n")
    file.Close()
    fmt.Println("\nINFO: saved as "+filename)
    A.Save(setting)
}
//#### Save ##################################################################//
func (A Aset) Save(setting int) {
    var i1,i2 int
    var filename string
    var s [][]string
    var file *os.File
    var writer *csv.Writer
    s=append(s,[]string{strconv.FormatInt(int64(len(A)),10)})
    if len(A)>0 {
        s=append(s,[]string{strconv.FormatInt(int64(A[0].Mat.Size(1)),10)})
    }
    for i1=range A {
        s=append(s,[]string{A[i1].Name})
        s=append(s,[]string{strconv.FormatFloat(A[i1].Basin,'f',-1,64)})
        for i2=range A[i1].Mat {
            s=append(s,A[i1].Mat[i2].VtoS())
        }
    }
    switch setting {
        case 0: filename="A_physio.csv"
        case 1: filename="A_patho.csv"
        case 2: filename="A_versus.csv"
    }
    if Exist(filename) {
        fmt.Println("\nWARNING: "+filename+" will be overwritten")
    }
    file,_=os.Create(filename)
    writer=csv.NewWriter(file)
    writer.Comma=','
    writer.UseCRLF=false
    writer.WriteAll(s)
    file.Close()
    fmt.Println("\nINFO: saved as "+filename)
}
//#### Sort ##################################################################//
func (a *Attractor) Sort() {
    var i int
    var jmin Vector
    jmin=ItoV(Range(0,(*a).Mat.Size(2)))
    for i=range (*a).Mat {
        jmin=jmin.Pos((*a).Mat[i].Pos(jmin.VtoI()).MinPos())
        if len(jmin)==1 {
            (*a).Mat.CircShift(int(jmin[0]))
            break
        }
    }
}
//#### Sort ##################################################################//
func (A *Aset) Sort() {
    var repass bool
    var i1,i2 int
    for {
        repass=false
        for i1=0;i1<len(*A)-1;i1++ {
            if (*A)[i1].Mat.Size(2)>(*A)[i1+1].Mat.Size(2) {
                (*A).Swap(i1,i1+1)
                repass=true
            } else if (*A)[i1].Mat.Size(2)==(*A)[i1+1].Mat.Size(2) {
                for i2=range (*A)[i1].Mat {
                    if (*A)[i1].Mat[i2][0]>(*A)[i1+1].Mat[i2][0] {
                        (*A).Swap(i1,i1+1)
                        repass=true
                        break
                    } else if (*A)[i1].Mat[i2][0]<(*A)[i1+1].Mat[i2][0] {
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
func (A *Aset) Swap(i1,i2 int) {
    var a Attractor
    if len(*A)>0 {
        a=(*A)[i1].Copy()
        (*A)[i1]=(*A)[i2].Copy()
        (*A)[i2]=a
    }
}
//#### Union #################################################################//
func Union(A1,A2 Aset) Aset {
    var i int
    var A Aset
    for i=range A1 {
        A=append(A,A1[i].Copy())
    }
    for i=range A2 {
        A=append(A,A2[i].Copy())
    }
    return A
}
//#### Versus ################################################################//
func (Aversus *Aset) Versus(Apatho Aset) {
    var i int
    (*Aversus)=Aset{}
    for i=range Apatho {
        if Apatho[i].IsPatho() {
            (*Aversus)=append((*Aversus),Apatho[i].Copy())
        }
    }
}
