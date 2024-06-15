package ascii

var Colors = map[string]string{
	"red":                  "rgb(255,0,0)",
	"green":                "rgb(0,255,0)",
	"blue":                 "rgb(0,0,255)",
	"yellow":               "rgb(255,255,0)",
	"cyan":                 "rgb(0,255,255)",
	"magenta":              "rgb(255,0,255)",
	"black":                "rgb(0,0,0)",
	"white":                "rgb(255,255,255)",
	"gray":                 "rgb(128,128,128)",
	"orange":               "rgb(255,165,0)",
	"purple":               "rgb(128,0,128)",
	"brown":                "rgb(165,42,42)",
	"pink":                 "rgb(255,192,203)",
	"lime":                 "rgb(0,255,0)",
	"navy":                 "rgb(0,0,128)",
	"darkred":              "rgb(139,0,0)",
	"darkgreen":            "rgb(0,100,0)",
	"darkblue":             "rgb(0,0,139)",
	"gold":                 "rgb(255,215,0)",
	"silver":               "rgb(192,192,192)",
	"lightgray":            "rgb(211,211,211)",
	"darkgray":             "rgb(169,169,169)",
	"lightgreen":           "rgb(144,238,144)",
	"lightblue":            "rgb(173,216,230)",
	"darkcyan":             "rgb(0,139,139)",
	"darkmagenta":          "rgb(139,0,139)",
	"beige":                "rgb(245,245,220)",
	"khaki":                "rgb(240,230,140)",
	"teal":                 "rgb(0,128,128)",
	"coral":                "rgb(255,127,80)",
	"salmon":               "rgb(250,128,114)",
	"turquoise":            "rgb(64,224,208)",
	"plum":                 "rgb(221,160,221)",
	"orchid":               "rgb(218,112,214)",
	"goldenrod":            "rgb(218,165,32)",
	"lavender":             "rgb(230,230,250)",
	"tan":                  "rgb(210,180,140)",
	"chocolate":            "rgb(210,105,30)",
	"sienna":               "rgb(160,82,45)",
	"maroon":               "rgb(128,0,0)",
	"olive":                "rgb(128,128,0)",
	"peachpuff":            "rgb(255,218,185)",
	"mintcream":            "rgb(245,255,250)",
	"mistyrose":            "rgb(255,228,225)",
	"lavenderblush":        "rgb(255,240,245)",
	"honeydew":             "rgb(240,255,240)",
	"aliceblue":            "rgb(240,248,255)",
	"ghostwhite":           "rgb(248,248,255)",
	"whitesmoke":           "rgb(245,245,245)",
	"seashell":             "rgb(255,245,238)",
	"floralwhite":          "rgb(255,250,240)",
	"oldlace":              "rgb(253,245,230)",
	"linen":                "rgb(250,240,230)",
	"antiquewhite":         "rgb(250,235,215)",
	"papayawhip":           "rgb(255,239,213)",
	"blanchedalmond":       "rgb(255,235,205)",
	"bisque":               "rgb(255,228,196)",
	"wheat":                "rgb(245,222,179)",
	"moccasin":             "rgb(255,228,181)",
	"navajowhite":          "rgb(255,222,173)",
	"peach":                "rgb(255,218,185)",
	"yellowgreen":          "rgb(154,205,50)",
	"darkolivegreen":       "rgb(85,107,47)",
	"olivedrab":            "rgb(107,142,35)",
	"greenyellow":          "rgb(173,255,47)",
	"lawngreen":            "rgb(124,252,0)",
	"chartreuse":           "rgb(127,255,0)",
	"mediumseagreen":       "rgb(60,179,113)",
	"springgreen":          "rgb(0,255,127)",
	"seagreen":             "rgb(46,139,87)",
	"mediumaquamarine":     "rgb(102,205,170)",
	"mediumspringgreen":    "rgb(0,250,154)",
	"palegreen":            "rgb(152,251,152)",
	"lightseagreen":        "rgb(32,178,170)",
	"forestgreen":          "rgb(34,139,34)",
	"limegreen":            "rgb(50,205,50)",
	"darkseagreen":         "rgb(143,188,143)",
	"darkslategray":        "rgb(47,79,79)",
	"mediumturquoise":      "rgb(72,209,204)",
	"darkturquoise":        "rgb(0,206,209)",
	"lightcyan":            "rgb(224,255,255)",
	"paleturquoise":        "rgb(175,238,238)",
	"azure":                "rgb(240,255,255)",
	"powderblue":           "rgb(176,224,230)",
	"cadetblue":            "rgb(95,158,160)",
	"steelblue":            "rgb(70,130,180)",
	"cornflowerblue":       "rgb(100,149,237)",
	"deepskyblue":          "rgb(0,191,255)",
	"dodgerblue":           "rgb(30,144,255)",
	"lightsteelblue":       "rgb(176,196,222)",
	"skyblue":              "rgb(135,206,235)",
	"lightskyblue":         "rgb(135,206,250)",
	"midnightblue":         "rgb(25,25,112)",
	"slateblue":            "rgb(106,90,205)",
	"royalblue":            "rgb(65,105,225)",
	"blueviolet":           "rgb(138,43,226)",
	"darkviolet":           "rgb(148,0,211)",
	"darkorchid":           "rgb(153,50,204)",
	"mediumorchid":         "rgb(186,85,211)",
	"mediumpurple":         "rgb(147,112,219)",
	"amethyst":             "rgb(153,102,204)",
	"thistle":              "rgb(216,191,216)",
	"violet":               "rgb(238,130,238)",
	"mediumvioletred":      "rgb(199,21,133)",
	"palevioletred":        "rgb(219,112,147)",
	"hotpink":              "rgb(255,105,180)",
	"deepink":              "rgb(255,20,147)",
	"lightpink":            "rgb(255,182,193)",
	"rosybrown":            "rgb(188,143,143)",
	"indianred":            "rgb(205,92,92)",
	"firebrick":            "rgb(178,34,34)",
	"lightcoral":           "rgb(240,128,128)",
	"darkorange":           "rgb(255,140,0)",
	"tomato":               "rgb(255,99,71)",
	"orangered":            "rgb(255,69,0)",
	"crimson":              "rgb(220,20,60)",
	"palegoldenrod":        "rgb(238,232,170)",
	"lemonchiffon":         "rgb(255,250,205)",
	"lightgoldenrodyellow": "rgb(250,250,210)",
	"lightyellow":          "rgb(255,255,224)",
}
