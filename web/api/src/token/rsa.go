package token

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
)

func GenerateRSAKeys(bitSize int) (priv, pub, kid string) {
	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	//publicKey := key.PublicKey

	//saveGobKey("private.key", key)
	//savePEMKey("private.pem", key)

	//saveGobKey("public.key", publicKey)
	//savePublicPEMKey("public.pem", publicKey)
	asn1Bytes, err := asn1.Marshal(key.PublicKey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	//pemfile, err := os.Create(fileName)
	//checkError(err)
	//defer pemfile.Close()

	pub = string(pem.EncodeToMemory(pemkey))
	//checkError(err)

	//asn1Bytes, _ := asn1.Marshal(publicKey)
	kid = fmt.Sprintf("%x", md5.Sum(asn1Bytes))
	//tokenString, err := token.SignedString(x509.MarshalPKCS1PrivateKey(key))
	priv = string(x509.MarshalPKCS1PrivateKey(key))

	return priv, pub, kid
}
