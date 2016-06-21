// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "math"
    "strconv"
)
func DoStateSpace(n,maxS int,vals Vector) {
    var (
        wholeS int
        cardS float64
    )
    cardS=math.Pow(float64(len(vals)),float64(n))
    if cardS>float64(maxS) {
        wholeS=GetInt("\nS cardinality ("+strconv.FormatFloat(cardS,'f',-1,64)+") > maxS ("+strconv.FormatInt(int64(maxS),10)+")\n\nLimit to maxS [0] or compute whole S [1]? [0/1] ",[]int{0,1})
    } else {
        wholeS=1
    }
    if wholeS==0 {
        vals.GenArrangMat(n,maxS).Save("S.csv")
    } else if wholeS==1 {
        vals.GenS(n).Save("S.csv")
    }
    fmt.Println("\nINFO: S generated")
}
