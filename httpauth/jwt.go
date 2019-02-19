package httpauth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/function61/gokit/csrf"
	"net/http"
	"time"
)

type jwtSigner struct {
	privKey *ecdsa.PrivateKey
}

func NewEcJwtSigner(privateKey []byte) (Signer, error) {
	privKey, err := jwt.ParseECPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	return &jwtSigner{
		privKey: privKey,
	}, nil
}

func (j *jwtSigner) Sign(userDetails UserDetails, now time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
		"sub": userDetails.Id,
		"exp": now.Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(j.privKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

type jwtAuthenticator struct {
	pkey *ecdsa.PublicKey
}

func NewEcJwtAuthenticator(validatorPublicKey []byte) (HttpRequestAuthenticator, error) {
	pkey, err := jwt.ParseECPublicKeyFromPEM(validatorPublicKey)
	if err != nil {
		return nil, err
	}

	return &jwtAuthenticator{
		pkey: pkey,
	}, nil
}

func (j *jwtAuthenticator) AuthenticateWithCsrfProtection(r *http.Request) (*UserDetails, error) {
	authCookie, err := r.Cookie(loginCookieName)
	if err != nil {
		return nil, errors.New("auth: cookie " + loginCookieName + " missing")
	}

	claims, err := j.getValidatedClaims(authCookie.Value)
	if err != nil {
		return nil, err
	}

	if err := csrf.ValidateCookie(r); err != nil {
		return nil, err
	}

	return &UserDetails{
		Id: claims["sub"].(string),
	}, nil
}

func (j *jwtAuthenticator) getValidatedClaims(jwtString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return j.pkey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
