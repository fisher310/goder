package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
}

func main() {
	//flag.VisitAll(func(f *flag.Flag) {
	//	format := "\t-%s: %s (Default: '%s')\n"
	//	fmt.Printf(format, f.Name, f.Usage, f.DefValue)
	//})
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hala: %s !\n", *name)
	} else {
		fmt.Printf("Hello: %s !\n", *name)
	}
}
