package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
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
	Username           string `json:"username"`
	Auth               string `json:"auth"`
	KeysPattern        string `json:"keysPattern"`
	NamespaceSeparator string `json:"namespaceSeparator"`
	TimeoutConnect     int64  `json:"timeoutConnect"`
	TimeoutExecute     int64  `json:"timeoutExecute"`
	DbScanLimit        int32  `json:"dbScanLimit"`
	DataScanLimit      int32  `json:"dataScanLimit"`
	Tls                Tls    `json:"tls"`
}

type Tls struct {
	Enable bool   `json:"enable"`
	Cert   string `json:"cert"`
	Key    string `json:"key"`
	Ca     string `json:"ca"`
}

// Command .
type Command struct {
	ID       string          `json:"id"`
	Commands [][]interface{} `json:"commands"`
}

var DefaultConnection = &Connection{}

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
func (c *Connection) Test() error {
	client, err := c.client(*c)
	if nil != err {
		return err
	}
	_, err = client.Ping(context.Background()).Result()
	if nil != err {
		return err
	}
	return nil
}

// New .
func (c *Connection) New() (Connection, error) {
	if len(c.ID) == 0 {
		c.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
		connections = append([]Connection{*c}, connections...)
	} else {
		for i := 0; i < len(connections); i++ {
			if connections[i].ID == c.ID {
				connections[i] = *c
				break
			}
		}
	}

	client, err := c.client(*c)
	if nil != err {
		return *c, err
	}

	// write file
	if err := c.writeConnections(); nil != err {
		return *c, err
	}

	Clients[c.ID] = client
	return *c, nil
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
		_client, err := c.client(*c.findByID(id))
		if nil != err {
			return nil, err
		}
		Clients[id] = _client
		client = _client
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
func (c *Connection) Copy() (Connection, error) {
	c.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	connections = append(connections, *c)
	if err := c.writeConnections(); nil != err {
		return *c, err
	}
	return *c, nil
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
		client, err := c.client(*c.findByID(cmd.ID))
		if nil != err {
			return nil, err
		}
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

func (c *Connection) client(connection Connection) (*redis.Client, error) {
	// options
	options := redis.Options{
		Addr:        fmt.Sprintf("%s:%d", connection.Host, connection.Port),
		Password:    connection.Auth,
		ReadTimeout: time.Duration(connection.TimeoutExecute) * time.Millisecond,
		DialTimeout: time.Duration(connection.TimeoutConnect) * time.Millisecond,
	}

	// tls enable
	if c.Tls.Enable {
		cer, err := tls.X509KeyPair([]byte(c.Tls.Cert), []byte(c.Tls.Key))
		if nil != err {
			return nil, err
		}
		tls := tls.Config{
			Certificates:       []tls.Certificate{cer},
			InsecureSkipVerify: true,
		}
		if len(c.Tls.Ca) > 0 {
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM([]byte(c.Tls.Ca))
			tls.RootCAs = caCertPool
		}
		options.TLSConfig = &tls
	}

	return redis.NewClient(&options), nil
}

func (c *Connection) loadConnections() error {
	err := os.MkdirAll(ROOT_PATH, os.ModePerm)
	if nil != err {
		return err
	}
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
