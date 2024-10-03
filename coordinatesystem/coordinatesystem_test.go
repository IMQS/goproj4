package coordinatesystem

import (
	"testing"

	"github.com/IMQS/goproj4/globals"
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

	from := &globals.XYZ{
		X: 128410.08539230417,
		Y: 445806.50796043128,
		Z: 100,
	}
	plh := cs.Inverse(from)
	// Convert back to actual
	plh.Phi = plh.Phi * globals.RAD_TO_DEG
	plh.Lam = plh.Lam * globals.RAD_TO_DEG

	expected := &globals.PLH{
		Phi:    52.00096942804018,
		Lam:    5.00037918470534,
		Height: 100,
	}

	if plh.Phi != expected.Phi {
		t.Errorf("Expected %v : Found %v diff %v", expected.Phi, plh.Phi, plh.Phi-expected.Phi)
	}

	if plh.Lam != expected.Lam {
		t.Errorf("Expected %v : Found %v diff %v", expected.Lam, plh.Lam, plh.Lam-expected.Lam)
	}

	if plh.Height != expected.Height {
		t.Errorf("Expected %v : Found %v diff %v", expected.Height, plh.Height, plh.Height-expected.Height)
	}
}
