package dither

type Filter struct {
	Name   string
	Matrix [][]float32
}

func NewFilter(name string, matrix [][]float32) *Filter {
	return &Filter{
		Name:   name,
		Matrix: matrix,
	}
}

var FloydSteinberg = NewFilter(
	"floydSteinberg",
	[][]float32{
		[]float32{0.0, 0.0, 0.0, 7.0 / 48.0, 5.0 / 48.0},
		[]float32{3.0 / 48.0, 5.0 / 48.0, 7.0 / 48.0, 5.0 / 48.0, 3.0 / 48.0},
		[]float32{1.0 / 48.0, 3.0 / 48.0, 5.0 / 48.0, 3.0 / 48.0, 1.0 / 48.0},
	},
)

var Burkes = NewFilter(
	"burkes",
	[][]float32{
		{0.0, 0.0, 0.0, 8.0 / 32.0, 4.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 8.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
	},
)

var SierraLite = NewFilter(
	"sirra",
	[][]float32{
		{0.0, 0.0, 2.0 / 4.0},
		{1.0 / 4.0, 1.0 / 4.0, 0.0},
		{0.0, 0.0, 0.0},
	},
)

var Sierra2 = NewFilter(
	"sierra2",
	[][]float32{
		{0.0, 0.0, 0.0, 4.0 / 16.0, 3.0 / 16.0},
		{1.0 / 16.0, 2.0 / 16.0, 3.0 / 16.0, 2.0 / 16.0, 1.0 / 16.0},
		{0.0, 0.0, 0.0, 0.0, 0.0},
	},
)

var Sierra3 = NewFilter(
	"sierra3",
	[][]float32{
		{0.0, 0.0, 0.0, 5.0 / 32.0, 3.0 / 32.0},
		{2.0 / 32.0, 4.0 / 32.0, 5.0 / 32.0, 4.0 / 32.0, 2.0 / 32.0},
		{0.0, 2.0 / 32.0, 3.0 / 32.0, 2.0 / 32.0, 0.0},
	},
)

var Stucki = NewFilter(
	"stucki",
	[][]float32{
		{0.0, 0.0, 0.0, 8.0 / 42.0, 4.0 / 42.0},
		{2.0 / 42.0, 4.0 / 42.0, 8.0 / 42.0, 4.0 / 42.0, 2.0 / 42.0},
		{1.0 / 42.0, 2.0 / 42.0, 4.0 / 42.0, 2.0 / 42.0, 1.0 / 42.0},
	},
)

var Atkinson = NewFilter(
	"atkinson",
	[][]float32{
		{0.0, 0.0, 1.0 / 8.0, 1.0 / 8.0},
		{1.0 / 8.0, 1.0 / 8.0, 1.0 / 8.0, 0.0},
		{0.0, 1.0 / 8.0, 0.0, 0.0},
	},
)