package controller

import (
	"io/ioutil"
	"os"

	"github.com/couchbase/gocb"
	"github.com/imyslx/go-idp-api/app"
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

// GetCbBucket : Get config from yaml file.
func GetCbBucket(filename string) gocb.Bucket {

	if filename == "" {
		filename = "conf/couchbase.yaml"
	}

	cec := new(CbEnvConfig)
	cc := new(CbConfig)

	yamlBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Could not get config file.")
	}

	err = yaml.Unmarshal(yamlBytes, cec)
	if err == nil {
		env := os.Getenv("IDP_ENV")
		switch env {
		case "staging":
			cc = &cec.Stg
		case "production":
			cc = &cec.Prd
		default:
			cc = &cec.Dev
		}
	} else {
		zlog.Fatal().Err(err).Msg("Could not unmarshal config file.")
	}

	bucket := cc.GetConnection()

	return *bucket
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
	zlog.Debug().Msg("Try to connect " + cc.User + ":" + cc.Password)
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: cc.User,
		Password: cc.Password,
	})
	bucket, err := cluster.OpenBucket(cc.Bucket, "")
	if err != nil {
		zlog.Fatal().Err(err).Msg("Cannot connect to bucket.")
	}
	zlog.Debug().Msg("Success to Connect !")
	return bucket
}

// ExecuteQuery : Execute the query.
func ExecuteQuery(params *app.HostsPayload, query string) gocb.QueryResults {

	// Get connect to couchbase bucket.
	bucket := GetCbBucket("")

	// Execute
	rows, err := bucket.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), nil)
	if err != nil {
		zlog.Error().Err(err).Msg("Could not query in N1QL.")
	}

	return rows
}
