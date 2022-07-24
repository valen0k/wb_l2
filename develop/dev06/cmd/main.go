package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cut struct {
	fields, delimiter string
	separated         bool
	column            []int
}

func NewCut() (*cut, error) {
	c := cut{}
	flag.StringVar(&c.fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&c.delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&c.separated, "s", false, "только строки с разделителем")
	flag.Parse()

	if c.fields != "" {
		if err := c.checkColumn(); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

func (c *cut) checkColumn() error {
	split := strings.Split(c.fields, ",")
	c.column = make([]int, 0, len(split))

	for _, val := range split {
		valSplit := strings.Split(val, "-")
		if len(valSplit) == 1 {
			atoi, err := strconv.Atoi(valSplit[0])
			if err != nil {
				return err
			}
			c.column = append(c.column, atoi)
			continue
		}
		n1, err := strconv.Atoi(valSplit[0])
		if err != nil {
			return err
		}
		n2, err := strconv.Atoi(valSplit[1])
		if err != nil {
			return err
		}
		for ; n1 <= n2; n1++ {
			c.column = append(c.column, n1)
		}
	}
	sort.Ints(c.column)
	return nil
}

func (c *cut) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res := make([]string, 0)
		split := strings.Split(scanner.Text(), c.delimiter)
		for i, val := range split {
			if c.column == nil {
				res = append(res, val)
			} else if slices.Contains(c.column, i+1) {
				res = append(res, val)
			}
		}
		fmt.Println(strings.Join(res, " "))
	}
}

func main() {
	c, err := NewCut()
	if err != nil {
		log.Fatalln(err)
	}
	c.Start()
}
