// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
//#### DoBtest ###############################################################//
func DoBtest(n,ntarg,maxtarg,maxmoda int,vals Vector) {
    var Targ,Moda Matrix
    Targ=GenCombiMat(ItoV(Range(0,n)),ntarg,maxtarg)
    Targ.Save("Targ.csv")
    Moda=GenArrangMat(vals,ntarg,maxmoda)
    Moda.Save("Moda.csv")
    fmt.Println("\nINFO: bullets to test (Targ and Moda) generated")
}
