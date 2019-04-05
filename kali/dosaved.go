// Copyright (C) 2013-2019 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "os"
    "strconv"
    "strings"
)
func DoSaved() {
    var toclear int
    toclear=GetInt(strings.Join([]string{
        "\nSaved:",
        "    S:        "+strconv.FormatBool(Exist("S.csv")),
        "    A_physio: "+strconv.FormatBool(Exist("A_physio.csv")),
        "    A_patho:  "+strconv.FormatBool(Exist("A_patho.csv")),
        "    A_versus: "+strconv.FormatBool(Exist("A_versus.csv")),
        "    Targ:     "+strconv.FormatBool(Exist("Targ.csv")),
        "    Moda:     "+strconv.FormatBool(Exist("Moda.csv")),
        "    B_therap: "+strconv.FormatBool(Exist("B_therap.txt")),
        "\nClear:",
        "    [1] csv files",
        "    [2] all",
        "    [0] nothing",
        "\nTo clear [0-2] ",
    },"\n"),Range(0,3))
    if toclear==1 {
        os.Remove("S.csv")
        os.Remove("A_physio.csv")
        os.Remove("A_patho.csv")
        os.Remove("A_versus.csv")
        os.Remove("Targ.csv")
        os.Remove("Moda.csv")
    } else if toclear==2 {
        os.Remove("S.csv")
        os.Remove("A_physio.csv")
        os.Remove("A_physio.txt")
        os.Remove("A_patho.csv")
        os.Remove("A_patho.txt")
        os.Remove("A_versus.csv")
        os.Remove("A_versus.txt")
        os.Remove("Targ.csv")
        os.Remove("Moda.csv")
        os.Remove("B_therap.txt")
    }
}
