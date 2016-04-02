// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "math"
import "math/rand"
import "strconv"
import "strings"
import "time"
//#### DoTheJob ##############################################################//
func DoTheJob(fphysio,fpatho func(Matrix,int) Vector,maxtarg,maxmoda,maxD int,nodes []string,vals Vector) {
    var todo,wholeS,setting,rmin,rmax,tochange,gen,save int
    var D Matrix
    var nullb Bullet
    var Aphysio,Apatho,Aversus,nullset Aset
    var Btherap Bset
    rand.Seed(int64(time.Now().Nanosecond()))
    todo=-1
    for todo!=0 {// TODO case 0 does not break it...
        todo=int(Prompt(strings.Join([]string{
            "\nWhat to do:",
            "    [1] generate or load the state space",
            "    [2] compute an attractor set (A_physio or A_patho)",
            "    [3] compute the pathological attractors (A_versus)",
            "    [4] compute therapeutic bullets (B_therap)",
            "    [5] change parameter values (maxtarg, maxmoda, maxD)",
            "    [6] check what is already saved (state space, A_physio, A_patho, A_versus, B_therap)",
            "    [7] help",
            "    [8] license",
            "    [0] quit",
            "\nTo do [0-8] ",
        },"\n"),ToV(Range(0,9))))
        switch todo {
            case 1:
                gen=int(Prompt("\nState space: load [0] or generate [1]? [0/1] ",Vector{0.0,1.0}))
                switch gen {
                    case 0:
                        if !Exist("state_space.csv") {
                            fmt.Println("\nERROR: nothing to load")
                        } else {
                            D.Load("state_space.csv")
                            fmt.Println("\nINFO: state space loaded")
                        }
                    case 1:
                        wholeS=int(Prompt("\nState space cardinality: "+strconv.FormatFloat(math.Pow(float64(len(vals)),float64(len(nodes))),'f',-1,64)+"\n\nCompute the whole state space? [0/1] ",Vector{0.0,1.0}))
                        switch wholeS {
                            case 0: D=GenArrangMat(vals,len(nodes),maxD).T()
                            case 1: D=GenS(vals,len(nodes))
                        }
                        fmt.Println("\nINFO: state space generated")
                        save=int(Prompt("\nSave? (optional) [0/1] ",Vector{0.0,1.0}))
                        if save==1 {
                            D.Save("state_space.csv")
                            fmt.Println("\nINFO: state space saved as state_space.csv")
                        }
                }
            case 2:
                if len(D)==0 {
                    fmt.Println("\nERROR: the state space must be generated or loaded")
                } else {
                    setting=int(Prompt("\nSetting: physiological [0] or pathological [1]? [0/1] ",Vector{0.0,1.0}))
                    switch setting {
                        case 0:
                            Aphysio.Compute(fphysio,D,nullb,nullset,0)
                            Aphysio.Report(0,nodes)
                        case 1:
                            if !Exist("A_physio.csv") {
                                fmt.Println("\nERROR: the physiological attractor set (A_physio) must be computed and saved")
                            } else {
                                Aphysio.Load(0)
                                Apatho.Compute(fpatho,D,nullb,Aphysio,1)
                                Apatho.Report(1,nodes)
                            }
                    }
                }
            case 3:
                if !Exist("A_patho.csv") {
                    fmt.Println("\nERROR: the pathological attractor set (A_patho) must be computed ans saved")
                } else {
                    Apatho.Load(1)
                    Aversus.Versus(Apatho)
                    Aversus.Report(2,nodes)
                }
            case 4:
                if len(D)==0 {
                    fmt.Println("\nERROR: the state space must be generated or loaded")
                } else if !Exist("A_physio.csv") || !Exist("A_patho.csv") || !Exist("A_versus.csv") {
                    fmt.Println("\nERROR: the physiological attractor set (A_physio), the pathological attractor set (A_patho) and the pathological attractors (A_versus) must be computed and saved")
                } else {
                    Aversus.Load(2)
                    if len(Aversus)==0 {
                        fmt.Println("\nWARNING: there are no pathological attractors to remove (A_versus is empty)")
                    } else {
                        Aphysio.Load(0)
                        Apatho.Load(1)
                        rmin=int(Prompt("\nNumber of targets per bullet (lower bound): ",ToV(Range(1,len(nodes)+1))))
                        rmax=int(Prompt("\nNumber of targets per bullet (upper bound): ",ToV(Range(rmin,len(nodes)+1))))
                        Btherap.Compute(fpatho,D,rmin,rmax,maxtarg,maxmoda,Aphysio,Apatho,Aversus,vals)
                        Btherap.Report(nodes,Aphysio,Aversus,rmin,rmax)
                    }
                }
            case 5:
                tochange=-1
                for tochange!=0 {// TODO case 0 does not break it...
                    tochange=int(Prompt(strings.Join([]string{
                        "\nChange:",
                        "    [1] maxtarg ("+strconv.FormatInt(int64(maxtarg),10)+")",
                        "    [2] maxmoda ("+strconv.FormatInt(int64(maxmoda),10)+")",
                        "    [3] maxD    ("+strconv.FormatInt(int64(maxD),10)+")",
                        "    [0] done",
                        "\nTo change: ",
                    },"\n"),ToV(Range(0,4))))
                    switch tochange {
                        case 1:
                            maxtarg=int(Prompt("\nmaxtarg=",Vector{}))
                            fmt.Println("\nWARNING: you should recompute therapeutic bullets, if any")
                        case 2:
                            maxmoda=int(Prompt("\nmaxmoda=",Vector{}))
                            fmt.Println("\nWARNING: you should recompute therapeutic bullets, if any")
                        case 3:
                            maxD=int(Prompt("\nmaxD=",Vector{}))
                            fmt.Println("\nWARNING: you should regenerate the state space, if any")
                        case 0:
                            fmt.Println("\nINFO: OK!")
                            // break
                    }
                }
            case 6:
                fmt.Println(strings.Join([]string{
                    "\nAlready saved:",
                    "    state space: "+strconv.FormatBool(Exist("state_space.csv")),
                    "    A_physio:    "+strconv.FormatBool(Exist("A_physio.csv")),
                    "    A_patho:     "+strconv.FormatBool(Exist("A_patho.csv")),
                    "    A_versus:    "+strconv.FormatBool(Exist("A_versus.csv")),
                    "    B_therap:    "+strconv.FormatBool(Exist("B_therap.txt")),
                },"\n"))
            case 7:
                fmt.Println(strings.Join([]string{
                    "\nHow to use the algorithm:",
                    "    1) read my article, it explains:",
                    "        * the algorithm",
                    "        * its purpose",
                    "        * how it works (illustrated with the example)",
                    "        * its applications (illustrated with a Boolean model of Fanconi anemia)",
                    "        * its strengths and weaknesses",
                    "        * its possible further improvements/contributions",
                    "       freely available on arXiv and HAL:",
                    "        * http://arxiv.org/abs/1407.4374",
                    "        * https://hal.archives-ouvertes.fr/hal-01024788",
                    "    2) generate or load the state space: [1]",
                    "        * you can regenerate or reload the state space whenever you want",
                    "        * when prompted, you can save the state space (optional)",
                    "    3) compute the physiological attractor set: [2]",
                    "        * when prompted, set the setting to physiological",
                    "        * will return A_physio",
                    "        * when prompted, save A_physio (required for the next steps)",
                    "        * will create the files A_physio.csv (for the algorithm) and A_physio.txt (for you)",
                    "    4) compute the pathological attractor set: [2]",
                    "        * when prompted, set the setting to pathological",
                    "        * will return A_patho",
                    "        * when prompted, save A_patho (required for the next steps)",
                    "        * will create the files A_patho.csv (for the algorithm) and A_patho.txt (for you)",
                    "    5) compute the pathological attractors: [3]",
                    "        * will return A_versus",
                    "        * when prompted, save A_versus (required for the next steps)",
                    "        * will create the files A_versus.csv (for the algorithm) and A_versus.txt (for you)",
                    "    6) compute therapeutic bullets: [4]",
                    "        * will return B_therap",
                    "        * when prompted, you can save B_therap (optional)",
                    "        * will create the file B_therap.txt (for you)",
                    "        * therapeutic bullets are reported as follow:",
                    "              x1[y1] x2[y2] x3[y3] ...",
                    "          meaning that the variable x has to be set to the value y",
                    "    * you can change parameter values (maxtarg, maxmoda and maxD): [5]",
                    "    * you can check what is already saved (state space, A_physio, A_patho, A_versus and B_therap): [6]",
                    "\nIf you rename, move or delete the csv files created by the algorithm then it will not recognize them when required, if any.",
                    "\nThe algorithm is tested with Go version go1.6 linux/amd64 (Arch Linux).",
                },"\n"))
            case 8:
                fmt.Println(strings.Join([]string{
                    "\nkali: a tool for in silico therapeutic target discovery",
                    "Copyright (C) 2013-2016 Arnaud Poret",
                    "\nThis program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.",
                    "\nThis program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.",
                    "\nYou should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/gpl.html",
                },"\n"))
            case 0:
                fmt.Println("\nINFO: Goodbye!\n")
                // break
        }
    }
}
