package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var lambda = 0.0135
var bestFit = 1.5608

func main() {
	f, err := os.Create("servinsp23Generate.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	datawriter := bufio.NewWriter(f)

	for i := 0; i < 300; i++ {
		var ri float64
		ri = float64(i) + 1
		//  xi := 0.6084 * math.Exp((lambda * ri))
		xi := bestFit * math.Exp(lambda*ri)
		// xi := (1 / lambda) * (math.Ln2 * ri)

		_, err := datawriter.WriteString(fmt.Sprintf("%f\n", xi))
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}
	datawriter.Flush()
	f.Close()
}
