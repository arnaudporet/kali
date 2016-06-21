// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html

//#### HOWTO #################################################################//

// 1) read my article (all is explained inside), freely available at:
//        * arXiv: https://arxiv.org/abs/1407.4374
//        * HAL:   https://hal.archives-ouvertes.fr/hal-01024788
// 2) read the following comments
// 3) replace the content with your own stuff
// 4) run (tested with Go version go1.6.2 linux/amd64 (Arch Linux)):
//        go run example.go
//    it is possible that the Go package has a different name depending on your
//    OS/Linux distribution
//    for example, with Ubuntu, it is named "golang", so it may be
//    "golang-go run yourfile.go" instead of "go run yourfile.go"
// 5) check the help proposed by kali

// This example is a generic Boolean model of the mammalian cell cycle by Adrien
// Faure and coworkers [1].

//############################################################################//

package main
import "./kali"// import kali, change the path if you move it
func main() {
    var (
        ntarg,maxtarg,maxmoda,maxS,threshold int
        nodes []string
        vals kali.Vector
    )
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
    // The domain of value.
    // Values are real numbers in [0;1].
    // {0,1} for Boolean logic or, for example, {0,0.5,1} for three-valued
    // logic.
    vals=kali.Vector{0.0,1.0}
    // The number of targets per bullet.
    // ntarg is an integer in [1;number of nodes].
    // Can be changed at run-time.
    ntarg=2
    // The maximum number of target combinations to test.
    // maxtarg is an integer > 0.
    // If it exceeds its maximal possible value then kali will automatically
    // decrease it to its maximal possible value.
    // Can be changed at run-time.
    maxtarg=int(1e2)
    // The maximum number of modality arrangements to test for each target
    // combination.
    // maxmoda is an integer > 0.
    // If it exceeds its maximal possible value then kali will automatically
    // decrease it to its maximal possible value.
    // Can be changed at run-time.
    maxmoda=int(1e2)
    // The maximum number of initial states to test when searching an attractor.
    // maxS is an integer > 0.
    // If it exceeds its maximal possible value then kali will automatically
    // decrease it to its maximal possible value.
    // Can be changed at run-time.
    maxS=int(1e4)
    // The goal of therapeutic bullets is to increase the coverage of the
    // pathological state space by the physiological one.
    // To be considered therapeutic, this increase must be >= threshold (in
    // percents of the pathological state space).
    // threshold is an integer in [1;100].
    // Can be changed at run-time.
    threshold=5
    kali.DoTheJob(fphysio,fpatho,ntarg,maxtarg,maxmoda,maxS,threshold,nodes,vals)
}

// THE NETWORK MUST BE DETERMINISTIC (implementation of asynchronous updating
// scheme in progress).

// To cope with both Boolean and multivalued logic, the Zadeh fuzzy logic
// operators are used:
//     x AND y = min(x,y)
//     x OR y  = max(x,y)
//     NOT x   = 1-x

// The transition function of the physiological variant.
func fphysio(x kali.Matrix,k int) kali.Vector {
    return kali.Vector{
        // replace the following equations with your own stuff
        // your equations coded in the same way
        // note that the variable numbering starts at 0
        x[0][k],// CycD
        kali.Max(kali.Min(1.0-x[0][k],1.0-x[3][k],1.0-x[4][k],1.0-x[9][k]),kali.Min(x[5][k],1.0-x[0][k],1.0-x[9][k])),// Rb
        kali.Max(kali.Min(1.0-x[1][k],1.0-x[4][k],1.0-x[9][k]),kali.Min(x[5][k],1.0-x[1][k],1.0-x[9][k])),// E2F
        kali.Min(x[2][k],1.0-x[1][k]),// CycE
        kali.Max(kali.Min(x[2][k],1.0-x[1][k],1.0-x[6][k],1.0-kali.Min(x[7][k],x[8][k])),kali.Min(x[4][k],1.0-x[1][k],1.0-x[6][k],1.0-kali.Min(x[7][k],x[8][k]))),// CycA
        kali.Max(kali.Min(1.0-x[0][k],1.0-x[3][k],1.0-x[4][k],1.0-x[9][k]),kali.Min(x[5][k],1.0-kali.Min(x[3][k],x[4][k]),1.0-x[9][k],1.0-x[0][k])),// p27
        x[9][k],// Cdc20
        kali.Max(kali.Min(1.0-x[4][k],1.0-x[9][k]),x[6][k],kali.Min(x[5][k],1.0-x[9][k])),// Cdh1
        kali.Max(1.0-x[7][k],kali.Min(x[7][k],x[8][k],kali.Max(x[6][k],x[4][k],x[9][k]))),// UbcH10
        kali.Min(1.0-x[6][k],1.0-x[7][k]),// CycB
    }
}

// The transition function of the pathological variant.
func fpatho(x kali.Matrix,k int) kali.Vector {
    return kali.Vector{
        // replace the following equations with your own stuff
        // your equations coded in the same way
        // note that the variable numbering starts at 0
        x[0][k],// CycD
        0.0,// Rb
        kali.Max(kali.Min(1.0-x[1][k],1.0-x[4][k],1.0-x[9][k]),kali.Min(x[5][k],1.0-x[1][k],1.0-x[9][k])),// E2F
        kali.Min(x[2][k],1.0-x[1][k]),// CycE
        kali.Max(kali.Min(x[2][k],1.0-x[1][k],1.0-x[6][k],1.0-kali.Min(x[7][k],x[8][k])),kali.Min(x[4][k],1.0-x[1][k],1.0-x[6][k],1.0-kali.Min(x[7][k],x[8][k]))),// CycA
        kali.Max(kali.Min(1.0-x[0][k],1.0-x[3][k],1.0-x[4][k],1.0-x[9][k]),kali.Min(x[5][k],1.0-kali.Min(x[3][k],x[4][k]),1.0-x[9][k],1.0-x[0][k])),// p27
        x[9][k],// Cdc20
        kali.Max(kali.Min(1.0-x[4][k],1.0-x[9][k]),x[6][k],kali.Min(x[5][k],1.0-x[9][k])),// Cdh1
        kali.Max(1.0-x[7][k],kali.Min(x[7][k],x[8][k],kali.Max(x[6][k],x[4][k],x[9][k]))),// UbcH10
        kali.Min(1.0-x[6][k],1.0-x[7][k]),// CycB
    }
}

// [1] Adrien Faure, Aurelien Naldi, Claudine Chaouiya, Denis Thieffry.
// Dynamical analysis of a generic Boolean model for the control of the
// mammalian cell cycle. Bioinformatics 22(14):e124-e131. Oxford Univ Press.
