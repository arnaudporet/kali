// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "strconv"
import "strings"
//#### DoSaved ###############################################################//
func DoSaved() {
    fmt.Println(strings.Join([]string{
        "\nAlready saved:",
        "    S:           "+strconv.FormatBool(Exist("S.csv")),
        "    A_physio:    "+strconv.FormatBool(Exist("A_physio.csv")),
        "    A_patho:     "+strconv.FormatBool(Exist("A_patho.csv")),
        "    A_versus:    "+strconv.FormatBool(Exist("A_versus.csv")),
        "    Targ:        "+strconv.FormatBool(Exist("Targ.csv")),
        "    Moda:        "+strconv.FormatBool(Exist("Moda.csv")),
        "    B_therap:    "+strconv.FormatBool(Exist("B_therap.txt")),
    },"\n"))
}
