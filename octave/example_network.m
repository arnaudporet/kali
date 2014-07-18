# How to:
#    1) read the comments
#    2) fill the template
#    3) open a terminal
#    4) launch octave: cd ~/kali-targ/octave/ && octave
#    5) run the example network: run("~/kali-targ/octave/example_network.m")

# GNU Octave (http://www.gnu.org/software/octave/) is a high-level interpreted
# language, primarily intended for numerical computations. 

# The example network is a boolean model of the mammalian cell cycle [1].

# [1] Fauré, A., Naldi, A., Chaouiya, C., & Thieffry, D. (2006). Dynamical
# analysis of a generic Boolean model for the control of the mammalian cell
# cycle. Bioinformatics, 22(14), e124-e131.

clear all
clc
more off
addpath("~/kali-targ/octave/")

# the node names
V={"CycD","Rb","E2F","CycE","CycA","p27","Cdc20","Cdh1","UbcH10","CycB"};

# the domain of values, for example [0,1] for boolean logic and [0,0.5,1] for
# three valued logic
value=[0,1];

# the size of the subset of the state space to start from
size_D=50;

# the maximum number of target combinations to test
max_targ=50;

# the maximum number of modality arrangements to test for each target
# combination
max_moda=50;

# the boolean transition function of the physiological variant
# to cope with both boolean and multivalued logic, the Zadeh fuzzy logic
# operators are used
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

# the boolean transition function of the pathological variant
# to cope with both boolean and multivalued logic, the Zadeh fuzzy logic
# operators are used
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

# pass either f_physio (for computing the physiological attractor set) or
# f_patho (for computing the phathological attractor set or to compute therapeutic bullets)
what_to_do("f_physio",V,size_D,max_targ,max_moda,value)

