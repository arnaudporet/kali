// Copyright (C) 2013-2016 Arnaud Poret
// This work is licensed under the GNU General Public License.
// To view a copy of this license, visit https://www.gnu.org/licenses/gpl.html
package kali
import (
    "fmt"
    "strings"
)
func DoNotice() {
    fmt.Println(strings.Join([]string{
        "\nkali: a tool for in silico therapeutic target discovery",
        "Copyright (C) 2013-2016 Arnaud Poret\n",
        "This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.\n",
        "This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.\n",
        "You should have received a copy of the GNU General Public License along with this program. If not, see https://www.gnu.org/licenses/gpl.html.",
    },"\n"))
}
