package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/slowrookie/redis-web-manager/api"
	"github.com/stretchr/testify/suite"
)

type AppSuite struct {
	suite.Suite
}

var app *App

func (sut *AppSuite) SetupTest() {
	app = NewApp()
}

func (sut *AppSuite) TestConnections() {
	connections := app.Connections()
	sut.GreaterOrEqual(len(connections), 1)
}

func (sut *AppSuite) TestOpenConnection() {
	var con *api.Connection
	for _, v := range app.Connections() {
		if v.Name == "localhost" {
			con = v
		}
	}
	connection, err := api.GetConnection(con.ID)
	sut.ErrorIs(err, nil)
	sut.NotEqual(connection, nil)
	err = connection.Open()
	sut.ErrorIs(err, nil)
	log.Println(fmt.Sprintf("%#v", connection))
	commands := [][]interface{}{{"ping"}}
	ret, err := connection.Command(commands)
	sut.NotEmpty(ret)
}

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}
