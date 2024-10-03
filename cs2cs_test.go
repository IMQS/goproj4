package goproj4

import (
	"testing"

	"github.com/IMQS/goproj4/globals"
)

const paramstr = `+proj=sterea +lat_0=52.15616055555555 +lon_0=5.38763888888889 +k=0.9999079 +x_0=155000 +y_0=463000 +ellps=bessel +towgs84=565.4171,50.3319,465.5524,-0.398957388243134,0.343987817378283,-1.87740163998045,4.0725 +units=m +no_defs +to +proj=latlon +datum=WGS84 +ellps=WGS84`

func TestCs2Cs(t *testing.T) {
	cs2cs, err := NewCs2Cs(paramstr)
	if err != nil {
		t.Error(err)
	}

	if cs2cs.FromCS.FloatParams["lat_0"] != 52.15616055555555 {
		t.Errorf("Expected %v : Found %v", 52.15616055555555, cs2cs.FromCS.FloatParams["lat_0"])
	}

	/*
		I ran some tests with the transform being done by the original proj4 (C version) and this one
		and compared the result:
		Proj.4 : 52.000000017476815, 5.000000001137788, 143.61679825000465
		goproj4 : 52.00000000963378, 5.000000001518494, 143.61519583221525

		and then calculated the
		distance between these coordinates: 0.0000008725 km (0.8725mm)
		and the height difference : 0.0016024177893996239m (1.6mm)

		Should be good enough?
	*/

	temp := []*globals.XYZ{{
		X: 128410.08539230417,
		Y: 445806.50796043128,
		Z: 100}}
	cs2cs.Transform(temp)

	out := []*globals.XYZ{{
		X: 52.00000000963378,
		Y: 5.000000001518494,
		Z: 143.61519583221525}}

	if temp[0].X != out[0].X {
		t.Errorf("Expected %v : Found %v", out[0].X, temp[0].X)
	}

	if temp[0].Y != out[0].Y {
		t.Errorf("Expected %v : Found %v", out[0].Y, temp[0].Y)
	}

	if temp[0].Z != out[0].Z {
		t.Errorf("Expected %v : Found %v", out[0].Z, temp[0].Z)
	}

}
