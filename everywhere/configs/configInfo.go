package config

import "gopkg.in/mgo.v2"

const (
	DBAddress = "localhost"
)

/*
  ConfigInfo: configurations of this app
*/
type ConfigInfo struct {
	Conn *mgo.Session
}

// Init: initialize configuration
func (config *ConfigInfo) Init() {
	conn, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	config.Conn = conn
}
