PI3K// AKT
DNAdamage && !E2F1_lvl1 && !E2F1_lvl2// ATM_lvl1
(E2F1_lvl1 || E2F1_lvl2) && DNAdamage// ATM_lvl2
!CHEK1_2_lvl1 && !CHEK1_2_lvl2 && !RBL2 && (E2F1_lvl1 || E2F1_lvl2 || E2F3_lvl1 || E2F3_lvl2)// CDC25A
(ATM_lvl1 || ATM_lvl2) && !E2F1_lvl1 && !E2F1_lvl2// CHEK1_2_lvl1
(E2F1_lvl1 || E2F1_lvl2) && (ATM_lvl1 || ATM_lvl2)// CHEK1_2_lvl2
!RBL2 && !p21CIP && CDC25A && (E2F1_lvl1 || E2F1_lvl2 || E2F3_lvl1 || E2F3_lvl2)// CyclinA
(RAS || AKT) && !p16INK4a && !p21CIP// CyclinD1
!RBL2 && !p21CIP && CDC25A && (E2F1_lvl1 || E2F1_lvl2 || E2F3_lvl1 || E2F3_lvl2)// CyclinE1
!RB1 && !RBL2 && (((!CHEK1_2_lvl2 || !ATM_lvl2) && (RAS || E2F3_lvl1 || E2F3_lvl2)) || (CHEK1_2_lvl2 && ATM_lvl2 && !RAS && E2F3_lvl1))// E2F1_lvl1
!RBL2 && !RB1 && ATM_lvl2 && CHEK1_2_lvl2 && (RAS || E2F3_lvl2)// E2F1_lvl2
!RB1 && !CHEK1_2_lvl2 && RAS// E2F3_lvl1
!RB1 && CHEK1_2_lvl2 && RAS// E2F3_lvl2
(EGFRstimulus || SPRY) && !FGFR3 && !GRB2// EGFR
!EGFR && FGFR3stimulus && !GRB2// FGFR3
(FGFR3 && !GRB2 && !SPRY) || EGFR// GRB2
(TP53 || AKT) && !p14ARF && !ATM_lvl1 && !ATM_lvl2 && !RB1// MDM2
E2F1_lvl1 || E2F1_lvl2// p14ARF
GrowthInhibitors && !RB1// p16INK4a
!CyclinE1 && (GrowthInhibitors || TP53) && !AKT// p21CIP
GRB2 && RAS && !PTEN// PI3K
TP53// PTEN
EGFR || FGFR3 || GRB2// RAS
!CyclinD1 && !CyclinE1 && !p16INK4a && !CyclinA// RB1
!CyclinD1 && !CyclinE1// RBL2
RAS// SPRY
!MDM2 && (((ATM_lvl1 || ATM_lvl2) && (CHEK1_2_lvl1 || CHEK1_2_lvl2)) || E2F1_lvl2)// TP53
