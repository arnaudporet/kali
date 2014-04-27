#Copyright (c) 2013-2014, Arnaud Poret
#All rights reserved.

#read the following comments, fill the following template, open a terminal, launch octave, past this command and press Enter: run("~/kali-targ/octave/example_network.m")

#value: the domain of the variables, for example [0,1] for boolean logic and [0,0.5,1] for three valued logic

#max_targ: the maximum number of target combinations to test

#max_moda: the maximum number of modality arrangements to test for each target combination

#size_D: the size of the subset of the state space to start from

#V: the node names

#f_physio: the boolean transition function of the physiological variant

#f_patho: the boolean transition function of the pathological variant

#to cope with both boolean and multivalued logic, the Zadeh fuzzy logic operators are used

#at line 69, pass either f_physio (for computing the physiological attractor set) or f_patho (for computing the phathological attractor set or to compute therapeutic bullets) to the subroutine what_to_do

#this example network is an implementation of a boolean model of the mammalian cell cycle proposed by Adrien Faure et al: Aurelien Naldi, Claudine Chaouiya, and Denis Thieffry. Dynamical analysis of a generic boolean model for the control of the mammalian cell cycle. Bioinformatics, 22(14):e124–e131, 2006.

clear all
clc
more off
addpath("~/kali-targ/octave/")
cd("~/kali-targ/octave/")

V={"CycD","Rb","E2F","CycE","CycA","p27","Cdc20","Cdh1","UbcH10","CycB"};

value=[0,1];
size_D=50;
max_targ=50;
max_moda=50;

function y=f_physio(x,k)
    y=[
    x(1,k);#CycD
    max([min([1-x(1,k),1-x(4,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-x(1,k),1-x(10,k)])]);#Rb
    max([min([1-x(2,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-x(2,k),1-x(10,k)])]);#E2F
    min([x(3,k),1-x(2,k)]);#CycE
    max([min([x(3,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])]),min([x(5,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])])]);#CycA
    max([min([1-x(1,k),1-x(4,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-min([x(4,k),x(5,k)]),1-x(10,k),1-x(1,k)])]);#p27
    x(10,k);#Cdc20
    max([min([1-x(5,k),1-x(10,k)]),x(7,k),min([x(6,k),1-x(10,k)])]);#Cdh1
    max([1-x(8,k),min([x(8,k),x(9,k),max([x(7,k),x(5,k),x(10,k)])])]);#UbcH10
    min([1-x(7,k),1-x(8,k)])#CycB
    ];
endfunction

function y=f_patho(x,k)
    y=[
    x(1,k);#CycD
    0;#Rb
    max([min([1-x(2,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-x(2,k),1-x(10,k)])]);#E2F
    min([x(3,k),1-x(2,k)]);#CycE
    max([min([x(3,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])]),min([x(5,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])])]);#CycA
    max([min([1-x(1,k),1-x(4,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-min([x(4,k),x(5,k)]),1-x(10,k),1-x(1,k)])]);#p27
    x(10,k);#Cdc20
    max([min([1-x(5,k),1-x(10,k)]),x(7,k),min([x(6,k),1-x(10,k)])]);#Cdh1
    max([1-x(8,k),min([x(8,k),x(9,k),max([x(7,k),x(5,k),x(10,k)])])]);#UbcH10
    min([1-x(7,k),1-x(8,k)])#CycB
    ];
endfunction

what_to_do("f_patho",V,size_D,max_targ,max_moda,value)

################################################################################
##############################      LICENSE       ##############################
##############################    BSD 3-Clause    ##############################
################################################################################

#Copyright (c) 2013-2014, Arnaud Poret
#All rights reserved.

#Redistribution and use in source and binary forms, with or without modification,
#are permitted provided that the following conditions are met:

#1. Redistributions of source code must retain the above copyright notice, this
#list of conditions and the following disclaimer.

#2. Redistributions in binary form must reproduce the above copyright notice,
#this list of conditions and the following disclaimer in the documentation and/or
#other materials provided with the distribution.

#3. Neither the name of the copyright holder nor the names of its contributors
#may be used to endorse or promote products derived from this software without
#specific prior written permission.

#THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
#ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
#WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
#DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
#ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
#(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
#LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
#ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
#(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
#SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

