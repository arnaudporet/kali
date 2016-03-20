// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "math"
import "strconv"
import "strings"
//#### DoTheJob ##############################################################//
func DoTheJob(fphysio,fpatho func(Matrix,int) Vector,maxtarg,maxmoda,maxD int,nodes []string,vals Vector) {
    var todo,wholeS,setting,rmin,rmax,tochange int
    var D Matrix
    var nullb Bullet
    var Aphysio,Apatho,Aversus,nullset Aset
    var Btherap Bset
    todo=-1
    for todo!=0 {// TODO case 0 does not break it...
        todo=int(Prompt("What to do:\n    [1] (re)generate the state space (or a subset of it)\n    [2] compute an attractor set\n    [3] compute the pathological attractors\n    [4] compute therapeutic bullets\n    [5] change parameter values\n    [6] check what is already saved\n    [7] help\n    [8] license\n    [0] quit\n\nTo do: ",ToV(Range(0,9))))
        switch todo {
            case 1:
                wholeS=int(Prompt("\nState space cardinality: "+strconv.FormatFloat(math.Pow(float64(len(vals)),float64(len(nodes))),'f',-1,64)+"\n\nCompute the whole state space? [0/1] ",Vector{0.0,1.0}))
                switch wholeS {
                    case 0: D=GenArrangMat(vals,len(nodes),maxD).T()
                    case 1: D=GenS(vals,len(nodes))
                }
                fmt.Println("\nState space (or a subset of it) (re)generated.\n")
            case 2:
                if len(D)==0 {
                    fmt.Println("\nThe state space (or a subset of it) must be generated to compute an attractor set. Ensure that the state space (or a subset of it) is already generated.\n")
                } else {
                    setting=int(Prompt("\nSetting: physiological [1], pathological [2]? [1/2] ",Vector{1.0,2.0}))
                    switch setting {
                        case 1:
                            Aphysio.Compute(fphysio,D,nullb,nullset,1)
                            Aphysio.Report(1,nodes)
                        case 2:
                            if !Exist("A_physio.csv") {
                                fmt.Println("\nThe file A_physio.csv is required to compute the pathological attractor set. Ensure that the physiological attractor set is already computed.\n")
                            } else {
                                Aphysio.Load(1)
                                Apatho.Compute(fpatho,D,nullb,Aphysio,2)
                                Apatho.Report(2,nodes)
                            }
                    }
                }
            case 3:
                if !Exist("A_patho.csv") {
                    fmt.Println("\nThe file A_patho.csv is required to compute the pathological attractors. Ensure that the pathological attractor set is already computed.\n")
                } else {
                    Apatho.Load(2)
                    Aversus.Versus(Apatho)
                    Aversus.Report(3,nodes)
                }
            case 4:
                if len(D)==0 {
                    fmt.Println("\nThe state space (or a subset of it) must be generated to compute therapeutic bullets. Ensure that the state space (or a subset of it) is already generated.\n")
                } else {
                    if !Exist("A_physio.csv") || !Exist("A_patho.csv") || !Exist("A_versus.csv") {
                        fmt.Println("\nThe files A_physio.csv, A_patho.csv and A_versus.csv are required to compute therapeutic bullets. Ensure that the physiological attractor set, the pathological attractor set and the pathological attractors are already computed.\n")
                    } else {
                        Aphysio.Load(1)
                        Apatho.Load(2)
                        Aversus.Load(3)
                        rmin=int(Prompt("\nNumber of targets per bullet (lower bound): ",Vector{}))
                        rmax=int(Prompt("Number of targets per bullet (upper bound): ",Vector{}))
                        Btherap.Compute(fpatho,D,rmin,rmax,maxtarg,maxmoda,Aphysio,Apatho,Aversus,vals)
                        Btherap.Report(nodes,Aphysio,Aversus,rmin,rmax)
                    }
                }
            case 5:
                tochange=-1
                for tochange!=0 {// TODO case 0 does not break it...
                    tochange=int(Prompt("\nChange:\n    [1] maxtarg ("+strconv.FormatInt(int64(maxtarg),10)+")\n    [2] maxmoda ("+strconv.FormatInt(int64(maxmoda),10)+")\n    [3] maxD ("+strconv.FormatInt(int64(maxD),10)+")\n    [0] done\n\nTo change: ",ToV(Range(0,4))))
                    switch tochange {
                        case 1:
                            maxtarg=int(Prompt("\nmaxtarg=",Vector{}))
                            fmt.Println("\nYou should recompute therapeutic bullets.")
                        case 2:
                            maxmoda=int(Prompt("\nmaxmoda=",Vector{}))
                            fmt.Println("\nYou should recompute therapeutic bullets.")
                        case 3:
                            maxD=int(Prompt("\nmaxD=",Vector{}))
                            fmt.Println("\nYou should regenerate the subset of the state space.")
                        case 0:
                            fmt.Println("\nOK\n")
                            // break
                    }
                }
            case 6: fmt.Println("\nAlready saved:\n    A_physio: "+strconv.FormatBool(Exist("A_physio.csv"))+"\n    A_patho: "+strconv.FormatBool(Exist("A_patho.csv"))+"\n    A_versus: "+strconv.FormatBool(Exist("A_versus.csv"))+"\n    B_therap: "+strconv.FormatBool(Exist("B_therap.txt"))+"\n")
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
                "    2) generate the state space (or a subset of it): [1]",
                "        * you can regenerate it whenever you want",
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
                "        * in case of multivalued logic, therapeutic bullets are reported as follow:",
                "              x1[y1] x2[y2] x3[y3] ...",
                "          meaning that the variable x has to be set to the value y",
                "    * you can change the parameter values (namely maxtarg, maxmoda and maxD): [5]",
                "    * you can check what is already saved: [6]",
                "\nIf you rename, move or delete the csv files created by the algorithm then it will not recognize them when required, if any.",
                "\nThe algorithm is tested with Go version go1.6 linux/amd64 (Arch Linux).\n",
                },"\n"))
            case 8:
                fmt.Println(strings.Join([]string{
                "\nkali: a tool for in silico therapeutic target discovery",
                "Copyright (C) 2013-2016 Arnaud Poret",
                "\nThis program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.",
                "\nThis program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.",
                "\nYou should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/gpl.html\n",
                },"\n"))
            case 0:
                fmt.Println("\nGoodbye!\n")
                // break
        }
    }
}
