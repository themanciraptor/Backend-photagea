// Interface for the account API
type Interface interface {
	SignIn(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
}

// accountAPI is the API for account related requests
type AccountAPI struct {
	accountService accountservice.Interface
}

func (u *AccountAPI) SignIn(http.ResponseWriter, *http.Request) {
	
}

type registerRequest struct {
	Email string
	Password string
}

func (u *AccountAPI) Register(http.ResponseWriter, *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	c := registerRequest{}

	err := d.Decode(&c)
	if err != nil {
		log.Printf("Unable to read request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}