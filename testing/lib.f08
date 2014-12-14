! Copyright (C) 2013-2014 Arnaud Poret
! This program is licensed under the GNU General Public License.
! To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.
module lib
    !##############    /!\ networks must be deterministic /!\    ##############!
    integer::max_targ,max_moda,size_D,n_node
    real,dimension(:),allocatable::value
    character(16),dimension(:),allocatable::V
    type::attractor
        real::popularity
        real,dimension(:,:),allocatable::a
    end type attractor
    type::bullet
        integer,dimension(:),allocatable::targ
        real,dimension(2)::gain
        real,dimension(:),allocatable::moda
    end type bullet
    contains
    !##########################################################################!
    !##########################    add_attractor    ###########################!
    !##########################################################################!
    function add_attractor(A_set,a,popularity) result(y)
        real::popularity
        real,dimension(:,:)::a
        type(attractor),dimension(:)::A_set
        type(attractor),dimension(size(A_set)+1)::y
        y(:size(A_set))=A_set
        y(size(A_set)+1)%a=a
        y(size(A_set)+1)%popularity=popularity
    end function add_attractor
    !##########################################################################!
    !############################    add_bullet    ############################!
    !##########################################################################!
    function add_bullet(bullet_set,targ,moda,gain) result(y)
        integer,dimension(:)::targ
        real,dimension(2)::gain
        real,dimension(:)::moda
        type(bullet),dimension(:)::bullet_set
        type(bullet),dimension(size(bullet_set)+1)::y
        y(:size(bullet_set))=bullet_set
        y(size(bullet_set)+1)%targ=targ
        y(size(bullet_set)+1)%moda=moda
        y(size(bullet_set)+1)%gain=gain
    end function add_bullet
    !##########################################################################!
    !########################    compare_attractor    #########################!
    !##########################################################################!
    function compare_attractor(a1,a2) result(differ)
        logical::differ,start_found
        integer::i1,i2,start1,start2
        real,dimension(:,:)::a1,a2
        differ=.false.
        if (size(a1,2)/=size(a2,2)) then
            differ=.true.
        else
            start_found=.false.
            do1:do i1=1,size(a1,2)
                do2:do i2=1,size(a2,2)
                    if (all(a1(:,i1)==a2(:,i2))) then
                        start_found=.true.
                        start1=i1
                        start2=i2
                        exit do1
                    end if
                end do do2
            end do do1
            if (.not. start_found) then
                differ=.true.
            else
                do i1=0,size(a1,2)-2
                    if (.not. all(a1(:,modulo(start1+i1,size(a1,2))+1)==a2(:,modulo(start2+i1,size(a2,2))+1))) then
                        differ=.true.
                        exit
                    end if
                end do
            end if
        end if
    end function compare_attractor
    !##########################################################################!
    !######################    compare_attractor_set    #######################!
    !##########################################################################!
!    function compare_attractor_set(A_set1,A_set2) result(differ)
!        logical::differ,z
!        integer::i1,i2
!        type(attractor),dimension(:)::A_set1,A_set2
!        logical,dimension(size(A_set1))::in_2
!        if (size(A_set1)/=size(A_set2)) then
!            differ=.true.
!        else
!            do i1=1,size(A_set1)
!                z=.false.
!                do i2=1,size(A_set2)
!                    if (.not. compare_attractor(A_set1(i1)%a,A_set2(i2)%a)) then
!                        z=.true.
!                        exit
!                    end if
!                end do
!                in_2(i1)=z
!            end do
!            differ=.not. all(in_2)
!        end if
!    end function compare_attractor_set
    !##########################################################################!
    !########################    compute_attractor    #########################!
    !##########################################################################!
    function compute_attractor(f,c_targ,c_moda,D) result(A_set)
        logical::a_found,in_A
        integer::i1,i2,k
        integer,dimension(:)::c_targ
        integer,dimension(:),allocatable::count
        real,dimension(:)::c_moda
        real,dimension(:,:)::D
        real,dimension(:,:),allocatable::a,x
        type(attractor),dimension(:),allocatable::A_set
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        allocate(A_set(0))
        allocate(count(0))
        do i1=1,size(D,2)
            x=reshape(D(:,i1),[size(D,1),1])
            k=1
            do
                x=concatenate(x,f(x,k),2)
                x(c_targ,k+1)=c_moda
                a_found=.false.
                do i2=k,1,-1
                    if (all(x(:,i2)==x(:,k+1))) then
                        a_found=.true.
                        a=reshape(x(:,i2:k),[size(D,1),k-i2+1])
                        exit
                    end if
                end do
                if (a_found) then
                    in_A=.false.
                    do i2=1,size(A_set)
                        if (.not. compare_attractor(a,A_set(i2)%a)) then
                            in_A=.true.
                            count(i2)=count(i2)+1
                            exit
                        end if
                    end do
                    if (.not. in_A) then
                        A_set=add_attractor(A_set,a,0.0)
                        count=int(reshape(concatenate(real(reshape(count,[1,size(count)])),reshape([1.0],[1,1]),2),[size(count)+1]))
                    end if
                    exit
                end if
                k=k+1
            end do
            deallocate(x,a)
        end do
        do i1=1,size(A_set)
            A_set(i1)%popularity=(real(count(i1))/real(size(D,2)))*100.0
        end do
        deallocate(count)
    end function compute_attractor
    !##########################################################################!
    !##################    compute_pathological_attractor    ##################!
    !##########################################################################!
    function compute_pathological_attractor(A_physio,A_patho) result(a_patho_set)
        logical::in_physio
        integer::i1,i2
        type(attractor),dimension(:)::A_physio,A_patho
        type(attractor),dimension(:),allocatable::a_patho_set
        allocate(a_patho_set(0))
        do i1=1,size(A_patho)
            in_physio=.false.
            do i2=1,size(A_physio)
                if (.not. compare_attractor(A_patho(i1)%a,A_physio(i2)%a)) then
                    in_physio=.true.
                    exit
                end if
            end do
            if (.not. in_physio) then
                a_patho_set=add_attractor(a_patho_set,A_patho(i1)%a,A_patho(i1)%popularity)
            end if
        end do
    end function compute_pathological_attractor
    !##########################################################################!
    !####################    compute_therapeutic_bullet    ####################!
    !##########################################################################!
    function compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio,A_patho) result(therapeutic_bullet_set)
        integer::r_min,r_max,i1,i2,i3,i4,i5,max_targ,max_moda,n_node
        integer,dimension(:,:),allocatable::C_targ
        real::patho_covering,test_covering
        real,dimension(:)::value
        real,dimension(:,:)::D
        real,dimension(:,:),allocatable::C_moda
        type(attractor),dimension(:)::A_physio,A_patho
        type(attractor),dimension(:),allocatable::A_test
        type(bullet),dimension(:),allocatable::therapeutic_bullet_set
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        allocate(therapeutic_bullet_set(0))
        patho_covering=0.0
        do i1=1,size(A_patho)
            do i2=1,size(A_physio)
                if (.not. compare_attractor(A_patho(i1)%a,A_physio(i2)%a)) then
                    patho_covering=patho_covering+A_patho(i1)%popularity
                end if
            end do
        end do
        do i1=r_min,min(r_max,n_node)
            C_targ=int(generate_combination(real(range_int(1,n_node)),i1,max_targ))
            C_moda=generate_arrangement(value,i1,max_moda)
            do i2=1,size(C_targ,1)
                do i3=1,size(C_moda,1)
                    A_test=compute_attractor(f,C_targ(i2,:),C_moda(i3,:),D)
                    test_covering=0.0
                    do i4=1,size(A_test)
                        do i5=1,size(A_physio)
                            if (.not. compare_attractor(A_test(i4)%a,A_physio(i5)%a)) then
                                test_covering=test_covering+A_test(i4)%popularity
                            end if
                        end do
                    end do
                    if (test_covering>patho_covering) then
                        therapeutic_bullet_set=add_bullet(therapeutic_bullet_set,C_targ(i2,:),C_moda(i3,:),[patho_covering,test_covering])
                    end if
                    deallocate(A_test)
                end do
            end do
            deallocate(C_targ,C_moda)
        end do
    end function compute_therapeutic_bullet
    !##########################################################################!
    !###########################    concatenate    ############################!
    !##########################################################################!
    function concatenate(x1,x2,d) result(y)
        integer::d
        real,dimension(:,:)::x1,x2
        real,dimension(:,:),allocatable::y
        select case (d)
            case (1)
                if (size(x1,1)==0) then
                    y=x2
                else if (size(x2,1)==0) then
                    y=x1
                else
                    allocate(y(size(x1,1)+size(x2,1),size(x1,2)))
                    y(1:size(x1,1),:)=x1
                    y(size(x1,1)+1:size(x1,1)+size(x2,1),:)=x2
                end if
            case (2)
                if (size(x1,2)==0) then
                    y=x2
                else if (size(x2,2)==0) then
                    y=x1
                else
                    allocate(y(size(x1,1),size(x1,2)+size(x2,2)))
                    y(:,1:size(x1,2))=x1
                    y(:,size(x1,2)+1:size(x1,2)+size(x2,2))=x2
                end if
        end select
    end function concatenate
    !##########################################################################!
    !##############################    facto    ###############################!
    !##########################################################################!
    function facto(x) result(y)
        integer::x,i
        real(8)::y
        if (x>170) then
            write (unit=*,fmt="(a)") "facto(x): x>170 unsupported"
            stop
        end if
        if (x==0) then
            y=real(1,8)
        else
            y=real(x,8)
            do i=1,x-1
                y=y*real(x-i,8)
            end do
        end if
    end function facto
    !##########################################################################!
    !#######################    generate_arrangement    #######################!
    !##########################################################################!
    function generate_arrangement(deck,k,n_arrang) result(arrang_mat)
        !#################    /!\ only with repetition /!\    #################!
        integer::k,i1,i2,n_arrang
        real,dimension(k)::arrang
        real,dimension(:)::deck
        real,dimension(:,:),allocatable::arrang_mat
        allocate(arrang_mat(min(n_arrang,int(min(real(size(deck),8)**real(k,8),real(huge(1),8)))),k))
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
        integer::k,i1,i2,n_combi
        real::z
        real,dimension(k)::combi
        real,dimension(:)::deck
        real,dimension(:,:),allocatable::combi_mat
        allocate(combi_mat(min(n_combi,int(min(facto(size(deck))/(facto(k)*facto(size(deck)-k)),real(huge(1),8)))),k))
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
        integer::n,i1,i2
        real,dimension(n,2**n)::y
        if (n>30) then
            write (unit=*,fmt="(a)") "generate_state_space(n): n>30 unsupported"
            stop
        end if
        do i1=1,n
            y(:i1-1,(2**i1)/2+1:2**i1)=y(:i1-1,:2**(i1-1))
            do i2=1,2**(i1-1)
                y(i1,i2)=0.0
            end do
            do i2=(2**i1)/2+1,2**i1
                y(i1,i2)=1.0
            end do
        end do
    end function generate_state_space
    !##########################################################################!
    !#########################    init_random_seed    #########################!
    !##########################################################################!
    subroutine init_random_seed()
        integer::seed_size,error
        integer,dimension(:),allocatable::seed
        call random_seed(size=seed_size)
        allocate(seed(seed_size))
        open(unit=1,file="/dev/urandom",status="old",access="stream",form="unformatted",action="read",iostat=error)
        if (error==0) then
            read (unit=1) seed
            close (unit=1)
        else
            write (unit=*,fmt="(a)") "Too bad, your operating system does not provide a random number generator."
            stop
        end if
        call random_seed(put=seed)
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
    !########################    load_attractor_set    ########################!
    !##########################################################################!
    function load_attractor_set(setting) result(A_set)
        integer::setting,i1,i2,z,n,m
        character(32)::set_name
        type(attractor),dimension(:),allocatable::A_set
        select case (setting)
            case (1)
                set_name="set_physio.csv"
            case (2)
                set_name="set_patho.csv"
        end select
        open (unit=1,file=trim(set_name),status="old")
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
                read (unit=1,fmt=*) A_set(i1)%a(i2,:)
            end do
        end do
        close (unit=1)
    end function load_attractor_set
    !##########################################################################!
    !#############################    rand_int    #############################!
    !##########################################################################!
    function rand_int(a,b) result(y)
        integer::a,b,y
        real::x
        call random_number(x)
        y=nint(real(a)+x*(real(b-a)))
    end function rand_int
    !##########################################################################!
    !############################    range_int    #############################!
    !##########################################################################!
    function range_int(a,b) result(y)
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
        real::x
        character(43)::z
        character(:),allocatable::y
        write (unit=z,fmt="(f43.2)") x
        y=trim(adjustl(z))
    end function real2char
    !##########################################################################!
    !#######################    report_attractor_set    #######################!
    !##########################################################################!
    subroutine report_attractor_set(A_set,setting,V,boolean)
        logical::boolean
        integer::setting,n_point,n_cycle,i1,i2,i3,save_
        character(32)::set_name,report_name
        character(16),dimension(:)::V
        character(:),allocatable::report,s
        type(attractor),dimension(:)::A_set
        n_point=0
        n_cycle=0
        report=repeat("-",80)//new_line("a")
        do i1=1,size(A_set)
            if (size(A_set(i1)%a,2)==1) then
                n_point=n_point+1
            else
                n_cycle=n_cycle+1
            end if
            report=report//"basin: "//real2char(A_set(i1)%popularity)//"% (of the state space)"//new_line("a")//new_line("a")
            do i2=1,size(A_set(i1)%a,1)
                report=report//V(i2)//" "
                do i3=1,size(A_set(i1)%a,2)-1
                    if (boolean) then
                        report=report//int2char(int(A_set(i1)%a(i2,i3)))//" "
                    else
                        report=report//real2char(A_set(i1)%a(i2,i3))//" "
                    end if
                end do
                if (boolean) then
                    report=report//int2char(int(A_set(i1)%a(i2,size(A_set(i1)%a,2))))//new_line("a")
                else
                    report=report//real2char(A_set(i1)%a(i2,size(A_set(i1)%a,2)))//new_line("a")
                end if
            end do
            report=report//repeat("-",80)//new_line("a")
        end do
        report=report//"found attractors: "//int2char(size(A_set))//" ("//int2char(n_point)//" points, "//int2char(n_cycle)//" cycles)"
        write (unit=*,fmt="(a)",advance="no") new_line("a")//report//new_line("a")//new_line("a")//"save [1/0] "
        read (unit=*,fmt=*) save_
        if (save_==1) then
            select case (setting)
                case (1)
                    set_name="set_physio.csv"
                    report_name="report_physio.txt"
                case (2)
                    set_name="set_patho.csv"
                    report_name="report_patho.txt"
                case (3)
                    set_name="set_versus.csv"
                    report_name="report_versus.txt"
            end select
            s=int2char(size(A_set))//new_line("a")
            do i1=1,size(A_set)
                s=s//int2char(size(A_set(i1)%a,1))//new_line("a")//int2char(size(A_set(i1)%a,2))//new_line("a")//real2char(A_set(i1)%popularity)//new_line("a")
            end do
            do i1=1,size(A_set)
                do i2=1,size(A_set(i1)%a,1)
                    do i3=1,size(A_set(i1)%a,2)-1
                        s=s//real2char(A_set(i1)%a(i2,i3))//","
                    end do
                    s=s//real2char(A_set(i1)%a(i2,size(A_set(i1)%a,2)))
                    if (i1/=size(A_set) .or. i2/=size(A_set(i1)%a,1)) then
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
            write (unit=*,fmt="(a)") new_line("a")//"set saved as: "//trim(set_name)//new_line("a")//"report saved as: "//trim(report_name)
            deallocate(s)
        end if
        deallocate(report)
    end subroutine report_attractor_set
    !##########################################################################!
    !##################    report_therapeutic_bullet_set    ###################!
    !##########################################################################!
    subroutine report_therapeutic_bullet_set(therapeutic_bullet_set,V,boolean)
        logical::boolean
        integer::i1,i2,save_
        character(1)::moda
        character(16),dimension(:)::V
        character(:),allocatable::report
        type(bullet),dimension(:)::therapeutic_bullet_set
        report=repeat("-",80)//new_line("a")
        do i1=1,size(therapeutic_bullet_set)
            do i2=1,size(therapeutic_bullet_set(i1)%targ)
                if (boolean) then
                    if (therapeutic_bullet_set(i1)%moda(i2)==1.0) then
                        moda="+"
                    else
                        moda="-"
                    end if
                    report=report//moda//trim(V(therapeutic_bullet_set(i1)%targ(i2)))//" "
                else
                    report=report//trim(V(therapeutic_bullet_set(i1)%targ(i2)))//"["//real2char(therapeutic_bullet_set(i1)%moda(i2))//"] "
                end if
            end do
            report=report//"("//real2char(therapeutic_bullet_set(i1)%gain(1))//"% --> "//real2char(therapeutic_bullet_set(i1)%gain(2))//"%)"//new_line("a")//repeat("-",80)//new_line("a")
        end do
        report=report//"found therapeutic bullets: "//int2char(size(therapeutic_bullet_set))
        write (unit=*,fmt="(a)",advance="no") new_line("a")//report//new_line("a")//new_line("a")//"save [1/0] "
        read (unit=*,fmt=*) save_
        if (save_==1) then
            open (unit=1,file="report_therapeutic_bullet.txt",status="replace")
            write (unit=1,fmt="(a)") report
            close (unit=1)
            write (unit=*,fmt="(a)") new_line("a")//"report saved as: report_therapeutic_bullet.txt"
        end if
        deallocate(report)
    end subroutine report_therapeutic_bullet_set
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
            z=y(i)
            y(i)=y(i_min)
            y(i_min)=z
        end do
    end function sort
    !##########################################################################!
    !############################    what_to_do    ############################!
    !##########################################################################!
    subroutine what_to_do(f,value,size_D,n_node,max_targ,max_moda,V)
        logical::boolean
        integer::to_do,r_min,r_max,setting,comprehensive_D,size_D,n_node,max_targ,max_moda
        integer,dimension(0)::dummy1
        real::start,finish
        real,dimension(0)::dummy2
        real,dimension(:)::value
        real,dimension(:,:),allocatable::D
        character(16),dimension(:)::V
        type(attractor),dimension(:),allocatable::A_set,A_physio,A_patho,a_patho_set
        type(bullet),dimension(:),allocatable::therapeutic_bullet_set
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        call init_random_seed()
        call cpu_time(start)
        boolean=all(value==[0.0,1.0])
        write (unit=*,fmt="(a)",advance="no") new_line("a")//"[1] compute attractors"//new_line("a")//"[2] compute pathological attractors"//new_line("a")//"[3] compute therapeutic bullets"//new_line("a")//"[4] help"//new_line("a")//"[5] license"//new_line("a")//new_line("a")//"what to do [1/2/3/4/5] "
        read (unit=*,fmt=*) to_do
        if (to_do==1 .or. to_do==3) then
            if (boolean) then
                write (unit=*,fmt="(a,es10.3e3,a)",advance="no") new_line("a")//"state space cardinality: ",real(2,8)**real(n_node,8),", comprehensive computation [1/0] "
                read (unit=*,fmt=*) comprehensive_D
                select case (comprehensive_D)
                    case (1)
                        D=generate_state_space(n_node)
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
                write (unit=*,fmt="(a)",advance="no") new_line("a")//"[1] physiological"//new_line("a")//"[2] pathological"//new_line("a")//new_line("a")//"setting [1/2] "
                read (unit=*,fmt=*) setting
                call report_attractor_set(A_set,setting,V,boolean)
                deallocate(A_set,D)
            case (2)
                A_physio=load_attractor_set(1)
                A_patho=load_attractor_set(2)
                a_patho_set=compute_pathological_attractor(A_physio,A_patho)
                call report_attractor_set(a_patho_set,3,V,boolean)
                deallocate(A_physio,A_patho,a_patho_set)
            case (3)
                A_physio=load_attractor_set(1)
                A_patho=load_attractor_set(2)
                write (unit=*,fmt="(a)",advance="no") new_line("a")//"number of targets per bullet (lower bound): "
                read (unit=*,fmt=*) r_min
                write (unit=*,fmt="(a)",advance="no") "number of targets per bullet (upper bound): "
                read (unit=*,fmt=*) r_max
                therapeutic_bullet_set=compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio,A_patho)
                call report_therapeutic_bullet_set(therapeutic_bullet_set,V,boolean)
                deallocate(A_physio,therapeutic_bullet_set,D)
            case (4)
                write (unit=*,fmt="(a)") new_line("a")//"1) do step 1 with f_physio"//new_line("a")//"2) do step 1 with f_patho"//new_line("a")//"3) eventually do step 2"//new_line("a")//"4) do step 3 with f_patho"//new_line("a")//new_line("a")//"do not forget to recompile the sources following any modification"
            case (5)
                write (unit=*,fmt="(a)") new_line("a")//"kali-targ: a tool for in silico target identification."//new_line("a")//"Copyright (C) 2013-2014 Arnaud Poret"//new_line("a")//new_line("a")//"This program is free software: you can redistribute it and/or modify it under"//new_line("a")//"the terms of the GNU General Public License as published by the Free Software"//new_line("a")//"Foundation, either version 3 of the License, or (at your option) any later"//new_line("a")//"version."//new_line("a")//new_line("a")//"This program is distributed in the hope that it will be useful, but WITHOUT ANY"//new_line("a")//"WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A"//new_line("a")//"PARTICULAR PURPOSE. See the GNU General Public License for more details."//new_line("a")//new_line("a")//"You should have received a copy of the GNU General Public License along with"//new_line("a")//"this program. If not, see https://www.gnu.org/licenses/gpl.html."
        end select
        call cpu_time(finish)
        write (unit=*,fmt="(a)") new_line("a")//"done in "//int2char(int(finish-start))//" seconds"//new_line("a")
    end subroutine what_to_do
end module lib
