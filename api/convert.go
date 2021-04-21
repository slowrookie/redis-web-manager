package api

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"unicode"
)

type Convert struct{}

var DefaultConvert = &Convert{}

// Length .
func (c *Convert) Length(param string) int {
	return len([]rune(param))
}

// ToHex .
func (c *Convert) ToHex(param string) string {
	return hex.EncodeToString([]byte(param))
}

// ToJson .
func (c *Convert) ToJson(param string) string {
	var formatOut bytes.Buffer
	err := json.Indent(&formatOut, []byte(param), "", "\t")
	if nil != err {
		return param
	}
	return formatOut.String()
}

// ToBinary .
func (c *Convert) ToBinary(param string) (bstr string) {
	for _, v := range param {
		bstr = fmt.Sprintf("%s%b", bstr, v)
	}
	return
}

// Base64ToText .
func (c *Convert) Base64ToText(param string) string {
	bts, err := base64.StdEncoding.DecodeString(param)
	if nil != err {
		return param
	}
	return string(bts)
}

// Base64ToJson .
func (c *Convert) Base64ToJson(param string) string {
	txt := c.Base64ToText(param)
	return c.ToJson(string(txt))
}

// banary
func (c *Convert) IsPrintable(param string) bool {
	for _, r := range param {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
