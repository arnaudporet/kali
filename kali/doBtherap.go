// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
//#### DoBtherap #############################################################//
func DoBtherap(fpatho func(Matrix,int) Vector,nodes []string) {
    var S,Targ,Moda Matrix
    var Aphysio,Apatho,Aversus Aset
    var Btherap Bset
    if !Exist("S.csv") {
        fmt.Println("\nERROR: S must be generated")
    } else if !(Exist("Targ.csv") && Exist("Moda.csv")) {
        fmt.Println("\nERROR: Targ and Moda must be generated")
    } else if !(Exist("A_physio.csv") && Exist("A_patho.csv") && Exist("A_versus.csv")) {
        fmt.Println("\nERROR: A_physio, A_patho and A_versus must be computed")
    } else {
        Aversus.Load(2)
        if len(Aversus)==0 {
            fmt.Println("\nWARNING: no pathological attractors to remove (A_versus is empty)")
        } else {
            S.Load("S.csv")
            Targ.Load("Targ.csv")
            Moda.Load("Moda.csv")
            Aphysio.Load(0)
            Apatho.Load(1)
            Btherap.Compute(fpatho,S,Targ,Moda,Aphysio,Apatho,Aversus)
            Btherap.Report(nodes,Aphysio,Aversus)
        }
    }
}
