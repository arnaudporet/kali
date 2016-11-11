// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "strings"
)
func DoHelp() {
    fmt.Println(strings.Join([]string{
        "\nHow to use kali:",
        "    1) read my article, freely available at: https://arxiv.org/abs/1611.03144",
        "    2) generate the state space (S)",
        "    3) compute the attractor set of the physiological variant (A_physio)",
        "        * when prompted, set the setting to physiological",
        "    4) compute the attractor set of the pathological variant (A_patho)",
        "        * when prompted, set the setting to pathological",
        "    5) get the pathological attractors (A_versus)",
        "    6) generate the bullets to test (Targ and Moda)",
        "    7) compute therapeutic bullets (B_therap)",
        "        * therapeutic bullets are reported as follow:",
        "              x1[y1] x2[y2] x3[y3] ...",
        "          meaning that the variable x has to be set to the value y",
        "    * you can change parameter values (ntarg, maxtarg, maxmoda, maxS, kmax, threshold, sync)",
        "    * you can check/clear what is saved (S, A_physio, A_patho, A_versus, Targ, Moda, B_therap)\n",
        "kali automatically saves/loads the files it creates/uses.\n",
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
        "The csv files are for kali while the txt files are for you.\n",
        "If a file already exists then kali will overwrite it.\n",
        "If you rename, move or delete the csv files created by kali then it will not be able to load them when required.\n",
    },"\n"))
}
