// Copyright (C) 2013-2021 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

package kali
import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)
func DoVersus(nodes []string) {
    var (
        err error
        help,usage,quiet bool
        fileName string
        flagSet *flag.FlagSet
    )
    fileName=filepath.Base(os.Args[0])
    flagSet=flag.NewFlagSet("",flag.ContinueOnError)
    flagSet.Usage=func() {}
    flagSet.BoolVar(&help,"help",false,"")
    flagSet.BoolVar(&help,"h",false,"")
    flagSet.BoolVar(&usage,"usage",false,"")
    flagSet.BoolVar(&usage,"u",false,"")
    flagSet.BoolVar(&quiet,"quiet",false,"")
    flagSet.BoolVar(&quiet,"q",false,"")
    err=flagSet.Parse(os.Args[2:])
    if err!=nil {
        fmt.Println("Error: "+fileName+": versus: "+err.Error())
    } else if help {
        fmt.Println(strings.Join([]string{
            "",
            "Get the set of the pathological attractors.",
            "",
            "Usage: "+fileName+" versus [options]",
            "",
            "Option:",
            "    * -q/-quiet: do not print results (but still save them to files)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print this help",
            "",
            "Output files:",
            "    * A_versus.csv",
            "    * A_versus.txt",
            "",
            "Cautions:",
            "    * A_versus is not the attractor set of the pathological variant",
            "      (i.e. A_patho): it is the set containing the pathological attractors,",
            "      namely the attractors specific to A_patho (A_patho can also contain",
            "      physiological attractors)",
            "    * if A_versus.csv, A_versus.txt already exist then they are overwritten",
            "    * if A_versus.csv, A_versus.txt are renamed, moved or deleted then they can",
            "      not be loaded when required",
            "",
            "The basins of the attractors are expressed in percents of the state space (note",
            "that it is an estimation).",
            "",
            "The csv files are for kali while the txt files are for you.",
            "",
        },"\n"))
    } else if usage {
        fmt.Println(strings.Join([]string{
            "",
            "Usage: "+fileName+" versus [options]",
            "",
            "Option:",
            "    * -q/-quiet: do not print results (but still save them to files)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print help",
            "",
        },"\n"))
    } else if len(flagSet.Args())!=0 {
        fmt.Println("Error: "+fileName+": versus: expecting no positional arguments")
    } else {
        if !Exist("A_patho.csv") {
            fmt.Println("Error: "+fileName+": versus: no pathological attractor set found (A_patho.csv)")
        } else {
            LoadAttractorSet("patho").GetVersus().Report(nodes,"versus",quiet)
        }
    }
}
