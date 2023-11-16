package filter

type Filter struct {
	ColorBalance                            [3]float32
	Contrast, Gamma, Saturation, Brightness float32
}

func New() *Filter {
	return &Filter{
		ColorBalance: ColorBalance(0, 0, 0),
		Contrast:     0,
		Gamma:        1,
		Saturation:   0,
		Brightness:   0,
	}
}
