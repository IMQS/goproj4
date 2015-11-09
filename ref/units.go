package ref

type Unit struct {
	Multiplier  float64
	Description string
}

var UnitList = map[string]*Unit{
	"km":     &Unit{1000., "Kilometer"},
	"m":      &Unit{1.0, "Meter"},
	"dm":     &Unit{1 / 10, "Decimeter"},
	"cm":     &Unit{1 / 100, "Centimeter"},
	"mm":     &Unit{1 / 1000, "Millimeter"},
	"kmi":    &Unit{1852.0, "International Nautical Mile"},
	"in":     &Unit{0.0254, "International Inch"},
	"ft":     &Unit{0.3048, "International Foot"},
	"yd":     &Unit{0.9144, "International Yard"},
	"mi":     &Unit{1609.344, "International Statute Mile"},
	"fath":   &Unit{1.8288, "International Fathom"},
	"ch":     &Unit{20.1168, "International Chain"},
	"link":   &Unit{0.201168, "International Link"},
	"us-in":  &Unit{1. / 39.37, "U.S. Surveyor's Inch"},
	"us-ft":  &Unit{0.304800609601219, "U.S. Surveyor's Foot"},
	"us-yd":  &Unit{0.914401828803658, "U.S. Surveyor's Yard"},
	"us-ch":  &Unit{20.11684023368047, "U.S. Surveyor's Chain"},
	"us-mi":  &Unit{1609.347218694437, "U.S. Surveyor's Statute Mile"},
	"ind-yd": &Unit{0.91439523, "Indian Yard"},
	"ind-ft": &Unit{0.30479841, "Indian Foot"},
	"ind-ch": &Unit{20.11669506, "Indian Chain"},
}
