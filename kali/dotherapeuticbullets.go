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
func DoTherapeuticBullets(nodes []string,Patho func(Vector) Vector) {
    var (
        err error
        help,usage,success bool
        nSteps,maxTry int
        th float64
        fileName,upd string
        Aphysio,Apatho,Aversus AttractorSet
        Btherap BulletSet
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
    flagSet.Float64Var(&th,"th",5,"")
    err=flagSet.Parse(os.Args[2:])
    if err!=nil {
        fmt.Println("Error: "+fileName+": target: "+err.Error())
    } else if help {
        fmt.Println(strings.Join([]string{
            "",
            "Compute a set of therapeutic bullets.",
            "",
            "Usage: "+fileName+" target [options]",
            "",
            "Options:",
            "    * -nstep: the number of steps performed during a random walk when searching",
            "              for an attractor (default: 1 000)",
            "    * -maxtry: the maximum number of random walks performed when searching for",
            "               an attractor (default: 10)",
            "    * -upd: the updating method to use (default: async)",
            "    * -th: the threshold for a bullet to be considered therapeutic (default: 5)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print this help",
            "",
            "Output file:",
            "    * B_therap.txt",
            "",
            "Cautions:",
            "    * nstep, maxtry and upd should be the same as the attractor set",
            "      computation step",
            "    * if B_therap.txt already exists then it is overwritten",
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
            "th:",
            "    * is expressed in percents of the state space",
            "    * the goal of therapeutic bullets is to increase the physiological basins",
            "      (i.e. to increase the union of the basins of the physiological attractors)",
            "    * to be therapeutic, this increase must be superior or equal to th (in",
            "      percents of the state space)",
            "",
            "Therapeutic bullets are reported as follows:",
            "    Bullet: x1[y1] x2[y2] ...",
            "    Gain: U1 --> U2",
            "    Physiological basins:",
            "        a_physio1: b_physio1",
            "        a_physio2: b_physio2",
            "        ...",
            "    Pathological basins:",
            "        a_patho1: b_patho1",
            "        a_patho2: b_patho2",
            "        ...",
            "where:",
            "    * x1[y1] x2[y2] ... means that the node x1 has to be set to the value y1,",
            "      the node x2 to the value y2, and so on",
            "    * U1 is the union of the physiological basins in the state space of the",
            "      pathological variant (in percents of it)",
            "    * U2 is the union of the physiological basins in the state space of the",
            "      pathological variant subjected to the effect of the bullet (in percents of",
            "      it)",
            "    * therefore, the gain of the bullet is the increase from U1 to U2 (this is a",
            "      condition for the bullet to be considered therapeutic: U2-U1>=th)",
            "    * the basin of the physiological and pathological attractors (b_physio and",
            "      b_patho respectively) in the state space of the pathological variant",
            "      subjected to the effect of the bullet are expressed in percents of it",
            "",
        },"\n"))
    } else if usage {
        fmt.Println(strings.Join([]string{
            "",
            "Usage: "+fileName+" target [options]",
            "",
            "Options:",
            "    * -nstep: the number of steps performed during a random walk when searching",
            "              for an attractor (default: 1 000)",
            "    * -maxtry: the maximum number of random walks performed when searching for",
            "               an attractor (default: 10)",
            "    * -upd: the updating method to use (default: async)",
            "    * -th: the threshold for a bullet to be considered therapeutic (default: 5)",
            "    * -u/-usage: print usage only",
            "    * -h/-help: print help",
            "",
        },"\n"))
    } else if len(flagSet.Args())!=0 {
        fmt.Println("Error: "+fileName+": target: expecting no positional arguments")
    } else if nSteps<1 {
        fmt.Println("Error: "+fileName+": target: nstep must be a positive integer")
    } else if maxTry<1 {
        fmt.Println("Error: "+fileName+": target: maxtry must be a positive integer")
    } else if (upd!="sync") && (upd!="async") {
        fmt.Println("Error: "+fileName+": target: unknown updating, expecting one of: sync, async")
    } else if (th<0) || (th>100) {
        fmt.Println("Error: "+fileName+": target: th must be a real number in [0;100]")
    } else {
        if !Exist("S.csv") {
            fmt.Println("Error: "+fileName+": target: no start states found (S.csv)")
        } else if !Exist("A_physio.csv") {
            fmt.Println("Error: "+fileName+": target: no physiological attractor set found (A_physio.csv)")
        } else if !Exist("A_patho.csv") {
            fmt.Println("Error: "+fileName+": target: no pathological attractor set found (A_patho.csv)")
        } else if !Exist("A_versus.csv") {
            fmt.Println("Error: "+fileName+": target: no pathological attractors found (A_versus.csv)")
        } else if !Exist("Targ.csv") {
            fmt.Println("Error: "+fileName+": target: no target node combinations found (Targ.csv)")
        } else if !Exist("Moda.csv") {
            fmt.Println("Error: "+fileName+": target: no modality arrangements found (Moda.csv)")
        } else {
            Aversus=LoadAttractorSet("versus")
            if len(Aversus)==0 {
                fmt.Println("Warning: "+fileName+": target: there are no pathological attractors (A_versus is empty)")
            } else {
                Aphysio=LoadAttractorSet("physio")
                Apatho=LoadAttractorSet("patho")
                Btherap,success=ComputeTherapeuticBullets(LoadMat("S.csv"),LoadMat("Targ.csv"),LoadMat("Moda.csv"),Patho,nSteps,maxTry,th,Aphysio,Apatho,Aversus,upd)
                if success {
                    Btherap.Report(nodes,Align(Aphysio.GetNames()," "),Align(Aversus.GetNames()," "))
                } else {
                    fmt.Println("Warning: "+fileName+": target: unable to find therapeutic bullets, try increasing nstep and/or maxtry")
                }
            }
        }
    }
}
