
module lib
    !##############    /!\ networks must be deterministic /!\    ##############!
    integer::max_targ,max_moda,size_D,n_node
    real,dimension(:),allocatable::value
    character(16),dimension(:),allocatable::V
    type::attractor
        real::popularity
        real,dimension(:,:),allocatable::att
        character(16)::name
    end type attractor
    type::bullet
        integer,dimension(:),allocatable::targ
        real,dimension(:),allocatable::moda
        character(16)::metal
        character(16),dimension(:),allocatable::unrecovered
    end type bullet
    contains
    !##########################################################################!
    !########################    compare_attractor    #########################!
    !##########################################################################!
    function compare_attractor(a1,a2) result(differ)
        !##########    /!\ attractors must be in sorted form /!\    ###########!
        logical::differ
        type(attractor)::a1,a2
        if (size(a1%att,2)/=size(a2%att,2)) then
            differ=.true.
        else
            differ=.not. all(a1%att==a2%att)
        end if
    end function compare_attractor
    !##########################################################################!
    !######################    compare_attractor_set    #######################!
    !##########################################################################!
    function compare_attractor_set(A_set1,A_set2) result(differ)
        !########    /!\ attractor sets must be in sorted form /!\    #########!
        !##########    /!\ attractors must be in sorted form /!\    ###########!
        logical::differ
        integer::i
        type(attractor),dimension(:)::A_set1,A_set2
        if (size(A_set1)/=size(A_set2)) then
            differ=.true.
        else
            differ=.false.
            do i=1,size(A_set1)
                if (compare_attractor(A_set1(i),A_set2(i))) then
                    differ=.true.
                    exit
                end if
            end do
        end if
    end function compare_attractor_set
    !##########################################################################!
    !########################    compute_attractor    #########################!
    !##########################################################################!
    function compute_attractor(f,D,setting,A_physio,bull) result(A_set)
        logical::a_found,in_A,in_physio
        integer::i1,i2,k,setting,i_patho
        integer,dimension(:),allocatable::c_targ
        real,dimension(:),allocatable::c_moda
        real,dimension(:,:)::D
        real,dimension(:,:),allocatable::x
        character(16)::attractor_name
        type(attractor)::a
        type(attractor),dimension(:),optional::A_physio
        type(attractor),dimension(:),allocatable::A_set,A_check
        type(bullet),optional::bull
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        allocate(A_set(0))
        if (present(bull)) then
            c_targ=bull%targ
            c_moda=bull%moda
        else
            allocate(c_targ(0))
            allocate(c_moda(0))
        end if
        do i1=1,size(D,2)
            a_found=.false.
            x=reshape(D(:,i1),[size(D,1),1])
            k=1
            do
                x=concatenate(x,f(x,k),2)
                x(c_targ,k+1)=c_moda
                do i2=k,1,-1
                    if (all(x(:,i2)==x(:,k+1))) then
                        a_found=.true.
                        a%att=reshape(x(:,i2:k),[size(D,1),k-i2+1])
                        a=sort_attractor(a)
                        exit
                    end if
                end do
                if (a_found) then
                    in_A=.false.
                    do i2=1,size(A_set)
                        if (.not. compare_attractor(A_set(i2),a)) then
                            in_A=.true.
                            A_set(i2)%popularity=A_set(i2)%popularity+1.0
                            exit
                        end if
                    end do
                    if (.not. in_A) then
                        a%popularity=1.0
                        A_set=[A_set,a]
                    end if
                    exit
                else
                    k=k+1
                end if
            end do
        end do
        deallocate(x,a%att,c_targ,c_moda)
        A_set=sort_attractor_set(A_set)
        if (present(A_physio)) then
            A_check=A_physio
        else
            allocate(A_check(0))
        end if
        select case (setting)
            case (1)
                attractor_name="a_physio"
            case (2)
                attractor_name="a_patho"
        end select
        i_patho=1
        do i1=1,size(A_set)
            A_set(i1)%popularity=A_set(i1)%popularity*100.0/real(size(D,2))
            in_physio=.false.
            do i2=1,size(A_check)
                if (.not. compare_attractor(A_set(i1),A_check(i2))) then
                    in_physio=.true.
                    exit
                end if
            end do
            if (in_physio) then
                A_set(i1)%name=trim(A_check(i2)%name)
            else
                A_set(i1)%name=trim(attractor_name)//int2char(i_patho)
                i_patho=i_patho+1
            end if
        end do
        deallocate(A_check)
    end function compute_attractor
    !##########################################################################!
    !##################    compute_pathological_attractor    ##################!
    !##########################################################################!
    function compute_pathological_attractor(A_patho) result(a_patho_set)
        integer::i
        type(attractor),dimension(:)::A_patho
        type(attractor),dimension(:),allocatable::a_patho_set
        allocate(a_patho_set(0))
        do i=1,size(A_patho)
            if (index(trim(A_patho(i)%name),"patho")/=0) then
                a_patho_set=[a_patho_set,A_patho(i)]
            end if
        end do
    end function compute_pathological_attractor
    !##########################################################################!
    !####################    compute_therapeutic_bullet    ####################!
    !##########################################################################!
    function compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio) result(therapeutic_bullet_set)
        logical::in_patho
        integer::r_min,r_max,max_targ,max_moda,n_node,i1,i2,i3,i4,i5
        integer,dimension(:,:),allocatable::C_targ
        real,dimension(:)::value
        real,dimension(:,:)::D
        real,dimension(:,:),allocatable::C_moda
        type(attractor),dimension(:)::A_physio
        type(attractor),dimension(:),allocatable::A_patho
        type(bullet)::bull
        type(bullet),dimension(:),allocatable::therapeutic_bullet_set
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        allocate(therapeutic_bullet_set(0))
        do i1=r_min,min(r_max,n_node)
            C_targ=int(generate_combination(real(range_int(1,n_node)),i1,max_targ))
            C_moda=generate_arrangement(value,i1,max_moda)
            do i2=1,size(C_targ,1)
                do i3=1,size(C_moda,1)
                    bull%targ=C_targ(i2,:)
                    bull%moda=C_moda(i3,:)
                    A_patho=compute_attractor(f,D,2,A_physio,bull)
                    if (size(compute_pathological_attractor(A_patho))==0) then
                        allocate(bull%unrecovered(0))
                        if (compare_attractor_set(A_physio,A_patho)) then
                            bull%metal="silver"
                            do i4=1,size(A_physio)
                                in_patho=.false.
                                do i5=1,size(A_patho)
                                    if (.not. compare_attractor(A_physio(i4),A_patho(i5))) then
                                        in_patho=.true.
                                        exit
                                    end if
                                end do
                                if (.not. in_patho) then
                                    bull%unrecovered=[bull%unrecovered,trim(A_physio(i4)%name)]
                                end if
                            end do
                        else
                            bull%metal="golden"
                        end if
                        therapeutic_bullet_set=[therapeutic_bullet_set,bull]
                        deallocate(bull%unrecovered)
                    end if
                end do
            end do
        end do
        deallocate(A_patho,C_targ,C_moda,bull%targ,bull%moda)
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
    !#######################    generate_arrangement    #######################!
    !##########################################################################!
    function generate_arrangement(deck,k,n_arrang) result(arrang_mat)
        !#################    /!\ only with repetition /!\    #################!
        integer::k,n_arrang,i1,i2
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
        integer::k,n_combi,i1,i2
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
                if (any(z==combi(:i2-1))) then
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
        else
            do i1=1,n
                y(:i1-1,(2**i1)/2+1:2**i1)=y(:i1-1,:2**(i1-1))
                do i2=1,2**(i1-1)
                    y(i1,i2)=0.0
                end do
                do i2=(2**i1)/2+1,2**i1
                    y(i1,i2)=1.0
                end do
            end do
        end if
    end function generate_state_space
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
            write (unit=*,fmt="(a)") "Too bad, your operating system does not provide a random number generator."
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
            read (unit=1,fmt=*) A_set(i1)%name
            allocate(A_set(i1)%att(n,m))
        end do
        do i1=1,size(A_set)
            do i2=1,size(A_set(i1)%att,1)
                read (unit=1,fmt=*) A_set(i1)%att(i2,:)
            end do
        end do
        close (unit=1)
    end function load_attractor_set
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
    !#######################    report_attractor_set    #######################!
    !##########################################################################!
    subroutine report_attractor_set(A_set,setting,V,boolean)
        logical::boolean
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
            if (size(A_set(i1)%att,2)==1) then
                n_point=n_point+1
            else
                n_cycle=n_cycle+1
            end if
            report=report//trim(A_set(i1)%name)//new_line("a")//"basin: "//real2char(A_set(i1)%popularity)//"% (of the state space)"//new_line("a")
            do i2=1,size(A_set(i1)%att,1)
                report=report//V(i2)//" "
                do i3=1,size(A_set(i1)%att,2)-1
                    if (boolean) then
                        report=report//int2char(int(A_set(i1)%att(i2,i3)))//" "
                    else
                        report=report//real2char(A_set(i1)%att(i2,i3))//" "
                    end if
                end do
                if (boolean) then
                    report=report//int2char(int(A_set(i1)%att(i2,size(A_set(i1)%att,2))))//new_line("a")
                else
                    report=report//real2char(A_set(i1)%att(i2,size(A_set(i1)%att,2)))//new_line("a")
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
                s=s//int2char(size(A_set(i1)%att,1))//new_line("a")//int2char(size(A_set(i1)%att,2))//new_line("a")//real2char(A_set(i1)%popularity)//new_line("a")//trim(A_set(i1)%name)//new_line("a")
            end do
            do i1=1,size(A_set)
                do i2=1,size(A_set(i1)%att,1)
                    do i3=1,size(A_set(i1)%att,2)-1
                        s=s//real2char(A_set(i1)%att(i2,i3))//","
                    end do
                    s=s//real2char(A_set(i1)%att(i2,size(A_set(i1)%att,2)))
                    if (i1/=size(A_set) .or. i2/=size(A_set(i1)%att,1)) then
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
        integer::n_gold,n_silv,i1,i2,save_
        character(1)::moda
        character(16),dimension(:)::V
        character(:),allocatable::report
        type(bullet),dimension(:)::therapeutic_bullet_set
        n_gold=0
        n_silv=0
        report=repeat("-",80)//new_line("a")
        do i1=1,size(therapeutic_bullet_set)
            if (trim(therapeutic_bullet_set(i1)%metal)=="golden") then
                n_gold=n_gold+1
            else
                n_silv=n_silv+1
            end if
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
            report=report//"("//trim(therapeutic_bullet_set(i1)%metal)
            if (trim(therapeutic_bullet_set(i1)%metal)=="silver") then
                report=report//", unrecovered:"
                do i2=1,size(therapeutic_bullet_set(i1)%unrecovered)-1
                    report=report//" "//trim(therapeutic_bullet_set(i1)%unrecovered(i2))//","
                end do
                report=report//" "//trim(therapeutic_bullet_set(i1)%unrecovered(size(therapeutic_bullet_set(i1)%unrecovered)))//")"//new_line("a")//repeat("-",80)//new_line("a")
            else
                report=report//")"//new_line("a")//repeat("-",80)//new_line("a")
            end if
        end do
        report=report//"found therapeutic bullets: "//int2char(size(therapeutic_bullet_set))//" ("//int2char(n_gold)//" golden bullets, "//int2char(n_silv)//" silver bullets)"
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
    !##########################    sort_attractor    ##########################!
    !##########################################################################!
    function sort_attractor(a) result(y)
        integer::i,j_min
        integer,dimension(:),allocatable::z
        type(attractor)::a,y
        y=a
        z=range_int(1,size(y%att,2))
        do i=1,size(y%att,1)
            z=z(minlocs(y%att(i,z)))
            if (size(z)==1) then
                j_min=z(1)
                exit
            end if
        end do
        y%att=cshift(y%att,j_min-1,2)
        deallocate(z)
    end function sort_attractor
    !##########################################################################!
    !########################    sort_attractor_set    ########################!
    !##########################################################################!
    function sort_attractor_set(A_set) result(y)
        !##########    /!\ attractors must be in sorted form /!\    ###########!
        logical::repass
        integer::i1,i2
        type(attractor)::z
        type(attractor),dimension(:)::A_set
        type(attractor),dimension(size(A_set))::y
        y=A_set
        allocate(z%att(0,0))
        do
            repass=.false.
            do i1=1,size(y)-1
                if (size(y(i1)%att,2)>size(y(i1+1)%att,2)) then
                    repass=.true.
                    z=y(i1)
                    y(i1)=y(i1+1)
                    y(i1+1)=z
                end if
            end do
            if (.not. repass) then
                exit
            end if
        end do
        do
            repass=.false.
            do i1=1,size(y)-1
                if (size(y(i1)%att,2)==size(y(i1+1)%att,2)) then
                    do i2=1,size(y(i1)%att,1)
                        if (y(i1)%att(i2,1)<y(i1+1)%att(i2,1)) then
                            exit
                        else if (y(i1)%att(i2,1)>y(i1+1)%att(i2,1)) then
                            repass=.true.
                            z=y(i1)
                            y(i1)=y(i1+1)
                            y(i1+1)=z
                            exit
                        end if
                    end do
                end if
            end do
            if (.not. repass) then
                exit
            end if
        end do
        deallocate(z%att)
    end function sort_attractor_set
    !##########################################################################!
    !############################    what_to_do    ############################!
    !##########################################################################!
    subroutine what_to_do(f,value,size_D,n_node,max_targ,max_moda,V)
        logical::boolean
        integer::size_D,n_node,max_targ,max_moda,to_do,r_min,r_max,setting,comprehensive_D
        real::start,finish
        real,dimension(:)::value
        real,dimension(:,:),allocatable::D
        character(16),dimension(:)::V
        type(attractor),dimension(:),allocatable::A_physio,A_patho,a_patho_set
        type(bullet),dimension(:),allocatable::therapeutic_bullet_set
        interface
            function f(x,k) result(y)
                integer::k
                real,dimension(:,:)::x
                real,dimension(size(x,1),1)::y
            end function f
        end interface
        call cpu_time(start)
        call init_random_seed()
        if (size(value)==2) then
            boolean=all(value==[0.0,1.0])
        else
            boolean=.false.
        end if
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
                write (unit=*,fmt="(a)",advance="no") new_line("a")//"[1] physiological"//new_line("a")//"[2] pathological"//new_line("a")//new_line("a")//"setting [1/2] "
                read (unit=*,fmt=*) setting
                select case (setting)
                    case (1)
                        A_physio=compute_attractor(f,D,1)
                        call report_attractor_set(A_physio,1,V,boolean)
                        deallocate(A_physio,D)
                    case (2)
                        A_physio=load_attractor_set(1)
                        A_patho=compute_attractor(f,D,2,A_physio)
                        call report_attractor_set(A_patho,2,V,boolean)
                        deallocate(A_physio,A_patho,D)
                end select
            case (2)
                A_patho=load_attractor_set(2)
                a_patho_set=compute_pathological_attractor(A_patho)
                call report_attractor_set(a_patho_set,3,V,boolean)
                deallocate(A_patho,a_patho_set)
            case (3)
                A_physio=load_attractor_set(1)
                write (unit=*,fmt="(a)",advance="no") new_line("a")//"number of targets per bullet (lower bound): "
                read (unit=*,fmt=*) r_min
                write (unit=*,fmt="(a)",advance="no") "number of targets per bullet (upper bound): "
                read (unit=*,fmt=*) r_max
                therapeutic_bullet_set=compute_therapeutic_bullet(f,D,r_min,r_max,max_targ,max_moda,n_node,value,A_physio)
                call report_therapeutic_bullet_set(therapeutic_bullet_set,V,boolean)
                deallocate(A_physio,therapeutic_bullet_set,D)
            case (4)
                write (unit=*,fmt="(a)") new_line("a")//"1) compute attractors with f_physio"//new_line("a")//"2) compute attractors with f_patho"//new_line("a")//"3) eventually compute pathological attractors"//new_line("a")//"4) compute therapeutic bullets with f_patho"//new_line("a")//new_line("a")//"do not forget to recompile the sources following any modification"
            case (5)
                write (unit=*,fmt="(a)") new_line("a")//'Copyright (c) 2013-2014, Arnaud Poret arnaud.poret@gmail.com'//new_line("a")//'All rights reserved.'//new_line("a")//new_line("a")//'Redistribution and use in source and binary forms, with or without modification,'//new_line("a")//'are permitted provided that the following conditions are met:'//new_line("a")//new_line("a")//'1. Redistributions of source code must retain the above copyright notice, this'//new_line("a")//'list of conditions and the following disclaimer.'//new_line("a")//new_line("a")//'2. Redistributions in binary form must reproduce the above copyright notice,'//new_line("a")//'this list of conditions and the following disclaimer in the documentation and/or'//new_line("a")//'other materials provided with the distribution.'//new_line("a")//new_line("a")//'3. Neither the name of the copyright holder nor the names of its contributors'//new_line("a")//'may be used to endorse or promote products derived from this software without'//new_line("a")//'specific prior written permission.'//new_line("a")//new_line("a")//'THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND'//new_line("a")//'ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED'//new_line("a")//'WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE'//new_line("a")//'DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR'//new_line("a")//'ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES'//new_line("a")//'(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;'//new_line("a")//'LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON'//new_line("a")//'ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT'//new_line("a")//'(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS'//new_line("a")//'SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.'
        end select
        call cpu_time(finish)
        write (unit=*,fmt="(a)") new_line("a")//"done in "//int2char(int(finish-start))//" CPU seconds"//new_line("a")
    end subroutine what_to_do
end module lib
