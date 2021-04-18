// Copyright (C) 2013-2021 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html.

//#### HOWTO #################################################################//

// 1) read my article: https://arxiv.org/pdf/1611.03144.pdf

// 2) read the "readme.md" file

// 3) read the present file

// 4) replace the content of the present file with your own stuff

// 5) run in a terminal emulator: go run example.go -help

// The Go package can have different names depending on your operating system.
// For example, with Ubuntu the Go package is named "golang". Consequently,
// running a Go file with Ubuntu might be "golang-go run yourfile.go" instead of
// "go run yourfile.go" with Arch Linux. Otherwise see
// https://golang.org/doc/install.

// This example is a simple and fictive Boolean model used to conveniently
// illustrate kali.

// A biological case study is also proposed in the "case_study" folder to
// address a concrete case: a published logic-based model of bladder
// tumorigenesis.

//############################################################################//

package main

// Import kali.
import (
    "github.com/arnaudporet/kali/kali"
)

func main() {
    var (
        nodes []string
        vals kali.Vector
    )
    // The domain of value of the used logic (i.e. the domain of value of the
    // model's variables).
    // {0,1} for Boolean logic or, for example, {0,0.5,1} for three-valued
    // logic.
    vals=kali.Vector{
        0,
        1,
    }
    // The node names.
    // The order in which the nodes are listed must be identical in:
    //     * the updating function of the physiological variant (see below)
    //     * the updating function of the pathological variant (see below)
    //     * this list of the node names
    nodes=[]string{
        "do",
        "factory",
        "energy",
        "locker",
        "releaser",
        "sequester",
        "activator",
        "effector",
        "task",
    }
    // Do the job.
    kali.DoTheJob(vals,nodes,Physio,Patho)
}

// To cope with both Boolean and multivalued logic, the Zadeh's logical
// operators are used:
//     x AND y = min(x,y)
//     x OR y  = max(x,y)
//     NOT x   = 1-x

// The updating function of the physiological variant.
func Physio(x kali.Vector) kali.Vector {
    return kali.Vector{
        // Replace the following equations with your own stuff.
        // Your equations encoded in the same way.
        // Note that the variable numbering starts at 0.
        x[0],// do
        x[1],// factory
        kali.Max(kali.Min(x[2],1-x[8]),x[1]),// energy
        1-x[2],// locker
        x[0],// releaser
        1-x[4],// sequester
        kali.Min(x[0],1-x[3]),// activator
        kali.Min(x[6],1-x[5]),// effector
        x[7],// task
    }
}

// The updating function of the pathological variant.
// In this example, it is obtained by knocking down the locker.
func Patho(x kali.Vector) kali.Vector {
    return kali.Vector{
        // Replace the following equations with your own stuff.
        // Your equations encoded in the same way.
        // Note that the variable numbering starts at 0.
        x[0],// do
        x[1],// factory
        kali.Max(kali.Min(x[2],1-x[8]),x[1]),// energy
        0,// locker (knocked down)
        x[0],// releaser
        1-x[4],// sequester
        kali.Min(x[0],1-x[3]),// activator
        kali.Min(x[6],1-x[5]),// effector
        x[7],// task
    }
}
