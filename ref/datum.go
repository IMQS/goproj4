package ref

type Datum struct {
	Definition  string
	Ellipse     string
	Description string
}

var DatumList = map[string]*Datum{
	"WGS84":         &Datum{"towgs84=0,0,0", "WGS84", ""},
	"GGRS87":        &Datum{"towgs84=-199.87,74.79,246.62", "GRS80", "Greek_Geodetic_Reference_System_1987"},
	"NAD83":         &Datum{"towgs84=0,0,0", "GRS80", "North_American_Datum_1983"},
	"NAD27":         &Datum{"nadgrids=@conus,@alaska,@ntv2_0.gsb,@ntv1_can.dat", "clrk66", "North_American_Datum_1927"},
	"potsdam":       &Datum{"towgs84=598.1,73.7,418.2,0.202,0.045,-2.455,6.7", "bessel", "Potsdam Rauenberg 1950 DHDN"},
	"carthage":      &Datum{"towgs84=-263.0,6.0,431.0", "clrk80ign", "Carthage 1934 Tunisia"},
	"hermannskogel": &Datum{"towgs84=577.326,90.129,463.919,5.137,1.474,5.297,2.4232", "bessel", "Hermannskogel"},
	"ire65":         &Datum{"towgs84=482.530,-130.596,564.557,-1.042,-0.214,-0.631,8.15", "mod_airy", "Ireland 1965"},
	"nzgd49":        &Datum{"towgs84=59.47,-5.04,187.44,0.47,-0.1,1.024,-4.5993", "intl", "New Zealand Geodetic Datum 1949"},
	"OSGB36":        &Datum{"towgs84=446.448,-125.157,542.060,0.1502,0.2470,0.8421,-20.4894", "airy", "Airy 1830"},
}
