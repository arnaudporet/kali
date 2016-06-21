// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
func DoAttractorSet(fphysio,fpatho func(Matrix,int) Vector,nodes []string) {
    var setting int
    if !Exist("S.csv") {
        fmt.Println("\nERROR: S must be generated")
    } else {
        setting=GetInt("\nSetting: physiological [0] or pathological [1]? [0/1] ",[]int{0,1})
        if setting==0 {
            ComputeAttractorSet(fphysio,LoadMatrix("S.csv"),Bullet{},AttractorSet{},0).Report(nodes,0)
        } else if setting==1 {
            if !Exist("A_physio.csv") {
                fmt.Println("\nERROR: A_physio must be computed")
            } else {
                ComputeAttractorSet(fpatho,LoadMatrix("S.csv"),Bullet{},LoadAttractorSet(0),1).Report(nodes,1)
            }
        }
    }
}
