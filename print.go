package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/gookit/color.v1"
)

type Format uint8

const (
	FormatHEX Format = iota
	FormatRGB
	FormatHSL
)

func print_palette(f Format, cs []string) {

	colorlen := 0
	if f == FormatHEX {
		colorlen = 6
	} else if f == FormatRGB {
		colorlen = 18 - 5
	} else if f == FormatHSL {
		colorlen = 20 - 5
	}

	width, _, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("cannot get terminal size")
	}

	stopat := len(cs)
	size := 0
	// color names
	for k, v := range cs {
		hex := material.Get500(v)

		clr, err := colorful.Hex(hex)
		if err != nil {
			continue
		}

		r, g, b := clr.RGB255()

		fl := float64(colorlen-len(v)) / 2
		x := 0

		left := " "
		right := " "
		if fl > 0 {
			x = int(fl)
			left = strings.Repeat(left, x+1)
			if math.Mod(fl, 1) != 0 {
				x++
			}

			right = strings.Repeat(right, x+1)
		}
		//name := v

		v = left + v + right

		bg := color.RGB(255, 255, 255, false)
		if ((float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 255) > 0.5 {
			bg = color.RGB(0, 0, 0, false)
		}

		y := len(v)
		if len(v) < colorlen {
			y = colorlen
		}
		fg := color.RGB(r, g, b, true)

		size += y

		if size > width {
			size = size - y
			stopat = k
			break
		}

		if !monochrome {
			v = bg.Sprint(fg.Sprint(v))
		}
		fmt.Print(v)
	}
	fmt.Println()

	// actual colors
	for _, s := range order {
		for _, v := range cs[:stopat] {
			if len(material[v][s]) > 0 {
				clr, err := colorful.Hex(material[v][s])
				if err != nil {
					continue
				}

				r, g, b := clr.RGB255()

				fg := color.RGB(255, 255, 255, false)
				if ((float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 255) > 0.5 {
					fg = color.RGB(0, 0, 0, false)
				}

				left := " "
				right := " "

				fl := float64((len(v) - colorlen)) / 2
				if fl > 0 {
					x := int(fl)
					left = strings.Repeat(left, x+1)
					if math.Mod(fl, 1) != 0 {
						x++
					}
					right = strings.Repeat(right, x+1)
				}

				str := ""
				if f == FormatHEX {
					str = material[v][s][1:]
				} else if f == FormatRGB {
					str = fmt.Sprintf("%03d, %03d, %03d", r, g, b)
				} else if f == FormatHSL {
					h, s, l := clr.Hsl()

					//str = "not implemented"
					str = fmt.Sprintf("%03.0f, %03.0f%%, %03.0f%%", h, s*100, l*100)
				}

				str = left + str + right
				if !monochrome {
					str = fg.Sprint(color.RGB(r, g, b, true).Sprint(str))
				}

				fmt.Print(str)
			}
		}
		fmt.Println()
	}

}
