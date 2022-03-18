package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	SelectConnectionStatement     = "SELECT `id`, `connection` FROM `rwm_connection`"
	SelectConnectionByIdStatement = "SELECT `id`, `connection` FROM `rwm_connection` WHERE `id` = $1"
	InsertConnectionStatement     = "INSERT INTO `rwm_connection` (`id`, `connection`) VALUES ($1, $2)"
	UpdateConnectionStatement     = "UPDATE `rwm_connection` SET `connection` = $1 WHERE `id` = $2"
	DeleteConnectionStatement     = "DELETE FROM `rwm_connection` WHERE `id` = $1"
)

// ConnectionEntry .
type ConnectionEntry struct {
	ID         string
	Connection string
}

// Connection .
type Connection struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Host               string   `json:"host"`
	Port               int32    `json:"port"`
	Addrs              []string `json:"addrs"`
	Username           string   `json:"username"`
	Auth               string   `json:"auth"`
	KeysPattern        string   `json:"keysPattern"`
	NamespaceSeparator string   `json:"namespaceSeparator"`
	TimeoutConnect     int64    `json:"timeoutConnect"`
	TimeoutExecute     int64    `json:"timeoutExecute"`
	DbScanLimit        int32    `json:"dbScanLimit"`
	DataScanLimit      int32    `json:"dataScanLimit"`
	Tls                Tls      `json:"tls"`
	IsCluster          bool     `json:"isCluster"`
	IsSentinel         bool     `json:"isSentinel"`
	SentinelPassword   string   `json:"sentinelPassword"`
	MasterName         string   `json:"masterName"`
	RouteByLatency     bool     `json:"routeByLatency"`
	RouteRandomly      bool     `json:"routeRandomly"`
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

// Clients .
var Clients = make(map[string]redis.UniversalClient)

// Test .
func (c *Connection) Test() error {
	client, err := c.client()
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
func (c *Connection) New() error {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	c.ID = id

	bts, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if _, err := DB.Exec(InsertConnectionStatement, id, string(bts)); err != nil {
		return err
	}
	return nil
}

// Edit .
func (c *Connection) Edit() error {
	bts, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if _, err := DB.Exec(UpdateConnectionStatement, string(bts), c.ID); err != nil {
		return err
	}
	return nil
}

// Delete .
func (c *Connection) Delete() error {
	if _, err := DB.Exec(DeleteConnectionStatement, c.ID); err != nil {
		return err
	}
	return nil
}

// Open .
func (c *Connection) Open() error {
	if _, exists := Clients[c.ID]; !exists {
		client, err := c.client()
		if nil != err {
			return err
		}
		Clients[c.ID] = client
	}
	_, err := Clients[c.ID].Ping(context.Background()).Result()
	if nil != err {
		return err
	}
	return nil
}

// Disconnection .
func (c *Connection) Disconnection() error {
	if _, exists := Clients[c.ID]; exists {
		err := Clients[c.ID].Close()
		delete(Clients, c.ID)
		if nil != err {
			return err
		}
	}
	return nil
}

// Command
func (c *Connection) Command(cmd Command) ([]interface{}, error) {
	size := len(cmd.Commands)
	if size == 0 {
		return nil, nil
	}

	if _, exists := Clients[cmd.ID]; !exists {
		client, err := c.client()
		if nil != err {
			return nil, err
		}
		Clients[cmd.ID] = client
	}

	ctx := context.Background()
	pipe := Clients[cmd.ID].Pipeline()
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

func (c *Connection) client() (redis.UniversalClient, error) {
	// options
	universalOptions := &redis.UniversalOptions{
		Addrs:       []string{fmt.Sprintf("%s:%d", c.Host, c.Port)},
		Username:    c.Username,
		Password:    c.Auth,
		ReadTimeout: time.Duration(c.TimeoutExecute) * time.Millisecond,
		DialTimeout: time.Duration(c.TimeoutConnect) * time.Millisecond,
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
		universalOptions.TLSConfig = &tls
	}

	// cluster
	if c.IsSentinel {
		universalOptions.SentinelPassword = c.SentinelPassword
		universalOptions.MasterName = c.MasterName
	}

	// cluster & sentinel
	if c.IsCluster || c.IsSentinel {
		universalOptions.Addrs = c.Addrs
		universalOptions.RouteByLatency = c.RouteByLatency
		universalOptions.RouteRandomly = c.RouteRandomly
	}

	return redis.NewUniversalClient(universalOptions), nil
}

func FindConnectionByID(id string) (*Connection, error) {
	var entry ConnectionEntry
	if err := DB.QueryRow(SelectConnectionByIdStatement, id).Scan(&entry.ID, &entry.Connection); err != nil {
		return nil, err
	}

	connection := Connection{}
	if err := json.Unmarshal([]byte(entry.Connection), &connection); err != nil {
		return nil, err
	}

	connection.ID = entry.ID

	return &connection, nil
}

// Connections get all connections
func Connections() ([]Connection, error) {
	rows, err := DB.Query(SelectConnectionStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	connections := []Connection{}
	for rows.Next() {
		entry := ConnectionEntry{}
		if err = rows.Scan(&entry.ID, &entry.Connection); err != nil {
			return nil, err
		}
		connection := Connection{ID: entry.ID}
		if err := json.Unmarshal([]byte(entry.Connection), &connection); err != nil {
			return nil, err
		}
		connections = append(connections, connection)
	}

	return connections, nil
}
