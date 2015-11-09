package ref

type Ellipse struct {
	Params      string
	Description string
}

var EllipseList = map[string]*Ellipse{
	"MERIT":     &Ellipse{"a=6378137.0 rf=298.257", "MERIT 1983"},
	"SGS85":     &Ellipse{"a=6378136.0 rf=298.257", "Soviet Geodetic System 85"},
	"GRS80":     &Ellipse{"a=6378137.0 rf=298.257222101", "GRS 1980(IUGG, 1980)"},
	"IAU76":     &Ellipse{"a=6378140.0 rf=298.257", "IAU 1976"},
	"airy":      &Ellipse{"a=6377563.396 b=6356256.910", "Airy 1830"},
	"APL4.9":    &Ellipse{"a=6378137.0. rf=298.25", "Appl. Physics. 1965"},
	"NWL9D":     &Ellipse{"a=6378145.0. rf=298.25", "Naval Weapons Lab., 1965"},
	"mod_airy":  &Ellipse{"a=6377340.189 b=6356034.446", "Modified Airy"},
	"andrae":    &Ellipse{"a=6377104.43 rf=300.0", "Andrae 1876 (Den., Iclnd.)"},
	"aust_SA":   &Ellipse{"a=6378160.0 rf=298.25", "Australian Natl & S. Amer. 1969"},
	"GRS67":     &Ellipse{"a=6378160.0 rf=298.2471674270", "GRS 67(IUGG 1967)"},
	"bessel":    &Ellipse{"a=6377397.155 rf=299.1528128", "Bessel 1841"},
	"bess_nam":  &Ellipse{"a=6377483.865 rf=299.1528128", "Bessel 1841 (Namibia)"},
	"clrk66":    &Ellipse{"a=6378206.4 b=6356583.8", "Clarke 1866"},
	"clrk80":    &Ellipse{"a=6378249.145 rf=293.4663", "Clarke 1880 mod."},
	"clrk80ign": &Ellipse{"a=6378249.2 rf=293.4660212936269", "Clarke 1880 (IGN)."},
	"CPM":       &Ellipse{"a=6375738.7 rf=334.29", "Comm. des Poids et Mesures 1799"},
	"delmbr":    &Ellipse{"a=6376428. rf=311.5", "Delambre 1810 (Belgium)"},
	"engelis":   &Ellipse{"a=6378136.05 rf=298.2566", "Engelis 1985"},
	"evrst30":   &Ellipse{"a=6377276.345 rf=300.8017", "Everest 1830"},
	"evrst48":   &Ellipse{"a=6377304.063 rf=300.8017", "Everest 1948"},
	"evrst56":   &Ellipse{"a=6377301.243 rf=300.8017", "Everest 1956"},
	"evrst69":   &Ellipse{"a=6377295.664 rf=300.8017", "Everest 1969"},
	"evrstSS":   &Ellipse{"a=6377298.556 rf=300.8017", "Everest (Sabah & Sarawak)"},
	"fschr60":   &Ellipse{"a=6378166. rf=298.3", "Fischer (Mercury Datum) 1960"},
	"fschr60m":  &Ellipse{"a=6378155. rf=298.3", "Modified Fischer 1960"},
	"fschr68":   &Ellipse{"a=6378150. rf=298.3", "Fischer 1968"},
	"helmert":   &Ellipse{"a=6378200. rf=298.3", "Helmert 1906"},
	"hough":     &Ellipse{"a=6378270.0 rf=297.", "Hough"},
	"intl":      &Ellipse{"a=6378388.0 rf=297.", "International 1909 (Hayford)"},
	"krass":     &Ellipse{"a=6378245.0 rf=298.3", "Krassovsky, 1942"},
	"kaula":     &Ellipse{"a=6378163. rf=298.24", "Kaula 1961"},
	"lerch":     &Ellipse{"a=6378139. rf=298.257", "Lerch 1979"},
	"mprts":     &Ellipse{"a=6397300. rf=191.", "Maupertius 1738"},
	"new_intl":  &Ellipse{"a=6378157.5 b=6356772.2", "New International 1967"},
	"plessis":   &Ellipse{"a=6376523.0 b=6355863.", "Plessis 1817 (France)"},
	"SEasia":    &Ellipse{"a=6378155.0 b=6356773.3205", "Southeast Asia"},
	"walbeck":   &Ellipse{"a=6376896.0 b=6355834.8467", "Walbeck"},
	"WGS60":     &Ellipse{"a=6378165.0 rf=298.3", "WGS 60"},
	"WGS66":     &Ellipse{"a=6378145.0 rf=298.25", "WGS 66"},
	"WGS72":     &Ellipse{"a=6378135.0 rf=298.26", "WGS 72"},
	"WGS84":     &Ellipse{"a=6378137.0 rf=298.257223563", "WGS 84"},
	"sphere":    &Ellipse{"a=6370997.0 b=6370997.0", "Normal Sphere (r=6370997)"},
}