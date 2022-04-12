package api

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type LuaSuite struct {
	suite.Suite
}

func (sut *LuaSuite) SetupTest() {

}

func (sut *LuaSuite) TestLoadLuas() {
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

	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	lua := &Lua{
		ConnectionID: id,
		ID:           id,
		Keys:         []string{},
		Args:         nil,
		Script:       incrByScript,
	}

	err := lua.New()
	sut.ErrorIs(err, nil)

	var luas = make([]Lua, 0)
	err = LoadLuas(&luas)
	sut.ErrorIs(err, nil)
	sut.GreaterOrEqual(len(luas), 1)

	err = lua.Delete()
	sut.ErrorIs(err, nil)
}

func TestLuaSuite(t *testing.T) {
	suite.Run(t, new(LuaSuite))
}
