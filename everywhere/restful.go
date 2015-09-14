package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
 * SampleHandler: test http handler just
 */
type BaseHandler struct {
	ConfigInfo *ConfigInfo
	Content    string // json string to be sent
}

func (baseHandler *BaseHandler) Init(config *ConfigInfo) {
	baseHandler.ConfigInfo = config
}

// HandleRequest
func (baseHandler *BaseHandler) HandleRequest(w http.ResponseWriter, req *http.Request) {
	// if Content == nil {
	// 	Content = ""
	// }
	fmt.Fprintf(w, baseHandler.Content)
}

/*
  CoderHandler: this is a coder handler
*/
type CoderHandler struct {
	BaseHandler
}

// HandleRequest,when request comes use this method to deal with it
func (coderHandler *CoderHandler) HandleRequest(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[E]", r)

			// error
			var errorJson = TemplateJson{"500", "Inner error", nil}
			b, _ := json.Marshal(errorJson)
			coderHandler.BaseHandler.Content = string(b)
			coderHandler.BaseHandler.HandleRequest(w, req)
		}
	}()

	var templateJson TemplateJson
	coderModelList, modelEerr := new(CoderModel).FindCoder(coderHandler.ConfigInfo.Conn, make(map[string]string))

	if modelEerr != nil {
		var errorJson = TemplateJson{"1", "DB error", nil}
		b, _ := json.Marshal(errorJson)
		coderHandler.BaseHandler.Content = string(b)
	} else {
		templateJson.Data = coderModelList
		coderJson, jsonErr := json.Marshal(templateJson)
		if jsonErr != nil {
			var errorJson = TemplateJson{State: "1", Message: "encode json error", Data: nil}
			b, _ := json.Marshal(errorJson)
			coderHandler.Content = string(b)
		}

		coderHandler.BaseHandler.Content = string(coderJson)
		coderHandler.BaseHandler.HandleRequest(w, req)
	}
}

func main() {
	fmt.Println("Hello world!")
	// init db
	configInfo := ConfigInfo{}
	configInfo.Init()

	conn, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := conn.DB("everywhere").C("user")
	err = c.Insert(&CoderModel{"Ale", 123, []string{"EN"}, []string{"C"}},
		&CoderModel{"Cla", 123, []string{"RU"}, []string{"Objective-C"}})
	if err != nil {
		log.Fatal(err)
	}

	result := CoderModel{}
	err = c.Find(bson.M{"name": "Bruce"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello world!")
	fmt.Println("Age:", result.Age)

	// var sampleHandler = new(BaseHandler)
	var coderHandler = new(CoderHandler)
	coderHandler.Init(&configInfo)

	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to the home page!")
	// })

	mux.HandleFunc("/", coderHandler.HandleRequest)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
