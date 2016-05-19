// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "math/rand"
import "strings"
import "time"
//#### DoTheJob ##############################################################//
func DoTheJob(fphysio,fpatho func(Matrix,int) Vector,ntarg,maxtarg,maxmoda,maxS int,nodes []string,vals Vector) {
    var todo int
    rand.Seed(int64(time.Now().Nanosecond()))
    for {
        todo=int(Prompt(strings.Join([]string{
            "\nWhat to do:",
            "    [1] generate the state space (S)",
            "    [2] compute an attractor set (A_physio or A_patho)",
            "    [3] compute the pathological attractors (A_versus)",
            "    [4] generate the bullets to test (Targ and Moda)",
            "    [5] compute therapeutic bullets (B_therap)",
            "    [6] change parameter values (ntarg, maxtarg, maxmoda and maxS)",
            "    [7] check what is already saved (S, A_physio, A_patho, A_versus, Targ, Moda and B_therap)",
            "    [8] help",
            "    [9] license",
            "    [0] quit",
            "\nTo do [0-9] ",
        },"\n"),ItoV(Range(0,10))))
        // "switch todo case 0: break" does not break the "for" loop...
        if todo==1 {DoS(len(nodes),maxS,vals)}
        if todo==2 {DoAset(fphysio,fpatho,nodes)}
        if todo==3 {DoVersus(nodes)}
        if todo==4 {DoBtest(len(nodes),ntarg,maxtarg,maxmoda,vals)}
        if todo==5 {DoBtherap(fpatho,nodes)}
        if todo==6 {DoParams(&ntarg,&maxtarg,&maxmoda,&maxS,len(nodes))}
        if todo==7 {DoSaved()}
        if todo==8 {DoHelp()}
        if todo==9 {DoNotice()}
        if todo==0 {break}
    }
    fmt.Println("\nINFO: Goodbye!\n")
}
