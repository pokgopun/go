/* https://theweeklychallenge.org/blog/perl-weekly-challenge-166/

Task 2: K-Directory Diff

Submitted by: [46]Ryan J Thompson
     __________________________________________________________________

   Given a few (three or more) directories (non-recursively), display a
   side-by-side difference of files that are missing from at least one of
   the directories. Do not display files that exist in every directory.

   Since the task is non-recursive, if you encounter a subdirectory,
   append a /, but otherwise treat it the same as a regular file.

Example

   Given the following directory structure:
dir_a:
Arial.ttf  Comic_Sans.ttf  Georgia.ttf  Helvetica.ttf  Impact.otf  Verdana.ttf
Old_Fonts/

dir_b:
Arial.ttf  Comic_Sans.ttf  Courier_New.ttf  Helvetica.ttf  Impact.otf  Tahoma.tt
f  Verdana.ttf

dir_c:
Arial.ttf  Courier_New.ttf  Helvetica.ttf  Impact.otf  Monaco.ttf  Verdana.ttf

   The output should look similar to the following:
dir_a          | dir_b           | dir_c
-------------- | --------------- | ---------------
Comic_Sans.ttf | Comic_Sans.ttf  |
               | Courier_New.ttf | Courier_New.ttf
Georgia.ttf    |                 |
               |                 | Monaco.ttf
Old_Fonts/     |                 |
               | Tahoma.ttf      |

go run ch-2.go ../../../challenge-001 ../../../challenge-002 ../../../challenge-003
../../../challenge-001|../../../challenge-002|../../../challenge-003|
----------------------|----------------------|----------------------|
abigail/              |abigail/              |abigail/              |
adam-russell/         |adam-russell/         |adam-russell/         |
ailbhe-tweedie/       |ailbhe-tweedie/       |ailbhe-tweedie/       |
alex-daniel/          |alex-daniel/          |alex-daniel/          |
alexander-karelas/    |alexander-karelas/    |alexander-karelas/    |
alexander-pankoff/    |alexander-pankoff/    |alexander-pankoff/    |
alexey-melezhik/      |alexey-melezhik/      |alexey-melezhik/      |
andrezgz/             |andrezgz/             |andrezgz/             |
antonio-gamiz/        |antonio-gamiz/        |antonio-gamiz/        |
arne-sommer/          |arne-sommer/          |arne-sommer/          |
arpad-toth/           |arpad-toth/           |arpad-toth/           |
athanasius/           |athanasius/           |athanasius/           |
aubrey-quarcoo/       |aubrey-quarcoo/       |aubrey-quarcoo/       |
                      |                      |bill-palmer/          |
bob-kleemann/         |bob-kleemann/         |bob-kleemann/         |
bob-lied/             |bob-lied/             |bob-lied/             |
                      |                      |cliveholloway/        |
conor-hoekstra/       |                      |                      |
daniel-mantovani/     |daniel-mantovani/     |daniel-mantovani/     |
dave-cross/           |dave-cross/           |dave-cross/           |
dave-jacoby/          |dave-jacoby/          |dave-jacoby/          |
david-kayal/          |david-kayal/          |david-kayal/          |
deadmarshal/          |                      |                      |
*/
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args,
			"../../../challenge-001",
			"../../../challenge-002",
			"../../../challenge-003",
			"../../../challenge-012",
			"../../../challenge-023",
			"../../../challenge-123")
	}
	var ds struct {
		dirNum     int
		name       string
		names      []string
		name2dir   map[string]int
		nameLen    int
		maxNameLen int
		ok         bool
	}
	ds.name2dir = map[string]int{}
	ds.dirNum = len(os.Args) - 1
	i := 1
	for _, v := range os.Args[1:] {
		ds.nameLen = len(v)
		if ds.maxNameLen < ds.nameLen {
			ds.maxNameLen = ds.nameLen
		}
		d, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		c, err := d.ReadDir(-1)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range c {
			ds.name = v.Name()
			if v.IsDir() {
				ds.name += "/"
			}
			_, ds.ok = ds.name2dir[ds.name]
			if !ds.ok {
				ds.names = append(ds.names, ds.name)
				ds.nameLen = len(ds.name)
				if ds.maxNameLen < ds.nameLen {
					ds.maxNameLen = ds.nameLen
				}
			}
			ds.name2dir[ds.name] += i
		}
		i *= 2
	}
	sort.Strings(ds.names)
	w := tabwriter.NewWriter(os.Stdout, ds.maxNameLen, ds.maxNameLen, 0, ' ', tabwriter.Debug)
	w.Write([]byte(strings.Join(os.Args[1:], "\t") + "\t\n"))
	w.Flush()
	w.Init(os.Stdout, ds.maxNameLen, ds.maxNameLen, 0, '-', tabwriter.Debug)
	w.Write([]byte(strings.Repeat("\t", ds.dirNum) + "\n"))
	w.Flush()
	w.Init(os.Stdout, ds.maxNameLen, ds.maxNameLen, 0, ' ', tabwriter.Debug)
	for _, v := range ds.names {
		flag := []byte(fmt.Sprintf("%0[1]*[2]b", ds.dirNum, ds.name2dir[v]))
		sort.SliceStable(flag, func(i, j int) bool {
			return true
		})
		for _, b := range flag {
			if b > 48 {
				w.Write([]byte(v))
			}
			w.Write([]byte("\t"))
		}
		w.Write([]byte("\n"))
	}
	w.Flush()
}
