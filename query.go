package main

import (
	"strconv"
	"strings"
)

// return slice of color names
func Query(str string) (ret []string) {
	spl := strings.Split(str, ",")
	for _, v := range spl {
		sec := strings.Split(v, ":")
		if len(sec) == 2 {
			// assume it's string
			onev, twov := sec[0], sec[1]
			onei, twoi := -1, -1

			onen, onee := strconv.Atoi(onev)
			twon, twoe := strconv.Atoi(twov)
			if onee == nil && twoe == nil {
				onei, twoi = onen, twon
			} else {
				for i, v := range colors {
					if v == onev {
						onei = i
					}

					if v == twov {
						twoi = i
						break
					}
				}
			}

			// it's not valid :9
			if onei == -1 || twoi == -1 || onei > twoi {
				continue
			}

			ret = append(ret, colors[onei:twoi+1]...)
		} else if len(sec) == 1 {
			num, err := strconv.Atoi(v)
			if err == nil {
				// number from index?
				if num < len(colors) && num >= 0 {
					// is number valid
					ret = append(ret, colors[num])
				}
			} else {
				// nah, color name
				_, ok := material[v]
				if ok {
					// does color name exist
					ret = append(ret, v)
				}
			}
		} else {
			continue
		}
	}

	return ret
}
