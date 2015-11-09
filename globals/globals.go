package globals

import (
	"math"
)

const (
	PI         = math.Pi
	TWOPI      = (PI * 2.0)
	DEG_TO_RAD = (PI / 180)
	RAD_TO_DEG = (180 / PI)
	SEC_TO_RAD = (PI / 180 / 3600)
	PI_OVER_2  = (PI / 2.0)
	PI_OVER_4  = (PI / 4.0)
	COS_67P5   = 0.38268343236508977 /* cosine of 67.5 degrees */
	AD_C       = 1.0026000           /* Toms region 1 constant */
	GENAU      = 1.E-12
	GENAU2     = (GENAU * GENAU)
)

type XYZ struct {
	X float64
	Y float64
	Z float64
}

type PLH struct {
	Phi    float64
	Lam    float64
	Height float64
}

type DatumType uint8

const (
	DATUMUNKNOWN DatumType = iota
	DATUMWGS84
	DATUM3PARAM
	DATUM7PARAM
)
