package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/eric2918/leaf-cluster/log"
)

var Server struct {
	LogLevel string
	LogPath  string

	Region     string
	MachineID  int64
	ServerName string
	WSAddr     string
	CertFile   string
	KeyFile    string
	TCPAddr    string
	MaxConnNum int
	Language   string

	ListenAddr      string
	ConnAddrs       map[string]string
	PendingWriteNum int

	ConsolePort   int
	ConsolePrompt string
	ProfilePath   string

	MongodbAddr       string
	MongodbSessionNum int

	RedisAddr     string
	RedisPassword string
	RedisDb       int

	JwtSecret  string
	JwtTimeout int
}

func init() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatal("os args of len(%v) less than 2", argsLen)
	}

	filePath := os.Args[1]
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
