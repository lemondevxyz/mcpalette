package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/gookit/color.v1"
)

func print_palette(p Palette) {

	width, _, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("cannot get terminal size")
	}

	stopat := len(colors)
	size := 0
	// color names
	for k, v := range colors {
		clr := p.Get500(v)

		rgb := color.HexToRGB(clr)
		r, g, b := uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])

		x := (6 - len(v))

		space := " "
		if x > 0 {
			space = strings.Repeat(space, x)
		}
		name := v

		v = space + v + space
		if len(name) <= 6 && len(v) > 8 {
			v = v[:8]
		} else if len(name) <= 5 && len(v) < 8 {
			v += strings.Repeat(" ", 8-len(v))
		}

		bg := color.RGB(255, 255, 255, false)
		if ((float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 255) > 0.5 {
			bg = color.RGB(0, 0, 0, false)
		}

		fg := color.RGB(r, g, b, true)
		size += len(v)
		if size > width {
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
		for _, v := range colors[:stopat] {
			if len(p[v][s]) > 0 {
				rgb := color.HexToRGB(p[v][s])
				r, g, b := uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2])

				fg := color.RGB(255, 255, 255, false)
				if ((float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 255) > 0.5 {
					fg = color.RGB(0, 0, 0, false)
				}

				left := " "
				right := " "

				fl := float64((len(v) - 6)) / 2
				if fl > 0 {
					x := int(fl)
					left = strings.Repeat(left, x+1)
					if math.Mod(fl, 1) != 0 {
						x++
					}
					right = strings.Repeat(right, x+1)
				}

				str := left + p[v][s][1:] + right
				if !monochrome {
					str = fg.Sprint(color.RGB(r, g, b, true).Sprint(str))
				}

				fmt.Print(str)
			}
		}
		fmt.Println()
	}

}
