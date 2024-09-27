package jwtFunc

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var (
	publicKey *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

func init (){
	keyPrivate, err := os.ReadFile("app.rsa")
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(keyPrivate)

	PriKey , err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	privateKey = PriKey.(*rsa.PrivateKey)


	keyPublic, err := os.ReadFile("app.rsa.pub")
	if err != nil {
		log.Fatal(err)
	}
	blockPub, _ := pem.Decode(keyPublic)
	
	PubKey, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	publicKey = PubKey.(*rsa.PublicKey)
}

func GenerateToken (id interface{}, w http.ResponseWriter){
	claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 4).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwtToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	cookies := &http.Cookie{
		Name: "jwt_token",
		Value: jwtToken,
		Path:     "/",
        HttpOnly: true,
        Secure:   false,
        MaxAge:   60 * 60 * 24 * 4,
	}
	http.SetCookie(w, cookies)
}

// func VerifyToken()