// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)
func DoTheJob(fphysio,fpatho func(Matrix,int) Vector,ntarg,maxtarg,maxmoda,maxS,threshold int,nodes []string,vals Vector) {
    var n,todo int
    n=len(nodes)
    if n==0 {
        panic("len(nodes)==0: nodes can not be empty (i.e. the network can not be empty)")
    } else if len(vals)<2 {
        panic("len(vals)<2: logic is at least two-valued (i.e. at least Boolean)")
    } else if ntarg<1 || ntarg>n {
        panic("ntarg<1 or ntarg>len(nodes): ntarg must be an integer in [1;number of nodes] (i.e. 1<=ntarg<=len(nodes))")
    } else if maxtarg<1 || maxmoda<1 || maxS<1 {
        panic("maxtarg<1 or maxmoda<1 or maxS<1: maxtarg, maxmoda and maxS must be strictly positive integers (i.e. maxtarg,maxmoda,maxS>0)")
    } else if threshold<1 || threshold>100 {
        panic("threshold<1 or threshold>100: threshold must be an integer in [1;100] (i.e. 1<=threshold<=100)")
    } else {
        rand.Seed(int64(time.Now().Nanosecond()))
        for {
            todo=GetInt(strings.Join([]string{
                "\nWhat to do:",
                "    [1] generate the state space (S)",
                "    [2] compute an attractor set (A_physio or A_patho)",
                "    [3] get the pathological attractors (A_versus)",
                "    [4] generate the bullets to test (Targ and Moda)",
                "    [5] compute therapeutic bullets (B_therap)",
                "    [6] change parameter values (ntarg, maxtarg, maxmoda, maxS and threshold)",
                "    [7] check what is already saved (S, A_physio, A_patho, A_versus, Targ, Moda and B_therap)",
                "    [8] help",
                "    [9] license",
                "    [0] quit",
                "\nTo do [0-9] ",
            },"\n"),Range(0,10))
            if todo==1 {
                DoStateSpace(n,maxS,vals)
            } else if todo==2 {
                DoAttractorSet(fphysio,fpatho,nodes)
            } else if todo==3 {
                DoVersus(nodes)
            } else if todo==4 {
                DoTestBullets(ntarg,maxtarg,maxmoda,n,vals)
            } else if todo==5 {
                DoTherapeuticBullets(fpatho,threshold,nodes)
            } else if todo==6 {
                DoParameters(&ntarg,&maxtarg,&maxmoda,&maxS,&threshold,n)
            } else if todo==7 {
                DoSaved()
            } else if todo==8 {
                DoHelp()
            } else if todo==9 {
                DoNotice()
            } else if todo==0 {
                break
            }
        }
        fmt.Println("\nINFO: Goodbye!\n")
    }
}
