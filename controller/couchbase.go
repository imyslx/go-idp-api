package controller

import (
	"io/ioutil"
	"os"

	"github.com/couchbase/gocb"
	zlog "github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v2"
)

// CbEnvConfig : CouchbaseConfigs
type CbEnvConfig struct {
	Dev CbConfig `yaml:"development"`
	Stg CbConfig `yaml:"staging"`
	Prd CbConfig `yaml:"production"`
}

// CbConfig : Config to connect couchbase servers.
type CbConfig struct {
	Hosts    []string `yaml:"hosts"`
	Bucket   string   `yaml:"bucket"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
}

// GetCbConfig : Get config from yaml file.
func GetCbConfig(filename string) (CbConfig, error) {
	if filename == "" {
		filename = "conf/couchbase.yaml"
	}

	cec := new(CbEnvConfig)
	cc := new(CbConfig)

	yamlBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return *cc, err
	}

	err = yaml.Unmarshal(yamlBytes, cec)
	if err == nil {
		env := os.Getenv("IDP_ENV")
		switch env {
		case "staging":
			return cec.Stg, nil
		case "production":
			return cec.Prd, nil
		default:
			return cec.Dev, nil
		}
	}
	return *cc, err
}

// GetConnection : Get connection to couchbase.
func (cc CbConfig) GetConnection() *gocb.Bucket {

	hosts := ""
	for _, host := range cc.Hosts {
		hosts += "couchbase://" + host + ","
	}
	cluster, err := gocb.Connect(hosts)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Cannot connect to couchbase.")
	}
	zlog.Debug().Msg("Try to connect " + cc.User + ":" + cc.Password + "¥n")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: cc.User,
		Password: cc.Password,
	})
	bucket, err := cluster.OpenBucket(cc.Bucket, "")
	if err != nil {
		zlog.Fatal().Err(err).Msg("Cannot connect to bucket.")
	}
	return bucket
}