package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
)

type Config struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
	Port     uint   `json:"port"`
}

var DefaultConfig = &Config{}

func (c *Config) Get() (Config, error) {
	err := os.MkdirAll(ROOT_PATH, os.ModePerm)
	if nil != err {
		return *c, err
	}
	jsonFile, err := os.OpenFile(ConfigFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if nil != err {
		return *c, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if nil != err {
		return *c, err
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, c)
		if nil != err {
			return *c, err
		}
	}
	if c.Port == 0 {
		c.Port = 8080
	}
	return *c, nil
}

func (c *Config) Set() error {
	jsonFile, err := os.OpenFile(ConfigFilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if nil != err {
		return err
	}
	defer jsonFile.Close()

	bts, err := json.Marshal(c)
	if nil != err {
		return err
	}
	var formatOut bytes.Buffer
	err = json.Indent(&formatOut, bts, "", "\t")
	if nil != err {
		return err
	}
	_, err = jsonFile.Write(formatOut.Bytes())
	if nil != err {
		return err
	}
	return nil
}

func (c *Config) CheckPort(port string) error {
	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return err
	}

	defer ln.Close()
	return nil
}
