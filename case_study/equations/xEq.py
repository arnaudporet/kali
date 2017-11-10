# This script takes the functional version of the model equations and returns
# them ready to be pasted in the main file at the appropriate place.

# Also returns the list of the nodes names ready to be pasted in the main file
# at the appropriate place.

# Run ``python eq_func.go'' in a terminal, where eq_func.go is the file
# containing the functional version of the model equations.

import re,sys
equations=[]
nodes=[]
for line in open(sys.argv[1],"rt").read().splitlines():
    line=line.split("// ")
    equations.append(line[0].replace("min(","kali.Min(").replace("max(","kali.Max("))
    nodes.append(line[1])
for j in range(len(nodes)):
    for i in range(len(equations)):
        equations[i]=re.sub("\\b"+nodes[j]+"\\b","x["+str(j)+"]",equations[i])
for i in range(len(equations)):
    equations[i]=equations[i]+",// "+nodes[i]
equations="\n".join(equations)+"\n"
nodes="\""+"\",\n\"".join(nodes)+"\",\n"
open(sys.argv[1].replace(".go","_xed.go"),"w").write("\n".join([nodes,equations]))
