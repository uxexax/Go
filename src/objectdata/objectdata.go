package objectdata

type ObjectData struct {
	x             int
	X             int
	internal_data []float64
}

func (o *ObjectData) GetSmallX() int {
	return o.x
}

func (o *ObjectData) GetBigX() int {
	return o.X
}

func (o *ObjectData) FillXs(x, X int) {
	o.fill_xs(x, X)
	o.x = x
	o.X = X
}

func (o *ObjectData) StoreData(data []float64) {
	o.internal_data = data
}

func (o *ObjectData) RetrieveData() []float64 {
	return o.internal_data
}

func (o *ObjectData) fill_xs(x, X int) {
	o.x = x
	o.X = X
}
