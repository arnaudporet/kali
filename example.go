// Copyright (C) 2013-2017 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html

//#### HOWTO #################################################################//

// 1) read my article, freely available at: https://arxiv.org/abs/1611.03144
// 2) read the following comments
// 3) replace the content of this file with your own stuff
// 4) run (in a terminal emulator): ``go run example.go''
// 5) check the help proposed by kali at runtime

// It is possible that the Go package has a different name depending on your
// operating system. For example, with Ubuntu it is named golang, so it might be
// ``golang-go run yourfile.go'' instead of ``go run yourfile.go'' with Arch
// Linux.

// This example is a fictive network. It is specifically designed for
// illustration and is not intended to model a real biological phenomenon.

//############################################################################//

package main
// import kali, change the path if you move it
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
    // vals: the domain of value
    //     * vals is an array of at least two real numbers in [0;1]
    //     * {0,1} for Boolean logic or, for example, {0,0.5,1} for three-valued
    //       logic
    vals=kali.Vector{0.0,1.0}
    // sync: the updating method for the variables
    //     * sync is an integer in {0,1}
    //     * sync=0: an asynchronous updating is used (one randomly selected
    //       variable is updated at each iteration, according to a uniform
    //       distribution)
    //     * sync=1: a synchronous updating is used (all the variables are
    //       updated simultaneously at each iteration)
    //     * can be changed at runtime
    sync=0
    // maxS: the maximum number of initial states to use
    //     * maxS is an integer > 0
    //     * maxS is the maximum number of initial states to use when computing
    //       an attractor set
    //     * if it exceeds its maximal possible value then kali will
    //       automatically decrease it to its maximal possible value
    //     * can be changed at runtime
    maxS=1000
    // kmax: the number of iterations performed during a random walk
    //     * only relevant in the asynchronous case
    //     * kmax is an integer > 0 (recommended to be > 1000)
    //     * when searching for an attractor according to an asynchronous
    //       updating, a long random walk is performed in order to reach an
    //       attractor with high probability (this candidate attractor will then
    //       be subjected to validation)
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
    //     * the goal of therapeutic bullets is to increase the basin of the
    //       physiological attractors in the pathological state space
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
        x[0],// do
        x[1],// factory
        kali.Max(kali.Min(x[2],1.0-x[8]),x[1]),// energy
        1.0-x[2],// locker
        x[0],// releaser
        1.0-x[4],// sequester
        kali.Min(x[0],1.0-x[3]),// activator
        kali.Min(x[6],1.0-x[5]),// effector
        x[7],// task
    }
}

// fpatho: the transition function of the pathological variant
//     * fpatho is a vector function from vals^{number of nodes} to itself
//     * in this example, fpatho is obtained by knocking down the locker
func fpatho(x kali.Vector) kali.Vector {
    return kali.Vector{
        // replace the following equations with your own stuff
        // your equations encoded in the same way
        // note that the variable numbering starts at 0
        x[0],// do
        x[1],// factory
        kali.Max(kali.Min(x[2],1.0-x[8]),x[1]),// energy
        0.0,// locker (knocked down)
        x[0],// releaser
        1.0-x[4],// sequester
        kali.Min(x[0],1.0-x[3]),// activator
        kali.Min(x[6],1.0-x[5]),// effector
        x[7],// task
    }
}
