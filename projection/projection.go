package projection

import (
	"math"
	"goproj4/globals"
)

// lam = lon
// phi = lat

// type Transform interface {
// 	Forward(map[string]float64, *PLH) *XYZ
// 	Inverse(map[string]float64, *XYZ) *PLH
// }

type Projection struct {
	Forward func(map[string]float64, *globals.PLH) *globals.XYZ
	Inverse func(map[string]float64, *globals.XYZ) *globals.PLH
}

var ProjectionList = map[string]func(map[string]float64) *Projection{
	"latlon": NewLatLon,
	"sterea": NewSterea,
}

func srat(esinp, exp float64) float64 {
	return math.Pow((1.0-esinp)/(1.0+esinp), exp)
}
