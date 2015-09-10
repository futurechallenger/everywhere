package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

const (
	DBAddress = "localhost:127071"
)

/*
  ConfigInfo: configurations of this app
*/
type ConfigInfo struct {
}

// Init: initialize configuration
func (config *ConfigInfo) Init() {

}

/*
  ErrorJson
*/
type ErrorJson struct {
	State   string      `json:"state"`   // 0: correct, other values means there're something wrong.
	Message string      `json:"message"` // message: success or error
	Data    interface{} `json:"data"`    // if success, contains a dic or array, or null.
}

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
	coderModelList := []CoderModel{CoderModel{"Jack", 20, []string{"CN"}, []string{"Oobjective-C"}}}
	return coderModelList, nil
}

/*
 * SampleHandler: test http handler just
 */
type BaseHandler struct {
	Content string // json string to be sent
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
	var templateJson ErrorJson
	coderModelList, modelEerr := new(CoderModel).FindCoder(make(map[string]string))

	if modelEerr != nil {
		var errorJson = ErrorJson{"1", "DB error", nil}
		b, _ := json.Marshal(errorJson)
		coderHandler.BaseHandler.Content = string(b)
	} else {
		templateJson.Data = coderModelList
		coderJson, jsonErr := json.Marshal(templateJson)
		if jsonErr != nil {
			var errorJson = ErrorJson{State: "1", Message: "encode json error", Data: nil}
			b, _ := json.Marshal(errorJson)
			coderHandler.Content = string(b)
		}

		coderHandler.BaseHandler.Content = string(coderJson)
		coderHandler.BaseHandler.HandleRequest(w, req)
	}
}

func main() {
	// var sampleHandler = new(BaseHandler)
	var coderHandler = new(CoderHandler)

	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to the home page!")
	// })

	mux.HandleFunc("/", coderHandler.HandleRequest)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
