// Copyright (C) 2013-2019 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
func DoTestBullets(n,ntarg,maxtarg,maxmoda int,vals Vector) {
    IntToVect(Range(0,n)).Combis(ntarg,maxtarg).Save("Targ.csv")
    vals.Arrangs(ntarg,maxmoda).Save("Moda.csv")
}
