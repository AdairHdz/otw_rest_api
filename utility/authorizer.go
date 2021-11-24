package utility

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"crypto/rsa"	
	"io/ioutil"		
)

type CustomClaims struct {
	UserID string
	UserType int
	EmailAddress string
	jwt.RegisteredClaims
}

var (	
	privateKey *rsa.PrivateKey	
	publicKey *rsa.PublicKey
)

func init() {
	privateKeyBytes, err := ioutil.ReadFile("./keys/app.rsa")

	if err != nil {
		panic(err.Error())
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)

	if err != nil {
		panic(err.Error())
	}

	publicKeyBytes, err := ioutil.ReadFile("./keys/app.rsa.pub")

	if err != nil {
		panic(err.Error())
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)

	if err != nil {
		panic(err.Error())
	}
}

func SignString(userID string, userType int, emailAddress string, duration time.Time) (string, error) {
	claims := CustomClaims{
		userID,
		userType,
		emailAddress,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(duration),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "OTW-Auth-Server",
			Subject: userID,
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	
	signedString, err := token.SignedString(privateKey)		

	return signedString, err
}

func ValidateSignedString(signedString string) bool {
	token, err := jwt.ParseWithClaims(signedString, &CustomClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return false
	}
	
	if _, ok := token.Claims.(*CustomClaims); !ok || token.Valid {
		return false
	}

	return true
}

