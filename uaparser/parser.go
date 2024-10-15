package uaparser

import (
	"github.com/dlclark/regexp2"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	RgxMapperIdle int32 = iota
	RgxMapperParsed
)

type parsedCall func(key, value string) string

type Parser struct {
	Regexps        []*regexp.Regexp
	ComplexRegexps []*regexp2.Regexp
	Props          []*Property
	ParsedCall     []parsedCall
}

func MakeMatchedCall(key string, call func(string) string) func(key, value string) string {
	return func(resKey, resValue string) string {
		if strings.Compare(resKey, key) != 0 {
			return resValue
		}

		return call(resValue)
	}
}

func (p *Parser) dealValue(key, value string) string {
	if len(p.ParsedCall) > 0 {
		for _, callback := range p.ParsedCall {
			value = callback(key, value)
		}
	}
	return value
}

type Property struct {
	Key   string
	Expr  *regexp.Regexp
	Value string
}

func BuildProperty(value ...string) *Property {
	if len(value) == 1 {
		return &Property{value[0], nil, ""}
	} else if len(value) == 2 {
		return &Property{value[0], nil, value[1]}
	} else if len(value) == 3 {
		return &Property{value[0], regexp.MustCompile(value[1]), value[2]}
	}
	return &Property{}
}

type rgxMapper struct {
	mux      sync.RWMutex
	itemType UaItemType
	Matches  map[string]string
	flag     int32
}

func newRgxMapper(itemType UaItemType) *rgxMapper {
	return &rgxMapper{
		itemType: itemType,
		Matches:  make(map[string]string),
		flag:     RgxMapperIdle,
	}
}

func (rgx *rgxMapper) clear() {
	rgx.Matches = make(map[string]string)
	atomic.SwapInt32(&rgx.flag, RgxMapperIdle)
}

func (rgx *rgxMapper) rgxParse(ua []byte, rgxMapper []*Parser) {
	if len(ua) == 0 {
		return
	}

	if len(rgxMapper) == 0 {
		rgxMapper = rgxMaps[rgx.itemType]
	}

	defer func() {
		if len(rgx.Matches) > 0 {
			atomic.SwapInt32(&rgx.flag, RgxMapperParsed)
		}
	}()

	var matched bool
	var matches = make([][]string, 0)
	for i := 0; i < len(rgxMapper); i++ {
		var (
			regexps        = rgxMapper[i].Regexps
			complexRegexps = rgxMapper[i].ComplexRegexps
			props          = rgxMapper[i].Props
		)

		// simple regexps
		for _, expr := range regexps {
			if expr == nil {
				continue
			}
			if matched = expr.Match(ua); matched {
				subMatch := expr.FindSubmatch(ua)
				if len(subMatch) > 0 {
					subMatches := make([]string, len(subMatch))
					for mi, match := range subMatch {
						subMatches[mi] = string(match)
					}
					matches = append(matches, subMatches)
				}

				// matched and break
				if matched {
					break
				}
			}
		}

		if !matched {
			// complex regexps
			for _, expr := range complexRegexps {
				if expr == nil {
					continue
				}
				if m, _ := expr.FindStringMatch(string(ua)); m != nil {
					matched = true
					subMatch := m.Groups()
					if len(subMatch) > 0 {
						subMatches := make([]string, len(subMatch))
						for mi, match := range subMatch {
							subMatches[mi] = match.String()
						}
						matches = append(matches, subMatches)
					}

					// matched and break
					if matched {
						break
					}
				}
			}
		}

		var matchesMap = make(map[string]string)
		for _, itemMatch := range matches {
			var k = 0
			for _, prop := range props {
				k++
				if len(prop.Key) == 0 {
					continue
				}

				var matchVal string
				if len(prop.Value) > 0 && prop.Expr == nil {
					matchVal = prop.Value
				} else {
					if len(itemMatch) > k {
						matchVal = itemMatch[k]
					}
				}

				if prop.Expr != nil {
					// sanitize match using given regex
					matchVal = prop.Expr.ReplaceAllString(matchVal, prop.Value)
				}
				matchesMap[prop.Key] = rgxMapper[i].dealValue(prop.Key, matchVal)
			}
		}

		// matched and break
		if matched {
			rgx.Matches = matchesMap
			break
		}
	}
}
