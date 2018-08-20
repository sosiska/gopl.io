// Cf converts its numeric argument to Meters and Foots.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sosiska/gopl.io/ch2/distconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		d, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "distconverter: %v\n", err)
			os.Exit(1)
		}
		f := distconv.Foots(d)
		m := distconv.Meters(d)
		fmt.Printf("%s = %s, %s = %s\n",
			f, distconv.FToM(f), m, distconv.MToF(m))
	}
}
