package main

import (
	"log"

	. "github.com/alimy/mir/v2/core"
	. "github.com/alimy/mir/v2/engine"

	_ "github.com/alimy/alexandrite/mirc/routes"
	_ "github.com/alimy/alexandrite/mirc/routes/v1"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialMode),
		GeneratorName(GeneratorMux),
		Cleanup(true),
		SinkPath("./gen"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
