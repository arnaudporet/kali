// Copyright (C) 2013-2020 Arnaud Poret
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
func DoTestBullets(vals Vector,nodes []string) {
    var (
        err error
        help,usage bool
        nTarg,maxTarg,maxModa int
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
    flagSet.IntVar(&nTarg,"ntarg",1,"")
    flagSet.IntVar(&maxTarg,"maxtarg",100,"")
    flagSet.IntVar(&maxModa,"maxmoda",100,"")
    err=flagSet.Parse(os.Args[2:])
    if err!=nil {
        fmt.Println("Error: "+fileName+": bullet: "+err.Error())
    } else if help {
        fmt.Println(strings.Join([]string{
            "",
            "Generate the set of the bullets to test:",
            "    * Targ: the set of the target node combinations to test",
            "    * Moda: the set of the modality arrangements (i.e. the state modifications)",
            "            to apply on each of the target node combinations",
            "",
            "The modalities are the perturbations to apply on the target nodes, typically",
            "inhibitions or stimulations.",
            "",
            "A bullet is a couple made of:",
            "    * one combination without repetition of ntarg target nodes",
            "    * one arrangement with repetition of ntarg modalities",
            "    * ex: ((node1,node2),(moda1,moda2))",
            "          where the modality moda1 is for the node node1 and the modality moda2",
            "          for the node node2, and ntarg=2",
            "",
            "Usage: "+fileName+" bullet [options]",
            "",
            "Options:",
            "    * -ntarg: the number of target nodes per bullet (default: 1)",
            "    * -maxtarg: the maximum number of target node combinations to test",
            "                (default: 100)",
            "    * -maxmoda: the maximum number of modality arrangements to test for each",
            "                target node combination (default: 100)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print this help",
            "",
            "Output files:",
            "    * Targ.csv",
            "    * Moda.csv",
            "",
            "Cautions:",
            "    * if Targ.csv, Moda.csv already exist then they are overwritten",
            "    * if Targ.csv, Moda.csv are renamed, moved or deleted then they can not be",
            "      loaded when required",
            "",
            "maxmoda is the maximum number of modality arrangements to test for each target",
            "node combination: there are maximum maxtarg*maxmoda bullets to test.",
            "",
            "If maxtarg and/or maxmoda exceeds its maximal possible value then it is",
            "automatically decreased to its maximal possible value.",
            "",
        },"\n"))
    } else if usage {
        fmt.Println(strings.Join([]string{
            "",
            "Usage: "+fileName+" bullet [options]",
            "",
            "Options:",
            "    * -ntarg: the number of target nodes per bullet (default: 1)",
            "    * -maxtarg: the maximum number of target node combinations to test",
            "                (default: 100)",
            "    * -maxmoda: the maximum number of modality arrangements to test for each",
            "                target node combination (default: 100)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print help",
            "",
        },"\n"))
    } else if len(flagSet.Args())!=0 {
        fmt.Println("Error: "+fileName+": bullet: expecting no positional arguments")
    } else if (nTarg<1) || (nTarg>len(nodes)) {
        fmt.Println("Error: "+fileName+": bullet: ntarg must be an integer in [1;number of nodes]")
    } else if maxTarg<1 {
        fmt.Println("Error: "+fileName+": bullet: maxtarg must be a positive integer")
    } else if maxModa<1 {
        fmt.Println("Error: "+fileName+": bullet: maxmoda must be a positive integer")
    } else {
        IntToVect(Range(0,len(nodes))).Combis(nTarg,maxTarg).Save("Targ.csv")
        vals.Arrangs(nTarg,maxModa).Save("Moda.csv")
    }
}
