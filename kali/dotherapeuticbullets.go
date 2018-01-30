// Copyright (C) 2013-2018 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
func DoTherapeuticBullets(fpatho func(Vector) Vector,kmax,threshold,sync int,nodes []string) {
    var Aphysio,Apatho,Aversus AttractorSet
    if !Exist("S.csv") {
        fmt.Println("\nERROR: S must be generated")
    } else if !Exist("Targ.csv") || !Exist("Moda.csv") {
        fmt.Println("\nERROR: Targ and Moda must be generated")
    } else if !Exist("A_physio.csv") || !Exist("A_patho.csv") || !Exist("A_versus.csv") {
        fmt.Println("\nERROR: A_physio, A_patho and A_versus must be computed")
    } else {
        Aversus=LoadAttractorSet(2)
        if len(Aversus)==0 {
            fmt.Println("\nWARNING: no pathological attractors to remove (A_versus is empty)")
        } else {
            Aphysio=LoadAttractorSet(0)
            Apatho=LoadAttractorSet(1)
            ComputeTherapeuticBullets(fpatho,LoadMat("S.csv"),LoadMat("Targ.csv"),LoadMat("Moda.csv"),kmax,threshold,sync,Aphysio,Apatho,Aversus).Report(nodes,Align(Aphysio.GetNames()," "),Align(Aversus.GetNames()," "))
        }
    }
}
