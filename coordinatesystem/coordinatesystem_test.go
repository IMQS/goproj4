package coordinatesystem

import (
	"goproj4/globals"
	"testing"
)

// 128410.08539230417 445806.50796043128 100
const paramstr = `+proj=sterea +lat_0=52.15616055555555 +lon_0=5.38763888888889 +k=0.9999079 +x_0=155000 +y_0=463000 +ellps=bessel +towgs84=565.4171,50.3319,465.5524,-0.398957388243134,0.343987817378283,-1.87740163998045,4.0725 +units=m +no_defs`

func TestNewCoordinateSystem(t *testing.T) {
	cs, err := NewCoordinateSystem(paramstr)
	if err != nil {
		t.Error(err)
	}

	if cs.FloatParams["lat_0"] != 52.15616055555555 {
		t.Errorf("Expected %v : Found %v", 52.15616055555555, cs.FloatParams["lat_0"])
	}

	plh := cs.ToLatLon(&globals.XYZ{128410.08539230417, 445806.50796043128, 100})
	// Convert back to actual
	plh.Phi = plh.Phi * globals.RAD_TO_DEG
	plh.Lam = plh.Lam * globals.RAD_TO_DEG
	if plh.Phi != 52.000000017476815 {
		t.Errorf("Expected %v : Found %v diff %v", 52.000000017476815, plh.Phi, plh.Phi-52.000000017476815)
	}

	if plh.Lam != 5.000000001137788 {
		t.Errorf("Expected %v : Found %v diff %v", 5.000000001137788, plh.Lam, plh.Lam-5.000000001137788)
	}

	if plh.Height != 143.61679825000465 {
		t.Errorf("Expected %v : Found %v diff %v", 143.61679825000465, plh.Height, plh.Height-143.61679825000465)
	}
}
