// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "strconv"
import "strings"
//#### DoParams ##############################################################//
func DoParams(ntarg,maxtarg,maxmoda,maxS *int,n int) {
    var tochange int
    tochange=-1
    for tochange!=0 {
        tochange=int(Prompt(strings.Join([]string{
            "\nChange:",
            "    [1] ntarg   ("+strconv.FormatInt(int64(*ntarg),10)+")",
            "    [2] maxtarg ("+strconv.FormatInt(int64(*maxtarg),10)+")",
            "    [3] maxmoda ("+strconv.FormatInt(int64(*maxmoda),10)+")",
            "    [4] maxS    ("+strconv.FormatInt(int64(*maxS),10)+")",
            "    [0] done",
            "\nTo change [0-4] ",
        },"\n"),ItoV(Range(0,5))))
        switch tochange {
            case 1:
                (*ntarg)=int(Prompt("\nntarg=",ItoV(Range(1,n+1))))
                fmt.Println("\nWARNING: you should regenerate Targ and Moda, if any")
            case 2:
                (*maxtarg)=int(Prompt("\nmaxtarg=",Vector{}))
                fmt.Println("\nWARNING: you should regenerate Targ and Moda, if any")
            case 3:
                (*maxmoda)=int(Prompt("\nmaxmoda=",Vector{}))
                fmt.Println("\nWARNING: you should regenerate Targ and Moda, if any")
            case 4:
                (*maxS)=int(Prompt("\nmaxS=",Vector{}))
                fmt.Println("\nWARNING: you should regenerate S, if any")
            case 0:
                fmt.Println("\nINFO: OK!")
        }
    }
}
