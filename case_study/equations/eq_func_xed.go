"AKT",
"ATM_lvl1",
"ATM_lvl2",
"CDC25A",
"CHEK1_2_lvl1",
"CHEK1_2_lvl2",
"CyclinA",
"CyclinD1",
"CyclinE1",
"E2F1_lvl1",
"E2F1_lvl2",
"E2F3_lvl1",
"E2F3_lvl2",
"EGFR",
"FGFR3",
"GRB2",
"MDM2",
"p14ARF",
"p16INK4a",
"p21CIP",
"PI3K",
"PTEN",
"RAS",
"RB1",
"RBL2",
"SPRY",
"TP53",

x[20],// AKT
kali.Min(DNAdamage,1-x[9],1-x[10]),// ATM_lvl1
kali.Min(kali.Max(x[9],x[10]),DNAdamage),// ATM_lvl2
kali.Min(1-x[4],1-x[5],1-x[24],kali.Max(x[9],x[10],x[11],x[12])),// CDC25A
kali.Min(kali.Max(x[1],x[2]),1-x[9],1-x[10]),// CHEK1_2_lvl1
kali.Min(kali.Max(x[9],x[10]),kali.Max(x[1],x[2])),// CHEK1_2_lvl2
kali.Min(1-x[24],1-x[19],x[3],kali.Max(x[9],x[10],x[11],x[12])),// CyclinA
kali.Min(kali.Max(x[22],x[0]),1-x[18],1-x[19]),// CyclinD1
kali.Min(1-x[24],1-x[19],x[3],kali.Max(x[9],x[10],x[11],x[12])),// CyclinE1
kali.Min(1-x[23],1-x[24],kali.Max(kali.Min(kali.Max(1-x[5],1-x[2]),kali.Max(x[22],x[11],x[12])),kali.Min(x[5],x[2],1-x[22],x[11]))),// E2F1_lvl1
kali.Min(1-x[24],1-x[23],x[2],x[5],kali.Max(x[22],x[12])),// E2F1_lvl2
kali.Min(1-x[23],1-x[5],x[22]),// E2F3_lvl1
kali.Min(1-x[23],x[5],x[22]),// E2F3_lvl2
kali.Min(kali.Max(EGFRstimulus,x[25]),1-x[14],1-x[15]),// EGFR
kali.Min(1-x[13],FGFR3stimulus,1-x[15]),// FGFR3
kali.Max(kali.Min(x[14],1-x[15],1-x[25]),x[13]),// GRB2
kali.Min(kali.Max(x[26],x[0]),1-x[17],1-x[1],1-x[2],1-x[23]),// MDM2
kali.Max(x[9],x[10]),// p14ARF
kali.Min(GrowthInhibitors,1-x[23]),// p16INK4a
kali.Min(1-x[8],kali.Max(GrowthInhibitors,x[26]),1-x[0]),// p21CIP
kali.Min(x[15],x[22],1-x[21]),// PI3K
x[26],// PTEN
kali.Max(x[13],x[14],x[15]),// RAS
kali.Min(1-x[7],1-x[8],1-x[18],1-x[6]),// RB1
kali.Min(1-x[7],1-x[8]),// RBL2
x[22],// SPRY
kali.Min(1-x[16],kali.Max(kali.Min(kali.Max(x[1],x[2]),kali.Max(x[4],x[5])),x[10])),// TP53
