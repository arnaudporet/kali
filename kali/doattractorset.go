// Copyright (C) 2013-2019 Arnaud Poret
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
func DoAttractorSet(nodes []string,Physio,Patho func(Vector) Vector) {
    var (
        err error
        help,usage,success bool
        nSteps,maxTry int
        fileName,upd string
        args []string
        refSet,A AttractorSet
        f func(Vector) Vector
        flagSet *flag.FlagSet
    )
    fileName=filepath.Base(os.Args[0])
    flagSet=flag.NewFlagSet("",flag.ContinueOnError)
    flagSet.Usage=func() {}
    flagSet.BoolVar(&help,"help",false,"")
    flagSet.BoolVar(&help,"h",false,"")
    flagSet.BoolVar(&usage,"usage",false,"")
    flagSet.BoolVar(&usage,"u",false,"")
    flagSet.IntVar(&nSteps,"nstep",1000,"")
    flagSet.IntVar(&maxTry,"maxtry",10,"")
    flagSet.StringVar(&upd,"upd","async","")
    err=flagSet.Parse(os.Args[2:])
    if err!=nil {
        fmt.Println("Error: "+fileName+": attractor: "+err.Error())
    } else if help {
        fmt.Println(strings.Join([]string{
            "",
            "Compute an attractor set:",
            "    * A_physio: the attractor set of the physiological variant",
            "    * A_patho: the attractor set of the pathological variant",
            "",
            "Usage: "+fileName+" attractor [options] <setting>",
            "",
            "Positional arguments:",
            "    * <setting>: compute the physiological attractor set (setting=physio) or the",
            "                 pathological attractor set (setting=patho)",
            "",
            "Options:",
            "    * -nstep: the number of steps performed during a random walk when searching",
            "              for an attractor (default: 1 000)",
            "    * -maxtry: the maximum number of random walks performed when searching for",
            "               an attractor (default: 10)",
            "    * -upd: the updating method to use (default: async)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print this help",
            "",
            "Output files:",
            "    if computing the physiological attractor set:",
            "        * A_physio.csv",
            "        * A_physio.txt",
            "    if computing the pathological attractor set:",
            "        * A_patho.csv",
            "        * A_patho.txt",
            "",
            "Cautions:",
            "    * if A_physio.csv, A_physio.txt, A_patho.csv, A_patho.txt already exist then",
            "      they are overwritten",
            "    * if A_physio.csv, A_physio.txt, A_patho.csv, A_patho.txt are renamed, moved",
            "      or deleted then they can not be loaded when required",
            "",
            "nstep:",
            "    * only relevant in the asynchronous case",
            "    * recommended to be at least 1 000",
            "    * when searching for an attractor according to an asynchronous updating, a",
            "      long random walk is performed in order to reach an attractor with high",
            "      probability (this candidate attractor is then subjected to validation)",
            "    * the smallest is nstep the smallest is the probability to reach an",
            "      attractor: this will cause kali to run for a too long time",
            "    * if nstep is too big then kali will also run for a too long time",
            "    * a compromise could be nstep in [1 000;10 000], depending on the size of",
            "      the model (i.e. the size of the state space)",
            "    * the bigger is the state space the bigger should be nstep",
            "",
            "maxtry:",
            "    * only relevant in the asynchronous case",
            "    * if nstep is too small regarding the size of the state space, it is",
            "      possible that no attractors can be reached from a given start state (i.e.",
            "      where starts a random walk in the state space)",
            "    * to prevent looping indefinitely, a maximum of maxtry random walks are",
            "      performed from each start state",
            "",
            "upd:",
            "    * async: an asynchronous updating is used, meaning that one randomly",
            "             selected variable is updated at each iteration of the simulation",
            "    * sync: a synchronous updating is used, meaning that all the variables are",
            "            updated simultaneously at each iteration of the simulation",
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
            "Usage: "+fileName+" attractor [options] <setting>",
            "",
            "Positional arguments:",
            "    * <setting>: compute the physiological attractor set (setting=physio) or the",
            "                 pathological attractor set (setting=patho)",
            "",
            "Options:",
            "    * -nstep: the number of steps performed during a random walk when searching",
            "              for an attractor (default: 1 000)",
            "    * -maxtry: the maximum number of random walks performed when searching for",
            "               an attractor (default: 10)",
            "    * -upd: the updating method to use (default: async)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print help",
            "",
        },"\n"))
    } else if len(flagSet.Args())!=1 {
        fmt.Println("Error: "+fileName+": attractor: wrong number of positional arguments, expecting: <setting>")
    } else if nSteps<1 {
        fmt.Println("Error: "+fileName+": attractor: nstep must be a positive integer")
    } else if maxTry<1 {
        fmt.Println("Error: "+fileName+": attractor: maxtry must be a positive integer")
    } else if (upd!="sync") && (upd!="async") {
        fmt.Println("Error: "+fileName+": attractor: unknown updating, expecting one of: sync, async")
    } else if (flagSet.Arg(0)!="physio") && (flagSet.Arg(0)!="patho") {
        fmt.Println("Error: "+fileName+": attractor: unknown setting, expecting one of: physio, patho")
    } else {
        args=flagSet.Args()
        if !Exist("S.csv") {
            fmt.Println("Error: "+fileName+": attractor: no start states found (S.csv)")
        } else if (args[0]=="patho") && !Exist("A_physio.csv") {
            fmt.Println("Error: "+fileName+": attractor: no physiological attractor set found (A_physio.csv)")
        } else {
            if args[0]=="physio" {
                f=Physio
                refSet=AttractorSet{}
            } else if args[0]=="patho" {
                f=Patho
                refSet=LoadAttractorSet("physio")
            }
            A,success=ComputeAttractorSet(LoadMat("S.csv"),f,Bullet{},nSteps,-1,maxTry,upd,args[0],refSet)
            if success {
                A.Report(nodes,args[0])
            } else {
                fmt.Println("Warning: "+fileName+": attractor: unable to find attractors, try increasing nstep and/or maxtry")
            }
        }
    }
}
