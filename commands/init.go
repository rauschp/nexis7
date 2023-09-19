package commands

import (
	"errors"
	"os"
)

func ProcessInit() error {
	if _, err := os.Stat("keys/nexis.key"); err != nil {
		return errors.New("nexis.key required. Please run the generate command before running this command")
	}

	return nil
}
