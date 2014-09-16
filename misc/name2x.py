input_file=input("input file: ")
output_file=input("output file: ")

lines=open(input_file,"r").read().splitlines()
names=[]
equations=[]

for i in range(len(lines)):
    names.append(lines[i].partition("=")[0])
    equations.append(lines[i].partition("=")[2])

for i in range(len(equations)):
    for j in range(len(names)):
        if names[j]+" " in equations[i]:
            equations[i]=equations[i].replace(names[j]+" ","x("+str(j+1)+",k) ")
        if names[j]+")" in equations[i]:
            equations[i]=equations[i].replace(names[j]+")","x("+str(j+1)+",k))")
        if names[j]+";" in equations[i]:
            equations[i]=equations[i].replace(names[j]+";","x("+str(j+1)+",k);")

lines=[]
for i in range(len(names)):
    lines.append("V("+str(i+1)+")=\""+names[i]+"\"")
for i in range(len(equations)):
    equations[i]=equations[i].replace(";","")
    lines.append("y("+str(i+1)+",1)="+equations[i]+"! "+names[i])

open(output_file,"w").write("\n".join(lines))
