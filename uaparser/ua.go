package uaparser

import "encoding/json"

// Borwser
type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Major   string `json:"major"`
}

// Cpu
type Cpu struct {
	Architecture string `json:"architecture"`
}

// Engine
type Engine struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// OS
type OS struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Device
type Device struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
	Type   string `json:"type"`
}

// UserAgent
type UserAgent struct {
	matcher *uaItem
	Ua      string `json:"ua"`
	parsedData
}

type parsedData struct {
	Browser `json:"browser"`
	Cpu     `json:"cpu"`
	Engine  `json:"engine"`
	Device  `json:"device"`
	OS      `json:"os"`
}

func ParseBrowser(ua string) Browser {
	u := NewUserAgent(ua)
	u.parse(UaBrowser)
	return u.GetBrowser()
}

func ParseCpu(ua string) Cpu {
	u := NewUserAgent(ua)
	u.parse(UaCpu)
	return u.GetCpu()
}

func ParseEngine(ua string) Engine {
	u := NewUserAgent(ua)
	u.parse(UaEngine)
	return u.GetEngine()
}

func ParseDevice(ua string) Device {
	u := NewUserAgent(ua)
	u.parse(UaDevice)
	return u.GetDevice()
}

func ParseOS(ua string) OS {
	u := NewUserAgent(ua)
	u.parse(UaOs)
	return u.GetOS()
}

func NewUserAgent(ua string) *UserAgent {
	userAgent := &UserAgent{
		Ua: ua,
	}
	userAgent.matcher = newUaItem(&userAgent.parsedData)
	return userAgent
}

func (ua *UserAgent) parse(itemType UaItemType) {
	ua.matcher.setUa(ua.Ua)
	ua.matcher.parseUa(itemType)
}

func (ua *UserAgent) GetResult() *UserAgent {
	ua.parse(UaResult)
	return ua
}

func (ua *UserAgent) SetUa(uaStr string) {
	ua.Ua = uaStr
}

func (ua *UserAgent) GetUa() string {
	return ua.Ua
}

func (ua *UserAgent) GetBrowser() Browser {
	return ua.Browser
}

func (ua *UserAgent) GetCpu() Cpu {
	return ua.Cpu
}

func (ua *UserAgent) GetEngine() Engine {
	return ua.Engine
}

func (ua *UserAgent) GetDevice() Device {
	return ua.Device
}

func (ua *UserAgent) GetOS() OS {
	return ua.OS
}

func (ua *UserAgent) String() string {
	s, _ := json.Marshal(ua)
	return string(s)
}

// IsMobile
func (d *Device) IsMobile() bool {
	if d.Type == MOBILE {
		return true
	}
	return false
}

// Parse
func Parse(ua string) *UserAgent {
	uagent := NewUserAgent(ua)
	uagent.parse(UaResult)
	return uagent
}

type uaItem struct {
	ua        []byte
	itemType  []UaItemType
	rgxMapper map[UaItemType]*rgxMapper

	d *parsedData
}

func newUaItem(data *parsedData) *uaItem {
	um := &uaItem{
		itemType:  make([]UaItemType, 0),
		rgxMapper: make(map[UaItemType]*rgxMapper),
		d:         data,
	}

	um.itemType = append(um.itemType, []UaItemType{UaBrowser, UaCpu, UaDevice, UaEngine, UaOs}...)
	um.createRgxMapper()
	return um
}

func (um *uaItem) createRgxMapper() {
	for _, itemType := range um.itemType {
		um.rgxMapper[itemType] = newRgxMapper(itemType)
	}
}

func (um *uaItem) setUa(ua string) {
	um.ua = []byte(ua)
}

func (um *uaItem) parseUa(itemType UaItemType) {
	switch itemType {
	case UaBrowser, UaCpu, UaDevice, UaEngine, UaOs:
		matcher := um.rgxMapper[itemType]
		if matcher != nil {
			matcher.rgxParse(um.ua, rgxMaps[itemType])
		}
	default:
		for t, matcher := range um.rgxMapper {
			matcher.rgxParse(um.ua, rgxMaps[t])
		}
	}
	um.extract()
}

func (um *uaItem) extract() {
	for itemType, matcher := range um.rgxMapper {
		if matcher.flag == RgxMapperIdle {
			continue
		}

		var matchData = matcher.Matches
		switch itemType {
		case UaBrowser:
			browser := &um.d.Browser
			extract(browser, matchData)
			browser.Major = majorize(browser.Version)
		case UaCpu:
			extract(&um.d.Cpu, matchData)
		case UaDevice:
			extract(&um.d.Device, matchData)
		case UaEngine:
			extract(&um.d.Engine, matchData)
		case UaOs:
			extract(&um.d.OS, matchData)
		}
		matcher.clear()
	}
}
