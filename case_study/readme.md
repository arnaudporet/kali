# Biological case study: bladder tumorigenesis

This folder contains a biological case study in order to address a concrete case: a published logic-based model of bladder tumorigenesis [1].

## Content

* `readme.md`: the file you are reading
* `bladder_tumorigenesis.go`: the case study

```sh
go run bladder_tumorigenesis.go -help
```

### Equations

The `equations` folder contains the model's equations of the case study in different versions:

* `eq_bool.go`: the Boolean version of the model's equations
* `eq_func.go`: the functional version of the model's equations (the Zadeh's logical operators are used instead of the Boolean ones)
* `readouts.go`: the equations of the three output phenotypes, use them to evaluate these output phenotypes from the returned attractors once the run has terminated

All the equations are and must be as follows:

* `<equation>// <node_name>`
* ex:
    * `(RAS || AKT) && !p16INK4a && !p21CIP// CyclinD1` (Boolean version)
    * `min(max(RAS,AKT),1-p16INK4a,1-p21CIP)// CyclinD1` (functional version)

However, kali only uses the functional version of the model's equations because they work with both Boolean and multivalued logic.

To do so, in the main file `bladder_tumorigenesis.go`:

* the variable names are replaced by their corresponding position in the state vector `x`
* the variable names are stored in the list `nodes`, respecting the __same order__ as in `x`

### Results

The `results` folder contains the results of this case study:

* `A_physio.txt`: the physiological attractor set
* `A_patho.txt`: the pathological attractor set
* `A_versus.txt`: the pathological attractors (i.e. those specific to `A_patho`)
* `B_therap_1.txt`: the found therapeutic bullets made of __one__ target node
* `B_therap_2.txt`: the found therapeutic bullets made of __two__ target nodes

[1] Elisabeth Remy, Sandra Rebouissou, Claudine Chaouiya, Andrei Zinovyev, Francois Radvanyi, Laurence Calzone (2015) A modeling approach to explain mutually exclusive and co-occurring genetic alterations in bladder tumorigenesis. Cancer Research 75(19).
