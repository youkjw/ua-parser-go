# ua-parser-go
Go language version of ua-parser-js

## Installing
This is a go-gettable library, so install is easy:

    go get github.com/youkjw/ua-parser-go@latest

# Usage

## Using Go

```go
var uastr = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63030522)";
res := Parse(uastr)

browser := res.GetBrowser()
os := res.GetOS()
engine := res.GetEngine()
device := res.GetDevice()
fmt.Println(res.String()) // {"ua":"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63030522)","browser":{"name":"WeChat","version":"7.0.20.1781","major":"7"},"cpu":{"architecture":"amd64"},"engine":{"name":"Blink","version":"81.0.4044.138"},"device":{"vendor":"","model":"","type":""},"os":{"name":"Windows","version":"7"}}

```

## Result case
```json
{
  "ua": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63030522)",
  "browser": {
    "name": "WeChat",
    "version": "7.0.20.1781",
    "major": "7"
  },
  "cpu": {
    "architecture": "amd64"
  },
  "engine": {
    "name": "Blink",
    "version": "81.0.4044.138"
  },
  "device": {
    "vendor": "",
    "model": "",
    "type": ""
  },
  "os": {
    "name": "Windows",
    "version": "7"
  }
}
```

## Uses

- [regexp2](https://github.com/dlclark/regexp2): regexp2 - full featured regular expressions for Go.

## Thinks

- [ua-parser-js](https://github.com/faisalman/ua-parser-js): UAParser.js
