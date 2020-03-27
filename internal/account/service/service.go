package accountservice

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	account "github.com/themanciraptor/Backend-photagea/internal/account/model"
	accountrepo "github.com/themanciraptor/Backend-photagea/internal/account/repo"
)

// Interface is the service interface
type Interface interface {
	SignIn(ctx context.Context, email string, password string) (string, error)
	Verify(r *http.Request) (int64, error)
	Create(ctx context.Context, Email string, Password string) error
	Update(ctx context.Context, accountID int64, Email string, Password string) error
	RefreshToken(r *http.Request) (string, error)
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

	return genToken(acc.AccountID)
}

// Verify checks that a token is valid and returns the accoundID attached to the jwt
func (a *Service) Verify(r *http.Request) (int64, error) {
	j := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	claims, err := getClaims(j)
	if err != nil {
		return 0, err
	}

	return claims.AccountID, nil
}

// RefreshToken issues a fresh token
func (a *Service) RefreshToken(r *http.Request) (string, error) {
	j := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	claims, err := getClaims(j)
	if err != nil {
		return "", err
	}

	return genToken(claims.AccountID)
}

func genToken(accountID int64) (string, error) {
	claims := accountClaims{
		AccountID: accountID,
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

func getClaims(token string) (*accountClaims, error) {
	t, err := jwt.ParseWithClaims(
		token,
		&accountClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secureSecretText"), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*accountClaims)
	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}
