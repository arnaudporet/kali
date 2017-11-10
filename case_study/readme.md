# Biological case study: bladder tumorigenesis

## Content

* `readme.md`: the file you are reading
* `bladder_tumorigenesis.go`: this is the main file of the biological case study; run `go run bladder_tumorigenesis.go` in a terminal emulator
* `equations/`: this folder contains the model equations
    * `eq_bool.go`: this file contains the Boolean version of the model equations
    * `eq_func.go`: this file contains the functional version of the model equations, namely using the Zadeh operators instead of the Boolean ones
    * `xEq.py`: this script takes the file `eq_func.go` and returns the file `eq_func_xed.go` (see below); run `python eq_func.go` in a terminal emulator
    * `eq_func_xed.go`: this file contains the equations of the file `eq_func.go` but with the node names replaced by the corresponding positions in the state vector x; also contains the list of the nodes names; the content of this file is intended to be directly pasted in the main file `bladder_tumorigenesis.go` at the appropriate places
    * `readouts.go`: this file contains the equations of three output phenotypes of the model; use them to evaluate these output phenotypes from the returned attractors once the run terminated
* `results/`: this folder contains the results; the used logic is the Boolean one; the used updating is the asynchronous one
    * `A_physio.txt`: this file contains the physiological attractor set
    * `A_patho.txt`: this file contains the pathological attractor set
    * `A_versus.txt`: this file contains the pathological attractors
    * `B_therap_1.txt`: this file contains the therapeutic bullets made of one target
    * `B_therap_2.txt`: this file contains the therapeutic bullets made of two targets
