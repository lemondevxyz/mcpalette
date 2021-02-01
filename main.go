package main

import (
	"flag"
)

var monochrome = false
var colorfilter = ""

func main() {

	f := FormatHEX
	rgb := false
	hsl := false
	{
		const fu = "display the hex codes but without color"
		flag.BoolVar(&monochrome, "no-color", false, fu)
		flag.BoolVar(&monochrome, "n", false, fu+" (shorthand)")
	}
	{
		const fu = "display rgb instead of hex codes"
		flag.BoolVar(&rgb, "rgb", false, fu)
		flag.BoolVar(&rgb, "r", false, fu+" (shorthand)")
	}
	{
		const fu = "display hsl instead of hex codes"
		flag.BoolVar(&hsl, "hsl", false, fu)
		flag.BoolVar(&hsl, "h", false, fu+" (shorthand)")
	}
	flag.Parse()

	if rgb == true {
		f = FormatRGB
	} else if hsl == true {
		f = FormatHSL
	}

	//print_palette(FormatRGB, material)
	print_palette(f, material)
}
