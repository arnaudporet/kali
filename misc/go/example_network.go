
// clear && golang-go run example_network.go

package main

import (
    "fmt"
    "math"
    "time"
    "math/rand"
    "sort"
    "strings"
    "strconv"
    "os"
    "encoding/csv"
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
    what_to_do()
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

////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////    lib    ///////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func all(x []bool) bool {
    for i,_:=range x {if x[i]==false {return false}}
    return true
}

func any(x []bool) bool {
    for i,_:=range x {if x[i]==true {return true}}
    return false
}

func compare_attractor(a1,a2 [][]bool) bool {
    // var start1,start2 int
    if len(a1[0])!=len(a2[0]) {return true} else {
        start_found:=false
        for j1,_:=range a1[0] {
            for j2,_:=range a2[0] {
                z:=[]bool{}
                for i,_:=range a1 {z=append(z,a1[i][j1]==a2[i][j2])}
                if all(z) {
                    start_found=true
                    start1:=j1
                    start2:=j2
                    break
                }
            }
            if start_found {break}
        }
        if !start_found {return true} else {
            for j:=1;j<=len(a1[0])-1;j++ {
                z:=[]bool{}
                for i,_:=range a1 {z=append(z,a1[i][(start1+j)%len(a1[0])]==a2[i][(start2+j)%len(a2[0])])}
                if !all(z) {return true}
            }
            return false
        }
    }
}

func compare_attractor_set(A1,A2 [][][]bool) bool {
    if len(A1)!=len(A2) {return true} else {
        in_2:=[]bool{}
        for i1,_:=range A1 {
            z:=false
            for i2,_:=range A2 {
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
    for i1,_:=range D[0] {
        x:=[][]bool{}
        for i,_:=range D {x=append(x,[]bool{D[i][i1]})}
        k:=0
        for {
            z:=f(x,k)
            for i,_:=range x {x[i]=append(x[i],z[i])}
            for i,_:=range c_targ {x[c_targ[i]][k+1]=c_moda[i]}
            a_found:=false
            for i2:=k;i2>=0;i2-- {
                z:=[]bool{}
                for i,_:=range x {z=append(z,x[i][i2]==x[i][k+1])}
                if all(z) {
                    a_found=true
                    a:=[][]bool{}
                    for i,_:=range x {a=append(a,x[i][i2:k+1])}
                    break
                }
            }
            if a_found {
                in_A:=false
                for i2,_:=range A {
                    if !compare_attractor(a,A[i2]) {
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
    for i1,_:=range A_patho {
        in_physio:=false
        for i2,_:=range A_physio {
            if !compare_attractor(A_patho[i1],A_physio[i2]) {
                in_physio=true
                break
            }
        }
        if !in_physio {a_patho_set=append(a_patho_set,A_patho[i1])}
    }
    return a_patho_set
}

func compute_therapeutic_bullet(f func(x [][]bool,k int) [][]bool,D [][]bool,r_min int,r_max int,max_targ int,max_moda int,n_node int,A_physio [][][]bool) ([][]int,[][]bool,[]string) {
    targ_set:=[][]int
    moda_set:=[][]bool
    metal_set:=[]string
    for i1:=r_min;i1<=int(math.Min(float64(r_max),float64(n_node)));i1++ {
        C_targ:=generate_combination(i1,n_node,max_targ)
        C_moda:=generate_arrangement(i1,max_moda)
        for i2:=range C_targ {
            for i3:=range C_moda {
                A_patho:=compute_attractor(f,C_targ[i2],C_moda[i3],D)
                if len(compute_pathological_attractor(A_physio,A_patho))==0 {
                    if compare_attractor_set(A_physio,A_patho) {
                        metal:="silver"
                    } else {
                        metal:="golden"
                    }
                    targ_set=append(targ_set,C_targ[i2])
                    moda_set=append(moda_set,C_moda[i3])
                    metal_set=append(metal_set,metal)
                }
            }
        }
    }
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
            for i2:=1;i2<=k;i2++ {
                z:=false
                if rand.Intn(2)==1 {z=true}
                arrang=append(arrang,z)
            }
            in_arrang_mat:=false
            for i2,_:=range arrang_mat {
                z:=[]bool{}
                for i,_:=range arrang {z=append(z,arrang[i]==arrang_mat[i2][i])}
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
            for i2:=1;i2<=k;i2++ {
                for {
                    z:=rand.Intn(n)
                    in_combi:=false
                    for i3,_:=range combi {
                        if combi[i3]==z {
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
            for i2,_:=range combi_mat {
                z:=[]bool{}
                for i,_:=range combi {z=append(z,combi[i]==combi_mat[i2][i])}
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
        for i2,_:=range y {y[i2]=append(y[i2],y[i2]...)}
        y=append(y,[]bool{})
        for j,_:=range y[0] {
            z:=false
            if j>=len(y[0])/2 {z=true}
            y[i1]=append(y[i1],z)
        }
    }
    return y
}

func load_attractor_set(setting int) [][][]bool {
    A:=[][][]bool{}
    switch setting {
        case 1: set_name:="set_physio.csv"
        case 2: set_name:="set_patho.csv"
    }
    csv_file,_:=os.Open(set_name)
    csv_reader:=csv.NewReader(csv_file)
    csv_reader.FieldsPerRecord=-1
    s,_:=csv_reader.ReadAll()
    csv_file.Close()
    s_bis:=[][]bool{}
    for i,_:=range s {
        if i>=1 {
            s_bis=append(s_bis,[]bool{})
            for j,_:=range s[i] {
                z,_:=strconv.ParseBool(s[i][j])
                s_bis[len(s_bis)-1]=append(s_bis[len(s_bis)-1],z)
            }
        }
    }
    n_attractor,_:=strconv.ParseInt(s[0][0],10,0)
    n_line:=(len(s)-1)/int(n_attractor)
    for i:=1;i<=int(n_attractor);i++ {A=append(A,s_bis[(i-1)*n_line:i*n_line])}
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
    fmt.Println(report)
    save:="y"
    fmt.Printf("save [Y/n]" )
    fmt.Scanf("%s",&save)
    if strings.ToLower(save)=="y" || strings.ToLower(save)=="yes" {
        switch setting {
            case 1: set_name,report_name:="set_physio.csv","report_physio.txt"
            case 2: set_name,report_name:="set_patho.csv","report_patho.txt"
            case 3: set_name,report_name:="set_versus.csv","report_versus.txt"
        }
        file,_:=os.Create(report_name)
        file.WriteString(report)
        file.Close()
        s:=[][]string{{strconv.FormatInt(int64(len(A)),10)}}
        for i1,_:=range A {
            for i2,_:=range A[i1] {
                s=append(s,[]string{})
                for i3,_:=range A[i1][i2] {s[len(s)-1]=append(s[len(s)-1],strconv.FormatBool(A[i1][i2][i3]))}
            }
        }
        csv_file,_:=os.Create(set_name)
        csv_writer:=csv.NewWriter(csv_file)
        csv_writer.WriteAll(s)
        csv_file.Close()
        fmt.Printf("set saved as: %s\nreport saved as: %s\n",set_name,report_name)
    }
}

func report_therapeutic_bullet_set(targ_set [][]int,moda_set [][]bool,metal_set []string,V []string) {
    n_gold:=0
    n_silv:=0
    report:=strings.Repeat("-",80)+"\n"
    for i1,_:=range targ_set {
        if metal_set[i1]=="golden" {n_gold+=1} else {n_silv+=1}
        for i2,_:=range targ_set[i1]) {
            moda:="-"
            if moda_set[i1][i2] {moda="+"}
            report+=moda+V[targ_set[i1][i2]]+" "
        }
        report+="("+metal[i1]+" bullet)\n"+strings.Repeat("-",80)+"\n"
    }
    report+="found therapeutic bullets: "+strconv.FormatInt(int64(len(targ_set)),10)+" ("+strconv.FormatInt(int64(n_gold),10)+" golden bullets, "+strconv.FormatInt(int64(n_silv),10)+" silver bullets)"
    fmt.Println(report)
    save:="y"
    fmt.Printf("save [Y/n]" )
    fmt.Scanf("%s",&save)
    if strings.ToLower(save)=="y" || strings.ToLower(save)=="yes" {
        file,_:=os.Create("report_therapeutic_bullet.txt")
        file.WriteString(report)
        file.Close()
        fmt.Println("report saved as: report_therapeutic_bullet.txt")
    }
}

func transpose(x [][]int) [][]int {
    y:=[][]int{}
    for i,_:=range x {
        y=append(y,[]int{})
        for j,_:=range x[i] {y[len(y)-1]=append(y[len(y)-1],x[j][i])}
    }
    return y
}

func what_to_do(f func(x [][]bool,k int) [][]bool,size_D int,n_node int,max_targ int,max_moda int,V []string) {
    rand.Seed(int64(time.Now().Nanosecond()))
    to_do:=5
    fmt.Printf("[1] compute attractors\n[2] compute pathological attractors\n[3] compute therapeutic bullets\n[4] help\n[5] license\nwhat to do: ")
    fmt.Scanf("%d",&to_do)
    if to_do==1 || to_do==3 {
        comprehensive_D:="n"
        fmt.Printf("size(S)=%e\ncomprehensive D [y/N] ",math.Pow(float64(2),float64(n_node)))
        fmt.Scanf("%s",&comprehensive_D)
        if strings.ToLower(comprehensive_D)=="y" || strings.ToLower(comprehensive_D)=="yes" {D:=generate_state_space(n_node)} else {D:=transpose(generate_arrangement(n_node,size_D))}
    }
    switch to_do {
        case: 1 {
            A:=compute_attractor(f,[]int{},[]bool{},D)
            setting:=1
            fmt.Printf("[1] physiological\n[2] pathological\nsetting: ")
            fmt.Scanf("%d",&setting)
            report_attractor_set(A,setting,V)
        }
        case: 2 {
            A_physio:=load_attractor_set(1)
            A_patho:=load_attractor_set(2)
            a_patho_set:=compute_pathological_attractor(A_physio,A_patho)
            report_attractor_set(a_patho_set,3,V)
        }
        case: 3 {
            A_physio:=load_attractor_set(1)
            r_min:=1
            fmt.Printf("r_min=")
            fmt.Scanf("%d",&r_min)
            r_max:=1
            fmt.Printf("r_max=")
            fmt.Scanf("%d",&r_max)
            targ_set,moda_set,metal_set:=compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,n_node,A_physio)
            report_therapeutic_bullet_set(targ_set,moda_set,metal_set,V)
        }
        case: 4 {
            fmt.Println("1) do step 1 with f_physio\n2) do step 1 with f_patho\n3) eventually do step 2\n4) do step 3 with f_patho\ndo not forget to recompile the sources following any modification")
        }
        case: 5 {
            fmt.Println("Copyright (c) 2013-2014, Arnaud Poret\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without modification,\nare permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice, this\nlist of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation and/or\nother materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\" AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR\nANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON\nANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
        }
    }
}
