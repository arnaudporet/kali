// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html

//#### HOWTO #################################################################//

// 1) read my article (all is explained inside), freely available at:
//        * arXiv: https://arxiv.org/abs/1407.4374
//        * HAL:   https://hal.archives-ouvertes.fr/hal-01024788
// 2) read the following comments
// 3) replace the content of this file with your own stuff
// 4) run (tested with Go version go1.6.3 linux/amd64 under Arch Linux):
//        go run example.go
//    it is possible that the Go package has a different name depending on your
//    OS/Linux distribution
//    for example, with Ubuntu, it is named "golang", so it may be
//    "golang-go run yourfile.go" instead of "go run yourfile.go"
// 5) check the help proposed by kali at runtime

// This example is a Boolean model of the ErbB receptor-regulated G1/S
// transition by Ozgur Sahin and colleagues [1].

//############################################################################//

package main
// Import kali, change the path if you move it.
import "./kali"
func main() {
    var (
        ntarg,maxtarg,maxmoda,maxS,kmax,threshold,sync int
        nodes []string
        vals kali.Vector
    )
    // nodes: the node names
    //     * nodes is an array of at least one string
    nodes=[]string{
        "ERBB1",
        "ERBB2",
        "ERBB3",
        "ERBB1_2",
        "ERBB1_3",
        "ERBB2_3",
        "IGF1R",
        "ER_alpha",
        "c_MYC",
        "AKT1",
        "MEK1",
        "CDK2",
        "CDK4",
        "CDK6",
        "Cyclin_D1",
        "Cyclin_E1",
        "p21",
        "p27",
        "pRB",
    }
    // vals: the domain of value
    //     * vals is an array of at least two real numbers in [0;1]
    //     * {0,1} for Boolean logic or for example {0,0.5,1} for three-valued
    //       logic
    vals=kali.Vector{0.0,1.0}
    // sync: the updating scheme
    //     * sync is an integer in {0,1}
    //     * sync=0: an asynchronous updating scheme is used (one randomly-
    //       selected variable is updated at each iteration following a uniform
    //       distribution)
    //     * sync=1: a synchronous updating scheme is used (all the variables
    //       are updated simultaneously at each iteration)
    //     * can be changed at runtime
    sync=0
    // maxS: the maximum number of initial states
    //     * maxS is an integer > 0
    //     * maxS is the maximum number of initial states to test when computing
    //       an attractor set
    //     * if it exceeds its maximal possible value then kali will
    //       automatically decrease it to its maximal possible value
    //     * can be changed at runtime
    maxS=1000
    // kmax: the number of iterations performed during a random walk
    //     * only relevant in the asynchronous case
    //     * kmax is an integer > 0 (recommended to be > 1000)
    //     * when searching for an attractor along an asynchronous updating
    //       scheme, a long random walk is performed in order to reach an
    //       attractor with a high probability (this candidate attractor will
    //       then be validated, or not)
    //     * the smallest is kmax the smallest is the probability to reach an
    //       attractor: this will cause kali to run for a too long time
    //     * on the other hand, if kmax is too big then kali will also run for a
    //       too long time
    //     * a good compromise could be 1000 < kmax < 10000
    //     * can be changed at runtime
    kmax=1000
    // ntarg: the number of targets per bullet
    //     * ntarg is an integer in [1;number of nodes]
    //     * can be changed at runtime
    ntarg=1
    // maxtarg: the maximum number of target combinations to test
    //     * maxtarg is an integer > 0
    //     * if it exceeds its maximal possible value then kali will
    //       automatically decrease it to its maximal possible value
    //     * can be changed at runtime
    maxtarg=100
    // maxmoda: the maximum number of modality arrangements to test
    //     * maxmoda is an integer > 0
    //     * maxmoda is the maximum number of modality arrangements to test for
    //       each target combination
    //     * if it exceeds its maximal possible value then kali will
    //       automatically decrease it to its maximal possible value
    //     * can be changed at runtime
    maxmoda=100
    // threshold: the threshold for a bullet to be considered therapeutic
    //     * threshold is an integer in [0;100]
    //     * the goal of therapeutic bullets is to increase the coverage of the
    //       pathological state space by the physiological one
    //     * to be therapeutic, this increase must be >= threshold (in percents
    //       of the pathological state space)
    //     * can be changed at runtime
    threshold=5
    kali.DoTheJob(fphysio,fpatho,ntarg,maxtarg,maxmoda,maxS,kmax,threshold,sync,nodes,vals)
}

// To cope with both Boolean and multivalued logic, the Zadeh fuzzy logic
// operators are used:
//     x AND y = min(x,y)
//     x OR y  = max(x,y)
//     NOT x   = 1-x

// fphysio: the transition function of the physiological variant
//     * fphysio is a vector function from vals^{number of nodes} to itself
func fphysio(x kali.Vector) kali.Vector {
    return kali.Vector{
        // replace the following equations with your own stuff
        // your equations encoded in the same way
        // note that the variable numbering starts at 0
        1.0,// ERBB1 (also stands for the homodimer)
        1.0,// ERBB2 (also stands for the homodimer)
        1.0,// ERBB3 (also stands for the homodimer)
        kali.Min(x[0],x[1]),// ERBB1:2
        kali.Min(x[0],x[2]),// ERBB1:3
        kali.Min(x[1],x[2]),// ERBB2:3
        kali.Min(kali.Max(x[7],x[9]),1.0-x[5]),// IGF1R
        kali.Max(x[9],x[10]),// ER alpha
        kali.Max(x[9],x[10],x[7]),// cMYC
        kali.Max(x[0],x[3],x[4],x[5],x[6]),// Akt1
        kali.Max(x[0],x[3],x[4],x[5],x[6]),// MEK1
        kali.Min(x[15],1.0-x[16],1.0-x[17]),// CDK2
        kali.Min(x[14],1.0-x[16],1.0-x[17]),// CDK4
        x[14],// CDK6
        kali.Min(x[7],x[8],kali.Max(x[9],x[10])),// Cyclin D1
        x[8],// Cyclin E1
        kali.Min(x[7],1.0-x[9],1.0-x[8],1.0-x[12]),// p21
        kali.Min(x[7],1.0-x[12],1.0-x[11],1.0-x[9],1.0-x[8]),// p27
        kali.Min(x[12],x[13]),// pRB
    }
}

// fpatho: the transition function of the pathological variant
//     * fpatho is a vector function from vals^{number of nodes} to itself
//     * in this example, fpatho is obtained by knocking out Akt1, but the
//       authors have tested a wide range of knockouts
func fpatho(x kali.Vector) kali.Vector {
    return kali.Vector{
        // replace the following equations with your own stuff
        // your equations encoded in the same way
        // note that the variable numbering starts at 0
        1.0,// ERBB1 (also stands for the homodimer)
        1.0,// ERBB2 (also stands for the homodimer)
        1.0,// ERBB3 (also stands for the homodimer)
        kali.Min(x[0],x[1]),// ERBB1:2
        kali.Min(x[0],x[2]),// ERBB1:3
        kali.Min(x[1],x[2]),// ERBB2:3
        kali.Min(kali.Max(x[7],x[9]),1.0-x[5]),// IGF1R
        kali.Max(x[9],x[10]),// ER alpha
        kali.Max(x[9],x[10],x[7]),// cMYC
        0.0,// Akt1
        kali.Max(x[0],x[3],x[4],x[5],x[6]),// MEK1
        kali.Min(x[15],1.0-x[16],1.0-x[17]),// CDK2
        kali.Min(x[14],1.0-x[16],1.0-x[17]),// CDK4
        x[14],// CDK6
        kali.Min(x[7],x[8],kali.Max(x[9],x[10])),// Cyclin D1
        x[8],// Cyclin E1
        kali.Min(x[7],1.0-x[9],1.0-x[8],1.0-x[12]),// p21
        kali.Min(x[7],1.0-x[12],1.0-x[11],1.0-x[9],1.0-x[8]),// p27
        kali.Min(x[12],x[13]),// pRB
    }
}

// [1] O. Sahin, H. Frohlich, C. Lobke, U. Korf, S. Burmester, M. Majety,
// J. Mattern, I. Schupp, C. Chaouiya, D. Thieffry, A. Poustka, S. Wiemann,
// T. Beissbarth, D. Arlt (2009) Modeling ERBB receptor-regulated G1/S
// transition to find novel targets for de novo trastuzumab resistance.
// BMC Systems Biology 3(1).
