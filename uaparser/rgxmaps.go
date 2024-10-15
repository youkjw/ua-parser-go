package uaparser

import (
	"github.com/dlclark/regexp2"
	"regexp"
	"strings"
)

var regMajor = regexp.MustCompile(`[^\d\.]`)

func majorize(version string) string {
	version = regMajor.ReplaceAllString(version, Empty)
	verArr := strings.Split(version, ".")
	if len(verArr) > 0 {
		return verArr[0]
	}
	return ""
}

var rgxMaps = map[UaItemType][]*Parser{
	UaBrowser: {
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(?:crmo|crios)\/([\w\.]+)`), // Chrome for Android/iOS
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, PREFIX_MOBILE+"Chrome"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)edg(?:e|ios|a)?\/([\w\.]+)`), // Microsoft Edge
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Edge"),
			},
		},

		// Presto based
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(opera mini)\/([-\w\.]+)`),                        // Opera Mini
				regexp.MustCompile(`(?i)(opera [mobiletab]{3,6})\b.+version\/([-\w\.]+)`), // Opera Mobi/Tablet
				regexp.MustCompile(`(?i)(opera)(?:.+version\/|[\/ ]+)([\w\.]+)`),          // Opera
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{ // Opera Mobi/Tablet
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)opios[\/ ]+([\w\.]+)`), // Opera mini on iphone >= 8.0
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, OPERA+" Mini"),
			},
		},
		&Parser{ // Opera
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bop(?:rg)?x\/([\w\.]+)`), // Opera GX
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, OPERA+" GX"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bopr\/([\w\.]+)`), // Opera Webkit
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, OPERA),
			},
		},

		// Mixed
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bb[ai]*d(?:uhd|[ub]*[aekoprswx]{5,6})[\/ ]?([\w\.]+)`), // Baidu
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Baidu"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(kindle)\/([\w\.]+)`),                                                 // Kindle
				regexp.MustCompile(`(?i)(lunascape|maxthon|netfront|jasmine|blazer|sleipnir)[\/ ]?([\w\.]*)`), // Lunascape/Maxthon/Netfront/Jasmine/Blazer/Sleipnir
				regexp.MustCompile(`(?i)(avant|iemobile|slim)\s?(?:browser)?[\/ ]?([\w\.]*)`),                 // Avant/IEMobile/SlimBrowser
				regexp.MustCompile(`(?i)(?:ms|\()(ie) ([\w\.]+)`),                                             // Internet Explorer
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(flock|rockmelt|midori|epiphany|silk|skyfire|ovibrowser|bolt|iron|vivaldi|iridium|phantomjs|bowser|qupzilla|falkon|rekonq|puffin|brave|whale(?!.+naver)|qqbrowserlite|duckduckgo|klar|helio)\/([-\w\.]+)`, regexp2.IgnoreCase), // Rekonq/Puffin/Brave/Whale/QQBrowserLite/QQ//Vivaldi/DuckDuckGo/Klar
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(heytap|ovi)browser\/([\d\.]+)`), // HeyTap/Ovi
				regexp.MustCompile(`(?i)(weibo)__([\d\.]+)`),             // Weibo
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)quark(?:pc)?\/([-\w\.]+)`), // quarkpc
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Quark"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bddg\/([\w\.]+)`), // DuckDuckGo
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "DuckDuckGo"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(?:\buc? ?browser|(?:juc.+)ucweb)[\/ ]?([\w\.]+)`), // UCBrowser
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "UCBrowser"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)microm.+\bqbcore\/([\w\.]+)`), // WeChat Desktop for Windows Built-in Browser
				regexp.MustCompile(`(?i)\bqbcore\/([\w\.]+).+microm`),
				regexp.MustCompile(`(?i)micromessenger\/([\w\.]+)`), // WeChat

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "WeChat"),
			},
			ParsedCall: []parsedCall{
				MakeMatchedCall(Version, windowsVersionMapper),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)konqueror\/([\w\.]+)`), // Konqueror

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Konqueror"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)trident.+rv[: ]([\w\.]{1,9})\b.+like gecko`), // IE11

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "IE"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)ya(?:search)?browser\/([\w\.]+)`), // Yandex

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Yandex"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)slbrowser\/([\w\.]+)`), // Smart Lenovo Browser

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Smart "+LENOVO+SUFFIX_BROWSER),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(avast|avg)\/([\w\.]+)`), // Avast/AVG Secure Browser

			},
			Props: []*Property{
				BuildProperty(Name, "(.+)", "${1} Secure"+SUFFIX_BROWSER),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bfocus\/([\w\.]+)`), // Firefox Focus

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, FIREFOX+" Focus"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bopt\/([\w\.]+)`), // Opera Touch

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, OPERA+" Touch"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)coc_coc\w+\/([\w\.]+)`), // Opera Touch

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Coc Coc"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)dolfin\/([\w\.]+)`), // Dolphin

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Dolphin"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)coast\/([\w\.]+)`), // Opera Coast

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, OPERA+" Coast"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)miuibrowser\/([\w\.]+)`), // MIUI Browser

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "MIUI"+SUFFIX_BROWSER),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)fxios\/([\w\.-]+)`), // Firefox for iOS

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, PREFIX_MOBILE+FIREFOX),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bqihu|(qi?ho?o?|360)browser`), // 360

			},
			Props: []*Property{
				BuildProperty(Name, "360"+SUFFIX_BROWSER),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(qq)\/([\w\.]+)`), // QQ

			},
			Props: []*Property{
				BuildProperty(Name, "(.+)", "${1}Browser"),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(oculus|sailfish|huawei|vivo|pico)browser\/([\w\.]+)`), // Oculus/Sailfish/HuaweiBrowser/VivoBrowser/PicoBrowser

			},
			Props: []*Property{
				BuildProperty(Name, "(.+)", "${1}"+SUFFIX_BROWSER),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)samsungbrowser\/([\w\.]+)`), // Samsung Internet

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, SAMSUNG+" Internet"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(comodo_dragon)\/([\w\.]+)`), // Comodo Dragon

			},
			Props: []*Property{
				BuildProperty(Name, "_", " "),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)metasr[\/ ]?([\d\.]+)`), // Sogou Explorer

			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, SOGOU+" Explorer"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(sogou)mo\w+\/([\d\.]+)`), // Sogou Mobile

			},
			Props: []*Property{
				BuildProperty(Name, SOGOU+" Mobile"),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(electron)\/([\w\.]+) safari`),                  // Electron-based App
				regexp.MustCompile(`(?i)(tesla)(?: qtcarbrowser|\/(20\d\d\.[-\w\.]+))`), // Tesla
				regexp.MustCompile(`(?i)m?(qqbrowser|2345Explorer)[\/ ]?([\w\.]+)`),     // QQBrowser/2345 Browser
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(lbbrowser|rekonq)`), // LieBao Browser/Rekonq
				regexp.MustCompile(`(?i)\[(linkedin)app\]`),  // LinkedIn App for iOS & Android
			},
			Props: []*Property{
				BuildProperty(Name),
			},
		},

		// WebView
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\[(linkedin)app\]`), // LinkedIn App for iOS & Android
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`((?:fban\/fbios|fb_iab\/fb4a)(?!.+fbav)|;fbav\/([\w\.]+);)`, regexp2.IgnoreCase), // Facebook App for iOS & Android
			},
			Props: []*Property{
				BuildProperty(Name, FACEBOOK),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(Klarna)\/([\w\.]+)`),                          // Klarna Shopping Browser for iOS & Android
				regexp.MustCompile(`(?i)(kakao(?:talk|story))[\/ ]([\w\.]+)`),          // Kakao App
				regexp.MustCompile(`(?i)(naver)\(.*?(\d+\.[\w\.]+).*\)`),               // Naver InApp
				regexp.MustCompile(`(?i)safari (line)\/([\w\.]+)`),                     // Line App for iOS
				regexp.MustCompile(`(?i)\b(line)\/([\w\.]+)\/iab`),                     // Line App for Android
				regexp.MustCompile(`(?i)(alipay)client\/([\w\.]+)`),                    // Alipay
				regexp.MustCompile(`(?i)(twitter)(?:and| f.+e\/([\w\.]+))`),            // Twitter
				regexp.MustCompile(`(?i)(chromium|instagram|snapchat)[\/ ]([-\w\.]+)`), // Chromium/Instagram/Snapchat
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bgsa\/([\w\.]+) .*safari\/`), // Klarna Shopping Browser for iOS & Android
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "GSA"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bgsa\/([\w\.]+) .*safari\/`), // Google Search Appliance on iOS
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "GSA"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)musical_ly(?:.+app_?version\/|_)([\w\.]+)`), // TikTok
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "TikTok"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)headlesschrome(?:\/([\w\.]+)| )`), // Chrome Headless
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, CHROME+" Headless"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i) wv\).+(chrome)\/([\w\.]+)`), // Chrome WebView
			},
			Props: []*Property{
				BuildProperty(Name, CHROME+" WebView"),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+ version\/([\w\.]+)\b.+(?:mobile safari|safari)`), // Android Browser
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Android"+SUFFIX_BROWSER),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)chrome\/([\w\.]+) mobile`), // Chrome Mobile
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, PREFIX_MOBILE+"Chrome"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(chrome|omniweb|arora|[tizenoka]{5} ?browser)\/v?([\w\.]+)`), // Chrome/OmniWeb/Arora/Tizen/Nokia
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)version\/([\w\.\,]+) .*mobile(?:\/\w+ | ?)safari`), // Safari Mobile
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, PREFIX_MOBILE+"Safari"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)iphone .*mobile(?:\/\w+ | ?)safari`), // Safari Mobile
			},
			Props: []*Property{
				BuildProperty(Name, PREFIX_MOBILE+"Safari"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)version\/([\w\.\,]+) .*(safari)`), // Safari
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)webkit.+?(mobile ?safari|safari)(\/[\w\.]+)`), // Safari < 3.0
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version, "1"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(webkit|khtml)\/([\w\.]+)`), // Safari < 3.0
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},

		// Gecko based
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(?:mobile|tablet);.*(firefox)\/([\w\.-]+)`), // Firefox Mobile
			},
			Props: []*Property{
				BuildProperty(Name, PREFIX_MOBILE+FIREFOX),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(navigator|netscape\d?)\/([-\w\.]+)`), // Netscape
			},
			Props: []*Property{
				BuildProperty(Name, "Netscape"),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(wolvic)\/([\w\.]+)`), // Wolvic
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)mobile vr; rv:([\w\.]+)\).+firefox`), // Firefox Reality
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, FIREFOX+" Reality"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)ekiohf.+(flow)\/([\w\.]+)`), // Flow
				regexp.MustCompile(`(?i)(swiftfox)`),                // Swiftfox
				regexp.MustCompile(`(?i)(icedragon|iceweasel|camino|chimera|fennec|maemo browser|minimo|conkeror)[\/ ]?([\w\.\+]+)`),  // IceDragon/Iceweasel/Camino/Chimera/Fennec/Maemo/Minimo/Conkeror
				regexp.MustCompile(`(?i)(seamonkey|k-meleon|icecat|iceape|firebird|phoenix|palemoon|basilisk|waterfox)\/([-\w\.]+)$`), // Firefox/SeaMonkey/K-Meleon/IceCat/IceApe/Firebird/Phoenix
				regexp.MustCompile(`(?i)(firefox)\/([\w\.]+)`),                                                                                        // Other Firefox-based
				regexp.MustCompile(`(?i)(mozilla)\/([\w\.]+) .+rv\:.+gecko\/\d+`),                                                                     // Mozilla
				regexp.MustCompile(`(?i)(polaris|lynx|dillo|icab|doris|amaya|w3m|netsurf|obigo|mosaic|(?:go|ice|up)[\. ]?browser)[-\/ ]?v?([\w\.]+)`), // Polaris/Lynx/Dillo/iCab/Doris/Amaya/w3m/NetSurf/Obigo/Mosaic/Go/ICE/UP.Browser
				regexp.MustCompile(`(?i)(links) \(([\w\.]+)`),                                                                                         // Links
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version, "_", "."),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(cobalt)\/([\w\.]+)`), // Cobalt
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version, "[^\\d\\.]+.", Empty),
			},
		},
	},
	UaCpu: {
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(?:(amd|x|x86[-_]?|wow|win)64)\b`), // AMD64 (x64)
			},
			Props: []*Property{
				BuildProperty(Architecture, "amd64"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{

				regexp.MustCompile(`(?i)((?:i[346]|x)86)[;\)]`), // IA32 (x86)
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(ia32(?=;))`, regexp2.IgnoreCase), // IA32 (quicktime)
			},
			Props: []*Property{
				BuildProperty(Architecture, "ia32"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(aarch64|arm(v?8e?l?|_?64))\b`), // ARM64
			},
			Props: []*Property{
				BuildProperty(Architecture, "armhf"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(arm(?:v[67])?ht?n?[fl]p?)\b`), // ARMHF
			},
			Props: []*Property{
				BuildProperty(Architecture, "arm64"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\bwindows (ce|mobile); ppc;`), // PocketPC mistakenly identified as PowerPC
			},
			Props: []*Property{
				BuildProperty(Architecture, "arm"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)((?:ppc|powerpc)(?:64)?)(?: mac|;|\))`), // PowerPC
			},
			Props: []*Property{
				BuildProperty(Architecture, "ower", Empty),
			},
			ParsedCall: []parsedCall{
				MakeMatchedCall(Architecture, strings.ToLower),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(sun4\w)[;\)]`), // SPARC
			},
			Props: []*Property{
				BuildProperty(Architecture, "sparc"),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`((?:avr32|ia64(?=;))|68k(?=\))|\barm(?=v(?:[1-7]|[5-7]1)l?|;|eabi)|(?=atmel )avr|(?:irix|mips|sparc)(?:64)?\b|pa-risc)`, regexp2.IgnoreCase), // IA64, 68K, ARM/64, AVR/32, IRIX/64, MIPS/64, SPARC/64, PA-RISC
			},
			Props: []*Property{
				BuildProperty(Architecture),
			},
			ParsedCall: []parsedCall{
				MakeMatchedCall(Architecture, strings.ToLower),
			},
		},
	},
	UaDevice: {

		//////////////////////////
		// MOBILES & TABLETS
		/////////////////////////

		// Samsung
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(sch-i[89]0\d|shw-m380s|sm-[ptx]\w{2,4}|gt-[pn]\d{2,4}|sgh-t8[56]9|nexus 10)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SAMSUNG),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`\b((?:s[cgp]h|gt|sm)-(?![lr])\w+|sc[g-]?[\d]+a?|galaxy nexus)`, regexp2.IgnoreCase),
				regexp2.MustCompile(`samsung[- ]((?!sm-[lr])[-\w]+)`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SAMSUNG),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)sec-(sgh\w+)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SAMSUNG),
				BuildProperty(Type, MOBILE),
			},
		},

		// Apple
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(?:\/|\()(ip(?:hone|od)[\w, ]*)(?:\/|;)`), // iPod/iPhone
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, APPLE),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\((ipad);[-\w\),; ]+apple`), // iPad
				regexp.MustCompile(`(?i)applecoremedia\/[\w\.]+ \((ipad)`),
				regexp.MustCompile(`(?i)\b(ipad)\d\d?,\d\d?[;\]].+ios`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, APPLE),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(macintosh);`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, APPLE),
			},
		},

		// Sharp
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(sh-?[altvz]?\d\d[a-ekm]?)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SHARP),
				BuildProperty(Type, MOBILE),
			},
		},

		// Huawei
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`\b((?:ag[rs][23]?|bah2?|sht?|btv)-a?[lw]\d{2})\b(?!.+d\/s)`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, HUAWEI),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(?:huawei|honor)([-\w ]+)[;\)]`),
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`\b(nexus 6p|\w{2,4}e?-[atu]?[ln][\dx][012359c][adn]?)\b(?!.+d\/s)`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, HUAWEI),
				BuildProperty(Type, MOBILE),
			},
		},

		// Xiaomi
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(poco[\w ]+|m2\d{3}j\d\d[a-z]{2})(?: bui|\))`),                                                           // Xiaomi POCO
				regexp.MustCompile(`(?i)\b(hm[-_ ]?note?[_ ]?(?:\d\w)?) bui`),                                                                      // Xiaomi Hongmi
				regexp.MustCompile(`(?i)\b(redmi[\-_ ]?(?:note|k)?[\w_ ]+)(?: bui|\))`),                                                            // Xiaomi Redmi
				regexp.MustCompile(`(?i)oid[^\)]+; (m?[12][0-389][01]\w{3,6}[c-y])( bui|; wv|\))`),                                                 // Xiaomi Redmi 'numeric' models
				regexp.MustCompile(`(?i)\b(mi[-_ ]?(?:a\d|one|one[_ ]plus|note lte|max|cc)?[_ ]?(?:\d?\w?)[_ ]?(?:plus|se|lite|pro)?)(?: bui|\))`), // Xiaomi Mi 				// Xiaomi Hongmi 'numeric' models
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`\b; (\w+) build\/hm\1`, regexp2.IgnoreCase), // Xiaomi Hongmi 'numeric' models
			},
			Props: []*Property{
				BuildProperty(Model, "_", " "),
				BuildProperty(Vendor, XIAOMI),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)oid[^\)]+; (2\d{4}(283|rpbf)[cgl])( bui|\))`), // Redmi Pad
				regexp.MustCompile(`(?i)\b(mi[-_ ]?(?:pad)(?:[\w_ ]+))(?: bui|\))`),   // Mi Pad tablets
			},
			Props: []*Property{
				BuildProperty(Model, "_", " "),
				BuildProperty(Vendor, XIAOMI),
				BuildProperty(Type, TABLET),
			},
		},

		// OPPO
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i); (\w+) bui.+ oppo`),
				regexp.MustCompile(`(?i)\b(cph[12]\d{3}|p(?:af|c[al]|d\w|e[ar])[mt]\d0|x9007|a101op)\b`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, OPPO),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(opd2\d{3}a?) bui`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, OPPO),
				BuildProperty(Type, TABLET),
			},
		},

		// Vivo
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)vivo (\w+)(?: bui|\))`),
				regexp.MustCompile(`(?i)\b(v[12]\d{3}\w?[at])(?: bui|;)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Vivo"),
				BuildProperty(Type, MOBILE),
			},
		},

		// Realme
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(rmx[1-3]\d{3})(?: bui|;|\))`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Realme"),
				BuildProperty(Type, MOBILE),
			},
		},

		// Motorola
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(milestone|droid(?:[2-4x]| (?:bionic|x2|pro|razr))?:?( 4g)?)\b[\w ]+build\/`),
				regexp.MustCompile(`(?i)\bmot(?:orola)?[- ](\w*)`),
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`((?:moto[\w\(\) ]+|xt\d{3,4}|nexus 6)(?= bui|\)))`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, MOTOROLA),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(mz60\d|xoom[2 ]{0,2}) build\/`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, MOTOROLA),
				BuildProperty(Type, TABLET),
			},
		},

		// LG
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`((?=lg)?[vl]k\-?\d{3}) bui| 3\.[-\w; ]{10}lg?-([06cv9]{3,4})`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, LG),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\blg-?([\d\w]+) bui`),
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(lm(?:-?f100[nv]?|-[\w\.]+)(?= bui|\))|nexus [45])`, regexp2.IgnoreCase),
				regexp2.MustCompile(`\blg[-e;\/ ]+((?!browser|netcast|android tv)\w+)`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, LG),
				BuildProperty(Type, MOBILE),
			},
		},

		// Lenovo
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(ideatab[-\w ]+)`),
				regexp.MustCompile(`(?i)lenovo ?(s[56]000[-\w]+|tab(?:[\w ]+)|yt[-\d\w]{6}|tb[-\d\w]{6})`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, LENOVO),
				BuildProperty(Type, TABLET),
			},
		},

		// Nokia
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(?:maemo|nokia).*(n900|lumia \d+)`),
				regexp.MustCompile(`(?i)nokia[-_ ]?([-\w\.]*)`),
			},
			Props: []*Property{
				BuildProperty(Model, "_", " "),
				BuildProperty(Vendor, NOKIA),
				BuildProperty(Type, MOBILE),
			},
		},

		// Google
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(pixel c)\b`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, GOOGLE),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; (pixel[\daxl ]{0,6})(?: bui|\))`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, GOOGLE),
				BuildProperty(Type, MOBILE),
			},
		},

		// Sony
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`droid.+ (a?\d[0-2]{2}so|[c-g]\d{4}|so[-gl]\w+|xq-a\w[4-7][12])(?= bui|\).+chrome\/(?![1-6]{0,1}\d\.))`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SONY),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)sony tablet [ps]`),
				regexp.MustCompile(`(?i)\b(?:sony)?sgp\w+(?: bui|\))`),
			},
			Props: []*Property{
				BuildProperty(Model, "Xperia Tablet"),
				BuildProperty(Vendor, SONY),
				BuildProperty(Type, TABLET),
			},
		},

		// OnePlus
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i) (kb2005|in20[12]5|be20[12][59])\b`),
				regexp.MustCompile(`(?i)(?:one)?(?:plus)? (a\d0\d\d)(?: b|\))`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "OnePlus"),
				BuildProperty(Type, MOBILE),
			},
		},

		// Amazon
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(kf[a-z]{2}wi|aeo(?!bc)\w\w)( bui|\))`, regexp2.IgnoreCase), // Kindle Fire without Silk / Echo Show
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, AMAZON),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(alexa)webm`),
				regexp.MustCompile(`(?i)(kf[a-z]+)( bui|\)).+silk\/`), // Kindle Fire HD
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, AMAZON),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)((?:sd|kf)[0349hijorstuw]+)( bui|\)).+silk\/`), // Fire Phone
			},
			Props: []*Property{
				BuildProperty(Model, "(.+)", "Fire Phone ${1}"),
				BuildProperty(Vendor, AMAZON),
				BuildProperty(Type, MOBILE),
			},
		},

		// BlackBerry
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(playbook);[-\w\),; ]+(rim)`), // BlackBerry PlayBook
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b((?:bb[a-f]|st[hv])100-\d)`),
				regexp.MustCompile(`(?i)\(bb10; (\w+)`), // BlackBerry 10
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, BLACKBERRY),
				BuildProperty(Type, MOBILE),
			},
		},

		// Asus
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(?:\b|asus_)(transfo[prime ]{4,10} \w+|eeepc|slider \w+|nexus 7|padfone|p00[cj])`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, ASUS),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i) (z[bes]6[027][012][km][ls]|zenfone \d\w?)\b`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, ASUS),
				BuildProperty(Type, MOBILE),
			},
		},

		// HTC
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(nexus 9)`), // HTC Nexus 9
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, HTC),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(zte)[- ]([\w ]+?)(?: bui|\/|\))`), // ZTE
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(htc)[-;_ ]{1,2}([\w ]+(?=\)| bui)|\w+)`, regexp2.IgnoreCase),                                      // HTC
				regexp2.MustCompile(`(alcatel|geeksphone|nexian|panasonic(?!(?:;|\.))|sony(?!-bra))[-_ ]?([-\w]*)`, regexp2.IgnoreCase), // Alcatel/GeeksPhone/Nexian/Panasonic/Sony
			},
			Props: []*Property{
				BuildProperty(Vendor, HTC),
				BuildProperty(Model, "_", " "),
				BuildProperty(Type, MOBILE),
			},
		},

		// tcl
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)tcl (xess p17aa)`),
				regexp.MustCompile(`(?i)droid [\w\.]+; ((?:8[14]9[16]|9(?:0(?:48|60|8[01])|1(?:3[27]|66)|2(?:6[69]|9[56])|466))[gqswx])(_\w(\w|\w\w))?(\)| bui)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, TCL),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid [\w\.]+; (418(?:7d|8v)|5087z|5102l|61(?:02[dh]|25[adfh]|27[ai]|56[dh]|59k|65[ah])|a509dl|t(?:43(?:0w|1[adepqu])|50(?:6d|7[adju])|6(?:09dl|10k|12b|71[efho]|76[hjk])|7(?:66[ahju]|67[hw]|7[045][bh]|71[hk]|73o|76[ho]|79w|81[hks]?|82h|90[bhsy]|99b)|810[hs]))(_\w(\w|\w\w))?(\)| bui)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, TCL),
				BuildProperty(Type, MOBILE),
			},
		},

		// itel
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(itel) ((\w+))`),
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type),
			},
			ParsedCall: []parsedCall{
				MakeMatchedCall(Vendor, strings.ToLower),
				MakeMatchedCall(Type, itelTypeMapper),
			},
		},

		// Acer
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; ([ab][1-7]-?[0178a]\d\d?)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, ACER),
				BuildProperty(Type, TABLET),
			},
		},

		// Meizu
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; (m[1-5] note) bui`),
				regexp.MustCompile(`(?i)\bmz-([-\w]{2,})`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, MEIZU),
				BuildProperty(Type, MOBILE),
			},
		},

		// Ulefone
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i); ((?:power )?armor(?:[\w ]{0,8}))(?: bui|\))`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Ulefone"),
				BuildProperty(Type, MOBILE),
			},
		},

		// Nothing
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; (a(?:015|06[35]|142p?))`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Nothing"),
				BuildProperty(Type, MOBILE),
			},
		},

		// MIXED
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(blackberry|benq|palm(?=\-)|sonyericsson|acer|asus|dell|meizu|motorola|polytron|infinix|tecno)[-_ ]?([-\w]*)`, regexp2.IgnoreCase), // BlackBerry/BenQ/Palm/Sony-Ericsson/Acer/Asus/Dell/Meizu/Motorola/Polytron
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(hp) ([\w ]+\w)`),            // HP iPAQ
				regexp.MustCompile(`(?i)(asus)-?(\w+)`),              // Asus
				regexp.MustCompile(`(?i)(microsoft); (lumia[\w ]+)`), // Microsoft Lumia
				regexp.MustCompile(`(?i)(lenovo)[-_ ]?([-\w]+)`),     // Lenovo
				regexp.MustCompile(`(?i)(jolla)`),                    // Jolla
				regexp.MustCompile(`(?i)(oppo) ?([\w ]+) bui`),       // OPPO
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(kobo)\s(ereader|touch)`), // Kobo
				regexp.MustCompile(`(?i)(archos) (gamepad2?)`),    // Archos
				regexp.MustCompile(`(?i)(kindle)\/([\w\.]+)`),     // Kindle
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(hp).+(touchpad(?!.+tablet)|tablet)`, regexp2.IgnoreCase), // HP TouchPad
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(surface duo)`), // Surface Duo
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, MICROSOFT),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid [\d\.]+; (fp\du?)(?: b|\))`), // Fairphone
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Fairphone"),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(shield[\w ]+) b`), // Nvidia Shield Tablets
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Nvidia"),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(sprint) (\w+)`), // Sprint Phones
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(kin\.[onetw]{3})`), // Microsoft Kin
			},
			Props: []*Property{
				BuildProperty(Model, "\\.", " "),
				BuildProperty(Vendor, MICROSOFT),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; ([c6]+|et5[16]|mc[239][23]x?|vc8[03]x?)\)`), // Zebra
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, ZEBRA),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; (ec30|ps20|tc[2-8]\d[kx])\)`),
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, ZEBRA),
				BuildProperty(Type, MOBILE),
			},
		},

		///////////////////
		// SMARTTVS
		///////////////////

		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)smart-tv.+(samsung)`), // Samsung
			},
			Props: []*Property{
				BuildProperty(Vendor, ZEBRA),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)hbbtv.+maple;(\d+)`),
			},
			Props: []*Property{
				BuildProperty(Model, "^", "SmartTV"),
				BuildProperty(Vendor, SAMSUNG),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(nux; netcast.+smarttv|lg (netcast\.tv-201\d|android tv))`), // LG SmartTV
			},
			Props: []*Property{
				BuildProperty(Vendor, LG),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(apple) ?tv`), // Apple TV
			},
			Props: []*Property{
				BuildProperty(Vendor, LG),
				BuildProperty(Model, APPLE+" TV"),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(apple) ?tv`), // Apple TV
			},
			Props: []*Property{
				BuildProperty(Vendor, LG),
				BuildProperty(Model, APPLE+" TV"),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)crkey.*devicetype\/chromecast`), // Google Chromecast Third Generation
			},
			Props: []*Property{
				BuildProperty(Model, CHROMECAST+" Third Generation"),
				BuildProperty(Vendor, GOOGLE),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)crkey.*devicetype\/([^/]*)`), // Google Chromecast with specific device type
			},
			Props: []*Property{
				BuildProperty(Model, "^", "Chromecast "),
				BuildProperty(Vendor, GOOGLE),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)fuchsia.*crkey`), // Google Chromecast Nest Hub
			},
			Props: []*Property{
				BuildProperty(Model, CHROMECAST+" Nest Hub"),
				BuildProperty(Vendor, GOOGLE),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)crkey`), // Google Chromecast, Linux-based or unknown
			},
			Props: []*Property{
				BuildProperty(Model, CHROMECAST),
				BuildProperty(Vendor, GOOGLE),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+aft(\w+)( bui|\))`), // Fire TV
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, AMAZON),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\(dtv[\);].+(aquos)`),
				regexp.MustCompile(`(?i)(aquos-tv[\w ]+)\)`), // Sharp
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SHARP),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(bravia[\w ]+)( bui|\))`), // Sony
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SONY),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(mitv-\w{5}) bui`), // Xiaomi
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, XIAOMI),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)Hbbtv.*(technisat) (.*);`), // TechniSAT
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, SMARTTV),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(roku)[\dx]*[\)\/]((?:dvp-)?[\d\.]*)`),                   // Roku
				regexp.MustCompile(`(?i)hbbtv\/\d+\.\d+\.\d+ +\([\w\+ ]*; *([\w\d][^;]*);([^;]*)`), // HbbTV devices
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, SMARTTV),
			},
			ParsedCall: []parsedCall{
				MakeMatchedCall(Vendor, trim),
				MakeMatchedCall(Model, trim),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(android tv|smart[- ]?tv|opera tv|tv; rv:)\b`), // SmartTV from Unidentified Vendors
			},
			Props: []*Property{
				BuildProperty(Type, SMARTTV),
			},
		},

		///////////////////
		// CONSOLES
		///////////////////
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(ouya)`),       // Ouya
				regexp.MustCompile(`(nintendo) (\w+)`), // Nintendo
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, CONSOLE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; (shield) bui`), // Nvidia
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Nvidia"),
				BuildProperty(Type, CONSOLE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(playstation \w+)`), // Playstation
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SONY),
				BuildProperty(Type, CONSOLE),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`\b(xbox(?: one)?(?!; xbox))[\); ]`, regexp2.IgnoreCase), // Microsoft Xbox
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, MICROSOFT),
				BuildProperty(Type, CONSOLE),
			},
		},

		///////////////////
		// WEARABLES
		///////////////////

		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\b(sm-[lr]\d\d[05][fnuw]?s?)`), // Samsung Galaxy Watch
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, SAMSUNG),
				BuildProperty(Type, WEARABLE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)((pebble))app`), // Pebble
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, WEARABLE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(watch)(?: ?os[,\/]|\d,\d\/)[\d\.]+`), // Apple Watch
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, APPLE),
				BuildProperty(Type, WEARABLE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; (wt63?0{2,3})\)`), // Apple Watch
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, ZEBRA),
				BuildProperty(Type, WEARABLE),
			},
		},

		///////////////////
		// XR
		///////////////////
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid.+; (glass) \d`), // Google Glass
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, GOOGLE),
				BuildProperty(Type, XR),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(pico) (4|neo3(?: link|pro)?)`), // Pico
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Model),
				BuildProperty(Type, XR),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(quest( \d| pro)?)`), // Oculus Quest
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, FACEBOOK),
				BuildProperty(Type, XR),
			},
		},

		///////////////////
		// EMBEDDED
		///////////////////

		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(tesla)(?: qtcarbrowser|\/[-\w\.]+)`), // Tesla
			},
			Props: []*Property{
				BuildProperty(Vendor),
				BuildProperty(Type, EMBEDDED),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(aeobc)\b`), // Echo Dot
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, AMAZON),
				BuildProperty(Type, EMBEDDED),
			},
		},

		////////////////////
		// MIXED (GENERIC)
		///////////////////

		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid .+?; ([^;]+?)(?: bui|; wv\)|\) applew).+? mobile safari`), // Android Phones from Unidentified Vendors
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`droid .+?; ([^;]+?)(?: bui|\) applew).+?(?! mobile) safari`, regexp2.IgnoreCase), // Android Tablets from Unidentified Vendors
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`\b((tablet|tab)[;\/]|focus\/\d(?!.+mobile))`, regexp2.IgnoreCase), // Unidentifiable Tablet
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Type, TABLET),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(phone|mobile(?:[;\/]| [ \w\/\.]*safari)|pda(?=.+windows ce))`, regexp2.IgnoreCase), // Unidentifiable Mobile
			},
			Props: []*Property{
				BuildProperty(Type, MOBILE),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(android[-\w\. ]{0,9});.+buil`), // Generic Android Device
			},
			Props: []*Property{
				BuildProperty(Model),
				BuildProperty(Vendor, "Generic"),
			},
		},
	},
	UaEngine: {
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)windows.+ edge\/([\w\.]+)`), // EdgeHTML
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, EDGE+"HTML"),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`webkit\/537\.36.+chrome\/(?!27)([\w\.]+)`, regexp2.IgnoreCase), // Blink
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Blink"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(presto)\/([\w\.]+)`),                                                // Presto
				regexp.MustCompile(`(?i)(webkit|trident|netfront|netsurf|amaya|lynx|w3m|goanna)\/([\w\.]+)`), // WebKit/Trident/NetFront/NetSurf/Amaya/Lynx/w3m/Goanna
				regexp.MustCompile(`(?i)ekioh(flow)\/([\w\.]+)`),                                             // Flow
				regexp.MustCompile(`(?i)(khtml|tasman|links)[\/ ]\(?([\w\.]+)`),                              // KHTML/Tasman/Links
				regexp.MustCompile(`(?i)(icab)[\/ ]([23]\.[\d\.]+)`),                                         // iCab
				regexp.MustCompile(`(?i)\b(libweb)`),
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)rv\:([\w\.]{1,9})\b.+(gecko)`), // Gecko
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name),
			},
		},
	},
	UaOs: {
		// Windows
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)microsoft (windows) (vista|xp)`), // Windows (iTunes)
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(windows (?:phone(?: os)?|mobile))[\/ ]?([\d\.\w ]*)`), // Windows Phone
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
			ParsedCall: []parsedCall{
				MakeMatchedCall(Version, windowsVersionMapper),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)windows nt 6\.2; (arm)`), // Windows RT
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`windows[\/ ]?([ntce\d\. ]+\w)(?!.+xbox)`, regexp2.IgnoreCase),
				regexp2.MustCompile(`(?:win(?=3|9|n)|win 9x )([nt\d\.]+)`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, WINDOWS),
			},
			ParsedCall: []parsedCall{
				MakeMatchedCall(Version, windowsVersionMapper),
			},
		},

		// iOS/macOS
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)ip[honead]{2,4}\b(?:.*os ([\w]+) like mac|; opera)`), // iOS
				regexp.MustCompile(`(?i)(?:ios;fbsv\/|iphone.+ios[\/ ])([\d\.]+)`),
				regexp.MustCompile(`(?i)cfnetwork\/.+darwin`),
			},
			Props: []*Property{
				BuildProperty(Version, "_", "."),
				BuildProperty(Name, "iOS"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(mac os x) ?([\w\. ]*)`), // Mac OS
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(macintosh|mac_powerpc\b)(?!.+haiku)`, regexp2.IgnoreCase),
			},
			Props: []*Property{
				BuildProperty(Name, "macOS"),
				BuildProperty(Version, "_", "."),
			},
		},

		// Google Chromecast
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)android ([\d\.]+).*crkey`), // Google Chromecast, Android-based
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, CHROMECAST+" Android"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)fuchsia.*crkey\/([\d\.]+)`), // Google Chromecast, Fuchsia-based
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, CHROMECAST+" Fuchsia"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)crkey\/([\d\.]+).*devicetype\/smartspeaker`), // Google Chromecast, Linux-based Smart Speaker
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, CHROMECAST+" SmartSpeaker"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)linux.*crkey\/([\d\.]+)`), // Google Chromecast, Legacy Linux-based
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, CHROMECAST+" Linux"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)crkey\/([\d\.]+)`), // Google Chromecast, Legacy Linux-based
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, CHROMECAST),
			},
		},

		// Mobile OSes
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid ([\w\.]+)\b.+(android[- ]x86|harmonyos)`), // Android-x86/HarmonyOS
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)droid ([\w\.]+)\b.+(android[- ]x86|harmonyos)`), // Android-x86/HarmonyOS
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(android|webos|qnx|bada|rim tablet os|maemo|meego|sailfish)[-\/ ]?([\w\.]*)`), // Android/WebOS/QNX/Bada/RIM/Maemo/MeeGo/Sailfish OS
				regexp.MustCompile(`(?i)(blackberry)\w*\/([\w\.]*)`),                                                  // Blackberry
				regexp.MustCompile(`(?i)(tizen|kaios)[\/ ]([\w\.]+)`),                                                 // Tizen/KaiOS
				regexp.MustCompile(`(?i)\((series40);`),                                                               // Series 40
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)\(bb(10);`), // BlackBerry 10
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, BLACKBERRY),
			},
		},
		&Parser{
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(?:symbian ?os|symbos|s60(?=;)|series60)[-\/ ]?([\w\.]*)`, regexp2.IgnoreCase), // Symbian
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "Symbian"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)mozilla\/[\d\.]+ \((?:mobile|tablet|tv|mobile; [\w ]+); rv:.+ gecko\/([\w\.]+)`), // Firefox OS
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, FIREFOX+" OS"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)web0s;.+rt(tv)`), // WebOS
				regexp.MustCompile(`(?i)\b(?:hp)?wos(?:browser)?\/([\w\.]+)`),
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "webOS"),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)watch(?: ?os[,\/]|\d,\d\/)([\d\.]+)`), // watchOS
			},
			Props: []*Property{
				BuildProperty(Version),
				BuildProperty(Name, "watchOS"),
			},
		},

		// Google ChromeOS
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(cros) [\w]+(?:\)| ([\w\.]+)\b)`), // Chromium OS
			},
			Props: []*Property{
				BuildProperty(Name, "Chrome OS"),
				BuildProperty(Version),
			},
		},

		&Parser{
			Regexps: []*regexp.Regexp{
				// Smart TVs
				regexp.MustCompile(`(?i)panasonic;(viera)`),       // Panasonic Viera
				regexp.MustCompile(`(?i)(netrange)mmh`),           // Netrange
				regexp.MustCompile(`(?i)(nettv)\/(\d+\.[\w\.]+)`), // NetTV
				// Console
				regexp.MustCompile(`(?i)(nintendo|playstation) (\w+)`), // Nintendo/Playstation
				regexp.MustCompile(`(?i)(xbox); +xbox ([^\);]+)`),      // Microsoft Xbox (360, One, X, S, Series X, Series S)
				regexp.MustCompile(`(?i)(pico) .+os([\w\.]+)`),         // Pico
				// Other
				regexp.MustCompile(`(?i)\b(joli|palm)\b ?(?:os)?\/?([\w\.]*)`), // Joli/Palm
				regexp.MustCompile(`(?i)(mint)[\/\(\) ]?(\w*)`),                // Mint
				regexp.MustCompile(`(?i)(mageia|vectorlinux)[; ]`),             // Mageia/VectorLinux
				regexp.MustCompile(`(?i)(hurd|linux) ?([\w\.]*)`),              // Hurd/Linux
				regexp.MustCompile(`(?i)(gnu) ?([\w\.]*)`),                     // GNU
				regexp.MustCompile(`(?i)(haiku) (\w+)`),                        // Haiku

			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`([kxln]?ubuntu|debian|suse|opensuse|gentoo|arch(?= linux)|slackware|fedora|mandriva|centos|pclinuxos|red ?hat|zenwalk|linpus|raspbian|plan 9|minix|risc os|contiki|deepin|manjaro|elementary os|sabayon|linspire)(?: gnu\/linux)?(?: enterprise)?(?:[- ]linux)?(?:-gnu)?[-\/ ]?(?!chrom|package)([-\w\.]*)`, regexp2.IgnoreCase), // Ubuntu/Debian/SUSE/Gentoo/Arch/Slackware/Fedora/Mandriva/CentOS/PCLinuxOS/RedHat/Zenwalk/Linpus/Raspbian/Plan9/Minix/RISCOS/Contiki/Deepin/Manjaro/elementary/Sabayon/Linspire
				regexp2.MustCompile(`\b([-frentopcghs]{0,5}bsd|dragonfly)[\/ ]?(?!amd|[ix346]{1,2}86)([\w\.]*)`, regexp2.IgnoreCase), // FreeBSD/NetBSD/OpenBSD/PC-BSD/GhostBSD/DragonFly
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)(sunos) ?([\w\.\d]*)`), // Solaris
			},
			Props: []*Property{
				BuildProperty(Name, "Solaris"),
				BuildProperty(Version),
			},
		},
		&Parser{
			Regexps: []*regexp.Regexp{
				regexp.MustCompile(`(?i)((?:open)?solaris)[-\/ ]?([\w\.]*)`),                              // Solaris
				regexp.MustCompile(`(?i)\b(beos|os\/2|amigaos|morphos|openvms|fuchsia|hp-ux|serenityos)`), // BeOS/OS2/AmigaOS/MorphOS/OpenVMS/Fuchsia/HP-UX/SerenityOS
				regexp.MustCompile(`(?i)(unix) ?([\w\.]*)`),                                               // UNIX
			},
			ComplexRegexps: []*regexp2.Regexp{
				regexp2.MustCompile(`(aix) ((\d)(?=\.|\)| )[\w\.])*`, regexp2.IgnoreCase), // AIX
			},
			Props: []*Property{
				BuildProperty(Name),
				BuildProperty(Version),
			},
		},
	},
}
