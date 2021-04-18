// Copyright (C) 2013-2021 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

package kali
import (
    "flag"
    "fmt"
    "math/rand"
    "os"
    "path/filepath"
    "strings"
    "time"
)
func DoTheJob(vals Vector,nodes []string,Physio,Patho func(Vector) Vector) {
    var (
        err error
        help,usage,license,ok bool
        fileName string
        args []string
        flagSet *flag.FlagSet
    )
    fileName=filepath.Base(os.Args[0])
    flagSet=flag.NewFlagSet("",flag.ContinueOnError)
    flagSet.Usage=func() {}
    flagSet.BoolVar(&help,"help",false,"")
    flagSet.BoolVar(&help,"h",false,"")
    flagSet.BoolVar(&usage,"usage",false,"")
    flagSet.BoolVar(&usage,"u",false,"")
    flagSet.BoolVar(&license,"license",false,"")
    flagSet.BoolVar(&license,"l",false,"")
    err=flagSet.Parse(os.Args[1:])
    if err!=nil {
        fmt.Println("Error: "+fileName+": "+err.Error())
    } else if help {
        fmt.Println(strings.Join([]string{
            "",
            "In silico therapeutic target discovery using network attractors.",
            "",
            "Currently, kali can operate on qualitative models, namely Boolean networks and",
            "multi-valued ones (a generalization of Boolean networks).",
            "",
            "Usage:",
            "    * "+fileName+" [options]",
            "    * "+fileName+" <command> [options] <arguments>",
            "",
            "Commands (should be run in that order):",
            "    * start: generate the set of the start states",
            "    * attractor: compute an attractor set",
            "    * versus: get the set of the pathological attractors",
            "    * bullet: generate the set of the bullets to test",
            "    * target: compute a set of therapeutic bullets",
            "",
            "Positional arguments: see the command-specific help ("+fileName+" <command> -help)",
            "",
            "Options:",
            "    * non command-specific options:",
            "        * -l/-license: print the GNU General Public License under which kali is",
            "        * -u/-usage: print usage only",
            "        * -h/-help: print this help",
            "    * command-specific options: "+fileName+" <command> -help",
            "",
            "Output files: see the command-specific help ("+fileName+" <command> -help)",
            "",
            "Cautions:",
            "    * non command-specific cautions:",
            "        * kali automatically saves and loads the files it creates and uses",
            "        * if one of these files already exists then it is overwritten",
            "        * if one of these files is renamed, moved or deleted then it can not be",
            "          loaded when required",
            "    * command-specific cautions: "+fileName+" <command> -help",
            "",
            "For command-specific help, run: "+fileName+" <command> -help",
            "",
            "For more information, see https://github.com/arnaudporet/kali",
            "",
            "For full explanation, see https://arxiv.org/pdf/1611.03144.pdf",
            "",
        },"\n"))
    } else if usage {
        fmt.Println(strings.Join([]string{
            "",
            "Usage:",
            "    * "+fileName+" [options]",
            "    * "+fileName+" <command> [options] <arguments>",
            "",
            "Commands (should be run in that order):",
            "    * start: generate the set of the start states",
            "    * attractor: compute an attractor set",
            "    * versus: get the set of the pathological attractors",
            "    * bullet: generate the set of the bullets to test",
            "    * target: compute a set of therapeutic bullets",
            "",
            "Positional arguments: see the command-specific help ("+fileName+" <command> -help)",
            "",
            "Options:",
            "    * non command-specific options:",
            "        * -l/-license: print the GNU General Public License under which kali is",
            "        * -u/-usage: print usage only",
            "        * -h/-help: print help",
            "    * command-specific options: "+fileName+" <command> -help",
            "",
        },"\n"))
    } else if license {
        fmt.Println(strings.Join([]string{
            "",
            "kali: in silico therapeutic target discovery using network attractors",
            "",
            "Copyright (C) 2013-2021 Arnaud Poret",
            "",
            "This program is free software: you can redistribute it and/or modify it under",
            "the terms of the GNU General Public License as published by the Free Software",
            "Foundation, either version 3 of the License, or (at your option) any later",
            "version.",
            "",
            "This program is distributed in the hope that it will be useful, but WITHOUT ANY",
            "WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A",
            "PARTICULAR PURPOSE. See the GNU General Public License for more details.",
            "",
            "You should have received a copy of the GNU General Public License along with",
            "this program. If not, see <https://www.gnu.org/licenses/>.",
            "",
        },"\n"))
    } else if len(flagSet.Args())==0 {
        fmt.Println("Error: "+fileName+": missing command, expecting one of: start, attractor, versus, bullet, target")
    } else if len(vals)<2 {
        fmt.Println("Error: "+fileName+": the used logic must be at least two-valued (i.e. at least Boolean)")
    } else if len(nodes)==0 {
        fmt.Println("Error: "+fileName+": no nodes provided")
    } else {
        ok=true
        CheckF(Physio,vals,len(nodes),&ok)
        if !ok {
            fmt.Println("Error: "+fileName+": something is wrong with the physiological function")
        } else {
            CheckF(Patho,vals,len(nodes),&ok)
            if !ok {
                fmt.Println("Error: "+fileName+": something is wrong with the pathological function")
            } else {
                rand.Seed(int64(time.Now().Nanosecond()))
                args=flagSet.Args()
                if args[0]=="start" {
                    DoStartStates(vals,nodes)
                } else if args[0]=="attractor" {
                    DoAttractorSet(nodes,Physio,Patho)
                } else if args[0]=="versus" {
                    DoVersus(nodes)
                } else if args[0]=="bullet" {
                    DoTestBullets(vals,nodes)
                } else if args[0]=="target" {
                    DoTherapeuticBullets(nodes,Patho)
                } else {
                    fmt.Println("Error: "+fileName+": "+args[0]+": unknown command, expecting one of: start, attractor, versus, bullet, target")
                }
            }
        }
    }
}
