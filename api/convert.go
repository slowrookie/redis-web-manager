package api

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"unicode"
)

type Convert struct {
	Data string `json:"data"`
}

var DefaultConvert = &Convert{}

// Length .
func (c *Convert) Length() int {
	return len([]rune(c.Data))
}

// ToHex .
func (c *Convert) ToHex() string {
	return hex.EncodeToString([]byte(c.Data))
}

// ToJson .
func (c *Convert) ToJson() string {
	var formatOut bytes.Buffer
	err := json.Indent(&formatOut, []byte(c.Data), "", "\t")
	if nil != err {
		return c.Data
	}
	return formatOut.String()
}

// ToBinary .
func (c *Convert) ToBinary() (bstr string) {
	for _, v := range c.Data {
		bstr = fmt.Sprintf("%s%b", bstr, v)
	}
	return
}

// Base64ToText .
func (c *Convert) Base64ToText() string {
	bts, err := base64.StdEncoding.DecodeString(c.Data)
	if nil != err {
		return c.Data
	}
	return string(bts)
}

// Base64ToJson .
func (c *Convert) Base64ToJson() string {
	txt := c.Base64ToText()
	c.Data = string(txt)
	return c.ToJson()
}

// banary
func (c *Convert) IsPrintable() bool {
	for _, r := range c.Data {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
