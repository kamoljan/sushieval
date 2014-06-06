package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func totalDuplicates(file string) map[string]int {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	m := make(map[string]int)
	i := 0
	for s.Scan() {
		i = i + 1
		_, ok := m[s.Text()]
		if !ok {
			m[s.Text()] = 1
		} else {
			m[s.Text()] = m[s.Text()] + 1
		}
	}

	fmt.Println(flag.Arg(0))
	fmt.Printf("Total      : %d\n", i)
	fmt.Printf("Unique     : %d\n", len(m))
	fmt.Printf("Duplicates : %d, %d%%\n", i-len(m), 100-len(m)*100/i)

	return m
}

func plotDuplicates(m map[string]int, file string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	i := 0
	for _, v := range m {
		i = i + 1
		if _, err = f.WriteString(fmt.Sprintf("%d,%d\n", i, v)); err != nil {
			panic(err)
		}
	}
}

func main() {
	flag.Parse()
	m := totalDuplicates(flag.Arg(0))
	plotDuplicates(m, flag.Arg(0)+".plot")
}
