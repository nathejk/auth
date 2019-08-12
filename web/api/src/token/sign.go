package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	activePrivateKey  string
	activeFingerprint string
	publicKeys        map[string]string
)

func Sign(claims map[string]string, priv string, kid string) string {

	mapClaims := jwt.MapClaims{
		"iss": "auth",
		"iat": time.Now().Unix(),
	}
	for key, value := range claims {
		mapClaims[key] = value
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	token.Header["kid"] = kid
	tokenString, _ := token.SignedString([]byte(priv))

	return tokenString
}

func pubkeyHandler(w http.ResponseWriter, r *http.Request) {
	kid := strings.Replace(r.URL.Path, "/", "", 1)

	fmt.Fprintf(w, publicKeys[kid])
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}
func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		//		os.Exit(1)
	}
}
