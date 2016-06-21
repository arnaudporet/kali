// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "strconv"
    "strings"
)
func DoParameters(ntarg,maxtarg,maxmoda,maxS,threshold *int,n int) {
    var tochange int
    for {
        tochange=GetInt(strings.Join([]string{
            "\nWhat to change:",
            "    [1] ntarg        ("+strconv.FormatInt(int64(*ntarg),10)+")",
            "    [2] maxtarg      ("+strconv.FormatInt(int64(*maxtarg),10)+")",
            "    [3] maxmoda      ("+strconv.FormatInt(int64(*maxmoda),10)+")",
            "    [4] maxS         ("+strconv.FormatInt(int64(*maxS),10)+")",
            "    [5] threshold    ("+strconv.FormatInt(int64(*threshold),10)+")",
            "    [0] done",
            "\nTo change [0-5] ",
        },"\n"),Range(0,6))
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
            fmt.Println("\nWARNING: you should regenerate S, then all recompute")
        } else if tochange==5 {
            *threshold=GetInt("\nthreshold=",Range(1,101))
            fmt.Println("\nWARNING: you should recompute B_therap")
        } else if tochange==0 {
            if *maxtarg<1 || *maxmoda<1 || *maxS<1 {
                fmt.Println("\nERROR: maxtarg<1 or maxmoda<1 or maxS<1 (they must be strictly positive integers)")
            } else {
                break
            }
        }
    }
    fmt.Println("\nINFO: OK!")
}
