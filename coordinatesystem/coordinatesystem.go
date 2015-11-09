package coordinatesystem

import (
	"fmt"
	"github.com/IMQS/goproj4/globals"
	"github.com/IMQS/goproj4/projection"
	"github.com/IMQS/goproj4/ref"
	"math"
	"strconv"
	"strings"
)

type CoordinateSystem struct {
	FloatParams  map[string]float64
	StringParams map[string]string
	BoolParams   map[string]bool // This is actually more of an existensial list
	ToWGS84      []float64
	Projection   *projection.Projection
}

func NewCoordinateSystem(paramstr string) (*CoordinateSystem, error) {
	cs := &CoordinateSystem{
		FloatParams:  make(map[string]float64, 0),
		StringParams: make(map[string]string, 0),
		BoolParams:   make(map[string]bool, 0),
	}

	params := parseParams(paramstr)
	for key, val := range params {
		switch val.(type) {
		case float64:
			cs.FloatParams[key] = val.(float64)
		case string:
			cs.StringParams[key] = val.(string)
		case bool:
			cs.BoolParams[key] = val.(bool)
		case []float64:
			cs.ToWGS84 = val.([]float64)
		}
	}
	err := cs.init()
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (cs *CoordinateSystem) IsGeocent() bool {
	if geoc, ok := cs.StringParams["proj"]; ok {
		return (geoc == "geocent")
	}
	return false
}

func (cs *CoordinateSystem) IsLatlon() bool {
	if latlon, ok := cs.StringParams["proj"]; ok {
		return (latlon == "latlon")
	}
	return false
}

func (cs *CoordinateSystem) DatumType() globals.DatumType {
	if cs.StringParams["datum"] == "WGS84" {
		return globals.DATUMWGS84
	}
	if len(cs.ToWGS84) == 3 {
		return globals.DATUM3PARAM
	}
	if len(cs.ToWGS84) == 7 {
		return globals.DATUM7PARAM
	}
	return globals.DATUMUNKNOWN
}

func (cs *CoordinateSystem) Forward(plh *globals.PLH) *globals.XYZ {
	plh.Lam = plh.Lam - cs.FloatParams["lam0"]
	if !cs.BoolParams["over"] {
		plh.Lam = cs.adjlon(plh.Lam)
	}
	xyz := cs.Projection.Forward(cs.FloatParams, plh)

	xyz.X = cs.FloatParams["from_meter"] * (cs.FloatParams["a"]*xyz.X + cs.FloatParams["x_0"])
	xyz.Y = cs.FloatParams["from_meter"] * (cs.FloatParams["a"]*xyz.Y + cs.FloatParams["x_0"])
	return xyz
}

func (cs *CoordinateSystem) Inverse(xyz *globals.XYZ) *globals.PLH {
	xyz.X = (xyz.X*cs.FloatParams["to_meter"] - cs.FloatParams["x_0"]) * cs.FloatParams["ra"]
	xyz.Y = (xyz.Y*cs.FloatParams["to_meter"] - cs.FloatParams["y_0"]) * cs.FloatParams["ra"]

	plh := cs.Projection.Inverse(cs.FloatParams, xyz)

	plh.Lam = plh.Lam + cs.FloatParams["lam0"]

	if !cs.BoolParams["over"] {
		plh.Lam = cs.adjlon(plh.Lam)
	}

	return plh
}

func (cs *CoordinateSystem) init() error {

	// have to have a proj to do anything
	if _, ok := cs.StringParams["proj"]; !ok {
		return fmt.Errorf("Must specify projection")
	}

	var err error
	err = cs.initDatum()
	if err != nil {
		return err
	}

	err = cs.initEllipse()
	if err != nil {
		return err
	}

	// Unashamedly copied from proj.4
	cs.FloatParams["a_orig"] = cs.FloatParams["a"]
	cs.FloatParams["es_orig"] = cs.FloatParams["es"]
	cs.FloatParams["e"] = math.Sqrt(cs.FloatParams["es"])
	cs.FloatParams["ra"] = 1.0 / cs.FloatParams["a"]
	cs.FloatParams["one_es"] = 1.0 - cs.FloatParams["es"]
	if cs.FloatParams["one_es"] == 0.0 {
		return fmt.Errorf("one_es cannot be zero")
	}
	cs.FloatParams["rone_es"] = 1.0 / cs.FloatParams["one_es"]

	cs.FloatParams["phi0"] = cs.FloatParams["lat_0"] * globals.DEG_TO_RAD
	cs.FloatParams["lam0"] = cs.FloatParams["lon_0"] * globals.DEG_TO_RAD

	if _, ok := cs.BoolParams["over"]; !ok {
		cs.BoolParams["over"] = false
	}

	// Units
	unitparams := [][]string{
		[]string{"unit", "to_meter", "from_meter"},
		[]string{"vunit", "vto_meter", "vfrom_meter"},
	}

	for _, unitparam := range unitparams {
		if unitstr, ok := cs.StringParams[unitparam[0]]; ok {
			if unit, ok := ref.UnitList[unitstr]; ok {
				cs.FloatParams[unitparam[1]] = unit.Multiplier
				cs.FloatParams[unitparam[2]] = 1.0 / unit.Multiplier
			}
		} else if unitval, ok := cs.FloatParams[unitparam[1]]; ok {
			cs.FloatParams[unitparam[2]] = 1.0 / unitval
		} else {
			// Set defaults
			cs.FloatParams[unitparam[1]] = 1.0
			cs.FloatParams[unitparam[2]] = 1.0

		}
	}

	if projstr, ok := cs.StringParams["proj"]; ok {
		if proj, ok := projection.ProjectionList[projstr]; ok {
			cs.Projection = proj(cs.FloatParams)
		}
	}

	return nil
}

func (cs *CoordinateSystem) initDatum() error {
	// check if we have datum param
	if datumstr, ok := cs.StringParams["datum"]; ok {
		if datum, ok := ref.DatumList[datumstr]; ok {
			params := parseParams(datum.Definition)
			if towgs, ok := params["towgs"]; ok {
				cs.ToWGS84 = towgs.([]float64)
			}
		}
	}
	if len(cs.ToWGS84) == 7 {
		cs.ToWGS84[3] = cs.ToWGS84[3] * globals.SEC_TO_RAD
		cs.ToWGS84[4] = cs.ToWGS84[4] * globals.SEC_TO_RAD
		cs.ToWGS84[5] = cs.ToWGS84[5] * globals.SEC_TO_RAD
		cs.ToWGS84[6] = (cs.ToWGS84[6] / 1000000.0) + 1
	}
	return nil
}

func (cs *CoordinateSystem) initEllipse() error {
	if R, ok := cs.FloatParams["R"]; ok {
		cs.FloatParams["a"] = R
		cs.FloatParams["es"] = 0.0
	}
	if ellipsestr, ok := cs.StringParams["ellps"]; ok {
		if ellipse, ok := ref.EllipseList[ellipsestr]; ok {
			params := parseParams(ellipse.Params)
			for key, val := range params {
				cs.FloatParams[key] = val.(float64)
			}

			// For now we are only doing a and rf
			// a already set
			if es, ok := cs.FloatParams["rf"]; ok {
				es = 1.0 / es
				es = es * (2.0 - es)
				cs.FloatParams["es"] = es
			}

		}
	}

	// Sanity checks
	if es, ok := cs.FloatParams["es"]; ok {
		if es < 0.0 {
			return fmt.Errorf("es cannot be smaller than zero")
		}
	} else {
		return fmt.Errorf("es must exist")
	}

	if a, ok := cs.FloatParams["a"]; ok {
		if a <= 0.0 {
			return fmt.Errorf("a cannot be smaller or equal to zero")
		}
	} else {
		return fmt.Errorf("a must exist")
	}

	return nil
}

func (cs *CoordinateSystem) SetGeocentricParameters() {
	var b float64
	a := cs.FloatParams["a_orig"]
	es := cs.FloatParams["es_orig"]
	if es == 0.0 {
		b = a
	} else {
		b = a * math.Sqrt(1-es)
	}

	a2 := a * a
	b2 := b * b
	cs.FloatParams["e2"] = (a2 - b2) / a2
	cs.FloatParams["b"] = b
}

func (cs *CoordinateSystem) GeodeticToGeocentric(plh *globals.PLH) *globals.XYZ {
	xyz := &globals.XYZ{}
	if _, ok := cs.FloatParams["e2"]; !ok {
		cs.SetGeocentricParameters()
	}
	e2 := cs.FloatParams["e2"]
	a := cs.FloatParams["a_orig"]
	if plh.Phi < -globals.PI_OVER_2 && plh.Phi > -1.001*globals.PI_OVER_2 {
		plh.Phi = -globals.PI_OVER_2
	} else if plh.Phi > globals.PI_OVER_2 && plh.Phi < 1.001*globals.PI_OVER_2 {
		plh.Phi = globals.PI_OVER_2
	}

	if plh.Lam > globals.PI {
		plh.Lam = plh.Lam - 2*globals.PI
	}
	sin_lat := math.Sin(plh.Phi)
	cos_lat := math.Cos(plh.Phi)
	sin2_lat := sin_lat * sin_lat
	rn := a / (math.Sqrt(1.0 - e2*sin2_lat))
	xyz.X = (rn + plh.Height) * cos_lat * math.Cos(plh.Lam)
	xyz.Y = (rn + plh.Height) * cos_lat * math.Sin(plh.Lam)
	xyz.Z = ((rn * (1 - e2)) + plh.Height) * sin_lat
	return xyz
}

func (cs *CoordinateSystem) GeocentricToGeodetic(xyz *globals.XYZ) *globals.PLH {
	plh := &globals.PLH{}
	const MAXITER = 30
	if _, ok := cs.FloatParams["e2"]; !ok {
		cs.SetGeocentricParameters()
	}
	e2 := cs.FloatParams["e2"]
	a := cs.FloatParams["a_orig"]
	b := cs.FloatParams["b"]

	p := math.Sqrt(xyz.X*xyz.X + xyz.Y*xyz.Y)
	rr := math.Sqrt(xyz.X*xyz.X + xyz.Y*xyz.Y + xyz.Z*xyz.Z)

	if (p / a) < globals.GENAU {
		plh.Lam = 0.0
		if (rr / a) < globals.GENAU {
			plh.Phi = globals.PI_OVER_2
			plh.Height = -b
			return plh
		}
	} else {
		plh.Lam = math.Atan2(xyz.Y, xyz.X)
	}

	ct := xyz.Z / rr
	st := p / rr
	rx := 1.0 / math.Sqrt(1.0-e2*(2.0-e2)*st*st)
	cphi0 := st * (1.0 - e2) * rx
	sphi0 := ct * rx
	cphi := 0.0
	sphi := 0.0

	iter := 0
	for {
		iter = iter + 1
		rn := a / math.Sqrt(1.0-e2*sphi0*sphi0)
		plh.Height = p*cphi0 + xyz.Z*sphi0 - rn*(1.0-e2*sphi0*sphi0)
		rk := e2 * rn / (rn + plh.Height)
		rx = 1.0 / math.Sqrt(1.0-rk*(2.0-rk)*st*st)
		cphi = st * (1.0 - rk) * rx
		sphi = ct * rx
		sdphi := sphi*cphi0 - cphi*sphi0
		cphi0 = cphi
		sphi0 = sphi
		if sdphi*sdphi < globals.GENAU2 || iter >= MAXITER {
			break
		}
	}
	plh.Phi = math.Atan(sphi / math.Abs(cphi))

	return plh
}

func (cs *CoordinateSystem) GeocentricToWGS84(in *globals.XYZ) *globals.XYZ {
	out := &globals.XYZ{}
	if len(cs.ToWGS84) == 3 {
		out.X = in.X + cs.ToWGS84[0]
		out.Y = in.Y + cs.ToWGS84[1]
		out.Z = in.Z + cs.ToWGS84[2]

	}

	if len(cs.ToWGS84) == 7 {
		out.X = cs.ToWGS84[6]*(in.X-cs.ToWGS84[5]*in.Y+cs.ToWGS84[4]*in.Z) + cs.ToWGS84[0]
		out.Y = cs.ToWGS84[6]*(cs.ToWGS84[5]*in.X+in.Y-cs.ToWGS84[3]*in.Z) + cs.ToWGS84[1]
		out.Z = cs.ToWGS84[6]*(-cs.ToWGS84[4]*in.X+cs.ToWGS84[3]*in.Y+in.Z) + cs.ToWGS84[2]
	}
	return out
}

func (cs *CoordinateSystem) GeocentricFromWGS84(in *globals.XYZ) *globals.XYZ {
	out := &globals.XYZ{}
	if len(cs.ToWGS84) == 3 {
		out.X = in.X - cs.ToWGS84[0]
		out.Y = in.Y - cs.ToWGS84[1]
		out.Z = in.Z - cs.ToWGS84[2]

	}
	if len(cs.ToWGS84) == 7 {
		x := (in.X - cs.ToWGS84[0]) / cs.ToWGS84[6]
		y := (in.Y - cs.ToWGS84[1]) / cs.ToWGS84[6]
		z := (in.Z - cs.ToWGS84[2]) / cs.ToWGS84[6]

		out.X = x + cs.ToWGS84[5]*y - cs.ToWGS84[4]*z
		out.Y = -cs.ToWGS84[5]*x + y + cs.ToWGS84[3]*z
		out.Z = cs.ToWGS84[4]*x - cs.ToWGS84[3]*y + z
	}
	return out
}

func (cs *CoordinateSystem) adjlon(lon float64) float64 {
	if math.Abs(lon) <= globals.PI {
		return lon
	}
	lon = lon + globals.PI
	lon = lon - globals.TWOPI*math.Floor(lon/globals.TWOPI)
	lon = lon - globals.PI
	return lon
}

func parseParams(paramstr string) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	paramstr = strings.Replace(paramstr, "+", "", -1)
	paramstr = strings.Replace(paramstr, "  ", " ", -1) // Keeps from doing trim later
	params := strings.Split(paramstr, " ")
	for _, param := range params {
		if strings.Contains(param, "=") {
			keyval := strings.Split(param, "=")
			// towgs84 special case
			if keyval[0] == "towgs84" {
				vals := strings.Split(keyval[1], ",")
				towgs84 := make([]float64, len(vals))
				for indx, val := range vals {
					towgs84[indx], _ = strconv.ParseFloat(val, 64)
				}
				result[keyval[0]] = towgs84
			} else {

				// check if we can convert to float
				fval, err := strconv.ParseFloat(keyval[1], 64)
				if err != nil {
					result[keyval[0]] = keyval[1]
				} else {
					result[keyval[0]] = fval
				}
			}
		} else {
			result[param] = true
		}
	}
	return result

}
