/*
Package config provides a function FiberConfig that returns a Fiber
configuration with predefined settings.
*/
package config

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

// The function initializes a new instance of Fiber configuration
func FiberConfig() fiber.Config {

	fiberConf := fiber.Config{
		AppName:     "Reddit Article Scraper",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}

	return fiberConf
}
