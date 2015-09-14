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
