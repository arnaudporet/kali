
input_file=input("input file: ")
output_file=input("output file: ")

lines=open(input_file,"r").read().splitlines()
names=[]
sorted_names=[]
equations=[]

for i in range(len(lines)):
    names.append(lines[i].partition("=")[0])
    equations.append(lines[i].partition("=")[2])

copy_names=names.copy()
while len(copy_names)>0:
    big_i=0
    for i in range(len(copy_names)):
        if len(copy_names[i])>len(copy_names[big_i]):
            big_i=i
    sorted_names.append(copy_names.pop(big_i))

for i in range(len(sorted_names)):
    for j in range(len(equations)):
        if sorted_names[i] in equations[j]:
            equations[j]=equations[j].replace(sorted_names[i],"x("+str(names.index(sorted_names[i])+1)+",k)")

lines=[]
for i in range(len(names)):
    lines.append("V("+str(i+1)+")=\""+names[i]+"\"")
lines.append("")
for i in range(len(equations)):
    lines.append("y("+str(i+1)+",1)="+equations[i]+"! "+names[i])

open(output_file,"w").write("\n".join(lines))
