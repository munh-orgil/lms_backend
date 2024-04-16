package session

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	private *rsa.PrivateKey
	pub     *rsa.PublicKey
}

var Jwt *JWT

func Write() {
	g := JWT{}
	g.private, _ = rsa.GenerateKey(rand.Reader, 2048)
	g.pub = &g.private.PublicKey

	var privateKeyBytes = x509.MarshalPKCS1PrivateKey(g.private)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}

	_ = pem.Encode(privatePem, privateKeyBlock)

	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(g.pub)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}
	_ = pem.Encode(publicPem, publicKeyBlock)
}

func (j *JWT) GenerateToken(claims *Token, expiresMinut time.Duration) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		Subject:   "lms token",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * expiresMinut)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.private)
}

func (j *JWT) ReadToken(inToken string) (*Token, error) {
	if j == nil {
		return nil, fmt.Errorf("does not have jwt")
	}
	var t Token
	if _, err := jwt.ParseWithClaims(inToken, &t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.pub, nil
	}); err != nil {
		return nil, err
	}
	return &t, nil
}

type ResToken struct {
	Token string `json:"token"`
}

func GetToken(tokenInfo *Token) (string, error) {
	token, err := Jwt.GenerateToken(tokenInfo, time.Duration(1440*30))
	if err != nil {
		return "", err
	}
	return token, nil
}
