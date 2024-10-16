package uaparser

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	TestBrowser(t)
	TestCpu(t)
	TestEngine(t)
	TestDevice(t)
	TestOs(t)
	TestCli(t)
	TestCrawlers(t)
	TestFetchers(t)
}

var testTable = []string{
	// wechat
	//"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63030522)",
	// Mac
	//"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
	// iPhone
	//"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
	// Android
	//"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
	//"Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.99 Mobile Safari/537.36",
	// FreeBSD
	//"Mozilla/5.0 (compatible; Konqueror/4.5; FreeBSD) KHTML/4.5.4 (like Gecko)",
	// Windows phone
	//"Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; NOKIA; Lumia 630)",
	//OviBrowser
	//"Mozilla/5.0 (Series40; NokiaX3-02/le6.32; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/1.0.0.11.8",
	// Avast Secure Browser
	//"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36 Avast/72.0.1174.122",
	// Avant Browser
	//"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; GTB5; Avant Browser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	//
	//"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
	// CriOS
	//"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3",
	// Opera
	//"Opera/9.80 (X11; Linux x86_64; U; Linux Mint; en) Presto/2.2.15 Version/10.10",
	// Safari
	//"Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  Mobile/15E148 Safari/604.1",
	// quarkpc
	//"mozilla/5.0 (windows nt 10.0; win64; x64) applewebkit/537.36 (khtml, like gecko) chrome/112.0.0.0 safari/537.36 quarkpc/1.5.5.75",
	//
	"Mozilla/5.0 (Linux; Android 11; 21061119AG Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.131 Mobile Safari/537.36 trill_2022109040 JsSdk/1.0 NetType/MOBILE Channel/googleplay AppName/musical_ly app_version/21.9.4 ByteLocale/ru-RU ByteFullLocale/ru-RU Region/KG BytedanceWebview/d8a21c6",
}

func TestParse(t *testing.T) {
	//wait := sync.WaitGroup{}
	//u := NewUserAgent(testTable[0])
	//for i := 0; i <= 100; i++ {
	//	wait.Add(1)
	//	go func() {
	//		fmt.Println(u.GetUa())
	//		wait.Done()
	//	}()
	//}
	//wait.Wait()

	for _, testUa := range testTable {
		info := Parse(testUa, SetExtensions([]string{CRAWLER}))
		fmt.Println(info)
	}

}

var testBrowser = []byte(`
[
    {
        "desc"    : "360 Browser on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_1 like Mac OS X) AppleWebKit/607.3.9 (KHTML, like Gecko) Mobile/16G102 QHBrowser/317 QihooBrowser/4.0.10",
        "expect"  :
        {
            "name"    : "360 Browser",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Alipay",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 10; zh-CN; V2034A Build/QP1A.190711.020) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 UWS/3.22.2.33 Mobile Safari/537.36 UCBS/3.22.2.33_211025173018 NebulaSDK/1.8.100112 Nebula AlipayDefined(nt:WIFI,ws:360|0|2.0) AliApp(AP/10.2.51.7100) AlipayClient/10.2.51.7100 Language/zh-Hans useStatusBar/true isConcaveScreen/true Region/CNAriver/1.0.0",
        "expect"  :
        {
            "name"    : "Alipay",
            "version" : "10.2.51.7100",
            "major"   : "10"
        }
    },
    {
        "desc"    : "Alipay",
        "ua"      : "Mozilla/5.0 (Linux; Android 10; VOG-AL00 Build/HUAWEIVOG-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/105.0.5195.148 MYWeb/0.2.103.0_20230131112530 UWS/3.22.2.9999 UCBS/3.22.2.9999_220000000000 Mobile Safari/537.36 NebulaSDK/1.8.100112 Nebula AlipayDefined(nt:WIFI,ws:360|0|3.0) AliApp(AP/10.3.50.9999) AlipayClient/10.3.50.9999 Language/en isConcaveScreen/true Region/CN ProductType/devAriver/1.0.0",
        "expect"  :
        {
            "name"    : "Alipay",
            "version" : "10.3.50.9999",
            "major"   : "10"
        }
    },
    {
        "desc"    : "Android Browser on Galaxy Nexus",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 4.0.2; en-us; Galaxy Nexus Build/ICL53F) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "expect"  :
        {
            "name"    : "Android Browser",
            "version" : "4.0",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Android Browser on Galaxy S3",
        "ua"      : "Mozilla/5.0 (Linux; Android 4.4.4; en-us; SAMSUNG GT-I9300I Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/1.5 Chrome/28.0.1500.94 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Android Browser",
            "version" : "1.5",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Android Browser on HTC Flyer (P510E)",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 3.2.1; ru-ru; HTC Flyer P510e Build/HTK75C) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
        "expect"  :
        {
            "name"    : "Android Browser",
            "version" : "4.0",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Android Browser on Huawei Honor Glory II (U9508)",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 4.0.4; ru-by; HUAWEI U9508 Build/HuaweiU9508) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 ACHEETAHI/2100050044",
        "expect"  :
        {
            "name"    : "Android Browser",
            "version" : "4.0",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Android Browser on Huawei P8 (H891L)",
        "ua"      : "Mozilla/5.0 (Linux; Android 4.4.4; HUAWEI H891L Build/HuaweiH891L) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Android Browser",
            "version" : "4.0",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Android Browser on Samsung S6 (SM-G925F)",
        "ua"      : "Mozilla/5.0 (Linux; Android 5.0.2; SAMSUNG SM-G925F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/3.0 Chrome/38.0.2125.102 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Samsung Internet",
            "version" : "3.0",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Sailfish Browser",
        "ua"      : "Mozilla/5.0 (Linux; U; Sailfish 3.0; Mobile; rv:45.0) Gecko/45.0 Firefox/45.0 SailfishBrowser/1.0",
        "expect"  :
        {
            "name"    : "Sailfish Browser",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Arora",
        "ua"      : "Mozilla/5.0 (Windows; U; Windows NT 5.1; de-CH) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.2",
        "expect"  :
        {
            "name"    : "Arora",
            "version" : "0.2",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Avant",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; GTB5; Avant Browser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
        "expect"  :
        {
            "name"    : "Avant",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Avast Secure Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36 Avast/72.0.1174.122",
        "expect"  :
        {
            "name"    : "Avast Secure Browser",
            "version" : "72.0.1174.122",
            "major"   : "72"
        }
    },
    {
        "desc"    : "AVG Secure Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36 AVG/72.0.719.123",
        "expect"  :
        {
            "name"    : "AVG Secure Browser",
            "version" : "72.0.719.123",
            "major"   : "72"
        }
    },
    {
        "desc"    : "Baidu",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; baidubrowser 1.x)",
        "expect"  :
        {
            "name"    : "Baidu",
            "version" : "1.x",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Baidu",
        "ua"      : "Mozilla/5.0 (Linux; Android 9; Redmi Note 5 Build/PKQ1.180904.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/110.0.5481.153 Mobile Safari/537.36 bdbrowser/6.4.0.4",
        "expect"  :
        {
            "name"    : "Baidu",
            "version" : "6.4.0.4",
            "major"   : "6"
        }
    },
    {
        "desc"    : "Baidu",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.4.9999.1900 Safari/537.31 BDSpark/26.4",
        "expect"  :
        {
            "name"    : "Baidu",
            "version" : "26.4",
            "major"   : "26"
        }
    },
    {
        "desc"    : "Baidu",
        "ua"      : "Mozilla/5.0 (iPad; CPU OS 9_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) BaiduHD/5.4.0.0 Mobile/10A406 Safari/8536.25",
        "expect"  :
        {
            "name"    : "Baidu",
            "version" : "5.4.0.0",
            "major"   : "5"
        }
    },
    {
        "desc"    : "BaiDu Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 BIDUBrowser/8.7 Safari/537.36",
        "expect"  :
        {
            "name"    : "Baidu",
            "version" : "8.7",
            "major"   : "8"
        }
    },
    {
        "desc"    : "baidu app on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16C101 main%2F1.0 baiduboxapp/11.12.0.18 (Baidu; P2 12.1.2)",
        "expect"  :
        {
            "name"    : "Baidu",
            "version" : "11.12.0.18",
            "major"   : "11"
        }
    },
    {
        "desc"    : "baidu app on Android",
        "ua"      : "Mozilla/5.0 (Linux; Android 8.1.0; BKK-AL10 Build/HONORBKK-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/63.0.3239.83 Mobile Safari/537.36 T7/11.11 baiduboxapp/11.11.0.0 (Baidu; P1 8.1.0)",
        "expect"  :
        {
            "name"    : "Baidu",
            "version" : "11.11.0.0",
            "major"   : "11"
        }
    },
    {
        "desc"    : "Blazer",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/hspr-H102; Blazer/4.0) 16;320x320",
        "expect"  :
        {
            "name"    : "Blazer",
            "version" : "4.0",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Bolt",
        "ua"      : "Mozilla/5.0 (X11; 78; CentOS; US-en) AppleWebKit/527+ (KHTML, like Gecko) Bolt/0.862 Version/3.0 Safari/523.15",
        "expect"  :
        {
            "name"    : "Bolt",
            "version" : "0.862",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Bowser",
        "ua"      : "Mozilla/5.0 (iOS; like Mac OS X) AppleWebKit/536.36 (KHTML, like Gecko) not Chrome/27.0.1500.95 Mobile/10B141 Safari/537.36 Bowser/0.2.1",
        "expect"  :
        {
            "name"    : "Bowser",
            "version" : "0.2.1",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Camino",
        "ua"      : "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.4; en; rv:1.9.0.19) Gecko/2011091218 Camino/2.0.9 (like Firefox/3.0.19)",
        "expect"  :
        {
            "name"    : "Camino",
            "version" : "2.0.9",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Camino on Mac",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:2.0.1) Gecko/20100101 Firefox/4.0.1 Camino/2.2.1",
        "expect"  :
        {
            "name"    : "Camino",
            "version" : "2.2.1",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Chimera",
        "ua"      : "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; pl-PL; rv:1.0.1) Gecko/20021111 Chimera/0.6",
        "expect"  :
        {
            "name"    : "Chimera",
            "version" : "0.6",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Chrome",
        "ua"      : "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
        "expect"  :
        {
            "name"    : "Chrome",
            "version" : "20.0.1090.0",
            "major"   : "20"
        }
    },
    {
        "desc"    : "Chrome",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4758.102 Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome",
            "version" : "100.0.4758.102",
            "major"   : "100"
        }
    },
    {
        "desc"    : "Chrome 112.0 on Win10",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome",
            "version" : "112.0.0.0",
            "major"   : "112"
        }
    },
    {
        "desc"    : "Chrome 112.0 on macOS",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome",
            "version" : "112.0.0.0",
            "major"   : "112"
        }
    },
    {
        "desc"    : "Chrome 111.0 on Linux",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome",
            "version" : "111.0.0.0",
            "major"   : "111"
        }
    },
    {
        "desc"    : "Chrome 111.0 on ChromeOS",
        "ua"      : "Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome",
            "version" : "111.0.0.0",
            "major"   : "111"
        }
    },
    {
            "desc"    : "Chrome Headless",
            "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome Safari/537.36",
            "expect"  :
        {
           "name"    : "Chrome Headless",
           "version" : "",
           "major"   : ""
        }
    },
    {
            "desc"    : "Chrome Headless",
            "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/60.0.3112.113 Safari/537.36",
            "expect"  :
        {
           "name"    : "Chrome Headless",
           "version" : "60.0.3112.113",
           "major"   : "60"
        }
    },
    {
        "desc"    : "Chrome WebView",
        "ua"      : "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome WebView",
            "version" : "43.0.2357.65",
            "major"   : "43"
        }
    },
    {
        "desc"    : "Chrome on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3",
        "expect"  :
        {
            "name"    : "Mobile Chrome",
            "version" : "19.0.1084.60",
            "major"   : "19"
        }
    },
    {
        "desc"    : "Chromium",
        "ua"      : "Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Ubuntu/11.10 Chromium/16.0.912.21 Chrome/16.0.912.21 Safari/535.7",
        "expect"  :
        {
            "name"    : "Chromium",
            "version" : "16.0.912.21",
            "major"   : "16"
        }
    },
    {
        "desc"    : "Chrome on Android",
        "ua"      : "Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7",
        "expect"  :
        {
            "name"    : "Mobile Chrome",
            "version" : "16.0.912.75",
            "major"   : "16"
        }
    },
    {
        "desc"    : "Coc Coc Browser",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) coc_coc_browser/78.0.129 Chrome/72.0.3626.129 Safari/537.36",
        "expect"  :
        {
            "name"    : "Coc Coc",
            "version" : "78.0.129",
            "major"   : "78"
        }
    },
    {
        "desc"    : "Comodo Dragon",
        "ua"      : "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/535.7 (KHTML, like Gecko) Comodo_Dragon/16.1.1.0 Chrome/16.0.912.63 Safari/535.7",
        "expect"  :
        {
            "name"    : "Comodo Dragon",
            "version" : "16.1.1.0",
            "major"   : "16"
        }
    },
    {
        "desc"    : "Conkeror",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; rv:6.0.1) Gecko/20110831 conkeror/0.9.3",
        "expect"  :
        {
            "name"    : "conkeror",
            "version" : "0.9.3",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Dillo",
        "ua"      : "Dillo/2.2",
        "expect"  :
        {
            "name"    : "Dillo",
            "version" : "2.2",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Dolphin",
        "ua"      : "Mozilla/5.0 (SCH-F859/F859DG12;U;NUCLEUS/2.1;Profile/MIDP-2.1 Configuration/CLDC-1.1;480*800;CTC/2.0) Dolfin/2.0",
        "expect"  :
        {
            "name"    : "Dolphin",
            "version" : "2.0",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Doris",
        "ua"      : "Doris/1.15 [en] (Symbian)",
        "expect"  :
        {
            "name"    : "Doris",
            "version" : "1.15",
            "major"   : "1"
        }
    },
    {
        "desc"    : "DuckDuckGo",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Safari/605.1.1517.4.1 Ddg/17.4.1",
        "expect"  :
        {
            "name"    : "DuckDuckGo",
            "version" : "17.4.1",
            "major"   : "17"
        }
    },
    {
        "desc"    : "DuckDuckGo",
        "ua"      : "Mozilla/5.0 (Linux; Android 8.1.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.131 Mobile DuckDuckGo/5 Safari/537.36",
        "expect"  :
        {
            "name"    : "DuckDuckGo",
            "version" : "5",
            "major"   : "5"
        }
    },
    {
        "desc"    : "Epiphany",
        "ua"      : "Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.7) Gecko/20040628 Epiphany/1.2.6",
        "expect"  :
        {
            "name"    : "Epiphany",
            "version" : "1.2.6",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Flow",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) EkiohFlow/5.7.4.30559 Flow/5.7.4 (like Gecko Firefox/53.0 rv:53.0)",
        "expect"  :
        {
            "name"    : "Flow",
            "version" : "5.7.4",
            "major"   : "5"
        }
    },
    {
        "desc"    : "Go Browser",
        "ua"      : "NokiaE66/GoBrowser/2.0.297",
        "expect"  :
        {
            "name"    : "GoBrowser",
            "version" : "2.0.297",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Waterfox",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; rv:55.0) Gecko/20100101 Firefox/55.2.2 Waterfox/55.2.2",
        "expect"  :
        {
            "name"    : "Waterfox",
            "version" : "55.2.2",
            "major"   : "55"
        }
    },
    {
        "desc"    : "PaleMoon",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; rv:52.9) Gecko/20100101 Goanna/3.4 Firefox/52.9 PaleMoon/27.6.1",
        "expect"  :
        {
            "name"    : "PaleMoon",
            "version" : "27.6.1",
            "major"   : "27"
        }
    },
    {
        "desc"    : "Basilisk",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; rv:55.0) Gecko/20100101 Goanna/4.0 Firefox/55.0 Basilisk/20171113",
        "expect"  :
        {
            "name"    : "Basilisk",
            "version" : "20171113",
            "major"   : "20171113"
        }
    },
    {
        "desc"    : "Facebook in-App Browser for Android with version",
        "ua"      : "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.121 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/35.0.0.48.273;]",
        "expect"  :
        {
            "name"    : "Facebook",
            "version" : "35.0.0.48.273",
            "major"   : "35"
        }
    },
    {
        "desc"    : "Facebook in-App Browser for iOS with version",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Mobile/14E304 [FBAN/FBIOS;FBAV/91.0.0.41.73;FBBV/57050710;FBDV/iPhone8,1;FBMD/iPhone;FBSN/iOS;FBSV/10.3.1;FBSS/2;FBCR/Telekom.de;FBID/phone;FBLC/de_DE;FBOP/5;FBRV/0])",
        "expect"  :
        {
            "name"    : "Facebook",
            "version" : "91.0.0.41.73",
            "major"   : "91"
        }
    },
    {
        "desc"    : "Facebook in-App Browser for iOS without version",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone10,2;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
        "expect"  :
        {
            "name"    : "Facebook",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Klarna in-App Browser for iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Klarna/23.36.223",
        "expect"  :
        {
            "name"    : "Klarna",
            "version" : "23.36.223",
            "major"   : "23"
        }
    },
    {
        "desc"    : "Klarna in-App Browser for Android",
        "ua"      : "Mozilla/5.0 (Linux; Android 12; moto g(60)s Build/S3RLS32.114-25-13; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/116.0.0.0 Mobile Safari/537.36 Klarna/23.36.215",
        "expect"  :
        {
            "name"    : "Klarna",
            "version" : "23.36.215",
            "major"   : "23"
        }
    },
    {
        "desc"    : "Instagram in-App Browser for iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 14_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Instagram 142.0.0.22.109 (iPhone12,5; iOS 14_1; en_US; en-US; scale=3.00; 1242x2688; 214888322) NW/1",
        "expect"  :
        {
            "name"    : "Instagram",
            "version" : "142.0.0.22.109",
            "major"   : "142"
        }
    },
    {
        "desc"    : "Falkon",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Falkon/3.0.0 Chrome/61.0.3163.140 Safari/537.36",
        "expect"  :
        {
            "name"    : "Falkon",
            "version" : "3.0.0",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Firebird",
        "ua"      : "Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.5) Gecko/20031007 Firebird/0.7",
        "expect"  :
        {
            "name"    : "Firebird",
            "version" : "0.7",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Firefox",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; rv:15.0) Gecko/20120716 Firefox/15.0a2",
        "expect"  :
        {
            "name"    : "Firefox",
            "version" : "15.0a2",
            "major"   : "15"
        }
    },
    {
        "desc"    : "Firefox",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:100.0) Gecko/20100101 Firefox/100.0",
        "expect"  :
        {
            "name"    : "Firefox",
            "version" : "100.0",
            "major"   : "100"
        }
    },
    {
        "desc"    : "Firefox Reality",
        "ua"      : "Mozilla/5.0 (Android 7.1.2; Mobile VR; rv:65.0) Gecko/65.0 Firefox/65.0",
        "expect"  :
        {
            "name"    : "Firefox Reality",
            "version" : "65.0",
            "major"   : "65"
        }
    },
    {
        "desc"    : "Firefox-based browser",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; rv:80.0) Gecko/20100101 Firefox/80.0 AppName/1.0",
        "expect"  :
        {
            "name"    : "Firefox",
            "version" : "80.0",
            "major"   : "80"
        }
    },
    {
        "desc"    : "Fennec",
        "ua"      : "Mozilla/5.0 (X11; U; Linux armv61; en-US; rv:1.9.1b2pre) Gecko/20081015 Fennec/1.0a1",
        "expect"  :
        {
            "name"    : "Fennec",
            "version" : "1.0a1",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Firefox for Maemo (Nokia N900)",
        "ua"      : "Mozilla/5.0 (Maemo; Linux armv7l; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 Fennec/10.0.1",
        "expect"  :
        {
            "name"    : "Fennec",
            "version" : "10.0.1",
            "major"   : "10"
        }
    },
    {
        "desc": "Firefox Focus",
        "ua": "Mozilla/5.0 (Linux; Android 7.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Focus/6.1.1 Chrome/68.0.3440.91 Mobile Safari/537.36",
        "expect": {
            "name": "Firefox Focus",
            "version": "6.1.1",
            "major": "6"
        }
    },
    {
        "desc"    : "Flock",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.3) Gecko/2008100716 Firefox/3.0.3 Flock/2.0",
        "expect"  :
        {
            "name"    : "Flock",
            "version" : "2.0",
            "major"   : "2"
        }
    },
    {
        "desc"    : "GoBrowser",
        "ua"      : "Nokia5700XpressMusic/GoBrowser/1.6.91",
        "expect"  :
        {
            "name"    : "GoBrowser",
            "version" : "1.6.91",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Helio",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36 Helio/0.98.20",
        "expect"  :
        {
            "name"    : "Helio",
            "version" : "0.98.20",
            "major"   : "0"
        }
    },
    {
        "desc"    : "HeyTap",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.61 Safari/537.36 HeyTapBrowser/40.8.10.1",
        "expect"  :
        {
            "name"    : "HeyTap",
            "version" : "40.8.10.1",
            "major"   : "40"
        }
    },
    {
        "desc"    : "HuaweiBrowser",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00；HMSCore/4.0.0 GMS/10.4 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Huawei Browser",
            "version" : "10.0.3.102",
            "major"   : "10"
        }
    },
    {
        "desc"    : "HuaweiBrowser",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00；HMSCore/4.0.0 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Huawei Browser",
            "version" : "10.0.3.102",
            "major"   : "10"
        }
    },
    {
        "desc"    : "HuaweiBrowser",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00；GMS/10.4 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Huawei Browser",
            "version" : "10.0.3.102",
            "major"   : "10"
        }
    },
    {
        "desc"    : "HuaweiBrowser",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Huawei Browser",
            "version" : "10.0.3.102",
            "major"   : "10"
        }
    },
    {
        "desc"    : "IceApe",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1.19) Gecko/20110817 Iceape/2.0.14",
        "expect"  :
        {
            "name"    : "Iceape",
            "version" : "2.0.14",
            "major"   : "2"
        }
    },
    {
        "desc"    : "ICEBrowser",
        "ua"      : "Mozilla/5.0 (Java 1.6.0_01; Windows XP 5.1 x86; en) ICEbrowser/v6_1_2",
        "expect"  :
        {
            "name"    : "ICEbrowser",
            "version" : "6.1.2",
            "major"   : "6"
        }
    },
    {
        "desc"    : "IceCat",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.3) Gecko/2008092921 IceCat/3.0.3-g1",
        "expect"  :
        {
            "name"    : "IceCat",
            "version" : "3.0.3-g1",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Iceweasel",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; de; rv:1.9.0.16) Gecko/2009121610 Iceweasel/3.0.6 (Debian-3.0.6-3)",
        "expect"  :
        {
            "name"    : "Iceweasel",
            "version" : "3.0.6",
            "major"   : "3"
        }
    },
    {
        "desc"    : "iCab",
        "ua"      : "iCab/2.9.5 (Macintosh; U; PPC; Mac OS X)",
        "expect"  :
        {
            "name"    : "iCab",
            "version" : "2.9.5",
            "major"   : "2"
        }
    },
    {
        "desc"    : "IEMobile",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.11) 320x240; VZW; Motorola-Q9c; Windows Mobile 6.1 Standard",
        "expect"  :
        {
            "name"    : "IEMobile",
            "version" : "7.11",
            "major"   : "7"
        }
    },
    {
        "desc"    : "IE 11 with IE token",
        "ua"      : "Mozilla/5.0 (IE 11.0; Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko",
        "expect"  :
        {
            "name"    : "IE",
            "version" : "11.0",
            "major"   : "11"
        }
    },
    {
        "desc"    : "IE 11 without IE token",
        "ua"      : "Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv 11.0) like Gecko",
        "expect"  :
        {
            "name"    : "IE",
            "version" : "11.0",
            "major"   : "11"
        }
    },
    {
        "desc"    : "Iron",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.4 (KHTML, like Gecko) Chrome/22.0.1250.0 Iron/22.0.2150.0 Safari/537.4",
        "expect"  :
        {
            "name"    : "Iron",
            "version" : "22.0.2150.0",
            "major"   : "22"
        }
    },
    {
        "desc"    : "Jasmine",
        "ua"      : "SAMSUNG-S8000/S8000XXIF3 SHP/VPP/R5 Jasmine/1.0 Nextreaming SMM-MMS/1.2.0 profile/MIDP-2.1 configuration/CLDC-1.1",
        "expect"  :
        {
            "name"    : "Jasmine",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "K-Meleon",
        "ua"      : "Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.5) Gecko/20031016 K-Meleon/0.8.2",
        "expect"  :
        {
            "name"    : "K-Meleon",
            "version" : "0.8.2",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Kindle Browser",
        "ua"      : "Mozilla/4.0 (compatible; Linux 2.6.22) NetFront/3.4 Kindle/2.5 (screen 600x800; rotate)",
        "expect"  :
        {
            "name"    : "Kindle",
            "version" : "2.5",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Klar < 4.1",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Klar/1.0 Chrome/58.0.3029.83 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Klar",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Konqueror",
        "ua"      : "Mozilla/5.0 (compatible; Konqueror/3.5; Linux; X11; x86_64) KHTML/3.5.6 (like Gecko) (Kubuntu)",
        "expect"  :
        {
            "name"    : "Konqueror",
            "version" : "3.5",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Konqueror",
        "ua"      : "Mozilla/5.0 (X11; Linux i686) AppleWebKit/534.34 (KHTML, like Gecko) konqueror/5.0.97 Safari/534.34",
        "expect"  :
        {
            "name"    : "Konqueror",
            "version" : "5.0.97",
            "major"   : "5"
        }
    },
    {
        "desc"    : "PicoBrowser",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; Pico Neo3 Link OS5.8.4.0 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.22 Chrome/105.0.5195.68 VR Safari/537.36",
        "expect"  :
        {
            "name"    : "Pico Browser",
            "version" : "3.3.22",
            "major"   : "3"
        }
    },
    {
        "desc"    : "PicoBrowser",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; PICO 4 OS5.8.2 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.38 Chrome/105.0.5195.68 VR Safari/537.36",
        "expect"  :
        {
            "name"    : "Pico Browser",
            "version" : "3.3.38",
            "major"   : "3"
        }
    },
    {
        "desc"    : "PicoBrowser",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; PICO 4 OS5.4.0 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.22 Chrome/105.0.5195.68 VR Safari/537.36 OculusBrowser/7.0",
        "expect"  :
        {
            "name"    : "Pico Browser",
            "version" : "3.3.22",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Rekonq",
        "ua"      : "Mozilla/5.0 (X11; U; Linux x86_64; cs-CZ) AppleWebKit/533.3 (KHTML, like Gecko) rekonq Safari/533.3",
        "expect"  :
        {
            "name"    : "rekonq",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Smart Lenovo Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 SLBrowser/8.0.0.10171 SLBChan/8",
        "expect"  :
        {
            "name"    : "Smart Lenovo Browser",
            "version" : "8.0.0.10171",
            "major"   : "8"
        }
    },
    {
        "desc"    : "Smart Lenovo Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 SLBrowser/9.0.0.9011 SLBChan/10",
        "expect"  :
        {
            "name"    : "Smart Lenovo Browser",
            "version" : "9.0.0.9011",
            "major"   : "9"
        }
    },
    {
        "desc"    : "LINE on Android",
        "ua"      : "Mozilla/5.0 (Linux; Android 5.0; ASUS_Z00AD Build/LRX21V; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 Line/6.5.1/IAB",
        "expect"  :
        {
            "name"    : "Line",
            "version" : "6.5.1",
            "major"   : "6"
        }
    },
    {
        "desc"    : "LINE on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_6 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) Mobile/15D100 Safari Line/8.4.1",
        "expect"  :
        {
            "name"    : "Line",
            "version" : "8.4.1",
            "major"   : "8"
        }
    },
    {
        "desc"    : "Lunascape",
        "ua"      : "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090804 Firefox/3.5.2 Lunascape/5.1.4.5",
        "expect"  :
        {
            "name"    : "Lunascape",
            "version" : "5.1.4.5",
            "major"   : "5"
        }
    },
    {
        "desc"    : "Lynx",
        "ua"      : "Lynx/2.8.5dev.16 libwww-FM/2.14 SSL-MM/1.4.1 OpenSSL/0.9.6b",
        "expect"  :
        {
            "name"    : "Lynx",
            "version" : "2.8.5dev.16",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Maemo Browser",
        "ua"      : "Mozilla/5.0 (X11; U; Linux armv7l; ru-RU; rv:1.9.2.3pre) Gecko/20100723 Firefox/3.5 Maemo Browser 1.7.4.8 RX-51 N900",
        "expect"  :
        {
            "name"    : "Maemo Browser",
            "version" : "1.7.4.8",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Maxthon",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; Maxthon; .NET CLR 1.1.4322)",
        "expect"  :
        {
            "name"    : "Maxthon",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Midori",
        "ua"      : "Midori/0.2.2 (X11; Linux i686; U; en-us) WebKit/531.2+",
        "expect"  :
        {
            "name"    : "Midori",
            "version" : "0.2.2",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Minimo",
        "ua"      : "Mozilla/5.0 (X11; U; Linux armv6l; rv 1.8.1.5pre) Gecko/20070619 Minimo/0.020",
        "expect"  :
        {
            "name"    : "Minimo",
            "version" : "0.020",
            "major"   : "0"
        }
    },
    {
        "desc"    : "MIUI Browser on Xiaomi Hongmi WCDMA (HM2013023)",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 4.2.2; ru-ru; 2013023 Build/HM2013023) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 XiaoMi/MiuiBrowser/1.0",
        "expect"  :
        {
            "name"    : "MIUI Browser",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Mobile Safari",
        "ua"      : "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7",
        "expect"  :
        {
            "name"    : "Mobile Safari",
            "version" : "4.0.5",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Mobile Safari",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)  Mobile/15E148 Safari/604.1",
        "expect"  :
        {
            "name"    : "Mobile Safari",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Mosaic",
        "ua"      : "NCSA_Mosaic/2.6 (X11; SunOS 4.1.3 sun4m)",
        "expect"  :
        {
            "name"    : "Mosaic",
            "version" : "2.6",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Mozilla",
        "ua"      : "Mozilla/5.0 (X11; U; SunOS sun4u; en-US; rv:1.7) Gecko/20070606",
        "expect"  :
        {
            "name"    : "Mozilla",
            "version" : "5.0",
            "major"   : "5"
        }
    },
    {
        "desc"    : "MSIE",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 5.0b1; Mac_PowerPC)",
        "expect"  :
        {
            "name"    : "IE",
            "version" : "5.0b1",
            "major"   : "5"
        }
    },
    {
        "desc"    : "NetFront",
        "ua"      : "Mozilla/4.0 (PDA; Windows CE/1.0.1) NetFront/3.0",
        "expect"  :
        {
            "name"    : "NetFront",
            "version" : "3.0",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Netscape on Windows ME",
        "ua"      : "Mozilla/5.0 (Windows; U; Win 9x 4.90; en-US; rv:1.8.1.8pre) Gecko/20071015 Firefox/2.0.0.7 Navigator/9.0",
        "expect"  :
        {
            "name"    : "Netscape",
            "version" : "9.0",
            "major"   : "9"
        }
    },
    {
        "desc"    : "Netscape on Windows 2000",
        "ua"      : "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.7.5) Gecko/20050519 Netscape/8.0.1",
        "expect"  :
        {
            "name"    : "Netscape",
            "version" : "8.0.1",
            "major"   : "8"
        }
    },
    {
        "desc"    : "Netscape 6",
        "ua"      : "Mozilla/5.0 (Windows; U; Win95; de-DE; rv:0.9.2) Gecko/20010726 Netscape6/6.1",
        "expect"  :
        {
            "name"    : "Netscape",
            "version" : "6.1",
            "major"   : "6"
        }
    },
    {
        "desc"    : "NetSurf in Plan9",
        "ua"      : "Mozilla/5.0 (Plan9) NetSurf/3.12",
        "expect"  :
        {
            "name"    : "NetSurf",
            "version" : "3.12",
            "major"   : "3"
        }
    },
    {
        "desc"    : "NetSurf in Linux",
        "ua"      : "NetSurf/3.10 (Linux; Arch Linux)",
        "expect"  :
        {
            "name"    : "NetSurf",
            "version" : "3.10",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Nokia Browser",
        "ua"      : "Mozilla/5.0 (Symbian/3; Series60/5.2 NokiaN8-00/025.007; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.37 Mobile Safari/533.4 3gpp-gba",
        "expect"  :
        {
            "name"    : "NokiaBrowser",
            "version" : "7.3.1.37",
            "major"   : "7"
        }
    },
    {
        "desc"    : "Obigo",
        "ua"      : "LG-GS290/V100 Obigo/WAP2.0 Profile/MIDP-2.1 Configuration/CLDC-1.1",
        "expect"  :
        {
            "name"    : "Obigo",
            "version" : "WAP2.0",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Obigo",
        "ua"      : "LG/KU990i/v10a Browser/Obigo-Q05A/3.6 MMS/LG-MMS-V1.0/1.2 Java/ASVM/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1",
        "expect"  :
        {
            "name"    : "Obigo",
            "version" : "Q05A",
            "major"   : "05"
        }
    },
    {
        "desc"    : "Oculus Browser",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.0; SM-G920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/3.4.9 SamsungBrowser/4.0 Chrome/57.0.2987.146 Mobile VR Safari/537.36",
        "expect"  :
        {
            "name"    : "Oculus Browser",
            "version" : "3.4.9",
            "major"   : "3"
        }
    },
    {
        "desc"    : "Oculus Browser",
        "ua"      : "Mozilla/5.0 (Linux; Android 10; Quest 2) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/15.0.0.0.22.280317669 SamsungBrowser/4.0 Chrome/89.0.4389.90 VR Safari/537.36",
        "expect"  :
        {
            "name"    : "Oculus Browser",
            "version" : "15.0.0.0.22.280317669",
            "major"   : "15"
        }
    },
    {
        "desc"    : "OmniWeb",
        "ua"      : "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-US) AppleWebKit/85 (KHTML, like Gecko) OmniWeb/v558.48",
        "expect"  :
        {
            "name"    : "OmniWeb",
            "version" : "558.48",
            "major"   : "558"
        }
    },
    {
        "desc"    : "Opera > 9.80",
        "ua"      : "Opera/9.80 (X11; Linux x86_64; U; Linux Mint; en) Presto/2.2.15 Version/10.10",
        "expect"  :
        {
            "name"    : "Opera",
            "version" : "10.10",
            "major"   : "10"
        }
    },
    {
        "desc"    : "Opera < 9.80 on Windows",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 5.0; Windows 95) Opera 6.01 [en]",
        "expect"  :
        {
            "name"    : "Opera",
            "version" : "6.01",
            "major"   : "6"
        }
    },
    {
        "desc"    : "Opera < 9.80 on OSX",
        "ua"      : "Opera/8.5 (Macintosh; PPC Mac OS X; U; en)",
        "expect"  :
        {
            "name"    : "Opera",
            "version" : "8.5",
            "major"   : "8"
        }
    },
    {
        "desc"    : "Opera Mobile",
        "ua"      : "Opera/9.80 (Android 2.3.5; Linux; Opera Mobi/ADR-1111101157; U; de) Presto/2.9.201 Version/11.50",
        "expect"  :
        {
            "name"    : "Opera Mobi",
            "version" : "11.50",
            "major"   : "11"
        }
    },
    {
        "desc"    : "Opera Webkit",
        "ua"      : "Mozilla/5.0 AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.123 Mobile Safari/537.22 OPR/14.0.1025.52315",
        "expect"  :
        {
            "name"    : "Opera",
            "version" : "14.0.1025.52315",
            "major"   : "14"
        }
    },
    {
        "desc"    : "Opera Mini",
        "ua"      : "Opera/9.80 (J2ME/MIDP; Opera Mini/5.1.21214/19.916; U; en) Presto/2.5.25",
        "expect"  :
        {
            "name"    : "Opera Mini",
            "version" : "5.1.21214",
            "major"   : "5"
        }
    },
    {
        "desc"    : "Opera Mini 8 above on iPhone",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 9_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) OPiOS/12.1.1.98980 Mobile/13C75 Safari/9537.53",
        "expect"  :
        {
            "name"    : "Opera Mini",
            "version" : "12.1.1.98980",
            "major"   : "12"
        }
    },
    {
        "desc"    : "Opera GX on Android",
        "ua"      : "Mozilla/5.0 (Linux; Android 10; Redmi Note 8 Pro Build/QP1A.190711.020) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.5790.168 Mobile Safari/537.36 OPX/2",
        "expect"  :
        {
            "name"    : "Opera GX",
            "version" : "2",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Opera GX on Windows",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36 OPR/60.0.3255.50747 OPRGX/60.0.3255.50747",
        "expect"  :
        {
            "name"    : "Opera GX",
            "version" : "60.0.3255.50747",
            "major"   : "60"
        }
    },
    {
        "desc"    : "Opera Tablet",
        "ua"      : "Opera/9.80 (Windows NT 6.1; Opera Tablet/15165; U; en) Presto/2.8.149 Version/11.1",
        "expect"  :
        {
            "name"    : "Opera Tablet",
            "version" : "11.1",
            "major"   : "11"
        }
    },
    {
        "desc"    : "Opera Coast",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 9_3_2 like Mac OS X; en) AppleWebKit/601.1.46 (KHTML, like Gecko) Coast/5.04.110603 Mobile/13F69 Safari/7534.48.3",
        "expect"  :
        {
            "name"    : "Opera Coast",
            "version" : "5.04.110603",
            "major"   : "5"
        }
    },
    {
        "desc"    : "Opera Touch",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.0; Lenovo P2a42 Build/NRD90N) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 OPT/1.10.33",
        "expect"  :
        {
            "name"    : "Opera Touch",
            "version" : "1.10.33",
            "major"   : "1"
        }
    },
    {
        "desc"    : "OviBrowser",
        "ua"      : "Mozilla/5.0 (Series40; NokiaX3-02/le6.32; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/1.0.0.11.8",
        "expect"  :
        {
            "name"    : "OviBrowser",
            "version" : "1.0.0.11.8",
            "major"   : "1"
        }
    },
    {
        "desc"    : "PhantomJS",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.34 (KHTML, like Gecko) PhantomJS/1.9.2 Safari/534.34",
        "expect"  :
        {
            "name"    : "PhantomJS",
            "version" : "1.9.2",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Phoenix",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.2b) Gecko/20021029 Phoenix/0.4",
        "expect"  :
        {
            "name"    : "Phoenix",
            "version" : "0.4",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Polaris",
        "ua"      : "LG-LX600 Polaris/6.0 MMP/2.0 Profile/MIDP-2.1 Configuration/CLDC-1.1",
        "expect"  :
        {
            "name"    : "Polaris",
            "version" : "6.0",
            "major"   : "6"
        }
    },
    {
        "desc"    : "QQBrowser",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; OPPO R7s Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 Chrome/37.0.0.0 MQQBrowser/7.1 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "QQBrowser",
            "version" : "7.1",
            "major"   : "7"
        }
    },
    {
        "desc"    : "QQBrowser",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 9; zh-cn; vivo X21 Build/PKQ1.180819.001) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/9.9 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "QQBrowser",
            "version" : "9.9",
            "major"   : "9"
        }
    },
    {
        "desc"    : "Quark",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 12; zh-Hans-CN; JLH-AN00 Build/HONORJLH-AN00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Quark/5.8.2.221 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Quark",
            "version" : "5.8.2.221",
            "major"   : "5"
        }
    },
    {
        "desc"    : "Quark",
        "ua"      : "mozilla/5.0 (windows nt 10.0; win64; x64) applewebkit/537.36 (khtml, like gecko) chrome/112.0.0.0 safari/537.36 quarkpc/1.5.5.75",
        "expect"  :
        {
            "name"    : "Quark",
            "version" : "1.5.5.75",
            "major"   : "1"
        }
    },
    {
        "desc"    : "QupZilla",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) QupZilla/1.8.9 Safari/538.1",
        "expect"  :
        {
            "name"    : "QupZilla",
            "version" : "1.8.9",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Rekonq 2",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.21 (KHTML, like Gecko) rekonq/2.2.1 Safari/537.21",
        "expect"  :
        {
            "name"    : "rekonq",
            "version" : "2.2.1",
            "major"   : "2"
        }
    },
    {
        "desc"    : "RockMelt",
        "ua"      : "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.7 (KHTML, like Gecko) RockMelt/0.8.36.78 Chrome/7.0.517.44 Safari/534.7",
        "expect"  :
        {
            "name"    : "RockMelt",
            "version" : "0.8.36.78",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Safari",
        "ua"      : "Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/533.17.8 (KHTML, like Gecko) Version/5.0.1 Safari/533.17.8",
        "expect"  :
        {
            "name"    : "Safari",
            "version" : "5.0.1",
            "major"   : "5"
        }
    },
    {
        "desc"    : "Safari < 3.0",
        "ua"      : "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; sv-se) AppleWebKit/419 (KHTML, like Gecko) Safari/419.3",
        "expect"  :
        {
            "name"    : "Safari",
            "version" : "1",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Samsung Internet for Android",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-G925A Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Samsung Internet",
            "version" : "4.0",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Samsung Internet for Tizen Mobile",
        "ua"      : "Mozilla/5.0 (Linux; Tizen 2.3; SAMSUNG SM-Z130H) AppleWebKit/537.3 (KHTML, like Gecko) SamsungBrowser/1.0 Mobile Safari/537.3",
        "expect"  :
        {
            "name"    : "Samsung Internet",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Samsung Internet for Smart-TV",
        "ua"      : "Mozilla/5.0 (SMART-TV; Linux; Tizen 2.3) AppleWebkit/538.1 (KHTML, like Gecko) SamsungBrowser/1.0 TV Safari/538.1",
        "expect"  :
        {
            "name"    : "Samsung Internet",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Samsung Internet for Gear VR",
        "ua"      : "Mozilla/5.0 (Linux; Android 5.0.2; SAMSUNG SM-G925K Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile VR Safari/537.36",
        "expect"  :
        {
            "name"    : "Samsung Internet",
            "version" : "4.0",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Samsung Internet in Redmi 8A",
        "ua"      : "Mozilla/5.0 (Linux; Android 10; Redmi 8A) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/23.0 Chrome/115.0.0.0 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Samsung Internet",
            "version" : "23.0",
            "major"   : "23"
        }
    },
    {
        "desc"    : "SeaMonkey",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1b4pre) Gecko/20090405 SeaMonkey/2.0b1pre",
        "expect"  :
        {
            "name"    : "SeaMonkey",
            "version" : "2.0b1pre",
            "major"   : "2"
        }
    },
    {
        "desc"    : "SeaMonkey on Mac",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 SeaMonkey/2.7.1",
        "expect"  :
        {
            "name"    : "SeaMonkey",
            "version" : "2.7.1",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Silk Browser",
        "ua"      : "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-84)",
        "expect"  :
        {
            "name"    : "Silk",
            "version" : "1.1.0-84",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Skyfire",
        "ua"      : "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_7; en-us) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Safari/530.17 Skyfire/2.0",
        "expect"  :
        {
            "name"    : "Skyfire",
            "version" : "2.0",
            "major"   : "2"
        }
    },    
    {
        "desc"    : "Sleipnir",
        "ua"      : "Mozilla/5.0 (Linux; Android 10; SOV37 Build/52.1.C.0.220; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/123.0.6312.120 Mobile Safari/537.36 Sleipnir/3.7.5",
        "expect"  :
        {
            "name"    : "Sleipnir",
            "version" : "3.7.5",
            "major"   : "3"
        }
    },

    {
        "desc"    : "Sleipnir",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; Sleipnir 2.8.4)",
        "expect"  :
        {
            "name"    : "Sleipnir",
            "version" : "2.8.4",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Sleipnir",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.04506.648; .NET CLR 3.5.21022) Sleipnir/2.8.4",
        "expect"  :
        {
            "name"    : "Sleipnir",
            "version" : "2.8.4",
            "major"   : "2"
        }
    },
    {
        "desc"    : "SlimBrowser",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SlimBrowser)",
        "expect"  :
        {
            "name"    : "Slim",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Swiftfox",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1) Gecko/20061024 Firefox/2.0 (Swiftfox)",
        "expect"  :
        {
            "name"    : "Swiftfox",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Tesla",
        "ua"      : "Mozilla/5.0 (X11; GNU/Linux) AppleWebKit/601.1 (KHTML, like Gecko) Tesla QtCarBrowser Safari/601.1",
        "expect"  :
        {
            "name"    : "Tesla",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Tesla",
        "ua"      : "Mozilla/5.0 (X11; GNU/Linux) AppleWebKit/537.36 (KHTML, like Gecko) Chromium/79.0.3945.130 Chrome/79.0.3945.130 Safari/537.36 Tesla/2020.16.2.1-e99c70fff409",
        "expect"  :
        {
            "name"    : "Tesla",
            "version" : "2020.16.2.1-e99c70fff409",
            "major"   : "2020"
        }
    },
    {
        "desc"    : "Tizen Browser",
        "ua"      : "Mozilla/5.0 (Linux; U; Tizen/1.0 like Android; en-us; AppleWebKit/534.46 (KHTML, like Gecko) Tizen Browser/1.0 Mobile",
        "expect"  :
        {
            "name"    : "Tizen Browser",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "UC Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 UBrowser/5.6.12860.7 Safari/537.36",
        "expect"  :
        {
            "name"    : "UCBrowser",
            "version" : "5.6.12860.7",
            "major"   : "5"
        }
    },
    {
        "desc"    : "UC Browser",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 6.0.1; en-US; Lenovo P2a42 Build/MMB29M) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 UCBrowser/11.2.0.915 U3/0.8.0 Mobile Safari/534.30",
        "expect"  :
        {
            "name"    : "UCBrowser",
            "version" : "11.2.0.915",
            "major"   : "11"
        }
    },
    {
        "desc"    : "UC Browser on Samsung",
        "ua"      : "Mozilla/5.0 (Java; U; Pt-br; samsung-gt-s5620) UCBrowser8.2.1.144/69/352/UCWEB Mobile UNTRUSTED/1.0",
        "expect"  :
        {
            "name"    : "UCBrowser",
            "version" : "8.2.1.144",
            "major"   : "8"
        }
    },
    {
        "desc"    : "UC Browser on Nokia",
        "ua"      : "Mozilla/5.0 (S60V3; U; en-in; NokiaN73)/UC Browser8.4.0.159/28/351/UCWEB Mobile",
        "expect"  :
        {
            "name"    : "UCBrowser",
            "version" : "8.4.0.159",
            "major"   : "8"
        }
    },
    {
        "desc"    : "UC Browser J2ME",
        "ua"      : "UCWEB/2.0 (MIDP-2.0; U; zh-CN; HTC EVO 3D X515m) U2/1.0.0 UCBrowser/10.4.0.558 U2/1.0.0 Mobile",
        "expect"  :
        {
            "name"    : "UCBrowser",
            "version" : "10.4.0.558",
            "major"   : "10"
        }
    },
    {
        "desc"    : "UC Browser J2ME 2",
        "ua"      : "JUC (Linux; U; 2.3.5; zh-cn; GT-I9100; 480*800) UCWEB7.9.0.94/139/800",
        "expect"  :
        {
            "name"    : "UCBrowser",
            "version" : "7.9.0.94",
            "major"   : "7"
        }
    },
    {
        "desc"    : "UP.Browser",
        "ua"      : "BenQ-CF61/1.00/WAP2.0/MIDP2.0/CLDC1.0 UP.Browser/6.3.0.4.c.1.102 (GUI) MMP/2.0",
        "expect"  :
        {
            "name"    : "UP.Browser",
            "version" : "6.3.0.4.c.1.102",
            "major"   : "6"
        }
    },
    {
        "desc": "WeChat on iOS",
        "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12H321 MicroMessenger/6.3.6 NetType/WIFI Language/zh_CN",
        "expect":
        {
            "name": "WeChat",
            "version": "6.3.6",
            "major": "6"
        }
    },
    {
        "desc": "WeChat on Android",
        "ua": "Mozilla/5.0 (Linux; U; Android 5.1; zh-cn; Lenovo K50-t5 Build/LMY47D) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/5.4 TBS/025478 Mobile Safari/533.1 MicroMessenger/6.3.5.50_r1573191.640 NetType/WIFI Language/zh_CN",
        "expect":
        {
            "name": "WeChat",
            "version": "6.3.5.50_r1573191.640",
            "major": "6"
        }
    },
    {
        "desc": "WeiBo on Android",
        "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16A366 Weibo (iPhone8,2__weibo__8.9.3__iphone__os12.0)",
        "expect":
        {
            "name": "weibo",
            "version": "8.9.3",
            "major": "8"
        }
    },
    {
        "desc"    : "Vivaldi",
        "ua"      : "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Vivaldi/1.0.83.38 Safari/537.36",
        "expect"  :
        {
            "name"    : "Vivaldi",
            "version" : "1.0.83.38",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Vivaldi on Mac",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.88 Safari/537.36 Vivaldi/2.4.1488.36",
        "expect"  :
        {
            "name"    : "Vivaldi",
            "version" : "2.4.1488.36",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Vivo Browser",
        "ua"      : "Mozilla/5.0 (Linux; Android 13; 23049RAD8C; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.141 Mobile Safari/537.36 VivoBrowser/16.7.1.1",
        "expect"  :
        {
            "name"    : "Vivo Browser",
            "version" : "16.7.1.1",
            "major"   : "16"
        }
    },
    {
        "desc"    : "w3m",
        "ua"      : "w3m/0.5.1",
        "expect"  :
        {
            "name"    : "w3m",
            "version" : "0.5.1",
            "major"   : "0"
        }
    },
    {
        "desc"    : "Wolvic",
        "ua"      : "Mozilla/5.0 (Android 12; Mobile VR; rv:121.0) Gecko/121.0 Firefox/121.0 Wolvic/1.6.1",
        "expect"  :
        {
            "name"    : "Wolvic",
            "version" : "1.6.1",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Yandex",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/536.5 (KHTML, like Gecko) YaBrowser/1.0.1084.5402 Chrome/19.0.1084.5402 Safari/536.5",
        "expect"  :
        {
            "name"    : "Yandex",
            "version" : "1.0.1084.5402",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Yandex",
        "ua"      : "Mozilla/5.0 (Linux; arm_64; Android 11; M2101K7AG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.125 YaApp_Android/22.70 YaSearchBrowser/22.70 BroPP/1.0 SA/3 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Yandex",
            "version" : "22.70",
            "major"   : "22"
        }
    },
    {
        "desc"    : "Yandex",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.0.2246 Yowser/2.5 Safari/537.36",
        "expect"  :
        {
            "name"    : "Yandex",
            "version" : "23.3.0.2246",
            "major"   : "23"
        }
    },
    {
        "desc"    : "Yandex on Android",
        "ua"      : "Mozilla/5.0 (Linux; arm_64; Android 13; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 YaBrowser/21.3.4.59 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Yandex",
            "version" : "21.3.4.59",
            "major"   : "21"
        }
    },
    {
        "desc"    : "Yandex on iPhone",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.3.330 Mobile/15E148 Safari/604.1",
        "expect"  :
        {
            "name"    : "Yandex",
            "version" : "23.3.3.330",
            "major"   : "23"
        }
    },
    {
        "desc"    : "Yandex on iPad",
        "ua"      : "Mozilla/5.0 (iPad; CPU OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.3.330 Mobile/15E148 Safari/605.1",
        "expect"  :
        {
            "name"    : "Yandex",
            "version" : "23.3.3.330",
            "major"   : "23"
        }
    },
    {
        "desc"    : "Yandex on iPod",
        "ua"      : "Mozilla/5.0 (iPod touch; CPU iPhone 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.3.330 Mobile/15E148 Safari/605.1",
        "expect"  :
        {
            "name"    : "Yandex",
            "version" : "23.3.3.330",
            "major"   : "23"
        }
    },
    {
        "desc"    : "Puffin",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0.1; Lenovo P2a42 Build/MMB29M; en-us) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Puffin/6.0.8.15804AP",
        "expect"  :
        {
            "name"    : "Puffin",
            "version" : "6.0.8.15804AP",
            "major"   : "6"
        }
    },
    {
        "desc"    : "Puffin",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.1.1; ZTE BLADE A0620 Build/NMF26F; ru-ru) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36 Puffin/9.2.0.50586AP",
        "expect"  :
        {
            "name"    : "Puffin",
            "version" : "9.2.0.50586AP",
            "major"   : "9"
        }
    },
    {
        "desc"    : "Microsoft Edge 0.1",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36 Edge/12.0",
        "expect"  :
        {
            "name"    : "Edge",
            "version" : "12.0",
            "major"   : "12"
        }
    },
    {
        "desc"    : "Microsoft Edge 42",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/17.17134",
        "expect"  :
        {
            "name"    : "Edge",
            "version" : "17.17134",
            "major"   : "17"
        }
    },
    {
        "desc"    : "Microsoft Edge 44",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763",
        "expect"  :
        {
            "name"    : "Edge",
            "version" : "18.17763",
            "major"   : "18"
        }
    },
    {
        "desc"    : "Microsoft Edge 100",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.1108.55 Safari/537.36 Edg/100.0.1108.55",
        "expect"  :
        {
            "name"    : "Edge",
            "version" : "100.0.1108.55",
            "major"   : "100"
        }
    },
    {
        "desc"    : "Microsoft Edge on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 11_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 EdgiOS/42.1.1.0 Mobile/15F79 Safari/605.1.15",
        "expect"  :
        {
            "name"    : "Edge",
            "version" : "42.1.1.0",
            "major"   : "42"
        }
    },
    {
        "desc"    : "Microsoft Edge on Android",
        "ua"      : "Mozilla/5.0 (Linux; Android 8.0.0; G8441 Build/47.1.A.12.270) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.123 Mobile Safari/537.36 EdgA/42.0.0.2529",
        "expect"  :
        {
            "name"    : "Edge",
            "version" : "42.0.0.2529",
            "major"   : "42"
        }
    },
    {
        "desc"    : "Microsoft Edge Chromium",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24",
        "expect"  :
        {
            "name"    : "Edge",
            "version" : "74.1.96.24",
            "major"   : "74"
        }
    },
    {
        "desc"    : "Iridium",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Iridium/43.8 Safari/537.36 Chrome/43.0.2357.132",
        "expect"  :
        {
            "name"    : "Iridium",
            "version" : "43.8",
            "major"   : "43"
        }
    },
    {
        "desc"    : "Firefox iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) FxiOS/1.1 Mobile/13B143 Safari/601.1.46",
        "expect"  :
        {
            "name"    : "Mobile Firefox",
            "version" : "1.1",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Firefox on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/112.0 Mobile/15E148 Safari/605.1.15",
        "expect"  :
        {
            "name"    : "Mobile Firefox",
            "version" : "112.0",
            "major"   : "112"
        }
    },
    {
        "desc"    : "Firefox iOS using iPad",
        "ua"      : "Mozilla/5.0 (iPad; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4",
        "expect"  :
        {
            "name"    : "Mobile Firefox",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "QQ on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 10_0_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Mobile/14A456 QQ/6.5.3.410 V1_IPH_SQ_6.5.3_1_APP_A Pixel/1080 Core/UIWebView NetType/WIFI Mem/26",
        "expect"  :
        {
            "name"    : "QQBrowser",
            "version" : "6.5.3.410",
            "major"   : "6"
        }
    },
    {
        "desc"    : "QQ on Android",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0; PRO 6 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/37.0.0.0 Mobile MQQBrowser/6.8 TBS/036824 Safari/537.36 V1_AND_SQ_6.5.8_422_YYB_D PA QQ/6.5.8.2910 NetType/WIFI WebP/0.3.0 Pixel/1080",
        "expect"  :
        {
            "name"    : "QQBrowser",
            "version" : "6.5.8.2910",
            "major"   : "6"
        }
    },
    {
        "desc"    : "WeChat Desktop for Windows Built-in Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 MicroMessenger/6.5.2.501 NetType/WIFI WindowsWechat QBCore/3.43.901.400 QQBrowser/9.0.2524.400",
        "expect"  :
        {
            "name"    : "WeChat",
            "version" : "3.43.901.400",
            "major"   : "3"
        }
    },
    {
        "desc"    : "WeChat Desktop for Windows Built-in Browser major version in 4",
        "ua"      : "mozilla/5.0 (windows nt 10.0; wow64) applewebkit/537.36 (khtml, like gecko) chrome/53.0.2785.116 safari/537.36 qbcore/4.0.1301.400 qqbrowser/9.0.2524.400 mozilla/5.0 (windows nt 6.1; wow64) applewebkit/537.36 (khtml, like gecko) chrome/81.0.4044.138 safari/537.36 nettype/wifi micromessenger/7.0.20.1781(0x6700143b) windowswechat",
        "expect"  :
        {
            "name"    : "WeChat",
            "version" : "4.0.1301.400",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Supposed not to be detected as WeChat",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.124 Safari/537.36 qblink wegame.exe WeGame/5.1.1.11100 QBCore/3.70.107.400 QQBrowser/9.0.2524.400",
        "expect"  :
        {
            "name"    : "QQBrowser",
            "version" : "9.0.2524.400",
            "major"   : "9"
        }
    },
    {
        "desc"    : "GSA on iOS",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) GSA/30.1.161623614 Mobile/14F89 Safari/602.1",
        "expect"  :
        {
            "name"    : "GSA",
            "version" : "30.1.161623614",
            "major"   : "30"
        }
    },
    {
        "desc"    : "Sogou Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.221 Safari/537.36 SE 2.X MetaSr 1.0",
        "expect"  :
        {
            "name"    : "Sogou Explorer",
            "version" : "1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Sogou Mobile Browser",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 SogouMSE,SogouMobileBrowser/3.7.4",
        "expect"  :
        {
            "name"    : "Sogou Mobile",
            "version" : "3.7.4",
            "major"   : "3"
        }
    },
    {
        "desc"    : "LieBao Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.154 Safari/537.36 LBBROWSER",
        "expect"  :
        {
            "name"    : "LBBROWSER"
        }
    },
    {
        "desc"    : "2345 Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.90 Safari/537.36 2345Explorer/9.2.1.17116",
        "expect"  :
        {
            "name"    : "2345Explorer",
            "version" : "9.2.1.17116",
            "major"   : "9"
        }
    },
    {
        "desc"    : "QQBrowserLite",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/10.0.1 Safari/602.2.14 QQBrowserLite/1.1.0",
        "expect"  :
        {
            "name"    : "QQBrowserLite",
            "version" : "1.1.0",
            "major"   : "1"
        }
    },
    {
        "desc"    : "Brave Browser",
        "ua"      : "Brave/4.5.16 CFNetwork/893.13.1 Darwin/17.3.0 (x86_64)",
        "expect"  :
        {
            "name"    : "Brave",
            "version" : "4.5.16",
            "major"   : "4"
        }
    },
    {
        "desc"    : "Whale Browser",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.146 Whale/2.6.90.14 Safari/537.36",
        "expect"  :
        {
            "name"    : "Whale",
            "version" : "2.6.90.14",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Electron",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Atom/1.41.0 Chrome/69.0.3497.128 Electron/4.2.7 Safari/537.36",
        "expect"  :
        {
            "name"    : "Electron",
            "version" : "4.2.7",
            "major"   : "4"
        }
    },
    {
        "desc"    : "IE11 on Windows 7 (ua length >255)",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64; APCPMS=^N201205020840572565478A37A6F9C41BD33F_9975^; Trident/7.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E; MARKANYEPS#25118; Zoom 3.6.0; rv:11.0) like Gecko",
        "expect"  :
        {
            "name"    : "IE",
            "version" : "11.0",
            "major"   : "11"
        }
    },
    {
        "desc"    : "LinkedIn",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [LinkedInApp]",
        "expect"  :
        {
            "name"    : "LinkedIn"
        }
    },
    {
        "desc"    : "Links in Linux",
        "ua"      : "Links (2.xpre7; Linux 2.4.18 i586; x)",
        "expect"  :
        {
            "name"    : "Links",
            "version" : "2.xpre7",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Links in Mac",
        "ua"      : "Links (2.1pre33; Darwin 8.11.0 Power Macintosh; 169x55)",
        "expect"  :
        {
            "name"    : "Links",
            "version" : "2.1pre33",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Links in NetBSD",
        "ua"      : "Links (2.29; NetBSD 10.0 i386; GNU C 10.5; x)",
        "expect"  :
        {
            "name"    : "Links",
            "version" : "2.29",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Links in FreeBSD",
        "ua"      : "Links (2.1pre15; FreeBSD 5.3-RELEASE i386; 196x84)",
        "expect"  :
        {
            "name"    : "Links",
            "version" : "2.1pre15",
            "major"   : "2"
        }
    },
    {
        "desc"    : "Safari including comma in minor version number",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6,2 Safari/605.1.15",
        "expect"  :
        {
            "name"    : "Safari",
            "version" : "15.6,2",
            "major"   : "15"
        }
    },
    {
        "desc"    : "Mobile Safari including comma in minor version number",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 15_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6,2 Mobile/15E148 Safari/604.1",
        "expect"  :
        {
            "name"    : "Mobile Safari",
            "version" : "15.6,2",
            "major"   : "15"
        }
    },
    {
        "desc"     : "Cobalt 23 Master",
        "ua"       : "Mozilla/5.0 (X11; Linux x86_64) Cobalt/23.master.0.0-devel (unlike Gecko) v8/8.8.278.8-jit gles Starboard/15",
        "expect"   : {
            "name"   : "Cobalt",
            "version": "23.0.0",
            "major"  : "23"
        }
    },
    {
        "desc"     : "Cobalt 23 LTS",
        "ua"       : "Mozilla/5.0 (X11; Linux x86_64) Cobalt/23.lts.1.0-qa (unlike Gecko) v8/8.8.278.8-jit gles Starboard/15",
        "expect"   : {
            "name"   : "Cobalt",
            "version": "23.1.0",
            "major"  : "23"
        }
    },
    {
        "desc"     : "Cobalt 11",
        "ua"       : "Mozilla/5.0 (X11; Linux x86_64) Cobalt/11.0-qa (unlike Gecko) Starboard/6",
        "expect"   : {
            "name"   : "Cobalt",
            "version": "11.0",
            "major"  : "11"
        }
    },
    {
        "desc"     : "Cobalt 9",
        "ua"       : "Mozilla/5.0 (X11; Linux x86_64) Cobalt/9.0-qa (unlike Gecko) Starboard/4",
        "expect"   : {
            "name"   : "Cobalt",
            "version": "9.0",
            "major"  : "9"
        }
    },
    {
        "desc"     : "KakaoTalk App Android",
        "ua"       : "Mozilla/5.0 (Linux; Android 12; SM-G988N Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/100.0.4896.79 Mobile Safari/537.36;KAKAOTALK 2409760",
        "expect"   : {
            "name"   : "KAKAOTALK",
            "version": "2409760",
            "major"  : "2409760"
        }
    },
    {
        "desc"     : "KakaoStory App Android",
        "ua"       : "Mozilla/5.0 (Linux; Android 12; SM-G988N Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/100.0.4896.79 Mobile Safari/537.36 KAKAOSTORY/6.8.3_21046",
        "expect"   : {
            "name"   : "KAKAOSTORY",
            "version": "6.8.3_21046",
            "major"  : "6"
        }
    },
    {
        "desc"     : "KakaoTalk App iOS",
        "ua"       : "Mozilla/5.0 (iPhone; CPU; iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 BizWebView KAKAOTALK 9.7.6",
        "expect"   : {
            "name"   : "KAKAOTALK",
            "version": "9.7.6",
            "major"  : "9"
        }
    },
    {
        "desc"     : "Naver App Android",
        "ua"       : "Mozilla/5.0 (Linux; Android 12; SM-G988N Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/90.0.4430.232 Whale/1.0.0.0 Crosswalk/26.90.3.21 Mobile Safari/537.36 NAVER(inapp; search; 1010; 11.11.2)",
        "expect"   : {
            "name"   : "NAVER",
            "version": "11.11.2",
            "major"  : "11"
        }
    },
    {
        "desc"     : "Naver App iOS",
        "ua"       : "Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Mobile/15E148 Safari/605.1 NAVER(inapp; search; 720; 10.25.0; 11PRO)",
        "expect"   : {
            "name"   : "NAVER",
            "version": "10.25.0",
            "major"  : "10"
        }
    },
    {
        "desc"     : "TikTok",
        "ua"       : "Mozilla/5.0 (Linux; Android 11; 21061119AG Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.131 Mobile Safari/537.36 trill_2022109040 JsSdk/1.0 NetType/MOBILE Channel/googleplay AppName/musical_ly app_version/21.9.4 ByteLocale/ru-RU ByteFullLocale/ru-RU Region/KG BytedanceWebview/d8a21c6",
        "expect"   : {
            "name"   : "TikTok",
            "version": "21.9.4",
            "major"  : "21"
        }
    },
    {
        "desc"     : "TikTok",
        "ua"       : "Mozilla/5.0 (iPhone; CPU iPhone OS 14_8 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 musical_ly_21.1.0 JsSdk/2.0 NetType/4G Channel/App Store ByteLocale/ru Region/RU ByteFullLocale/ru-RU isDarkMode/1 WKWebView/1 BytedanceWebview/d8a21c6",
        "expect"   : {
            "name"   : "TikTok",
            "version": "21.1.0",
            "major"  : "21"
        }
    },
    {
        "desc"     : "TikTok",
        "ua"       : "Mozilla/5.0 (Linux; Android 10; STK-LX1 Build/HONORSTK-LX1; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/110.0.5481.153 Mobile Safari/537.36 musical_ly_2022803040 JsSdk/1.0 NetType/WIFI Channel/huaweiadsglobal_int AppName/musical_ly app_version/28.3.4 ByteLocale/en ByteFullLocale/en Region/IQ Spark/1.2.7-alpha.8 AppVersion/28.3.4 PIA/1.5.11 BytedanceWebview/d8a21c6",
        "expect"   : {
            "name"   : "TikTok",
            "version": "28.3.4",
            "major"  : "28"
        }
    },
    {
        "desc"    : "Chrome Mobile",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.1.2; Nexus 5X Build/N2G47W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.83 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "Mobile Chrome",
            "version" : "58.0.3029.83",
            "major"   : "58"
        }
    },
    {
        "desc"    : "Firefox Mobile",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.1.2; Nexus 5X Build/N2G47W) AppleWebKit/537.36 (KHTML, like Gecko) FxiOS/7.5b3349 Mobile/14F89 Safari/603.2.4",
        "expect"  :
        {
            "name"    : "Mobile Firefox",
            "version" : "7.5b3349",
            "major"   : "7"
        }
    },
    {
        "desc"    : "Firefox Mobile",
        "ua"      : "Mozilla/5.0 (Android 5.0; Mobile; rv:41.0) Gecko/41.0 Firefox/41.0",
        "expect"  :
        {
            "name"    : "Mobile Firefox",
            "version" : "41.0",
            "major"   : "41"
        }
    },
    {
        "desc"    : "Snapchat",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Mobile/15E148 Snapchat/12.33.0.36 (like Safari/8614.1.25.0.31, panda)",
        "expect"  :
        {
            "name"    : "Snapchat",
            "version" : "12.33.0.36",
            "major"   : "12"
        }
    },
    {
        "desc"    : "Twitter",
        "ua"      : "Mozilla/5.0 (Linux; Android 13; CPH2531 Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/123.0.6312.120 Mobile Safari/537.36 TwitterAndroid",
        "expect"  :
        {
            "name"    : "Twitter",
            "version" : "",
            "major"   : ""
        }
    },
    {
        "desc"    : "Twitter",
        "ua"      : "Mozilla/5.0 (iPad; CPU OS 15_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19H12 Twitter for iPhone/10.34",
        "expect"  :
        {
            "name"    : "Twitter",
            "version" : "10.34",
            "major"   : "10"
        }
    }
]
`)

type BrowserCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version,omitempty"`
		Major   string `json:"major,omitempty"`
	} `json:"expect"`
}

func TestBrowser(t *testing.T) {
	var UACases []BrowserCase
	_ = json.Unmarshal(testBrowser, &UACases)
	for _, uacase := range UACases {
		u := Parse(uacase.Ua)

		if !assert.Equal(t, uacase.Expect.Name, u.GetBrowser().Name) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Version, u.GetBrowser().Version) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Major, u.GetBrowser().Major) {
			fmt.Println(uacase.Ua)
		}
	}
}

var testCpu = []byte(`
[
    {
        "desc"    : "i686",
        "ua"      : "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:19.0) Gecko/20100101 Firefox/19.0",
        "expect"  :
        {
            "architecture"  : "ia32"
        }
    },
    {
        "desc"    : "i386",
        "ua"      : "Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.7) Gecko/20040628 Epiphany/1.2.6",
        "expect"  :
        {
            "architecture"  : "ia32"
        }
    },
    {
        "desc"    : "x86-64",
        "ua"      : "Opera/9.80 (X11; Linux x86_64; U; Linux Mint; en) Presto/2.2.15 Version/10.10",
        "expect"  :
        {
            "architecture"  : "amd64"
        }
    },
    {
        "desc"    : "Vivaldi on Windows",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Vivaldi/6.0.2979.18",
        "expect"  :
        {
            "architecture"  : "amd64"
        }
    },
    {
        "desc"    : "Vivaldi on Windows",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Vivaldi/6.0.2979.18",
        "expect"  :
        {
            "architecture"  : "amd64"
        }
    },
    {
        "desc"    : "Vivaldi on Linux",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Vivaldi/6.0.2979.18",
        "expect"  :
        {
            "architecture"  : "amd64"
        }
    },
    {
        "desc"    : "Vivaldi on Linux",
        "ua"      : "Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Vivaldi/6.0.2979.18",
        "expect"  :
        {
            "architecture"  : "ia32"
        }
    },
    {
        "desc": "Xiaomi POCO M2 Pro",
        "ua": "Mozilla/5.0 (Linux; arm_64; Android 11; POCO M2 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 YaBrowser/22.11.7.42.00 SA/3 Mobile Safari/537.36",
        "expect"  :
        {
            "architecture"  : "arm64"
        }
    },
    {
        "desc"    : "win64",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; Win64; x64; Trident/6.0; .NET4.0E; .NET4.0C)",
        "expect"  :
        {
            "architecture"  : "amd64"
        }
    },
    {
        "desc"    : "WOW64",
        "ua"      : "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
        "expect"  :
        {
            "architecture"  : "amd64"
        }
    },
    {
        "desc"    : "ARM",
        "ua"      : "Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 635) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "ARMv61",
        "ua"      : "Mozilla/5.0 (X11; U; Linux armv61; en-US; rv:1.9.1b2pre) Gecko/20081015 Fennec/1.0a1",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "ARMv7",
        "ua"      : "Mozilla/5.0 (Linux ARMv7) WebKitGTK+/3.4.9 vimprobable2",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "ARMv7l",
        "ua"      : "Mozilla/5.0 (SMART-TV; X11; Linux armv7l) AppleWebKit/537.42 (KHTML, like Gecko) Chromium/25.0.1349.2 Chrome/25.0.1349.2 Safari/537.42",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "ARMv7l",
        "ua"      : "Mozilla/5.0 (X11; CrOS armv7l 9765.85.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.123 Safari/537.36",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "Nokia N900 Linux mobile",
        "ua"      : "Mozilla/5.0 (Maemo; Linux armv7l; rv:10.0) Gecko/20100101 Firefox/10.0 Fennec/10.0",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "ARMEABI",
        "ua"      : "[FBAN/FB4A;FBAV/237.0.0.44.120;FBBV/170693408;FBDM/{density=1.75,width=720,height=1280};FBLC/en_US;FBRV/172067074;FBCR/ ;FBMF/samsung;FBBD/samsung;FBPN/com.facebook.katana;FBDV/SM-S367VL;FBSV/9;FBBK/1;FBOP/19;FBCA/armeabi-v7a:armeabi;]",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "ARMv8",
        "ua"      : "Mozilla/5.0 (X11; Linux armv8l; rv:45.0) Gecko/20100101 Firefox/45.0",
        "expect"  :
        {
            "architecture"  : "arm64"
        }
    },
    {
        "desc"    : "AARCH64",
        "ua"      : "Mozilla/5.0 (X11; CrOS aarch64 13310.93.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.133 Safari/537.36",
        "expect"  :
        {
            "architecture"  : "arm64"
        }
    },
    {
        "desc"    : "ARM64",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; ARM64; RM-1096) AppleWebKit/537.36 (KHTML like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393",
        "expect"  :
        {
            "architecture"  : "arm64"
        }
    },
    {
        "desc"    : "ARM64",
        "ua"      : "Mozilla/5.0 (Linux; arm_64; Android 9; HRY-LX1T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 YaBrowser/19.12.1.121.00 Mobile Safari/537.36",
        "expect"  :
        {
            "architecture"  : "arm64"
        }
    },
    {
        "desc"    : "Google Search App",
        "ua"      : "Mozilla/5.0 (Linux; Android 6.0; M5s Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/44.0.2403.147 Mobile Safari/537.36 GSA/12.40.17.23.arm64",
        "expect"  :
        {
            "architecture"  : "arm64"
        }
    },
    {
        "desc"    : "Pocket PC",
        "ua"      : "Opera/9.7 (Windows Mobile; PPC; Opera Mobi/35166; U; en) Presto/2.2.1",
        "expect"  :
        {
            "architecture"  : "arm"
        }
    },
    {
        "desc"    : "Mac PowerPC",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 4.5; Mac_PowerPC)",
        "expect"  :
        {
            "architecture"  : "ppc"
        }
    },
    {
        "desc"    : "Mac PowerPC",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC Mac OS; en)",
        "expect"  :
        {
            "architecture"  : "ppc"
        }
    },
    {
        "desc"    : "Mac PowerPC",
        "ua"      : "iCab/2.9.5 (Macintosh; U; PPC; Mac OS X)",
        "expect"  :
        {
            "architecture"  : "ppc"
        }
    },
    {
        "desc"    : "Mac OS X on PowerPC using Firefox",
        "ua"      : "Mozilla/5.0 (Macintosh; PPC Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
        "expect"  :
        {
            "architecture"  : "ppc"
        }
    },
    {
        "desc"    : "UltraSPARC",
        "ua"      : "Mozilla/5.0 (X11; U; SunOS sun4u; en-US; rv:1.9b5) Gecko/2008032620 Firefox/3.0b5",
        "expect"  :
        {
            "architecture"  : "sparc"
        }
    },
    {
        "desc"    : "sparc64",
        "ua"      : "ELinks (0.4.3; NetBSD 3.0.2PATCH sparc64; 141x19)",
        "expect"  :
        {
            "architecture"  : "sparc64"
        }
    },
    {
        "desc"    : "QuickTime",
        "ua"      : "QuickTime/7.5.6 (qtver=7.5.6;cpu=IA32;os=Mac 10.5.8)",
        "expect"  : 
        {
            "architecture"  : "ia32"
        }
    },
    {
        "desc"    : "XBMC",
        "ua"      : "XBMC/12.0 Git:20130127-fb595f2 (Windows NT 6.1;WOW64;Win64;x64; http://www.xbmc.org)",
        "expect"  : 
        {
            "architecture"  : "amd64"
        }
    },
    {
        "desc"    : "IRIX64",
        "ua"      : "Mozilla/4.8C-SGI [en] (X11; U; IRIX64 6.5 IP27",
        "expect"  : 
        {
            "architecture"  : "irix64"
        }
    },
    {
        "desc"    : "68k",
        "ua"      : "'Mozilla/1.1 (Macintosh; U; 68K)'",
        "expect"  : 
        {
            "architecture"  : "68k"
        }
    }
]

`)

type CpuCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Architecture string `json:"architecture"`
	} `json:"expect"`
}

func TestCpu(t *testing.T) {
	var cpuCases []CpuCase
	_ = json.Unmarshal(testCpu, &cpuCases)
	for _, uacase := range cpuCases {
		u := Parse(uacase.Ua)

		if !assert.Equal(t, uacase.Expect.Architecture, u.GetCpu().Architecture) {
			fmt.Println(uacase.Ua)
		}
	}
}

var testDevice = []byte(`
[
    {
        "desc": "K",
        "ua": "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "",
            "model": "K",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Nexus 7",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 7 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "Nexus 7",
            "type": "tablet"
        }
    },
    {
        "desc": "ASUS Padfone",
        "ua": "Mozilla/5.0 (Linux; Android 4.1.1; PadFone 2 Build/JRO03L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.117 Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "PadFone",
            "type": "tablet"
        }
    },
    {
        "desc": "ASUS ZenPad 10",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; P00C Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "P00C",
            "type": "tablet"
        }
    },
    {
        "desc": "ASUS ZenPad Z8s",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; ASUS_P00J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.111 Safari/537.36\n",
        "expect": {
            "vendor": "ASUS",
            "model": "P00J",
            "type": "tablet"
        }
    },
    {
        "desc": "ASUS ROG",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; ZS600KL Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.126 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "ZS600KL",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS ROG II",
        "ua": "Mozilla/5.0 (Linux; Android 9; ASUS_I001DA Build/PKQ1.190414.001; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "I001DA",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone 2",
        "ua": "Mozilla/5.0 (Linux; Android 5.0; ASUS ZenFone 2 Build/LRX22C) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/37.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "ZenFone 2",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone 3 Deluxe",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; ASUS_Z016D Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "Z016D",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone 5",
        "ua": "Mozilla/5.0 (Linux; Android 8.0; ZE620KL Build/OPR1.170623.032) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.158 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "ZE620KL",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone 7",
        "ua": "Mozilla/5.0 (Linux; Android 10; ASUS_I002D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.81 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "I002D",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone 7 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; ZS671KS) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "ZS671KS",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone Max Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; ZB602KL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "ZB602KL",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone Max Pro (M1)",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; ASUS_X00TD Build/OPM1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.137 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "X00TD",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone Max M2",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; ASUS_X01AD) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.99 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "X01AD",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone Max Pro M2",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; ASUS_X01BDA) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.99 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "X01BDA",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone Go",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; ASUS_X009DA Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "X009DA",
            "type": "mobile"
        }
    },
    {
        "desc": "ASUS Zenfone 2 Laser",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; ASUS_Z00ED) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "ASUS",
            "model": "Z00ED",
            "type": "mobile"
        }
    },
    {
        "desc": "Acer Iconia A1-810",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.2; A1-810 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.109 Safari/537.36",
        "expect": {
            "vendor": "Acer",
            "model": "A1-810",
            "type": "tablet"
        }
    },
    {
        "desc": "BlackBerry Priv",
        "ua": "User-Agent: Mozilla/5.0 (Linux; Android 5.1.1; STV100-1 Build/LMY47V; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/46.0.2490.76 Mobile Safari/537.36",
        "expect": {
            "vendor": "BlackBerry",
            "model": "STV100-1",
            "type": "mobile"
        }
    },
    {
        "desc": "BlackBerry Keyone",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; BBB100-1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.111 Mobile Safari/537.36",
        "expect": {
            "vendor": "BlackBerry",
            "model": "BBB100-1",
            "type": "mobile"
        }
    },
    {
        "desc": "BlackBerry Key2",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; BBF100-1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36",
        "expect": {
            "vendor": "BlackBerry",
            "model": "BBF100-1",
            "type": "mobile"
        }
    },
    {
        "desc": "BlackBerry Key2 LE",
        "ua": "User-Agent: Mozilla/5.0 (Linux; Android 8.1.0; BBE100-1 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497 Mobile Safari/537.36",
        "expect": {
            "vendor": "BlackBerry",
            "model": "BBE100-1",
            "type": "mobile"
        }
    },
    {
        "desc": "Blackview 4900Pro",
        "ua": "Mozilla/5.0 (Linux; Android 12; BV4900Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "",
            "model": "BV4900Pro",
            "type": "mobile"
        }
    },
    {
        "desc": "Desktop (IE11 with Tablet string)",
        "ua": "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; .NET4.0E; .NET4.0C; .NET CLR 3.5.30729; .NET CLR 2.0.50727; .NET CLR 3.0.30729; Tablet PC 2.0; GWX:MANAGED; rv:11.0) like Gecko",
        "expect": {
            "vendor": "",
            "model": "",
            "type": ""
        }
    },
    {
        "desc": "Mobile (DuckDuckGo mobile browser)",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.131 Mobile DuckDuckGo/5 Safari/537.36",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "mobile"
        }
    },
    {
        "desc": "Fairphone 1U",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.2.2; FP1U Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "expect": {
            "vendor": "Fairphone",
            "model": "FP1U",
            "type": "mobile"
        }
    },
    {
        "desc": "Fairphone 2",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.2; FP2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Fairphone",
            "model": "FP2",
            "type": "mobile"
        }
    },
    {
        "desc": "Fairphone 3",
        "ua": "Mozilla/5.0 (Linux; Android 9; FP3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.93 Mobile Safari/537.36",
        "expect": {
            "vendor": "Fairphone",
            "model": "FP3",
            "type": "mobile"
        }
    },
    {
        "desc": "HTC Desire 820",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; HTC Desire 820 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
        "expect": {
            "vendor": "HTC",
            "model": "Desire 820",
            "type": "mobile"
        }
    },
    {
        "desc": "HTC Evo Shift 4G",
        "ua": "Mozilla/5.0 (Linux; U; Android 2.3.4; en-us; Sprint APA7373KT Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0",
        "expect": {
            "vendor": "Sprint",
            "model": "APA7373KT",
            "type": "mobile"
        }
    },
    {
        "desc": "HTC Nexus 9",
        "ua": "Mozilla/5.0 (Linux; Android 5.0; Nexus 9 Build/LRX21R) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Mobile Crosswalk/7.36.154.13 Safari/537.36",
        "expect": {
            "vendor": "HTC",
            "model": "Nexus 9",
            "type": "tablet"
        }
    },
    {
        "desc": "Huawei Honor",
        "ua": "Mozilla/5.0 (Linux; U; Android 2.3; xx-xx; U8860 Build/HuaweiU8860) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
        "expect": {
            "vendor": "Huawei",
            "model": "U8860",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Honor 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; YAL-L41) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YAL-L41",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Honor 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; YAL-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YAL-AL10",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nexus 6P",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 6P Build/MTC19V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.81 Mobile Safari/537",
        "expect": {
            "vendor": "Huawei",
            "model": "Nexus 6P",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P10",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; VTR-L09 Build/HUAWEIVTR-L09; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "VTR-L09",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Y3II",
        "ua": "Mozilla/5.0 (Linux; U; Android 5.1; xx-xx; HUAWEI LUA-L03 Build/HUAWEILUA-L03) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/39.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LUA-L03",
            "type": "mobile"
        }
    },
    {
        "desc": "HUAWEI MediaPad M3 Lite 10",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; BAH-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "BAH-L09",
            "type": "tablet"
        }
    },
    {
        "desc": "HUAWEI MediaPad M5 Lite",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; BAH2-W19 Build/HUAWEIBAH2-W19; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "BAH2-W19",
            "type": "tablet"
        }
    },
    {
        "desc": "HUAWEI MediaPad M5",
        "ua": "Mozilla/5.0 (Linux; Android 9; SHT-AL09 Build/HUAWEISHT-AL09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "SHT-AL09",
            "type": "tablet"
        }
    },
    {
        "desc": "HUAWEI MediaPad T5",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; AGS2-L09 Build/HUAWEIAGS2-L09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/84.0.4147.125 Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "AGS2-L09",
            "type": "tablet"
        }
    },
    {
        "desc": "HUAWEI MediaPad T10",
        "ua": "Mozilla/5.0 (Linux; Android 10; AGR-W09 Build/HUAWEIAGR-W09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "AGR-W09",
            "type": "tablet"
        }
    },
    {
        "desc": "HUAWEI MediaPad T10s",
        "ua": "Mozilla/5.0 (Linux; Android 10; AGS3-W09 Build/HUAWEIAGS3-W09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "AGS3-W09",
            "type": "tablet"
        }
    },
    {
        "desc": "Huawei MatePad T 10",
        "ua": "Mozilla/5.0 (Linux; Android 10; AGR-L09; HMSCore 5.0.4.301) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 HuaweiBrowser/11.0.3.304 Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "AGR-L09",
            "type": "tablet"
        }
    },
    {
        "desc": "Huawei M3",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; BTV-W09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "BTV-W09",
            "type": "tablet"
        }
    },
    {
        "desc": "Huawei Mate 10 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 8.0; BLA-L29 Build/HUAWEIBLA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3236.6 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "BLA-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate X",
        "ua": "Mozilla/5.0 (Linux; Android 9; TAH-AN00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.111 Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "TAH-AN00",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate X2",
        "ua": "Mozilla/5.0 (Linux; Android 10; TET-AN00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "TET-AN00",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 X",
        "ua": "Mozilla/5.0 (Linux; Android 9; EVR-L29 Build/HUAWEIEVR-L29; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "EVR-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; LYA-L09) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LYA-L09",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; LYA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LYA-AL00",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; LYA-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LYA-AL10",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; LYA-L0C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LYA-L0C",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; LYA-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LYA-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; LYA-TL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LYA-TL00",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 50 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 12; DCO-LX9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "DCO-LX9",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P20 Lite",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; ANE-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.143 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "ANE-LX1",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P20",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; EML-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "EML-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P20 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; CLT-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "CLT-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P30",
        "ua": "Mozilla/5.0 (Linux; Android 9; ELE-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "ELE-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P30 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.143 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "VOG-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P40",
        "ua": "Mozilla/5.0 (Linux; Android 10; ANA-AN00 Build/HUAWEIANA-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/11.26 SP-engine/2.22.0 baiduboxapp/11.26.0.10 (Baidu; P1 10) NABar/1.0",
        "expect": {
            "vendor": "Huawei",
            "model": "ANA-AN00",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P40 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; ELS-AN00 Build/HUAWEIELS-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Mobile Safari/537.36 mailapp/6.0.0",
        "expect": {
            "vendor": "Huawei",
            "model": "ELS-AN00",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei 30 Pro+",
        "ua": "Mozilla/5.0 (Linux; Android 10; EBG-AN10 Build/HUAWEIEBG-AN10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36 EdgA/42.0.0.2741",
        "expect": {
            "vendor": "Huawei",
            "model": "EBG-AN10",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei 30S",
        "ua": "Mozilla/5.0 (Linux; Android 10; CDY-AN90 Build/HUAWEICDY-AN90; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Mobile Safari/537.36 mailapp/5.8.0",
        "expect": {
            "vendor": "Huawei",
            "model": "CDY-AN90",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nova 5T",
        "ua": "Mozilla/5.0 (Linux; Android 10; YAL-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YAL-L21",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nova 5T",
        "ua": "Mozilla/5.0 (Linux; Android 10; YAL-L61) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YAL-L61",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nova 5T",
        "ua": "Mozilla/5.0 (Linux; Android 10; YAL-L71) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YAL-L71",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nova 5T",
        "ua": "Mozilla/5.0 (Linux; Android 10; YAL-L61D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YAL-L61D",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nova 5T",
        "ua": "Mozilla/5.0 (Linux; Android 10; YALE-L61A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YALE-L61A",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nova 5T",
        "ua": "Mozilla/5.0 (Linux; Android 10; YALE-L61D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YALE-L61D",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Nova 5T",
        "ua": "Mozilla/5.0 (Linux; Android 10; YALE-L71A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "YALE-L71A",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Enjoy10e",
        "ua": "Dalvik/2.1.0 (Linux; U; Android 10; MED-AL00 Build/HUAWEIMED-AL00)",
        "expect": {
            "vendor": "Huawei",
            "model": "MED-AL00",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Honor 6A",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; DLI-L22 Build/HONORDLI-L22; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.116 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/252.0.0.22.355;]",
        "expect": {
            "vendor": "Huawei",
            "model": "DLI-L22",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Honor 7",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; PLK-L01 Build/HONORPLK-L01; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "PLK-L01",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei 10 Lite",
        "ua": "Mozilla/5.0 (Linux; Android 9; HRY-LX1 Build/HONORHRY-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.91 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "HRY-LX1",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Y7 2018",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; LDN-L01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.62 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "LDN-L01",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Honor 8X",
        "ua": "Mozilla/5.0 (Linux; Android 9; JSN-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "JSN-L21",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Y6 2019",
        "ua": "Mozilla/5.0 (Linux; Android 9; MRD-LX1N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "MRD-LX1N",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Y9 2019",
        "ua": "Mozilla/5.0 (Linux; Android 9; JKM-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "JKM-LX2",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Y5",
        "ua": "Mozilla/5.0 (Linux; Android 9; AMN-LX3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "AMN-LX3",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Y7p",
        "ua": "Mozilla/5.0 (Linux; Android 9; ART-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.92 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "ART-L29",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Mate 20 Lite",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; SNE-LX1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "SNE-LX1",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei P10 Lite",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; WAS-LX1A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "WAS-LX1A",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Y5 Lite 2018",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; DRA-LX5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "DRA-LX5",
            "type": "mobile"
        }
    },
    {
        "desc": "Huawei Honor 8C",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; BKK-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Huawei",
            "model": "BKK-LX2",
            "type": "mobile"
        }
    },
    {
        "desc": "Infinix Hot 7 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; Infinix X625C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Infinix",
            "model": "X625C",
            "type": "mobile"
        }
    },
    {
        "desc": "Infinix Hot 10T",
        "ua": "Mozilla/5.0 (Linux; Android 11; Infinix X689C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Infinix",
            "model": "X689C",
            "type": "mobile"
        }
    },
    {
        "desc": "Infinix Hot 11s",
        "ua": "Mozilla/5.0 (Linux; Android 11; Infinix X6812 Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/111.0.5563.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Infinix",
            "model": "X6812",
            "type": "mobile"
        }
    },
    {
        "desc": "Infinix Smart 5",
        "ua": "Mozilla/5.0 (Linux; Android 10; Infinix X657C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Infinix",
            "model": "X657C",
            "type": "mobile"
        }
    },
    {
        "desc": "Infinix Zero 5G",
        "ua": "Mozilla/5.0 (Linux; Android 12; Infinix X6815B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Infinix",
            "model": "X6815B",
            "type": "mobile"
        }
    },
    {
        "desc": "Apple Desktop",
        "ua": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
        "expect": {
            "vendor": "Apple",
            "model": "Macintosh",
            "type": ""
        }
    },
    {
        "desc": "Apple Watch",
        "ua": "atc/1.0 watchOS/7.3.3 model/Watch4,2 hwp/t8006 build/18S830 (6; dt:191)",
        "expect": {
            "vendor": "Apple",
            "model": "watch",
            "type": "wearable"
        }
    },
    {
        "desc": "iPad using UCBrowser",
        "ua": "Mozilla/5.0 (iPad; U; CPU OS 11_2 like Mac OS X; zh-CN; iPad5,3) AppleWebKit/534.46 (KHTML, like Gecko) UCBrowser/3.0.1.776 U3/ Mobile/10A403 Safari/7543.48.3",
        "expect": {
            "vendor": "Apple",
            "model": "iPad",
            "type": "tablet"
        }
    },
    {
        "desc": "iPad Air",
        "ua": "Mozilla/5.0 (iPad; CPU OS 12_4_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPad4,1;FBMD/iPad;FBSN/iOS;FBSV/12.4.5;FBSS/2;FBID/tablet;FBLC/en_US;FBOP/5;FBCR/]",
        "expect": {
            "vendor": "Apple",
            "model": "iPad",
            "type": "tablet"
        }
    },
    {
        "desc": "iPad using Facebook Browser",
        "ua": "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) WebKit/8610 (KHTML, like Gecko) Mobile/18D70 [FBAN/FBIOS;FBDV/iPad7,11;FBMD/iPad;FBSN/iOS;FBSV/14.4.2;FBSS/2;FBID/tablet;FBLC/en_US;FBOP/5]",
        "expect": {
            "vendor": "Apple",
            "model": "iPad",
            "type": "tablet"
        }
    },
    {
        "desc": "iPod",
        "ua": "Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B554a Safari/9537.53",
        "expect": {
            "vendor": "Apple",
            "model": "iPod touch",
            "type": "mobile"
        }
    },
    {
        "desc": "JVC LT-43V55LFA Smart TV",
        "ua": "Mozilla/5.0 (Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 OPR/40.0.2207.0 OMI/4.9.0.237.DOM3-OPT.245 Model/Vestel-MB211 VSTVB MB200 HbbTV/1.2.1 (; JVC; MB211; 3.19.4.2; _TV_NT72563_2017 SmartTvA/3.0.0",
        "expect": {
            "vendor": "JVC",
            "model": "MB211",
            "type": "smarttv"
        }
    },
    {
        "desc": "JVC LT-43V65LUA Smart TV",
        "ua": "Mozilla/5.0 (Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 OPR/40.0.2207.0 OMI/4.9.0.237.DOM3-OPT.245 Model/Vestel-MB130 VSTVB MB100 HbbTV/1.2.1 (; JVC; MB130; 5.7.20.0; _TV_G10_2017;) SmartTvA/3.0.0",
        "expect": {
            "vendor": "JVC",
            "model": "MB130",
            "type": "smarttv"
        }
    },
    {
        "desc": "Kobo eReader",
        "ua": "Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Kobo eReader Safari/538.1",
        "expect": {
            "vendor": "Kobo",
            "model": "eReader",
            "type": "tablet"
        }
    },
    {
        "desc": "Kobo Touch",
        "ua": "Mozilla/5.0 (Linux; U; Android 2.0; en-us;) AppleWebKit/538.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/538.1 (Kobo Touch 0377/4.20.14622)",
        "expect": {
            "vendor": "Kobo",
            "model": "Touch",
            "type": "tablet"
        }
    },
    {
        "desc": "Lenovo Tab 2",
        "ua": "Mozilla/5.0 (Linux; Android 5.0.1; Lenovo TAB 2 A7-30HC Build/LRX21M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.157 Safari/537.36",
        "expect": {
            "vendor": "Lenovo",
            "model": "TAB 2 A7",
            "type": "tablet"
        }
    },
    {
        "desc": "Lenovo Phone",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; Lenovo PB2-650M Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.105 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/311.0.0.44.117;]",
        "expect": {
            "vendor": "Lenovo",
            "model": "PB2-650M",
            "type": "mobile"
        }
    },
    {
        "desc": "Lenovo Tab 3 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; Lenovo YT3-X90F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.99 Safari/537.36",
        "expect": {
            "vendor": "Lenovo",
            "model": "YT3-X90F",
            "type": "tablet"
        }
    },
    {
        "desc": "Lenovo Tab 4",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.1; Lenovo TB-X304F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.99 Safari/537.36",
        "expect": {
            "vendor": "Lenovo",
            "model": "TB-X304F",
            "type": "tablet"
        }
    },
    {
        "desc": "Lenovo Tab 4",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; Lenovo TAB 2 A7-30HC) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
        "expect": {
            "vendor": "Lenovo",
            "model": "TAB 2 A7",
            "type": "tablet"
        }
    },
    {
        "desc": "Lenovo Tab 7 Essential",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; Lenovo TB-7304X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
        "expect": {
            "vendor": "Lenovo",
            "model": "TB-7304X",
            "type": "tablet"
        }
    },
    {
        "desc": "Lenovo Tab M10",
        "ua": "Mozilla/5.0 (Linux; arm_64; Android 9; Lenovo TB-X606F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 YaBrowser/20.9.4.99.01 Safari/537.36",
        "expect": {
            "vendor": "Lenovo",
            "model": "TB-X606F",
            "type": "tablet"
        }
    },
    {
        "desc": "Lenovo IdeaTab S6000",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.2; IdeaTab S6000-H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 YaBrowser/18.11.1.1011.01 Safari/537.36",
        "expect": {
            "vendor": "Lenovo",
            "model": "IdeaTab S6000-H",
            "type": "tablet"
        }
    },
    {
        "desc": "LG V40 ThinQ",
        "ua": "Mozilla/5.0 (Linux; Android 9; LM-V405) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LM-V405",
            "type": "mobile"
        }
    },
    {
        "desc": "LG K30",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; LM-X410.F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LM-X410.F",
            "type": "mobile"
        }
    },
    {
        "desc": "LG K30",
        "ua": "Mozilla/5.0 (Linux; Android 9; LM-X410.FGN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.93 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LM-X410.FGN",
            "type": "mobile"
        }
    },
    {
        "desc": "LG K40",
        "ua": "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.57 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LM-X420",
            "type": "mobile"
        }
    },
    {
        "desc": "LG Stylo 4",
        "ua": "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.57 Mobile Safari/537.36",
        "expect": {
            "vendor": "",
            "model": "LM-Q710(FGN)",
            "type": "mobile"
        }
    },
    {
        "desc": "LG Stylo 5",
        "ua": "Mozilla/5.0 (Linux; Android 9; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LM-Q720",
            "type": "mobile"
        }
    },
    {
        "desc": "LG G7 ThinQ",
        "ua": "Mozilla/5.0 (Linux; Android 9; LM-G710VM Build/PKQ1.181105.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LM-G710VM",
            "type": "mobile"
        }
    },
    {
        "desc": "LG K20",
        "ua": "Mozilla/5.0 (Android 13; Mobile; LG-M255; rv:111.0) Gecko/111.0 Firefox/111.0",
        "expect": {
            "vendor": "LG",
            "model": "M255",
            "type": "mobile"
        }
    },
    {
        "desc": "LG K500",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; LG-K500 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "K500",
            "type": "mobile"
        }
    },
    {
        "desc": "LG Nexus 4",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.1; Nexus 4 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
        "expect": {
            "vendor": "LG",
            "model": "Nexus 4",
            "type": "mobile"
        }
    },
    {
        "desc": "LG Nexus 4",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.3; en-us; Google Nexus 4 - 4.3 - API 18 - 768x1280 Build/JLS36G) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "expect": {
            "vendor": "LG",
            "model": "Nexus 4",
            "type": "mobile"
        }
    },
    {
        "desc": "LG Nexus 5",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 5 Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
        "expect": {
            "vendor": "LG",
            "model": "Nexus 5",
            "type": "mobile"
        }
    },
    {
        "desc": "LG Wing",
        "ua": "Mozilla/5.0 (Linux; Android 10; LM-F100N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LM-F100N",
            "type": "mobile"
        }
    },
    {
        "desc": "LG Smart TV",
        "ua": "Mozilla/5.0 (DirectFB; U; Linux mips; en) AppleWebKit/528.5+ (KHTML, like Gecko, Safari/528.5+) LG Browser (; LG NetCast.TV-2011)",
        "expect": {
            "vendor": "LG",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "LG Smart TV",
        "ua": "Mozilla/5.0 (Linux; NetCast; U) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/53.0.2785 34 Safari/537.31 SmartTV/8.5",
        "expect": {
            "vendor": "LG",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "LG Android TV",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.2.2; zh-cn; LG Android TV Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
        "expect": {
            "vendor": "LG",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Loewe Smart TV",
        "ua": "Mozilla/5.0 (Linux; U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36 OPR/46.0.2207.0 LOEWE-SL410/5.2.0.0 HbbTV/1.4.1 (; LOEWE; SL410; LOH/5.2.0.0;;) FVC/3.0 (LOEWE; SL410;) CE-HTML/1.0 Config (L:deu,CC:DEU) NETRANGEMMH",
        "expect": {
            "vendor": "LOEWE",
            "model": "SL410",
            "type": "smarttv"
        }
    },
    {
        "desc": "Meizu M5 Note",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; M5 Note Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043024 Safari/537.36 MicroMessenger/6.5.7.1040 NetType/WIFI Language/zh_CN",
        "expect": {
            "vendor": "Meizu",
            "model": "M5 Note",
            "type": "mobile"
        }
    },
    {
        "desc": "Microsoft Lumia 950",
        "ua": "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586",
        "expect": {
            "vendor": "Microsoft",
            "model": "Lumia 950",
            "type": "mobile"
        }
    },
    {
        "desc": "Microsoft Surface Duo",
        "ua": "Dalvik/2.1.0 (Linux; U; Android 10; Surface Duo Build/2020.1014.61)",
        "expect": {
            "vendor": "Microsoft",
            "model": "Surface Duo",
            "type": "tablet"
        }
    },
    {
        "desc": "Motorola Moto X",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.4; XT1097 Build/KXE21.187-38) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.109 Mobile Safari/537.36",
        "expect": {
            "vendor": "Motorola",
            "model": "XT1097",
            "type": "mobile"
        }
    },
    {
        "desc": "Motorola Moto Z3 Play",
        "ua": "Mozilla/5.0 (Linux; Android 9; Moto Z3 Play) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Motorola",
            "model": "Moto Z3 Play",
            "type": "mobile"
        }
    },
    {
        "desc": "Meizu M3S",
        "ua": "Mozilla/5.0 (X11; Linux; Android 5.1; MZ-M3s Build/LMY47I) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrom/45.0.2454.94 Mobile Safari/537.36",
        "expect": {
            "vendor": "Meizu",
            "model": "M3s",
            "type": "mobile"
        }
    },
    {
        "desc": "Microsoft Lumia 950",
        "ua": "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/13.10586",
        "expect": {
            "vendor": "Microsoft",
            "model": "Lumia 950",
            "type": "mobile"
        }
    },
    {
        "desc": "Motorola Nexus 6",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.20 Mobile Safari/537.36",
        "expect": {
            "vendor": "Motorola",
            "model": "Nexus 6",
            "type": "mobile"
        }
    },
    {
        "desc": "Motorola Droid RAZR 4G",
        "ua": "Mozilla/5.0 (Linux; U; Android 2.3; xx-xx; DROID RAZR 4G Build/6.5.1-73_DHD-11_M1-29) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
        "expect": {
            "vendor": "Motorola",
            "model": "DROID RAZR 4G",
            "type": "mobile"
        }
    },
    {
        "desc": "Motorola RAZR 2019",
        "ua": "Mozilla/5.0 (Linux; Android 9; motorola razr) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/11.1 Chrome/75.0.3770.143 Mobile Safari/537.36",
        "expect": {
            "vendor": "Motorola",
            "model": "razr",
            "type": "mobile"
        }
    },
    {
        "desc": "iPhone",
        "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
        "expect": {
            "vendor": "Apple",
            "model": "iPhone",
            "type": "mobile"
        }
    },
    {
        "desc": "iPhone SE",
        "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone8,4;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/2;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
        "expect": {
            "vendor": "Apple",
            "model": "iPhone",
            "type": "mobile"
        }
    },
    {
        "desc": "iPhone SE using Facebook App",
        "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone8,4;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/2;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
        "expect": {
            "vendor": "Apple",
            "model": "iPhone",
            "type": "mobile"
        }
    },
    {
        "desc": "iPhone 11 Pro Max",
        "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone12,5;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
        "expect": {
            "vendor": "Apple",
            "model": "iPhone",
            "type": "mobile"
        }
    },
    {
        "desc": "iPhone XS",
        "ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone11,2;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
        "expect": {
            "vendor": "Apple",
            "model": "iPhone",
            "type": "mobile"
        }
    },
    {
        "desc": "iPod touch",
        "ua": "Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0_2 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A501 Safari/9537.53",
        "expect": {
            "vendor": "Apple",
            "model": "iPod touch",
            "type": "mobile"
        }
    },
    {
        "desc": "itel A25",
        "ua": "Mozilla/5.0 (Linux; Android 9; itel L5002) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.130 Mobile Safari/537.36 OPR/63.3.3216.58675",
        "expect": {
            "vendor": "itel",
            "model": "L5002",
            "type": "mobile"
        }
    },
    {
        "desc": "itel A50",
        "ua": "Mozilla/5.0 (Linux; U; Android 14; itel A667L Build/UP1A.231005.007; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/127.0.6533.103 Mobile Safari/537.36 OPR/83.1.2254.73239",
        "expect": {
            "vendor": "itel",
            "model": "A667L",
            "type": "mobile"
        }
    },
    {
        "desc": "itel KidPad 1",
        "ua": "Mozilla/5.0 (Linux; Android 10; Itel W7001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
        "expect": {
            "vendor": "itel",
            "model": "W7001",
            "type": "tablet"
        }
    },
    {
        "desc": "itel Pad One",
        "ua": "Mozilla/5.0 (Linux; Android 12; itel P10001L Build/SP1A.210812.016) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.172 Safari/537.36",
        "expect": {
            "vendor": "itel",
            "model": "P10001L",
            "type": "tablet"
        }
    },
    {
        "desc": "itel RS4",
        "ua": "Mozilla/5.0 (Linux; Android 13; itel S666LN Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/125.0.6422.165 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/468.1.0.56.78;]",
        "expect": {
            "vendor": "itel",
            "model": "S666LN",
            "type": "mobile"
        }
    },
    {
        "desc": "itel Vision 2S",
        "ua": "Mozilla/5.0 (Linux; Android 11; itel P651L Build/RP1A.201005.001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
        "expect": {
            "vendor": "itel",
            "model": "P651L",
            "type": "mobile"
        }
    },
    {
        "desc": "Moto X",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.2; xx-xx; XT1058 Build/13.9.0Q2.X-70-GHOST-ATT_LE-2) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "expect": {
            "vendor": "Motorola",
            "model": "XT1058",
            "type": "mobile"
        }
    },
    {
        "desc": "Motorola Moto g(6) Play",
        "ua": "Mozilla/5.0 (Linux; Android 9; moto g(6) play) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Motorola",
            "model": "moto g(6) play",
            "type": "mobile"
        }
    },
    {
        "desc": "Motorola Moto g(7) Supra",
        "ua": "Mozilla/5.0 (Linux; Android 9; moto g(7) supra Build/PCOS29.114-134-2; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/73.0.3683.90 Mobile Safari/537.36",
        "expect": {
            "vendor": "Motorola",
            "model": "moto g(7) supra",
            "type": "mobile"
        }
    },
    {
        "desc": "Motorola Moto E",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.1; Moto E (4) Build/NDQS26.69-64-11-7; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36",
        "expect": {
            "vendor": "Motorola",
            "model": "Moto E (4)",
            "type": "mobile"
        }
    },
    {
        "desc": "Nokia3xx",
        "ua": "Nokia303/14.87 CLDC-1.1",
        "expect": {
            "vendor": "Nokia",
            "model": "303",
            "type": "mobile"
        }
    },
    {
        "desc": "Nokia 3.2",
        "ua": "Mozilla/5.0 (Linux; Android 10; Nokia 3.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36",
        "expect": {
            "vendor": "Nokia",
            "model": "3.2",
            "type": "mobile"
        }
    },
    {
        "desc": "Nokia 7",
        "ua": "Mozilla/5.0 (Linux; Android 11; Nokia 7.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Nokia",
            "model": "7.2",
            "type": "mobile"
        }
    },
    {
        "desc": "Nokia N9",
        "ua": "Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13",
        "expect": {
            "vendor": "Nokia",
            "model": "N9",
            "type": "mobile"
        }
    },
    {
        "desc": "Nokia 2720 Flip",
        "ua": "Mozilla/5.0 (Mobile; Nokia_2720_Flip; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5.2",
        "expect": {
            "vendor": "Nokia",
            "model": "2720 Flip",
            "type": "mobile"
        }
    },
    {
        "desc": "Nothing 1",
        "ua": "Mozilla/5.0 (Linux; Android 13; A063) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/22.0 Chrome/111.0.5563.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Nothing",
            "model": "A063",
            "type": "mobile"
        }
    },
    {
        "desc": "Nothing 2",
        "ua": "Mozilla/5.0 (Linux; Android 14; A065 Build/UP1A.231005.007; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/125.0.6422.53 Mobile Safari/537.36",
        "expect": {
            "vendor": "Nothing",
            "model": "A065",
            "type": "mobile"
        }
    },
    {
        "desc": "Nothing 2a",
        "ua": "Mozilla/5.0 (Linux; Android 14; A142 Build/UP1A.231005.007; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/126.0.6478.71 Mobile Safari/537.36",
        "expect": {
            "vendor": "Nothing",
            "model": "A142",
            "type": "mobile"
        }
    },
    {
        "desc": "Oculus Quest",
        "ua": "Mozilla/5.0 (Linux; Android 10; Quest) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/15.0.0.0.22.280317669 SamsungBrowser/4.0 Chrome/89.0.4389.90 VR Safari/537.36",
        "expect": {
            "vendor": "Facebook",
            "model": "Quest",
            "type": "xr"
        }
    },
    {
        "desc": "Oculus Quest 2",
        "ua": "Mozilla/5.0 (Linux; Android 10; Quest 2) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/15.0.0.0.22.280317669 SamsungBrowser/4.0 Chrome/89.0.4389.90 VR Safari/537.36",
        "expect": {
            "vendor": "Facebook",
            "model": "Quest 2",
            "type": "xr"
        }
    },
    {
        "desc": "Oculus Quest 3",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; Quest 3) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/31.4.0.6.51.566757996 Chrome/120.0.6099.283 VR Safari/537.36",
        "expect": {
            "vendor": "Facebook",
            "model": "Quest 3",
            "type": "xr"
        }
    },
    {
        "desc": "Oculus Quest Pro",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; Quest Pro) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/24.4.0.22.60.426469926 SamsungBrowser/4.0 Chrome/106.0.5249.181 VR Safari/537.36",
        "expect": {
            "vendor": "Facebook",
            "model": "Quest Pro",
            "type": "xr"
        }
    },
    {
        "desc": "Issue #747",
        "ua": "python-requests/2.25.1",
        "expect": {
            "vendor": "",
            "model": "",
            "type": ""
        }
    },
    {
        "desc": "OnePlus One",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.4; A0001 Build/KTU84Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.59 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "A0001",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus One",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; OnePlus One A0001 Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.117 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "A0001",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 2",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; ONE A2003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.93 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "A2003",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 3",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.1; ONEPLUS A3000 Build/NMF26F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.98 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "A3000",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 6",
        "ua": "Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.89 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "A6003",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 6T",
        "ua": "Mozilla/5.0 (Linux; Android 9; ONEPLUS A6010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "A6010",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 7T Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.57 Mobile Safari/537.36 EdgA/110.0.1587.66",
        "expect": {
            "vendor": "",
            "model": "HD1913",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 8T",
        "ua": "Mozilla/5.0 (Linux; Android 11; KB2005) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "KB2005",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 8 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; IN2025) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.119 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "IN2025",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus 10RT",
        "ua": "Mozilla/5.0 (Linux; Android 13; CPH2413) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "CPH2413",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus Nord N100",
        "ua": "Mozilla/5.0 (Linux; Android 10; BE2015 Build/QKQ1.200719.002; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "BE2015",
            "type": "mobile"
        }
    },
    {
        "desc": "OnePlus Nord N10 5G",
        "ua": "Mozilla/5.0 (Linux; Android 10; BE2029) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.185 Mobile Safari/537.36",
        "expect": {
            "vendor": "OnePlus",
            "model": "BE2029",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Pad",
        "ua": "Mozilla/5.0 (Linux; U; Android 13; zh-CN; OPD2101 Build/TP1A.220905.001) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/100.0.4896.58 UCBrowser/16.3.9.1290 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "OPD2101",
            "type": "tablet"
        }
    },
    {
        "desc": "OPPO Neo",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.2.2; zh-cn; R831T Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 OppoBrowser/3.3.2 Mobile Safari/534.30",
        "expect": {
            "vendor": "OPPO",
            "model": "R831T",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO R7s",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; OPPO R7s Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 Chrome/37.0.0.0 MQQBrowser/7.1 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "R7s",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO A3s",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; CPH1803 Build/OPM1.171019.026; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "CPH1803",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO A12",
        "ua": "Mozilla/5.0 (Linux; Android 9; CPH2083) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "CPH2083",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Reno",
        "ua": "Mozilla/5.0 (Linux; Android 9; PCAT00 Build/PKQ1.190101.001; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "PCAT00",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Reno3 Pro 5G",
        "ua": "Mozilla/5.0 (Linux; Android 10; PCLM50) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.117 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "PCLM50",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Reno4 SE",
        "ua": "Mozilla/5.0 (Linux; U; Android 10; xx-xx; PEAM00 Build/QP1A.190711.020) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "PEAM00",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Reno4 5G",
        "ua": "Mozilla/5.0 (Linux; Android 10; PDPM00 Build/QKQ1.200216.002; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "PDPM00",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Reno4 Pro 5G",
        "ua": "Mozilla/5.0 (Linux; U; Android 10; xx-xx; PDNT00 Build/QKQ1.200216.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "PDNT00",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Reno5 A",
        "ua": "Mozilla/5.0 (Linux; Android 11; A101OP) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "A101OP",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Find X",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; PAFM00 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "PAFM00",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO Find 7a",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.3; xx-xx; X9007 Build/JLS36C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "expect": {
            "vendor": "OPPO",
            "model": "X9007",
            "type": "mobile"
        }
    },
    {
        "desc": "OPPO F5",
        "ua": "ozilla/5.0 (Linux; Android 7.1.1; CPH1723) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "OPPO",
            "model": "CPH1723",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme C1",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; RMX1811 Build/OPM1.171019.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.126 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX1811",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme C2",
        "ua": "Mozilla/5.0 (Linux; Android 9; RMX1941) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX1941",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme Narzo 20",
        "ua": "Mozilla/5.0 (Linux; U; Android 10; xx-xx; RMX2193 Build/QP1A.190711.020) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX2193",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme 2 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 9; RMX1801) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX1801",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme 3 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 11; RMX1851) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX1851",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme 8",
        "ua": "Mozilla/5.0 (Linux; Android 12; RMX3085) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX3085",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme 9 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 13; RMX3471) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX3471",
            "type": "mobile"
        }
    },
    {
        "desc": "Realme GT Master",
        "ua": "Mozilla/5.0 (Linux; Android 13; RMX3363) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Realme",
            "model": "RMX3363",
            "type": "mobile"
        }
    },
    {
        "desc": "Panasonic T31",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.2; Panasonic T31 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.170 Mobile Safari/537.36 ",
        "expect": {
            "vendor": "Panasonic",
            "model": "T31",
            "type": "mobile"
        }
    },
    {
        "desc": "Panasonic TX-32CSW514 SmartTV",
        "ua": "HbbTV/1.2.1 (;Panasonic;VIERA 2015;3.014;a001-003 4000-0000;)",
        "expect": {
            "vendor": "Panasonic",
            "model": "VIERA 2015",
            "type": "smarttv"
        }
    },
    {
        "desc": "Panasonic TX-40FXW724 SmartTV",
        "ua": "HbbTV/1.4.1 (+DRM;Panasonic;SmartTV2018mid;3.024;4301-0003 0002-0000;SmartTV2018;)",
        "expect": {
            "vendor": "Panasonic",
            "model": "SmartTV2018mid",
            "type": "smarttv"
        }
    },
    {
        "desc": "Panasonic TX-43HXW904 SmartTV",
        "ua": "HbbTV/1.5.1 (+DRM;Panasonic;SmartTV2020mid;3.326;4301-0003 0008-0000;com.panasonic.SmartTV2020mid;)",
        "expect": {
            "vendor": "Panasonic",
            "model": "SmartTV2020mid",
            "type": "smarttv"
        }
    },
    {
        "desc": "Panasonic DMR-HST130 SAT receiver",
        "ua": "HbbTV/1.1.1 (+PVR;Panasonic;DIGA WebKit M8658;3.420;;)",
        "expect": {
            "vendor": "Panasonic",
            "model": "DIGA WebKit M8658",
            "type": "smarttv"
        }
    },
    {
        "desc": "Philips SmartTV",
        "ua": "Opera/9.80 HbbTV/1.1.1 (; Philips; ; ; ; ) NETTV/4.0.2; en) Version/11.60",
        "expect": {
            "vendor": "Philips",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Philips 32PFL6606K/02 SmartTV (2011)",
        "ua": "Opera/9.80 (Linux mips ; U; HbbTV/1.1.1 (; Philips; ; ; ; ) CE-HTML/1.0 NETTV/3.1.0; en) Presto/2.6.33 Version/10.70",
        "expect": {
            "vendor": "Philips",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Philips 32PFL6606K/02 SmartTV (2013)",
        "ua": "Opera/9.80 (Linux mips ; U; HbbTV/1.1.1 (; Philips; ; ; ; ) CE-HTML/1.0 NETTV/3.1.0; en) Presto/2.6.33 Version/10.70",
        "expect": {
            "vendor": "Philips",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Philips 32PHS5301/12 SmartTV (2016)",
        "ua": "Mozilla/5.0 (Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36 OPR/29.0.1803.0 OMI/4.5.23.37.MOT2.13 HbbTV/1.2.1 (;Philips;32PHS5301/12;;_TV_MT5800;) Firmware/TPM161E_012.002.045.001 en",
        "expect": {
            "vendor": "Philips",
            "model": "32PHS5301/12",
            "type": "smarttv"
        }
    },
    {
        "desc": "Pico 4",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; PICO 4 OS5.8.2 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.38 Chrome/105.0.5195.68 VR Safari/537.36",
        "expect": {
            "vendor": "PICO",
            "model": "4",
            "type": "xr"
        }
    },
    {
        "desc": "Pico 4",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; PICO 4 OS5.4.0 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.22 Chrome/105.0.5195.68 VR Safari/537.36 OculusBrowser/7.0",
        "expect": {
            "vendor": "PICO",
            "model": "4",
            "type": "xr"
        }
    },
    {
        "desc": "Pico Neo3 Link",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; Pico Neo3 Link OS5.8.4.0 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.22 Chrome/105.0.5195.68 VR Safari/537.36",
        "expect": {
            "vendor": "Pico",
            "model": "Neo3 Link",
            "type": "xr"
        }
    },
    {
        "desc": "Roku",
        "ua": "Mozilla/5.0 (Roku) AppleWebKit/537.36 (KHTML, like Gecko) Web/1.1 Safari/537.36",
        "expect": {
            "vendor": "Roku",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Roku",
        "ua": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36 Roku/DVP-8.10 (468.10E04145A)",
        "expect": {
            "vendor": "Roku",
            "model": "DVP-8.10",
            "type": "smarttv"
        }
    },
    {
        "desc": "Roku",
        "ua": "Roku4640X/DVP-7.70 (297.70E04154A)",
        "expect": {
            "vendor": "Roku",
            "model": "DVP-7.70",
            "type": "smarttv"
        }
    },
    {
        "desc": "Xiaomi TV",
        "ua": "Mozilla/5.0 (Linux; Android 10; MiTV-MOOQ0 Build/QTG3.200305.006; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/94.0.4606.61 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MiTV-MOOQ0",
            "type": "smarttv"
        }
    },
    {
        "desc": "Kindle Fire HD",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; KFTT Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.4 Mobile Safari/535.19 Silk-Accelerated=true",
        "expect": {
            "vendor": "Amazon",
            "model": "KFTT",
            "type": "tablet"
        }
    },
    {
        "desc": "Kindle Fire HD",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; KFTT) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.4 Mobile Safari/535.19 Silk-Accelerated=true",
        "expect": {
            "vendor": "Amazon",
            "model": "KFTT",
            "type": "tablet"
        }
    },
    {
        "desc": "Echo Show 5",
        "ua": "Mozilla/5.0 (Linux; Android 5.1; AEORK Build/LVY48F; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "AEORK",
            "type": "tablet"
        }
    },
    {
        "desc": "Echo Show 8",
        "ua": "Mozilla/5.0 (Linux; Android 7.1; AEOCH) AppleWebKit/537.36 (KHTML, like Gecko) Silk/77.2.21 like Chrome/77.0.3865.92 Mobile Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "AEOCH",
            "type": "tablet"
        }
    },
    {
        "desc": "Echo Show 8",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.2; AEOCW) AppleWebKit/537.36 (KHTML, like Gecko) Silk/106.3.3 like Chrome/106.0.5249.170 Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "AEOCW",
            "type": "tablet"
        }
    },
    {
        "desc": "Echo Show 15",
        "ua": "Mozilla/5.0 (Linux; Android 9; AEOHY) AppleWebKit/537.36 (KHTML, like Gecko) Silk/112.6.3 like Chrome/112.0.5615.213 Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "AEOHY",
            "type": "tablet"
        }
    },
    {
        "desc": "Echo Dot",
        "ua": "Dalvik/2.1.0 (Linux; U; Android 5.1.1; AEOBC Build/LVY48F)",
        "expect": {
            "vendor": "Amazon",
            "model": "AEOBC",
            "type": "embedded"
        }
    },
    {
        "desc": "Samsung Galaxy A21s",
        "ua": "Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-A217F) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/11.0 Chrome/75.0.3770.143 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-A217F",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy A31",
        "ua": "Mozilla/5.0 (Linux; Android 10; SM-A315G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-A315G",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy A50",
        "ua": "Mozilla/5.0 (Linux; Android 9; SM-A505F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.105 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-A505F",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy A50s",
        "ua": "Mozilla/5.0 (Linux; Android 11; SM-A507FN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-A507FN",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy A52s",
        "ua": "Mozilla/5.0 (Linux; Android 13; SM-A528B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-A528B",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy A80",
        "ua": "Mozilla/5.0 (Linux; Android 9; SM-A805F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.112 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-A805F",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy Fold",
        "ua": "Mozilla/5.0 (Linux; Android 9; SAMSUNG SM-F900U Build/PPR1.180610.011) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/9.2 Chrome/67.0.3396.87 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-F900U",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy Z Flip",
        "ua": "Mozilla/5.0 (Linux; Android 10; SM-F700N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-F700N",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy Z Fold2",
        "ua": "Mozilla/5.0 (Linux; Android 10; SM-F916B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-F916B",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy S10E",
        "ua": "Mozilla/5.0 (Linux; Android 9; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-G970F",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy S20 5G",
        "ua": "Mozilla/5.0 (Linux; Android 10; SCG01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SCG01",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy Note 10+",
        "ua": "Mozilla/5.0 (Linux; Android 9; SM-N976V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.89 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-N976V",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung SM-C5000",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; SM-C5000 Build/MMB29M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 wkbrowser 4.1.35 3065",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-C5000",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung C8",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.1; SM-C7108) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-C7108",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy Note 8",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.2; GT-N5100 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "GT-N5100",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung SM-T231",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; SM-T231 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-T231",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Tab 6 Lite",
        "ua": "Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-P610) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/12.0 Chrome/79.0.3945.136 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-P610",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Tab A 9.7",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.1; SM-P550 Build/NMF26X; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.90 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-P550",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Tab A 10.1",
        "ua": " Mozilla/5.0 (Linux; Android 10; SAMSUNG SM-T515) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/13.0 Chrome/83.0.4103.106 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-T515",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Tab S7",
        "ua": "Mozilla/5.0 (Linux; Android 10; SM-T870) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-T870",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Tab S8",
        "ua": "Mozilla/5.0 (Linux; Android 12; SM-X706B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-X706B",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Tab S",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; SM-T700 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-T700",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Tab Pro 10.1",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; SM-T520 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-T520",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung Galaxy Watch",
        "ua": "Mozilla/5.0 (Linux; Tizen 5.5; SAMSUNG SM-R805W) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/2.0 Chrome/69.0.3497.106 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-R805W",
            "type": "wearable"
        }
    },
    {
        "desc": "Samsung Galaxy Watch Active 2",
        "ua": "Mozilla/5.0 (Linux; Tizen 5.5; SAMSUNG SM-R820) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/2.0 Chrome/69.0.3497.106 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-R820",
            "type": "wearable"
        }
    },
    {
        "desc": "Samsung Galaxy Watch4",
        "ua": "Mozilla/5.0 (Linux; Android 11; SAMSUNG SM-R875U) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/2.2. Chrome/102.0.5005.125 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-R875U",
            "type": "wearable"
        }
    },
    {
        "desc": "Samsung Galaxy Watch5 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 11; SAMSUNG SM-R925U) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/3.2. Chrome/111.0.5563.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-R925U",
            "type": "wearable"
        }
    },
    {
        "desc": "Samsung Note 10.1",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; SM-P605) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-P605",
            "type": "tablet"
        }
    },
    {
        "desc": "Samsung SmartTV2011",
        "ua": "HbbTV/1.1.1 (;;;;;) Maple;2011",
        "expect": {
            "vendor": "Samsung",
            "model": "SmartTV2011",
            "type": "smarttv"
        }
    },
    {
        "desc": "Samsung SmartTV2012",
        "ua": "HbbTV/1.1.1 (;Samsung;SmartTV2012;;;) WebKit",
        "expect": {
            "vendor": "Samsung",
            "model": "SmartTV2012",
            "type": "smarttv"
        }
    },
    {
        "desc": "Samsung SmartTV2014",
        "ua": "HbbTV/1.1.1 (;Samsung;SmartTV2014;T-NT14UDEUC-1060.4;;) WebKit",
        "expect": {
            "vendor": "Samsung",
            "model": "SmartTV2014",
            "type": "smarttv"
        }
    },
    {
        "desc": "Samsung SmartTV",
        "ua": "Mozilla/5.0 (SMART-TV; X11; Linux armv7l) AppleWebkit/537.42 (KHTML, like Gecko) Safari/537.42",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Samsung SmartTV",
        "ua": "Mozilla/5.0 (SMART-TV; Linux; Tizen 2.3) AppleWebkit/538.1 (KHTML, like Gecko) SamsungBrowser/1.0 TV Safari/538.1",
        "expect": {
            "vendor": "Samsung",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Samsung SmartTV HBBTV",
        "ua": "HbbTV/1.5.1 (+DRM;Samsung;SmartTV2021:UAU7000;T-KSU2EDEUC-1506.0;KantSU2e;urn:samsungtv:familyname:21_KANTSU2E_UHD_BASIC:2021;) Tizen/6.0 (+TVPLUS+SmartHubLink) Chrome/76 LaTivu_1.0.1_2021 RVID/17",
        "expect": {
            "vendor": "Samsung",
            "model": "SmartTV2021:UAU7000",
            "type": "smarttv"
        }
    },
    {
        "desc": "Sharp AQUOS-TVX19B",
        "ua": "Mozilla/5.0 (Linux; Android 9; AQUOS-TVX19B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sharp",
            "model": "AQUOS-TVX19B",
            "type": "smarttv"
        }
    },
    {
        "desc": "Sharp Aquos B10",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; SH-A01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sharp",
            "model": "SH-A01",
            "type": "mobile"
        }
    },
    {
        "desc": "Sharp Aquos L2",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; SH-L02 Build/S4045) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sharp",
            "model": "SH-L02",
            "type": "mobile"
        }
    },
    {
        "desc": "Sharp Aquos L2",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; SH-L02) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sharp",
            "model": "SH-L02",
            "type": "mobile"
        }
    },
    {
        "desc": "Sharp Aquos R2",
        "ua": "Mozilla/5.0 (Linux; Android 8.0; SHV42) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.92 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sharp",
            "model": "SHV42",
            "type": "mobile"
        }
    },
    {
        "desc": "SONY Xperia 1 III",
        "ua": "Mozilla/5.0 (Linux; Android 11; A101SO) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "A101SO",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony G8141 (Xperia XZ1)",
        "ua": "Mozilla/5.0 (Linux; Android 9; SO-01K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "SO-01K",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony G8141 (Xperia XZ Premium)",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; G8141) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "G8141",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony C5303 (Xperia SP)",
        "ua": "Mozilla/5.0 (Linux; Android 4.3; C5303 Build/12.1.A.1.205) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.93 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "C5303",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony SO-02F (Xperia Z1 F)",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.2; SO-02F Build/14.1.H.2.119) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "SO-02F",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony D6653 (Xperia Z3)",
        "ua": "Mozilla/5.0 (Linux; Android 4.4; D6653 Build/23.0.A.0.376) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "D6653",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony Xperia SOL25 (ZL2)",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.4; SOL25 Build/17.1.1.C.1.64) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "expect": {
            "vendor": "Sony",
            "model": "SOL25",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony Xperia SP",
        "ua": "Mozilla/5.0 (Linux; Android 4.3; C5302 Build/12.1.A.1.201) AppleWebkit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "C5302",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony Xperia L4",
        "ua": "Mozilla/5.0 (Linux; Android 9; XQ-AD51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.83 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "XQ-AD51",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony Xperia 1ii",
        "ua": "Mozilla/5.0 (Linux; Android 10; XQ-AT51) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "XQ-AT51",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony Xperia 1ii",
        "ua": "Mozilla/5.0 (Linux; Android 10; SOG01) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "SOG01",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony Xperia 10ii",
        "ua": "Mozilla/5.0 (Linux; Android 10; XQ-AU52) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "XQ-AU52",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony Xperia Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; XQ-AQ52) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.185 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "XQ-AQ52",
            "type": "mobile"
        }
    },
    {
        "desc": "Sony SGP521 (Xperia Z2 Tablet)",
        "ua": "Mozilla/5.0 (Linux; Android 4.4; SGP521 Build/17.1.A.0.432) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "Xperia Tablet",
            "type": "tablet"
        }
    },
    {
        "desc": "Sony Xperia Z2 Tablet",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; SGP561) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.99 Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "Xperia Tablet",
            "type": "tablet"
        }
    },
    {
        "desc": "Sony Tablet S",
        "ua": "Mozilla/5.0 (Linux; U; Android 3.1; Sony Tablet S Build/THMAS10000) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
        "expect": {
            "vendor": "Sony",
            "model": "Xperia Tablet",
            "type": "tablet"
        }
    },
    {
        "desc": "Sony Tablet Z LTE",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.1; SonySGP321 Build/10.2.C.0.143) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
        "expect": {
            "vendor": "Sony",
            "model": "Xperia Tablet",
            "type": "tablet"
        }
    },
    {
        "desc": "Sony BRAVIA 4K GB ATV3",
        "ua": "Mozilla/5.0 (Linux; Andr0id 9; BRAVIA 4K GB ATV3 Build/PTT1.190515.001.S38) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36 OPR/46.0.2207.0 OMI/4.13.0.180.DIA5.104 Model/Sony-BRAVIA-4K-GB-ATV3",
        "expect": {
            "vendor": "Sony",
            "model": "BRAVIA 4K GB ATV3",
            "type": "smarttv"
        }
    },
    {
        "desc": "Sony BRAVIA 4K GB ATV3",
        "ua": "Mozilla/5.0 (Linux; Android 9; BRAVIA 4K GB ATV3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sony",
            "model": "BRAVIA 4K GB ATV3",
            "type": "smarttv"
        }
    },
    {
        "desc": "Sony Bravia 4k UR2",
        "ua": "Mozilla/5.0 (Linux: Andr0id 9: BRAVIA 4K UR2 Build/PTT1.190515.001.S104) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36 OPR/46.0.2207.0 OMI/4.13.5.431.DIA5HBBTV.250 Model/Sony-BRAVIA-4K-UR2",
        "expect": {
            "vendor": "Sony",
            "model": "BRAVIA 4K UR2",
            "type": "smarttv"
        }
    },
    {
        "desc": "TCL 10 5G",
        "ua": "Mozilla/5.0 (Linux; Android 11; T790Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36 EdgA/114.0.1823.43",
        "expect": {
            "vendor": "TCL",
            "model": "T790Y",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 10 5G UW",
        "ua": "Mozilla/5.0 (Linux; Android 10; T790S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T790S",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 10 Plus",
        "ua": "Mozilla/5.0 (Linux; Android 11; T782H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Mobile Safari/537.36 OPR/64.3.3282.60839",
        "expect": {
            "vendor": "TCL",
            "model": "T782H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 10 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; T799B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.87 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T799B",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 10 SE",
        "ua": "Mozilla/5.0 (Linux; Android 10; T766H_RU) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.85 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T766H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 10 TabMax",
        "ua": "Mozilla/5.0 (Linux; Android 11; 9296Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9296Q",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL 10 TabMax 4G",
        "ua": "Mozilla/5.0 (Linux; Android 10; 9295G_EEA) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9295G",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL 10 TabMax WiFi",
        "ua": "Mozilla/5.0 (Linux; Android 10; 9296G_TR) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9296G",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL 10L",
        "ua": "Mozilla/5.0 (Linux; Android 10; T770B Build/QKQ1.200329.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.141 Mobile Safari/537.36 GSA/11.41.10.23.arm64",
        "expect": {
            "vendor": "TCL",
            "model": "T770B",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 10L",
        "ua": "Mozilla/5.0 (Linux; Android 11; T770H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T770H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20 5G",
        "ua": "Mozilla/5.0 (Linux; Android 11; T781) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T781",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20 Pro 5G",
        "ua": "Mozilla/5.0 (Linux; Android 12; T810S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Mobile Safari/537.36 EdgA/113.0.1774.63",
        "expect": {
            "vendor": "TCL",
            "model": "T810S",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20 SE",
        "ua": "Mozilla/5.0 (Linux; Android 11; T671H Build/RKQ1.201112.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/99.0.4844.73 Mobile Safari/537.36 GoogleApp/13.9.7.23.arm64",
        "expect": {
            "vendor": "TCL",
            "model": "T671H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20 XE",
        "ua": "Mozilla/5.0 (Linux; Android 11; 5087Z) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "5087Z",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20B",
        "ua": "Mozilla/5.0 (Linux; Android 11; 6159K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "6159K",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 205",
        "ua": "Mozilla/5.0 (Linux; Android 11; 4187D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "4187D",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20E",
        "ua": "Mozilla/5.0 (Linux; Android 11; 6125A) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/18.0 Chrome/99.0.4844.88 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "6125A",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20L",
        "ua": "Mozilla/5.0 (Linux; Android 11; T774H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.59 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T774H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20L Plus",
        "ua": "Mozilla/5.0 (Linux; Android 11; T775H Build/RKQ1.210107.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/114.0.5735.61 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T775H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20R 5G",
        "ua": "Mozilla/5.0 (Linux; Android 11; T767H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.97 Mobile Safari/537.36 OPR/71.3.3718.67322",
        "expect": {
            "vendor": "TCL",
            "model": "T767H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20S",
        "ua": "Mozilla/5.0 (Linux; Android 11; T773O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T773O",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 20Y",
        "ua": "Mozilla/5.0 (Linux; Android 11; 6156D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.134 Mobile Safari/537.36 OPR/70.3.3653.66287",
        "expect": {
            "vendor": "TCL",
            "model": "6156D",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30 V 5G",
        "ua": "Mozilla/5.0 (Linux; Android 11; T781S Build/RKQ1.210614.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/115.0.5790.166 Mobile Safari/537.36[FBAN/EMA;FBLC/en_US;FBAV/369.0.0.5.110;]",
        "expect": {
            "vendor": "TCL",
            "model": "T781S",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30 XE 5G",
        "ua": "Mozilla/5.0 (Linux; Android 12; T767W Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/113.0.5672.162 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/416.0.0.35.85;]",
        "expect": {
            "vendor": "TCL",
            "model": "T767W",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 305",
        "ua": "Mozilla/5.0 (Linux; arm; Android 11; 6102D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.167 YaBrowser/22.7.6.96.00 SA/3 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "6102D",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 306",
        "ua": "Mozilla/5.0 (Linux; Android 12; 6102H Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/107.0.5304.141 Mobile Safari/537.36[FBAN/EMA;FBLC/it_IT;FBAV/332.0.0.22.108;]",
        "expect": {
            "vendor": "TCL",
            "model": "6102H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30",
        "ua": "Mozilla/5.0 (Linux; Android 12; T676H Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/113.0.5672.162 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T676H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30+",
        "ua": "Mozilla/5.0 (Linux; Android 12; T676J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T676J",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30 5G",
        "ua": "Mozilla/5.0 (Linux; Android 12; T776H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.104 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T776H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30 LE",
        "ua": "Mozilla/5.0 (Linux; Android 12; 4188V Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/112.0.5615.136 Mobile Safari/537.36[FBAN/EMA;FBLC/en_US;FBAV/352.0.0.14.108;]",
        "expect": {
            "vendor": "TCL",
            "model": "4188V",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30 SE",
        "ua": "Mozilla/5.0 (Linux; Android 12; 6165H Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/108.0.5359.128 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/396.1.0.28.104;]",
        "expect": {
            "vendor": "TCL",
            "model": "6165H",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 30E",
        "ua": "Mozilla/5.0 (Linux; Android 12; 6127I) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "6127I",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 40 NxtPaper",
        "ua": "Mozilla/5.0 (Linux; Android 13; T612B Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/125.0.6422.53 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T612B",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL A3",
        "ua": "Mozilla/5.0 (Linux; Android 11; A509DL Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/101.0.4951.61 Mobile Safari/537.36 GSA/13.18.7.23.arm64",
        "expect": {
            "vendor": "TCL",
            "model": "A509DL",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL A30",
        "ua": "Mozilla/5.0 (Linux; Android 11; 5102L Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/112.0.5615.136 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/413.0.0.30.104;]",
        "expect": {
            "vendor": "TCL",
            "model": "5102L",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 40 SE",
        "ua": "Mozilla/5.0 (Linux; Android 13; T610K Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/115.0.5790.166 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T610K",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 40 XE 5G",
        "ua": "Mozilla/5.0 (Linux; Android 13; T609DL Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/125.0.6422.136 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/466.1.0.57.85;]",
        "expect": {
            "vendor": "TCL",
            "model": "T609DL",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 403",
        "ua": "Mozilla/5.0 (Linux; Android 12; T431D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T431D",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 405",
        "ua": "Mozilla/5.0 (Linux; Android 12; T506D Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/113.0.5672.162 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/418.0.0.33.69;]",
        "expect": {
            "vendor": "TCL",
            "model": "T506D",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 408",
        "ua": "Mozilla/5.0 (Linux; U; Android 12; T507U Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/105.0.5195.136 Mobile Safari/537.36 OPR/75.0.2254.68857",
        "expect": {
            "vendor": "TCL",
            "model": "T507U",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL 40R 5G",
        "ua": "Mozilla/5.0 (Linux; Android 12; T771K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36 EdgA/114.0.1823.37",
        "expect": {
            "vendor": "TCL",
            "model": "T771K",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL Ion X",
        "ua": "Mozilla/5.0 (Linux; Android 12; T430W Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/114.0.5735.60 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T430W",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL NxtPaper 11",
        "ua": "Mozilla/5.0 (Linux; Android 13; 9466X Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/126.0.6478.179 Safari/537.36 [FB_IAB/FB4A;FBAV/473.0.0.41.81;]",
        "expect": {
            "vendor": "TCL",
            "model": "9466X",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Stylus 5G",
        "ua": "Mozilla/5.0 (Linux; Android 12; T779W Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/127.0.6533.2 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "T779W",
            "type": "mobile"
        }
    },
    {
        "desc": "TCL Tab 8 4G",
        "ua": "Mozilla/5.0 (Linux; Android 10; 9048S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9048S",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Tab 8 LE",
        "ua": "Mozilla/5.0 (Linux; Android 12; 9137W Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/114.0.5735.61 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9137W",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Tab 10 FHD 4G",
        "ua": "Mozilla/5.0 (Linux; Android 11; 9060G Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/114.0.5735.196 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9060G",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Tab 10 HD 4G",
        "ua": "Mozilla/5.0 (Linux; Android 11; 9060X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9060X",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Tab 10 LTE",
        "ua": "Mozilla/5.0 (Linux; Android 13; 8196G Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/126.0.6478.162 Safari/537.36 [FB_IAB/FB4A;FBAV/471.0.0.35.80;]",
        "expect": {
            "vendor": "TCL",
            "model": "8196G",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Tab 10 WiFi",
        "ua": "Mozilla/5.0 (Linux; Android 13; 8496G Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/127.0.6533.61 Safari/537.36 [FB_IAB/FB4A;FBAV/474.0.0.52.74;]",
        "expect": {
            "vendor": "TCL",
            "model": "8496G",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Tab 10L",
        "ua": "Mozilla/5.0 (Linux; Android 11; 8491X_EEA Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/107.0.5304.105 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "8491X",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Tab 10s 4G",
        "ua": "Mozilla/5.0 (Linux; Android 11; 9080G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "9080G",
            "type": "tablet"
        }
    },
    {
        "desc": "TCL Xess P17AA",
        "ua": "Mozilla/5.0 (Linux; Android 5.1; TCL Xess P17AA Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.85 Safari/537.36",
        "expect": {
            "vendor": "TCL",
            "model": "Xess P17AA",
            "type": "tablet"
        }
    },
    {
        "desc": "Tecno KC8",
        "ua": "Mozilla/5.0 (Linux; Android 10; TECNO KC8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TECNO",
            "model": "KC8",
            "type": "mobile"
        }
    },
    {
        "desc": "Tecno Spark 8C",
        "ua": "Mozilla/5.0 (Linux; Android 11; TECNO KG5n) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "TECNO",
            "model": "KG5n",
            "type": "mobile"
        }
    },
    {
        "desc": "Tesla",
        "ua": "Mozilla/5.0 (X11; GNU/Linux) AppleWebKit/601.1 (KHTML, like Gecko) Tesla QtCarBrowser Safari/601.1",
        "expect": {
            "vendor": "Tesla",
            "model": "",
            "type": "embedded"
        }
    },
    {
        "desc": "Tesla",
        "ua": "Mozilla/5.0 (X11; GNU/Linux) AppleWebKit/537.36 (KHTML, like Gecko) Chromium/79.0.3945.130 Chrome/79.0.3945.130 Safari/537.36 Tesla/2020.16.2.1-e99c70fff409",
        "expect": {
            "vendor": "Tesla",
            "model": "",
            "type": "embedded"
        }
    },
    {
        "desc": "TechniSAT Digit ISIO S SAT receiver",
        "ua": "Opera/9.80 (Linux sh4; U; HbbTV/1.1.1 (;;;;;); CE-HTML; TechniSat Digit ISIO S; de) Presto/2.9.167 Version/11.50",
        "expect": {
            "vendor": "TechniSat",
            "model": "Digit ISIO S",
            "type": "smarttv"
        }
    },
    {
        "desc": "TechniSAT MultyVision SmartTV",
        "ua": "Opera/9.80 (Linux i686; U; HbbTV/1.1.1 (;;;;;); CE-HTML; TechniSat MultyVision ISIO; de) Presto/2.9.167 Version/11.50",
        "expect": {
            "vendor": "TechniSat",
            "model": "MultyVision ISIO",
            "type": "smarttv"
        }
    },
    {
        "desc": "Ulefone Armor",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; Armor Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.107 Mobile Safari/537.36",
        "expect": {
            "vendor": "Ulefone",
            "model": "Armor",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Armor",
        "ua": "Mozilla/5.0 (Linux; arm_64; Android 6.0; Armor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 YaBrowser/20.4.2.101.00 SA/1 Mobile Safari/537.36",
        "expect": {
            "vendor": "Ulefone",
            "model": "Armor",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Armor 8 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 11; Armor 8 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.192 Mobile Safari/537.36 OPR/74.1.3922.71199",
        "expect": {
            "vendor": "Ulefone",
            "model": "Armor 8 Pro",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Armor 12 5G",
        "ua": "Mozilla/5.0 (Linux; Android 11; Armor 12 5G Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/115.0.5790.166 Mobile Safari/537.36",
        "expect": {
            "vendor": "Ulefone",
            "model": "Armor 12 5G",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Armor 20WT",
        "ua": "Mozilla/5.0 (Linux; Android 12; Armor 20WT) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/22.0 Chrome/111.0.5563.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Ulefone",
            "model": "Armor 20WT",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Armor Pad",
        "ua": "Mozilla/5.0 (Linux; Android 12; Armor Pad Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/116.0.0.0 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/431.0.0.30.108;]",
        "expect": {
            "vendor": "Ulefone",
            "model": "Armor Pad",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Armor X5 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; Armor X5 Pro Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/116.0.0.0 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/430.0.0.23.113;]",
        "expect": {
            "vendor": "Ulefone",
            "model": "Armor X5 Pro",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Power Armor 14 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 12; Power Armor14 Pro Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/115.0.5790.138 Mobile Safari/537.36",
        "expect": {
            "vendor": "Ulefone",
            "model": "Power Armor14 Pro",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Power Armor 18T",
        "ua": "Mozilla/5.0 (Linux; Android 12; Power Armor 18T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Ulefone",
            "model": "Power Armor 18T",
            "type": "mobile"
        }
    },
    {
        "desc": "Ulefone Power Armor 19T",
        "ua": "Mozilla/5.0 (Linux; Android 12; Power Armor 19T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.192 Mobile Safari/537.36 OPR/74.3.3922.71982",
        "expect": {
            "vendor": "Ulefone",
            "model": "Power Armor 19T",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi 2201117TG",
        "ua": "Mozilla/5.0 (Linux; Android 11; 2201117TG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.98 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "2201117TG",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi M2004J19C",
        "ua": "Mozilla/5.0 (Linux; Android 11; M2004J19C Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/113.0.5672.77 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2004J19C",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi M2006C3MNG",
        "ua": "Mozilla/5.0 (Linux; Android 11; M2006C3MNG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.85 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2006C3MNG",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi 21061119DG",
        "ua": "Mozilla/5.0 (Linux; arm_64; Android 11; 21061119DG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.3.7.24.00 SA/3 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "21061119DG",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi 2013023",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.2.2; en-US; 2013023 Build/HM2013023) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/10.0.1.512 U3/0.8.0 Mobile Safari/533.1",
        "expect": {
            "vendor": "Xiaomi",
            "model": "2013023",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Hongmi Note 1W",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.2.2; zh-CN; HM NOTE 1W Build/JDQ39) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.7.9.439 U3/0.8.0 Mobile Safari/533.1",
        "expect": {
            "vendor": "Xiaomi",
            "model": "HM NOTE 1W",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi 3C",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.3; zh-CN; MI 3C Build/JLS36C) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.7.9.439 U3/0.8.0 Mobile Safari/533.1",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI 3C",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi 5",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; MI 5 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.83 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI 5",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi 6",
        "ua": "Mozilla/5.0 (Linux; Android 7.1; MI 6 Build/NMF26X; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI 6",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi 10 Pro",
        "ua": "Linux; U; Android 13; Mi 10 Pro Build/TKQ1.221114.001",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Mi 10 Pro",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi 5s Plus",
        "ua": "Mozilla/5.0 (Linux; U; Android 6.0.1; zh-cn; MI 5s Plus Build/MXB48T) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.146 Mobile Safari/537.36 XiaoMi/MiuiBrowser/8.7.1",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI 5s Plus",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi A1",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; Mi A1 Build/OPR1.170623.026) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Mi A1",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi Note",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.4; MI NOTE LTE Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.76 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI NOTE LTE",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi One Plus",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; MI-ONE Plus Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI-ONE Plus",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi Max 3",
        "ua": "Mozilla/5.0 (Linux; Android 9; MI MAX 3 Build/PKQ1.181007.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI MAX 3",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi A1",
        "ua": "Mozilla/5.0 (Linux; Android 9; Mi A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.101 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Mi A1",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi A2 Lite",
        "ua": "Mozilla/5.0 (Linux; Android 9; Mi A2 Lite) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.62 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Mi A2 Lite",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi 9 SE",
        "ua": "Mozilla/5.0 (Linux; Android 9; Mi 9 SE) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Mi 9 SE",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi A2",
        "ua": "Mozilla/5.0 (Linux; Android 9; Mi A2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Mi A2",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Mi CC9",
        "ua": "Mozilla/5.0 (Linux; U; Android 11; zh-cn; MI CC 9 Build/RKQ1.200826.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.116 Mobile Safari/537.36 XiaoMi/MiuiBrowser/15.5.18",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI CC 9",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi MI PAD 2",
        "ua": "Mozilla/5.0 (Linux; Android 5.1; MI PAD 2 Build/LMY47I; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/60.0.3112.107 Safari/537.36 [FB_IAB/FB4A;FBAV/137.0.0.24.91;]",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI PAD 2",
            "type": "tablet"
        }
    },
    {
        "desc": "Xiaomi MI PAD 4 PLUS",
        "ua": "Mozilla/5.0 (Linux; Android 8.1; MI PAD 4 PLUS) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "MI PAD 4 PLUS",
            "type": "tablet"
        }
    },
    {
        "desc": "Xiaomi POCO X2",
        "ua": "Mozilla/5.0 (Linux; Android 10; POCO X2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "POCO X2",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi POCO X3 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 11; M2102J20SI) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2102J20SI",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi POCO X3 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 12; M2102J20SG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2102J20SG",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi POCO X3 NFC",
        "ua": "Mozilla/5.0 (Linux; Android 12; M2007J20CG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2007J20CG",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi POCO M2 Pro",
        "ua": "Mozilla/5.0 (Linux; arm_64; Android 11; POCO M2 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 YaBrowser/22.11.7.42.00 SA/3 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "POCO M2 Pro",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi POCO M3",
        "ua": "Mozilla/5.0 (Linux; Android 10; M2010J19CI) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2010J19CI",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi POCOPHONE F1",
        "ua": "Mozilla/5.0 (Linux; Android 10; POCOPHONE F1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "POCOPHONE F1",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Redmi 4A",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; Redmi 4A Build/MMB29M; xx-xx) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/56.0.2924.87 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Redmi 4A",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Redmi 10C",
        "ua": "Mozilla/5.0 (Linux; Android 12; 220333QAG) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "220333QAG",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Redmi K30 5G",
        "ua": "Mozilla/5.0 (Linux; Android 10; Redmi K30 5G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Redmi K30 5G",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Redmi K30 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 10; Redmi K30 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Redmi K30 Pro",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Redmi Note 3",
        "ua": "Mozilla/5.0 (Linux; Android 6.0.1; Redmi Note 3 Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.116 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Redmi Note 3",
            "type": "mobile"
        }
    },
    {
        "desc": "Xiaomi Redmi Note 9 Pro Max",
        "ua": "Mozilla/5.0 (Linux; Android 10; Redmi Note 9 Pro Max) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.99 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Redmi Note 9 Pro Max",
            "type": "mobile"
        }
    },
    {
        "desc": "XiaoMi Redmi Note 9S",
        "ua": "Mozilla/5.0 (Linux; Android 10; Redmi Note 9S) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Redmi Note 9S",
            "type": "mobile"
        }
    },
    {
        "desc": "XiaoMi Redmi Note 10 5G",
        "ua": "Mozilla/5.0 (Linux; Android 12; M2103K19C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.88 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2103K19C",
            "type": "mobile"
        }
    },
    {
        "desc": "XiaoMi Redmi Note 10 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 13; M2101K6P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2101K6P",
            "type": "mobile"
        }
    },
    {
        "desc": "XiaoMi Redmi Note 10 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 12; M2101K6G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "M2101K6G",
            "type": "mobile"
        }
    },
    {
        "desc": "XiaoMi Redmi Note 8",
        "ua": "Mozilla/5.0 (Linux; Android 10; Redmi Note 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Xiaomi",
            "model": "Redmi Note 8",
            "type": "mobile"
        }
    },
    {
        "desc": "XiaoMi Redmi Note 12 Turbo",
        "ua": "Mozilla/5.0 (Linux; Android 13; 23049RAD8C; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.141 Mobile Safari/537.36 VivoBrowser/16.7.1.1",
        "expect": {
            "vendor": "Xiaomi",
            "model": "23049RAD8C",
            "type": "mobile"
        }
    },
    {
        "desc": "ZTE Blade A6",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.1; ZTE BLADE A0620 Build/NMF26F; ru-ru) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36 Puffin/9.2.0.50586AP",
        "expect": {
            "vendor": "ZTE",
            "model": "BLADE A0620",
            "type": "mobile"
        }
    },
    {
        "desc": "PlayStation 4",
        "ua": "Mozilla/5.0 (PlayStation 4 3.00) AppleWebKit/537.73 (KHTML, like Gecko)",
        "expect": {
            "vendor": "Sony",
            "model": "PlayStation 4",
            "type": "console"
        }
    },
    {
        "desc": "PlayStation 5",
        "ua": "Mozilla/5.0 (Playstation; Playstation 5/1.05) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Safari/605.1.15",
        "expect": {
            "vendor": "Sony",
            "model": "Playstation 5",
            "type": "console"
        }
    },
    {
        "desc": "PlayStation Vita",
        "ua": "Mozilla/5.0 (PlayStation Vita 3.52) AppleWebKit/537.73 (KHTML, like Gecko) Silk/3.2",
        "expect": {
            "vendor": "Sony",
            "model": "PlayStation Vita",
            "type": "console"
        }
    },
    {
        "desc": "Nintendo Switch",
        "ua": "Mozilla/5.0 (Nintendo Switch; WifiWebAuthApplet) AppleWebKit/606.4 (KHTML, like Gecko) NF/6.0.1.15.4 NintendoBrowser/5.1.0.20393",
        "expect": {
            "vendor": "Nintendo",
            "model": "Switch",
            "type": "console"
        }
    },
    {
        "desc": "Nintendo WiiU",
        "ua": "Mozilla/5.0 (Nintendo WiiU) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.4.2.9 NintendoBrowser/4.2.0.11146.EU",
        "expect": {
            "vendor": "Nintendo",
            "model": "WiiU",
            "type": "console"
        }
    },
    {
        "desc": "Nintendo Wii",
        "ua": "Opera/9.10 (Nintendo Wii; U; ; 1621; en)",
        "expect": {
            "vendor": "Nintendo",
            "model": "Wii",
            "type": "console"
        }
    },
    {
        "desc": "Nintendo 3DS",
        "ua": "Mozilla/5.0 (Nintendo 3DS; U; ; en) Version/1.7610.EU",
        "expect": {
            "vendor": "Nintendo",
            "model": "3DS",
            "type": "console"
        }
    },
    {
        "desc": "Nintendo 3DS",
        "ua": "Mozilla/5.0 (New Nintendo 3DS like iPhone) AppleWebKit/536.30 (KHTML, like Gecko) NX/3.0.0.5.15 Mobile NintendoBrowser/1.3.10126.EU",
        "expect": {
            "vendor": "Nintendo",
            "model": "3DS",
            "type": "console"
        }
    },
    {
        "desc": "Galaxy Nexus",
        "ua": "Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19",
        "expect": {
            "vendor": "Samsung",
            "model": "Galaxy Nexus",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy C9 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 6.0; SAMSUNG SM-C900F Build/MMB29M) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.2 Chrome/44.0.2403.133 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-C900F",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy S5",
        "ua": "Mozilla/5.0 (Linux; Android 5.0; SM-G900F Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.78 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-G900F",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy J7 Prime",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; SM-G610F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-G610F",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy S6",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; SM-G920I Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-G920I",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy S6 Edge",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; SM-G925I Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-G925I",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy Note 5 Chrome",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; SM-N920C Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.91 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-N920C",
            "type": "mobile"
        }
    },
    {
        "desc": "Samsung Galaxy Note 5 Samsung Browser",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; SAMSUNG SM-N920C Build/LMY47X) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36",
        "expect": {
            "vendor": "Samsung",
            "model": "SM-N920C",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Chromecast with Google TV",
        "ua": "Mozilla/5.0 (Linux; Android 12.0; Build/STTL.240206.002) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.0 Safari/537.36 CrKey/1.56.500000 DeviceType/AndroidTV",
        "expect": {
            "vendor": "Google",
            "model": "Chromecast AndroidTV",
            "type": "smarttv"
        }
    },
    {
        "desc": "Google Chromecast Mini Smart Speaker",
        "ua": "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.225 Safari/537.36 CrKey/1.56.500000 DeviceType/SmartSpeaker",
        "expect": {
            "vendor": "Google",
            "model": "Chromecast SmartSpeaker",
            "type": "smarttv"
        }
    },
    {
        "desc": "Google Chromecast Third Generation",
        "ua": "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.225 Safari/537.36 CrKey/1.56.500000 DeviceType/Chromecast",
        "expect": {
            "vendor": "Google",
            "model": "Chromecast Third Generation",
            "type": "smarttv"
        }
    },
    {
        "desc": "Google Chromecast Nest Hub",
        "ua": "Mozilla/5.0 (Fuchsia) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 CrKey/1.56.500000",
        "expect": {
            "vendor": "Google",
            "model": "Chromecast Nest Hub",
            "type": "smarttv"
        }
    },
    {
        "desc": "Google Chromecast",
        "ua": "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.84 Safari/537.36 CrKey/1.22.79313",
        "expect": {
            "vendor": "Google",
            "model": "Chromecast",
            "type": "smarttv"
        }
    },
    {
        "desc": "Google Pixel C",
        "ua": "Mozilla/5.0 (Linux; Android 7.0; Pixel C Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.98 Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel C",
            "type": "tablet"
        }
    },
    {
        "desc": "Google Pixel C",
        "ua": "Mozilla/5.0 (Linux; Android 8.0.0; Pixel C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel C",
            "type": "tablet"
        }
    },
    {
        "desc": "Google Pixel",
        "ua": "Mozilla/5.0 (Linux; Android 7.1; Pixel Build/NDE63V) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.85 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel XL",
        "ua": "Mozilla/5.0 (Linux; Android 7.1; Pixel XL Build/NDE63X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.85 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel XL",
        "ua": "Mozilla/5.0 (Linux; Android 9; Pixel XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 2",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; Pixel 2 Build/OPM1.171019.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 2",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 2 XL",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; Pixel 2 XL Build/OPM1.171019.013) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 2 XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 2 XL",
        "ua": "Mozilla/5.0 (Linux; Android 9; Pixel 2 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 2 XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 3",
        "ua": "Mozilla/5.0 (Linux; Android 9; Pixel 3 Build/PD1A.180720.030) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 3",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 3 XL",
        "ua": "Mozilla/5.0 (Linux; Android 9; Pixel 3 XL Build/PD1A.180720.030) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 3 XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 3 XL",
        "ua": "Mozilla/5.0 (Linux; Android 9; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 3 XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 3a",
        "ua": "Mozilla/5.0 (Linux; Android 10; Pixel 3a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 3a",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 3a XL",
        "ua": "Mozilla/5.0 (Linux; Android 10; Pixel 3a XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 3a XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 4",
        "ua": "Mozilla/5.0 (Linux; Android 10; Pixel 4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 4",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 4a",
        "ua": "Mozilla/5.0 (Linux; Android 10; Pixel 4a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.83 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 4a",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 4 XL",
        "ua": "Mozilla/5.0 (Linux; Android 10; Pixel 4 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 4 XL",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 5",
        "ua": "Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 5",
            "type": "mobile"
        }
    },
    {
        "desc": "Google Pixel 7",
        "ua": "Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Google",
            "model": "Pixel 7",
            "type": "mobile"
        }
    },
    {
        "desc": "Generic Android Device",
        "ua": "Mozilla/5.0 (Linux; U; Android 6.0.1; i980 Build/MRA58K)",
        "expect": {
            "vendor": "Generic",
            "model": "Android 6.0.1"
        }
    },
    {
        "desc": "Android Phone Unidentified Vendor (docomo F-04K)",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; F-04K Build/V15R060P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.137 Mobile Safari/537.36",
        "expect": {
            "model": "F-04K",
            "type": "mobile"
        }
    },
    {
        "desc": "docomo SH-02M",
        "ua": "Mozilla/5.0 (Linux; Android 9; SH-02M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Mobile Safari/537.36",
        "expect": {
            "vendor": "Sharp",
            "model": "SH-02M",
            "type": "mobile"
        }
    },
    {
        "desc": "Android Tablet Unidentified Vendor (docomo F-02K)",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; F-02K Build/V44R059G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.109 Safari/537.36",
        "expect": {
            "model": "F-02K",
            "type": "tablet"
        }
    },
    {
        "desc": "Android Tablet Unidentified Vendor (docomo d-02K)",
        "ua": "Mozilla/5.0 (Linux; Android 9; d-02K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.136 Safari/537.36",
        "expect": {
            "model": "d-02K",
            "type": "tablet"
        }
    },
    {
        "desc": "LG VK Series Tablet",
        "ua": "Mozilla/5.0 (Linux; Android 5.0.2; VK700 Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.84 Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "VK700",
            "type": "tablet"
        }
    },
    {
        "desc": "LG LK Series Tablet",
        "ua": "Mozilla/5.0 (Linux; Android 5.0.1; LGLK430 Build/LRX21Y) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/38.0.2125.102 Safari/537.36",
        "expect": {
            "vendor": "LG",
            "model": "LK430",
            "type": "tablet"
        }
    },
    {
        "desc": "Amazon Alexa Echo Show",
        "ua": "AlexaWebMediaPlayer/1.0.200641.0 (Linux;Android 5.1.1)",
        "expect": {
            "vendor": "Amazon",
            "model": "Alexa",
            "type": "tablet"
        }
    },
    {
        "desc": "Amazon Kindle Fire Tablet",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.4.3; en-us; KFSAWI Build/KTU84M) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.66 like Chrome/39.0.2171.93 Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "KFSAWI",
            "type": "tablet"
        }
    },
    {
        "desc": "Amazon Kindle Fire Tablet",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.4.3; en-us; KFSAWI) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.66 like Chrome/39.0.2171.93 Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "KFSAWI",
            "type": "tablet"
        }
    },
    {
        "desc": "Amazon Kindle Fire Tablet",
        "ua": "Mozilla/5.0 (Linux; Android 9; KFMAWI Build/PS7312; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "KFMAWI",
            "type": "tablet"
        }
    },
    {
        "desc": "Amazon Fire TV",
        "ua": "Mozilla/5.0 (Linux; Android 4.2.2; AFTB Build/JDQ39) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.173 Mobile Safari/537.22",
        "expect": {
            "vendor": "Amazon",
            "model": "B",
            "type": "smarttv"
        }
    },
    {
        "desc": "Amazon Fire TV",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; AFTT) AppleWebKit/537.36 (KHTML, like Gecko) Silk/86.3.20 like Chrome/86.0.4240.198 Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "T",
            "type": "smarttv"
        }
    },
    {
        "desc": "Amazon Fire TV",
        "ua": "Mozilla/5.0 (Linux; Android 9; AFTKA Build/PS7633.3445N; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/108.0.5359.160 Mobile Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "KA",
            "type": "smarttv"
        }
    },
    {
        "desc": "Android TV",
        "ua": "Mozilla/5.0 (Linux; Android 10; 2020/2021 UHD Android TV Build/QTG3.201102.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) version/4.0 Chrome/83.0.4103.101 Mobile Safari/537.36",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "Amazon Fire 7",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; KFAUWI) AppleWebKit/537.36 (KHTML, like Gecko) Silk/80.5.3 like Chrome/80.0.3987.162 Safari/537.36",
        "expect": {
            "vendor": "Amazon",
            "model": "KFAUWI",
            "type": "tablet"
        }
    },
    {
        "desc": "FaceBook Mobile App",
        "ua": "[FBAN/FBIOS;FBAV/283.0.0.44.117;FBBV/238386386;FBDV/iPhone12,1;FBMD/iPhone;FBSN/iOS;FBSV/13.6.1;FBSS/2;FBID/phone;FBLC/en_US;FBOP/5;FBRV/240127608]",
        "expect": {
            "vendor": "Apple",
            "model": "iPhone12,1",
            "type": "mobile"
        }
    },
    {
        "desc": "Issue #519",
        "ua": "ios/iPhone/14.2/SOME_CUSTOM_APP_VERSION",
        "expect": {
            "vendor": "Apple",
            "model": "iPhone",
            "type": "mobile"
        }
    },
    {
        "desc": "Issue #454",
        "ua": "Mosamzilla/5.0 (Windows; U; Win98; en-US; rv:1.7.5) Gecko/20050603 Netscape/8.0.2",
        "expect": {
            "vendor": "",
            "model": "",
            "type": ""
        }
    },
    {
        "desc": "Alcatel",
        "ua": "Mozilla/5.0 (Linux; Android 4.4.2; ALCATEL A564C Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.133 Mobile Safari/537.36",
        "expect": {
            "vendor": "ALCATEL",
            "model": "A564C",
            "type": "mobile"
        }
    },
    {
        "desc": "Alcatel Go Flip",
        "ua": "Mozilla/5.0 (Mobile; ALCATEL4044T; rv:37.0) Gecko/37.0 Firefox/37.0 KaiOS/1.0",
        "expect": {
            "vendor": "ALCATEL",
            "model": "4044T",
            "type": "mobile"
        }
    },
    {
        "desc": "Jolla",
        "ua": "Mozilla/5.0 (Maemo; Linux; U; Jolla; Sailfish; Mobile; rv:31.0) Gecko/31.0 Firefox/31.0 SailfishBrowser/1.0",
        "expect": {
            "vendor": "Jolla",
            "model": "",
            "type": "mobile"
        }
    },
    {
        "desc": "Xbox One",
        "ua": "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; Xbox; Xbox One)",
        "expect": {
            "vendor": "Microsoft",
            "model": "Xbox One",
            "type": "console"
        }
    },
    {
        "desc": "Xbox",
        "ua": "Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Xbox)",
        "expect": {
            "vendor": "Microsoft",
            "model": "Xbox",
            "type": "console"
        }
    },
    {
        "desc": "Nvidia Shield Tablet",
        "ua": "Mozilla/5.0 (Linux; Android 5.1.1; SHIELD Tablet Build/LVY48E; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/45.0.2454.19 Safari/537.36",
        "expect": {
            "vendor": "Nvidia",
            "model": "SHIELD Tablet",
            "type": "tablet"
        }
    },
    {
        "desc": "Ouya",
        "ua": "Mozilla/5.0 (Linux; Android 4.1.2; OUYA Console Build/JZO54L-OUYA) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.84 Safari/537.36",
        "expect": {
            "vendor": "OUYA",
            "model": "",
            "type": "console"
        }
    },
    {
        "desc": "Vivo S1 Pro",
        "ua": "Mozilla/5.0 (Linux; Android 11; vivo 1920) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Vivo",
            "model": "1920",
            "type": "mobile"
        }
    },
    {
        "desc": "Vivo Y52s",
        "ua": "Mozilla/5.0 (Linux; Android 10; V2057A Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/12.10 SP-engine/2.28.0 baiduboxapp/12.10.0.10 (Baidu; P1 10) NABar/1.0",
        "expect": {
            "vendor": "Vivo",
            "model": "V2057A",
            "type": "mobile"
        }
    },
    {
        "desc": "Vivo X60",
        "ua": "Mozilla/5.0 (Linux; Android 11; V2046A; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/8.8.71.0",
        "expect": {
            "vendor": "Vivo",
            "model": "V2046A",
            "type": "mobile"
        }
    },
    {
        "desc": "Vivo Y79A",
        "ua": "Mozilla/5.0 (Linux; Android 7.1.2; vivo Y79A Build/N2G47H; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/9.0.14.0",
        "expect": {
            "vendor": "Vivo",
            "model": "Y79A",
            "type": "mobile"
        }
    },
    {
        "desc": "Vivo Y93",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; vivo 1814) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Vivo",
            "model": "1814",
            "type": "mobile"
        }
    },
    {
        "desc": "Vivo Y97",
        "ua": "Mozilla/5.0 (Linux; Android 8.1.0; V1813T Build/O11019; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/9.0.14.0",
        "expect": {
            "vendor": "Vivo",
            "model": "V1813T",
            "type": "mobile"
        }
    },
    {
        "desc": "Vivo iQOO Pro",
        "ua": "Mozilla/5.0 (Linux; Android 11; V1916A; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/9.1.10.6",
        "expect": {
            "vendor": "Vivo",
            "model": "V1916A",
            "type": "mobile"
        }
    },
    {
        "desc": "Vivo 1906 (Y11)",
        "ua": "Mozilla/5.0 (Linux; Android 11; vivo 1906) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",
        "expect": {
            "vendor": "Vivo",
            "model": "1906",
            "type": "mobile"
        }
    },
    {
        "desc": "Unknown Mobile using Firefox",
        "ua": "Mozilla/5.0 (Android 4.4; Mobile; rv:41.0) Gecko/41.0 Firefox/41.0",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "mobile"
        }
    },
    {
        "desc": "Unknown Tablet using Firefox",
        "ua": "Mozilla/5.0 (Android 4.4; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "tablet"
        }
    },
    {
        "desc": "Unknown Mobile using Focus for Android",
        "ua": "Mozilla/5.0 (Linux; Android 7.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Focus/1.0 Chrome/59.0.3029.83 Mobile Safari/537.36",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "mobile"
        }
    },
    {
        "desc": "Unknown Tablet using Focus for Android",
        "ua": "Mozilla/5.0 (Linux; Android 7.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Focus/1.0 Chrome/59.0.3029.83 Safari/537.36",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "tablet"
        }
    },
    {
        "desc": "Unknown Device using Focus for Android with GeckoView",
        "ua": "Mozilla/5.0 (Android 7.0; Mobile; rv:62.0) Gecko/62.0 Firefox/62.0",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "mobile"
        }
    },
    {
        "desc": "Unknown Mobile using Firefox OS",
        "ua": "Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "mobile"
        }
    },
    {
        "desc": "Unknown Tablet using Firefox OS",
        "ua": "Mozilla/5.0 (Tablet; rv:26.0) Gecko/26.0 Firefox/26.0",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "tablet"
        }
    },
    {
        "desc": "Unknown TV using Firefox OS",
        "ua": "Mozilla/5.0 (TV; rv:44.0) Gecko/44.0 Firefox/44.0",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "smarttv"
        }
    },
    {
        "desc": "PDA with Windows CE",
        "ua": "Mozilla/4.0 (PDA; Windows CE/1.0.1) NetFront/3.0",
        "expect": {
            "vendor": "",
            "model": "",
            "type": "mobile"
        }
    }
]
`)

type DeviceCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Vendor string `json:"vendor"`
		Model  string `json:"model,omitempty"`
		Type   string `json:"type,omitempty"`
	} `json:"expect"`
}

func TestDevice(t *testing.T) {
	var deviceCases []DeviceCase
	_ = json.Unmarshal(testDevice, &deviceCases)
	for _, uacase := range deviceCases {
		u := Parse(uacase.Ua)

		if !assert.Equal(t, uacase.Expect.Vendor, u.GetDevice().Vendor) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Model, u.GetDevice().Model) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Type, u.GetDevice().Type) {
			fmt.Println(uacase.Ua)
		}
	}
}

var testEngine = []byte(`
[
    {
        "desc"    : "Blink",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.0; SM-G920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/3.4.9 SamsungBrowser/4.0 Chrome/57.0.2987.146 Mobile VR Safari/537.36",
        "expect"  :
        {
            "name"    : "Blink",
            "version" : "57.0.2987.146"
        }
    },
    {
        "desc"    : "EdgeHTML",
        "ua"      : "Mozilla/5.0 (Windows NT 6.4; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36 Edge/12.0",
        "expect"  :
        {
            "name"    : "EdgeHTML",
            "version" : "12.0"
        }
    },
    {
        "desc"    : "Flow",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) EkiohFlow/5.7.4.30559 Flow/5.7.4 (like Gecko Firefox/53.0 rv:53.0)",
        "expect"  :
        {
            "name"    : "Flow",
            "version" : "5.7.4.30559"
        }
    },
    {
        "desc"    : "Gecko",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; rv:2.0b9pre) Gecko/20110111 Firefox/4.0b9pre",
        "expect"  :
        {
            "name"    : "Gecko",
            "version" : "2.0b9pre"
        }
    },
    {
        "desc"    : "Goanna",
        "ua"      : "Mozilla/5.0 (Windows NT 5.1; rv:38.9) Gecko/20100101 Goanna/2.2 Firefox/38.9 PaleMoon/26.5.0",
        "expect"  :
        {
            "name"    : "Goanna",
            "version" : "2.2"
        }
    },
    {
        "desc"    : "KHTML",
        "ua"      : "Mozilla/5.0 (compatible; Konqueror/4.5; FreeBSD) KHTML/4.5.4 (like Gecko)",
        "expect"  :
        {
            "name"    : "KHTML",
            "version" : "4.5.4"
        }
    },
    {
        "desc"    : "LibWeb",
        "ua"      : "Mozilla/4.0 (SerenityOS; x86) LibWeb+LibJS (Not KHTML, nor Gecko) LibWeb",
        "expect"  :
        {
            "name"    : "LibWeb",
            "version" : ""
        }
    },
    {
        "desc"    : "NetFront",
        "ua"      : "Mozilla/4.0 (PDA; Windows CE/1.0.1) NetFront/3.0",
        "expect"  :
        {
            "name"    : "NetFront",
            "version" : "3.0"
        }
    },
    {
        "desc"    : "Presto",
        "ua"      : "Opera/9.80 (Windows NT 6.1; Opera Tablet/15165; U; en) Presto/2.8.149 Version/11.1",
        "expect"  :
        {
            "name"    : "Presto",
            "version" : "2.8.149"
        }
    },
    {
        "desc"    : "Tasman",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 6.0; PPC Mac OS X 10.4.7; Tasman 1.0)",
        "expect"  :
        {
            "name"    : "Tasman",
            "version" : "1.0"
        }
    },
    {
        "desc"    : "Trident",
        "ua"      : "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
        "expect"  :
        {
            "name"    : "Trident",
            "version" : "6.0"
        }
    },
    {
        "desc"    : "WebKit",
        "ua"      : "Mozilla/5.0 (Windows; U; Windows NT 6.1; sv-SE) AppleWebKit/533.19.4 (KHTML, like Gecko) Version/5.0.3 Safari/533.19.4",
        "expect"  :
        {
            "name"    : "WebKit",
            "version" : "533.19.4"
        }
    },
    {
        "desc"    : "WebKit",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/27.0.1453.110 Safari/537.36",
        "expect"  :
        {
            "name"    : "WebKit",
            "version" : "537.36"
        }
    },
    {
        "desc"    : "WebOS TV 5.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36 WebAppManager",
        "expect"  :
        {
            "name"    : "Blink",
            "version" : "68.0.3440.106"
        }
    },
    {
        "desc"    : "WebOS TV 4.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.34 Safari/537.36 WebAppManager",
        "expect"  :
        {
            "name"    : "Blink",
            "version" : "53.0.2785.34"
        }
    },
    {
        "desc"    : "WebOS TV 3.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) QtWebEngine/5.2.1 Chrome/38.0.2125.122 Safari/537.36 WebAppManager",
        "expect"  :
        {
            "name"    : "Blink",
            "version" : "38.0.2125.122"
        }
    },
    {
        "desc"    : "WebOS TV 2.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/538.2 (KHTML, like Gecko) Large Screen WebAppManager Safari/538.2",
        "expect"  :
        {
            "name"    : "WebKit",
            "version" : "538.2"
        }
    },
    {
        "desc"    : "WebOS TV 1.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.41 (KHTML, like Gecko) Large Screen WebAppManager Safari/537.41",
        "expect"  :
        {
            "name"    : "WebKit",
            "version" : "537.41"
        }
    }
]
`)

type EngineCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"expect"`
}

func TestEngine(t *testing.T) {
	var engineCases []EngineCase
	_ = json.Unmarshal(testEngine, &engineCases)
	for _, uacase := range engineCases {
		u := Parse(uacase.Ua)

		if !assert.Equal(t, uacase.Expect.Name, u.GetEngine().Name) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Version, u.GetEngine().Version) {
			fmt.Println(uacase.Ua)
		}
	}
}

var testOs = []byte(`
[
    {
        "desc"    : "Windows 95",
        "ua"      : "Mozilla/1.22 (compatible; MSIE 2.0; Windows 95)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "95"
        }
    },
    {
        "desc"    : "Windows 98",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 4.01; Windows 98)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "98"
        }
    },
    {
        "desc"    : "Windows ME",
        "ua"      : "Mozilla/5.0 (Windows; U; Win 9x 4.90) Gecko/20020502 CS 2000 7.0/7.0",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "ME"
        }
    },
    {
        "desc"    : "Windows 2000",
        "ua"      : "Mozilla/3.0 (compatible; MSIE 3.0; Windows NT 5.0)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "2000"
        }
    },
    {
        "desc"    : "Windows XP",
        "ua"      : "Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 5.2)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "XP"
        }
    },
    {
        "desc"    : "Windows Vista",
        "ua"      : "Mozilla/5.0 (compatible; MSIE 7.0; Windows NT 6.0; fr-FR)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "Vista"
        }
    },
    {
        "desc"    : "Windows 7",
        "ua"      : "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "7"
        }
    },
    {
        "desc"    : "Windows 8",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; Win64; x64; Trident/6.0; .NET4.0E; .NET4.0C)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "8"
        }
    },
    {
        "desc"    : "Windows 10",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36 Edge/12.0",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "10"
        }
    },
    {
        "desc"    : "WeChat Desktop for Windows Built-in Browser",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 MicroMessenger/6.5.2.501 NetType/WIFI WindowsWechat QBCore/3.43.901.400 QQBrowser/9.0.2524.400",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "7"
        }
    },
    {
        "desc"    : "WeChat Desktop for Windows Built-in Browser major version in 4",
        "ua"      : "mozilla/5.0 (windows nt 6.1; wow64) applewebkit/537.36 (khtml, like gecko) chrome/81.0.4044.138 safari/537.36 nettype/wifi micromessenger/7.0.20.1781(0x6700143b) windowswechat",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "7"
        }
    },
    {
        "desc"    : "Windows RT",
        "ua"      : "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "RT"
        }
    },
    {
        "desc"    : "Windows CE",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.11)",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "CE"
        }
    },
    {
        "desc"    : "Windows Mobile",
        "ua"      : "Mozilla/5.0 (ZTE-E_N72/N72V1.0.0B02;U;Windows Mobile/6.1;Profile/MIDP-2.0 Configuration/CLDC-1.1;320*240;CTC/2.0) IE/6.0 (compatible; MSIE 4.01; Windows CE; PPC)/UC Browser7.7.1.88",
        "expect"  :
        {
            "name"    : "Windows Mobile",
            "version" : "6.1"
        }
    },
    {
        "desc"    : "Windows Mobile",
        "ua"      : "Opera/9.80 (Windows Mobile; WCE; Opera Mobi/WMD-50433; U; en) Presto/2.4.13 Version/10.00",
        "expect"  :
        {
            "name"    : "Windows Mobile",
            "version" : ""
        }
    },
    {
        "desc"    : "Windows Phone",
        "ua"      : "Opera/9.80 (Windows Phone; Opera Mini/7.6.8/35.7518; U; ru) Presto/2.8.119 Version/11.10",
        "expect"  :
        {
            "name"    : "Windows Phone",
            "version" : ""
        }
    },
    {
        "desc"    : "Windows Phone OS",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; DELL; Venue Pro)",
        "expect"  :
        {
            "name"    : "Windows Phone OS",
            "version" : "7.0"
        }
    },
    {
        "desc"    : "Windows Phone 8",
        "ua"      : "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; HTC; Windows Phone 8X by HTC)",
        "expect"  :
        {
            "name"    : "Windows Phone",
            "version" : "8.0"
        }
    },
    {
        "desc"    : "Windows NT on x86 or aarch64 CPU using Firefox",
        "ua"      : "Mozilla/5.0 (Windows NT x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "NT x"
        }
    },
    {
        "desc"    : "Windows NT on x64 CPU using Firefox",
        "ua"      : "Mozilla/5.0 (Windows NT x.y; Win64; x64; rv:10.0) Gecko/20100101 Firefox/10.0",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "NT x"
        }
    },
    {
        "desc"    : "BlackBerry",
        "ua"      : "BlackBerry9300/5.0.0.912 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/378",
        "expect"  :
        {
            "name"    : "BlackBerry",
            "version" : "5.0.0.912"
        }
    },
    {
        "desc"    : "BlackBerry 10",
        "ua"      : "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.3+ (KHTML, like Gecko) Version/10.0.9.386 Mobile Safari/537.3+",
        "expect"  :
        {
            "name"    : "BlackBerry",
            "version" : "10"
        }
    },
    {
        "desc"    : "Tizen",
        "ua"      : "Mozilla/5.0 (SMART-TV; Linux; Tizen 2.3) AppleWebkit/538.1 (KHTML, like Gecko) SamsungBrowser/1.0 TV Safari/538.1",
        "expect"  :
        {
            "name"    : "Tizen",
            "version" : "2.3"
        }
    },
    {
        "desc"    : "Tizen",
        "ua"      : "Mozilla/5.0 (Linux; Tizen 2.3; SAMSUNG SM-Z130H) AppleWebKit/537.3 (KHTML, like Gecko) Version/2.3 Mobile Safari/537.3",
        "expect"  :
        {
            "name"    : "Tizen",
            "version" : "2.3"
        }
    },
    {
        "desc"    : "Tizen 6.0",
        "ua"      : "HbbTV/1.5.1 (+DRM;Samsung;SmartTV2021:UAU7000;T-KSU2EDEUC-1506.0;KantSU2e;urn:samsungtv:familyname:21_KANTSU2E_UHD_BASIC:2021;) Tizen/6.0 (+TVPLUS+SmartHubLink) Chrome/76 LaTivu_1.0.1_2021 RVID/17",
        "expect"  :
        {
            "name"    : "Tizen",
            "version" : "6.0"
        }
    },
    {
        "desc"    : "Android",
        "ua"      : "Mozilla/5.0 (Linux; U; Android 2.2.2; en-us; VM670 Build/FRG83G) AppleWebKit/533.1 (KHTML, like Gecko)",
        "expect"  :
        {
            "name"    : "Android",
            "version" : "2.2.2"
        }
    },
    {
        "desc"    : "HarmonyOS",
        "ua"      : "Mozilla/5.0 (Linux; Android 10; HarmonyOS; YAL-AL10; HMSCore 6.3.0.327; GMSCore 21.48.15) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.105 HuaweiBrowser/12.0.3.310 Mobile Safari/537.36",
        "expect"  :
        {
            "name"    : "HarmonyOS",
            "version" : "10"
        }
    },
    {
        "desc"    : "Sailfish",
        "ua"      : "Mozilla/5.0 (Linux; U; Sailfish 3.0; Mobile; rv:45.0) Gecko/45.0 Firefox/45.0 SailfishBrowser/1.0",
        "expect"  :
        {
            "name"    : "Sailfish",
            "version" : "3.0"
        }
    },
    {
        "desc"    : "WebOS",
        "ua"      : "Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.5; U; en-US) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/234.83 Safari/534.6 TouchPad/1.0",
        "expect"  :
        {
            "name"    : "webOS",
            "version" : "3.0.5"
        }
    },
    {
        "desc"    : "WebOS",
        "ua"      : "Mozilla/5.0 (webOS/1.4.5; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
        "expect"  :
        {
            "name"    : "webOS",
            "version" : "1.4.5"
        }
    },
    {
        "desc"    : "WebOS TV 5.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36 WebAppManager",
        "expect"  :
        {
            "name"    : "webOS",
            "version" : "TV"
        }
    },
    {
        "desc"    : "WebOS TV 4.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.34 Safari/537.36 WebAppManager",
        "expect"  :
        {
            "name"    : "webOS",
            "version" : "TV"
        }
    },
    {
        "desc"    : "WebOS TV 3.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) QtWebEngine/5.2.1 Chrome/38.0.2125.122 Safari/537.36 WebAppManager",
        "expect"  :
        {
            "name"    : "webOS",
            "version" : "TV"
        }
    },
    {
        "desc"    : "WebOS TV 2.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/538.2 (KHTML, like Gecko) Large Screen WebAppManager Safari/538.2",
        "expect"  :
        {
            "name"    : "webOS",
            "version" : "TV"
        }
    },
    {
        "desc"    : "WebOS TV 1.x",
        "ua"      : "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.41 (KHTML, like Gecko) Large Screen WebAppManager Safari/537.41",
        "expect"  :
        {
            "name"    : "webOS",
            "version" : "TV"
        }
    },
    {
        "desc"    : "QNX",
        "ua"      : "Mozilla/5.0 (Photon; U; QNX x86pc; en-US; rv:1.8.1.20) Gecko/20090127 BonEcho/2.0.0.20",
        "expect"  :
        {
            "name"    : "QNX",
            "version" : "x86pc"
        }
    },
    {
        "desc"    : "Bada",
        "ua"      : "Mozilla/5.0 (SAMSUNG; SAMSUNG-GT-S5253/S5253DDKC1; U; Bada/1.0; en-us) AppleWebKit/533.1 (KHTML, like Gecko) Dolfin/2.0 Mobile WQVGA SMM-MMS/1.2.0 OPN-B",
        "expect"  :
        {
            "name"    : "Bada",
            "version" : "1.0"
        }
    },
    {
        "desc"    : "RIM Tablet OS",
        "ua"      : "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+",
        "expect"  :
        {
            "name"    : "RIM Tablet OS",
            "version" : "2.1.0"
        }
    },
    {
        "desc"    : "Nokia N900 Linux mobile, on the Fennec browser",
        "ua"      : "Mozilla/5.0 (Maemo; Linux armv7l; rv:10.0) Gecko/20100101 Firefox/10.0 Fennec/10.0",
        "expect"  :
        {
            "name"    : "Maemo",
            "version" : ""
        }
    },
    {
        "desc"    : "MeeGo",
        "ua"      : "Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13",
        "expect"  :
        {
            "name"    : "MeeGo",
            "version" : ""
        }
    },
    {
        "desc"    : "Symbian",
        "ua"      : "Nokia5250/10.0.011 (SymbianOS/9.4; U; Series60/5.0 Mozilla/5.0; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Safari/525 3gpp-gba",
        "expect"  :
        {
            "name"    : "Symbian",
            "version" : "9.4"
        }
    },
    {
        "desc"    : "Symbian",
        "ua"      : "Mozilla/5.0 (Symbian/3; Series60/5.2 NokiaC7-00/024.001; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.37 Mobile Safari/533.4 3gpp-gba",
        "expect"  :
        {
            "name"    : "Symbian",
            "version" : "5.2"
        }
    },
    {
        "desc"    : "Series40",
        "ua"      : "Mozilla/5.0 (Series40; Nokia2055/03.20; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/2.2.0.0.34",
        "expect"  :
        {
            "name"    : "Series40",
            "version" : ""
        }
    },
    {
        "desc"    : "Firefox OS",
        "ua"      : "Mozilla/5.0 (Mobile; rv:14.0) Gecko/14.0 Firefox/14.0",
        "expect"  :
        {
            "name"    : "Firefox OS",
            "version" : "14.0"
        }
    },
    {
        "desc"    : "Firefox OS on Tablet",
        "ua"      : "Mozilla/5.0 (Tablet; rv:26.0) Gecko/26.0 Firefox/26.0",
        "expect"  :
        {
            "name"    : "Firefox OS",
            "version" : "26.0"
        }
    },
    {
        "desc"    : "Firefox OS on TV",
        "ua"      : "Mozilla/5.0 (TV; rv:44.0) Gecko/44.0 Firefox/44.0",
        "expect"  :
        {
            "name"    : "Firefox OS",
            "version" : "44.0"
        }
    },
    {
        "desc"    : "Google Chromecast with Google TV",
        "ua"      : "Mozilla/5.0 (Linux; Android 12.0; Build/STTL.240206.002) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.0 Safari/537.36 CrKey/1.56.500000 DeviceType/AndroidTV",
        "expect"  :
        {
            "name"    : "Chromecast Android",
            "version" : "12.0"
        }
    },
    {
        "desc"    : "Google Chromecast Nest Hub",
        "ua"      : "Mozilla/5.0 (Fuchsia) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 CrKey/1.56.500000",
        "expect"  :
        {
            "name"    : "Chromecast Fuchsia",
            "version" : "1.56.500000"
        }
    },
    {
        "desc"    : "Google Chromecast Mini Smart Speaker",
        "ua"      : "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.225 Safari/537.36 CrKey/1.56.500000 DeviceType/SmartSpeaker",
        "expect"  :
        {
            "name"    : "Chromecast SmartSpeaker",
            "version" : "1.56.500000"
        }
    },
    {
        "desc"    : "Google Chromecast Legacy Linux-Based",
        "ua"      : "Mozilla/5.0 (X11; Linux aarch64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.81 Safari/537.36 CrKey/1.42.183786",
        "expect"  :
        {
            "name"    : "Chromecast Linux",
            "version" : "1.42.183786"
        }
    },
    {
        "desc"    : "Nintendo Switch",
        "ua"      : "Mozilla/5.0 (Nintendo Switch; WifiWebAuthApplet) AppleWebKit/606.4 (KHTML, like Gecko) NF/6.0.1.15.4 NintendoBrowser/5.1.0.20393",
        "expect"  :
        {
            "name"    : "Nintendo",
            "version" : "Switch"
        }
    },
    {
        "desc"    : "PlayStation 4",
        "ua"      : "Mozilla/5.0 (PlayStation 4 3.00) AppleWebKit/537.73 (KHTML, like Gecko)",
        "expect"  :
        {
            "name"    : "PlayStation",
            "version" : "4"
        }
    },
    {
        "desc"    : "PlayStation 5",
        "ua"      : "Mozilla/5.0 (PlayStation 5/SmartTV) AppleWebKit/605.1.15 (KHTML, like Gecko)",
        "expect"  :
        {
            "name"    : "PlayStation",
            "version" : "5"
        }
    },
    {
        "desc": "Pico 4",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; PICO 4 OS5.8.2 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.38 Chrome/105.0.5195.68 VR Safari/537.36",
        "expect": {
            "name"    : "PICO",
            "version" : "5.8.2"
        }
    },
    {
        "desc": "Pico 4",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; PICO 4 OS5.4.0 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.22 Chrome/105.0.5195.68 VR Safari/537.36 OculusBrowser/7.0",
        "expect": {
            "name"    : "PICO",
            "version" : "5.4.0"
        }
    },
    {
        "desc": "Pico Neo3 Link",
        "ua": "Mozilla/5.0 (X11; Linux x86_64; Pico Neo3 Link OS5.8.4.0 like Quest) AppleWebKit/537.36 (KHTML, like Gecko) PicoBrowser/3.3.22 Chrome/105.0.5195.68 VR Safari/537.36",
        "expect": {
            "name"    : "Pico",
            "version" : "5.8.4.0"
        }
    },
    {
        "desc"    : "Xbox 360",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox 360) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
        "expect"  :
        {
            "name"    : "Xbox",
            "version" : "360"
        }
    },
    {
        "desc"    : "Xbox One",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One; WebView/3.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.19041",
        "expect"  :
        {
            "name"    : "Xbox",
            "version" : "One"
        }
    },
    {
        "desc"    : "Xbox X",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.82 Safari/537.36 Edge/20.02",
        "expect"  :
        {
            "name"    : "Xbox",
            "version" : "X"
        }
    },
    {
        "desc"    : "Xbox Series X",
        "ua"      : "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox Series X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.82 Safari/537.36 Edge/20.02 ",
        "expect"  :
        {
            "name"    : "Xbox",
            "version" : "Series X"
        }
    },
    {
        "desc"    : "Xbox Series S",
        "ua"      : "Mozilla/5.0 (Compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; Xbox; Xbox Series S)",
        "expect"  :
        {
            "name"    : "Xbox",
            "version" : "Series S"
        }
    },
    {
        "desc"    : "Mint",
        "ua"      : "Opera/9.80 (X11; Linux x86_64; Edition Linux Mint) Presto/2.12.388 Version/12.16",
        "expect"  :
        {
            "name"    : "Mint",
            "version" : ""
        }
    },
    {
        "desc"    : "Mint",
        "ua"      : "Opera/9.64 (X11; Linux i686; U; Linux Mint; nb) Presto/2.1.1",
        "expect"  :
        {
            "name"    : "Mint",
            "version" : ""
        }
    },
    {
        "desc"    : "Mint",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.5) Gecko/2008121622 Linux Mint/6 (Felicia) Firefox/3.0.4",
        "expect"  :
        {
            "name"    : "Mint",
            "version" : "6"
        }
    },
    {
        "desc"    : "Ubuntu",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.22+ (KHTML, like Gecko) Chromium/17.0.963.56 Chrome/17.0.963.56 Safari/535.22+ Ubuntu/12.04 (3.4.1-0ubuntu1) Epiphany/3.4.1",
        "expect"  :
        {
            "name"    : "Ubuntu",
            "version" : "12.04"
        }
    },
    {
        "desc"    : "Ubuntu",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/31.0.1650.63 Chrome/31.0.1650.63 Safari/537.36",
        "expect"  :
        {
            "name"    : "Ubuntu",
            "version" : ""
        }
    },
    {
        "desc"    : "Kubuntu",
        "ua"      : "Mozilla/5.0 (compatible; Konqueror/4.4; Linux 2.6.32-22-generic; X11; en_US) KHTML/4.4.3 (like Gecko) Kubuntu",
        "expect"  :
        {
            "name"    : "Kubuntu",
            "version" : ""
        }
    },
    {
        "desc"    : "Debian",
        "ua"      : "Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.7 (like Gecko) (Debian)",
        "expect"  :
        {
            "name"    : "Debian",
            "version" : ""
        }
    },
    {
        "desc"    : "Debian",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; Debian GNU/Linux 8.1 (jessie)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.0 Maxthon/1.0.5.3 Safari/537.36",
        "expect"  :
        {
            "name"    : "Debian",
            "version" : "8.1"
        }
    },
    {
        "desc"    : "Debian",
        "ua"      : "ELinks/0.12~pre5-4 (textmode; Debian; Linux 3.2.0-4-amd64 x86_64 192x47-2)",
        "expect"  :
        {
            "name"    : "Debian",
            "version" : ""
        }
    },
    {
        "desc"    : "Debian",
        "ua"      : "w3m/0.5.3+debian-19",
        "expect"  :
        {
            "name"    : "debian",
            "version" : "19"
        }
    },
    {
        "desc"    : "Debian",
        "ua"      : "Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 (Debian-3.0.1-1)",
        "expect"  :
        {
            "name"    : "Debian",
            "version" : "3.0.1-1"
        }
    },
    {
        "desc"    : "Debian",
        "ua"      : "Mozilla/5.0 (compatible; Konqueror/3.5; Linux 2.6.24.4; X11) KHTML/3.5.9 (like Gecko) (Debian package 4:3.5.9.dfsg.1-2+b1)",
        "expect"  :
        {
            "name"    : "Debian",
            "version" : ""
        }
    },
    {
        "desc"    : "OpenSUSE",
        "ua"      : "Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.17) Gecko/20110420 SUSE/3.6.17-0.2.1 Firefox/3.6.17",
        "expect"  :
        {
            "name"    : "SUSE",
            "version" : "3.6.17-0.2.1"
        }
    },
    {
        "desc"    : "Gentoo",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1.16) Gecko/20080716 (Gentoo) Galeon/2.0.6",
        "expect"  :
        {
            "name"    : "Gentoo",
            "version" : ""
        }
    },
    {
        "desc"    : "Gentoo",
        "ua"      : "Xombrero (X11; U; Gentoo Linux amd64; en-US) Webkit/2.8.5",
        "expect"  :
        {
            "name"    : "Gentoo",
            "version" : "amd64"
        }
    },
    {
        "desc"    : "Gentoo",
        "ua"      : "Xombrero/1.6.4 (Linux amd64; en; Gentoo)",
        "expect"  :
        {
            "name"    : "Gentoo",
            "version" : ""
        }
    },
    {
        "desc"    : "Gentoo",
        "ua"      : "Links (2.8; Linux 3.17.2-gentoo-x86 i686; GNU C 4.8.2; x)",
        "expect"  :
        {
            "name"    : "gentoo",
            "version" : "x86"
        }
    },
    {
        "desc"    : "Arch",
        "ua"      : "Uzbl (Webkit 1.1.10) (Arch Linux)",
        "expect"  :
        {
            "name"    : "Arch",
            "version" : ""
        }
    },
    {
        "desc"    : "Slackware",
        "ua"      : "Mozilla/5.0 Slackware/13.37 (X11; U; Linux x86_64; en-US) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.41",
        "expect"  :
        {
            "name"    : "Slackware",
            "version" : "13.37"
        }
    },
    {
        "desc"    : "Fedora",
        "ua"      : "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:40.0) Gecko/20100101 Firefox/40.0",
        "expect"  :
        {
            "name"    : "Fedora",
            "version" : ""
        }
    },
    {
        "desc"    : "Fedora",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-GB; rv:2.0) Gecko/20110404 Fedora/16-dev Firefox/4.0",
        "expect"  :
        {
            "name"    : "Fedora",
            "version" : "16-dev"
        }
    },
    {
        "desc"    : "Fedora",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; sk; rv:1.9.0.4) Gecko/2008111217 Fedora/3.0.4-1.fc10 Firefox/3.0.4",
        "expect"  :
        {
            "name"    : "Fedora",
            "version" : "3.0.4-1.fc10"
        }
    },
    {
        "desc"    : "Mandriva",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.22) Gecko/20110907 Mandriva Linux/1.9.2.22-0.1mdv2010.2 (2010.2) Firefox/3.6.22",
        "expect"  :
        {
            "name"    : "Mandriva",
            "version" : "1.9.2.22-0.1mdv2010.2"
        }
    },
    {
        "desc"    : "Chrome OS",
        "ua"      : "Mozilla/5.0 (X11; CrOS x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.0.0 Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome OS",
            "version" : ""
        }
    },
    {
        "desc"    : "Chromium OS",
        "ua"      : "Mozilla/5.0 (X11; CrOS x86_64 10575.58.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
        "expect"  :
        {
            "name"    : "Chrome OS",
            "version" : "10575.58.0"
        }
    },
    {
        "desc"    : "Fuchsia",
        "ua"      : "Mozilla/5.0 (X11; Fuchsia x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3557.0 Safari/537.36",
        "expect"  :
        {
            "name"    : "Fuchsia",
            "version" : ""
        }
    },
    {
        "desc"    : "Solaris",
        "ua"      : "Mozilla/5.0 (X11; U; SunOS sun4u; en-US; rv:1.7) Gecko/20070606",
        "expect"  :
        {
            "name"    : "Solaris",
            "version" : "sun4u"
        }
    },
    {
        "desc"    : "FreeBSD",
        "ua"      : "Mozilla/5.0 (X11; U; FreeBSD x86_64; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.204 Safari/534.16",
        "expect"  :
        {
            "name"    : "FreeBSD",
            "version" : ""
        }
    },
    {
        "desc"    : "OpenBSD",
        "ua"      : "Mozilla/5.0 (X11; U; OpenBSD i386; en-US; rv:1.9.1) Gecko/20090702 Firefox/3.5",
        "expect"  :
        {
            "name"    : "OpenBSD",
            "version" : ""
        }
    },
    {
        "desc"    : "NetBSD",
        "ua"      : "ELinks (0.4.3; NetBSD 3.0.2PATCH sparc64; 141x19)",
        "expect"  :
        {
            "name"    : "NetBSD",
            "version" : "3.0.2PATCH"
        }
    },
    {
        "desc"    : "DragonFly",
        "ua"      : "Mozilla/5.0 (X11; U; DragonFly i386; de; rv:1.9.1) Gecko/20090720 Firefox/3.5.1",
        "expect"  :
        {
            "name"    : "DragonFly",
            "version" : ""
        }
    },
    {
        "desc"    : "iOS in App",
        "ua"      : "AppName/version CFNetwork/version Darwin/version",
        "expect"  :
        {
            "name"    : "iOS",
            "version" : ""
        }
    },
    {
        "desc"    : "iOS with Chrome",
        "ua"      : "Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3",
        "expect"  :
        {
            "name"    : "iOS",
            "version" : "5.1.1"
        }
    },
    {
        "desc"    : "iOS with Opera Mini",
        "ua"      : "Opera/9.80 (iPhone; Opera Mini/7.1.32694/27.1407; U; en) Presto/2.8.119 Version/11.10",
        "expect"  :
        {
            "name"    : "iOS",
            "version" : ""
        }
    },
    {
        "desc": "iOS with FaceBook Mobile App",
        "ua": "[FBAN/FBIOS;FBAV/283.0.0.44.117;FBBV/238386386;FBDV/iPhone12,1;FBMD/iPhone;FBSN/iOS;FBSV/13.6.1;FBSS/2;FBID/phone;FBLC/en_US;FBOP/5;FBRV/240127608]",
        "expect":
        {
            "name"    : "iOS",
            "version" : "13.6.1"
        }
    },
    {
        "desc": "iOS with Slack App",
        "ua": "com.tinyspeck.chatlyio/23.04.10 (iPhone; iOS 16.4.1; Scale/3.00)",
        "expect":
        {
            "name"    : "iOS",
            "version" : "16.4.1"
        }
    },
    {
        "desc"    : "watchOS",
        "ua"      : "server-bag [Watch OS,8.4,19S546,Watch3,4]",
        "expect"  :
        {
            "name"    : "watchOS",
            "version" : "8.4"
        }
    },
    {
        "desc"    : "watchOS",
        "ua"      : "atc/1.0 watchOS/7.4.1 model/Watch3,3 hwp/t8004 build/18T201 (6; dt:155)",
        "expect"  :
        {
            "name"    : "watchOS",
            "version" : "7.4.1"
        }
    },
    {
        "desc"    : "watchOS",
        "ua"      : "Watch4,3/5.3.8 (16U680)",
        "expect"  :
        {
            "name"    : "watchOS",
            "version" : "5.3.8"
        }
    },
    {
        "desc"    : "Mac OS on PowerPC",
        "ua"      : "Mozilla/4.0 (compatible; MSIE 5.0b1; Mac_PowerPC)",
        "expect"  :
        {
            "name"    : "macOS",
            "version" : ""
        }
    },
    {
        "desc"    : "Mac OS X on x86, x86_64, or aarch64 using Firefox",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
        "expect"  :
        {
            "name"    : "macOS",
            "version" : "x.y"
        }
    },
    {
        "desc"    : "Mac OS X on PowerPC using Firefox",
        "ua"      : "Mozilla/5.0 (Macintosh; PPC Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
        "expect"  :
        {
            "name"    : "macOS",
            "version" : "x.y"
        }
    },
    {
        "desc"    : "Mac OS",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.95 Safari/537.36",
        "expect"  :
        {
            "name"    : "macOS",
            "version" : "10.6.8"
        }
    },
    {
        "desc"    : "Haiku",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Haiku R1 x86) AppleWebKit/602.1.1 (KHTML, like Gecko) WebPositive/1.2 Version/8.0 Safari/602.1.1",
        "expect"  :
        {
            "name"    : "Haiku",
            "version" : "R1"
        }
    },
    {
        "desc"    : "KaiOS",
        "ua"      : "Mozilla/5.0 (Mobile; Nokia_8110_4G; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5",
        "expect"  :
        {
            "name"    : "KAIOS",
            "version" : "2.5"
        }
    },
    {
        "desc"    : "iTunes Windows Vista",
        "ua"      : "iTunes/10.7 (Windows; Microsoft Windows Vista Home Premium Edition Service Pack 1 (Build 6001)) AppleWebKit/536.26.9",
        "expect"  :
        {
            "name"    : "Windows",
            "version" : "Vista"
        }
    },
    {
        "desc"    : "iOS BE App",
        "ua"      : "APP-BE Test/1.0 (iPad; Apple; CPU iPhone OS 7_0_2 like Mac OS X)",
        "expect"  :
        {
            "name"    : "iOS",
            "version" : "7.0.2"
        }
    },
    {
        "desc"    : "KTB-Nexus 5",
        "ua"      : "APP-My App/1.0 (Linux; Android 4.2.1; Nexus 5 Build/JOP40D)",
        "expect"  :
        {
            "name"    : "Android",
            "version" : "4.2.1"
        }
    },
    {
        "desc"    : "Solaris",
        "ua"      : "NCSA Mosaic/1.0 (X11;SunOS 4.1.4 sun4m)",
        "expect"  :
        {
            "name"    : "Solaris",
            "version" : "4.1.4"
        }
    },
    {
        "desc"    : "Raspbian",
        "ua"      : "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Raspbian Chromium/72.0.3626.121 HeadlessChrome/72.0.3626.121 Safari/537.36",
        "expect"  :
        {
            "name"    : "Raspbian",
            "version" : ""
        }
    },
    {
        "desc"    : "Raspbian",
        "ua"      : "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/538.15 (KHTML, like Gecko) Version/8.0 Safari/538.15 Raspbian/9.0 (1:3.8.2.0-0rpi28) Epiphany/3.8.2",
        "expect"  :
        {
            "name"    : "Raspbian",
            "version" : "9.0"
        }
    },
    {
        "desc"    : "AIX",
        "ua"      : "Mozilla/5.0 (X11; U; AIX 000138384C00; en-US; rv:1.0.1) Gecko/20030213 Netscape/7.0",
        "expect"  :
        {
            "name"    : "AIX",
            "version" : ""
        }
    },
    {
        "desc"    : "Plan9",
        "ua"      : "NCSA_Mosaic/5.0 (X11;Plan 9 4.0)",
        "expect"  :
        {
            "name"    : "Plan 9",
            "version" : "4.0"
        }
    },
    {
        "desc"    : "Minix",
        "ua"      : "Mozilla/5.0 (X11; Original ; Minix 3.3 ; rv:3.0)",
        "expect"  :
        {
            "name"    : "Minix",
            "version" : "3.3"
        }
    },
    {
        "desc"    : "BeOS",
        "ua"      : "Mozilla/5.0 (BeOS; U; BeOS BePC; en-US; rv:1.8.1.8pre) Gecko/20070926 SeaMonkey/1.1.5pre",
        "expect"  :
        {
            "name"    : "BeOS",
            "version" : ""
        }
    },
    {
        "desc"    : "OS/2",
        "ua"      : "Links (2.1pre14; OS/2 1 i386; 80x33)",
        "expect"  :
        {
            "name"    : "OS/2",
            "version" : ""
        }
    },
    {
        "desc"    : "AmigaOS",
        "ua"      : "Mozilla/4.0 (compatible; AWEB 3.4 SE; AmigaOS)",
        "expect"  :
        {
            "name"    : "AmigaOS",
            "version" : ""
        }
    },
    {
        "desc"    : "MorphOS",
        "ua"      : "AmigaVoyager/3.4.4 (MorphOS/PPC native)",
        "expect"  :
        {
            "name"    : "MorphOS",
            "version" : ""
        }
    },
    {
        "desc"    : "UNIX",
        "ua"      : "Surf/0.4.1 (X11; U; Unix; en-US) AppleWebKit/531.2+ Compatible (Safari)",
        "expect"  :
        {
            "name"    : "Unix",
            "version" : ""
        }
    },
    {
        "desc"    : "Joli",
        "ua"      : "Mozilla/5.0 (X11; Jolicloud Linux i686) AppleWebKit/537.6 (KHTML, like Gecko) Joli OS/1.2 Chromium/23.0.1240.0 Chrome/23.0.1240.0 Safari/537.6",
        "expect"  :
        {
            "name"    : "Joli",
            "version" : "1.2"
        }
    },
    {
        "desc"    : "CentOS",
        "ua"      : "Konqueror/15.13 (CentOS Linux 7.4; cs-CZ;)",
        "expect"  :
        {
            "name"    : "CentOS",
            "version" : "7.4"
        }
    },
    {
        "desc"    : "PCLinuxOS",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.13) Gecko/20101209 PCLinuxOS/1.9.2.13-1pclos2010 (2010) Firefox/3.6.13",
        "expect"  :
        {
            "name"    : "PCLinuxOS",
            "version" : "1.9.2.13-1pclos2010"
        }
    },
    {
        "desc"    : "RedHat",
        "ua"      : "Mozilla/5.0 (compatible; Konqueror/4.3; Linux) KHTML/4.3.4 (like Gecko) Red Hat Enterprise Linux/4.3.4-11.el6_1.4",
        "expect"  :
        {
            "name"    : "Red Hat",
            "version" : "4.3.4-11.el6_1.4"
        }
    },
    {
        "desc"    : "RedHat",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.13pre) Gecko/20070717 Red Hat/1.0.9-4.el4 SeaMonkey/1.0.9",
        "expect"  :
        {
            "name"    : "Red Hat",
            "version" : "1.0.9-4.el4"
        }
    },
    {
        "desc"    : "RedHat",
        "ua"      : "iTunes/4.7.1 (Linux; N; Red Hat; x86_64-linux; EN; utf8) SqueezeCenter, Squeezebox Server, Logitech Media Server/7.9.1/1522157629",
        "expect"  :
        {
            "name"    : "Red Hat",
            "version" : ""
        }
    },
    {
        "desc"    : "RedHat",
        "ua"      : "curl/7.20.0 (x86_64-redhat-linux-gnu) libcurl/7.20.0 OpenSSL/0.9.8b zlib/1.2.3 libidn/0.6.5",
        "expect"  :
        {
            "name"    : "redhat",
            "version" : ""
        }
    },
    {
        "desc"    : "RISC OS",
        "ua"      : "Mozilla/1.10 [en] (Compatible; RISC OS 3.70; Oregano 1.10)",
        "expect"  :
        {
            "name"    : "RISC OS",
            "version" : "3.70"
        }
    },
    {
        "desc"    : "Zenwalk",
        "ua"      : "Flock/2.16 (Zenwalk 7.3; es_PR;)",
        "expect"  :
        {
            "name"    : "Zenwalk",
            "version" : "7.3"
        }
    },
    {
        "desc"    : "Hurd",
        "ua"      : "Mozilla/5.0 (X11; Hurd 0.9 i386; en-US) libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/3.7.0 Safari/696.96",
        "expect"  :
        {
            "name"    : "Hurd",
            "version" : "0.9"
        }
    },
    {
        "desc"    : "Linux",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36",
        "expect"  :
        {
            "name"    : "Linux",
            "version" : "x86_64"
        }
    },
    {
        "desc"    : "Deepin",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64; Deepin 15.5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36 NFSBrowser/5.0.0.1886",
        "expect"  :
        {
            "name"    : "Deepin",
            "version" : "15.5"
        }
    },
    {
        "desc"    : "Palm OS",
        "ua"      : "Mozilla/4.76 [en] (PalmOS; U; WebPro3.0; Palm-Arz1)",
        "expect"  :
        {
            "name"    : "Palm",
            "version" : ""
        }
    },
    {
        "desc"    : "Panasonic Viera",
        "ua"      : "HbbTV/1.2.1 (;Panasonic;VIERA 2015;3.014;a001-003 4000-0000;)",
        "expect"  :
        {
            "name"    : "VIERA",
            "version" : ""
        }
    },
    {
        "desc"    : "Netrange Smart TV",
        "ua"      : "Mozilla/5.0 (Linux; U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36 OPR/46.0.2207.0 LOEWE-SL410/5.2.0.0 HbbTV/1.4.1 (; LOEWE; SL410; LOH/5.2.0.0;;) FVC/3.0 (LOEWE; SL410;) CE-HTML/1.0 Config (L:deu,CC:DEU) NETRANGEMMH",
        "expect"  :
        {
            "name"    : "NETRANGE",
            "version" : ""
        }
    },
    {
        "desc"    : "NetTV 3.2.1",
        "ua"      : "Opera/9.80 (Linux mips ; U; HbbTV/1.1.1 (; Philips; ; ; ; ) CE-HTML/1.0 NETTV/3.2.1; en) Presto/2.6.33 Version/10.70",
        "expect"  :
        {
            "name"    : "NETTV",
            "version" : "3.2.1"
        }
    },
    {
        "desc"    : "HP-UX",
        "ua"      : "Mozilla/5.0 (X11; U; HP-UX 9000/785; es-ES; rv:1.0.1) Gecko/20020827 Netscape/7.0",
        "expect"  :
        {
            "name"    : "HP-UX",
            "version" : ""
        }
    },
    {
        "desc"    : "Contiki",
        "ua"      : "Contiki/1.0 (Commodore 64; http://dunkels.com/adam/contiki/)",
        "expect"  :
        {
            "name"    : "Contiki",
            "version" : "1.0"
        }
    },
    {
        "desc"    : "Linpus",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9b5pre) Gecko/2008032619 Linpus/3.0-0.49",
        "expect"  :
        {
            "name"    : "Linpus",
            "version" : "3.0-0.49"
        }
    },
    {
        "desc"    : "Manjaro",
        "ua"      : "Mozilla/5.0 (X11; Manjaro 19.0.2; Arch; x64; rv:84.0) Gecko/20100101 Firefox/84.0",
        "expect"  :
        {
            "name"    : "Manjaro",
            "version" : "19.0.2"
        }
    },
    {
        "desc"    : "elementary OS",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/604.1 (KHTML, like Gecko) Version/11.0 Safari/604.1 elementary OS/0.4 (Loki) Epiphany/3.18.11",
        "expect"  :
        {
            "name"    : "elementary OS",
            "version" : "0.4"
        }
    },
    {
        "desc"    : "GhostBSD",
        "ua"      : "Mozilla/5.0 (X11; GhostBSD/10.3; x86_64; rv:50.0.1) Gecko/20100101 Firefox/50.0.1",
        "expect"  :
        {
            "name"    : "GhostBSD",
            "version" : "10.3"
        }
    },
    {
        "desc"    : "Android-x86",
        "ua"      : "Mozilla/5.0 (Linux; Android 7.1.2; Generic Android-x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36 OPR/61.2.3076.56749",
        "expect"  :
        {
            "name"    : "Android-x86",
            "version" : "7.1.2"
        }
    },
    {
        "desc"    : "Sabayon",
        "ua"      : "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Sabayon Chrome/19.0.1084.46 Safari/536.5",
        "expect"  :
        {
            "name"    : "Sabayon",
            "version" : ""
        }
    },
    {
        "desc"    : "Linspire",
        "ua"      : "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.4) Gecko/20060803 Firefox/1.5.0.4 Linspire/1.5.0.4",
        "expect"  :
        {
            "name"    : "Linspire",
            "version" : "1.5.0.4"
        }
    },
    {
        "desc"    : "SerenityOS",
        "ua"      : "Mozilla/4.0 (SerenityOS; x86) LibWeb+LibJS (Not KHTML, nor Gecko) LibWeb",
        "expect"  :
        {
            "name"    : "SerenityOS",
            "version" : ""
        }
    }
]
`)

type OsCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"expect"`
}

func TestOs(t *testing.T) {
	var osCases []OsCase
	_ = json.Unmarshal(testOs, &osCases)
	for _, uacase := range osCases {
		u := Parse(uacase.Ua)

		if !assert.Equal(t, uacase.Expect.Name, u.GetOS().Name) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Version, u.GetOS().Version) {
			fmt.Println(uacase.Ua)
		}
	}
}

var testCli = []byte(`
[
    {
        "desc"    : "curl",
        "ua"      : "curl/7.38.0",
        "expect"  :
        {
            "name"    : "curl",
            "version" : "7.38.0",
            "type"    : "cli"
        }
    },
    {
        "desc"    : "lynx",
        "ua"      : "Lynx 2.8.8dev.3",
        "expect"  :
        {
            "name"    : "Lynx",
            "version" : "2.8.8dev.3",
            "type"    : "cli"
        }
    },
    {
        "desc"    : "lynx",
        "ua"      : "Lynx/2.6",
        "expect"  :
        {
            "name"    : "Lynx",
            "version" : "2.6",
            "type"    : "cli"
        }
    },
    {
        "desc"    : "wget",
        "ua"      : "Wget/1.21.1",
        "expect"  :
        {
            "name"    : "Wget",
            "version" : "1.21.1",
            "type"    : "cli"
        }
    }
]
`)

type CliCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Type    string `json:"type"`
	} `json:"expect"`
}

func TestCli(t *testing.T) {
	var cliCases []CliCase
	_ = json.Unmarshal(testCli, &cliCases)
	for _, uacase := range cliCases {
		u := Parse(uacase.Ua, SetExtensions([]string{CLI}))

		if !assert.Equal(t, uacase.Expect.Name, u.GetBrowser().Name) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Version, u.GetBrowser().Version) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Type, u.GetBrowser().Type) {
			fmt.Println(uacase.Ua)
		}
	}
}

var testCrawlers = []byte(`
[
    {
        "desc"    : "AhrefsBot",
        "ua"      : "Mozilla/5.0 (compatible; AhrefsBot/7.0; +http://ahrefs.com/robot/)",
        "expect"  :
        {
            "name"    : "AhrefsBot",
            "version" : "7.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Applebot",
        "ua"      : "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B410 Safari/600.1.4 (Applebot/0.1;+http://www.apple.com/go/applebot)",
        "expect"  :
        {
            "name"    : "Applebot",
            "version" : "0.1",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Amazonbot",
        "ua"      : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Amazonbot/0.1; +https://developer.amazon.com/support/amazonbot)",
        "expect"  :
        {
            "name"    : "Amazonbot",
            "version" : "0.1",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Bytespider",
        "ua"      : "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.1511.1269 Mobile Safari/537.36; Bytespider",
        "expect"  :
        {
            "name"    : "Bytespider",
            "version" : "",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "ClaudeBot",
        "ua"      : "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; ClaudeBot/1.0; +claudebot@anthropic.com)",
        "expect"  :
        {
            "name"    : "ClaudeBot",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Coc Coc Bot (web)",
        "ua"      : "Mozilla/5.0 (compatible; coccocbot-web/1.0; +http://help.coccoc.com/searchengine)",
        "expect"  :
        {
            "name"    : "coccocbot-web",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Coc Coc Bot (image)",
        "ua"      : "Mozilla/5.0 (compatible; coccocbot-image/1.0; +http://help.coccoc.com/searchengine)",
        "expect"  :
        {
            "name"    : "coccocbot-image",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "ClaudeWeb",
        "ua"      : "Claude-Web/1.0 (web crawler; +https://www.anthropic.com/; bots@anthropic.com)",
        "expect"  :
        {
            "name"    : "Claude-Web",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Dotbot",
        "ua"      : "Mozilla/5.0 (compatible; DotBot/1.2; +https://opensiteexplorer.org/dotbot; help@moz.com)",
        "expect"  :
        {
            "name"    : "DotBot",
            "version" : "1.2",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "FacebookBot",
        "ua"      : "Mozilla/5.0 (compatible; FacebookBot/1.0; +https://developers.facebook.com/docs/sharing/webmasters/facebookbot/",
        "expect"  :
        {
            "name"    : "FacebookBot",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Googlebot-Video",
        "ua"      : "Googlebot-Video/1.0",
        "expect"  :
        {
            "name"    : "Googlebot-Video",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "GPTBot",
        "ua"      : "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; GPTBot/1.0; +https://openai.com/gptbot)",
        "expect"  :
        {
            "name"    : "GPTBot",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "MJ12bot",
        "ua"      : "Mozilla/5.0 (compatible; MJ12bot/v1.4.8; http://mj12bot.com/)",
        "expect"  :
        {
            "name"    : "MJ12bot",
            "version" : "v1.4.8",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "OpenAI Search",
        "ua"      : "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko); compatible; OAI-SearchBot/1.0; +https://openai.com/searchbot",
        "expect"  :
        {
            "name"    : "OAI-SearchBot",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "SemrushBot",
        "ua"      : "Mozilla/5.0 (compatible; SemrushBot/7~bl; +http://www.semrush.com/bot.html)",
        "expect"  :
        {
            "name"    : "SemrushBot",
            "version" : "7",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "Yahoo! Japan",
        "ua"      : "Y!J-BRW/1.0 (https://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716)",
        "expect"  :
        {
            "name"    : "Y!J-BRW",
            "version" : "1.0",
            "type"    : "crawler"
        }
    },
    {
        "desc"    : "YandexBot",
        "ua"      : "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
        "expect"  :
        {
            "name"    : "YandexBot",
            "version" : "3.0",
            "type"    : "crawler"
        }
    }
]
`)

type CrawlersCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Type    string `json:"type"`
	} `json:"expect"`
}

func TestCrawlers(t *testing.T) {
	var crawlersCases []CrawlersCase
	_ = json.Unmarshal(testCrawlers, &crawlersCases)
	for _, uacase := range crawlersCases {
		u := Parse(uacase.Ua, SetExtensions([]string{CRAWLER}))

		if !assert.Equal(t, uacase.Expect.Name, u.GetBrowser().Name) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Version, u.GetBrowser().Version) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Type, u.GetBrowser().Type) {
			fmt.Println(uacase.Ua)
		}
	}
}

var testFetchers = []byte(`
[
    {
        "desc"    : "AhrefsSiteAudit",
        "ua"      : "Mozilla/5.0 (Linux; Android 13) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.128 Mobile Safari/537.36 (compatible; AhrefsSiteAudit/6.1; +http://ahrefs.com/robot/site-audit)",
        "expect"  :
        {
            "name"    : "AhrefsSiteAudit",
            "version" : "6.1",
            "type"    : "fetcher"
        }
    },
    {
        "desc"    : "BingPreview",
        "ua"      : "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko) BingPreview/1.0b",
        "expect"  :
        {
            "name"    : "BingPreview",
            "version" : "1.0b",
            "type"    : "fetcher"
        }
    },
    {
        "desc"    : "ChatGPT-User",
        "ua"      : "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko); compatible; ChatGPT-User/1.0; +https://openai.com/bot",
        "expect"  :
        {
            "name"    : "ChatGPT-User",
            "version" : "1.0",
            "type"    : "fetcher"
        }
    },
    {
        "desc"    : "Rogerbot",
        "ua"      : "Mozilla/5.0 (compatible; rogerBot/1.0; UrlCrawler; http://www.seomoz.org/dp/rogerbot)",
        "expect"  :
        {
            "name"    : "rogerBot",
            "version" : "1.0",
            "type"    : "fetcher"
        }
    },
    {
        "desc"    : "UptimeRobot",
        "ua"      : "Mozilla/5.0 (compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)",
        "expect"  :
        {
            "name"    : "UptimeRobot",
            "version" : "2.0",
            "type"    : "fetcher"
        }
    }
]`)

type FetchersCase struct {
	Desc   string `json:"desc"`
	Ua     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Type    string `json:"type"`
	} `json:"expect"`
}

func TestFetchers(t *testing.T) {
	var fetchersCases []FetchersCase
	_ = json.Unmarshal(testFetchers, &fetchersCases)
	for _, uacase := range fetchersCases {
		u := Parse(uacase.Ua, SetExtensions([]string{FETCHER}))

		if !assert.Equal(t, uacase.Expect.Name, u.GetBrowser().Name) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Version, u.GetBrowser().Version) {
			fmt.Println(uacase.Ua)
		}
		if !assert.Equal(t, uacase.Expect.Type, u.GetBrowser().Type) {
			fmt.Println(uacase.Ua)
		}
	}
}
