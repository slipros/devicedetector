package devicedetector

import (
	"embed"
	"path"
	"strings"

	regexp "github.com/dlclark/regexp2"
	gover "github.com/mcuadros/go-version"

	"github.com/slipros/devicedetector/parser"
	"github.com/slipros/devicedetector/parser/client"
	"github.com/slipros/devicedetector/parser/device"
)

//go:embed regexes/*
var EmbeddedRegexes embed.FS

const UNKNOWN = "UNK"
const VERSION = `6.4.7`

var desktopOsArray = []string{
	`AmigaOS`, `IBM`, `GNU/Linux`, `Mac`, `Unix`, `Windows`, `BeOS`, `Chrome OS`, `Chromium OS`,
}

var (
	chrMobReg = regexp.MustCompile(fixUserAgentRegEx(`Chrome/[\.0-9]* Mobile`), regexp.IgnoreCase)
	chrTabReg = regexp.MustCompile(fixUserAgentRegEx(`Chrome/[\.0-9]* (?!Mobile)`), regexp.IgnoreCase)
	opaTabReg = regexp.MustCompile(fixUserAgentRegEx(`Opera Tablet`), regexp.IgnoreCase)
	opaTvReg  = regexp.MustCompile(fixUserAgentRegEx(`Opera TV Store`), regexp.IgnoreCase)
)

func fixUserAgentRegEx(regex string) string {
	reg := strings.ReplaceAll(regex, `/`, `\/`)
	reg = strings.ReplaceAll(reg, `++`, `+`)

	return `(?:^|[^A-Z_-])(?:` + reg + `)`
}

type DeviceDetector struct {
	deviceParsers         []device.DeviceParser
	clientParsers         []client.ClientParser
	botParsers            []parser.BotParser
	osParsers             []parser.OsParser
	vendorParser          *parser.VendorFragments
	DiscardBotInformation bool
	SkipBotDetection      bool
}

func NewDeviceDetector() (*DeviceDetector, error) {
	parser.ReadFile = EmbeddedRegexes.ReadFile

	return newDeviceDetector("./regexes")
}

func newDeviceDetector(dir string) (*DeviceDetector, error) {
	vp, err := parser.NewVendor(path.Join(dir, parser.FixtureFileVendor))
	if err != nil {
		return nil, err
	}
	osp, err := parser.NewOss(path.Join(dir, parser.FixtureFileOs))
	if err != nil {
		return nil, err
	}

	return &DeviceDetector{
		vendorParser: vp,
		osParsers:    []parser.OsParser{osp},
		clientParsers: client.NewClientParsers(
			path.Join(dir, "client"),
			[]string{
				client.ParserNameFeedReader,
				client.ParserNameMobileApp,
				client.ParserNameMediaPlayer,
				client.ParserNamePim,
				client.ParserNameBrowser,
				client.ParserNameLibrary,
			},
		),
		deviceParsers: device.NewDeviceParsers(
			path.Join(dir, "device"),
			[]string{
				device.ParserNameHbbTv,
				device.ParserNameConsole,
				device.ParserNameCar,
				device.ParserNameCamera,
				device.ParserNamePortableMediaPlayer,
				device.ParserNameMobile,
			},
		),
		botParsers: []parser.BotParser{
			parser.NewBot(path.Join(dir, parser.FixtureFileBot)),
		},
	}, nil
}

func (d *DeviceDetector) AddClientParser(cp client.ClientParser) {
	d.clientParsers = append(d.clientParsers, cp)
}

func (d *DeviceDetector) GetClientParser() []client.ClientParser {
	return d.clientParsers
}

func (d *DeviceDetector) AddDeviceParser(dp device.DeviceParser) {
	d.deviceParsers = append(d.deviceParsers, dp)
}

func (d *DeviceDetector) GetDeviceParsers() []device.DeviceParser {
	return d.deviceParsers
}

func (d *DeviceDetector) AddBotParser(op parser.BotParser) {
	d.botParsers = append(d.botParsers, op)
}

func (d *DeviceDetector) GetBotParsers() []parser.BotParser {
	return d.botParsers
}

func (d *DeviceDetector) ParseBot(ua string) *parser.BotMatchResult {
	if !d.SkipBotDetection {
		for _, p := range d.botParsers {
			p.DiscardDetails(d.DiscardBotInformation)
			if r := p.Parse(ua); r != nil {
				return r
			}
		}
	}
	return nil
}

func (d *DeviceDetector) ParseOs(ua string) *parser.OsMatchResult {
	for _, p := range d.osParsers {
		if r := p.Parse(ua); r != nil {
			return r
		}
	}
	return nil
}

func (d *DeviceDetector) ParseClient(ua string) *client.ClientMatchResult {
	for _, p := range d.clientParsers {
		if r := p.Parse(ua); r != nil {
			return r
		}
	}
	return nil
}

func (d *DeviceDetector) ParseDevice(ua string) *device.DeviceMatchResult {
	for _, p := range d.deviceParsers {
		if r := p.Parse(ua); r != nil {
			return r
		}
	}
	return nil
}

func (d *DeviceDetector) parseInfo(info *DeviceInfo) {
	ua := info.userAgent
	if r := d.ParseDevice(ua); r != nil {
		info.Type = r.Type
		info.Model = r.Model
		info.Brand = r.Brand
	}
	// If no brand has been assigned try to match by known vendor fragments
	if info.Brand == "" && d.vendorParser != nil {
		info.Brand = d.vendorParser.Parse(ua)
	}

	os := info.GetOs()
	osShortName := os.ShortName
	osFamily := parser.GetOsFamily(osShortName)
	osVersion := os.Version
	cmr := info.GetClient()

	if info.Brand == "" && (osShortName == `ATV` || osShortName == `IOS` || osShortName == `MAC`) {
		info.Brand = `AP`
	}

	deviceType := parser.GetDeviceType(info.Type)
	// Chrome on Android passes the device type based on the keyword 'Mobile'
	// If it is present the device should be a smartphone, otherwise it's a tablet
	// See https://developer.chrome.com/multidevice/user-agent#chrome_for_android_user_agent
	if deviceType == parser.DEVICE_TYPE_INVALID && osFamily == `Android` {
		if browserName, ok := client.GetBrowserFamily(cmr.ShortName); ok && browserName == `Chrome` {
			if ok, _ := chrMobReg.MatchString(ua); ok {
				deviceType = parser.DEVICE_TYPE_SMARTPHONE
			} else if ok, _ = chrTabReg.MatchString(ua); ok {
				deviceType = parser.DEVICE_TYPE_TABLET
			}
		}
	}

	if deviceType == parser.DEVICE_TYPE_INVALID {
		if info.HasAndroidMobileFragment() {
			deviceType = parser.DEVICE_TYPE_TABLET
		} else if ok, _ := opaTabReg.MatchString(ua); ok {
			deviceType = parser.DEVICE_TYPE_TABLET
		} else if info.HasAndroidMobileFragment() {
			deviceType = parser.DEVICE_TYPE_SMARTPHONE
		} else if osShortName == "AND" && osVersion != "" {
			if gover.CompareSimple(osVersion, `2.0`) == -1 {
				deviceType = parser.DEVICE_TYPE_SMARTPHONE
			} else if gover.CompareSimple(osVersion, `3.0`) >= 0 &&
				gover.CompareSimple(osVersion, `4.0`) == -1 {
				deviceType = parser.DEVICE_TYPE_TABLET
			}
		}
	}

	// All detected feature phones running android are more likely a smartphone
	if deviceType == parser.DEVICE_TYPE_FEATURE_PHONE && osFamily == `Android` {
		deviceType = parser.DEVICE_TYPE_SMARTPHONE
	}

	// According to http://msdn.microsoft.com/en-us/library/ie/hh920767(v=vs.85).aspx
	if deviceType == parser.DEVICE_TYPE_INVALID &&
		(osShortName == `WRT` || (osShortName == `WIN` && gover.CompareSimple(osVersion, `8`) >= 0)) &&
		info.IsTouchEnabled() {
		deviceType = parser.DEVICE_TYPE_TABLET
	}

	// All devices running Opera TV Store are assumed to be a tv
	if ok, _ := opaTvReg.MatchString(ua); ok {
		deviceType = parser.DEVICE_TYPE_TV
	}

	// Devices running Kylo or Espital TV Browsers are assumed to be a TV
	if deviceType == parser.DEVICE_TYPE_INVALID {
		if cmr.Name == `Kylo` || cmr.Name == `Espial TV Browser` {
			deviceType = parser.DEVICE_TYPE_TV
		} else if info.IsDesktop() {
			deviceType = parser.DEVICE_TYPE_DESKTOP
		}
	}

	if deviceType != parser.DEVICE_TYPE_INVALID {
		info.Type = parser.GetDeviceName(deviceType)
	}
	return
}

func (d *DeviceDetector) Parse(ua string) *DeviceInfo {
	// skip parsing for empty useragents or those not containing any letter
	if !parser.StringContainsLetter(ua) {
		return nil
	}

	info := &DeviceInfo{
		userAgent: ua,
	}

	info.bot = d.ParseBot(ua)
	if info.IsBot() {
		return info
	}

	info.os = d.ParseOs(ua)

	// Parse Clients
	// Clients might be browsers, Feed Readers, Mobile Apps, Media Players or
	// any other application accessing with an parseable UA
	info.client = d.ParseClient(ua)

	d.parseInfo(info)

	return info
}
