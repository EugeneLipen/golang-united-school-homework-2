package square

import (
	"math"
)

type mytype int

const (
	SidesCircle   mytype = 0
	SidesTriangle mytype = 3
	SidesSquare   mytype = 4
)

func CalcSquare(sideLen float64, sideNum mytype) float64 {
	var res float64
	switch sideNum {
	case SidesSquare:
		res = math.Pow(sideLen, 2)
	case SidesTriangle:
		res = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case SidesCircle:
		res = math.Pi * math.Pow(sideLen, 2)
	default:
		res = 0
	}

	return res
}
