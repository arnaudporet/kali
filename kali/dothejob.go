// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "math/rand"
import "strings"
import "time"
//#### DoTheJob ##############################################################//
func DoTheJob(fphysio,fpatho func(Matrix,int) Vector,ntarg,maxtarg,maxmoda,maxS int,nodes []string,vals Vector) {
    var todo int
    rand.Seed(int64(time.Now().Nanosecond()))
    todo=-1
    for todo!=0 {
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
        switch todo {
            case 1: DoS(len(nodes),maxS,vals)
            case 2: DoAset(fphysio,fpatho,nodes)
            case 3: DoVersus(nodes)
            case 4: DoBtest(len(nodes),ntarg,maxtarg,maxmoda,vals)
            case 5: DoBtherap(fpatho,nodes)
            case 6: DoParams(&ntarg,&maxtarg,&maxmoda,&maxS,len(nodes))
            case 7: DoSaved()
            case 8: DoHelp()
            case 9: DoNotice()
            case 0: fmt.Println("\nINFO: Goodbye!\n")
        }
    }
}
