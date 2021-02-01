package main

import (
	"flag"
)

var monochrome = false
var colorfilter = ""

func main() {
	{
		const fu = "display the hex codes but without color"
		flag.BoolVar(&monochrome, "no-color", false, fu)
		flag.BoolVar(&monochrome, "nc", false, fu+" (shorthand)")
	}
	/*
		{
			const fu = "display rgb codes instead of hex codes"
			flag.StringVar(&colorfilter, "color", "")
		}
	*/
	flag.Parse()

	//print_palette(FormatRGB, material)
	print_palette(FormatHSL, material)
}
