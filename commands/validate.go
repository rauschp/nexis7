package commands

import (
	"github.com/rauschp/nexis7/util"
	"github.com/rs/zerolog/log"
	"os"
)

func ProcessValidate() error {
	f, err := os.ReadFile(util.DefaultPrivateKeyPath)
	if err != nil {
		return err
	}

	contents := string(f)

	key, err := util.GeneratePrivateKeyFromBase64(contents)
	if err != nil {
		return err
	}

	p := key.Public().GetAddress().ToString()

	log.Info().Msg("Successfully validated private key")
	log.Info().Msgf("Nexis Address: %s", p)

	return nil
}
