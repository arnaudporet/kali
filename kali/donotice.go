// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "io/ioutil"
//#### DoNotice ##############################################################//
func DoNotice() {
    var notice []byte
    notice,_=ioutil.ReadFile("./kali/notice.txt")
    fmt.Printf("%s",notice)
}
