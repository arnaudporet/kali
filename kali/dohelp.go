// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import "fmt"
import "strings"
//#### DoHelp ################################################################//
func DoHelp() {
    fmt.Println(strings.Join([]string{
        "\nHow to use the algorithm:",
        "    1) read my article (all is explained inside), freely available on:",
        "        * arXiv: https://arxiv.org/abs/1407.4374",
        "        * HAL: https://hal.archives-ouvertes.fr/hal-01024788",
        "    2) generate the state space (S)",
        "    3) compute the attractor set of the physiological variant (A_physio)",
        "        * when prompted, set the setting to physiological",
        "    4) compute the attractor set of the pathological variant (A_patho)",
        "        * when prompted, set the setting to pathological",
        "    5) compute the pathological attractors (A_versus)",
        "        * A_versus should not be confused with the attractor set of the pathological variant (A_patho)",
        "        * A_versus is not an attractor set: it is the set containing the pathological attractors",
        "    6) generate the bullets to test (Targ and Moda)",
        "    7) compute therapeutic bullets (B_therap)",
        "        * therapeutic bullets are reported as follow:",
        "              x1[y1] x2[y2] x3[y3] ...",
        "          meaning that the variable x has to be set to the value y",
        "    * you can change parameter values (ntarg, maxtarg, maxmoda and maxS)",
        "    * you can check what is already saved (S, A_physio, A_patho, A_versus, Targ, Moda and B_therap)\n",
        "The algorithm automatically save/load the files it creates/uses.\n",
        "These files are:",
        "    * S.csv",
        "    * A_physio.csv",
        "    * A_physio.txt",
        "    * A_patho.csv",
        "    * A_patho.txt",
        "    * A_versus.csv",
        "    * A_versus.txt",
        "    * Targ.csv",
        "    * Moda.csv",
        "    * B_therap.txt\n",
        "The csv files are for the algorithm while the txt files are for you.\n",
        "If a file already exists then the algorithm will overwrite it.\n",
        "If you rename, move or delete the csv files created by the algorithm then it will not be able to load them when required.\n",
        "The algorithm is tested with Go version 1.6.2 under Arch Linux.",
    },"\n"))
}
