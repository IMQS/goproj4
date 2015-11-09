package projection

import (
	"goproj4/globals"
)

func NewLatLon(params map[string]float64) *Projection {
	proj := &Projection{
		Forward: LatLonForward,
		Inverse: LatLonInverse,
	}
	return proj
}

func LatLonForward(params map[string]float64, plh *globals.PLH) *globals.XYZ {
	xyz := &globals.XYZ{}
	xyz.X = plh.Lam / params["a"]
	xyz.Y = plh.Phi / params["a"]
	return xyz
}

func LatLonInverse(params map[string]float64, xyz *globals.XYZ) *globals.PLH {
	plh := &globals.PLH{}
	plh.Lam = xyz.X * params["a"]
	plh.Phi = xyz.Y * params["a"]
	return plh
}
