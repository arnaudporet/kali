// Copyright (C) 2013-2018 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
func DoVersus(nodes []string) {
    if !Exist("A_patho.csv") {
        fmt.Println("\nERROR: A_patho must be computed")
    } else {
        LoadAttractorSet(1).GetVersus().Report(nodes,2)
    }
}
