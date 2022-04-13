package api

import (
	"fmt"
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
