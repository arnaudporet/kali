// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
//#### DoAset ################################################################//
func DoAset(fphysio,fpatho func(Matrix,int) Vector,nodes []string) {
    var setting int
    var S Matrix
    var nullb Bullet
    var Aphysio,Apatho,nullset Aset
    if !Exist("S.csv") {
        fmt.Println("\nERROR: S must be generated")
    } else {
        setting=int(Prompt("\nSetting: physiological [0], pathological [1]? [0/1] ",Vector{0.0,1.0}))
        switch setting {
            case 0:
                S.Load("S.csv")
                Aphysio.Compute(fphysio,S,nullb,nullset,0)
                Aphysio.Report(0,nodes)
            case 1:
                if !Exist("A_physio.csv") {
                    fmt.Println("\nERROR: A_physio must be computed")
                } else {
                    S.Load("S.csv")
                    Aphysio.Load(0)
                    Apatho.Compute(fpatho,S,nullb,Aphysio,1)
                    Apatho.Report(1,nodes)
                }
        }
    }
}
