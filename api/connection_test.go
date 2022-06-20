package api

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ConnectionSuite struct {
	suite.Suite
}

func (sut *ConnectionSuite) SetupTest() {

}

func (sut *ConnectionSuite) TestOpen() {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	connection := &Connection{
		ID:   id,
		Host: "localhost",
		Port: 6379,
	}

	err := connection.Open()
	sut.ErrorIs(err, nil)
	sut.NotEqual(connection._client, nil)
}

func (sut *ConnectionSuite) TestConnection() {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	connection := &Connection{
		ID:   id,
		Host: "localhost",
		Port: 6379,
	}
	err := connection.Open()
	sut.ErrorIs(err, nil)
	commands := [][]interface{}{
		{"CONFIG", "GET", "*"},
		{"INFO"},
	}
	ret, err := connection.Command(commands)
	sut.ErrorIs(err, nil)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	retJson, err := json.Marshal(ret)
	sut.ErrorIs(err, nil)
	fmt.Println(string(retJson))
}

func (sut *ConnectionSuite) TestScripting() {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	connection := &Connection{
		ID:   id,
		Host: "localhost",
		Port: 6379,
	}

	incrByScript := `
	local key = KEYS[1]
	local change = ARGV[1]
	local value = redis.call("GET", key)
	if not value then
	  value = 0
	end
	value = value + change
	redis.call("SET", key, value)
	return value
	`
	// args
	var args = make([]interface{}, 0)
	args = append(args, 1)

	lua := &Lua{
		ConnectionID: id,
		ID:           id,
		Keys:         []string{"my_counter"},
		Args:         args,
		Script:       incrByScript,
	}

	err := connection.Scripting(lua)
	sut.ErrorIs(err, nil)
	sut.GreaterOrEqual(lua.Result, int64(1))
	fmt.Println(fmt.Sprintf("%#v", lua))
	fmt.Println(fmt.Sprintf("%+v", lua.Result))
}

func TestConnectionSuite(t *testing.T) {
	suite.Run(t, new(ConnectionSuite))
}
