package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kor44/prefix"
)

var Usage = func() {
	fmt.Printf("Usage of %s: <begin> <end>\n", os.Args[0])
	fmt.Println("  Need specifiy <begin> and <end> range to generate list of prefixes")
	fmt.Printf("  Example (will generate 201,202,203,2040): %s -s , 2010 2040\n", os.Args[0])
}

func main() {
	help := flag.Bool("h", false, "show help")
	separator := flag.String("s", "\n", "seprator")
	flag.Parse()

	if *help {
		Usage()
		os.Exit(0)
	}

	if len(flag.Args()) < 2 {
		Usage()
		fmt.Fprintln(os.Stderr, "Not enough input arguments")
		os.Exit(1)
	}

	prefixList, err := prefix.FromRange(flag.Args()[0], flag.Args()[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(strings.Join(prefixList, *separator))
}
