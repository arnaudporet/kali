// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "math"
import "strconv"
//#### DoS ###################################################################//
func DoS(n,maxS int,vals Vector) {
    var whole int
    var cardS float64
    var S Matrix
    cardS=math.Pow(float64(len(vals)),float64(n))
    if cardS>float64(maxS) {
        whole=int(Prompt("\nS cardinality ("+strconv.FormatFloat(cardS,'f',-1,64)+") > maxS ("+strconv.FormatInt(int64(maxS),10)+")\n\nCompute whole S [1], limit to maxS [0]? [0/1] ",Vector{0.0,1.0}))
    } else {
        whole=1
    }
    switch whole {
        case 0: S=GenArrangMat(vals,n,maxS).T()
        case 1: S=GenS(vals,n)
    }
    S.Save("S.csv")
    fmt.Println("\nINFO: state space (S) generated")
}
