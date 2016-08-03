// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "strconv"
    "strings"
)
func DoParameters(ntarg,maxtarg,maxmoda,maxS,kmax,threshold,sync *int,n int) {
    var tochange int
    for {
        tochange=GetInt(strings.Join([]string{
            "\nWhat to change:",
            "    [1] ntarg        ("+strconv.FormatInt(int64(*ntarg),10)+")",
            "    [2] maxtarg      ("+strconv.FormatInt(int64(*maxtarg),10)+")",
            "    [3] maxmoda      ("+strconv.FormatInt(int64(*maxmoda),10)+")",
            "    [4] maxS         ("+strconv.FormatInt(int64(*maxS),10)+")",
            "    [5] kmax         ("+strconv.FormatInt(int64(*kmax),10)+")",
            "    [6] threshold    ("+strconv.FormatInt(int64(*threshold),10)+")",
            "    [7] sync         ("+strconv.FormatInt(int64(*sync),10)+")",
            "    [0] done",
            "\nTo change [0-7] ",
        },"\n"),Range(0,8))
        if tochange==1 {
            *ntarg=GetInt("\nntarg=",Range(1,n+1))
            fmt.Println("\nWARNING: you should regenerate Targ and Moda, then recompute B_therap")
        } else if tochange==2 {
            *maxtarg=GetInt("\nmaxtarg=",[]int{})
            fmt.Println("\nWARNING: you should regenerate Targ and Moda, then recompute B_therap")
        } else if tochange==3 {
            *maxmoda=GetInt("\nmaxmoda=",[]int{})
            fmt.Println("\nWARNING: you should regenerate Targ and Moda, then recompute B_therap")
        } else if tochange==4 {
            *maxS=GetInt("\nmaxS=",[]int{})
            fmt.Println("\nWARNING: you should regenerate S, then recompute A_physio, A_patho, A_versus and B_therap")
        } else if tochange==5 {
            *kmax=GetInt("\nkmax=",[]int{})
            fmt.Println("\nWARNING: you should recompute A_physio and A_patho, then recompute A_versus and B_therap")
        } else if tochange==6 {
            *threshold=GetInt("\nthreshold=",Range(0,101))
            fmt.Println("\nWARNING: you should recompute B_therap")
        } else if tochange==7 {
            *sync=GetInt("\nsync=",[]int{0,1})
            fmt.Println("\nWARNING: you should recompute A_physio and A_patho, then recompute A_versus and B_therap")
        } else if tochange==0 {
            if *maxtarg<1 || *maxmoda<1 || *maxS<1 || *kmax<1 {
                fmt.Println("\nERROR: maxtarg<1 or maxmoda<1 or maxS<1 or kmax<1 (they must be strictly positive integers)")
            } else {
                break
            }
        }
    }
}
