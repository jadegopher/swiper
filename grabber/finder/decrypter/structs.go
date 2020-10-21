package decrypter

import "encoding/asn1"

type passwordItem struct {
	GlobalSalt []byte
	SecretASN1 []byte
}

type nssPrivate struct {
	a11  []byte
	a102 []byte
}

type asnSecretWithSalt struct {
	ObjectSeq objectSequence
	Data      []byte
}

type asnCredentials struct {
	CkaId         []byte
	IvSeq         ivSequence
	EncryptedData []byte
}

type ivSequence struct {
	ObjectIdentifier asn1.ObjectIdentifier
	InitVector       []byte
}

type objectSequence struct {
	ObjectIdentifier asn1.ObjectIdentifier
	SaltSeq          SaltSequence
}

type SaltSequence struct {
	Salt []byte
}
