package dotenvdecoder

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func LoadENV() error {
	file, err := ioutil.ReadFile(".env")
	if err != nil {
		return errors.New("Err: Can't load .env file")
	}

	for _, l := range strings.Split(string(file), "\n") {
		words := strings.Split(l, `="`)
		key := words[0]
		value := words[len(words) - 1]

		os.Setenv(key, value)
	}

	return nil
}