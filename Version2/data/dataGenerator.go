package data

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)


func Generate(file string, lambda float64, bestFit float64, x int) {
	var dataArray []float64
	f, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	datawriter := bufio.NewWriter(f)

	for i := 0; i < x; i++ {
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
