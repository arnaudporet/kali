################################################################################
########    An in silico target identification using boolean network    ########
########          attractors: avoiding pathological phenotypes          ########
################################################################################

################################    LICENSE    #################################

Copyright (C) 2013-2015 Arnaud Poret

This program is licensed under the GNU General Public License.

To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html or
see the GNU_General_Public_License.txt file in the present folder.

################################    CONTACT    #################################

Arnaud Poret: arnaud.poret@gmail.com

If you are interested in my work, or ultimately if you use it, feel free to let
me know.

#######################    RECOMMENDED PREREQUISITES    ########################

Operating system: Linux, such as Ubuntu (http://www.ubuntu.com/)

Compiler: GFortran (https://gcc.gnu.org/fortran/)

Text editor: provided that it handle Fortran syntax highlighting, such as Gedit
(https://wiki.gnome.org/Apps/Gedit)

Distributed version control system: Git (http://git-scm.com/)

################################    VERSIONS    ################################

* 1.1.0: see the corresponding changelog

* 2.0.0: see the corresponding changelog

#################################    HOWTO    ##################################

1) read my article: freely available (and more readable than the published
   version) at http://arxiv.org/abs/1407.4374 or
   https://hal.archives-ouvertes.fr/hal-01024788

2) clone kali-targ:
       git clone https://github.com/arnaudporet/kali-targ.git
   or download it if you do not use Git:
       https://github.com/arnaudporet/kali-targ/archive/master.zip
   but I strongly encourage you to use Git

3) go to the folder corresponding to the version of your choice

4) read the corresponding changelog

5) open the template file example_network.f08

################################    CREDITS    #################################

The example network is a published Boolean model of the mammalian cell cycle:

    Faure A., Naldi A., Chaouiya C., Thieffry D. (2006) Dynamical analysis of a
    generic Boolean model for the control of the mammalian cell cycle.
    Bioinformatics 22(14) e124-e131.
