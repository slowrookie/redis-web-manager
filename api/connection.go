package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// Connection .
type Connection struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Host               string `json:"host"`
	Port               int32  `json:"port"`
	Auth               string `json:"auth"`
	KeysPattern        string `json:"keysPattern"`
	NamespaceSeparator string `json:"namespaceSeparator"`
	TimeoutConnect     int64  `json:"timeoutConnect"`
	TimeoutExecute     int64  `json:"timeoutExecute"`
	DbScanLimit        int32  `json:"dbScanLimit"`
	DataScanLimit      int32  `json:"dataScanLimit"`
}

// Command .
type Command struct {
	ID       string          `json:"id"`
	Commands [][]interface{} `json:"commands"`
}

// Clients .
var Clients = make(map[string]*redis.Client)

var connections = make([]Connection, 0, 10)

// Connections get all connections
func (c *Connection) Connections() ([]Connection, error) {
	if err := c.loadConnections(); nil != err {
		return connections, err
	}
	return connections, nil
}

// Test .
func (c *Connection) Test(connection Connection) error {
	client := c.client(connection)
	_, err := client.Ping(context.Background()).Result()
	if nil != err {
		return err
	}
	return nil
}

// New .
func (c *Connection) New(connection Connection) ([]interface{}, error) {
	if len(connection.ID) == 0 {
		connection.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
		connections = append([]Connection{connection}, connections...)
	} else {
		for i := 0; i < len(connections); i++ {
			if connections[i].ID == connection.ID {
				connections[i] = connection
				break
			}
		}
	}

	client := c.client(connection)
	conf, err := client.ConfigGet(context.Background(), "databases").Result()
	if nil != err {
		return nil, err
	}

	// write file
	if err := c.writeConnections(); nil != err {
		return nil, err
	}

	Clients[connection.ID] = client
	return conf, nil
}

// Delete .
func (c *Connection) Delete(id string) error {
	for i := 0; i < len(connections); i++ {
		if connections[i].ID == id {
			connections = append(connections[:i], connections[i+1:]...)
			if _, exists := Clients[id]; exists {
				Clients[id].Close()
				delete(Clients, id)
			}
			break
		}
	}

	if err := c.writeConnections(); nil != err {
		return err
	}
	return nil
}

// Open .
func (c *Connection) Open(id string) (map[string]interface{}, error) {
	var client *redis.Client
	if _, exists := Clients[id]; exists {
		client = Clients[id]
	} else {
		client = c.client(*c.findByID(id))
		Clients[id] = client
	}
	ctx := context.Background()
	databases, err := client.ConfigGet(ctx, "databases").Result()
	if nil != err {
		return nil, err
	}
	info, err := client.Info(ctx).Result()
	if nil != err {
		return nil, err
	}
	var ret = make(map[string]interface{})
	ret["database"] = databases
	ret["info"] = info
	return ret, nil
}

// Disconnection .
func (c *Connection) Disconnection(id string) error {
	if _, exists := Clients[id]; exists {
		err := Clients[id].Close()
		delete(Clients, id)
		if nil != err {
			return err
		}
	}
	return nil
}

// Copy .
func (c *Connection) Copy(connection Connection) error {
	connection.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	connections = append(connections, connection)
	if err := c.writeConnections(); nil != err {
		return err
	}
	return nil
}

// Command
func (c *Connection) Command(cmd Command) ([]interface{}, error) {
	size := len(cmd.Commands)
	if size == 0 {
		return nil, nil
	}

	var client *redis.Client
	if _, exists := Clients[cmd.ID]; exists {
		client = Clients[cmd.ID]
	} else {
		client = c.client(*c.findByID(cmd.ID))
		Clients[cmd.ID] = client
	}
	ctx := context.Background()
	pipe := client.Pipeline()
	retCmds := make([]*redis.Cmd, size)
	rets := make([]interface{}, size)

	for i := 0; i < len(cmd.Commands); i++ {
		retCmds[i] = pipe.Do(ctx, cmd.Commands[i]...)
	}
	if _, err := pipe.Exec(ctx); nil != err {
		return nil, err
	}
	for i := 0; i < len(retCmds); i++ {
		ret, err := retCmds[i].Result()
		if nil != err {
			return nil, err
		}
		rets[i] = ret
	}
	return rets, nil
}

func (c *Connection) findByID(id string) *Connection {
	for i := 0; i < len(connections); i++ {
		if connections[i].ID == id {
			return &connections[i]
		}
	}
	return nil
}

func (c *Connection) client(connection Connection) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", connection.Host, connection.Port),
		Password:    connection.Auth,
		ReadTimeout: time.Duration(connection.TimeoutExecute) * time.Millisecond,
		DialTimeout: time.Duration(connection.TimeoutConnect) * time.Millisecond,
	})
}

func (c *Connection) loadConnections() error {
	jsonFile, err := os.OpenFile(ConnectionsFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if nil != err {
		return err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if nil != err {
		return err
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &connections)
		if nil != err {
			return err
		}
	}
	return nil
}

func (c *Connection) writeConnections() error {
	jsonFile, err := os.OpenFile(ConnectionsFilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if nil != err {
		return err
	}
	defer jsonFile.Close()

	bts, err := json.Marshal(connections)
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
