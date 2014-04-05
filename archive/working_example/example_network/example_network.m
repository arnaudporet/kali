#Copyright (c) 2013, Arnaud Poret
#All rights reserved.

#run("~/kali/kali-targ/working_example/example_network/a.m")

clear all
clc
more off
addpath("~/kali/kali-targ/lib/")
cd("~/kali/kali-targ/working_example/example_network/")

V={"CycD","Rb","E2F","CycE","CycA","p27","Cdc20","Cdh1","UbcH10","CycB"};

value=[0,1];
size_D=50;
max_targ=50;
max_moda=50;

function y=f(x,k)
    y=[
    x(1,k);#CycD
    or(and(not(x(1,k)),not(x(4,k)),not(x(5,k)),not(x(10,k))),and(x(6,k),not(x(1,k)),not(x(10,k))));#Rb
    or(and(not(x(2,k)),not(x(5,k)),not(x(10,k))),and(x(6,k),not(x(2,k)),not(x(10,k))));#E2F
    and(x(3,k),not(x(2,k)));#CycE
    or(and(x(3,k),not(x(2,k)),not(x(7,k)),not(and(x(8,k),x(9,k)))),and(x(5,k),not(x(2,k)),not(x(7,k)),not(and(x(8,k),x(9,k)))));#CycA
    or(and(not(x(1,k)),not(x(4,k)),not(x(5,k)),not(x(10,k))),and(x(6,k),not(and(x(4,k),x(5,k))),not(x(10,k)),not(x(1,k))));#p27
    x(10,k);#Cdc20
    or(and(not(x(5,k)),not(x(10,k))),x(7,k),and(x(6,k),not(x(10,k))));#Cdh1
    or(not(x(8,k)),and(x(8,k),x(9,k),or(x(7,k),x(5,k),x(10,k))));#UbcH10
    and(not(x(7,k)),not(x(8,k)))#CycB
    ];
endfunction

what_to_do("f",V,size_D,max_targ,max_moda,value)

