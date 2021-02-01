package main

import (
	"flag"
)

var monochrome = false

func main() {
	nocolor := flag.Bool("no-color", false, "display the hex codes but without color")
	if *nocolor == true {
		monochrome = true
	}

	flag.Parse()
	print_palette(material)
}
