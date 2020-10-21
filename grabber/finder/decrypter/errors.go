package decrypter

import "errors"

var (
	WrongMasterPasswordErr = errors.New("wrong master password")
	ProtocolSupportErr     = errors.New("crypto protocol not supported")
)
