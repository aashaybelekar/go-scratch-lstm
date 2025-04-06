package cal

import (
	"fmt"
	"math"
)

type State struct {
	St float64
	Lt float64
}

func Sigmoid(v float64) float64 {
	x := 1 / (1 + math.Pow(math.E, v))
	fmt.Println("power is ", x)
	return x
}

func forget(i float64, s State) State {
	return s //temp
}
