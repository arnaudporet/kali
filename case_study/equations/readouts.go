// Below are the equations of three output phenotypes of the model. They can be
// evaluated from the returned attractors once the run terminated.

//#### Boolean version #######################################################//

CyclinE1 || CyclinA// Proliferation
p21CIP || RB1 || RBL2// GrowthArrest
TP53 || E2F1_lvl2// Apoptosis

//#### functional version ####################################################//

max(CyclinE1,CyclinA)// Proliferation
max(p21CIP,RB1,RBL2)// GrowthArrest
max(TP53,E2F1_lvl2)// Apoptosis
