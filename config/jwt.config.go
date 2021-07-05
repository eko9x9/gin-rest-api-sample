package config

import (
	"io/ioutil"
	"os"
)

func GetJwtKey() ([]byte, error) {
	pwd, _ := os.Getwd()

	keyPath := pwd + "/config/secret/jwt.key"
	key, err := ioutil.ReadFile(keyPath)

	return key, err
}
