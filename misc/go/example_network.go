
// clear && golang-go run example_network.go

package main

import (
    "encoding/csv"
    "fmt"
    "math"
    "math/rand"
    "os"
    "sort"
    "strconv"
    "strings"
    "time"
)

func main() {
    max_targ:=int(1e2)
    max_moda:=int(1e2)
    size_D:=int(1e4)
    V:=[]string{
        "CycD",
        "Rb",
        "E2F",
        "CycE",
        "CycA",
        "p27",
        "Cdc20",
        "Cdh1",
        "UbcH10",
        "CycB",
    }
    what_to_do(f_patho,size_D,max_targ,max_moda,V)
}

func f_physio(x [][]bool,k int) [][]bool {
    return [][]bool{
        {x[0][k]},// CycD
        {(!x[0][k] && !x[3][k] && !x[4][k] && !x[9][k]) || (x[5][k] && !x[0][k] && !x[9][k])},// Rb
        {(!x[1][k] && !x[4][k] && !x[9][k]) || (x[5][k] && !x[1][k] && !x[9][k])},// E2F
        {x[2][k] && !x[1][k]},// CycE
        {(x[2][k] && !x[1][k] && !x[6][k] && !(x[7][k] && x[8][k])) || (x[4][k] && !x[1][k] && !x[6][k] && !(x[7][k] && x[8][k]))},// CycA
        {(!x[0][k] && !x[3][k] && !x[4][k] && !x[9][k]) || (x[5][k] && !(x[3][k] && x[4][k]) && !x[9][k] && !x[0][k])},// p27
        {x[9][k]},// Cdc20
        {(!x[4][k] && !x[9][k]) || x[6][k] || (x[5][k] && !x[9][k])},// Cdh1
        {!x[7][k] || (x[7][k] && x[8][k] && (x[6][k] || x[4][k] || x[9][k]))},// UbcH10
        {!x[6][k] && !x[7][k]},// CycB
    }
}

func f_patho(x [][]bool,k int) [][]bool {
    return [][]bool{
        {x[0][k]},// CycD
        {false},// Rb
        {(!x[1][k] && !x[4][k] && !x[9][k]) || (x[5][k] && !x[1][k] && !x[9][k])},// E2F
        {x[2][k] && !x[1][k]},// CycE
        {(x[2][k] && !x[1][k] && !x[6][k] && !(x[7][k] && x[8][k])) || (x[4][k] && !x[1][k] && !x[6][k] && !(x[7][k] && x[8][k]))},// CycA
        {(!x[0][k] && !x[3][k] && !x[4][k] && !x[9][k]) || (x[5][k] && !(x[3][k] && x[4][k]) && !x[9][k] && !x[0][k])},// p27
        {x[9][k]},// Cdc20
        {(!x[4][k] && !x[9][k]) || x[6][k] || (x[5][k] && !x[9][k])},// Cdh1
        {!x[7][k] || (x[7][k] && x[8][k] && (x[6][k] || x[4][k] || x[9][k]))},// UbcH10
        {!x[6][k] && !x[7][k]},// CycB
    }
}

////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////    lib    ///////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func all(x []bool) bool {
    for i1:=range x {if x[i1]==false {return false}}
    return true
}

func any(x []bool) bool {
    for i1:=range x {if x[i1]==true {return true}}
    return false
}

func compare_attractor(a1,a2 [][]bool) bool {
    if len(a1[0])!=len(a2[0]) {return true} else {
        start_found:=false
        start1:=-1
        start2:=-1
        for j1:=range a1[0] {
            for j2:=range a2[0] {
                z:=[]bool{}
                for i1:=range a1 {z=append(z,a1[i1][j1]==a2[i1][j2])}
                if all(z) {
                    start_found=true
                    start1=j1
                    start2=j2
                    break
                }
            }
            if start_found {break}
        }
        if !start_found {return true} else {
            for j1:=1;j1<=len(a1[0])-1;j1++ {
                z:=[]bool{}
                for i1:=range a1 {z=append(z,a1[i1][(start1+j1)%len(a1[i1])]==a2[i1][(start2+j1)%len(a2[i1])])}
                if !all(z) {return true}
            }
            return false
        }
    }
}

func compare_attractor_set(A1,A2 [][][]bool) bool {
    if len(A1)!=len(A2) {return true} else {
        in_2:=[]bool{}
        for i1:=range A1 {
            z:=false
            for i2:=range A2 {
                if !compare_attractor(A1[i1],A2[i2]) {
                    z=true
                    break
                }
            }
            in_2=append(in_2,z)
        }
        return !all(in_2)
    }
}

func compute_attractor(f func(x [][]bool,k int) [][]bool,c_targ []int,c_moda []bool,D [][]bool) [][][]bool {
    A:=[][][]bool{}
    for j1:=range D[0] {
        x:=[][]bool{}
        for i1:=range D {x=append(x,[]bool{D[i1][j1]})}
        k:=0
        a_found:=false
        a:=[][]bool{}
        for {
            z:=f(x,k)
            for i1:=range x {x[i1]=append(x[i1],z[i1]...)}
            for i1:=range c_targ {x[c_targ[i1]][k+1]=c_moda[i1]}
            for j2:=k;j2>=0;j2-- {
                z:=[]bool{}
                for i1:=range x {z=append(z,x[i1][j2]==x[i1][k+1])}
                if all(z) {
                    a_found=true
                    for i1:=range x {a=append(a,x[i1][j2:k+1])}
                    break
                }
            }
            if a_found {
                in_A:=false
                for i1:=range A {
                    if !compare_attractor(a,A[i1]) {
                        in_A=true
                        break
                    }
                }
                if !in_A {A=append(A,a)}
                break
            }
            k+=1
        }
    }
    return A
}

func compute_pathological_attractor(A_physio,A_patho [][][]bool) [][][]bool {
    a_patho_set:=[][][]bool{}
    for i1:=range A_patho {
        in_physio:=false
        for i2:=range A_physio {
            if !compare_attractor(A_patho[i1],A_physio[i2]) {
                in_physio=true
                break
            }
        }
        if !in_physio {a_patho_set=append(a_patho_set,A_patho[i1])}
    }
    return a_patho_set
}

func compute_therapeutic_bullet(f func(x [][]bool,k int) [][]bool,D [][]bool,r_min int,r_max int,max_targ int,max_moda int,A_physio [][][]bool) ([][]int,[][]bool,[]string) {
    targ_set:=[][]int{}
    moda_set:=[][]bool{}
    metal_set:=[]string{}
    for i1:=r_min;i1<=int(math.Min(float64(r_max),float64(len(D))));i1++ {
        C_targ:=generate_combination(i1,len(D),max_targ)
        C_moda:=generate_arrangement(i1,max_moda)
        for i2:=range C_targ {
            for i3:=range C_moda {
                A_patho:=compute_attractor(f,C_targ[i2],C_moda[i3],D)
                if len(compute_pathological_attractor(A_physio,A_patho))==0 {
                    metal:=""
                    if compare_attractor_set(A_physio,A_patho) {metal="silver"} else {metal="gold"}
                    targ_set=append(targ_set,C_targ[i2])
                    moda_set=append(moda_set,C_moda[i3])
                    metal_set=append(metal_set,metal)
                }
            }
        }
    }
    return targ_set,moda_set,metal_set
}

func factorial(x float64) float64 {
    if x==float64(0) {return float64(1)} else {return x*factorial(x-float64(1))}
}

func generate_arrangement(k,n_arrang int) [][]bool {
    ////////////////////    /!\ only with repetition /!\    ////////////////////
    arrang_mat:=[][]bool{}
    for i1:=1;i1<=int(math.Min(float64(n_arrang),math.Pow(float64(2),float64(k))));i1++ {
        for {
            arrang:=[]bool{}
            for j1:=1;j1<=k;j1++ {
                z:=false
                if rand.Intn(2)==1 {z=true}
                arrang=append(arrang,z)
            }
            in_arrang_mat:=false
            for i2:=range arrang_mat {
                z:=[]bool{}
                for j1:=range arrang {z=append(z,arrang[j1]==arrang_mat[i2][j1])}
                if all(z) {
                    in_arrang_mat=true
                    break
                }
            }
            if !in_arrang_mat {
                arrang_mat=append(arrang_mat,arrang)
                break
            }
        }
    }
    return arrang_mat
}

func generate_combination(k,n,n_combi int) [][]int {
    //////////////////    /!\ only without repetition /!\    ///////////////////
    combi_mat:=[][]int{}
    for i1:=1;i1<=int(math.Min(float64(n_combi),factorial(float64(n))/(factorial(float64(k))*factorial(float64(n-k)))));i1++ {
        for {
            combi:=[]int{}
            for j1:=1;j1<=k;j1++ {
                for {
                    z:=rand.Intn(n)
                    in_combi:=false
                    for i2:=range combi {
                        if combi[i2]==z {
                            in_combi=true
                            break
                        }
                    }
                    if !in_combi {
                        combi=append(combi,z)
                        break
                    }
                }
            }
            sort.Ints(combi)
            in_combi_mat:=false
            for i2:=range combi_mat {
                z:=[]bool{}
                for j1:=range combi {z=append(z,combi[j1]==combi_mat[i2][j1])}
                if all(z) {
                    in_combi_mat=true
                    break
                }
            }
            if !in_combi_mat {
                combi_mat=append(combi_mat,combi)
                break
            }
        }
    }
    return combi_mat
}

func generate_state_space(n int) [][]bool {
    y:=[][]bool{{false,true}}
    for i1:=1;i1<=n-1;i1++ {
        for i2:=range y {y[i2]=append(y[i2],y[i2]...)}
        y=append(y,[]bool{})
        for j1:=range y[i1-1] {
            z:=false
            if j1>=len(y[i1-1])/2 {z=true}
            y[i1]=append(y[i1],z)
        }
    }
    return y
}

func load_attractor_set(setting int) [][][]bool {
    A:=[][][]bool{}
    set_name:=""
    switch setting {
        case 1: set_name="set_physio.csv"
        case 2: set_name="set_patho.csv"
    }
    csv_file,_:=os.Open(set_name)
    csv_reader:=csv.NewReader(csv_file)
    csv_reader.FieldsPerRecord=-1
    s,_:=csv_reader.ReadAll()
    csv_file.Close()
    s_bis:=[][]bool{}
    for i1:=range s {
        if i1>=1 {
            s_bis=append(s_bis,[]bool{})
            for j1:=range s[i1] {
                z,_:=strconv.ParseBool(s[i1][j1])
                s_bis[i1-1]=append(s_bis[i1-1],z)
            }
        }
    }
    n_attractor,_:=strconv.ParseInt(s[0][0],10,0)
    n_line:=(len(s)-1)/int(n_attractor)
    for i1:=1;i1<=int(n_attractor);i1++ {A=append(A,s_bis[(i1-1)*n_line:i1*n_line])}
    return A
}

func report_attractor_set(A [][][]bool,setting int,V []string) {
    n_point:=0
    n_cycle:=0
    report:=strings.Repeat("-",80)+"\n"
    for i1:=range A {
        if len(A[i1][0])==1 {n_point+=1} else {n_cycle+=1}
        for i2:=range A[i1] {
            report+=V[i2]+": "
            z:=[]string{}
            for i3:=range A[i1][i2] {z=append(z,strconv.FormatBool(A[i1][i2][i3]))}
            report+=strings.Join(z," ")+"\n"
        }
        report+=strings.Repeat("-",80)+"\n"
    }
    report+="found attractors: "+strconv.FormatInt(int64(len(A)),10)+" ("+strconv.FormatInt(int64(n_point),10)+" points, "+strconv.FormatInt(int64(n_cycle),10)+" cycles)"
    fmt.Println("\n"+report+"\n")
    save:="y"
    fmt.Printf("save [Y/n] ")
    fmt.Scanf("%s",&save)
    if strings.ToLower(save)=="y" || strings.ToLower(save)=="yes" {
        set_name:=""
        report_name:=""
        switch setting {
            case 1:
                set_name="set_physio.csv"
                report_name="report_physio.txt"
            case 2:
                set_name="set_patho.csv"
                report_name="report_patho.txt"
            case 3:
                set_name="set_versus.csv"
                report_name="report_versus.txt"
        }
        file,_:=os.Create(report_name)
        file.WriteString(report)
        file.Close()
        s:=[][]string{{strconv.FormatInt(int64(len(A)),10)}}
        for i1:=range A {
            for i2:=range A[i1] {
                s=append(s,[]string{})
                for i3:=range A[i1][i2] {s[len(s)-1]=append(s[len(s)-1],strconv.FormatBool(A[i1][i2][i3]))}
            }
        }
        csv_file,_:=os.Create(set_name)
        csv_writer:=csv.NewWriter(csv_file)
        csv_writer.WriteAll(s)
        csv_file.Close()
        fmt.Printf("\nset saved as: %s\nreport saved as: %s\n\n",set_name,report_name)
    }
}

func report_therapeutic_bullet_set(targ_set [][]int,moda_set [][]bool,metal_set []string,V []string) {
    n_gold:=0
    n_silv:=0
    report:=strings.Repeat("-",80)+"\n"
    for i1:=range targ_set {
        if metal_set[i1]=="gold" {n_gold+=1} else {n_silv+=1}
        for i2:=range targ_set[i1] {
            moda:=""
            if moda_set[i1][i2] {moda="+"} else {moda="-"}
            report+=moda+V[targ_set[i1][i2]]+" "
        }
        report+="("+metal_set[i1]+" bullet)\n"+strings.Repeat("-",80)+"\n"
    }
    report+="found therapeutic bullets: "+strconv.FormatInt(int64(len(targ_set)),10)+" ("+strconv.FormatInt(int64(n_gold),10)+" gold bullets, "+strconv.FormatInt(int64(n_silv),10)+" silver bullets)"
    fmt.Println("\n"+report+"\n")
    save:="y"
    fmt.Printf("save [Y/n] ")
    fmt.Scanf("%s",&save)
    if strings.ToLower(save)=="y" || strings.ToLower(save)=="yes" {
        report_name:="report_therapeutic_bullet.txt"
        file,_:=os.Create(report_name)
        file.WriteString(report)
        file.Close()
        fmt.Printf("\nreport saved as: %s\n\n",report_name)
    }
}

func transpose(x [][]bool) [][]bool {
    y:=[][]bool{}
    for j1:=range x[0] {
        y=append(y,[]bool{})
        for i1:=range x {y[j1]=append(y[j1],x[i1][j1])}
    }
    return y
}

func what_to_do(f func(x [][]bool,k int) [][]bool,size_D int,max_targ int,max_moda int,V []string) {
    rand.Seed(int64(time.Now().Nanosecond()))
    to_do:=5
    fmt.Printf("\n[1] compute attractors\n[2] compute pathological attractors\n[3] compute therapeutic bullets\n[4] help\n[5] license\n\nwhat to do: ")
    fmt.Scanf("%d",&to_do)
    D:=[][]bool{}
    if to_do==1 || to_do==3 {
        comprehensive_D:="n"
        fmt.Printf("\nsize(S)=%e, comprehensive D [y/N] ",math.Pow(float64(2),float64(len(V))))
        fmt.Scanf("%s",&comprehensive_D)
        if strings.ToLower(comprehensive_D)=="y" || strings.ToLower(comprehensive_D)=="yes" {D=generate_state_space(len(V))} else {D=transpose(generate_arrangement(len(V),size_D))}
    }
    switch to_do {
        case 1:
            A:=compute_attractor(f,[]int{},[]bool{},D)
            setting:=1
            fmt.Printf("\n[1] physiological\n[2] pathological\n\nsetting: ")
            fmt.Scanf("%d",&setting)
            report_attractor_set(A,setting,V)
        case 2:
            A_physio:=load_attractor_set(1)
            A_patho:=load_attractor_set(2)
            a_patho_set:=compute_pathological_attractor(A_physio,A_patho)
            report_attractor_set(a_patho_set,3,V)
        case 3:
            A_physio:=load_attractor_set(1)
            r_min:=1
            fmt.Printf("\nr_min=")
            fmt.Scanf("%d",&r_min)
            r_max:=1
            fmt.Printf("\nr_max=")
            fmt.Scanf("%d",&r_max)
            targ_set,moda_set,metal_set:=compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,A_physio)
            report_therapeutic_bullet_set(targ_set,moda_set,metal_set,V)
        case 4:
            fmt.Println("\n1) do step 1 with f_physio\n2) do step 1 with f_patho\n3) eventually do step 2\n4) do step 3 with f_patho\n\ndo not forget to recompile the sources following any modification\n")
        case 5:
            fmt.Println("\nCopyright (c) 2013-2014, Arnaud Poret\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without modification,\nare permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice, this\nlist of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation and/or\nother materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\" AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR\nANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON\nANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
    }
}

