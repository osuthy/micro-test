package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	os.Args = []string{"osargstrial.exe", "-p", "5"} // stub command line args
	run()
}

func run() {
	parallelNumber := flag.Int("p", 1, "pはparallelの略で並列数を示してるんだぞい。つまり整数や！")
	flag.Parse()
	fmt.Printf("execute by %d parallel.\n", *parallelNumber)
}
