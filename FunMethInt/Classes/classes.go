package main

import (
	"data"
	"fmt"
	"objectdata"
)

type RegModel struct {
	name             string
	intercept, slope float64
}

func (rm RegModel) Init(name string, intercept, slope float64) {
	rm.name = name
	rm.intercept = intercept
	rm.slope = slope
}

func (rm RegModel) Predict(predictors []float64) []float64 {
	predictees := make([]float64, len(predictors))
	for i, x := range predictors {
		predictees[i] = rm.intercept + rm.slope*x
	}
	return predictees
}

func main() {
	// R := new(RegModel)
	R := RegModel{"RegModel1", 2, 7}
	fmt.Println("Object R of class RegModel:", R.name)
	fmt.Println("Prediction:", R.Predict([]float64{1, 2, 3}))

	fmt.Println("---------------------------------------------")
	fmt.Println("Data package, small X:", data.GetSmallX())
	fmt.Println("Data package, big X:", data.X)
	data.NullData()
	fmt.Println("Data package, small X:",
		data.GetSmallX(),
		"big X:",
		data.GetBigX())
	data.StoreData([]float64{122.110, 423.001, 389.63})
	fmt.Println("Data package, data:", data.RetrieveData())

	fmt.Println("---------------------------------------------")
	d1 := new(objectdata.ObjectData)
	d1.FillXs(2, 150)
	fmt.Println("Objectdata package, small X:", d1.GetSmallX())
	fmt.Println("Objectdata package, big X", d1.GetBigX())
	d1.X = 200
	fmt.Println("Objectdata package, big X", d1.GetBigX())
	d1.StoreData([]float64{122.110, 423.001, 389.63})
	fmt.Println("Objectdata package, data:", d1.RetrieveData())

	fmt.Println("---------------------------------------------")
	d2 := new(objectdata.ObjectData)
	fmt.Println("Objectdata package, 2, object:", &d2)
	fmt.Println("Objectdata package, 2, small X:",
		d2.GetSmallX(),
		"big X:",
		d2.GetBigX())
}
