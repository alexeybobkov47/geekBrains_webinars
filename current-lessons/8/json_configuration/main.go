package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/go-playground/validator.v8"
)

type Configuration struct {
	SiteName string `json:"site_name" check:"required"`
	DBLink   string `json:"db_link" check:"required"`
	LogFile  string `json:"log_file" check:"required"`
}

func init() {
	bytes, err := ioutil.ReadFile("app.json")
	if err != nil {
		log.Fatal(err)
	}

	Config := new(Configuration)
	if err = json.Unmarshal(bytes, Config); err != nil {
		log.Fatal(err)
	}

	if err := validate.Struct(Config); err != nil {
		log.Printf("try to load configuration: %v", err)
		Config.useDefault()
		log.Print("use default")
	}

	log.Printf("%+v", Config)
}

var validate = validator.New(&validator.Config{
	TagName: "check",
})

func (c *Configuration) useDefault() {
	*c = Configuration{
		SiteName: "Списки задач",
		DBLink:   "mongodb://localhost:27017/",
		LogFile:  "app.log",
	}
}

func main() {}
