package commands

import (
	"errors"
	"nexis7/util"
	"os"
)

func ProcessGenerate() error {
	if _, err := os.Stat(util.PrivateKeyPath); err == nil {
		return errors.New("nexis.key file already exists. delete it, move it, or rename it if you want to generate a new one")
	} else {
		key := util.GenerateNewPrivateKey()
		encoded := key.ToBase64()

		f, err := os.Create(util.PrivateKeyPath)
		if err != nil {
			return err
		}

		defer f.Close()
		_, err = f.WriteString(encoded)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}
