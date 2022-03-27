package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/ssh"
)

const (
	ConnectionC = "connections"
)

// ConnectionEntry .
type ConnectionEntry struct {
	ID         string
	Connection string
}

// Connection .
type Connection struct {
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	Host               string     `json:"host"`
	Port               int        `json:"port"`
	Addrs              []string   `json:"addrs"`
	Username           string     `json:"username"`
	Auth               string     `json:"auth"`
	KeysPattern        string     `json:"keysPattern"`
	NamespaceSeparator string     `json:"namespaceSeparator"`
	TimeoutConnect     int64      `json:"timeoutConnect"`
	TimeoutExecute     int64      `json:"timeoutExecute"`
	DbScanLimit        int32      `json:"dbScanLimit"`
	DataScanLimit      int32      `json:"dataScanLimit"`
	Tls                Tls        `json:"tls"`
	SSHOptions         SSHOptions `json:"ssh"`
	IsCluster          bool       `json:"isCluster"`
	IsSentinel         bool       `json:"isSentinel"`
	SentinelPassword   string     `json:"sentinelPassword"`
	MasterName         string     `json:"masterName"`
	RouteByLatency     bool       `json:"routeByLatency"`
	RouteRandomly      bool       `json:"routeRandomly"`
}

type Tls struct {
	Enable bool   `json:"enable"`
	Cert   string `json:"cert"`
	Key    string `json:"key"`
	Ca     string `json:"ca"`
}

type SSHOptions struct {
	Enable     bool   `json:"enable"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Privatekey string `json:"privatekey"`
}

// Command .
type Command struct {
	ID       string          `json:"id"`
	Commands [][]interface{} `json:"commands"`
}

var Connections = make([]Connection, 0)

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
	return GlobalStorage.Write(ConnectionC, id, c)
}

// Edit .
func (c *Connection) Edit() error {
	return GlobalStorage.Write(ConnectionC, c.ID, c)
}

// Delete .
func (c *Connection) Delete() error {
	return GlobalStorage.Delete(ConnectionC, c.ID)
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

	// ssh enable
	if c.SSHOptions.Enable {
		cli, err := sshClient(c.SSHOptions)
		if nil != err {
			return nil, err
		}
		universalOptions.ReadTimeout = -1
		universalOptions.WriteTimeout = -1
		universalOptions.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return cli.Dial(network, addr)
		}
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

// sshClient create ssh client
func sshClient(opts SSHOptions) (*ssh.Client, error) {
	addr := net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))

	sshConfig := &ssh.ClientConfig{
		User:            opts.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         15 * time.Second,
	}

	if len(opts.Password) != 0 && len(opts.Privatekey) != 0 {
		signer, err := ssh.ParsePrivateKeyWithPassphrase([]byte(opts.Privatekey), []byte(opts.Password))
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else if len(opts.Privatekey) != 0 {
		signer, err := ssh.ParsePrivateKey([]byte(opts.Privatekey))
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	} else if len(opts.Password) != 0 {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(opts.Password)}
	}

	return ssh.Dial("tcp", addr, sshConfig)
}

func GetConnection(id string) (*Connection, error) {
	for _, con := range Connections {
		if con.ID == id {
			return &con, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Connection %s not exists", id))
}

func LoadConnections() error {
	var byts [][]byte
	err := GlobalStorage.ReadAll(ConnectionC, &byts)
	if err != nil {
		return err
	}
	var _connections = make([]Connection, 0)
	if err = RecordsToStruct(byts, &_connections); err != nil {
		return err
	}
	Connections = _connections
	return nil
}
