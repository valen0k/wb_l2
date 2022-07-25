package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	k             int
	n, r, u       bool
	input, output string
}

func NewFlags() (*flags, error) {
	fl := &flags{}
	flag.IntVar(&fl.k, "k", 0, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&fl.n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&fl.r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&fl.u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if len(flag.Args()) != 2 {
		return nil, errors.New("need 2 files for input and output")
	}

	fl.input = flag.Arg(0)
	fl.output = flag.Arg(1)

	return fl, nil
}

func main() {
	f, err := NewFlags()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := readFile(f.input)
	if err != nil {
		log.Fatalln(err)
	}

	if f.u {
		file = checkUnique(file)
	}

	sliceFile := make([][]string, 0, len(file))

	for _, val := range file {
		sliceFile = append(sliceFile, strings.Split(val, " "))
	}

	if f.n {
		sort.Slice(sliceFile, func(i, j int) bool {
			first, _ := strconv.ParseFloat(checkInfo(sliceFile, i, f.k), 64)
			second, _ := strconv.ParseFloat(checkInfo(sliceFile, j, f.k), 64)
			if f.r {
				return first > second
			}
			return first < second
		})
	} else {
		sort.Slice(sliceFile, func(i, j int) bool {
			if f.r {
				return checkInfo(sliceFile, i, f.k) > checkInfo(sliceFile, j, f.k)
			}
			return checkInfo(sliceFile, i, f.k) < checkInfo(sliceFile, j, f.k)
		})
	}

	if err = writeFile(f.output, sliceFile); err != nil {
		log.Fatalln(err)
	}
}

func checkInfo(info [][]string, i, j int) string {
	if j < len(info[i]) {
		return info[i][j]
	}
	return ""
}

func checkUnique(file []string) []string {
	set, res := make(map[string]struct{}), make([]string, 0)
	for _, val := range file {
		if _, ok := set[val]; !ok {
			res = append(res, val)
			set[val] = struct{}{}
		}
	}
	return res
}

func readFile(name string) ([]string, error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(file), "\n")

	return split, nil
}

func writeFile(name string, info [][]string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	for _, val := range info {
		_, err = file.WriteString(strings.Join(val, " ") + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
