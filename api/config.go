package api

import (
	"encoding/json"
)

type Config struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

const (
	SelectConfigStatement = "SELECT `config` FROM `rwm_config`"
	InsertConfigStatement = "INSERT INTO `config` VALUES ($1)"
	CountConfigStatement  = "SELECT COUNT(*) FROM `rwm_config`"
	UpdateConfigStatement = "UPDATE `rwm_config` SET `config` = $1"
)

var DefaultConfig = &Config{}

func (c *Config) Get() (Config, error) {
	var config string
	if err := DB.QueryRow(SelectConfigStatement).Scan(&config); err != nil {
		return *c, err
	}

	if len(config) > 0 {
		err := json.Unmarshal([]byte(config), c)
		if nil != err {
			return *c, err
		}
	}

	return *c, nil
}

func (c *Config) Set() error {
	bts, err := json.Marshal(c)
	if nil != err {
		return err
	}

	var count int
	if err := DB.QueryRow(CountConfigStatement).Scan(&count); err != nil {
		return err
	}

	if count >= 0 {
		_, err = DB.Exec(UpdateConfigStatement, string(bts))
	} else {
		_, err = DB.Exec(InsertConfigStatement, string(bts))
	}
	if nil != err {
		return err
	}
	return nil
}
