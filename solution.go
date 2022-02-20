package square

import (
	"math"
)

type mytype int

func CalcSquare(sideLen float64, sideNum mytype) float64 {
	var res float64
	switch sideNum {
	case 4:
		res = math.Pow(sideLen, 2)
	case 3:
		res = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	case 0:
		res = math.Pi * math.Pow(sideLen, 2)
	default:
		res = 0
	}

	return res
}
