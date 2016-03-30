// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit http://www.gnu.org/licenses/gpl.html

//#### HOWTO #################################################################//

// 1) read my article (all is explained inside), freely available at:
//        * http://arxiv.org/abs/1407.4374
//        * https://hal.archives-ouvertes.fr/hal-01024788
// 2) read the following comments
// 3) replace the contents with your own stuff
// 4) run (tested with Go version go1.6 linux/amd64 (Arch Linux)):
//        go run example.go
//    it is possible that the Go package has a different name depending on your
//    OS/Linux distribution
//    for example, with Ubuntu, it is named "golang", so it may be
//    "golang-go run yourfile.go" instead of "go run yourfile.go"
// 5) check the help proposed by the algorithm

// This example is a generic Boolean model of the mammalian cell cycle by Adrien
// Faure and coworkers [1].

//############################################################################//

package main
import "./kali"//import the algorithm, change the path if you move it
func main() {
    var maxtarg,maxmoda,maxD int
    var nodes []string
    var vals kali.Vector
    // The node names.
    nodes=[]string{
        "CycD",
        "Rb",
        "E2F",
        "CycE",
        "CycA",
        "p27",
        "Cdc20",
        "Cdh1",
        "UbcH10",
        "CycB",
    }
    // The domain of values.
    // {0,1} for Boolean logic or, for example, {0,0.5,1} for three-valued
    // logic.
    vals=kali.Vector{0.0,1.0}
    // The maximum number of target combinations to test.
    // If it exceeds its maximal possible value then the algorithm will
    // automatically decrease it to its maximal possible value.
    // Can be changed at run-time.
    maxtarg=int(1e2)
    // The maximum number of modality arrangements to test for each target
    // combination.
    // If it exceeds its maximal possible value then the algorithm will
    // automatically decrease it to its maximal possible value.
    // Can be changed at run-time.
    maxmoda=int(1e2)
    // The maximum number of initial conditions to test.
    // If it exceeds its maximal possible value then the algorithm will
    // automatically decrease it to its maximal possible value.
    // Can be changed at run-time.
    maxD=int(1e4)
    kali.DoTheJob(fphysio,fpatho,maxtarg,maxmoda,maxD,nodes,vals)
}

// The network must be deterministic.

// To cope with both Boolean and multivalued logic, the Zadeh fuzzy logic
// operators are used:
//     x AND y = min(x,y)
//     x OR y = max(x,y)
//     NOT x = 1-x

//#### fphysio ###############################################################//

// The transition function of the physiological variant.
func fphysio(x kali.Matrix,k int) kali.Vector {
    var min,max func(...float64) float64
    var y kali.Vector
    min=kali.Min
    max=kali.Max
    y=make(kali.Vector,x.Size(1))
    // replace the following equations with your own stuff
    // your equations coded in the same way: y[i]=f(x[0][k],x[1][k],x[2][k],...)
    // note that the numbering starts at 0
    y[0]=x[0][k]// CycD
    y[1]=max(min(1.0-x[0][k],1.0-x[3][k],1.0-x[4][k],1.0-x[9][k]),min(x[5][k],1.0-x[0][k],1.0-x[9][k]))// Rb
    y[2]=max(min(1.0-x[1][k],1.0-x[4][k],1.0-x[9][k]),min(x[5][k],1.0-x[1][k],1.0-x[9][k]))// E2F
    y[3]=min(x[2][k],1.0-x[1][k])// CycE
    y[4]=max(min(x[2][k],1.0-x[1][k],1.0-x[6][k],1.0-min(x[7][k],x[8][k])),min(x[4][k],1.0-x[1][k],1.0-x[6][k],1.0-min(x[7][k],x[8][k])))// CycA
    y[5]=max(min(1.0-x[0][k],1.0-x[3][k],1.0-x[4][k],1.0-x[9][k]),min(x[5][k],1.0-min(x[3][k],x[4][k]),1.0-x[9][k],1.0-x[0][k]))// p27
    y[6]=x[9][k]// Cdc20
    y[7]=max(min(1.0-x[4][k],1.0-x[9][k]),x[6][k],min(x[5][k],1.0-x[9][k]))// Cdh1
    y[8]=max(1.0-x[7][k],min(x[7][k],x[8][k],max(x[6][k],x[4][k],x[9][k])))// UbcH10
    y[9]=min(1.0-x[6][k],1.0-x[7][k])// CycB
    return y
}

//#### fpatho ################################################################//

// The transition function of the pathological variant.
func fpatho(x kali.Matrix,k int) kali.Vector {
    var min,max func(...float64) float64
    var y kali.Vector
    min=kali.Min
    max=kali.Max
    y=make(kali.Vector,x.Size(1))
    // replace the following equations with your own stuff
    // your equations coded in the same way: y[i]=f(x[0][k],x[1][k],x[2][k],...)
    // note that the numbering starts at 0
    y[0]=x[0][k]// CycD
    y[1]=0.0// Rb
    y[2]=max(min(1.0-x[1][k],1.0-x[4][k],1.0-x[9][k]),min(x[5][k],1.0-x[1][k],1.0-x[9][k]))// E2F
    y[3]=min(x[2][k],1.0-x[1][k])// CycE
    y[4]=max(min(x[2][k],1.0-x[1][k],1.0-x[6][k],1.0-min(x[7][k],x[8][k])),min(x[4][k],1.0-x[1][k],1.0-x[6][k],1.0-min(x[7][k],x[8][k])))// CycA
    y[5]=max(min(1.0-x[0][k],1.0-x[3][k],1.0-x[4][k],1.0-x[9][k]),min(x[5][k],1.0-min(x[3][k],x[4][k]),1.0-x[9][k],1.0-x[0][k]))// p27
    y[6]=x[9][k]// Cdc20
    y[7]=max(min(1.0-x[4][k],1.0-x[9][k]),x[6][k],min(x[5][k],1.0-x[9][k]))// Cdh1
    y[8]=max(1.0-x[7][k],min(x[7][k],x[8][k],max(x[6][k],x[4][k],x[9][k])))// UbcH10
    y[9]=min(1.0-x[6][k],1.0-x[7][k])// CycB
    return y
}

// [1] Adrien Faure, Aurelien Naldi, Claudine Chaouiya, Denis Thieffry.
// Dynamical analysis of a generic Boolean model for the control of the
// mammalian cell cycle. Bioinformatics 22(14):e124-e131. Oxford Univ Press.
