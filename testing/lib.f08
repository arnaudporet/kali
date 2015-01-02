! Copyright (C) 2013-2014 Arnaud Poret
! This program is licensed under the GNU General Public License.
! To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.
module lib
    !############    /!\ the network must be deterministic /!\    #############!
    !#######    /!\ the default integer kind must be 4 (32 bits) /!\    #######!
    !########    /!\ the default real kind must be 4 (32 bits) /!\    #########!
    integer::max_targ,max_moda,size_D,n_node
    real::min_gain
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
        real,dimension(:,:),allocatable::gain
    end type bullet
    contains
    !##########################################################################!
    !############################    compare_a    #############################!
    !##########################################################################!
    function compare_a(a1,a2) result(diff)
        !########    /!\ the attractors must be in sorted form /!\    #########!
        logical::diff
        type(attractor)::a1,a2
        if (size(a1%mat,2)/=size(a2%mat,2)) then
            diff=.true.
        else
            diff=any(a1%mat/=a2%mat)
        end if
    end function compare_a
    !##########################################################################!
    !############################    compute_a    #############################!
    !##########################################################################!
    function compute_a(f,x0,b) result(a)
        !###########    /!\ the network must be deterministic /!\    ###########!
        integer::k,i
        real,dimension(:,:)::x0
        real,dimension(:,:),allocatable::x
        type(attractor)::a
        type(bullet)::b
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        x=x0
        k=1
        do1:do
            x=concat(x,f(x,k),2)
            x(b%targ,k+1)=b%moda
            do i=k,1,-1
                if (all(x(:,i)==x(:,k+1))) then
                    a%mat=reshape(x(:,i:k),[size(x,1),k-i+1])
                    a=sort_a(a)
                    deallocate(x)
                    exit do1
                end if
            end do
            k=k+1
        end do do1
    end function compute_a
    !##########################################################################!
    !##########################    compute_A_set    ###########################!
    !##########################################################################!
    function compute_A_set(f,D,setting,ref_set,b) result(A_set)
        integer::i,in_A,setting
        real,dimension(:,:)::D
        type(attractor)::a
        type(attractor),dimension(:)::ref_set
        type(attractor),dimension(:),allocatable::A_set
        type(bullet)::b
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        allocate(A_set(0))
        do i=1,size(D,2)
            a=compute_a(f,reshape(D(:,i),[size(D,1),1]),b)
            in_A=in_A_set(a,A_set)
            if (in_A>0) then
                A_set(in_A)%basin=A_set(in_A)%basin+1.0
            else
                a%basin=1.0
                A_set=[A_set,a]
            end if
        end do
        deallocate(a%mat)
        do i=1,size(A_set)
            A_set(i)%basin=A_set(i)%basin*100.0/real(size(D,2))
        end do
        A_set=sort_A_set(A_set)
        A_set=name_A_set(A_set,ref_set,setting)
    end function compute_A_set
    !##########################################################################!
    !#######################    compute_a_patho_set    ########################!
    !##########################################################################!
    function compute_a_patho_set(A_patho) result(a_patho_set)
        !#############    /!\ the attractors must be named /!\    #############!
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
    !#########################    compute_B_therap    #########################! TODO modularize?
    !##########################################################################!
    function compute_B_therap(f,D,r_min,r_max,max_targ,max_moda,min_gain,de_novo,n_node,value,A_physio,A_patho,a_patho_set) result(B_therap)
        logical::allowed
        integer::i1,i2,i3,i4,r_min,r_max,max_targ,max_moda,de_novo,n_node
        integer,dimension(:,:),allocatable::C_targ
        real::min_gain
        real,dimension(:)::value
        real,dimension(:,:)::D
        real,dimension(:,:),allocatable::C_moda
        type(attractor),dimension(:)::A_physio,A_patho,a_patho_set
        type(attractor),dimension(:),allocatable::A_test
        type(bullet)::b
        type(bullet),dimension(:),allocatable::B_therap
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        allocate(B_therap(0))
        allocate(b%gain(size(A_physio)+size(a_patho_set),2))
        b%gain(:size(A_physio),1)=compute_cover(A_physio,A_patho)
        b%gain(size(A_physio)+1:,1)=compute_cover(a_patho_set,A_patho)
        do i1=r_min,min(r_max,n_node)
            C_targ=int(gen_combi(real(range_int(1,n_node)),i1,max_targ))
            C_moda=gen_arrang(value,i1,max_moda)
            do i2=1,size(C_targ,1)
                do i3=1,size(C_moda,1)
                    b%targ=C_targ(i2,:)
                    b%moda=C_moda(i3,:)
                    A_test=compute_A_set(f,D,2,A_physio,b)
                    b%gain(:size(A_physio),2)=compute_cover(A_physio,A_test)
                    b%gain(size(A_physio)+1:,2)=compute_cover(a_patho_set,A_test)
                    if (sum(b%gain(:size(A_physio),2))>sum(b%gain(:size(A_physio),1))+min_gain) then
                        allowed=.true.
                        if (de_novo==0) then
                            do i4=1,size(A_test)
                                if (index(trim(A_test(i4)%name),"patho")/=0) then
                                    if (in_A_set(A_test(i4),a_patho_set)==0) then
                                        allowed=.false.
                                        exit
                                    end if
                                end if
                            end do
                        end if
                        if (allowed) then
                            B_therap=[B_therap,b]
                        end if
                    end if
                end do
            end do
        end do
        deallocate(C_targ,C_moda,A_test,b%targ,b%moda,b%gain)
        B_therap=sort_B_therap(B_therap)
    end function compute_B_therap
    !##########################################################################!
    !##########################    compute_cover    ###########################!
    !##########################################################################!
    function compute_cover(A_set1,A_set2) result(cover)
        integer::i,in_2
        type(attractor),dimension(:)::A_set1,A_set2
        real,dimension(size(A_set1))::cover
        do i=1,size(A_set1)
            in_2=in_A_set(A_set1(i),A_set2)
            if (in_2>0) then
                cover(i)=A_set2(in_2)%basin
            else
                cover(i)=0.0
            end if
        end do
    end function compute_cover
    !##########################################################################!
    !##############################    concat    ##############################!
    !##########################################################################!
    function concat(x1,x2,d) result(y)
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
                    y(:size(x1,1),:)=x1
                    y(size(x1,1)+1:,:)=x2
                end if
            case (2)
                if (size(x1,2)==0) then
                    y=x2
                else if (size(x2,2)==0) then
                    y=x1
                else
                    allocate(y(size(x1,1),size(x1,2)+size(x2,2)))
                    y(:,:size(x1,2))=x1
                    y(:,size(x1,2)+1:)=x2
                end if
        end select
    end function concat
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
                y=y*real(x-i,8)
            end do
        end if
    end function facto
    !##########################################################################!
    !############################    gen_arrang    ############################!
    !##########################################################################!
    function gen_arrang(deck,k,n_arrang) result(arrang_mat)
        !#################    /!\ only with repetition /!\    #################!
        logical::in_arrang_mat
        integer::k,n_arrang,i1,i2
        real,dimension(k)::arrang
        real,dimension(:)::deck
        real,dimension(:,:),allocatable::arrang_mat
        allocate(arrang_mat(min(n_arrang,int(min(real(size(deck),8)**real(k,8),real(huge(1),8)))),k))
        do i1=1,size(arrang_mat,1)
            do
                do i2=1,k
                    arrang(i2)=deck(rand_int(1,size(deck)))
                end do
                in_arrang_mat=.false.
                do i2=1,i1-1
                    if (all(arrang_mat(i2,:)==arrang)) then
                        in_arrang_mat=.true.
                        exit
                    end if
                end do
                if (.not. in_arrang_mat) then
                    arrang_mat(i1,:)=arrang
                    exit
                end if
            end do
        end do
    end function gen_arrang
    !##########################################################################!
    !############################    gen_combi    #############################!
    !##########################################################################!
    function gen_combi(deck,k,n_combi) result(combi_mat)
        !###############    /!\ only without repetition /!\    ################!
        logical::in_combi_mat
        integer::k,n_combi,i1,i2
        real::z
        real,dimension(k)::combi
        real,dimension(:)::deck
        real,dimension(:,:),allocatable::combi_mat
        allocate(combi_mat(min(n_combi,int(min(facto(size(deck))/(facto(k)*facto(size(deck)-k)),real(huge(1),8)))),k))
        do i1=1,size(combi_mat,1)
            do
                do i2=1,k
                    do
                        z=deck(rand_int(1,size(deck)))
                        if (all(z/=combi(:i2-1))) then
                            combi(i2)=z
                            exit
                        end if
                    end do
                end do
                combi=sort(combi)
                in_combi_mat=.false.
                do i2=1,i1-1
                    if (all(combi_mat(i2,:)==combi)) then
                        in_combi_mat=.true.
                        exit
                    end if
                end do
                if (.not. in_combi_mat) then
                    combi_mat(i1,:)=combi
                    exit
                end if
            end do
        end do
    end function gen_combi
    !##########################################################################!
    !##############################    gen_S    ###############################!
    !##########################################################################!
    function gen_S(n,value) result(S)
        integer::i1,i2,n
        real,dimension(:)::value
        real,dimension(:,:),allocatable::S
        if (real(size(value),8)**real(n,8)>real(huge(1),8)) then
            write (unit=*,fmt="(a)") "gen_S(n,value): S too big to be generated."//new_line("a")
            stop
        else
            allocate(S(n,size(value)**n))
            do i1=1,n
                do i2=1,size(value)-1
                    S(:size(value)-1,i2*size(value)**(i1-1)+1:(i2+1)*size(value)**(i1-1))=S(:i1-1,:size(value)**(i1-1))
                end do
                do i2=1,size(value)
                    S(i1,(i2-1)*size(value)**(i1-1)+1:i2*size(value)**(i1-1))=value(i2)
                end do
            end do
        end if
    end function gen_S
    !##########################################################################!
    !#############################    in_A_set    #############################!
    !##########################################################################!
    function in_A_set(a,A_set) result(in_A)
        integer::in_A,i
        type(attractor)::a
        type(attractor),dimension(:)::A_set
        in_A=0
        do i=1,size(A_set)
            if (.not. compare_a(a,A_set(i))) then
                in_A=i
                exit
            end if
        end do
    end function in_A_set
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
        !############    /!\ x must be of kind 4 (32 bits) /!\    #############!
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
        integer::setting,i1,i2,size1
        integer,dimension(2)::size2
        character(16)::set_name
        type(attractor),dimension(:),allocatable::A_set
        select case (setting)
            case (1)
                set_name="A_physio.csv"
            case (2)
                set_name="A_patho.csv"
            case (3)
                set_name="A_versus.csv"
        end select
        open (unit=1,file=trim(set_name),status="old")
        read (unit=1,fmt=*) size1
        allocate(A_set(size1))
        do i1=1,size(A_set)
            read (unit=1,fmt=*) size2
            read (unit=1,fmt=*) A_set(i1)%basin
            read (unit=1,fmt=*) A_set(i1)%name
            allocate(A_set(i1)%mat(size2(1),size2(2)))
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
    !############################    name_A_set    ############################!
    !##########################################################################!
    function name_A_set(A_set,ref_set,setting) result(y)
        integer::setting,k,i,in_ref
        character(16)::a_name
        type(attractor),dimension(:)::A_set,ref_set
        type(attractor),dimension(size(A_set))::y
        y=A_set
        select case (setting)
            case (1)
                a_name="a_physio"
            case (2)
                a_name="a_patho"
        end select
        k=1
        do i=1,size(y)
            in_ref=in_A_set(y(i),ref_set)
            if (in_ref>0) then
                y(i)%name=trim(ref_set(in_ref)%name)
            else
                y(i)%name=trim(a_name)//int2char(k)
                k=k+1
            end if
        end do
    end function name_A_set
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
    function real2char(x,dec) result(y)
        !############    /!\ x must be of kind 4 (32 bits) /!\    #############!
        integer::dec
        real::x
        character(41+dec)::z
        character(:),allocatable::y
        write (unit=z,fmt="(f"//int2char(41+dec)//"."//int2char(dec)//")") x
        y=trim(adjustl(z))
    end function real2char
    !##########################################################################!
    !###########################    report_A_set    ###########################!
    !##########################################################################!
    subroutine report_A_set(A_set,setting,V,bool,dec)
        logical::bool
        integer::n_point,n_cycle,setting,i1,i2,i3,save_,dec
        character(16)::set_type,space_type,report_name
        character(16),dimension(:)::V
        character(:),allocatable::report
        type(attractor),dimension(:)::A_set
        n_point=0
        n_cycle=0
        select case (setting)
            case (1)
                set_type="A_physio"
                space_type="physiological"
            case (2)
                set_type="A_patho"
                space_type="pathological"
            case (3)
                set_type="A_versus"
                space_type="pathological"
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
            report=report//"Name: "//trim(A_set(i1)%name)//new_line("a")//"Basin: "//real2char(A_set(i1)%basin,1)//"% (of the "//trim(space_type)//" state space)"//new_line("a")//"Matrix:"//new_line("a")
            do i2=1,size(A_set(i1)%mat,1)
                report=report//"    "//V(i2)//" "
                do i3=1,size(A_set(i1)%mat,2)-1
                    if (bool) then
                        report=report//int2char(int(A_set(i1)%mat(i2,i3)))//" "
                    else
                        report=report//real2char(A_set(i1)%mat(i2,i3),dec)//" "
                    end if
                end do
                if (bool) then
                    report=report//int2char(int(A_set(i1)%mat(i2,size(A_set(i1)%mat,2))))//new_line("a")
                else
                    report=report//real2char(A_set(i1)%mat(i2,size(A_set(i1)%mat,2)),dec)//new_line("a")
                end if
            end do
            report=report//repeat("-",80)//new_line("a")
        end do
        report=report//"Found attractors: "//int2char(size(A_set))//new_line("a")//"    point(s): "//int2char(n_point)//new_line("a")//"    cycle(s): "//int2char(n_cycle)
        write (unit=*,fmt="(a)") new_line("a")//report//new_line("a")//new_line("a")//"Save?"//new_line("a")//"    [0] no"//new_line("a")//"    [1] yes"
        read (unit=*,fmt=*) save_
        if (save_==1) then
            write (unit=*,fmt="(a)") ""
            call save_A_set(A_set,setting,dec)
            select case (setting)
                case (1)
                    report_name="A_physio.txt"
                case (2)
                    report_name="A_patho.txt"
                case (3)
                    report_name="A_versus.txt"
            end select
            open (unit=1,file=trim(report_name),status="replace")
            write (unit=1,fmt="(a)") report
            close (unit=1)
            deallocate(report)
            write (unit=*,fmt="(a)") "Report saved as "//trim(report_name)//"."
        end if
    end subroutine report_A_set
    !##########################################################################!
    !#########################    report_B_therap    ##########################!
    !##########################################################################!
    subroutine report_B_therap(B_therap,V,bool,dec,A_physio,a_patho_set,min_gain,r_min,r_max)
        logical::bool
        integer::i1,i2,save_,dec,r_min,r_max
        integer,dimension(r_max-r_min+1)::barrel
        real::min_gain
        character(1)::moda
        character(16),dimension(:)::V
        character(:),allocatable::report
        type(attractor),dimension(:)::A_physio,a_patho_set
        type(bullet),dimension(:)::B_therap
        report="Minimum gain: "//real2char(min_gain,1)//"%"//new_line("a")//"Number of targets per bullet: "//int2char(r_min)//"-"//int2char(r_max)//new_line("a")//repeat("-",80)//new_line("a")
        barrel=0
        do i1=1,size(B_therap)
            barrel(size(B_therap(i1)%targ)-r_min+1)=barrel(size(B_therap(i1)%targ)-r_min+1)+1
            report=report//"Bullet: "
            do i2=1,size(B_therap(i1)%targ)
                if (bool) then
                    if (B_therap(i1)%moda(i2)==1.0) then
                        moda="+"
                    else
                        moda="-"
                    end if
                    report=report//moda//trim(V(B_therap(i1)%targ(i2)))//" "
                else
                    report=report//trim(V(B_therap(i1)%targ(i2)))//"["//real2char(B_therap(i1)%moda(i2),dec)//"] "
                end if
            end do
            report=report//new_line("a")//"Union of the physiological basins: "//real2char(sum(B_therap(i1)%gain(:size(A_physio),1)),1)//"% --> "//real2char(sum(B_therap(i1)%gain(:size(A_physio),2)),1)//"%"//new_line("a")//"Physiological basins:"//new_line("a")
            do i2=1,size(A_physio)
                report=report//"    "//trim(A_physio(i2)%name)//": "//real2char(B_therap(i1)%gain(i2,1),1)//"% --> "//real2char(B_therap(i1)%gain(i2,2),1)//"%"//new_line("a")
            end do
            report=report//"Pathological basins:"//new_line("a")
            do i2=1,size(a_patho_set)
                report=report//"    "//trim(a_patho_set(i2)%name)//": "//real2char(B_therap(i1)%gain(i2+size(A_physio),1),1)//"% --> "//real2char(B_therap(i1)%gain(i2+size(A_physio),2),1)//"%"//new_line("a")
            end do
            report=report//repeat("-",80)//new_line("a")
        end do
        report=report//"Found therapeutic bullets: "//int2char(size(B_therap))//new_line("a")
        do i1=1,size(barrel)-1
            report=report//"    "//int2char(i1+r_min-1)//"-bullet(s): "//int2char(barrel(i1)+r_min-1)//new_line("a")
        end do
        report=report//"    "//int2char(size(barrel)+r_min-1)//"-bullet(s): "//int2char(barrel(size(barrel))+r_min-1)
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
    !############################    save_A_set    ############################!
    !##########################################################################!
    subroutine save_A_set(A_set,setting,dec)
        integer::setting,i1,i2,i3,dec
        character(16)::set_name
        character(:),allocatable::s
        type(attractor),dimension(:)::A_set
        select case (setting)
            case (1)
                set_name="A_physio.csv"
            case (2)
                set_name="A_patho.csv"
            case (3)
                set_name="A_versus.csv"
        end select
        s=int2char(size(A_set))//new_line("a")
        do i1=1,size(A_set)
            s=s//int2char(size(A_set(i1)%mat,1))//","//int2char(size(A_set(i1)%mat,2))//new_line("a")//real2char(A_set(i1)%basin,1)//new_line("a")//trim(A_set(i1)%name)//new_line("a")
        end do
        do i1=1,size(A_set)
            do i2=1,size(A_set(i1)%mat,1)
                do i3=1,size(A_set(i1)%mat,2)-1
                    s=s//real2char(A_set(i1)%mat(i2,i3),dec)//","
                end do
                s=s//real2char(A_set(i1)%mat(i2,size(A_set(i1)%mat,2)),dec)
                if (i1/=size(A_set) .or. i2/=size(A_set(i1)%mat,1)) then
                    s=s//new_line("a")
                end if
            end do
        end do
        open (unit=1,file=trim(set_name),status="replace")
        write (unit=1,fmt="(a)") s
        close (unit=1)
        deallocate(s)
        write (unit=*,fmt="(a)") "Set saved as "//trim(set_name)//"."
    end subroutine save_A_set
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
    !##############################    sort_a    ##############################!
    !##########################################################################!
    function sort_a(a) result(y)
        !###########    /!\ the network must be deterministic /!\    ###########!
        integer::i
        integer,dimension(:),allocatable::j_min
        type(attractor)::a,y
        y=a
        j_min=range_int(1,size(y%mat,2))
        do i=1,size(y%mat,1)
            j_min=j_min(minlocs(y%mat(i,j_min)))
            if (size(j_min)==1) then
                y%mat=cshift(y%mat,j_min(1)-1,2)
                deallocate(j_min)
                exit
            end if
        end do
    end function sort_a
    !##########################################################################!
    !############################    sort_A_set    ############################!
    !##########################################################################!
    function sort_A_set(A_set) result(y)
        logical::repass
        integer::i1,i2
        type(attractor)::a
        type(attractor),dimension(:)::A_set
        type(attractor),dimension(size(A_set))::y
        y=A_set
        allocate(a%mat(0,0))
        do
            repass=.false.
            do i1=1,size(y)-1
                if (size(y(i1)%mat,2)>size(y(i1+1)%mat,2)) then
                    repass=.true.
                    a=y(i1)
                    y(i1)=y(i1+1)
                    y(i1+1)=a
                else if (size(y(i1)%mat,2)==size(y(i1+1)%mat,2)) then
                    do i2=1,size(y(i1)%mat,1)
                        if (y(i1)%mat(i2,1)<y(i1+1)%mat(i2,1)) then
                            exit
                        else if (y(i1)%mat(i2,1)>y(i1+1)%mat(i2,1)) then
                            repass=.true.
                            a=y(i1)
                            y(i1)=y(i1+1)
                            y(i1+1)=a
                            exit
                        end if
                    end do
                end if
            end do
            if (.not. repass) then
                deallocate(a%mat)
                exit
            end if
        end do
    end function sort_A_set
    !##########################################################################!
    !##########################    sort_B_therap    ###########################!
    !##########################################################################!
    function sort_B_therap(B_therap) result(y)
        logical::repass
        integer::i1,i2
        type(bullet)::b
        type(bullet),dimension(:)::B_therap
        type(bullet),dimension(size(B_therap))::y
        y=B_therap
        allocate(b%targ(0))
        allocate(b%moda(0))
        allocate(b%gain(0,0))
        do
            repass=.false.
            do i1=1,size(y)-1
                if (size(y(i1)%targ)>size(y(i1+1)%targ)) then
                    repass=.true.
                    b=y(i1)
                    y(i1)=y(i1+1)
                    y(i1+1)=b
                else if (size(y(i1)%targ)==size(y(i1+1)%targ)) then
                    if (any(y(i1)%targ/=y(i1+1)%targ)) then
                        do i2=1,size(y(i1)%targ)
                            if (y(i1)%targ(i2)<y(i1+1)%targ(i2)) then
                                exit
                            else if (y(i1)%targ(i2)>y(i1+1)%targ(i2)) then
                                repass=.true.
                                b=y(i1)
                                y(i1)=y(i1+1)
                                y(i1+1)=b
                                exit
                            end if
                        end do
                    else
                        do i2=1,size(y(i1)%moda)
                            if (y(i1)%moda(i2)<y(i1+1)%moda(i2)) then
                                exit
                            else if (y(i1)%moda(i2)>y(i1+1)%moda(i2)) then
                                repass=.true.
                                b=y(i1)
                                y(i1)=y(i1+1)
                                y(i1+1)=b
                                exit
                            end if
                        end do
                    end if
                end if
            end do
            if (.not. repass) then
                deallocate(b%targ,b%moda,b%gain)
                exit
            end if
        end do
    end function sort_B_therap
    !##########################################################################!
    !############################    what_to_do    ############################!
    !##########################################################################!
    subroutine what_to_do(f1,f2,value,size_D,n_node,max_targ,max_moda,min_gain,V)
        logical::bool,exist1,exist2,exist3
        integer::size_D,n_node,max_targ,max_moda,to_do,r_min,r_max,setting,whole_S,de_novo,dec,i1,i2
        real::min_gain
        real,dimension(:)::value
        integer,dimension(size(value))::z
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
                real,dimension(size(x,1),1)::y
            end function f1
        end interface
        interface
            function f2(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f2
        end interface
        call init_random_seed()
        if (size(value)==2) then
            bool=all(value==[0.0,1.0])
        else
            bool=.false.
        end if
        z=3
        do i1=1,size(value)
            do i2=1,2
                if (modulo(value(i1)*real(10**i2),1.0)==0.0) then
                    z(i1)=i2
                    exit
                end if
            end do
        end do
        dec=maxval(z)
        do
            write (unit=*,fmt="(a)") new_line("a")//"What to do?"//new_line("a")//"    [1] compute attractor set"//new_line("a")//"    [2] compute pathological attractors"//new_line("a")//"    [3] compute therapeutic bullets"//new_line("a")//"    [4] help"//new_line("a")//"    [5] license"//new_line("a")//"    [6] quit"
            read (unit=*,fmt=*) to_do
            if (to_do==1 .or. to_do==3) then
                write (unit=*,fmt="(a,es10.3e3,a)") new_line("a")//"State space cardinality: ",real(size(value),8)**real(n_node,8),", compute the whole state space?"//new_line("a")//"    [0] no"//new_line("a")//"    [1] yes"
                read (unit=*,fmt=*) whole_S
                select case (whole_S)
                    case (1)
                        D=gen_S(n_node,value)
                    case (0)
                        D=transpose(gen_arrang(value,n_node,size_D))
                end select
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
                            call report_A_set(A_physio,1,V,bool,dec)
                            deallocate(A_physio)
                        case (2)
                            inquire (file="A_physio.csv",exist=exist1)
                            if (.not. exist1) then
                                write (unit=*,fmt="(a)") new_line("a")//"The file A_physio.csv is required for computing the pathological attractor set."//new_line("a")//"Ensure that the physiological attractor set is already computed."
                            else
                                A_physio=load_A_set(1)
                                A_patho=compute_A_set(f2,D,2,A_physio,null_b)
                                call report_A_set(A_patho,2,V,bool,dec)
                                deallocate(A_physio,A_patho)
                            end if
                    end select
                    deallocate(D,null_b%targ,null_b%moda)
                case (2)
                    inquire (file="A_patho.csv",exist=exist1)
                    if (.not. exist1) then
                        write (unit=*,fmt="(a)") new_line("a")//"The file A_patho.csv is required for computing the pathological attractors."//new_line("a")//"Ensure that the pathological attractor set is already computed."
                    else
                        A_patho=load_A_set(2)
                        a_patho_set=compute_a_patho_set(A_patho)
                        call report_A_set(a_patho_set,3,V,bool,dec)
                        deallocate(A_patho,a_patho_set)
                    end if
                case (3)
                    inquire (file="A_physio.csv",exist=exist1)
                    inquire (file="A_patho.csv",exist=exist2)
                    inquire (file="A_versus.csv",exist=exist3)
                    if (.not. (exist1 .and. exist2 .and. exist3)) then
                        write (unit=*,fmt="(a)") new_line("a")//"The files A_physio.csv, A_patho.csv and A_versus.csv are required for computing"//new_line("a")//"the therapeutic bullets. Ensure that the physiological attractor set, the"//new_line("a")//"pathological attractor set and the pathological attractors are already computed."
                    else
                        A_physio=load_A_set(1)
                        A_patho=load_A_set(2)
                        a_patho_set=load_A_set(3)
                        !write (unit=*,fmt="(a)") new_line("a")//"Allow de novo attractors?"//new_line("a")//"    [0] no"//new_line("a")//"    [1] yes"
                        de_novo=0!read (unit=*,fmt=*) de_novo<<<<<<<<<<<<<<<<<<<<<<<<<
                        write (unit=*,fmt="(a)",advance="no") new_line("a")//"Number of targets per bullet (lower bound): "
                        read (unit=*,fmt=*) r_min
                        write (unit=*,fmt="(a)",advance="no") "Number of targets per bullet (upper bound): "
                        read (unit=*,fmt=*) r_max
                        B_therap=compute_B_therap(f2,D,r_min,r_max,max_targ,max_moda,min_gain,de_novo,n_node,value,A_physio,A_patho,a_patho_set)
                        call report_B_therap(B_therap,V,bool,dec,A_physio,a_patho_set,min_gain,r_min,r_max)
                        deallocate(A_physio,A_patho,a_patho_set,B_therap)
                    end if
                    deallocate(D)
                case (4)
                    write (unit=*,fmt="(a)") new_line("a")//"How to:"//new_line("a")//"    1) compute the physiological attractor set: [1]"//new_line("a")//"        * when prompted by the algorithm, set the setting to physiological"//new_line("a")//"        * returns A_physio"//new_line("a")//"        * when prompted by the algorithm, saving A_physio is necessary for the"//new_line("a")//"          next steps"//new_line("a")//"    2) compute the pathological attractor set: [1]"//new_line("a")//"        * when prompted by the algorithm, set the setting to pathological"//new_line("a")//"        * returns A_patho"//new_line("a")//"        * when prompted by the algorithm, saving A_patho is necessary for the"//new_line("a")//"          next steps"//new_line("a")//"    3) compute the pathological attractors: [2]"//new_line("a")//"        * returns A_versus"//new_line("a")//"        * when prompted by the algorithm, saving A_versus is necessary for the"//new_line("a")//"          next steps"//new_line("a")//"    4) compute the therapeutic bullets: [3]"//new_line("a")//"        * returns B_therap"//new_line("a")//"        * in case of multivalued logic, therapeutic bullets are reported as"//new_line("a")//"          follow: ... X[y] ... where the variable X has to be set to the value y"//new_line("a")//new_line("a")//"If you rename/move/delete the csv files generated by the algorithm, it will not"//new_line("a")//"recognize them when requiered, if any."//new_line("a")//new_line("a")//"Do not forget to recompile the sources following any modification."
                case (5)
                    write (unit=*,fmt="(a)") new_line("a")//"kali-targ: a tool for in silico target identification."//new_line("a")//"Copyright (C) 2013-2014 Arnaud Poret"//new_line("a")//new_line("a")//"This program is free software: you can redistribute it and/or modify it under"//new_line("a")//"the terms of the GNU General Public License as published by the Free Software"//new_line("a")//"Foundation, either version 3 of the License, or (at your option) any later"//new_line("a")//"version."//new_line("a")//new_line("a")//"This program is distributed in the hope that it will be useful, but WITHOUT ANY"//new_line("a")//"WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A"//new_line("a")//"PARTICULAR PURPOSE. See the GNU General Public License for more details."//new_line("a")//new_line("a")//"You should have received a copy of the GNU General Public License along with"//new_line("a")//"this program. If not, see https://www.gnu.org/licenses/gpl.html."
                case (6)
                    write (unit=*,fmt="(a)") new_line("a")//"Goodbye."//new_line("a")
                    exit
            end select
        end do
    end subroutine what_to_do
end module lib
