package main

import (
	"log"

	"github.com/joqd/slovo/internal/adapter/bootstrap"
	"github.com/joqd/slovo/internal/adapter/config"
)

// @title           Go Russian
// @version         1.0
// @description     Russian Dictionary & Learning Platform.
// @contact.name    Abolfazl Shahbazi
// @contact.url     https://github.com/joqd
// @contact.email   rodia2559@example.com
// @license.name    MIT License
// @license.url     https://opensource.org/licenses/MIT
// @schemes         http
func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	bootstrap.Run(conf)
}
