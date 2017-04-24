// Copyright (C) 2013-2017 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)
func DoTheJob(fphysio,fpatho func(Vector) Vector,ntarg,maxtarg,maxmoda,maxS,kmax,threshold,sync int,nodes []string,vals Vector) {
    var n,todo int
    n=len(nodes)
    if n==0 {
        panic("len(nodes)==0: nodes can not be empty")
    } else if len(vals)<2 {
        panic("len(vals)<2: the logic is at least two-valued")
    } else if sync!=0 && sync!=1 {
        panic("sync!=0 and sync!=1: sync must be 0 or 1")
    } else if ntarg<1 || ntarg>n {
        panic("ntarg<1 or ntarg>len(nodes): ntarg must be an integer in [1;number of nodes]")
    } else if threshold<0 || threshold>100 {
        panic("threshold<0 or threshold>100: threshold must be an integer in [0;100]")
    } else if maxtarg<1 || maxmoda<1 || maxS<1 || kmax<1 {
        panic("maxtarg<1 or maxmoda<1 or maxS<1 or kmax<1: maxtarg, maxmoda, maxS and kmax must be strictly positive integers")
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
                "    [6] change parameter values (ntarg, maxtarg, maxmoda, maxS, kmax, threshold, sync)",
                "    [7] check/clear what is saved (S, A_physio, A_patho, A_versus, Targ, Moda, B_therap)",
                "    [8] help",
                "    [9] license",
                "    [0] quit",
                "\nTo do [0-9] ",
            },"\n"),Range(0,10))
            if todo==1 {
                DoStateSpace(n,maxS,vals)
            } else if todo==2 {
                DoAttractorSet(fphysio,fpatho,kmax,sync,nodes)
            } else if todo==3 {
                DoVersus(nodes)
            } else if todo==4 {
                DoTestBullets(n,ntarg,maxtarg,maxmoda,vals)
            } else if todo==5 {
                DoTherapeuticBullets(fpatho,kmax,threshold,sync,nodes)
            } else if todo==6 {
                DoParameters(&ntarg,&maxtarg,&maxmoda,&maxS,&kmax,&threshold,&sync,n)
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
        fmt.Println("\nGoodbye!\n")
    }
}
