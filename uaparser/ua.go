package uaparser

import "encoding/json"

type OptFunc func(ua *UserAgent)

// Borwser
type Browser struct {
	Name    string `json:"name" field:"name"`
	Version string `json:"version" field:"version"`
	Major   string `json:"major" field:"major"`
	Type    string `json:"-" field:"type"`
}

// Cpu
type Cpu struct {
	Architecture string `json:"architecture" field:"architecture"`
}

// Engine
type Engine struct {
	Name    string `json:"name" field:"name"`
	Version string `json:"version" field:"version"`
}

// OS
type OS struct {
	Name    string `json:"name" field:"name"`
	Version string `json:"version" field:"version"`
}

// Device
type Device struct {
	Vendor string `json:"vendor" field:"vendor"`
	Model  string `json:"model" field:"model"`
	Type   string `json:"type" field:"type"`
}

// UserAgent
type UserAgent struct {
	matcher          *uaItem
	supportExtension []string

	Ua string `json:"ua"`
	parsedData
}

// parsedData
type parsedData struct {
	Browser `json:"browser"`
	Cpu     `json:"cpu"`
	Engine  `json:"engine"`
	Device  `json:"device"`
	OS      `json:"os"`
}

func ParseBrowser(ua string, opts ...OptFunc) Browser {
	u := NewUserAgent(ua, opts...)
	u.parse(UaBrowser)
	return u.GetBrowser()
}

func ParseCpu(ua string, opts ...OptFunc) Cpu {
	u := NewUserAgent(ua, opts...)
	u.parse(UaCpu)
	return u.GetCpu()
}

func ParseEngine(ua string, opts ...OptFunc) Engine {
	u := NewUserAgent(ua, opts...)
	u.parse(UaEngine)
	return u.GetEngine()
}

func ParseDevice(ua string, opts ...OptFunc) Device {
	u := NewUserAgent(ua, opts...)
	u.parse(UaDevice)
	return u.GetDevice()
}

func ParseOS(ua string, opts ...OptFunc) OS {
	u := NewUserAgent(ua, opts...)
	u.parse(UaOs)
	return u.GetOS()
}

func SetExtensions(extSupport []string) OptFunc {
	return func(ua *UserAgent) {
		ua.supportExtension = extSupport
	}
}

func NewUserAgent(ua string, opts ...OptFunc) *UserAgent {
	userAgent := &UserAgent{
		Ua:               ua,
		supportExtension: make([]string, 0),
	}

	for _, fc := range opts {
		fc(userAgent)
	}

	userAgent.matcher = newUaItem(&userAgent.parsedData, userAgent.supportExtension)
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

func (ua *UserAgent) IsCli() bool {
	return ua.Browser.IsCli()
}

func (ua *UserAgent) IsBot() bool {
	return ua.Browser.IsBot()
}

func (ua *UserAgent) IsMobile() bool {
	return ua.Device.IsMobile()
}

// IsCli
func (b *Browser) IsCli() bool {
	if b.Type == CLI {
		return true
	}
	return false
}

// IsBot
func (b *Browser) IsBot() bool {
	if b.Type == CRAWLER || b.Type == FETCHER {
		return true
	}
	return false
}

// IsMobile
func (d *Device) IsMobile() bool {
	if d.Type == MOBILE {
		return true
	}
	return false
}

// Parse
func Parse(ua string, opts ...OptFunc) *UserAgent {
	uagent := NewUserAgent(ua, opts...)
	uagent.parse(UaResult)
	return uagent
}

type uaItem struct {
	ua         []byte
	itemType   []UaItemType
	rgxMapper  map[UaItemType]*rgxMapper
	extensions []string

	d *parsedData
}

func newUaItem(data *parsedData, supportExtension []string) *uaItem {
	um := &uaItem{
		itemType:   make([]UaItemType, 0),
		rgxMapper:  make(map[UaItemType]*rgxMapper),
		d:          data,
		extensions: supportExtension,
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
			matcher.rgxParse(um.ua, um.extend(itemType, rgxMaps[itemType]))
		}
	default:
		for t, matcher := range um.rgxMapper {
			matcher.rgxParse(um.ua, um.extend(t, rgxMaps[t]))
		}
	}
	um.extract()
}

func (um *uaItem) extend(itemType UaItemType, rgxParsers []*Parser) []*Parser {
	for _, ext := range um.extensions {
		var rgxExtensions map[UaItemType][]*Parser
		switch ext {
		case CLI:
			rgxExtensions = rgxExtensionCli
		case CRAWLER:
			rgxExtensions = rgxCrawlersExtensions
		case FETCHER:
			rgxExtensions = rgxFetchersExtensions
		}
		if extendParses, ok := rgxExtensions[itemType]; ok {
			rgxParsers = append(extendParses, rgxParsers...)
		}
	}
	return rgxParsers
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
