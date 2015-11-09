package goproj4

import (
	"fmt"
	"github.com/IMQS/goproj4/coordinatesystem"
	"github.com/IMQS/goproj4/globals"
	"strings"
)

type Cs2Cs struct {
	FromCS *coordinatesystem.CoordinateSystem
	ToCS   *coordinatesystem.CoordinateSystem
}

func NewCs2Cs(params string) (*Cs2Cs, error) {
	params2 := strings.Split(params, " +to ")
	if len(params2) != 2 {
		return nil, fmt.Errorf("Not enough parameters to work with")
	}
	params2[0] = strings.TrimSpace(params2[0])
	params2[1] = strings.TrimSpace(params2[1])

	if len(params2[0]) == 0 || len(params2[1]) == 0 {
		return nil, fmt.Errorf("Not enough parameters to work with")
	}
	var err error
	cs2cs := &Cs2Cs{}
	cs2cs.FromCS, err = coordinatesystem.NewCoordinateSystem(params2[0])
	if err != nil {
		return nil, err
	}

	cs2cs.ToCS, err = coordinatesystem.NewCoordinateSystem(params2[1])
	if err != nil {
		return nil, err
	}

	return cs2cs, nil
}

func (cs2cs *Cs2Cs) Transform(xyz []*globals.XYZ) { //  (plh []*globals.PLH) {
	// TODO axis
	// if axis, ok := cs2cs.FromCS.StringParams["axis"]; ok {
	//	if axis != "enu" {
	//		//Adjust axis
	//	}
	// }

	plh := make([]*globals.PLH, len(xyz))

	if cs2cs.FromCS.FloatParams["vto_meter"] != 1.0 {
		for idx, _ := range xyz {
			xyz[idx].Z = xyz[idx].Z * cs2cs.FromCS.FloatParams["vto_meter"]
		}
	}

	if cs2cs.FromCS.IsGeocent() {
	} else if !cs2cs.FromCS.IsLatlon() {
		for idx, _ := range xyz {
			plh[idx] = cs2cs.FromCS.Inverse(xyz[idx])
		}
	}

	if cs2cs.FromCS.DatumType() == globals.DATUM3PARAM ||
		cs2cs.FromCS.DatumType() == globals.DATUM7PARAM ||
		cs2cs.ToCS.DatumType() == globals.DATUM3PARAM ||
		cs2cs.ToCS.DatumType() == globals.DATUM7PARAM {

		for idx, _ := range plh {
			xyz[idx] = cs2cs.FromCS.GeodeticToGeocentric(plh[idx])

			if cs2cs.FromCS.DatumType() == globals.DATUM3PARAM ||
				cs2cs.FromCS.DatumType() == globals.DATUM7PARAM {
				xyz[idx] = cs2cs.FromCS.GeocentricToWGS84(xyz[idx])
			}

			if cs2cs.ToCS.DatumType() == globals.DATUM3PARAM ||
				cs2cs.ToCS.DatumType() == globals.DATUM7PARAM {
				xyz[idx] = cs2cs.ToCS.GeocentricFromWGS84(xyz[idx])
			}

			plh[idx] = cs2cs.ToCS.GeocentricToGeodetic(xyz[idx])
		}
	}

	if cs2cs.ToCS.IsGeocent() {
	} else if !cs2cs.ToCS.IsLatlon() {
		for idx, _ := range plh {
			xyz[idx] = cs2cs.ToCS.Forward(plh[idx])
		}
	}

	if cs2cs.ToCS.IsLatlon() {
		for idx, _ := range xyz {
			xyz[idx].X = plh[idx].Phi * globals.RAD_TO_DEG
			xyz[idx].Y = plh[idx].Lam * globals.RAD_TO_DEG
			xyz[idx].Z = plh[idx].Height
		}
	}
}
