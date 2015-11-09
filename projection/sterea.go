package projection

import (
	"fmt"
	"goproj4/globals"
	"math"
)

// I am including gauss in here as it is the only place it is used

func NewSterea(params map[string]float64) *Projection {
	proj := &Projection{
		Forward: StereaForward,
		Inverse: StereaInverse,
	}

	// Gauss initialisation
	sphi := math.Sin(params["phi0"])
	cphi := math.Cos(params["phi0"])
	cphi = cphi * cphi
	params["R"] = math.Sqrt(1.0-params["es"]) / (1.0 - params["es"]*sphi*sphi)
	params["C"] = math.Sqrt(1.0 + params["es"]*cphi*cphi/(1.0-params["es"]))
	params["chi"] = math.Asin(sphi / params["C"])
	params["latc0"] = math.Asin(sphi / params["C"])
	params["ratexp"] = 0.5 * params["C"] * params["e"]
	params["K"] = math.Tan(0.5*params["chi"]+globals.PI_OVER_4) / (math.Pow(math.Tan(0.5*params["phi0"]+globals.PI_OVER_4), params["C"]) * srat(params["e"]*sphi, params["ratexp"]))

	// Sterea initialisation
	params["sinc0"] = math.Sin(params["chi"])
	params["cosc0"] = math.Cos(params["chi"])
	params["2R"] = 2.0 * params["R"]

	return proj
}

func StereaForward(params map[string]float64, plh *globals.PLH) *globals.XYZ { // lat, lon float64) (x, y float64) {
	xyz := &globals.XYZ{}
	// Gauss
	// local_phi := lat * DEG_TO_RAD
	// local_lam := lon * DEG_TO_RAD
	
	plh.Phi = 2.0*math.Atan(params["K"]*math.Pow(math.Tan(0.5*plh.Phi+globals.PI_OVER_4), params["C"])*srat(params["e"]*math.Sin(plh.Phi), params["ratexp"])) - globals.PI_OVER_2
	plh.Lam = params["C"] * plh.Lam

	// Sterea
	sinc := math.Sin(plh.Phi)
	cosc := math.Cos(plh.Phi)
	cosl := math.Cos(plh.Lam)
	k := params["k"] * params["2R"] / (1.0 + params["sinc0"]*sinc + params["cosc0"]*cosc*cosl)
	xyz.X = k * cosc * math.Sin(plh.Lam)
	xyz.Y = k * (params["cosc0"]*sinc - params["sinc0"]*cosc*cosl)
	return xyz
}

func StereaInverse(params map[string]float64, xyz *globals.XYZ) *globals.PLH { // x, y float64) (lat, lon float64) {
	splh := &globals.PLH{}
	// sterea
	const MAX_ITER = 20
	const DEL_TOL = 1e-14
	xyz.X = xyz.X / params["k"]
	xyz.Y = xyz.Y / params["k"]
	rho := math.Hypot(xyz.X, xyz.Y)

	if rho != 0 {
		c := 2.0 * math.Atan2(rho, params["2R"])
		sinc := math.Sin(c)
		cosc := math.Cos(c)
		splh.Phi = math.Asin(cosc*params["sinc0"] + xyz.Y*sinc*params["cosc0"]/rho)
		splh.Lam = math.Atan2(xyz.X*sinc, rho*params["cosc0"]*cosc-xyz.Y*params["sinc0"]*sinc)
	} else {
		splh.Phi = params["latc0"]
		splh.Lam = 0.0
	}

	// inv_gaus
	eplh := &globals.PLH{}
	eplh.Lam = splh.Lam / params["C"]
	num := math.Pow(math.Tan(0.5*splh.Phi+globals.PI_OVER_4)/params["K"], 1.0/params["C"])
	var i uint8
	for i = 0; i < MAX_ITER; i++ {
		eplh.Phi = 2.0*math.Atan(num*srat(params["e"]*math.Sin(splh.Phi), -0.5*params["e"])) - globals.PI_OVER_2
		if math.Abs(eplh.Phi-splh.Phi) < DEL_TOL {
			break
		}
		splh.Phi = eplh.Phi
	}
	if i == MAX_ITER-1 {
		fmt.Printf("Convergence failed")
	}
	eplh.Height = xyz.Z
	return eplh
}
