!Copyright (c) 2013-2014, Arnaud Poret
!All rights reserved.

!read the following comments, fill the following template, open a terminal, past this command and press Enter: cd ~/kali-targ/ && gfortran lib.f95 example_network.f95 -o example_network && ./example_network

!at line 31, allocate to V the number of nodes

!value: the domain of the variables, for example [0.0,1.0] for boolean logic and [0.0,0.5,1.0] for three valued logic

!max_targ: the maximum number of target combinations to test

!max_moda: the maximum number of modality arrangements to test for each target combination

!size_D: the size of the subset of the state space to start from

!V: the node names

!f_physio: the boolean transition function of the physiological variant

!f_patho: the boolean transition function of the pathological variant

!to cope with boolean and multivalued logic, the Zadeh fuzzy logic operators are used

!at line 46, pass either f_physio (for computing the physiological attractor set) or f_patho (for computing the phathological attractor set or to compute therapeutic bullets) to the subroutine what_to_do

!this example network is an implementation of a boolean model of the mammalian cell cycle proposed by Adrien Faure et al: Aurelien Naldi, Claudine Chaouiya, and Denis Thieffry. Dynamical analysis of a generic boolean model for the control of the mammalian cell cycle. Bioinformatics, 22(14):e124–e131, 2006.

program example_network
    use lib
    implicit none
    allocate(V(10))
    value=[0.0,1.0]
    max_targ=1e4
    max_moda=1e4
    size_D=1e4
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
    call what_to_do(f_patho,V,max_targ,max_moda,size_D,value)
    deallocate(value,V)
    contains
    !##########################################################################!
    !#############################    f_physio    #############################!
    !##########################################################################!
    function f_physio(x,k) result(y)
        implicit none
        real,dimension(:,:)::x
        integer::k
        real,dimension(size(x,1),1)::y
        y(1,1)=x(1,k)!CycD
        y(2,1)=max(min(1.0-x(1,k),1.0-x(4,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-x(1,k),1.0-x(10,k)))!Rb
        y(3,1)=max(min(1.0-x(2,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-x(2,k),1.0-x(10,k)))!E2F
        y(4,1)=min(x(3,k),1.0-x(2,k))!CycE
        y(5,1)=max(min(x(3,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),x(9,k))),min(x(5,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),&
        x(9,k))))!CycA
        y(6,1)=max(min(1.0-x(1,k),1.0-x(4,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-min(x(4,k),x(5,k)),1.0-x(10,k),1.0-x(1,k)))!p27
        y(7,1)=x(10,k)!Cdc20
        y(8,1)=max(min(1.0-x(5,k),1.0-x(10,k)),x(7,k),min(x(6,k),1.0-x(10,k)))!Cdh1
        y(9,1)=max(1.0-x(8,k),min(x(8,k),x(9,k),max(x(7,k),x(5,k),x(10,k))))!UbcH10
        y(10,1)=min(1.0-x(7,k),1.0-x(8,k))!CycB
    end function f_physio
    !##########################################################################!
    !#############################    f_patho    ##############################!
    !##########################################################################!
    function f_patho(x,k) result(y)
        implicit none
        real,dimension(:,:)::x
        integer::k
        real,dimension(size(x,1),1)::y
        y(1,1)=x(1,k)!CycD
        y(2,1)=0.0!Rb
        y(3,1)=max(min(1.0-x(2,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-x(2,k),1.0-x(10,k)))!E2F
        y(4,1)=min(x(3,k),1.0-x(2,k))!CycE
        y(5,1)=max(min(x(3,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),x(9,k))),min(x(5,k),1.0-x(2,k),1.0-x(7,k),1.0-min(x(8,k),&
        x(9,k))))!CycA
        y(6,1)=max(min(1.0-x(1,k),1.0-x(4,k),1.0-x(5,k),1.0-x(10,k)),min(x(6,k),1.0-min(x(4,k),x(5,k)),1.0-x(10,k),1.0-x(1,k)))!p27
        y(7,1)=x(10,k)!Cdc20
        y(8,1)=max(min(1.0-x(5,k),1.0-x(10,k)),x(7,k),min(x(6,k),1.0-x(10,k)))!Cdh1
        y(9,1)=max(1.0-x(8,k),min(x(8,k),x(9,k),max(x(7,k),x(5,k),x(10,k))))!UbcH10
        y(10,1)=min(1.0-x(7,k),1.0-x(8,k))!CycB
    end function f_patho
end program example_network

!##############################################################################!
!#############################      LICENSE       #############################!
!#############################    BSD 3-Clause    #############################!
!##############################################################################!

!Copyright (c) 2013-2014, Arnaud Poret
!All rights reserved.

!Redistribution and use in source and binary forms, with or without modification,
!are permitted provided that the following conditions are met:

!1. Redistributions of source code must retain the above copyright notice, this
!list of conditions and the following disclaimer.

!2. Redistributions in binary form must reproduce the above copyright notice,
!this list of conditions and the following disclaimer in the documentation and/or
!other materials provided with the distribution.

!3. Neither the name of the copyright holder nor the names of its contributors
!may be used to endorse or promote products derived from this software without
!specific prior written permission.

!THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
!ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
!WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
!DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
!ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
!(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
!LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
!ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
!(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
!SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

