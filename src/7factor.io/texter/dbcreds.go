package texter

import (
	"os"
	"strings"
	"time"

	mysql "github.com/go-sql-driver/mysql"
	mgo "gopkg.in/mgo.v2"
)

// DBCreds stores credential information about a database. Note for Mongo
// assume that Hosts are stored in the environment as a comma separated list.
type DBCreds struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

// ToDialInfo translates a DBCreds object to a DialInfo object
func (creds DBCreds) ToDialInfo() *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs:    strings.Split(creds.Host, ","),
		Database: creds.DBName,
		Username: creds.User,
		Password: creds.Pass,
	}
}

// ToDSN converts a creds object into a DSN
func (creds DBCreds) ToDSN() string {
	config := mysql.Config{
		User:    creds.User,
		Passwd:  creds.Pass,
		DBName:  creds.DBName,
		Timeout: 30 * time.Second,
		Addr:    creds.Host,
		Net:     "tcp",
	}
	return config.FormatDSN()
}

// GetMongoCredsFromEnv grabs all the appropriate creds from the
// environment and creates an object for it.
func GetMongoCredsFromEnv() DBCreds {
	return DBCreds{
		Host:   os.Getenv("MONGO_URL"),
		User:   os.Getenv("MONGO_USER"),
		Pass:   os.Getenv("MONGO_PASS"),
		DBName: os.Getenv("MONGO_DB"),
	}
}

// GetCongressDBCredsFromEnv grabs all the appropriate creds from the
// environment and creates an object for it.
func GetCongressDBCredsFromEnv() DBCreds {
	return DBCreds{
		Host:   os.Getenv("CNG_URL"),
		User:   os.Getenv("CNG_USER"),
		Pass:   os.Getenv("CNG_PASS"),
		DBName: os.Getenv("CNG_DBNAME"),
	}
}
