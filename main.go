package main

import (
	"fmt"
	"github.com/rauschp/nexis7/commands"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := processCommand(); err != nil {
		log.Error().Err(err).Msg("failed to process command")
	}
}

func processCommand() error {
	args := os.Args
	cmd := args[1]

	if len(args) < 2 {
		return fmt.Errorf("no command specified")
	}

	switch cmd {
	case "generate":
		err := commands.ProcessGenerate()
		if err != nil {
			return err
		}

		log.Info().Msg("Successfully generated private key")

		return nil
	case "init":
		err := commands.ProcessInit()
		if err != nil {
			return err
		}

		return nil
	case "validate":
		err := commands.ProcessValidate()
		if err != nil {
			return err
		}

		return nil
	case "run":
		err := commands.ProcessRunCommand()
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("unknown command %s", cmd)
}
