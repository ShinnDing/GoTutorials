// pkg/errors example
// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch05/05_01/errors.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Config holds configuration
type Config struct {
	// configuration fields go here (redacted)
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}
	defer file.Close()

	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil

}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}

	// Normal operation (redacted)
	fmt.Println(cfg)
}
