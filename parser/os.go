package parser

import (
	"strings"
)

const ParserNameOs = "os"
const FixtureFileOs = "oss.yml"

type OsReg struct {
	Regular `yaml:",inline" json:",inline"`
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}

// Known operating systems mapped to their internal short codes
var OperatingSystems = map[string]string{
	`AIX`: `AIX`,
	`AND`: `Android`,
	`ADR`: `Android TV`,
	`ALP`: `Alpine Linux`,
	`AMZ`: `Amazon Linux`,
	`AMG`: `AmigaOS`,
	`ARM`: `Armadillo OS`,
	`ARO`: `AROS`,
	`ATV`: `tvOS`,
	`ARL`: `Arch Linux`,
	`AOS`: `AOSC OS`,
	`ASP`: `ASPLinux`,
	`AZU`: `Azure Linux`,
	`BTR`: `BackTrack`,
	`SBA`: `Bada`,
	`BYI`: `Baidu Yi`,
	`BEO`: `BeOS`,
	`BLB`: `BlackBerry OS`,
	`QNX`: `BlackBerry Tablet OS`,
	`PAN`: `blackPanther OS`,
	`BOS`: `Bliss OS`,
	`BMP`: `Brew`,
	`BSN`: `BrightSignOS`,
	`CAI`: `Caixa Mágica`,
	`CES`: `CentOS`,
	`CST`: `CentOS Stream`,
	`CLO`: `Clear Linux OS`,
	`CLR`: `ClearOS Mobile`,
	`COS`: `Chrome OS`,
	`CRS`: `Chromium OS`,
	`CHN`: `China OS`,
	`COL`: `Coolita OS`,
	`CYN`: `CyanogenMod`,
	`DEB`: `Debian`,
	`DEE`: `Deepin`,
	`DFB`: `DragonFly`,
	`DVK`: `DVKBuntu`,
	`ELE`: `ElectroBSD`,
	`EUL`: `EulerOS`,
	`FED`: `Fedora`,
	`FEN`: `Fenix`,
	`FOS`: `Firefox OS`,
	`FIR`: `Fire OS`,
	`FOR`: `Foresight Linux`,
	`FRE`: `Freebox`,
	`BSD`: `FreeBSD`,
	`FRI`: `FRITZ!OS`,
	`FYD`: `FydeOS`,
	`FUC`: `Fuchsia`,
	`GNT`: `Gentoo`,
	`GNX`: `GENIX`,
	`GEO`: `GEOS`,
	`GNS`: `gNewSense`,
	`GRI`: `GridOS`,
	`GTV`: `Google TV`,
	`HPX`: `HP-UX`,
	`HAI`: `Haiku OS`,
	`IPA`: `iPadOS`,
	`HAR`: `HarmonyOS`,
	`HAS`: `HasCodingOS`,
	`HEL`: `HELIX OS`,
	`IRI`: `IRIX`,
	`INF`: `Inferno`,
	`JME`: `Java ME`,
	`JOL`: `Joli OS`,
	`KOS`: `KaiOS`,
	`KAL`: `Kali`,
	`KAN`: `Kanotix`,
	`KIN`: `KIN OS`,
	`KOL`: `KolibriOS`,
	`KNO`: `Knoppix`,
	`KTV`: `KreaTV`,
	`KBT`: `Kubuntu`,
	`LIN`: `GNU/Linux`,
	`LEA`: `LeafOS`,
	`LND`: `LindowsOS`,
	`LNS`: `Linspire`,
	`LEN`: `Lineage OS`,
	`LIR`: `Liri OS`,
	`LOO`: `Loongnix`,
	`LBT`: `Lubuntu`,
	`LOS`: `Lumin OS`,
	`LUN`: `LuneOS`,
	`VLN`: `VectorLinux`,
	`MAC`: `Mac`,
	`MAE`: `Maemo`,
	`MAG`: `Mageia`,
	`MDR`: `Mandriva`,
	`SMG`: `MeeGo`,
	`MET`: `Meta Horizon`,
	`MCD`: `MocorDroid`,
	`MON`: `moonOS`,
	`EZX`: `Motorola EZX`,
	`MIN`: `Mint`,
	`MLD`: `MildWild`,
	`MOR`: `MorphOS`,
	`NBS`: `NetBSD`,
	`MTK`: `MTK / Nucleus`,
	`MRE`: `MRE`,
	`NXT`: `NeXTSTEP`,
	`NWS`: `NEWS-OS`,
	`WII`: `Nintendo`,
	`NDS`: `Nintendo Mobile`,
	`NOV`: `Nova`,
	`OS2`: `OS/2`,
	`T64`: `OSF1`,
	`OBS`: `OpenBSD`,
	`OHS`: `OpenHarmony`,
	`OVS`: `OpenVMS`,
	`OVZ`: `OpenVZ`,
	`OWR`: `OpenWrt`,
	`OTV`: `Opera TV`,
	`ORA`: `Oracle Linux`,
	`ORD`: `Ordissimo`,
	`PAR`: `Pardus`,
	`PCL`: `PCLinuxOS`,
	`PIC`: `PICO OS`,
	`PLA`: `Plasma Mobile`,
	`PSP`: `PlayStation Portable`,
	`PS3`: `PlayStation`,
	`PVE`: `Proxmox VE`,
	`PUF`: `Puffin OS`,
	`PUR`: `PureOS`,
	`QTP`: `Qtopia`,
	`PIO`: `Raspberry Pi OS`,
	`RAS`: `Raspbian`,
	`RHT`: `Red Hat`,
	`RST`: `Red Star`,
	`RED`: `RedOS`,
	`REV`: `Revenge OS`,
	`RIS`: `risingOS`,
	`ROS`: `RISC OS`,
	`ROC`: `Rocky Linux`,
	`ROK`: `Roku OS`,
	`RSO`: `Rosa`,
	`ROU`: `RouterOS`,
	`REM`: `Remix OS`,
	`RRS`: `Resurrection Remix OS`,
	`REX`: `REX`,
	`RZD`: `RazoDroiD`,
	`RXT`: `RTOS & Next`,
	`SAB`: `Sabayon`,
	`SSE`: `SUSE`,
	`SAF`: `Sailfish OS`,
	`SCI`: `Scientific Linux`,
	`SEE`: `SeewoOS`,
	`SER`: `SerenityOS`,
	`SIR`: `Sirin OS`,
	`SLW`: `Slackware`,
	`SOS`: `Solaris`,
	`SBL`: `Star-Blade OS`,
	`SYL`: `Syllable`,
	`SYM`: `Symbian`,
	`SYS`: `Symbian OS`,
	`S40`: `Symbian OS Series 40`,
	`S60`: `Symbian OS Series 60`,
	`SY3`: `Symbian^3`,
	`TEN`: `TencentOS`,
	`TDX`: `ThreadX`,
	`TIZ`: `Tizen`,
	`TIV`: `TiVo OS`,
	`TOS`: `TmaxOS`,
	`TUR`: `Turbolinux`,
	`UBT`: `Ubuntu`,
	`ULT`: `ULTRIX`,
	`UOS`: `UOS`,
	`VID`: `VIDAA`,
	`VIZ`: `ViziOS`,
	`WAS`: `watchOS`,
	`WER`: `Wear OS`,
	`WTV`: `WebTV`,
	`WHS`: `Whale OS`,
	`WIN`: `Windows`,
	`WCE`: `Windows CE`,
	`WIO`: `Windows IoT`,
	`WMO`: `Windows Mobile`,
	`WPH`: `Windows Phone`,
	`WRT`: `Windows RT`,
	`WPO`: `WoPhone`,
	`XBX`: `Xbox`,
	`XBT`: `Xubuntu`,
	`YNS`: `YunOS`,
	`ZEN`: `Zenwalk`,
	`ZOR`: `ZorinOS`,
	`IOS`: `iOS`,
	`POS`: `palmOS`,
	`WEB`: `Webian`,
	`WOS`: `webOS`,
}

// Operating system families mapped to the short codes of the associated operating systems
var OsFamilies = map[string][]string{
	`Android`: {
		`AND`, `CYN`, `FIR`, `REM`, `RZD`, `MLD`, `MCD`, `YNS`, `GRI`, `HAR`,
		`ADR`, `CLR`, `BOS`, `REV`, `LEN`, `SIR`, `RRS`, `WER`, `PIC`, `ARM`,
		`HEL`, `BYI`, `RIS`, `PUF`, `LEA`, `MET`, `OHS`,
	},
	`AmigaOS`:        {`AMG`, `MOR`, `ARO`},
	`BlackBerry`:     {`BLB`, `QNX`},
	`Brew`:           {`BMP`},
	`BeOS`:           {`BEO`, `HAI`},
	`Chrome OS`:      {`COS`, `CRS`, `FYD`, `SEE`},
	`Firefox OS`:     {`FOS`, `KOS`},
	`Gaming Console`: {`WII`, `PS3`},
	`Google TV`:      {`GTV`},
	`IBM`:            {`OS2`},
	`iOS`:            {`IOS`, `ATV`, `WAS`, `IPA`},
	`RISC OS`:        {`ROS`},
	`GNU/Linux`: {
		`LIN`, `ARL`, `DEB`, `KNO`, `MIN`, `UBT`, `KBT`, `XBT`, `LBT`, `FED`,
		`RHT`, `VLN`, `MDR`, `GNT`, `SAB`, `SLW`, `SSE`, `CES`, `BTR`, `SAF`,
		`ORD`, `TOS`, `RSO`, `DEE`, `FRE`, `MAG`, `FEN`, `CAI`, `PCL`, `HAS`,
		`LOS`, `DVK`, `ROK`, `OWR`, `OTV`, `KTV`, `PUR`, `PLA`, `FUC`, `PAR`,
		`FOR`, `MON`, `KAN`, `ZEN`, `LND`, `LNS`, `CHN`, `AMZ`, `TEN`, `CST`,
		`NOV`, `ROU`, `ZOR`, `RED`, `KAL`, `ORA`, `VID`, `TIV`, `BSN`, `RAS`,
		`UOS`, `PIO`, `FRI`, `LIR`, `WEB`, `SER`, `ASP`, `AOS`, `LOO`, `EUL`,
		`SCI`, `ALP`, `CLO`, `ROC`, `OVZ`, `PVE`, `RST`, `EZX`, `GNS`, `JOL`,
		`TUR`, `QTP`, `WPO`, `PAN`, `VIZ`, `AZU`, `COL`,
	},
	`Mac`:                   {`MAC`},
	`Mobile Gaming Console`: {`PSP`, `NDS`, `XBX`},
	`OpenVMS`:               {`OVS`},
	`Real-time OS`:          {`MTK`, `TDX`, `MRE`, `JME`, `REX`, `RXT`, `KOL`},
	`Other Mobile`:          {`WOS`, `POS`, `SBA`, `TIZ`, `SMG`, `MAE`, `LUN`, `GEO`},
	`Symbian`:               {`SYM`, `SYS`, `SY3`, `S60`, `S40`},
	`Unix`: {
		`SOS`, `AIX`, `HPX`, `BSD`, `NBS`, `OBS`, `DFB`, `SYL`, `IRI`, `T64`,
		`INF`, `ELE`, `GNX`, `ULT`, `NWS`, `NXT`, `SBL`,
	},
	`WebTV`:          {`WTV`},
	`Windows`:        {`WIN`},
	`Windows Mobile`: {`WPH`, `WMO`, `WCE`, `WRT`, `WIO`, `KIN`},
	`Other Smart TV`: {`WHS`},
}

const (
	PlatformTypeARM  = "ARM"
	PlatformTypeX64  = "x64"
	PlatformTypeX86  = "x86"
	PlatformTypeNONE = ""
)

type PlatformReg struct {
	Name string
	Regular
}

type OsMatchResult struct {
	Name      string `yaml:"name" json:"name"`
	ShortName string `yaml:"short_name" json:"short_name"`
	Version   string `yaml:"version" json:"version"`
	Platform  string `yaml:"platform" json:"platform"`
}

type OsParser interface {
	PreMatch(string) bool
	Parse(string) *OsMatchResult
}

// Parses the useragent for operating system information
type Oss struct {
	Regexes      []*OsReg
	platforms    []*PlatformReg
	overAllMatch Regular
}

func NewOss(file string) (*Oss, error) {
	var v []*OsReg
	err := ReadYamlFile(file, &v)
	if err != nil {
		return nil, err
	}
	ps := []*PlatformReg{
		{Name: PlatformTypeARM, Regular: Regular{Regex: "arm"}},
		{Name: PlatformTypeX64, Regular: Regular{Regex: "WOW64|x64|win64|amd64|x86_64"}},
		{Name: PlatformTypeX86, Regular: Regular{Regex: "i[0-9]86|i86pc"}},
	}
	for _, pp := range ps {
		pp.Compile()
	}
	return &Oss{
		Regexes:   v,
		platforms: ps,
	}, nil
}

func (o *Oss) ParsePlatform(ua string) string {
	for _, p := range o.platforms {
		if p.IsMatchUserAgent(ua) {
			return p.Name
		}
	}
	return PlatformTypeNONE
}

func (o *Oss) PreMatch(ua string) bool {
	if o.overAllMatch.Regexp == nil {
		count := len(o.Regexes)
		if count == 0 {
			return false
		}
		sb := strings.Builder{}
		sb.WriteString(o.Regexes[count-1].Regex)
		for i := count - 2; i >= 0; i-- {
			sb.WriteString("|")
			sb.WriteString(o.Regexes[i].Regex)
		}
		o.overAllMatch.Regex = sb.String()
		o.overAllMatch.Compile()
	}
	r := o.overAllMatch.IsMatchUserAgent(ua)
	return r
}

func (o *Oss) Parse(ua string) *OsMatchResult {
	var matches []string
	var osRegex *OsReg
	for _, osRegex = range o.Regexes {
		matches = osRegex.MatchUserAgent(ua)
		if len(matches) > 0 {
			break
		}
	}

	if len(matches) == 0 || osRegex == nil {
		return nil
	}

	name := BuildByMatch(osRegex.Name, matches)
	short := UnknownShort

	for osShort, osName := range OperatingSystems {
		if StringEqualIgnoreCase(name, osName) {
			name = osName
			short = osShort
			break
		}
	}

	result := &OsMatchResult{
		Name:      name,
		ShortName: short,
		Version:   BuildVersion(osRegex.Version, matches),
		Platform:  o.ParsePlatform(ua),
	}
	return result
}

func GetOsFamily(osLabel string) string {
	for k, vs := range OsFamilies {
		for _, v := range vs {
			if v == osLabel {
				return k
			}
		}
	}
	return ""
}

func GetOsNameFromId(os, ver string) string {
	if osFullName, ok := OperatingSystems[os]; ok {
		return strings.TrimSpace(osFullName + " " + ver)
	}
	return ""
}
