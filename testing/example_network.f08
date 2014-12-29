! Copyright (C) 2013-2014 Arnaud Poret
! This program is licensed under the GNU General Public License.
! To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

! gf08W lib.f08 example_network.f08 && ./a.out && rm a.out lib.mod

program example_network
    use lib
    n_node=7
    value=[0.0,0.5,1.0]
    max_targ=int(1e2)
    max_moda=int(1e2)
    size_D=int(1e4)
    allocate(V(n_node))
    V(1)="CycD"
    V(2)="Rb"
    V(3)="E2F"
    V(4)="CycA"
    V(5)="p27"
    V(6)="UbcH10"
    V(7)="CycB"
    call what_to_do(f_physio,f_patho,value,size_D,n_node,max_targ,max_moda,V)
    deallocate(value,V)
    contains
    !##########################################################################!
    !#############################    f_physio    #############################!
    !##########################################################################!
    function f_physio(x,k) result(y)
        integer::k
        real,dimension(:,:)::x
        real,dimension(size(x,1))::y
        y(1)=x(1,k)!CycD
        y(2)=1.0-max(x(1,k),x(7,k),min(max(min(x(3,k),1.0-x(2,k)),x(4,k)),1.0-x(5,k)))!Rb
        y(3)=1.0-max(x(2,k),x(7,k),min(x(4,k),1.0-x(5,k)))!E2F
        y(4)=1.0-max(x(2,k),x(7,k),min(max(x(7,k),min(x(4,k),1.0-x(5,k))),x(6,k)),1.0-max(x(3,k),x(4,k)))!CycA
        y(5)=1.0-max(x(1,k),x(7,k),min(max(min(x(3,k),1.0-x(2,k)),x(4,k)),1.0-x(5,k)))!p27
        y(6)=1.0-max(max(x(7,k),min(x(4,k),1.0-x(5,k))),min(max(x(7,k),min(x(4,k),1.0-x(5,k))),x(6,k),max(x(4,k),x(7,k))))!UbcH10
        y(7)=1.0-max(x(7,k),max(x(7,k),min(x(4,k),1.0-x(5,k))))!CycB
    end function f_physio
    !##########################################################################!
    !#############################    f_patho    ##############################!
    !##########################################################################!
    function f_patho(x,k) result(y)
        integer::k
        real,dimension(:,:)::x
        real,dimension(size(x,1))::y
        y(1)=x(1,k)!CycD
        y(2)=0.0!Rb
        y(3)=1.0-max(x(2,k),x(7,k),min(x(4,k),1.0-x(5,k)))!E2F
        y(4)=1.0-max(x(2,k),x(7,k),min(max(x(7,k),min(x(4,k),1.0-x(5,k))),x(6,k)),1.0-max(x(3,k),x(4,k)))!CycA
        y(5)=1.0-max(x(1,k),x(7,k),min(max(min(x(3,k),1.0-x(2,k)),x(4,k)),1.0-x(5,k)))!p27
        y(6)=1.0-max(max(x(7,k),min(x(4,k),1.0-x(5,k))),min(max(x(7,k),min(x(4,k),1.0-x(5,k))),x(6,k),max(x(4,k),x(7,k))))!UbcH10
        y(7)=1.0-max(x(7,k),max(x(7,k),min(x(4,k),1.0-x(5,k))))!CycB
    end function f_patho
end program example_network
