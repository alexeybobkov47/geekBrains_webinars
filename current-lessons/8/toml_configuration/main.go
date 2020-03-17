package main

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type TOML struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

func init() {
	conf := new(TOML)
	if _, err := toml.DecodeFile("config.toml", conf); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", conf)
}

func main() {

}
