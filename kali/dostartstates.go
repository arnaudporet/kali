// Copyright (C) 2013-2019 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.
package kali
import (
    "flag"
    "fmt"
    "math"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)
func DoStartStates(vals Vector,nodes []string) {
    var (
        err error
        help,usage bool
        maxS,whole int
        card float64
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
    flagSet.IntVar(&maxS,"maxS",1000,"")
    err=flagSet.Parse(os.Args[2:])
    if err!=nil {
        fmt.Println("Error: "+fileName+": start: "+err.Error())
    } else if help {
        fmt.Println(strings.Join([]string{
            "",
            "Generate the set of the start states.",
            "",
            "The start states are the initial states to use when computing an attractor set.",
            "",
            "Usage: "+fileName+" start [options]",
            "",
            "Options:",
            "    * -maxS: the maximal number of start states (default: 1 000)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print this help",
            "",
            "Output file:",
            "    * S.csv",
            "",
            "Cautions:",
            "    * if S.csv already exists then it is overwritten",
            "    * if S.csv is renamed, moved or deleted then it can not be loaded when",
            "      required",
            "",
            "maxS:",
            "    * recommended to be at least 1 000",
            "    * is the maximum number of initial states to use when computing an attractor",
            "      set",
            "    * if it exceeds its maximal possible value then it is automatically",
            "      decreased to its maximal possible value",
            "    * the smallest is maxS the smallest is the probability to find all the",
            "      attractors",
            "    * if maxS is too big then kali will run for a too long time",
            "    * a compromise could be maxS in [1 000;100 000], depending on the size of",
            "      the model (i.e. the size of the state space)",
            "    * the bigger is the state space the bigger should be maxS",
            "",
        },"\n"))
    } else if usage {
        fmt.Println(strings.Join([]string{
            "",
            "Usage: "+fileName+" start [options]",
            "",
            "Options:",
            "    * -maxS: the maximal number of start states (default: 1 000)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print help",
            "",
        },"\n"))
    } else if len(flagSet.Args())!=0 {
        fmt.Println("Error: "+fileName+": start: expecting no positional arguments")
    } else if maxS<1 {
        fmt.Println("Error: "+fileName+": start: maxS must be a positive integer")
    } else {
        card=math.Pow(float64(len(vals)),float64(len(nodes)))
        if card>float64(maxS) {
            whole=GetInt("state space cardinality ("+strconv.FormatFloat(card,'f',-1,64)+") > maxS ("+strconv.FormatInt(int64(maxS),10)+")\nlimit to maxS [0] or use all the state space [1]? [0/1] ",[]int{0,1})
        } else {
            whole=1
        }
        if whole==0 {
            vals.Arrangs(len(nodes),maxS).Save("S.csv")
        } else if whole==1 {
            vals.Space(len(nodes)).Save("S.csv")
        }
    }
}
