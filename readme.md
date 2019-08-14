# In silico therapeutic target discovery using network attractors

Copyright 2013-2019 [Arnaud Poret](https://github.com/arnaudporet)

This work is licensed under the [GNU General Public License](https://www.gnu.org/licenses/gpl.html).

## kali

kali is a tool for performing _in silico_ therapeutic target discovery using network attractors.

kali can currently operate on qualitative models, namely Boolean networks and multi-valued ones (a generalization of Boolean networks).

kali provides 5 commands, which should be run in that order:

1. `start`: generate the set of the start states
2. `attractor`: compute an attractor set
3. `versus`: get the set of the pathological attractors
4. `bullet`: generate the set of the bullets to test
5. `target`: compute a set of therapeutic bullets

kali is implemented in [Go](https://golang.org) (see at the end of this readme file).

## Publications

Arnaud Poret, Carito Guziolowski (2018) Therapeutic target discovery using Boolean network attractors: improvements of kali. Royal Society Open Science 5(2).

Arnaud Poret, Jean-Pierre Boissel (2014) An in silico target identification using Boolean network attractors: avoiding pathological phenotypes. Comptes Rendus Biologies 337(12).

## How to

1. read my [article](https://arxiv.org/pdf/1611.03144.pdf)
2. read this readme file
3. read the (template) file `example.go`
4. replace its content with your own stuff
5. run in a terminal emulator:

```sh
go run example.go -help
```
or

```sh
go build example.go
./example -help
```

Note that `go run` builds kali each time before running it: safer if the sources are frequently modified.

The Go package can have different names depending on your operating system. For example, with [Ubuntu](https://ubuntu.com) the Go package is named `golang`. Consequently, running a Go file with Ubuntu might be `golang-go run yourfile.go` instead of `go run yourfile.go` with [Arch Linux](https://www.archlinux.org). Otherwise, see https://golang.org/dl/ or https://golang.org/doc/install.

The (template) file `example.go` contains a simple and fictive Boolean network to conveniently illustrate kali. Next, you can replace its content with your own stuff.

A biological case study is also proposed in the `case_study` folder to address a concrete case (see below).

## Usage

In this section, it is assumed that the file `example.go` is used. It is the __main file__, meaning that it contains the `main` function together with your model, and that it imports kali in order to use it on your model.

You can name the main file as you wish, but avoid naming it "kali" in order to prevent confusion/conflict issues with kali itself. For example, in the `case_study` folder, the main file is named `bladder_tumorigenesis.go`.

### kali

Usage:

```
example [options]
example <command> [options] <arguments>
```

Commands (should be run in that order):

* `start`: generate the set of the start states
* `attractor`: compute an attractor set
* `versus`: get the set of the pathological attractors
* `bullet`: generate the set of the bullets to test
* `target`: compute a set of therapeutic bullets

Positional arguments: see the command-specific help (`example <command> -help`)

Options:

* non command-specific options:
    * `-l/-license`: print the GNU General Public License under which kali is
    * `-u/-usage`: print usage only
    * `-h/-help`: print help
* command-specific options: `example <command> -help`

Output files: see the command-specific help (`example <command> -help`)

Cautions:

* non command-specific cautions:
    * kali automatically saves and loads the files it creates and uses
    * if one of these files already exists then it is overwritten
    * if one of these files is renamed, moved or deleted then it can not be loaded when required
* command-specific cautions: `example <command> -help`

For command-specific help, run: `example <command> -help`

For full explanation, see the [article](https://arxiv.org/pdf/1611.03144.pdf).

### kali start

Generate the set of the start states.

The start states are the initial states to use when computing an attractor set.

Usage: `example start [options]`

Options:

* `-maxS`: the maximal number of start states (default: `1 000`)
* `-u/-usage`: print usage only
* `-h/-help`: print this help

Output file:

* `S.csv`

Cautions:

* if `S.csv` already exists then it is overwritten
* if `S.csv` is renamed, moved or deleted then it can not be loaded when required

`maxS`:

* recommended to be at least 1 000
* is the maximum number of initial states to use when computing an attractor set
* if it exceeds its maximal possible value then it is automatically decreased to its maximal possible value
* the smallest is `maxS` the smallest is the probability to find all the attractors
* if `maxS` is too big then kali will run for a too long time
* a compromise could be `maxS` in [1 000;100 000], depending on the size of the model (i.e. the size of the state space)
* the bigger is the state space the bigger should be `maxS`

### kali attractor

Compute an attractor set:

* `A_physio`: the attractor set of the physiological variant
* `A_patho`: the attractor set of the pathological variant

Usage: `example attractor [options] <setting>`

Positional arguments:

* `<setting>`: compute the physiological attractor set (`physio`) or the pathological attractor set (`patho`)

Options:

* `-nstep`: the number of steps performed during a random walk when searching for an attractor (default: `1 000`)
* `-maxtry`: the maximum number of random walks performed when searching for an attractor (default: `10`)
* `-upd`: the updating method to use (default: `async`)
* `-u/-usage`: print usage only
* `-h/-help`: print this help

Output files:

* if computing the physiological attractor set:
    * `A_physio.csv`
    * `A_physio.txt`
* if computing the pathological attractor set:
    * `A_patho.csv`
    * `A_patho.txt`

Cautions:

* if `A_physio.csv`, `A_physio.txt`, `A_patho.csv`, `A_patho.txt` already exist then they are overwritten
* if `A_physio.csv`, `A_physio.txt`, `A_patho.csv`, `A_patho.txt` are renamed, moved or deleted then they can not be loaded when required

`nstep`:

* only relevant in the asynchronous case
* recommended to be at least 1 000
* when searching for an attractor according to an asynchronous updating, a long random walk is performed in order to reach an attractor with high probability (this candidate attractor is then subjected to validation)
* the smallest is `nstep` the smallest is the probability to reach an attractor: this will cause kali to run for a too long time
* if `nstep` is too big then kali will also run for a too long time
* a compromise could be `nstep` in [1 000;10 000], depending on the size of the model (i.e. the size of the state space)
* the bigger is the state space the bigger should be `nstep`

`maxtry`:

* only relevant in the asynchronous case
* if `nstep` is too small regarding the size of the state space, it is possible that no attractors can be reached from a given start state (i.e. where starts a random walk in the state space)
* to prevent looping indefinitely, a maximum of `maxtry` random walks are performed from each start state

`upd`:

* `async`: an asynchronous updating is used, meaning that one randomly selected variable is updated at each iteration of the simulation
* `sync`: a synchronous updating is used, meaning that all the variables are updated simultaneously at each iteration of the simulation

The basins of the attractors are expressed in percents of the state space (note that it is an estimation).

The csv files are for kali while the txt files are for you.

### kali versus

Get the set of the pathological attractors.

Usage: `example versus [options]`

Option:

* `-u/-usage`: print usage only
* `-h/-help`: print this help

Output files:

* `A_versus.csv`
* `A_versus.txt`

Cautions:

* `A_versus` is not the attractor set of the pathological variant (i.e. `A_patho`): it is the set containing the pathological attractors, namely the attractors specific to `A_patho` (`A_patho` can also contain physiological attractors)
* if `A_versus.csv`, `A_versus.txt` already exist then they are overwritten
* if `A_versus.csv`, `A_versus.txt` are renamed, moved or deleted then they can not be loaded when required

The basins of the attractors are expressed in percents of the state space (note that it is an estimation).

The csv files are for kali while the txt files are for you.

### kali bullet

Generate the set of the bullets to test:

* `Targ`: the set of the target node combinations to test
* `Moda`: the set of the modality arrangements (i.e. the state modifications) to apply on each of the target node combinations

The modalities are the perturbations to apply on the target nodes, typically inhibitions or stimulations.

A bullet is a couple made of:

* one combination without repetition of `ntarg` target nodes
* one arrangement with repetition of `ntarg` modalities
* ex: `((node1,node2),(moda1,moda2))` where the modality `moda1` is for the node `node1` and the modality `moda2` for the node `node2`, and `ntarg=2`

Usage: `example bullet [options]`

Options:

* `-ntarg`: the number of target nodes per bullet (default: `1`)
* `-maxtarg`: the maximum number of target node combinations to test (default: `100`)
* `-maxmoda`: the maximum number of modality arrangements to test for each target node combination (default: `100`)
* `-u/-usage`: print usage only
* `-h/-help`: print this help

Output files:

* `Targ.csv`
* `Moda.csv`

Cautions:

* if `Targ.csv`, `Moda.csv` already exist then they are overwritten
* if `Targ.csv`, `Moda.csv` are renamed, moved or deleted then they can not be loaded when required

`maxmoda` is the maximum number of modality arrangements to test for each target node combination: there are maximum `maxtarg*maxmoda` bullets to test.

If `maxtarg` and/or `maxmoda` exceeds its maximal possible value then it is automatically decreased to its maximal possible value.

### kali target

Compute a set of therapeutic bullets.

Usage: `example target [options]`

Options:

* `-nstep`: the number of steps performed during a random walk when searching for an attractor (default: `1 000`)
* `-maxtry`: the maximum number of random walks performed when searching for an attractor (default: `10`)
* `-upd`: the updating method to use (default: `async`)
* `-th`: the threshold for a bullet to be considered therapeutic (default: `5`)
* `-u/-usage`: print usage only
* `-h/-help`: print this help

Output file:

* `B_therap.txt`

Cautions:

* `nstep`, `maxtry` and `upd` should be the same as the attractor set computation step
* if `B_therap.txt` already exists then it is overwritten

`nstep`:

* only relevant in the asynchronous case
* recommended to be at least 1 000
* when searching for an attractor according to an asynchronous updating, a long random walk is performed in order to reach an attractor with high probability (this candidate attractor is then subjected to validation)
* the smallest is `nstep` the smallest is the probability to reach an attractor: this will cause kali to run for a too long time
* if `nstep` is too big then kali will also run for a too long time
* a compromise could be `nstep` in [1 000;10 000], depending on the size of the model (i.e. the size of the state space)
* the bigger is the state space the bigger should be `nstep`

`maxtry`:

* only relevant in the asynchronous case
* if `nstep` is too small regarding the size of the state space, it is possible that no attractors can be reached from a given start state (i.e. where starts a random walk in the state space)
* to prevent looping indefinitely, a maximum of `maxtry` random walks are performed from each start state

`upd`:

* `async`: an asynchronous updating is used, meaning that one randomly selected variable is updated at each iteration of the simulation
* `sync`: a synchronous updating is used, meaning that all the variables are updated simultaneously at each iteration of the simulation

`th`:

* is expressed in percents of the state space
* the goal of therapeutic bullets is to increase the physiological basins (i.e. to increase the union of the basins of the physiological attractors)
* to be therapeutic, this increase must be superior or equal to `th` (in percents of the state space)

Therapeutic bullets are reported as follows:

```
Bullet: x1[y1] x2[y2] ...
Gain: U1 --> U2
Physiological basins:
    a_physio1: b_physio1
    a_physio2: b_physio2
    ...
Pathological basins:
    a_patho1: b_patho1
    a_patho2: b_patho2
    ...
```

where:

* `x1[y1] x2[y2] ...` means that the node `x1` has to be set to the value `y1`, the node `x2` to the value `y2`, and so on
* `U1` is the union of the physiological basins in the state space of the pathological variant (in percents of it)
* `U2` is the union of the physiological basins in the state space of the pathological variant __subjected to the effect of the bullet__ (in percents of it)
* therefore, the gain of the bullet is the increase from `U1` to `U2` (this is a condition for the bullet to be considered therapeutic: `U2-U1>=th`)
* the basin of the physiological and pathological attractors (`b_physio` and `b_patho` respectively) in the state space of the pathological variant __subjected to the effect of the bullet__ are expressed in percents of it

## Case study: bladder tumorigenesis

The folder `case_study` contains a biological case study in order to address a concrete case, namely a published logic-based model of bladder tumorigenesis.

See its own `readme.md` file.

## Forthcoming

## Credit

The functions used to handle the asynchronous updating are adapted from [BoolNet](https://cran.r-project.org/web/packages/BoolNet/) [1].

The biological case study is adapted from a published logic-based model of bladder tumorigenesis [2].

1. Christoph Mussel, Martin Hopfensitz, Hans Kestler (2010) BoolNet-an R package for generation, reconstruction and analysis of Boolean networks. Bioinformatics 26(10).
2. Elisabeth Remy, Sandra Rebouissou, Claudine Chaouiya, Andrei Zinovyev, Francois Radvanyi, Laurence Calzone (2015) A modeling approach to explain mutually exclusive and co-occurring genetic alterations in bladder tumorigenesis. Cancer Research 75(19).

## Go

Most [Linux distributions](https://distrowatch.com) provide Go in their official repositories. For example:

* `go` (Arch Linux)
* `golang` (Ubuntu)

Otherwise, see https://golang.org/dl/ or https://golang.org/doc/install
