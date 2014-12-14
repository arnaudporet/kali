################################################################################
########    An in silico target identification using boolean network    ########
########          attractors: avoiding pathological phenotypes          ########
################################################################################

Copyright (C) 2013-2014 Arnaud Poret
This program is licensed under the GNU General Public License.
To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html or
see the GNU_General_Public_License.txt file in the present folder.

Contact:
    Arnaud Poret
    arnaud.poret@gmail.com

If you are interested in my work, or ultimately if you use it, feel free to let
me know.

Recommended prerequisites:
    * operating system: Linux, such as Ubuntu (http://www.ubuntu.com/)
    * compiler: GFortran (https://gcc.gnu.org/fortran/)
    * text editor: provided that it handle Fortran syntax highlighting, such as
      Gedit (https://wiki.gnome.org/Apps/Gedit)
    * distributed version control system: Git (http://git-scm.com/)

How to:
    1) read my article: freely available at http://arxiv.org/abs/1407.4374 or
       https://hal.archives-ouvertes.fr/hal-01024788
    2) clone kali-targ: git clone https://github.com/arnaudporet/kali-targ.git
       (or download it if you do not have Git, but I strongly encourage you to
       use Git)
    3) with a text editor, open the template file: example_network.f08

The example network is a published Boolean model of the mammalian cell
cycle [1].

[1] Faure A., Naldi A., Chaouiya C., Thieffry D. (2006) Dynamical analysis of a
    generic Boolean model for the control of the mammalian cell cycle.
    Bioinformatics 22(14) e124-e131.
