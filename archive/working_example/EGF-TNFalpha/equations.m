EGF;#EGF
TNFalpha;#TNFalpha
EGF;#EGFR
TNFalpha;#TNFR
and(EGFR,not(PP));#SOS
EGFR;#PI3K
TNFR;#TRAF2
SOS;#Ras
PI3K;#Akt
TRAF2;#ASK1
TRAF2;#MKKK7
or(ERK,PP);#PP
Ras;#Raf1
not(Akt);#GSK3
MKKK7;#NIK
Raf1;#MEK
Ras;#MKKK1
NIK;#IKK
MEK;#ERK
or(MKKK1,ASK1);#MKK7
or(not(IKK),ex);#IkappaB
MKK7;#JNK
and(MKKK1,MKKK7);#MKK4
not(IkappaB);#NFkappaB
JNK;#cJun
MKK4;#p38
NFkappaB;#ex
cJun;#AP1
