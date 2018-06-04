# This script takes the file ``your_functional_equations.go'' containing the
# functional version of the model equations and returns the file
# ``your_functional_equations_xed.go''.

# The file ``your_functional_equations_xed.go'' contains the functional version
# of the model equations with the node names replaced by their corresponding
# position in the state vector x.

# This is the version of the model equations used by kali.

# It also contains the list of the nodes names.

# The equations must be as follows: equation// node_name

# Run ``python xEq.py your_functional_equations.go'' in a terminal emulator.

# This script is coded in Python3, no warranties that it works with Python2.

import sys
if len(sys.argv)!=2:
    print("usage: python xEq.py your_functional_equations.go")
elif not sys.argv[1].endswith(".go"):
    print("only accepts Go files (i.e. with the \".go\" file extension)")
else:
    equations=[]
    nodes=[]
    for line in open(sys.argv[1],"rt").read().splitlines():
        if line.count("//")!=1:
            quit("the equations must be as follows: equation// node_name")
        else:
            line=line.split("//")
            equations.append(line[0].strip().replace("min(","kali.Min(").replace("max(","kali.Max("))
            nodes.append(line[1].strip())
    if len(equations)==0:
        print("unable to find equations")
    else:
        import re
        for i in range(len(equations)):
            for j in range(len(nodes)):
                equations[i]=re.sub("\\b"+nodes[j]+"\\b","x["+str(j)+"]",equations[i])
        for i in range(len(equations)):
            equations[i]=equations[i]+",// "+nodes[i]
        equations="\n".join(equations)+"\n"
        nodes="\""+"\",\n\"".join(nodes)+"\",\n"
        open(sys.argv[1].replace(".go","_xed.go"),"w").write("\n".join([nodes,equations]))
