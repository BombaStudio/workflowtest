package main

import (
	"flag"
	"fmt"
)

func main() {
	cmd := flag.String("cmd", "", "")
	flag.Parse()
	fmt.Printf("my cmd: \"%v\"\n", string(*cmd))
}
