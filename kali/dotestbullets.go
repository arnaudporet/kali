// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
func DoTestBullets(ntarg,maxtarg,maxmoda,n int,vals Vector) {
    IntToVector(Range(0,n)).GenCombiMat(ntarg,maxtarg).Save("Targ.csv")
    vals.GenArrangMat(ntarg,maxmoda).Save("Moda.csv")
    fmt.Println("\nINFO: Targ and Moda generated")
}
