package main

import (
	"flag"
	"fmt"
)

func main() {
	run()
}

func run() {
	parallelNumber := flag.Int("p", 1, "pはparallelの略で並列数を示してるんだぞい。つまり整数や！")
	flag.Parse()
	fmt.Printf("execute by %d parallel.\n", *parallelNumber)
}
