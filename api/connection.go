package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/slowrookie/redis-web-manager/api/parser"
	"net"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/ssh"
)

const (
	connectionCollection = "connections"
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
	_client            redis.UniversalClient
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

var Connections = make([]*Connection, 0)

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
	return GlobalStorage.Write(connectionCollection, id, c)
}

// Edit .
func (c *Connection) Edit() error {
	return GlobalStorage.Write(connectionCollection, c.ID, c)
}

// Delete .
func (c *Connection) Delete() error {
	return GlobalStorage.Delete(connectionCollection, c.ID)
}

// Open .
func (c *Connection) Open() error {
	if nil == c._client {
		client, err := c.client()
		if nil != err {
			return err
		}
		c._client = client
	}
	_, err := c._client.Ping(context.Background()).Result()
	if nil != err {
		return err
	}
	return nil
}

// Disconnection .
func (c *Connection) Disconnection() error {
	if nil != c._client {
		err := c._client.Close()
		if nil != err {
			return err
		}
	}
	return nil
}

// Command .
func (c *Connection) Command(commands [][]interface{}) ([]interface{}, error) {
	size := len(commands)
	if size == 0 {
		return nil, nil
	}

	ctx := context.Background()
	pipe := c._client.Pipeline()
	retCmds := make([]*redis.Cmd, size)
	rets := make([]interface{}, size)

	for i := 0; i < size; i++ {
		retCmds[i] = pipe.Do(ctx, commands[i]...)
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

// Scripting lua script
func (c *Connection) Scripting(lua *Lua) error {
	start := time.Now()
	ctx := context.Background()
	ret, err := redis.NewScript(lua.Script).Run(ctx, c._client, lua.Keys, lua.Args...).Result()
	if nil != err {
		return err
	}
	lua.LastExecutionAt = start.UnixMilli()
	lua.Elapsed = fmt.Sprintf("%v", time.Since(start))
	lua.Result = ret
	return nil
}

func (c *Connection) Suggestions(command string) []string {
	input := antlr.NewInputStream(command)
	lexer := parser.NewRedisLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRedisParser(stream)
	p.BuildParseTrees = true

	redisErrorListener := &parser.SuggestionRedisErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(redisErrorListener)
	tree := p.Command()

	if nil == c || nil == c._client {
		return redisErrorListener.ExpectedTokens
	}

	suggestionVisitor := &parser.SuggestionRedisParserVisitor{
		ExpectedTokens: redisErrorListener.ExpectedTokens,
		Client:         c._client,
		Expects:        []string{},
	}

	suggestionVisitor.Visit(tree)

	if len(suggestionVisitor.Expects) <= 0 {
		return redisErrorListener.ExpectedTokens
	}
	return suggestionVisitor.Expects
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
		tlsConfig := tls.Config{
			Certificates:       []tls.Certificate{cer},
			InsecureSkipVerify: true,
		}
		if len(c.Tls.Ca) > 0 {
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM([]byte(c.Tls.Ca))
			tlsConfig.RootCAs = caCertPool
		}
		universalOptions.TLSConfig = &tlsConfig
	}

	// ssh enable
	if c.SSHOptions.Enable {
		cli, err := sshClient(c.SSHOptions)
		if nil != err {
			return nil, err
		}
		universalOptions.ReadTimeout = -2
		universalOptions.WriteTimeout = -2
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
			return con, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Connection %s not exists", id))
}

func LoadConnections() error {
	var bytes [][]byte
	err := GlobalStorage.ReadAll(connectionCollection, &bytes)
	if err != nil {
		return err
	}
	var _connections = make([]*Connection, 0)
	if err = RecordsToStruct(bytes, &_connections); err != nil {
		return err
	}
	for _, nc := range _connections {
		for _, c := range Connections {
			if nc.ID == c.ID {
				nc._client = c._client
			}
		}
	}
	Connections = _connections
	return nil
}
