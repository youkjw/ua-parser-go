package uaparser

import (
	"reflect"
	"regexp"
	"strings"
)

func extract(target any, source map[string]string) {
	targetType := reflect.TypeOf(target).Elem()
	targetValue := reflect.ValueOf(target).Elem()
	for i := 0; i < targetType.NumField(); i++ {
		var tag = targetType.Field(i).Tag.Get("json")
		v := targetValue.Field(i)
		switch v.Kind() {
		case reflect.Struct:
			extract(v.Interface(), source)
		case reflect.Ptr:
			extract(v.Elem().Interface(), source)
		case reflect.String:
			for itemType, value := range source {
				if itemType == tag {
					v.SetString(value)
					break
				}
			}
		}
	}
}

func trim(value string) string {
	reg := regexp.MustCompile(`^\s\s*`)
	value = reg.ReplaceAllString(value, Empty)
	if len(value) > UaMaxLength {
		return value[:UaMaxLength]
	}
	return value
}

// oldSafariBrowser
func oldSafariBrowser(b []byte) string {
	version := string(b)

	for fix, old := range safariMap {
		if strings.Index(version, old) >= 0 {
			return fix
		}
	}
	return version
}

func windowsVersionMapper(value string) string {
	for k, v := range windowsVersionMap {
		if has(v, value) {
			return k
		}
	}
	return value
}

func itelTypeMapper(value string) string {
	for k, v := range itelType {
		if has(v, value) {
			return k
		}
	}

	return wildcardsType["*"]
}

func has(s []string, v string) bool {
	v = strings.ToLower(v)

	for _, sval := range s {
		if strings.Contains(v, strings.ToLower(sval)) {
			return true
		}
	}
	return false
}
