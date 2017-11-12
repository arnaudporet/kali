# This script takes the file "your_functional_equations.go" containing the
# functional version of the model equations and returns the file
# "your_functional_equations_xed.go".

# The file "your_functional_equations_xed.go" contains the functional version of
# the model equations with the node names replaced by the corresponding
# positions in the state vector x.

# This is the version of the model equations used by kali.

# It also contains the list of the nodes names.

# The equations must be as follows: equation// node_name

# Run ``python your_functional_equations.go'' in a terminal emulator.

# This script is coded in Python3, no warranties that it works with Python2.

import re,sys
equations=[]
nodes=[]
for line in open(sys.argv[1],"rt").read().splitlines():
    line=line.split("//")
    equations.append(line[0].strip().replace("min(","kali.Min(").replace("max(","kali.Max("))
    nodes.append(line[1].strip())
for j in range(len(nodes)):
    for i in range(len(equations)):
        equations[i]=re.sub("\\b"+nodes[j]+"\\b","x["+str(j)+"]",equations[i])
for i in range(len(equations)):
    equations[i]=equations[i]+",// "+nodes[i]
equations="\n".join(equations)+"\n"
nodes="\""+"\",\n\"".join(nodes)+"\",\n"
open(sys.argv[1].replace(".go","_xed.go"),"w").write("\n".join([nodes,equations]))
