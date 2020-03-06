package userservice

import user "github.com/themanciraptor/Backend-photagea/internal/user/model"

// Interface is the service interface
type Interface interface {
	Get(string) (user.Model, error)
	Create(string, string, string) (user.Model, error)
	Update(string, string, string) (user.Model, error)
}
