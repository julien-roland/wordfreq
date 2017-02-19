package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"text/scanner"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Usage = usage
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: %s\n", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "wordfreq: missing input file\n")
		os.Exit(1)
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)

	wordCounts := make(map[string]int)
	for s.Scan() != scanner.EOF {
		wordCounts[s.TokenText()]++
	}

	for word, count := range wordCounts {
		fmt.Printf("%s : %d\n", word, count)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: wordfreq [options] file\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(2)
}
