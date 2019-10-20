package httpauth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/function61/gokit/csrf"
	"github.com/patrickmn/go-cache"
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
	publicKey *ecdsa.PublicKey
	// use caching for JWT validation, since at least ECDSA is really expensive, at least
	// when running on a Raspberry Pi Zero W each request takes seconds
	authCache *cache.Cache
}

func NewEcJwtAuthenticator(validatorPublicKey []byte) (HttpRequestAuthenticator, error) {
	publicKey, err := jwt.ParseECPublicKeyFromPEM(validatorPublicKey)
	if err != nil {
		return nil, err
	}

	return &jwtAuthenticator{
		publicKey: publicKey,
		// defaultExpiration doesn't matter, because we'll always push to cache with explicit TTLs
		authCache: cache.New(5*time.Minute, 10*time.Minute),
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

	if err := csrf.Validate(r); err != nil {
		return nil, err
	}

	return &UserDetails{
		Id: claims.Subject,
	}, nil
}

func (j *jwtAuthenticator) getValidatedClaims(jwtString string) (*jwt.StandardClaims, error) {
	cachedClaims, isCached := j.authCache.Get(jwtString)
	if isCached {
		return cachedClaims.(*jwt.StandardClaims), nil
	}

	token, err := jwt.ParseWithClaims(jwtString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return j.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// cache the claims until expiration, if we have expiration and is in the future
	if claims.ExpiresAt != 0 {
		untilExpiration := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())

		if untilExpiration > 0 {
			j.authCache.Set(jwtString, claims, untilExpiration)
		}
	}

	return claims, nil
}
