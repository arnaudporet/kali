#how to:
#    1) read the comments
#    2) fill the template
#    3) open a terminal
#    4) past: cd ~/kali-targ/octave/ && octave --eval "run('~/kali-targ/octave/example_network.m')"
#    5) press Enter

#this example network is an implementation of a boolean model of the mammalian
#cell cycle proposed by Adrien Faure et al: Aurelien Naldi, Claudine Chaouiya,
#and Denis Thieffry. Dynamical analysis of a generic boolean model for the
#control of the mammalian cell cycle. Bioinformatics, 22(14):e124–e131, 2006.

clear all
clc
more off
addpath("~/kali-targ/octave/")

#the node names
V={"CycD","Rb","E2F","CycE","CycA","p27","Cdc20","Cdh1","UbcH10","CycB"};

#the domain of values, for example [0,1] for boolean logic and [0,0.5,1] for
#three valued logic
value=[0,1];

#the size of the subset of the state space to start from
size_D=50;

#the maximum number of target combinations to test
max_targ=50;

#the maximum number of modality arrangements to test for each target combination
max_moda=50;

function y=f_physio(x,k)
    #the boolean transition function of the physiological variant
    #to cope with both boolean and multivalued logic, the Zadeh fuzzy logic
    #operators are used
    y=[
    #CycD
    x(1,k);
    #Rb
    max([min([1-x(1,k),1-x(4,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-x(1,k),1-x(10,k)])]);
    #E2F
    max([min([1-x(2,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-x(2,k),1-x(10,k)])]);
    #CycE
    min([x(3,k),1-x(2,k)]);
    #CycA
    max([min([x(3,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])]),min([x(5,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])])]);
    #p27
    max([min([1-x(1,k),1-x(4,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-min([x(4,k),x(5,k)]),1-x(10,k),1-x(1,k)])]);
    #Cdc20
    x(10,k);
    #Cdh1
    max([min([1-x(5,k),1-x(10,k)]),x(7,k),min([x(6,k),1-x(10,k)])]);
    #UbcH10
    max([1-x(8,k),min([x(8,k),x(9,k),max([x(7,k),x(5,k),x(10,k)])])]);
    #CycB
    min([1-x(7,k),1-x(8,k)])
    ];
endfunction

function y=f_patho(x,k)
    #the boolean transition function of the pathological variant
    #to cope with both boolean and multivalued logic, the Zadeh fuzzy logic
    #operators are used
    y=[
    #CycD
    x(1,k);
    #Rb
    0;
    #E2F
    max([min([1-x(2,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-x(2,k),1-x(10,k)])]);
    #CycE
    min([x(3,k),1-x(2,k)]);
    #CycA
    max([min([x(3,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])]),min([x(5,k),1-x(2,k),1-x(7,k),1-min([x(8,k),x(9,k)])])]);
    #p27
    max([min([1-x(1,k),1-x(4,k),1-x(5,k),1-x(10,k)]),min([x(6,k),1-min([x(4,k),x(5,k)]),1-x(10,k),1-x(1,k)])]);
    #Cdc20
    x(10,k);
    #Cdh1
    max([min([1-x(5,k),1-x(10,k)]),x(7,k),min([x(6,k),1-x(10,k)])]);
    #UbcH10
    max([1-x(8,k),min([x(8,k),x(9,k),max([x(7,k),x(5,k),x(10,k)])])]);
    #CycB
    min([1-x(7,k),1-x(8,k)])
    ];
endfunction

#pass either f_physio (for computing the physiological attractor set) or f_patho
#(for computing the phathological attractor set or to compute therapeutic
#bullets)
what_to_do("f_physio",V,size_D,max_targ,max_moda,value)

