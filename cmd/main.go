package main

import (
	"fmt"

	"github.com/macgeargear/ichigo/cfgldr"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := cfgldr.LoadConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "backend").
			Msg("failed to start service")
	}

	fmt.Println(conf)
}
