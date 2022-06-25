package nube

import (
	"fmt"
	"regexp"
)

var (
	TOKEN = "^[A-Za-z0-9!\\$#%&'*+-\\.\\^_`|~]+$"
	SP = " "
)

// Path related regex
var (
	PATH_UNRESERVED_CHARS = `([A-Za-z\-\._~])`
	PATH_PCT_ENCODED = `(%[0-9A-Fa-f]{2})`
	PATH_SUB_DELIMS = `([!$&'()*+,;=])`
	PATH_PCHAR = fmt.Sprintf("(%s|%s|%s|:|@)", PATH_PCT_ENCODED, PATH_SUB_DELIMS, PATH_UNRESERVED_CHARS)
	PATH_SEGMENT = fmt.Sprintf("(%s)*", PATH_PCHAR)
	PATH_QUERY = fmt.Sprintf("(%s|/|\\?)*", PATH_PCHAR)
	PATH = fmt.Sprintf("(/%s)+(\\?%s)?(#%s)?", PATH_SEGMENT, PATH_QUERY, PATH_QUERY)
)

func IsValidToken(str string) bool {
	match, _ := regexp.Match(TOKEN, []byte(str))
	return match
}

func IsValidPath(str string) bool {
	match, _ := regexp.Match(PATH, []byte(str))
	return match
}