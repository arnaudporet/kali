// Copyright (C) 2013-2018 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
func DoAttractorSet(fphysio,fpatho func(Vector) Vector,kmax,sync int,nodes []string) {
    var (
        setting int
        A AttractorSet
    )
    if !Exist("S.csv") {
        fmt.Println("\nERROR: S must be generated")
    } else {
        setting=GetInt("\nSetting: physiological [0] or pathological [1]? [0/1] ",[]int{0,1})
        if setting==0 {
            A,_=ComputeAttractorSet(fphysio,LoadMat("S.csv"),Bullet{},kmax,0,sync,-9,AttractorSet{})
            A.Report(nodes,0)
        } else if setting==1 {
            if !Exist("A_physio.csv") {
                fmt.Println("\nERROR: A_physio must be computed")
            } else {
                A,_=ComputeAttractorSet(fpatho,LoadMat("S.csv"),Bullet{},kmax,1,sync,-9,LoadAttractorSet(0))
                A.Report(nodes,1)
            }
        }
    }
}
