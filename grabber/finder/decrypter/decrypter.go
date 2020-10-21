package decrypter

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/sha1"
	"database/sql"
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"swiper/models"
)

type decrypt struct {
}

func New() *decrypt {
	return &decrypt{}
}

func (d *decrypt) Decrypt(db *sql.DB, data models.Login, masterPassword []byte) (models.Login, error) {
	rows, err := db.Query("SELECT item1,item2 FROM metadata WHERE id = 'password';")
	if err != nil {
		return models.Login{}, err
	}
	p := passwordItem{}
	if rows.Next() {
		if err := rows.Scan(&p.GlobalSalt, &p.SecretASN1); err != nil {
			return models.Login{}, err
		}
	}

	//Checking if master password valid
	pwdCheck, err := d.decryptSecret(masterPassword, p)
	if err != nil {
		if err == WrongMasterPasswordErr {
			return data, nil
		}
		return models.Login{}, err
	}
	if string(pwdCheck) != passwordCheck {
		return models.Login{}, WrongMasterPasswordErr
	}

	//Getting secret for key encryption
	rows, err = db.Query("SELECT a11,a102 FROM nssPrivate;")
	if err != nil {
		return models.Login{}, err
	}
	nss := nssPrivate{}
	if rows.Next() {
		if err := rows.Scan(&nss.a11, &nss.a102); err != nil {
			return models.Login{}, err
		}
	}
	if hex.EncodeToString(nss.a102) != ckaId {
		return models.Login{}, ProtocolSupportErr
	}

	//Key encryption
	p.SecretASN1 = nss.a11
	key, err := d.decryptSecret(masterPassword, p)
	if err != nil {
		return models.Login{}, err
	}
	key = key[:24]

	//Getting secret data
	username, err := d.decryptData(key, data.EncryptedUsername)
	if err != nil {
		return models.Login{}, err
	}
	password, err := d.decryptData(key, data.EncryptedPassword)
	if err != nil {
		return models.Login{}, err
	}

	data.UsernameField = string(username)
	data.PasswordField = string(password)
	return data, nil
}

func (d *decrypt) decryptData(key []byte, container string) ([]byte, error) {
	raw, err := d.fromBase64(container)
	if err != nil {
		return nil, err
	}
	asn := asnCredentials{}
	if _, err := asn1.Unmarshal(raw, &asn); err != nil {
		return nil, err
	}
	if asn.IvSeq.ObjectIdentifier.String() == tripleDesCbc {
		data, err := d.tripleDesDecrypt(key, asn.IvSeq.InitVector, asn.EncryptedData)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		return nil, ProtocolSupportErr
	}
}

func (d *decrypt) decryptPbeWithSha1andTripleDesCbc(globalSalt, entrySalt, masterPassword,
	cipherText []byte) ([]byte, error) {
	hp := sha1.Sum(append(globalSalt, masterPassword...))
	chp := sha1.Sum(append(hp[:], entrySalt...))

	k1, err := d.hmac(sha1.New, chp[:], append(entrySalt, entrySalt...))
	if err != nil {
		return nil, err
	}
	tk, err := d.hmac(sha1.New, chp[:], entrySalt)
	if err != nil {
		return nil, err
	}
	k2, err := d.hmac(sha1.New, chp[:], append(tk, entrySalt...))
	if err != nil {
		return nil, err
	}

	ret, err := d.tripleDesDecrypt(append(k1, k2...)[:24], k2[len(k2)-8:], cipherText)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (d *decrypt) hmac(hash func() hash.Hash, key, from []byte) ([]byte, error) {
	h := hmac.New(hash, key)
	_, err := h.Write(from)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func (d *decrypt) tripleDesDecrypt(key, iv, cipherText []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	return origData, nil
}

func (d *decrypt) fromBase64(msg string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func (d *decrypt) decryptSecret(masterPwd []byte, p passwordItem) ([]byte, error) {
	asn := asnSecretWithSalt{}
	if _, err := asn1.Unmarshal(p.SecretASN1, &asn); err != nil {
		return nil, err
	}

	if asn.ObjectSeq.ObjectIdentifier.String() == pbeWithSha1AndTripleDesCbc {
		res, err := d.decryptPbeWithSha1andTripleDesCbc(p.GlobalSalt, asn.ObjectSeq.SaltSeq.Salt, masterPwd, asn.Data)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		return nil, ProtocolSupportErr
	}
}
