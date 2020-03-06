package userrepo

// Interface is the interface for user repository interactions
type Interface interface {
	get(string) (string, error)
}
