package utility

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	REFRESH = 1
	EPHIMERAL = 2
)

type CustomClaims struct {
	UserID string
	SpecificID string
	UserType int
	EmailAddress string
	jwt.RegisteredClaims
	JWTType int
}

var (	
	privateKey *rsa.PrivateKey	
	PublicKey *rsa.PublicKey
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

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)

	if err != nil {
		panic(err.Error())
	}
}

func SignString(userID string, specificID string, userType int, emailAddress string, duration time.Time, jwtType int) (string, error) {
	claims := CustomClaims{
		userID,
		specificID,
		userType,
		emailAddress,			
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(duration),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "OTW-Auth-Server",
			Subject: userID,
		},
		jwtType,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	
	signedString, err := token.SignedString(privateKey)		

	return signedString, err
}

func ValidateSignedString(signedString string) bool {
	token, err := jwt.ParseWithClaims(signedString, &CustomClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})	

	if err != nil {
		return false
	}
	
	if _, ok := token.Claims.(*CustomClaims); !ok || !token.Valid {
		return false
	}

	return true
}

func ExtractCustomClaims(signedString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(signedString, &CustomClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return PublicKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid signed string")
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid signed string")
	}	

	return claims, nil
}

