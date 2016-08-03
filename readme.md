# Therapeutic target discovery using Boolean network attractors: avoiding pathological phenotypes

Copyright (C) 2013-2016 [Arnaud Poret](https://github.com/arnaudporet)

This work is licensed under the [GNU General Public License](https://www.gnu.org/licenses/gpl.html).

## How to

1. read my article (all is explained inside), freely available on [arXiv](https://arxiv.org) and [HAL](https://hal.archives-ouvertes.fr):
    * https://arxiv.org/abs/1407.4374
    * https://hal.archives-ouvertes.fr/hal-01024788
2. clone kali (or [download](https://github.com/arnaudporet/kali/archive/master.zip) it if you do not use [Git](https://git-scm.com)):
    * `git clone https://github.com/arnaudporet/kali.git`
3. open the file `example.go`

kali is tested with [Go](https://golang.org) version go1.6.3 linux/amd64 under [Arch Linux](https://www.archlinux.org).

How to get Go: https://golang.org/doc/install

Most [Linux distributions](https://distrowatch.com) provide Go in their official repositories. For example:
* golang ([Ubuntu](http://www.ubuntu.com))
* go ([Arch Linux](https://www.archlinux.org))

# Forthcoming

* updating the article (currently the article does not cover the asynchronous updating scheme)
* other examples and case studies

# References

The example is a published Boolean model of the ErbB receptor-regulated G1/S transition [1].

The algorithms used to handle the asynchronous updating scheme are adapted from those of [BoolNet](https://cran.r-project.org/web/packages/BoolNet/index.html) [2].

1. O. Sahin, H. Frohlich, C. Lobke, U. Korf, S. Burmester, M. Majety, J. Mattern, I. Schupp, C. Chaouiya, D. Thieffry, A. Poustka, S. Wiemann, T. Beissbarth, D. Arlt (2009) Modeling ERBB receptor-regulated G1/S transition to find novel targets for de novo trastuzumab resistance. BMC Systems Biology 3(1).

2. C. Mussel, M. Hopfensitz, H. A. Kestler (2010) BoolNet-an R package for generation, reconstruction and analysis of Boolean networks. Bioinformatics 26(10).
