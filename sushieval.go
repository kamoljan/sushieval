package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	m := make(map[string]int)
	i := 0
	for s.Scan() {
		i = i + 1
		m[s.Text()] = 1
	}

	fmt.Println(flag.Arg(0))
	fmt.Printf("Total      : %d\n", i)
	fmt.Printf("Unique     : %d\n", len(m))
	fmt.Printf("Duplicates : %d, %d%%\n", i-len(m), 100-len(m)*100/i)

}
