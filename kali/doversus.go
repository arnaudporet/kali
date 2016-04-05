// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
//#### DoVersus ##############################################################//
func DoVersus(nodes []string) {
    var Apatho,Aversus Aset
    if !Exist("A_patho.csv") {
        fmt.Println("\nERROR: A_patho must be computed")
    } else {
        Apatho.Load(1)
        Aversus.Versus(Apatho)
        Aversus.Report(2,nodes)
    }
}
