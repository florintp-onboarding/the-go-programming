package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
	fmt.Printf("File:\t%s\n",filename)
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
	    if counts[line] >1 && len(line) > 1  {
		    fmt.Printf("String Duplicate %s in ==>:\t%s\n",line,filename)
	    }
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
