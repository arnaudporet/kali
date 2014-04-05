#Copyright (c) 2013, Arnaud Poret
#All rights reserved.
################################################################################
###############################    what_to_do    ###############################
################################################################################
function what_to_do(f,V,size_D,max_targ,max_moda,value)
    todo=menu("what to do: ","compute attractors","compute pathological attractors","compute therapeutic bullets");
    if todo!=2
        if all(value==[0,1],2)
            disp(strcat("size(S)=",num2str(2**numel(V))))
            if yes_or_no("comprehensive D? ")
                D=generate_state_space(numel(V))';
            else
                D=generate_arrangement(value,numel(V),size_D)';
            endif
        else
            D=generate_arrangement(value,numel(V),size_D)';
        endif
    endif
    switch todo
        case {1}
            A=compute_attractor(f,[],[],D);
            report_attractor_set(A,V)
        case {2}
            A_physio=load_set("A_physio: ");
            A_patho=load_set("A_patho: ");
            a_patho_set=compute_pathological_attractor(A_physio,A_patho);
            report_attractor_set(a_patho_set,V)
        case {3}
            A_physio=load_set("A_physio: ");
            r_min=input("r_min: ");
            r_max=input("r_max: ");
            [Targ,Moda,Metal]=compute_therapeutic_bullet(r_min,r_max,max_targ,max_moda,A_physio,f,V,D,value);
            report_therapeutic_bullet_set(Targ,Moda,Metal,V)
    endswitch
endfunction
################################################################################
###########################    compare_attractor    ############################
################################################################################
function differ=compare_attractor(a1,a2)
    differ=false;
    if not(columns(a1)==columns(a2))
        differ=true;
        return
    endif
    start_found=false;
    for j1=1:columns(a1)
        for j2=1:columns(a2)
            if a1(:,j1)==a2(:,j2)
                start_found=true;
                start1=j1;
                start2=j2;
                break
            endif
        endfor
        if start_found
            break
        endif
    endfor
    if not(start_found)
        differ=true;
        return
    endif
    for j_a=0:columns(a1)-2
        if not(all(a1(:,mod(start1+j_a,columns(a1))+1)==a2(:,mod(start2+j_a,columns(a2))+1),1))
            differ=true;
            return
        endif
    endfor
endfunction
################################################################################
#########################    compare_attractor_set    ##########################
################################################################################
function differ=compare_attractor_set(A1,A2)
    if not(numel(A1)==numel(A2))
        differ=true;
        return
    endif
    differ=[];
    for i1=1:numel(A1)
        in_2=false;
        for i2=1:numel(A2)
            if not(compare_attractor(A1{i1},A2{i2}))
                in_2=true;
                break
            endif
        endfor
        differ(1,i1)=in_2;
    endfor
    differ=not(all(differ,2));
endfunction
################################################################################
###########################    compute_attractor    ############################
################################################################################
function A=compute_attractor(f,c_targ,c_moda,D)
    A={};
    for i_D=1:columns(D)
        x=D(:,i_D);
        k=1;
        while true
            x(:,k+1)=feval(f,x,k);
            x(c_targ,k+1)=c_moda'*eye(rows(c_moda),1);
            a_found=false;
            for j_x=k:-1:1
                if all(x(:,j_x)==x(:,k+1),1)
                    a_found=true;
                    a=x(:,[j_x:k]);
                    break
                endif
            endfor
            if a_found
                in_A=false;
                for i_A=1:numel(A)
                    if not(compare_attractor(a,A{i_A}))
                        in_A=true;
                        break
                    endif
                endfor
                if not(in_A)
                    A{numel(A)+1}=a;
                endif
                break
            endif
            k+=1;
        endwhile
    endfor
endfunction
################################################################################
#####################    compute_pathological_attractor    #####################
################################################################################
function a_patho_set=compute_pathological_attractor(A_physio,A_patho)
    a_patho_set={};
    for i_patho=1:numel(A_patho)
        in_physio=false;
        for i_physio=1:numel(A_physio)
            if not(compare_attractor(A_patho{i_patho},A_physio{i_physio}))
                in_physio=true;
                break
            endif
        endfor
        if not(in_physio)
            a_patho_set{numel(a_patho_set)+1}=A_patho{i_patho};
        endif
    endfor
endfunction
################################################################################
#######################    compute_therapeutic_bullet    #######################
################################################################################
function [Targ,Moda,Metal]=compute_therapeutic_bullet(r_min,r_max,max_targ,max_moda,A_physio,f,V,D,value)
    Targ={};
    Moda={};
    Metal={};
    r_max=min(r_max,numel(V));
    for i_r=r_min:r_max
        C_targ=generate_combination(1:numel(V),i_r,max_targ);
        C_moda=generate_arrangement(value,i_r,max_moda);
        for i_targ=1:rows(C_targ)
            for i_moda=1:rows(C_moda)
                A_patho=compute_attractor(f,C_targ(i_targ,:),C_moda(i_moda,:),D);
                if numel(compute_pathological_attractor(A_physio,A_patho))==0
                    Targ{numel(Targ)+1}=C_targ(i_targ,:);
                    Moda{numel(Moda)+1}=C_moda(i_moda,:);
                    if compare_attractor_set(A_physio,A_patho)
                        Metal{numel(Metal)+1}="silver";
                    else
                        Metal{numel(Metal)+1}="golden";
                    endif
                endif
            endfor
        endfor
    endfor
endfunction
################################################################################
##########################    generate_arrangement    ##########################
################################################################################
function arrang_mat=generate_arrangement(deck,k,n_arrang)
    ####################    /!\ only with repetition /!\    ####################
    n_arrang=min(n_arrang,columns(deck)**k);
    arrang_mat=[]*eye(0,k);
    while rows(arrang_mat)<n_arrang
        do
            arrang=discrete_rnd(deck,(1/columns(deck))*ones(size(deck)),1,k);
        until all(any(ones(rows(arrang_mat),1)*arrang!=arrang_mat,2),1)
        arrang_mat(rows(arrang_mat)+1,:)=arrang;
    endwhile
endfunction
################################################################################
##########################    generate_combination    ##########################
################################################################################
function combi_mat=generate_combination(deck,k,n_combi)
    ##################    /!\ only without repetition /!\    ###################
    n_combi=min(n_combi,factorial(columns(deck))/(factorial(k)*factorial(columns(deck)-k)));
    k=min(k,columns(deck));
    combi_mat=[]*eye(0,k);
    while rows(combi_mat)<n_combi
        do
            combi=eye(1,0)*[];
            while columns(combi)<k
                do
                    x=discrete_rnd(deck,(1/columns(deck))*ones(size(deck)),1,1);
                until all(x*ones(size(combi))!=combi,2)
                combi(1,columns(combi)+1)=x;
            endwhile
            combi=sort(combi,2,"ascend");
        until all(any(ones(rows(combi_mat),1)*combi!=combi_mat,2),1)
        combi_mat(rows(combi_mat)+1,:)=combi;
    endwhile
endfunction
################################################################################
##########################    generate_state_space    ##########################
################################################################################
function S=generate_state_space(n)
    S=[0;1];
    z=[];
    for i_n=1:n-1
        for i_S=1:rows(S)
            z(rows(z)+1,:)=[S(i_S,:),0];
            z(rows(z)+1,:)=[S(i_S,:),1];
        endfor
        S=z;
        z=[];
    endfor
endfunction
################################################################################
################################    load_set    ################################
################################################################################
function _set=load_set(prompt)
    ls()
    name=input(prompt,"s");
    load("-binary",name,name)
    eval(cstrcat("_set=",name,";"))
endfunction
################################################################################
##########################    report_attractor_set    ##########################
################################################################################
function report_attractor_set(A,V)
    name_len=[];
    for i_V=1:numel(V)
        name_len(1,i_V)=numel(V{i_V});
    endfor
    for i_V=1:numel(V)
        xtd_name=cstrcat(V{i_V},": ");
        for i_len=numel(V{i_V})+1:max(name_len)
            xtd_name=cstrcat(xtd_name," ");
        endfor
        V{i_V}=xtd_name;
    endfor
    n_point=0;
    n_cycle=0;
    sep="--------------------------------------------------------------------------------";
    report=cstrcat(sep,"\n");
    for i_a=1:numel(A)
        if columns(A{i_a})==1
            n_point+=1;
        else
            n_cycle+=1;
        endif
        for i_V=1:numel(V)
            report=cstrcat(report,cstrcat(V{i_V},mat2str(A{i_a}(i_V,:)),"\n"));
        endfor
        report=cstrcat(report,sep,"\n");
    endfor
    report=cstrcat(report,"found attractors: ",num2str(numel(A))," (",num2str(n_point)," points, ",num2str(n_cycle)," cycles)");
    disp(report)
    save_report(A,report)
endfunction
################################################################################
#####################    report_therapeutic_bullet_set    ######################
################################################################################
function report_therapeutic_bullet_set(Targ,Moda,Metal,V)
    n_gold=0;
    n_silv=0;
    sep="--------------------------------------------------------------------------------";
    report=cstrcat(sep,"\n");
    for i_targ=1:numel(Targ)
        if all(Metal{i_targ}=="golden",2)
            n_gold+=1;
        else
            n_silv+=1;
        endif
        for j_targ=1:columns(Targ{i_targ})
            report=cstrcat(report,V{Targ{i_targ}(:,j_targ)},"[",num2str(Moda{i_targ}(:,j_targ)),"] ");
        endfor
        report=cstrcat(report,"(",Metal{i_targ}," bullet)","\n",sep,"\n");
    endfor
    report=cstrcat(report,"found therapeutic bullets: ",num2str(numel(Targ))," (",num2str(n_gold)," golden bullets, ",num2str(n_silv)," silver bullets)");
    disp(report)
    save_report({Targ,Moda,Metal},report)
endfunction
################################################################################
##############################    save_report    ###############################
################################################################################
function save_report(_set,report)
    if yes_or_no("save? ")
        ls()
        name=input("name: ","s");
        eval(cstrcat(name,"=_set;"))
        save("-binary",name,name)
        eval(cstrcat(name,"=report;"))
        save("-text",cstrcat(name,".txt"),name)
    endif
endfunction
