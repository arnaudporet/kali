# Therapeutic target discovery using Boolean network attractors: avoiding pathological phenotypes

Copyright (C) 2013-2017 [Arnaud Poret](https://github.com/arnaudporet)

This work is licensed under the [GNU General Public License](https://www.gnu.org/licenses/gpl.html).

## How to

1. read my [article](https://arxiv.org/abs/1611.03144)
2. clone kali (or [download](https://github.com/arnaudporet/kali/archive/master.zip) it if you do not use [Git](https://git-scm.com)): `git clone https://github.com/arnaudporet/kali.git`
3. open the file `example.go`

The file `example.go` contains a simple and fictive Boolean network to conveniently illustrate kali.

The folder `case_study/` contains a biological case study in order to address a concrete case: a published logic-based model of bladder tumorigenesis.

The technical details about using kali are not recalled in the case study: the example should be consulted first.

kali is implemented in [Go](https://golang.org): https://golang.org/doc/install

Most [Linux distributions](https://distrowatch.com) provide Go in their official repositories. For example:
* golang ([Ubuntu](https://www.ubuntu.com))
* go ([Arch Linux](https://www.archlinux.org))

## Forthcoming

## References

The functions used to handle the asynchronous updating are adapted from [BoolNet](https://cran.r-project.org/web/packages/BoolNet/) [1].

The biological case study is adapted from a published logic-based model of bladder tumorigenesis [2].

1. Christoph Mussel, Martin Hopfensitz, Hans Kestler (2010) BoolNet-an R package for generation, reconstruction and analysis of Boolean networks. Bioinformatics 26(10):1378-1380.
2. Elisabeth Remy, Sandra Rebouissou, Claudine Chaouiya, Andrei Zinovyev, Francois Radvanyi, Laurence Calzone (2015) A modeling approach to explain mutually exclusive and co-occurring genetic alterations in bladder tumorigenesis. Cancer Research 75(19):4042-4052.
