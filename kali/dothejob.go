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
func DoTheJob(fphysio,fpatho func(Matrix,int) Vector,ntarg,maxtarg,maxmoda,maxS int,nodes []string,vals Vector) {
    var todo,whole,setting,tochange,gen,save int
    var S,Targ,Moda Matrix
    var nullb Bullet
    var Aphysio,Apatho,Aversus,nullset Aset
    var Btherap Bset
    rand.Seed(int64(time.Now().Nanosecond()))
    todo=-1
    for todo!=0 {// TODO case 0 does not break it...
        todo=int(Prompt(strings.Join([]string{
            "\nWhat to do:",
            "    [1] generate/load the state space (S)",
            "    [2] compute an attractor set (A_physio, A_patho)",
            "    [3] compute the pathological attractors (A_versus)",
            "    [4] generate/load the bullets to test ({Targ,Moda})",
            "    [5] compute therapeutic bullets (B_therap)",
            "    [6] change parameter values (ntarg, maxtarg, maxmoda, maxS)",
            "    [7] check what is already saved (S, A_physio, A_patho, A_versus, Targ, Moda, B_therap)",
            "    [8] help",
            "    [9] license",
            "    [0] quit",
            "\nTo do [0-9] ",
        },"\n"),ItoV(Range(0,10))))
        switch todo {
            case 1:
                gen=int(Prompt("\nS: load [0], generate [1]? [0/1] ",Vector{0.0,1.0}))
                switch gen {
                    case 0:
                        if !Exist("S.csv") {
                            fmt.Println("\nERROR: unable to load S.csv")
                        } else {
                            S.Load("S.csv")
                            fmt.Println("\nINFO: S loaded")
                        }
                    case 1:
                        whole=int(Prompt("\nS cardinality: "+strconv.FormatFloat(math.Pow(float64(len(vals)),float64(len(nodes))),'f',-1,64)+"\n\nCompute whole S [1], limit to maxS [0] ("+strconv.FormatInt(int64(maxS),10)+")? [0/1] ",Vector{0.0,1.0}))
                        switch whole {
                            case 0: S=GenArrangMat(vals,len(nodes),maxS).T()
                            case 1: S=GenS(vals,len(nodes))
                        }
                        fmt.Println("\nINFO: S generated")
                        save=int(Prompt("\nSave? (optional) [0/1] ",Vector{0.0,1.0}))
                        if save==1 {
                            S.Save("S.csv")
                            fmt.Println("\nINFO: set saved as S.csv")
                        }
                }
            case 2:
                if len(S)==0 {
                    fmt.Println("\nERROR: S must be generated/loaded")
                } else {
                    setting=int(Prompt("\nSetting: physiological [0], pathological [1]? [0/1] ",Vector{0.0,1.0}))
                    switch setting {
                        case 0:
                            Aphysio.Compute(fphysio,S,nullb,nullset,0)
                            Aphysio.Report(0,nodes)
                        case 1:
                            if !Exist("A_physio.csv") {
                                fmt.Println("\nERROR: unable to load A_physio.csv")
                            } else {
                                Aphysio.Load(0)
                                Apatho.Compute(fpatho,S,nullb,Aphysio,1)
                                Apatho.Report(1,nodes)
                            }
                    }
                }
            case 3:
                if !Exist("A_patho.csv") {
                    fmt.Println("\nERROR: unable to load A_patho.csv")
                } else {
                    Apatho.Load(1)
                    Aversus.Versus(Apatho)
                    Aversus.Report(2,nodes)
                }
            case 4:
                gen=int(Prompt("\n{Targ,Moda}: load [0], generate [1]? [0/1] ",Vector{0.0,1.0}))
                switch gen {
                    case 0:
                        if !(Exist("Targ.csv") && Exist("Moda.csv")) {
                            fmt.Println("\nERROR: unable to load Targ.csv, Moda.csv")
                        } else {
                            Targ.Load("Targ.csv")
                            Moda.Load("Moda.csv")
                            fmt.Println("\nINFO: {Targ,Moda} loaded")
                        }
                    case 1:
                        Targ=GenCombiMat(ItoV(Range(0,len(nodes))),ntarg,maxtarg)
                        Moda=GenArrangMat(vals,ntarg,maxmoda)
                        fmt.Println("\nINFO: {Targ,Moda} generated")
                        save=int(Prompt("\nSave? (optional) [0/1] ",Vector{0.0,1.0}))
                        if save==1 {
                            Targ.Save("Targ.csv")
                            Moda.Save("Moda.csv")
                            fmt.Println("\nINFO: sets saved as Targ.csv, Moda.csv")
                        }
                }
            case 5:
                if len(S)==0 {
                    fmt.Println("\nERROR: S must be generated/loaded")
                } else if len(Targ)==0 || len(Moda)==0 {
                    fmt.Println("\nERROR: {Targ,Moda} must be generated/loaded")
                } else if !(Exist("A_physio.csv") && Exist("A_patho.csv") && Exist("A_versus.csv")) {
                    fmt.Println("\nERROR: unable to load A_physio.csv, A_patho.csv, A_versus.csv")
                } else {
                    Aversus.Load(2)
                    if len(Aversus)==0 {
                        fmt.Println("\nWARNING: no pathological attractors to remove (A_versus empty)")
                    } else {
                        Aphysio.Load(0)
                        Apatho.Load(1)
                        Btherap.Compute(fpatho,S,Targ,Moda,Aphysio,Apatho,Aversus)
                        Btherap.Report(nodes,Aphysio,Aversus)
                    }
                }
            case 6:
                tochange=-1
                for tochange!=0 {// TODO case 0 does not break it...
                    tochange=int(Prompt(strings.Join([]string{
                        "\nChange:",
                        "    [1] ntarg   ("+strconv.FormatInt(int64(ntarg),10)+")",
                        "    [2] maxtarg ("+strconv.FormatInt(int64(maxtarg),10)+")",
                        "    [3] maxmoda ("+strconv.FormatInt(int64(maxmoda),10)+")",
                        "    [4] maxS    ("+strconv.FormatInt(int64(maxS),10)+")",
                        "    [0] done",
                        "\nTo change [0-4] ",
                    },"\n"),ItoV(Range(0,5))))
                    switch tochange {
                        case 1:
                            ntarg=int(Prompt("\nntarg=",ItoV(Range(1,len(nodes)+1))))
                            fmt.Println("\nWARNING: you should regenerate {Targ,Moda}, if any")
                        case 2:
                            maxtarg=int(Prompt("\nmaxtarg=",Vector{}))
                            fmt.Println("\nWARNING: you should regenerate {Targ,Moda}, if any")
                        case 3:
                            maxmoda=int(Prompt("\nmaxmoda=",Vector{}))
                            fmt.Println("\nWARNING: you should regenerate {Targ,Moda}, if any")
                        case 4:
                            maxS=int(Prompt("\nmaxS=",Vector{}))
                            fmt.Println("\nWARNING: you should regenerate S, if any")
                        case 0:
                            fmt.Println("\nINFO: OK!")
                            // break
                    }
                }
            case 7:
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
            case 8:
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
                    "    2) generate/load the state space (S): [1]",
                    "        * you can regenerate/reload S whenever you want",
                    "        * when prompted, you can save S (optional)",
                    "    3) compute the physiological attractor set (A_physio): [2]",
                    "        * when prompted, set the setting to physiological",
                    "        * when prompted, save A_physio (required for the next steps)",
                    "    4) compute the pathological attractor set (A_patho): [2]",
                    "        * when prompted, set the setting to pathological",
                    "        * when prompted, save A_patho (required for the next steps)",
                    "    5) compute the pathological attractors (A_versus): [3]",
                    "        * when prompted, save A_versus (required for the next steps)",
                    "    6) generate/load the bullets to test ({Targ,Moda}): [4]",
                    "        * you can regenerate/reload {Targ,Moda} whenever you want",
                    "        * when prompted, you can save {Targ,Moda} (optional)",
                    "    7) compute therapeutic bullets (B_therap): [5]",
                    "        * when prompted, you can save B_therap (optional)",
                    "        * therapeutic bullets are reported as follow:",
                    "              x1[y1] x2[y2] x3[y3] ...",
                    "          meaning that the variable x has to be set to the value y",
                    "    * you can change parameter values (ntarg, maxtarg, maxmoda, maxS): [6]",
                    "    * you can check what is already saved (S, A_physio, A_patho, A_versus, Targ, Moda, B_therap): [7]",
                    "\nIf you rename, move or delete the csv files created by the algorithm then it will not recognize them when required, if any.",
                    "\nThe algorithm is tested with Go version go1.6 linux/amd64 (Arch Linux).",
                },"\n"))
            case 9:
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
