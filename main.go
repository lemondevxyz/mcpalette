package main

import (
	"flag"
)

var monochrome = false
var colorfilter = ""

func main() {

	clr := ""

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
	{
		const fu = `print specific colors: "red, blue" or "1, 3, 4" or "red:lightgreen" or "5:7" - invalid data will be skipped without throwing an error`
		flag.StringVar(&clr, "color", "", fu)
		flag.StringVar(&clr, "c", "", fu+" (shorthand)")
	}
	flag.Parse()

	if rgb == true {
		f = FormatRGB
	} else if hsl == true {
		f = FormatHSL
	}

	cs := colors
	if len(clr) > 0 {
		cs = Query(clr)
	}

	print_palette(f, cs)
}
