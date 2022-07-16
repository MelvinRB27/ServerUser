package authorization

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/golang-jwt/jwt"
)

var (
	err error
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once sync.Once
)

func LoadFiles(privateFile, publicFile string) error {
	once.Do(func() {
		err = loadFiles(privateFile, publicFile)
	})
	return err
}


func loadFiles(privateFile, publicFile string) error {
	privateBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}

	publicBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}

	return parseRSA(privateBytes, publicBytes)
}


func parseRSA(privateBytes, publicBytes []byte) error {
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}