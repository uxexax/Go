package data

var x int = 1   // not accessible from outside of the package
var X int = 100 // accessible from outside of the package
var internal_data []float64

func GetSmallX() int {
	return x
}

func GetBigX() int {
	return X
}

func NullData() {
	null_data()
}

func StoreData(data []float64) {
	internal_data = data
}

func RetrieveData() []float64 {
	return internal_data
}

func null_data() { // not accessible from outside of the package
	x = 0
	X = 0
}
