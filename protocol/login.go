package protocol

import (
	"errors"
)
var (
	ErrorOutOfDateServer = errors.New("Server out of date")
	ErrorOutOfDateClient = errors.New("Client out of date")
	ErrorVerifyFailed    = errors.New("Verify Token incorrect")
	ErrorEncryption      = errors.New("Encryption error")
)

type Authenticator interface {
	Authenticate(username string, serverId string, sharedSecret []byte, publicKey []byte) (uuid string, err error)
}

func Login() (username string, uuid string, err error){

}