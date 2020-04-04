package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

var lambda = 0.0125
var bestFit = 0.8026
var dataArray []float64

func main() {
	f, err := os.Create("ws3Generate.dat")
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
		dataArray = append(dataArray, xi)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(dataArray), func(i, j int) { dataArray[i], dataArray[j] = dataArray[j], dataArray[i] })
	for i := 0; i < len(dataArray); i++ {
		_, err := datawriter.WriteString(fmt.Sprintf("%f\n", dataArray[i]))
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}
	datawriter.Flush()
	f.Close()
}
