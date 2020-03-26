package accountservice

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	account "github.com/themanciraptor/Backend-photagea/internal/account/model"
	accountrepo "github.com/themanciraptor/Backend-photagea/internal/account/repo"
)

// Interface is the service interface
type Interface interface {
	SignIn(ctx context.Context, email string, password string) (string, error)
	Verify(jwtstring string) (int64, error)
	Create(ctx context.Context, Email string, Password string) error
	Update(ctx context.Context, accountID int64, Email string, Password string) error
}

// Service implements account service interface
type Service struct {
	repo accountrepo.Interface
}

type accountClaims struct {
	AccountID int64 `json:"AccountID"`
	jwt.StandardClaims
}

// Initialize a new account Service
func Initialize(r accountrepo.Interface) Interface {
	return &Service{repo: r}
}

// Create a account
func (a *Service) Create(ctx context.Context, Email string, Password string) error {
	return a.repo.Create(ctx, &account.Model{
		Password: Password,
		Email:    Email,
	})
}

// Update a account
func (a *Service) Update(ctx context.Context, accountID int64, Email string, Password string) error {
	return a.repo.Update(ctx, &account.Model{
		AccountID: accountID,
		Email:     Email,
		Password:  Password,
	})
}

// SignIn creates a JWT and returns it, returns error if password incorrect
func (a *Service) SignIn(ctx context.Context, email string, password string) (string, error) {
	acc, err := a.repo.Get(ctx, email, password)
	if err != nil {
		return "", err
	}

	claims := accountClaims{
		AccountID: acc.AccountID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 1500,
			Issuer:    "photagea.com",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secureSecretText"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Verify checks that a token is valid and returns the accoundID attached to the jwt
func (a *Service) Verify(jwtstring string) (int64, error) {
	j := strings.SplitAfter(jwtstring, "Bearer ")
	token, err := jwt.ParseWithClaims(
		j[0],
		&accountClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secureSecretText"), nil
		},
	)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*accountClaims)
	if !ok {
		return 0, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return 0, errors.New("JWT is expired")
	}

	return claims.AccountID, nil
}
