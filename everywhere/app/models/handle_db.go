package models

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBSession struct {
	DBName string
	collectionName string
}

func query(this *DBSession) (collectionList: map[string]string, conditions: map[string]string) map[string]interface{} {
	
}