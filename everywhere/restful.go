package main

import (
	"everywhere/configs"
	"everywhere/models"
	"everywhere/routes"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/codegangsta/negroni"
)

func main() {

	// init db
	// configInfo := config.ConfigInfo{}
	// configInfo.Init()

	session, err := config.GetDBSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("everywhere").C("user")
	// err = c.Insert(&model.CoderModel{"Ale", 333},
	// 	&model.CoderModel{"Cla", 555})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// result := model.CoderModel{}
	var result model.CoderModel
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Age)

	// var sampleHandler = new(BaseHandler)
	var coderHandler = new(route.CoderHandler)
	// coderHandler.Init(&configInfo)

	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to the home page!")
	// })

	mux.HandleFunc("/", coderHandler.HandleRequest)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
