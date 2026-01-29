package main

import (
	"github.com/adhikag24/marketplace-booking/cmd/course/commands"

	"github.com/rs/zerolog/log"
)

func main() {
	err := commands.NewCommand().Execute()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to start server")
	}
}
