package p99

type IP99 interface {
	QuantileOf(pxx int) float64
	Add(val float64)
}

type P99 struct {
	//List [] //list[0...99]
}

func (p P99) Add(val float64) {
	panic("implement me")
}

func NewP99() *P99 {
	return &P99{}
}

func (p P99) QuantileOf(pxx int) float64 {
	panic("implement me")
}
