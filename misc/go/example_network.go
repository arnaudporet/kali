
// clear && golang-go run example_network.go

package main

import "fmt"

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
        if !start_found {return true}
        else {
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

func compute_attractor(f func(x [][]bool,k int) [][]bool,c_targ []int,c_moda []bool,D [][]bool) ([][][]bool,[]float64) {
    A:=[][][]bool{}
    count:=[]int{}
    popularity:=[]float64{}
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
                    for i,_:=range x {
                        a=append(a,x[i][i2:k+1])
                    }
                    break
                }
            }
            if a_found {
                in_A:=false
                for i2,_:=range A {
                    if !compare_attractor(a,A[i2]) {
                        in_A=true
                        count[i2]+=1
                        break
                    }
                }
                if !in_A {
                    A=append(A,a)
                    count=append(count,1)
                }
                break
            }
            k+=1
        }
    }
    for i,_:=range count {popularity=append(popularity,(float64(count[i])/float64(len(D[0])))*100.0)}
    return A,popularity
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
}

!##########################################################################!
!####################    compute_therapeutic_bullet    ####################!
!##########################################################################!
function compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio) result(therapeutic_bullet_set)
    implicit none
    integer::r_min,r_max,i1,i2,i3,max_targ,max_moda,n_node
    type(attractor),dimension(:)::A_physio
    real,dimension(:,:)::D
    type(bullet),dimension(:),allocatable::therapeutic_bullet_set
    type(attractor),dimension(:),allocatable::A_patho
    integer,dimension(:,:),allocatable::C_targ
    real,dimension(:,:),allocatable::C_moda
    real,dimension(:)::value
    character(6)::metal
    interface
        function f(x,k) result(y)
            implicit none
            real,dimension(:,:)::x
            integer::k
            real,dimension(size(x,1),1)::y
        end function f
    end interface
    allocate(therapeutic_bullet_set(0))
    do i1=r_min,min(r_max,n_node)
        C_targ=generate_combination(range_int(1,n_node),i1,max_targ)
        C_moda=generate_arrangement(value,i1,max_moda)
        do i2=1,size(C_targ,1)
            do i3=1,size(C_moda,1)
                A_patho=compute_attractor(f,C_targ(i2,:),C_moda(i3,:),D)
                if (size(compute_pathological_attractor(A_physio,A_patho))==0) then
                    if (compare_attractor_set(A_physio,A_patho)) then
                        metal="silver"
                    else
                        metal="golden"
                    end if
                    therapeutic_bullet_set=add_bullet(therapeutic_bullet_set,C_targ(i2,:),C_moda(i3,:),metal)
                end if
                deallocate(A_patho)
            end do
        end do
        deallocate(C_targ,C_moda)
    end do
end function compute_therapeutic_bullet

func concatenate(x,y [][]bool) [][]bool {
    for i,_:=range x {
        x[i]=append(x[i],y[i]...)
    }
    return x
}
!##########################################################################!
!##############################    facto    ###############################!
!##########################################################################!
function facto(x) result(y)
    implicit none
    integer::x,i
    real(kind=8)::y
    if (x>170) then
        write (unit=*,fmt="(a)") "facto(x): x>170 unsupported"
        stop
    end if
    if (x==0) then
        y=real(1,8)
    else
        y=real(x,8)
        do i=1,x-1
            y=y*(real(x,8)-real(i,8))
        end do
    end if
end function facto
!##########################################################################!
!#######################    generate_arrangement    #######################!
!##########################################################################!
function generate_arrangement(deck,k,n_arrang) result(arrang_mat)
    !#################    /!\ only with repetition /!\    #################!
    implicit none
    real,dimension(:)::deck
    integer::k,i1,i2,n_arrang
    real(kind=8)::max_arrang
    real,dimension(:,:),allocatable::arrang_mat
    real,dimension(k)::arrang
    max_arrang=min(real(size(deck),8)**real(k,8),real(huge(1),8))
    allocate(arrang_mat(min(n_arrang,int(max_arrang)),k))
    do i1=1,size(arrang_mat,1)
        1 continue
        do i2=1,k
            arrang(i2)=deck(rand_int(1,size(deck)))
        end do
        do i2=1,i1-1
            if (all(arrang_mat(i2,:)==arrang)) then
                go to 1
            end if
        end do
        arrang_mat(i1,:)=arrang
    end do
end function generate_arrangement
!##########################################################################!
!#######################    generate_combination    #######################!
!##########################################################################!
function generate_combination(deck,k,n_combi) result(combi_mat)
    !###############    /!\ only without repetition /!\    ################!
    implicit none
    integer,dimension(:)::deck
    integer::k,i1,i2,z,n_combi
    real(kind=8)::max_combi
    integer,dimension(:,:),allocatable::combi_mat
    integer,dimension(k)::combi
    max_combi=min(facto(size(deck))/(facto(k)*facto(size(deck)-k)),real(huge(1),8))
    allocate(combi_mat(min(n_combi,int(max_combi)),k))
    do i1=1,size(combi_mat,1)
        1 continue
        do i2=1,k
            2 continue
            z=deck(rand_int(1,size(deck)))
            if (any(z==combi(1:i2-1))) then
                go to 2
            else
                combi(i2)=z
            end if
        end do
        combi=sort(combi)
        do i2=1,i1-1
            if (all(combi_mat(i2,:)==combi)) then
                go to 1
            end if
        end do
        combi_mat(i1,:)=combi
    end do
end function generate_combination
!##########################################################################!
!#######################    generate_state_space    #######################!
!##########################################################################!
function generate_state_space(n) result(y)
    implicit none
    integer::n,i1,i2
    real,dimension(:,:),allocatable::y
    real,dimension(:,:),allocatable::z
    if (n>30) then
        write (unit=*,fmt="(a)") "generate_state_space(n): n>30 unsupported"
        stop
    end if
    allocate(y(2,1))
    y(1,1)=0.0
    y(2,1)=1.0
    do i1=1,n-1
        allocate(z(size(y,1)*2,size(y,2)+1))
        do i2=1,size(y,1)
            z(2*i2-1,:size(y,2))=y(i2,:)
            z(2*i2-1,size(y,2)+1)=0.0
            z(2*i2,:size(y,2))=y(i2,:)
            z(2*i2,size(y,2)+1)=1.0
        end do
        deallocate(y)
        y=z
        deallocate(z)
    end do
end function generate_state_space
!##########################################################################!
!############################    Heaviside    #############################!
!##########################################################################!
function Heaviside(x) result(y)
    implicit none
    real::x,y
    if (x<=0.0) then
        y=0.0
    else
        y=1.0
    end if
end function Heaviside
!##########################################################################!
!#########################    init_random_seed    #########################!
!##########################################################################!
subroutine init_random_seed()
    implicit none
    integer,allocatable::seed(:)
    integer::n,un,istat
    call random_seed(size=n)
    allocate(seed(n))
    open(newunit=un,file="/dev/urandom",access="stream",form="unformatted",action="read",status="old",iostat=istat)
    read(un) seed
    close(un)
    call random_seed(put=seed)
    deallocate(seed)
end subroutine init_random_seed
!##########################################################################!
!#############################    int2char    #############################!TODO merge with real2char?
!##########################################################################!
function int2char(x) result(y)
    implicit none
    integer::x
    character(:),allocatable::y
    character(9)::z
    if (x<0 .or. x>999999999) then
        write (unit=*,fmt="(a)") "int2char(x): x<0 or x>999 999 999 unsupported"!FIXME
        stop
    end if
    write (unit=z,fmt="(i9)") x
    y=trim(adjustl(z))
end function int2char
!##########################################################################!
!########################    load_attractor_set    ########################!
!##########################################################################!
function load_attractor_set(setting) result(A_set)
    implicit none
    integer::setting
    type(attractor),dimension(:),allocatable::A_set
    character(:),allocatable::set_name
    integer::i1,i2,z,n,m
    select case (setting)
        case (1)
            set_name="set_physio"
        case (2)
            set_name="set_patho"
    end select
    open (unit=1,file=set_name,status="old")
    read (unit=1,fmt=*) z
    allocate(A_set(z))
    do i1=1,size(A_set)
        read (unit=1,fmt=*) n
        read (unit=1,fmt=*) m
        read (unit=1,fmt=*) A_set(i1)%popularity
        allocate(A_set(i1)%a(n,m))
    end do
    do i1=1,size(A_set)
        do i2=1,size(A_set(i1)%a,1)
            read (unit=1,fmt="("//int2char(size(A_set(i1)%a,2))//"f3.1)") A_set(i1)%a(i2,:)
        end do
    end do
    close (unit=1)
    deallocate(set_name)
end function load_attractor_set
!##########################################################################!
!#############################    rand_int    #############################!
!##########################################################################!
function rand_int(a,b) result(y)
    implicit none
    integer::a,b,y
    real::x
    call random_number(x)
    y=nint(real(a)+x*(real(b)-real(a)))
end function rand_int
!##########################################################################!
!############################    range_int    #############################!
!##########################################################################!
function range_int(a,b) result(y)
    implicit none
    integer::a,b,i
    integer,dimension(b-a+1)::y
    do i=1,size(y)
        y(i)=a+i-1
    end do
end function range_int
!##########################################################################!
!############################    real2char    #############################!
!##########################################################################!
function real2char(x) result(y)
    !####################    /!\ only one digit /!\    ####################!
    implicit none
    real::x
    character(:),allocatable::y
    character(5)::z
    if (x<0.0 .or. x>999.9) then
        write (unit=*,fmt="(a)") "real2char(x): x<0.0 or x>999.9 unsupported"!FIXME
        stop
    end if
    write (unit=z,fmt="(f5.1)") x
    y=trim(adjustl(z))
end function real2char
!##########################################################################!
!#######################    report_attractor_set    #######################!
!##########################################################################!
subroutine report_attractor_set(A_set,setting,V)
    implicit none
    type(attractor),dimension(:)::A_set
    integer::setting,n_point,n_cycle,i1,i2,i3,save_
    character(16),dimension(:)::V
    character(:),allocatable::report
    character(:),allocatable::s
    character(:),allocatable::set_name
    character(:),allocatable::report_name
    n_point=0
    n_cycle=0
    report=repeat("-",80)//new_line("a")
    do i1=1,size(A_set)
        if (size(A_set(i1)%a,2)==1) then
            n_point=n_point+1
        else
            n_cycle=n_cycle+1
        end if
        report=report//"popularity: "//real2char(A_set(i1)%popularity)//"%"//new_line("a")//new_line("a")
        do i2=1,size(A_set(i1)%a,1)
            report=report//V(i2)//": "
            do i3=1,size(A_set(i1)%a,2)
                report=report//real2char(A_set(i1)%a(i2,i3))//" "
            end do
            report=report//new_line("a")
        end do
        report=report//repeat("-",80)//new_line("a")
    end do
    report=report//"found attractors: "//int2char(size(A_set))//" ("//int2char(n_point)//" points, "//int2char(n_cycle)//&
    " cycles)"
    write (unit=*,fmt="(a)") new_line("a")//report//new_line("a")
    write (unit=*,fmt="(a)") "save? [1/0]"//new_line("a")
    read (unit=*,fmt=*) save_
    if (save_==1) then
        select case (setting)
            case (1)
                set_name="set_physio"
                report_name="report_physio"
            case (2)
                set_name="set_patho"
                report_name="report_patho"
            case (3)
                set_name="set_versus"
                report_name="report_versus"
        end select
        s=int2char(size(A_set))//new_line("a")
        do i1=1,size(A_set)
            s=s//int2char(size(A_set(i1)%a,1))//new_line("a")//int2char(size(A_set(i1)%a,2))//new_line("a")//&
            real2char(A_set(i1)%popularity)//new_line("a")
        end do
        do i1=1,size(A_set)
            do i2=1,size(A_set(i1)%a,1)
                do i3=1,size(A_set(i1)%a,2)
                    s=s//real2char(A_set(i1)%a(i2,i3))
                end do
                if (i1/=size(A_set) .or. i2/=size(A_set(i1)%a,1)) then
                    s=s//new_line("a")
                end if
            end do
        end do
        open (unit=1,file=set_name,status="replace")
        write (unit=1,fmt="(a)") s
        close (unit=1)
        open (unit=1,file=report_name,status="replace")
        write (unit=1,fmt="(a)") report
        close (unit=1)
        write (unit=*,fmt="(a)") new_line("a")//"set saved as: "//set_name//new_line("a")//"report saved as: "//report_name//&
        new_line("a")
        deallocate(s,set_name,report_name)
    end if
    deallocate(report)
end subroutine report_attractor_set
!##########################################################################!
!##################    report_therapeutic_bullet_set    ###################!
!##########################################################################!
subroutine report_therapeutic_bullet_set(therapeutic_bullet_set,V)
    implicit none
    type(bullet),dimension(:)::therapeutic_bullet_set
    character(16),dimension(:)::V
    integer::n_gold,n_silv,i1,i2,save_
    character(:),allocatable::report
    n_gold=0
    n_silv=0
    report=repeat("-",80)//new_line("a")
    do i1=1,size(therapeutic_bullet_set)
        if (therapeutic_bullet_set(i1)%metal=="golden") then
            n_gold=n_gold+1
        else
            n_silv=n_silv+1
        end if
        do i2=1,size(therapeutic_bullet_set(i1)%targ)
            report=report//trim(V(therapeutic_bullet_set(i1)%targ(i2)))//"["//real2char(therapeutic_bullet_set(i1)%moda(i2))//&
            "] "
        end do
        report=report//"("//therapeutic_bullet_set(i1)%metal//" bullet)"//new_line("a")//repeat("-",80)//new_line("a")
    end do
    report=report//"found therapeutic bullets: "//int2char(size(therapeutic_bullet_set))//" ("//int2char(n_gold)//&
    " golden bullets, "//int2char(n_silv)//" silver bullets)"
    write (unit=*,fmt="(a)") new_line("a")//report//new_line("a")
    write (unit=*,fmt="(a)") "save? [1/0]"//new_line("a")
    read (unit=*,fmt=*) save_
    if (save_==1) then
        open (unit=1,file="report_therapeutic_bullet",status="replace")
        write (unit=1,fmt="(a)") report
        close (unit=1)
        write (unit=*,fmt="(a)") new_line("a")//"report saved as: report_therapeutic_bullet"//new_line("a")
    end if
    deallocate(report)
end subroutine report_therapeutic_bullet_set
!##########################################################################!
!###############################    sort    ###############################!
!##########################################################################!
function sort(x) result(y)
    implicit none
    integer,dimension(:)::x
    integer,dimension(size(x))::y
    integer,dimension(:),allocatable::z
    integer,dimension(:),allocatable::w
    integer,dimension(1)::p
    integer::i
    z=x
    do i=1,size(x)
        p=minloc(z)
        y(i)=z(p(1))
        w=z
        deallocate(z)
        allocate(z(size(x)-i))
        z(:p(1)-1)=w(:p(1)-1)
        z(p(1):)=w(p(1)+1:)
        deallocate(w)
    end do
    deallocate(z)
end function sort
!##########################################################################!
!############################    what_to_do    ############################!
!##########################################################################!
subroutine what_to_do(f,value,size_D,n_node,max_targ,max_moda,V)
    implicit none
    integer::to_do,r_min,r_max,setting,comprehensive_D,size_D,n_node,max_targ,max_moda
    real,dimension(:)::value
    character(16),dimension(:)::V
    real::start,finish
    real,dimension(:,:),allocatable::D
    type(attractor),dimension(:),allocatable::A_set
    integer,dimension(0)::dummy1
    real,dimension(0)::dummy2
    type(attractor),dimension(:),allocatable::A_physio
    type(attractor),dimension(:),allocatable::A_patho
    type(attractor),dimension(:),allocatable::a_patho_set
    type(bullet),dimension(:),allocatable::therapeutic_bullet_set
    interface
        function f(x,k) result(y)
            implicit none
            real,dimension(:,:)::x
            integer::k
            real,dimension(size(x,1),1)::y
        end function f
    end interface
    call init_random_seed()
    call cpu_time(start)
    write (unit=*,fmt="(a)") new_line("a")//"what to do: "//new_line("a")//new_line("a")//"    [1] compute attractors"//&
    new_line("a")//"    [2] compute pathological attractors"//new_line("a")//"    [3] compute therapeutic bullets"//&
    new_line("a")//"    [4] help"//new_line("a")//"    [5] license"//new_line("a")
    read (unit=*,fmt=*) to_do
    if (to_do/=2 .and. to_do/=4 .and. to_do/=5) then
        if (all(value==[0.0,1.0])) then
            write (unit=*,fmt="(a,es10.3e3,a)") new_line("a")//"size(S)=",real(2,8)**real(n_node,8),new_line("a")//&
            "comprehensive_D? [1/0]"//new_line("a")
            read (unit=*,fmt=*) comprehensive_D
            select case (comprehensive_D)
                case (1)
                    D=transpose(generate_state_space(n_node))
                case (0)
                    D=transpose(generate_arrangement(value,n_node,size_D))
            end select
        else
            D=transpose(generate_arrangement(value,n_node,size_D))
        end if
    end if
    select case (to_do)
        case (1)
            A_set=compute_attractor(f,dummy1,dummy2,D)
            write (unit=*,fmt="(a)") new_line("a")//"setting:"//new_line("a")//new_line("a")//"    [1] physiological"//&
            new_line("a")//"    [2] pathological"//new_line("a")
            read (unit=*,fmt=*) setting
            call report_attractor_set(A_set,setting,V)
            deallocate(A_set,D)
        case (2)
            A_physio=load_attractor_set(1)
            A_patho=load_attractor_set(2)
            a_patho_set=compute_pathological_attractor(A_physio,A_patho)
            call report_attractor_set(a_patho_set,3,V)
            deallocate(A_physio,A_patho,a_patho_set)
        case (3)
            A_physio=load_attractor_set(1)
            write (unit=*,fmt="(a)") new_line("a")//"r_min="//new_line("a")
            read (unit=*,fmt=*) r_min
            write (unit=*,fmt="(a)") new_line("a")//"r_max="//new_line("a")
            read (unit=*,fmt=*) r_max
            therapeutic_bullet_set=compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio)
            call report_therapeutic_bullet_set(therapeutic_bullet_set,V)
            deallocate(A_physio,therapeutic_bullet_set,D)
        case (4)
            write (unit=*,fmt="(a)") new_line("a")//"1) do step 1 with f_physio"//new_line("a")//"2) do step 1 with f_patho"//&
            new_line("a")//"3) eventually do step 2"//new_line("a")//"4) do step 3 with f_patho"//new_line("a")//&
            new_line("a")//"do not forget to recompile the sources following any "//"modification"//new_line("a")
        case (5)
            write (unit=*,fmt="(a)") new_line("a")//'Copyright (c) 2013-2014, Arnaud Poret'//new_line("a")//&
            'All rights reserved.'//new_line("a")//new_line("a")//&
            'Redistribution and use in source and binary forms, with or without modification,'//new_line("a")//&
            'are permitted provided that the following conditions are met:'//new_line("a")//new_line("a")//&
            '1. Redistributions of source code must retain the above copyright notice, this'//new_line("a")//&
            'list of conditions and the following disclaimer.'//new_line("a")//new_line("a")//&
            '2. Redistributions in binary form must reproduce the above copyright notice,'//new_line("a")//&
            'this list of conditions and the following disclaimer in the documentation and/or'//new_line("a")//&
            'other materials provided with the distribution.'//new_line("a")//new_line("a")//&
            '3. Neither the name of the copyright holder nor the names of its contributors'//new_line("a")//&
            'may be used to endorse or promote products derived from this software without'//new_line("a")//&
            'specific prior written permission.'//new_line("a")//new_line("a")//&
            'THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND'//new_line("a")//&
            'ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED'//new_line("a")//&
            'WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE'//new_line("a")//&
            'DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR'//new_line("a")//&
            'ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES'//new_line("a")//&
            '(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;'//new_line("a")//&
            'LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON'//new_line("a")//&
            'ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT'//new_line("a")//&
            '(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS'//new_line("a")//&
            'SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.'//new_line("a")
    end select
    call cpu_time(finish)
    write (unit=*,fmt="(a)") "done in "//int2char(int(finish-start))//" seconds"//new_line("a")
end subroutine what_to_do
end module lib
