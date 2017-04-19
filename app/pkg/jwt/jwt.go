package jwt

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	ajwt "github.com/someone1/gcp-jwt-go"
	"golang.org/x/net/context"
)

const HeaderKey = "JWT-TOKEN"

const (
	noToken         = "JWT token is missing"
	wrongSignMethod = "Wrong signing method %s"
	invalidClaims   = "Invalid claims"
)

type User struct {
	Name  string
	Email string
	//Account string
	//Role    string // Owner, member
}

//type Account struct {
//	ID   string // UUID
//	Name string
//}

type claims struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Email string `json:"email"`
	// Account string `json:"account"`
}

func SetShortToken(ctx context.Context, w http.ResponseWriter, u User) error {
	c := claims{
		StandardClaims: jwt.StandardClaims{
			//ID: UUID, from outside
			IssuedAt:  time.Now().UTC().Unix(),
			ExpiresAt: time.Now().UTC().Add(30 * time.Minute).Unix(),
			Issuer:    "test",
		},
		Name:  u.Name,
		Email: u.Email,
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod("AppEngine"), c)
	token, err := t.SignedString(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	w.Header().Set(HeaderKey, token)
	return nil
}

func GetUser(ctx context.Context, r *http.Request) (User, error) {
	t := r.Header.Get(HeaderKey)
	if t == "" {
		return User{}, errors.New(noToken)
	}

	token, err := jwt.ParseWithClaims(t, &claims{}, keyFunc(ctx))
	if err != nil {
		return User{}, errors.WithStack(err)
	}

	if c, ok := token.Claims.(*claims); ok && token.Valid {
		u := User{
			Name:  c.Name,
			Email: c.Email,
		}
		return u, nil
	}

	return User{}, errors.New(invalidClaims)
}

func keyFunc(ctx context.Context) func(*jwt.Token) (interface{}, error) {

	return func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*ajwt.SigningMethodAppEngine); !ok {
			return nil, errors.Errorf(wrongSignMethod, t.Header["alg"])
		}
		return ctx, nil
	}
}

//type StandardClaims struct {
//	Audience  string `json:"aud,omitempty"`
//	ExpiresAt int64  `json:"exp,omitempty"`
//	Id        string `json:"jti,omitempty"`
//	IssuedAt  int64  `json:"iat,omitempty"`
//	Issuer    string `json:"iss,omitempty"`
//	NotBefore int64  `json:"nbf,omitempty"`
//	Subject   string `json:"sub,omitempty"`
//}
