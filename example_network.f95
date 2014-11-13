! How to:
!     1) read the comments
!     2) fill the template
!     3) compile and execute: cd ~/kali-targ/ && gfortran lib.f95 example_network.f95 -o example_network && ./example_network

! GFortran (https://www.gnu.org/software/gcc/fortran/) is the Fortran compiler
! front end and run-time libraries for GCC, the GNU Compiler Collection.

! Do not forget to recompile the sources following any modification.

! The example network is a published boolean model of the mammalian cell
! cycle [1].

! [1] Fauré, A., Naldi, A., Chaouiya, C., & Thieffry, D. (2006). Dynamical
! analysis of a generic Boolean model for the control of the mammalian cell
! cycle. Bioinformatics, 22(14), e124-e131.

program example_network
    use lib
    implicit none
    ! the number of nodes
    n_node=10
    ! the domain of value, for example [0.0,1.0] for boolean logic and
    ! [0.0,0.5,1.0] for three-valued logic
    value=[0.0,1.0]
    ! the maximum number of target combinations to test
    max_targ=int(1e2)
    ! the maximum number of modality arrangements to test for each target
    ! combination
    max_moda=int(1e2)
    ! the size of the subset of the state space to compute
    size_D=int(1e4)
    ! the node names
    allocate(V(n_node))
    V(1)="CycD"
    V(2)="Rb"
    V(3)="E2F"
    V(4)="CycE"
    V(5)="CycA"
    V(6)="p27"
    V(7)="Cdc20"
    V(8)="Cdh1"
    V(9)="UbcH10"
    V(10)="CycB"
    ! pass either f_physio (for computing the physiological attractor set) or
    ! f_patho (for computing the phathological attractor set or to compute
    ! therapeutic bullets)
    call what_to_do(f_physio,value,size_D,n_node,max_targ,max_moda,V)
    deallocate(value,V)
    contains
    !##########################################################################!
    !#############################    f_physio    #############################!
    !##########################################################################!
    ! the transition function of the physiological variant
    ! to cope with both boolean and multivalued logic, the Zadeh fuzzy logic
    ! operators are used: x AND y = min(x,y), x OR y = max(x,y), NOT x = 1-x
    function f_physio(x,k) result(y)
        implicit none
        integer::k
        real,dimension(:,:)::x
        real,dimension(size(x,1),1)::y
        y(1,1)=x(1,k)!CycD
        y(2,1)=max(min(1.0-x(1,k),1.0-x(4,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-x(1,k),1.0-x(10,k)))!Rb
        y(3,1)=max(min(1.0-x(2,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-x(2,k),1.0-x(10,k)))!E2F
        y(4,1)=min(x(3,k),1.0-x(2,k))!CycE
        y(5,1)=max(min(x(3,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),x(9,k))),&
        min(x(5,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),x(9,k))))!CycA
        y(6,1)=max(min(1.0-x(1,k),1.0-x(4,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-min(x(4,k),x(5,k)),1.0-x(10,k),1.0-x(1,k)))!p27
        y(7,1)=x(10,k)!Cdc20
        y(8,1)=max(min(1.0-x(5,k),1.0-x(10,k)),x(7,k),min(x(6,k),1.0-x(10,k)))!Cdh1
        y(9,1)=max(1.0-x(8,k),min(x(8,k),x(9,k),max(x(7,k),x(5,k),x(10,k))))!UbcH10
        y(10,1)=min(1.0-x(7,k),1.0-x(8,k))!CycB
    end function f_physio
    !##########################################################################!
    !#############################    f_patho    ##############################!
    !##########################################################################!
    ! the transition function of the pathological variant
    ! to cope with both boolean and multivalued logic, the Zadeh fuzzy logic
    ! operators are used: x AND y = min(x,y), x OR y = max(x,y), NOT x = 1-x
    function f_patho(x,k) result(y)
        implicit none
        integer::k
        real,dimension(:,:)::x
        real,dimension(size(x,1),1)::y
        y(1,1)=x(1,k)!CycD
        y(2,1)=0.0!Rb
        y(3,1)=max(min(1.0-x(2,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-x(2,k),1.0-x(10,k)))!E2F
        y(4,1)=min(x(3,k),1.0-x(2,k))!CycE
        y(5,1)=max(min(x(3,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),x(9,k))),&
        min(x(5,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),x(9,k))))!CycA
        y(6,1)=max(min(1.0-x(1,k),1.0-x(4,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-min(x(4,k),x(5,k)),1.0-x(10,k),1.0-x(1,k)))!p27
        y(7,1)=x(10,k)!Cdc20
        y(8,1)=max(min(1.0-x(5,k),1.0-x(10,k)),x(7,k),min(x(6,k),1.0-x(10,k)))!Cdh1
        y(9,1)=max(1.0-x(8,k),min(x(8,k),x(9,k),max(x(7,k),x(5,k),x(10,k))))!UbcH10
        y(10,1)=min(1.0-x(7,k),1.0-x(8,k))!CycB
    end function f_patho
end program example_network
