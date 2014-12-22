! Copyright (C) 2013-2014 Arnaud Poret
! This program is licensed under the GNU General Public License.
! To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.
module lib
    !##############    /!\ networks must be deterministic /!\    ##############!
    integer::max_targ,max_moda,size_D,n_node
    real,dimension(:),allocatable::value
    character(16),dimension(:),allocatable::V
    type::attractor
        real::basin
        real,dimension(:,:),allocatable::mat
        character(16)::name
    end type attractor
    type::bullet
        integer,dimension(:),allocatable::targ
        real,dimension(:),allocatable::moda
        character(16)::metal
        character(16),dimension(:),allocatable::lost
    end type bullet
    contains
    !##########################################################################!
    !############################    compare_a    #############################!
    !##########################################################################!
    function compare_a(a1,a2) result(diff)
        !##########    /!\ attractors must be in sorted form /!\    ###########!
        logical::diff
        type(attractor)::a1,a2
        if (size(a1%mat,2)/=size(a2%mat,2)) then
            diff=.true.
        else
            diff=.not. all(a1%mat==a2%mat)
        end if
    end function compare_a
    !##########################################################################!
    !##########################    compute_A_set    ###########################!
    !##########################################################################!
    function compute_A_set(f,D,setting,ref_set,b) result(A_set)
        logical::a_found,in_A,in_ref,repass
        integer::i1,i2,i3,k,setting
        integer,dimension(:),allocatable::z2
        real,dimension(:,:)::D
        real,dimension(:,:),allocatable::x,z1
        character(16)::a_name
        type(attractor)::a
        type(attractor),dimension(:)::ref_set
        type(attractor),dimension(:),allocatable::A_set
        type(bullet)::b
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1))::y
            end function f
        end interface
        allocate(A_set(0))
        do i1=1,size(D,2)
            a_found=.false.
            x=reshape(D(:,i1),[size(D,1),1])
            k=1
            do
                allocate(z1(size(x,1),size(x,2)+1))
                z1(:,:size(z1,2)-1)=x(:,:)
                z1(:,size(z1,2))=f(x,k)
                x=z1
                deallocate(z1)
                x(b%targ,k+1)=b%moda
                do i2=k,1,-1
                    if (all(x(:,i2)==x(:,k+1))) then
                        a_found=.true.
                        a%mat=reshape(x(:,i2:k),[size(x,1),k-i2+1])
                        z2=range_int(1,size(a%mat,2))
                        do i3=1,size(a%mat,1)
                            z2=z2(minlocs(a%mat(i3,z2)))
                            if (size(z2)==1) then
                                a%mat=cshift(a%mat,z2(1)-1,2)
                                deallocate(z2)
                                exit
                            end if
                        end do
                        exit
                    end if
                end do
                if (a_found) then
                    in_A=.false.
                    do i2=1,size(A_set)
                        if (.not. compare_a(A_set(i2),a)) then
                            in_A=.true.
                            A_set(i2)%basin=A_set(i2)%basin+1.0
                            exit
                        end if
                    end do
                    if (.not. in_A) then
                        a%basin=1.0
                        A_set=[A_set,a]
                    end if
                    exit
                else
                    k=k+1
                end if
            end do
        end do
        deallocate(x)
        do
            repass=.false.
            do i1=1,size(A_set)-1
                if (size(A_set(i1)%mat,2)>size(A_set(i1+1)%mat,2)) then
                    repass=.true.
                    a=A_set(i1)
                    A_set(i1)=A_set(i1+1)
                    A_set(i1+1)=a
                else if (size(A_set(i1)%mat,2)==size(A_set(i1+1)%mat,2)) then
                    do i2=1,size(A_set(i1)%mat,1)
                        if (A_set(i1)%mat(i2,1)<A_set(i1+1)%mat(i2,1)) then
                            exit
                        else if (A_set(i1)%mat(i2,1)>A_set(i1+1)%mat(i2,1)) then
                            repass=.true.
                            a=A_set(i1)
                            A_set(i1)=A_set(i1+1)
                            A_set(i1+1)=a
                            exit
                        end if
                    end do
                end if
            end do
            if (.not. repass) then
                exit
            end if
        end do
        deallocate(a%mat)
        select case (setting)
            case (1)
                a_name="a_physio"
            case (2)
                a_name="a_patho"
        end select
        k=1
        do i1=1,size(A_set)
            A_set(i1)%basin=A_set(i1)%basin*100.0/real(size(D,2))
            in_ref=.false.
            do i2=1,size(ref_set)
                if (.not. compare_a(A_set(i1),ref_set(i2))) then
                    in_ref=.true.
                    exit
                end if
            end do
            if (in_ref) then
                A_set(i1)%name=trim(ref_set(i2)%name)
            else
                A_set(i1)%name=trim(a_name)//int2char(k)
                k=k+1
            end if
        end do
    end function compute_A_set
    !##########################################################################!
    !#######################    compute_a_patho_set    ########################!
    !##########################################################################!
    function compute_a_patho_set(A_patho) result(a_patho_set)
        integer::i
        type(attractor),dimension(:)::A_patho
        type(attractor),dimension(:),allocatable::a_patho_set
        allocate(a_patho_set(0))
        do i=1,size(A_patho)
            if (index(trim(A_patho(i)%name),"patho")/=0) then
                a_patho_set=[a_patho_set,A_patho(i)]
            end if
        end do
    end function compute_a_patho_set
    !##########################################################################!
    !#########################    compute_B_therap    #########################!
    !##########################################################################!
    function compute_B_therap(f,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio) result(B_therap)
        logical::in_patho,golden,repass
        integer::r_min,r_max,max_targ,max_moda,n_node,i1,i2,i3,i4,i5
        integer,dimension(:,:),allocatable::C_targ
        real,dimension(:)::value
        real,dimension(:,:)::D
        real,dimension(:,:),allocatable::C_moda
        type(attractor),dimension(:)::A_physio
        type(attractor),dimension(:),allocatable::A_patho
        type(bullet)::b
        type(bullet),dimension(:),allocatable::B_therap,B_golden,B_silver
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1))::y
            end function f
        end interface
        allocate(B_therap(0))
        do i1=r_min,min(r_max,n_node)
            C_targ=int(gen_combi(real(range_int(1,n_node)),i1,max_targ))
            C_moda=gen_arrang(value,i1,max_moda)
            do i2=1,size(C_targ,1)
                do i3=1,size(C_moda,1)
                    b%targ=C_targ(i2,:)
                    b%moda=C_moda(i3,:)
                    A_patho=compute_A_set(f,D,2,A_physio,b)
                    if (size(compute_a_patho_set(A_patho))==0) then
                        allocate(b%lost(0))
                        if (size(A_patho)/=size(A_physio)) then
                            golden=.false.
                        else
                            golden=.true.
                            do i4=1,size(A_patho)
                                if (compare_a(A_patho(i4),A_physio(i4))) then
                                    golden=.false.
                                    exit
                                end if
                            end do
                        end if
                        if (.not. golden) then
                            b%metal="silver"
                            do i4=1,size(A_physio)
                                in_patho=.false.
                                do i5=1,size(A_patho)
                                    if (.not. compare_a(A_physio(i4),A_patho(i5))) then
                                        in_patho=.true.
                                        exit
                                    end if
                                end do
                                if (.not. in_patho) then
                                    b%lost=[b%lost,trim(A_physio(i4)%name)]
                                end if
                            end do
                        else
                            b%metal="golden"
                        end if
                        B_therap=[B_therap,b]
                        deallocate(b%lost)
                    end if
                end do
            end do
        end do
        deallocate(A_patho,C_targ,C_moda)
        allocate(b%lost(0))
        do
            repass=.false.
            do i1=1,size(B_therap)-1
                if (size(B_therap(i1)%targ)>size(B_therap(i1+1)%targ)) then
                    repass=.true.
                    b=B_therap(i1)
                    B_therap(i1)=B_therap(i1+1)
                    B_therap(i1+1)=b
                else if (size(B_therap(i1)%targ)==size(B_therap(i1+1)%targ)) then
                    if (any(B_therap(i1)%targ/=B_therap(i1+1)%targ)) then
                        do i2=1,size(B_therap(i1)%targ)
                            if (B_therap(i1)%targ(i2)<B_therap(i1+1)%targ(i2)) then
                                exit
                            else if (B_therap(i1)%targ(i2)>B_therap(i1+1)%targ(i2)) then
                                repass=.true.
                                b=B_therap(i1)
                                B_therap(i1)=B_therap(i1+1)
                                B_therap(i1+1)=b
                                exit
                            end if
                        end do
                    else
                        do i2=1,size(B_therap(i1)%moda)
                            if (B_therap(i1)%moda(i2)<B_therap(i1+1)%moda(i2)) then
                                exit
                            else if (B_therap(i1)%moda(i2)>B_therap(i1+1)%moda(i2)) then
                                repass=.true.
                                b=B_therap(i1)
                                B_therap(i1)=B_therap(i1+1)
                                B_therap(i1+1)=b
                                exit
                            end if
                        end do
                    end if
                end if
            end do
            if (.not. repass) then
                exit
            end if
        end do
        deallocate(b%targ,b%moda,b%lost)
        allocate(B_golden(0))
        allocate(B_silver(0))
        do i1=1,size(B_therap)
            if (trim(B_therap(i1)%metal)=="golden") then
                B_golden=[B_golden,B_therap(i1)]
            else
                B_silver=[B_silver,B_therap(i1)]
            end if
        end do
        B_therap=[B_golden,B_silver]
        deallocate(B_golden,B_silver)
    end function compute_B_therap
    !##########################################################################!
    !##############################    facto    ###############################!
    !##########################################################################!
    function facto(x) result(y)
        integer::x,i
        real(8)::y
        if (x>170) then
            write (unit=*,fmt="(a)") "facto(x): x>170 unsupported."//new_line("a")
            stop
        else if (x==0) then
            y=real(1,8)
        else
            y=real(x,8)
            do i=1,x-1
                y=y*(real(x,8)-real(i,8))
            end do
        end if
    end function facto
    !##########################################################################!
    !############################    gen_arrang    ############################!
    !##########################################################################!
    function gen_arrang(deck,k,n_arrang) result(arrang)
        !#################    /!\ only with repetition /!\    #################!
        logical::in_arrang
        integer::k,n_arrang,i1,i2
        real,dimension(k)::z
        real,dimension(:)::deck
        real,dimension(:,:),allocatable::arrang
        allocate(arrang(min(n_arrang,int(min(real(size(deck),8)**real(k,8),real(huge(1),8)))),k))
        do i1=1,size(arrang,1)
            do
                do i2=1,k
                    z(i2)=deck(rand_int(1,size(deck)))
                end do
                in_arrang=.false.
                do i2=1,i1-1
                    if (all(arrang(i2,:)==z)) then
                        in_arrang=.true.
                        exit
                    end if
                end do
                if (.not. in_arrang) then
                    arrang(i1,:)=z
                    exit
                end if
            end do
        end do
    end function gen_arrang
    !##########################################################################!
    !############################    gen_combi    #############################!
    !##########################################################################!
    function gen_combi(deck,k,n_combi) result(combi)
        !###############    /!\ only without repetition /!\    ################!
        logical::in_combi
        integer::k,n_combi,i1,i2
        real::z1
        real,dimension(k)::z2
        real,dimension(:)::deck
        real,dimension(:,:),allocatable::combi
        allocate(combi(min(n_combi,int(min(facto(size(deck))/(facto(k)*facto(size(deck)-k)),real(huge(1),8)))),k))
        do i1=1,size(combi,1)
            do
                do i2=1,k
                    do
                        z1=deck(rand_int(1,size(deck)))
                        if (all(z1/=z2(:i2-1))) then
                            z2(i2)=z1
                            exit
                        end if
                    end do
                end do
                z2=sort(z2)
                in_combi=.false.
                do i2=1,i1-1
                    if (all(combi(i2,:)==z2)) then
                        in_combi=.true.
                        exit
                    end if
                end do
                if (.not. in_combi) then
                    combi(i1,:)=z2
                    exit
                end if
            end do
        end do
    end function gen_combi
    !##########################################################################!
    !##############################    gen_S    ###############################!
    !##########################################################################!
    function gen_S(n) result(S)
        !#####################    /!\ boolean only /!\    #####################!
        integer::n,i1,i2
        real,dimension(n,2**n)::S
        if (n>30) then
            write (unit=*,fmt="(a)") "gen_S(n): n>30 unsupported."//new_line("a")
            stop
        else
            do i1=1,n
                S(:i1-1,(2**i1)/2+1:2**i1)=S(:i1-1,:2**(i1-1))
                do i2=1,2**(i1-1)
                    S(i1,i2)=0.0
                end do
                do i2=(2**i1)/2+1,2**i1
                    S(i1,i2)=1.0
                end do
            end do
        end if
    end function gen_S
    !##########################################################################!
    !#########################    init_random_seed    #########################!
    !##########################################################################!
    subroutine init_random_seed()
        integer::seed_size,error
        integer,dimension(:),allocatable::seed
        open(unit=1,file="/dev/urandom",status="old",access="stream",form="unformatted",action="read",iostat=error)
        if (error==0) then
            call random_seed(size=seed_size)
            allocate(seed(seed_size))
            read (unit=1) seed
            close (unit=1)
            call random_seed(put=seed)
            deallocate(seed)
        else
            write (unit=*,fmt="(a)") "Too bad, your operating system does not provide a random number generator."//new_line("a")
            stop
        end if
    end subroutine init_random_seed
    !##########################################################################!
    !#############################    int2char    #############################!
    !##########################################################################!
    function int2char(x) result(y)
        integer::x
        character(11)::z
        character(:),allocatable::y
        write (unit=z,fmt="(i11)") x
        y=trim(adjustl(z))
    end function int2char
    !##########################################################################!
    !############################    load_A_set    ############################!
    !##########################################################################!
    function load_A_set(setting) result(A_set)
        integer::setting,i1,i2,z,n,m
        character(32)::set_name
        type(attractor),dimension(:),allocatable::A_set
        select case (setting)
            case (1)
                set_name="A_physio.csv"
            case (2)
                set_name="A_patho.csv"
        end select
        open (unit=1,file=trim(set_name),status="old")
        read (unit=1,fmt=*) z
        allocate(A_set(z))
        do i1=1,size(A_set)
            read (unit=1,fmt=*) n
            read (unit=1,fmt=*) m
            read (unit=1,fmt=*) A_set(i1)%basin
            read (unit=1,fmt=*) A_set(i1)%name
            allocate(A_set(i1)%mat(n,m))
        end do
        do i1=1,size(A_set)
            do i2=1,size(A_set(i1)%mat,1)
                read (unit=1,fmt=*) A_set(i1)%mat(i2,:)
            end do
        end do
        close (unit=1)
    end function load_A_set
    !##########################################################################!
    !#############################    minlocs    ##############################!
    !##########################################################################!
    function minlocs(x) result(y)
        integer::i,i_min
        integer,dimension(:),allocatable::y
        real,dimension(:)::x
        i_min=minloc(x,1)
        y=[i_min]
        do i=i_min+1,size(x)
            if (x(i)==x(i_min)) then
                y=[y,i]
            end if
        end do
    end function minlocs
    !##########################################################################!
    !#############################    rand_int    #############################!
    !##########################################################################!
    function rand_int(a,b) result(y)
        integer::a,b,y
        real::x
        call random_number(x)
        y=nint(real(a)+x*(real(b)-real(a)))
    end function rand_int
    !##########################################################################!
    !############################    range_int    #############################!
    !##########################################################################!
    function range_int(a,b) result(y)
        integer::a,b,i
        integer,dimension(b-a+1)::y
        do i=1,b-a+1
            y(i)=a+i-1
        end do
    end function range_int
    !##########################################################################!
    !############################    real2char    #############################!
    !##########################################################################!
    function real2char(x) result(y)
        real::x
        character(42)::z
        character(:),allocatable::y
        write (unit=z,fmt="(f42.1)") x
        y=trim(adjustl(z))
    end function real2char
    !##########################################################################!
    !###########################    report_A_set    ###########################!
    !##########################################################################!
    subroutine report_A_set(A_set,setting,V,bool)
        logical::bool
        integer::setting,n_point,n_cycle,i1,i2,i3,save_
        character(16)::set_type
        character(16),dimension(:)::V
        character(32)::set_name,report_name
        character(:),allocatable::report,s
        type(attractor),dimension(:)::A_set
        n_point=0
        n_cycle=0
        select case (setting)
            case (1)
                set_type="A_physio"
            case (2)
                set_type="A_patho"
            case (3)
                set_type="A_versus"
        end select
        report=trim(set_type)//"={"
        do i1=1,size(A_set)-1
            report=report//trim(A_set(i1)%name)//","
        end do
        report=report//trim(A_set(size(A_set))%name)//"}"//new_line("a")//repeat("-",80)//new_line("a")
        do i1=1,size(A_set)
            if (size(A_set(i1)%mat,2)==1) then
                n_point=n_point+1
            else
                n_cycle=n_cycle+1
            end if
            report=report//trim(A_set(i1)%name)//new_line("a")//"basin: "//real2char(A_set(i1)%basin)//"% (of the state space)"//new_line("a")
            do i2=1,size(A_set(i1)%mat,1)
                report=report//V(i2)//" "
                do i3=1,size(A_set(i1)%mat,2)-1
                    if (bool) then
                        report=report//int2char(int(A_set(i1)%mat(i2,i3)))//" "
                    else
                        report=report//real2char(A_set(i1)%mat(i2,i3))//" "
                    end if
                end do
                if (bool) then
                    report=report//int2char(int(A_set(i1)%mat(i2,size(A_set(i1)%mat,2))))//new_line("a")
                else
                    report=report//real2char(A_set(i1)%mat(i2,size(A_set(i1)%mat,2)))//new_line("a")
                end if
            end do
            report=report//repeat("-",80)//new_line("a")
        end do
        report=report//"found attractors: "//int2char(size(A_set))//" ("//int2char(n_point)//" points, "//int2char(n_cycle)//" cycles)"
        write (unit=*,fmt="(a)") new_line("a")//report//new_line("a")//new_line("a")//"Save?"//new_line("a")//"    [0] no"//new_line("a")//"    [1] yes"
        read (unit=*,fmt=*) save_
        if (save_==1) then
            select case (setting)
                case (1)
                    set_name="A_physio.csv"
                    report_name="A_physio.txt"
                case (2)
                    set_name="A_patho.csv"
                    report_name="A_patho.txt"
                case (3)
                    set_name="A_versus.csv"
                    report_name="A_versus.txt"
            end select
            s=int2char(size(A_set))//new_line("a")
            do i1=1,size(A_set)
                s=s//int2char(size(A_set(i1)%mat,1))//new_line("a")//int2char(size(A_set(i1)%mat,2))//new_line("a")//real2char(A_set(i1)%basin)//new_line("a")//trim(A_set(i1)%name)//new_line("a")
            end do
            do i1=1,size(A_set)
                do i2=1,size(A_set(i1)%mat,1)
                    do i3=1,size(A_set(i1)%mat,2)-1
                        s=s//real2char(A_set(i1)%mat(i2,i3))//","
                    end do
                    s=s//real2char(A_set(i1)%mat(i2,size(A_set(i1)%mat,2)))
                    if (i1/=size(A_set) .or. i2/=size(A_set(i1)%mat,1)) then
                        s=s//new_line("a")
                    end if
                end do
            end do
            open (unit=1,file=trim(set_name),status="replace")
            write (unit=1,fmt="(a)") s
            close (unit=1)
            open (unit=1,file=trim(report_name),status="replace")
            write (unit=1,fmt="(a)") report
            close (unit=1)
            write (unit=*,fmt="(a)") new_line("a")//"Set saved as "//trim(set_name)//"."//new_line("a")//"Report saved as "//trim(report_name)//"."
            deallocate(s)
        end if
        deallocate(report)
    end subroutine report_A_set
    !##########################################################################!
    !#########################    report_B_therap    ##########################!
    !##########################################################################!
    subroutine report_B_therap(B_therap,V,bool)
        logical::bool
        integer::n_gold,n_silv,i1,i2,save_
        character(1)::moda
        character(16),dimension(:)::V
        character(:),allocatable::report
        type(bullet),dimension(:)::B_therap
        n_gold=0
        n_silv=0
        report=repeat("-",80)//new_line("a")
        do i1=1,size(B_therap)
            if (trim(B_therap(i1)%metal)=="golden") then
                n_gold=n_gold+1
            else
                n_silv=n_silv+1
            end if
            do i2=1,size(B_therap(i1)%targ)
                if (bool) then
                    if (B_therap(i1)%moda(i2)==1.0) then
                        moda="+"
                    else
                        moda="-"
                    end if
                    report=report//moda//trim(V(B_therap(i1)%targ(i2)))//" "
                else
                    report=report//trim(V(B_therap(i1)%targ(i2)))//"["//real2char(B_therap(i1)%moda(i2))//"] "
                end if
            end do
            report=report//"("//trim(B_therap(i1)%metal)
            if (trim(B_therap(i1)%metal)=="silver") then
                report=report//", unrecovered:"
                do i2=1,size(B_therap(i1)%lost)-1
                    report=report//" "//trim(B_therap(i1)%lost(i2))//","
                end do
                report=report//" "//trim(B_therap(i1)%lost(size(B_therap(i1)%lost)))//")"//new_line("a")//repeat("-",80)//new_line("a")
            else
                report=report//")"//new_line("a")//repeat("-",80)//new_line("a")
            end if
        end do
        report=report//"found therapeutic bullets: "//int2char(size(B_therap))//" ("//int2char(n_gold)//" golden bullets, "//int2char(n_silv)//" silver bullets)"
        write (unit=*,fmt="(a)") new_line("a")//report//new_line("a")//new_line("a")//"Save?"//new_line("a")//"    [0] no"//new_line("a")//"    [1] yes"
        read (unit=*,fmt=*) save_
        if (save_==1) then
            open (unit=1,file="B_therap.txt",status="replace")
            write (unit=1,fmt="(a)") report
            close (unit=1)
            write (unit=*,fmt="(a)") new_line("a")//"Report saved as B_therap.txt."
        end if
        deallocate(report)
    end subroutine report_B_therap
    !##########################################################################!
    !###############################    sort    ###############################!
    !##########################################################################!
    function sort(x) result(y)
        integer::i,i_min
        real::z
        real,dimension(:)::x
        real,dimension(size(x))::y
        y=x
        do i=1,size(y)-1
            i_min=minloc(y(i:),1)+i-1
            if (i_min/=i) then
                z=y(i)
                y(i)=y(i_min)
                y(i_min)=z
            end if
        end do
    end function sort
    !##########################################################################!
    !############################    what_to_do    ############################!
    !##########################################################################!
    subroutine what_to_do(f1,f2,value,size_D,n_node,max_targ,max_moda,V)
        logical::bool,z
        integer::size_D,n_node,max_targ,max_moda,to_do,r_min,r_max,setting,whole_S
        real,dimension(:)::value
        real,dimension(:,:),allocatable::D
        character(16),dimension(:)::V
        type(attractor),dimension(0)::null_set
        type(attractor),dimension(:),allocatable::A_physio,A_patho,a_patho_set
        type(bullet)::null_b
        type(bullet),dimension(:),allocatable::B_therap
        interface
            function f1(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1))::y
            end function f1
        end interface
        interface
            function f2(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1))::y
            end function f2
        end interface
        call init_random_seed()
        if (size(value)==2) then
            bool=all(value==[0.0,1.0])
        else
            bool=.false.
        end if
        do
            write (unit=*,fmt="(a)") new_line("a")//"What to do?"//new_line("a")//"    [1] compute an attractor set"//new_line("a")//"    [2] compute the pathological attractors"//new_line("a")//"    [3] compute the therapeutic bullets"//new_line("a")//"    [4] show the help"//new_line("a")//"    [5] show the license"//new_line("a")//"    [6] quit the algorithm"
            read (unit=*,fmt=*) to_do
            if (to_do==1 .or. to_do==3) then
                if (bool) then
                    write (unit=*,fmt="(a,es10.3e3,a)") new_line("a")//"State space cardinality: ",real(2,8)**real(n_node,8),", compute the whole state space?"//new_line("a")//"    [0] no"//new_line("a")//"    [1] yes"
                    read (unit=*,fmt=*) whole_S
                    select case (whole_S)
                        case (1)
                            D=gen_S(n_node)
                        case (0)
                            D=transpose(gen_arrang(value,n_node,size_D))
                    end select
                else
                    D=transpose(gen_arrang(value,n_node,size_D))
                end if
            end if
            select case (to_do)
                case (1)
                    allocate(null_b%targ(0))
                    allocate(null_b%moda(0))
                    write (unit=*,fmt="(a)") new_line("a")//"Setting?"//new_line("a")//"    [1] physiological"//new_line("a")//"    [2] pathological"
                    read (unit=*,fmt=*) setting
                    select case (setting)
                        case (1)
                            A_physio=compute_A_set(f1,D,1,null_set,null_b)
                            call report_A_set(A_physio,1,V,bool)
                            deallocate(A_physio)
                        case (2)
                            inquire (file="A_physio.csv",exist=z)
                            if (.not. z) then
                                write (unit=*,fmt="(a)") new_line("a")//"The file A_physio.csv does not exist. Please compute the physiological attractor"//new_line("a")//"set before computing the pathological attractor set."
                            else
                                A_physio=load_A_set(1)
                                A_patho=compute_A_set(f2,D,2,A_physio,null_b)
                                call report_A_set(A_patho,2,V,bool)
                                deallocate(A_physio,A_patho)
                            end if
                    end select
                    deallocate(D,null_b%targ,null_b%moda)
                case (2)
                    inquire (file="A_patho.csv",exist=z)
                    if (.not. z) then
                        write (unit=*,fmt="(a)") new_line("a")//"The file A_patho.csv does not exist. Please compute the pathological attractor"//new_line("a")//"set before computing the pathological attractors."
                    else
                        A_patho=load_A_set(2)
                        a_patho_set=compute_a_patho_set(A_patho)
                        call report_A_set(a_patho_set,3,V,bool)
                        deallocate(A_patho,a_patho_set)
                    end if
                case (3)
                    inquire (file="A_physio.csv",exist=z)
                    if (.not. z) then
                        write (unit=*,fmt="(a)") new_line("a")//"The file A_physio.csv does not exist. Please compute the physiological attractor"//new_line("a")//"set before computing the therapeutic bullets."
                    else
                        A_physio=load_A_set(1)
                        write (unit=*,fmt="(a)",advance="no") new_line("a")//"Number of targets per bullet (lower bound): "
                        read (unit=*,fmt=*) r_min
                        write (unit=*,fmt="(a)",advance="no") "Number of targets per bullet (upper bound): "
                        read (unit=*,fmt=*) r_max
                        B_therap=compute_B_therap(f2,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio)
                        call report_B_therap(B_therap,V,bool)
                        deallocate(A_physio,B_therap)
                    end if
                    deallocate(D)
                case (4)
                    write (unit=*,fmt="(a)") new_line("a")//"How to:"//new_line("a")//"    1) compute the physiological attractor set ([1], returns A_physio)"//new_line("a")//"    2) compute the pathological attractor set ([1], returns A_patho)"//new_line("a")//"    3) compute the pathological attractors ([2], returns A_versus)"//new_line("a")//"    4) compute the therapeutic bullets ([3], returns B_therap)"//new_line("a")//new_line("a")//"Do not forget to recompile the sources following any modification."
                case (5)
                    write (unit=*,fmt="(a)") new_line("a")//"kali-targ: a tool for in silico target identification."//new_line("a")//"Copyright (C) 2013-2014 Arnaud Poret"//new_line("a")//new_line("a")//"This program is free software: you can redistribute it and/or modify it under"//new_line("a")//"the terms of the GNU General Public License as published by the Free Software"//new_line("a")//"Foundation, either version 3 of the License, or (at your option) any later"//new_line("a")//"version."//new_line("a")//new_line("a")//"This program is distributed in the hope that it will be useful, but WITHOUT ANY"//new_line("a")//"WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A"//new_line("a")//"PARTICULAR PURPOSE. See the GNU General Public License for more details."//new_line("a")//new_line("a")//"You should have received a copy of the GNU General Public License along with"//new_line("a")//"this program. If not, see https://www.gnu.org/licenses/gpl.html."
                case (6)
                    write (unit=*,fmt="(a)") new_line("a")//"Goodbye."//new_line("a")
                    exit
            end select
        end do
    end subroutine what_to_do
end module lib
