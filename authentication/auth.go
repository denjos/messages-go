package authentication

import (
	"crypto/rsa"
	"github.com/denjos/curso/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"time"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("error el lectura private.rsa")
	}
	publicBytes, err := ioutil.ReadFile("public.rsa.pub")
	if err != nil {
		log.Fatal("error el lectura public.rsa.pub")
	}
	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("error el parse privatekey")
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("error el parse publicKey")
	}
}
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "oscar proj",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Fatal("no se firmo el token", err)
	}
	return result
}
