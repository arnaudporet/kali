# Biological case study: bladder tumorigenesis

This folder contains a biological case study in order to address a concrete case, namely a published logic-based model of bladder tumorigenesis.

The technical details about using kali are not recalled: the example contained in the file `example.go` should be consulted first.

## `readme.md`

The file you are reading.

## `bladder_tumorigenesis.go`

This file contains the case study.

Run `go run bladder_tumorigenesis.go` in a terminal emulator.

## `equations/`

This folder contains the model equations of the case study in different versions.

It also contains a [Python](https://www.python.org) script to generate the version of the model equations present in the file `bladder_tumorigenesis.go`.

All the equations are, and must be, as follows: `equation// node_name`.

### `eq_bool.go`

This file contains the Boolean version of the model equations.

### `eq_func.go`

This file contains the functional version of the model equations: the Zadeh operators are used instead of the Boolean ones.

kali only uses the functional version of the model equations since it works with both Boolean and multivalued logic.

### `xEq.py`

This script takes the file containing the functional version of the model equations, here the file `eq_func.go`, and returns the file `eq_func_xed.go`.

Run `python eq_func.go` in a terminal emulator.

This script is coded in Python3, no warranties that it works with Python2.

### `eq_func_xed.go`

This file contains the functional version of the model equations contained in the file `eq_func.go` but with the node names replaced by the corresponding positions in the state vector **x**.

This is the version of the model equations present in the file `bladder_tumorigenesis.go`.

It also contains the list of the nodes names.

The content of this file is intended to be directly pasted into the file `bladder_tumorigenesis.go` at the appropriate places.

### `readouts.go`

This file contains the equations of the three output phenotypes.

Use them to evaluate these output phenotypes from the returned attractors once the run terminated.

## `results/`

This folder contains the results of this case study obtained by using the file `bladder_tumorigenesis.go`.

### `A_physio.txt`

This file contains the physiological attractor set.

### `A_patho.txt`

This file contains the pathological attractor set.

### `A_versus.txt`

This file contains the pathological attractors.

### `B_therap_1.txt`

This file contains the therapeutic bullets made of one target.

### `B_therap_2.txt`

This file contains the therapeutic bullets made of two targets.
