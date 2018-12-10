package util

import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// HandlePanic helper method to handle panics
func HandlePanic(name string, err *error) {
	if r := recover(); r != nil {

		logger.Warnf("%s: PANIC Occurred  : %v\n", name, r)

		// todo: useful for debugging
		logger.Debugf("StackTrace: %s", debug.Stack())

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}

// URLStringToFilePath convert fileURL to file path
func URLStringToFilePath(fileURL string) (string, bool) {

	if strings.HasPrefix(fileURL, "file://") {

		filePath := fileURL[7:]

		if runtime.GOOS == "windows" {
			if strings.HasPrefix(filePath, "/") {
				filePath = filePath[1:]
			}
			filePath = filepath.FromSlash(filePath)
		}

		filePath = strings.Replace(filePath, "%20", " ", -1)

		return filePath, true
	}

	return "", false
}

//ParseKeyValuePairs get key-value map from "key=value,key1=value1" str, if value have , please use quotes
func ParseKeyValuePairs(keyvalueStr string) map[string]string {
	m := make(map[string]string)
	parseKeyValue(removeQuote(keyvalueStr), m)
	return m
}

func parseKeyValue(keyvalueStr string, m map[string]string) {
	var key, value, rest string
	eidx := strings.Index(keyvalueStr, "=")
	if eidx >= 1 {
		//Remove space in case it has space between =
		key = strings.TrimSpace(keyvalueStr[:eidx])
	}

	afterKeyStr := strings.TrimSpace(keyvalueStr[eidx+1:])

	if len(afterKeyStr) > 0 {
		nextChar := afterKeyStr[0:1]
		if nextChar == "\"" || nextChar == "'" {
			//String value
			reststring := afterKeyStr[1:]
			endStrIdx := strings.Index(reststring, nextChar)
			value = reststring[:endStrIdx]
			rest = reststring[endStrIdx+1:]
			if strings.Index(rest, ",") == 0 {
				rest = rest[1:]
			}
		} else {
			cIdx := strings.Index(afterKeyStr, ",")
			//No value provide
			if cIdx == 0 {
				value = ""
				rest = afterKeyStr[1:]
			} else if cIdx < 0 {
				//no more properties
				value = afterKeyStr
				rest = ""
			} else {
				//more properties
				value = afterKeyStr[:cIdx]
				if cIdx < len(afterKeyStr) {
					rest = afterKeyStr[cIdx+1:]
				}
			}

		}
		m[key] = value
		if rest != "" {
			parseKeyValue(rest, m)
		}
	}
}

func removeQuote(quoteStr string) string {
	if (strings.HasPrefix(quoteStr, `"`) && strings.HasSuffix(quoteStr, `"`)) || (strings.HasPrefix(quoteStr, `'`) && strings.HasSuffix(quoteStr, `'`)) {
		quoteStr = quoteStr[1 : len(quoteStr)-1]
	}
	return quoteStr
}
