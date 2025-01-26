package operation

// Without generic
type CalculatorInt struct {
	X      int64
	Y      int64
	Result int64
}

func (value *CalculatorInt) Sum() {
	value.Result = value.X + value.Y
}

func New(valueX, valueY int64) CalculatorInt {
	return CalculatorInt{X: valueX, Y: valueY}
}

type CalculatorFloat struct {
	X      float64
	Y      float64
	Result float64
}

func (value *CalculatorFloat) Sum() {
	value.Result = value.X + value.Y
}

func NewFloat(x, y float64) CalculatorFloat {
	return CalculatorFloat{X: x, Y: y}
}

// with generic can use one struct.
type Number interface {
	int64 | float64
}

type CalculatorIntsOrFloats[K Number] struct {
	X, Y   K
	Result K
}

func NewCalculatorIntsOrFloats[T Number](x, y T) CalculatorIntsOrFloats[T] {
	newStruct := CalculatorIntsOrFloats[T]{X: x, Y: y}

	return newStruct
}

func (value *CalculatorIntsOrFloats[K]) Sum() {
	value.Result = value.X + value.Y
}
