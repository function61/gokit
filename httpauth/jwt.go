package httpauth

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/function61/gokit/csrf"
	"github.com/patrickmn/go-cache"
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
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.StandardClaims{
		Subject:   userDetails.Id,
		ExpiresAt: now.Add(time.Hour * 24).Unix(),
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

func (j *jwtAuthenticator) Authenticate(r *http.Request) (*UserDetails, error) {
	// grab JWT either from:
	// 1) bearer token OR
	// 2) cookie
	jwtString := func() string {
		// first check if we have an authorization header
		authorizationHeader := r.Header.Get("Authorization")

		if strings.HasPrefix(authorizationHeader, "Bearer ") {
			return authorizationHeader[len("Bearer "):]
		}

		authCookie, err := r.Cookie(loginCookieName)
		if err != nil {
			return ""
		}

		return authCookie.Value
	}()

	if jwtString == "" {
		return nil, fmt.Errorf("auth: either specify '%s' cookie or 'Authorization' header", loginCookieName)
	}

	return j.AuthenticateJwtString(jwtString)
}

func (j *jwtAuthenticator) AuthenticateJwtString(jwtString string) (*UserDetails, error) {
	claims, err := j.getValidatedClaims(jwtString)
	if err != nil {
		return nil, fmt.Errorf("JWT authentication: %w", err)
	}

	return &UserDetails{
		Id: claims.Subject,
	}, nil
}

func (j *jwtAuthenticator) AuthenticateWithCsrfProtection(r *http.Request) (*UserDetails, error) {
	if err := csrf.Validate(r); err != nil {
		return nil, err
	}

	return j.Authenticate(r)
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
		untilExpiration := time.Until(time.Unix(claims.ExpiresAt, 0))

		if untilExpiration > 0 {
			j.authCache.Set(jwtString, claims, untilExpiration)
		}
	}

	return claims, nil
}
