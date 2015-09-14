package model

import (
	"everywhere/configs"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
  CoderModel
*/
type CoderModel struct {
	Name            string
	Age             int
	SpeakLang       []string
	ProgrammingLang []string
}

// FindCoder find coder models by some condition
func (coderModel *CoderModel) FindCoder(condition map[string]string) ([]CoderModel, error) {
	var session *mgo.Session
	var err *error
	config.GetDBSession(session, err)

	var coderModelList = make([]CoderModel, 20, 20)
	var collection = session.DB("everywhere").C("user")
	var iter = collection.Find(bson.M{"name": "Bruce"}).Iter()

	var model CoderModel
	for iter.Next(&model) {
		fmt.Printf("Result: %v\n", model.Name)
		coderModelList = append(coderModelList, model)
	}
	// coderModelList := []CoderModel{CoderModel{"Jack", 20, []string{"CN"}, []string{"Oobjective-C"}}}
	// coderModelList = models
	return coderModelList, nil
}
