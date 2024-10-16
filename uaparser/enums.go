package uaparser

// UaItemType
type UaItemType string

const (
	UaBrowser UaItemType = "browser"
	UaCpu     UaItemType = "cpu"
	UaDevice  UaItemType = "device"
	UaEngine  UaItemType = "engine"
	UaOs      UaItemType = "os"
	UaResult  UaItemType = "result"
)

const (
	Empty       = ""
	Unknown     = "?"
	UaMaxLength = 500
)

// fields
const (
	Major        = "major"
	Model        = "model"
	Name         = "name"
	Type         = "type"
	Vendor       = "vendor"
	Version      = "version"
	Architecture = "architecture"
)

// device brand
const (
	AMAZON         = "Amazon"
	APPLE          = "Apple"
	ASUS           = "ASUS"
	BLACKBERRY     = "BlackBerry"
	BROWSER        = "Browser"
	CHROME         = "Chrome"
	CHROMECAST     = "Chromecast"
	EDGE           = "Edge"
	FIREFOX        = "Firefox"
	GOOGLE         = "Google"
	HUAWEI         = "Huawei"
	LENOVO         = "Lenovo"
	LG             = "LG"
	MICROSOFT      = "Microsoft"
	MOTOROLA       = "Motorola"
	OPERA          = "Opera"
	SAMSUNG        = "Samsung"
	SHARP          = "Sharp"
	SONY           = "Sony"
	XIAOMI         = "Xiaomi"
	ZEBRA          = "Zebra"
	PREFIX_MOBILE  = "Mobile "
	SUFFIX_BROWSER = " Browser"
	FACEBOOK       = "Facebook"
	IE             = "IE"
	SOGOU          = "Sogou"
	OPPO           = "OPPO"
	WINDOWS        = "Windows"
	NOKIA          = "Nokia"
	HTC            = "HTC"
	ACER           = "Acer"
	MEIZU          = "Meizu"
	TCL            = "TCL"
)

// device type
const (
	MOBILE   = "mobile"  // 手机
	TABLET   = "tablet"  // 平板
	SMARTTV  = "smarttv" // tv
	CONSOLE  = "console"
	WEARABLE = "wearable"
	XR       = "xr"
	EMBEDDED = "embedded"
)

// extensions
const (
	CLI     = "cli"
	CRAWLER = "crawler"
	FETCHER = "fetcher"
)

// old version to x.x.x
var safariMap = map[string]string{
	"1.0":   "/8",
	"1.2":   "/1",
	"1.3":   "/3",
	"2.0":   "/412",
	"2.0.2": "/416",
	"2.0.3": "/417",
	"2.0.4": "/419",
}

// special windowsVersion
var windowsVersionMap = map[string][]string{
	"ME":      {"4.90"},
	"NT 3.11": {"NT3.51"},
	"NT 4.0":  {"NT4.0"},
	"2000":    {"NT 5.0"},
	"XP":      {"NT 5.1", "NT 5.2"},
	"Vista":   {"NT 6.0"},
	"7":       {"NT 6.1"},
	"8":       {"NT 6.2"},
	"8.1":     {"NT 6.3"},
	"10":      {"NT 6.4", "NT 10.0"},
	"RT":      {"ARM"},
}

var wildcardsType = map[string]string{"*": "mobile"}

// itel special type
var itelType = map[string][]string{
	"tablet": {"p10001l", "w7001"},
}
